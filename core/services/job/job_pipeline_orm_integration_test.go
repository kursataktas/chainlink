package job_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"

	"github.com/smartcontractkit/chainlink/core/internal/cltest"
	"github.com/smartcontractkit/chainlink/core/internal/testutils/evmtest"
	"github.com/smartcontractkit/chainlink/core/internal/testutils/pgtest"
	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/services/job"
	"github.com/smartcontractkit/chainlink/core/services/pipeline"
	"github.com/smartcontractkit/chainlink/core/services/postgres"
	"github.com/smartcontractkit/chainlink/core/store/models"
)

func clearJobsDb(t *testing.T, db *gorm.DB) {
	t.Helper()
	// Ordering matters here to avoid deadlocks
	err := db.Exec(`TRUNCATE flux_monitor_round_stats_v2, jobs, pipeline_runs, pipeline_specs, pipeline_task_runs CASCADE`).Error
	require.NoError(t, err)
}

func TestPipelineORM_Integration(t *testing.T) {
	const DotStr = `
        // data source 1
        ds1          [type=bridge name=voter_turnout];
        ds1_parse    [type=jsonparse path="one,two"];
        ds1_multiply [type=multiply times=1.23];

        // data source 2
        ds2          [type=http method=GET url="https://chain.link/voter_turnout/USA-2020" requestData=<{"hi": "hello"}>];
        ds2_parse    [type=jsonparse path="three,four"];
        ds2_multiply [type=multiply times=4.56];

        ds1 -> ds1_parse -> ds1_multiply -> answer1;
        ds2 -> ds2_parse -> ds2_multiply -> answer1;

        answer1 [type=median                      index=0];
        answer2 [type=bridge name=election_winner index=1];
    `

	config := cltest.NewTestGeneralConfig(t)
	gdb := pgtest.NewGormDB(t)
	db := postgres.UnwrapGormDB(gdb)
	config.SetDB(gdb)
	config.Overrides.SetDefaultHTTPTimeout(30 * time.Millisecond)
	config.Overrides.DefaultMaxHTTPAttempts = null.IntFrom(1)
	keyStore := cltest.NewKeyStore(t, db)
	ethKeyStore := keyStore.Eth()

	_, transmitterAddress := cltest.MustInsertRandomKey(t, ethKeyStore)
	keyStore.OCR().Add(cltest.DefaultOCRKey)

	var specID int32

	answer1 := &pipeline.MedianTask{
		AllowedFaults: "",
	}
	answer2 := &pipeline.BridgeTask{
		Name: "election_winner",
	}
	ds1_multiply := &pipeline.MultiplyTask{
		Times: "1.23",
	}
	ds1_parse := &pipeline.JSONParseTask{
		Path: "one,two",
	}
	ds1 := &pipeline.BridgeTask{
		Name: "voter_turnout",
	}
	ds2_multiply := &pipeline.MultiplyTask{
		Times: "4.56",
	}
	ds2_parse := &pipeline.JSONParseTask{
		Path: "three,four",
	}
	ds2 := &pipeline.HTTPTask{
		URL:         "https://chain.link/voter_turnout/USA-2020",
		Method:      "GET",
		RequestData: `{"hi": "hello"}`,
	}

	answer1.BaseTask = pipeline.NewBaseTask(6, "answer1", []pipeline.Task{ds1_multiply, ds2_multiply}, nil, 0)
	answer2.BaseTask = pipeline.NewBaseTask(7, "answer2", nil, nil, 1)
	ds1_multiply.BaseTask = pipeline.NewBaseTask(2, "ds1_multiply", []pipeline.Task{ds1_parse}, []pipeline.Task{answer1}, 0)
	ds2_multiply.BaseTask = pipeline.NewBaseTask(5, "ds2_multiply", []pipeline.Task{ds2_parse}, []pipeline.Task{answer1}, 0)
	ds1_parse.BaseTask = pipeline.NewBaseTask(1, "ds1_parse", []pipeline.Task{ds1}, []pipeline.Task{ds1_multiply}, 0)
	ds2_parse.BaseTask = pipeline.NewBaseTask(4, "ds2_parse", []pipeline.Task{ds2}, []pipeline.Task{ds2_multiply}, 0)
	ds1.BaseTask = pipeline.NewBaseTask(0, "ds1", nil, []pipeline.Task{ds1_parse}, 0)
	ds2.BaseTask = pipeline.NewBaseTask(3, "ds2", nil, []pipeline.Task{ds2_parse}, 0)
	expectedTasks := []pipeline.Task{ds1, ds1_parse, ds1_multiply, ds2, ds2_parse, ds2_multiply, answer1, answer2}
	_, bridge := cltest.MustCreateBridge(t, db, cltest.BridgeOpts{})
	_, bridge2 := cltest.MustCreateBridge(t, db, cltest.BridgeOpts{})

	t.Run("creates task DAGs", func(t *testing.T) {
		clearJobsDb(t, gdb)

		orm := pipeline.NewORM(db)

		p, err := pipeline.Parse(DotStr)
		require.NoError(t, err)

		specID, err = orm.CreateSpec(*p, models.Interval(0))
		require.NoError(t, err)

		var specs []pipeline.Spec
		err = gdb.Find(&specs).Error
		require.NoError(t, err)
		require.Len(t, specs, 1)
		require.Equal(t, specID, specs[0].ID)
		require.Equal(t, DotStr, specs[0].DotDagSource)

		require.NoError(t, gdb.Exec(`DELETE FROM pipeline_specs`).Error)
	})

	t.Run("creates runs", func(t *testing.T) {
		lggr := logger.TestLogger(t)
		clearJobsDb(t, gdb)
		orm := pipeline.NewORM(db)
		cc := evmtest.NewChainSet(t, evmtest.TestChainOpts{Client: cltest.NewEthClientMockWithDefaultChain(t), DB: gdb, GeneralConfig: config})
		runner := pipeline.NewRunner(orm, config, cc, nil, nil, lggr)
		defer runner.Close()
		jobORM := job.NewTestORM(t, db, cc, orm, keyStore)

		dbSpec := makeVoterTurnoutOCRJobSpec(t, gdb, transmitterAddress, bridge.Name.String(), bridge2.Name.String())

		// Need a job in order to create a run
		err := jobORM.CreateJob(dbSpec)
		require.NoError(t, err)

		var pipelineSpecs []pipeline.Spec
		err = gdb.Find(&pipelineSpecs).Error
		require.NoError(t, err)
		require.Len(t, pipelineSpecs, 1)
		require.Equal(t, dbSpec.PipelineSpecID, pipelineSpecs[0].ID)
		pipelineSpecID := pipelineSpecs[0].ID

		// Create the run
		runID, _, err := runner.ExecuteAndInsertFinishedRun(context.Background(), pipelineSpecs[0], pipeline.NewVarsFrom(nil), lggr, true)
		require.NoError(t, err)

		// Check the DB for the pipeline.Run
		var pipelineRuns []pipeline.Run
		err = gdb.Where("id = ?", runID).Find(&pipelineRuns).Error
		require.NoError(t, err)
		require.Len(t, pipelineRuns, 1)
		require.Equal(t, pipelineSpecID, pipelineRuns[0].PipelineSpecID)
		require.Equal(t, runID, pipelineRuns[0].ID)

		// Check the DB for the pipeline.TaskRuns
		var taskRuns []pipeline.TaskRun
		err = gdb.Where("pipeline_run_id = ?", runID).Find(&taskRuns).Error
		require.NoError(t, err)
		require.Len(t, taskRuns, len(expectedTasks))

		for _, taskRun := range taskRuns {
			assert.Equal(t, runID, taskRun.PipelineRunID)
			assert.False(t, taskRun.Output.Valid)
			assert.False(t, taskRun.Error.IsZero())
		}
	})
}
