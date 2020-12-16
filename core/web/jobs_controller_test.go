package web_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/smartcontractkit/chainlink/core/services/eth"

	"github.com/pelletier/go-toml"

	"github.com/smartcontractkit/chainlink/core/internal/cltest"
	"github.com/smartcontractkit/chainlink/core/services"
	"github.com/smartcontractkit/chainlink/core/services/job"
	"github.com/smartcontractkit/chainlink/core/services/offchainreporting"
	"github.com/smartcontractkit/chainlink/core/store/models"
	"github.com/smartcontractkit/chainlink/core/web"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/guregu/null.v4"
)

func TestJobsController_Create_ValidationFailure(t *testing.T) {
	var (
		contractAddress    = cltest.NewEIP55Address()
		monitoringEndpoint = "chain.link:101"
	)

	var tt = []struct {
		name        string
		pid         models.PeerID
		kb          models.Sha256Hash
		taExists    bool
		expectedErr error
	}{
		{
			name:        "invalid keybundle",
			pid:         models.PeerID(cltest.DefaultP2PPeerID),
			kb:          models.Sha256Hash(cltest.Random32Byte()),
			taExists:    true,
			expectedErr: job.ErrNoSuchKeyBundle,
		},
		{
			name:        "invalid peerID",
			pid:         models.PeerID(cltest.NonExistentP2PPeerID),
			kb:          cltest.DefaultOCRKeyBundleIDSha256,
			taExists:    true,
			expectedErr: job.ErrNoSuchPeerID,
		},
		{
			name:        "invalid transmitter address",
			pid:         models.PeerID(cltest.DefaultP2PPeerID),
			kb:          cltest.DefaultOCRKeyBundleIDSha256,
			taExists:    false,
			expectedErr: job.ErrNoSuchTransmitterAddress,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ta, client, cleanup := setupJobsControllerTests(t)
			defer cleanup()

			var address models.EIP55Address
			if tc.taExists {
				key := cltest.MustInsertRandomKey(t, ta.Store.DB)
				address = key.Address
			} else {
				address = cltest.NewEIP55Address()
			}

			sp := cltest.MinimalOCRNonBootstrapSpec(contractAddress, address, tc.pid, monitoringEndpoint, tc.kb)
			body, _ := json.Marshal(models.CreateJobSpecRequest{
				TOML: sp,
			})
			resp, cleanup := client.Post("/v2/jobs", bytes.NewReader(body))
			defer cleanup()
			assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
			b, err := ioutil.ReadAll(resp.Body)
			require.NoError(t, err)
			assert.Contains(t, string(b), tc.expectedErr.Error())
		})
	}
}

func TestJobsController_Create_HappyPath_OffchainReportingSpec(t *testing.T) {
	app, client, cleanup := setupJobsControllerTests(t)
	defer cleanup()

	toml := string(cltest.MustReadFile(t, "testdata/oracle-spec.toml"))
	toml = strings.Replace(toml, "0x27548a32b9aD5D64c5945EaE9Da5337bc3169D15", app.Key.Address.Hex(), 1)
	body, _ := json.Marshal(models.CreateJobSpecRequest{
		TOML: toml,
	})
	response, cleanup := client.Post("/v2/jobs", bytes.NewReader(body))
	defer cleanup()
	require.Equal(t, http.StatusOK, response.StatusCode)

	jb := job.JobSpecV2{}
	require.NoError(t, app.Store.DB.Preload("OffchainreportingOracleSpec").First(&jb).Error)

	ocrJobSpec := job.JobSpecV2{}
	err := web.ParseJSONAPIResponse(cltest.ParseResponseBody(t, response), &ocrJobSpec)
	assert.NoError(t, err)

	assert.Equal(t, "web oracle spec", jb.Name.ValueOrZero())
	assert.Equal(t, jb.OffchainreportingOracleSpec.P2PPeerID, ocrJobSpec.OffchainreportingOracleSpec.P2PPeerID)
	assert.Equal(t, jb.OffchainreportingOracleSpec.P2PBootstrapPeers, ocrJobSpec.OffchainreportingOracleSpec.P2PBootstrapPeers)
	assert.Equal(t, jb.OffchainreportingOracleSpec.IsBootstrapPeer, ocrJobSpec.OffchainreportingOracleSpec.IsBootstrapPeer)
	assert.Equal(t, jb.OffchainreportingOracleSpec.EncryptedOCRKeyBundleID, ocrJobSpec.OffchainreportingOracleSpec.EncryptedOCRKeyBundleID)
	assert.Equal(t, jb.OffchainreportingOracleSpec.MonitoringEndpoint, ocrJobSpec.OffchainreportingOracleSpec.MonitoringEndpoint)
	assert.Equal(t, jb.OffchainreportingOracleSpec.TransmitterAddress, ocrJobSpec.OffchainreportingOracleSpec.TransmitterAddress)
	assert.Equal(t, jb.OffchainreportingOracleSpec.ObservationTimeout, ocrJobSpec.OffchainreportingOracleSpec.ObservationTimeout)
	assert.Equal(t, jb.OffchainreportingOracleSpec.BlockchainTimeout, ocrJobSpec.OffchainreportingOracleSpec.BlockchainTimeout)
	assert.Equal(t, jb.OffchainreportingOracleSpec.ContractConfigTrackerSubscribeInterval, ocrJobSpec.OffchainreportingOracleSpec.ContractConfigTrackerSubscribeInterval)
	assert.Equal(t, jb.OffchainreportingOracleSpec.ContractConfigTrackerSubscribeInterval, ocrJobSpec.OffchainreportingOracleSpec.ContractConfigTrackerSubscribeInterval)
	assert.Equal(t, jb.OffchainreportingOracleSpec.ContractConfigConfirmations, ocrJobSpec.OffchainreportingOracleSpec.ContractConfigConfirmations)
	assert.NotNil(t, ocrJobSpec.PipelineSpec.DotDagSource)

	// Sanity check to make sure it inserted correctly
	require.Equal(t, models.EIP55Address("0x613a38AC1659769640aaE063C651F48E0250454C"), jb.OffchainreportingOracleSpec.ContractAddress)
}

func TestJobsController_Create_HappyPath_EthRequestEventSpec(t *testing.T) {
	rpcClient, gethClient, _, assertMocksCalled := cltest.NewEthMocksWithStartupAssertions(t)
	defer assertMocksCalled()
	app, cleanup := cltest.NewApplicationWithKey(t,
		eth.NewClientWith(rpcClient, gethClient),
	)
	defer cleanup()
	require.NoError(t, app.Start())
	gethClient.On("SubscribeFilterLogs", mock.Anything, mock.Anything, mock.Anything).Maybe().Return(cltest.EmptyMockSubscription(), nil)

	client := app.NewHTTPClient()

	body, _ := json.Marshal(models.CreateJobSpecRequest{
		TOML: string(cltest.MustReadFile(t, "testdata/eth-request-event-spec.toml")),
	})
	response, cleanup := client.Post("/v2/jobs", bytes.NewReader(body))
	defer cleanup()
	require.Equal(t, http.StatusOK, response.StatusCode)

	jb := job.JobSpecV2{}
	require.NoError(t, app.Store.DB.Preload("EthRequestEventSpec").First(&jb).Error)

	jobSpec := job.JobSpecV2{}
	err := web.ParseJSONAPIResponse(cltest.ParseResponseBody(t, response), &jobSpec)
	assert.NoError(t, err)

	assert.Equal(t, "example eth request event spec", jb.Name.ValueOrZero())
	assert.NotNil(t, jobSpec.PipelineSpec.DotDagSource)

	// Sanity check to make sure it inserted correctly
	require.Equal(t, models.EIP55Address("0x613a38AC1659769640aaE063C651F48E0250454C"), jb.EthRequestEventSpec.ContractAddress)
}

func TestJobsController_Index_HappyPath(t *testing.T) {
	client, cleanup, ocrJobSpecFromFile, _, ereJobSpecFromFile, _ := setupJobSpecsControllerTestsWithJobs(t)
	defer cleanup()

	response, cleanup := client.Get("/v2/jobs")
	defer cleanup()
	cltest.AssertServerResponse(t, response, http.StatusOK)

	jobSpecs := []job.JobSpecV2{}
	err := web.ParseJSONAPIResponse(cltest.ParseResponseBody(t, response), &jobSpecs)
	assert.NoError(t, err)

	require.Len(t, jobSpecs, 2)

	runOCRJobSpecAssertions(t, ocrJobSpecFromFile, jobSpecs[0])
	runEthRequestEventJobSpecAssertions(t, ereJobSpecFromFile, jobSpecs[1])
}

func TestJobsController_Show_HappyPath(t *testing.T) {
	client, cleanup, ocrJobSpecFromFile, jobID, ereJobSpecFromFile, jobID2 := setupJobSpecsControllerTestsWithJobs(t)
	defer cleanup()

	response, cleanup := client.Get("/v2/jobs/" + fmt.Sprintf("%v", jobID))
	defer cleanup()
	cltest.AssertServerResponse(t, response, http.StatusOK)

	ocrJobSpec := job.JobSpecV2{}
	err := web.ParseJSONAPIResponse(cltest.ParseResponseBody(t, response), &ocrJobSpec)
	assert.NoError(t, err)

	runOCRJobSpecAssertions(t, ocrJobSpecFromFile, ocrJobSpec)

	response, cleanup = client.Get("/v2/jobs/" + fmt.Sprintf("%v", jobID2))
	defer cleanup()
	cltest.AssertServerResponse(t, response, http.StatusOK)

	ereJobSpec := job.JobSpecV2{}
	err = web.ParseJSONAPIResponse(cltest.ParseResponseBody(t, response), &ereJobSpec)
	assert.NoError(t, err)

	runEthRequestEventJobSpecAssertions(t, ereJobSpecFromFile, ereJobSpec)
}

func TestJobsController_Show_InvalidID(t *testing.T) {
	client, cleanup, _, _, _, _ := setupJobSpecsControllerTestsWithJobs(t)
	defer cleanup()

	response, cleanup := client.Get("/v2/jobs/uuidLikeString")
	defer cleanup()
	cltest.AssertServerResponse(t, response, http.StatusUnprocessableEntity)
}

func TestJobsController_Show_NonExistentID(t *testing.T) {
	client, cleanup, _, _, _, _ := setupJobSpecsControllerTestsWithJobs(t)
	defer cleanup()

	response, cleanup := client.Get("/v2/jobs/999999999")
	defer cleanup()
	cltest.AssertServerResponse(t, response, http.StatusNotFound)
}

func runOCRJobSpecAssertions(t *testing.T, ocrJobSpecFromFile offchainreporting.OracleSpec, ocrJobSpecFromServer job.JobSpecV2) {
	assert.Equal(t, ocrJobSpecFromFile.ContractAddress, ocrJobSpecFromServer.OffchainreportingOracleSpec.ContractAddress)
	assert.Equal(t, ocrJobSpecFromFile.P2PPeerID, ocrJobSpecFromServer.OffchainreportingOracleSpec.P2PPeerID)
	assert.Equal(t, ocrJobSpecFromFile.P2PBootstrapPeers, ocrJobSpecFromServer.OffchainreportingOracleSpec.P2PBootstrapPeers)
	assert.Equal(t, ocrJobSpecFromFile.IsBootstrapPeer, ocrJobSpecFromServer.OffchainreportingOracleSpec.IsBootstrapPeer)
	assert.Equal(t, ocrJobSpecFromFile.EncryptedOCRKeyBundleID, ocrJobSpecFromServer.OffchainreportingOracleSpec.EncryptedOCRKeyBundleID)
	assert.Equal(t, ocrJobSpecFromFile.MonitoringEndpoint, ocrJobSpecFromServer.OffchainreportingOracleSpec.MonitoringEndpoint)
	assert.Equal(t, ocrJobSpecFromFile.TransmitterAddress, ocrJobSpecFromServer.OffchainreportingOracleSpec.TransmitterAddress)
	assert.Equal(t, ocrJobSpecFromFile.ObservationTimeout, ocrJobSpecFromServer.OffchainreportingOracleSpec.ObservationTimeout)
	assert.Equal(t, ocrJobSpecFromFile.BlockchainTimeout, ocrJobSpecFromServer.OffchainreportingOracleSpec.BlockchainTimeout)
	assert.Equal(t, ocrJobSpecFromFile.ContractConfigTrackerSubscribeInterval, ocrJobSpecFromServer.OffchainreportingOracleSpec.ContractConfigTrackerSubscribeInterval)
	assert.Equal(t, ocrJobSpecFromFile.ContractConfigTrackerSubscribeInterval, ocrJobSpecFromServer.OffchainreportingOracleSpec.ContractConfigTrackerSubscribeInterval)
	assert.Equal(t, ocrJobSpecFromFile.ContractConfigConfirmations, ocrJobSpecFromServer.OffchainreportingOracleSpec.ContractConfigConfirmations)
	assert.Equal(t, ocrJobSpecFromFile.Pipeline.DOTSource, ocrJobSpecFromServer.PipelineSpec.DotDagSource)

	// Check that create and update dates are non empty values.
	// Empty date value is "0001-01-01 00:00:00 +0000 UTC" so we are checking for the
	// millenia and century characters to be present
	assert.Contains(t, ocrJobSpecFromServer.OffchainreportingOracleSpec.CreatedAt.String(), "20")
	assert.Contains(t, ocrJobSpecFromServer.OffchainreportingOracleSpec.UpdatedAt.String(), "20")
}

func runEthRequestEventJobSpecAssertions(t *testing.T, ereJobSpecFromFile services.EthRequestEventSpec, ereJobSpecFromServer job.JobSpecV2) {
	assert.Equal(t, ereJobSpecFromFile.ContractAddress, ereJobSpecFromServer.EthRequestEventSpec.ContractAddress)
	assert.Equal(t, ereJobSpecFromFile.Pipeline.DOTSource, ereJobSpecFromServer.PipelineSpec.DotDagSource)
	// Check that create and update dates are non empty values.
	// Empty date value is "0001-01-01 00:00:00 +0000 UTC" so we are checking for the
	// millenia and century characters to be present
	assert.Contains(t, ereJobSpecFromServer.EthRequestEventSpec.CreatedAt.String(), "20")
	assert.Contains(t, ereJobSpecFromServer.EthRequestEventSpec.UpdatedAt.String(), "20")
}

func setupJobsControllerTests(t *testing.T) (*cltest.TestApplication, cltest.HTTPClientCleaner, func()) {
	t.Parallel()
	rpcClient, gethClient, _, assertMocksCalled := cltest.NewEthMocksWithStartupAssertions(t)
	defer assertMocksCalled()
	app, cleanup := cltest.NewApplicationWithKey(t,
		eth.NewClientWith(rpcClient, gethClient),
	)
	require.NoError(t, app.Start())

	client := app.NewHTTPClient()
	return app, client, cleanup
}

func setupJobSpecsControllerTestsWithJobs(t *testing.T) (cltest.HTTPClientCleaner, func(), offchainreporting.OracleSpec, int32, services.EthRequestEventSpec, int32) {
	t.Parallel()
	rpcClient, gethClient, _, assertMocksCalled := cltest.NewEthMocksWithStartupAssertions(t)
	defer assertMocksCalled()
	app, cleanup := cltest.NewApplicationWithKey(t,
		eth.NewClientWith(rpcClient, gethClient),
	)
	require.NoError(t, app.Start())

	client := app.NewHTTPClient()

	var ocrJobSpecFromFile offchainreporting.OracleSpec
	tree, err := toml.LoadFile("testdata/oracle-spec.toml")
	require.NoError(t, err)
	err = tree.Unmarshal(&ocrJobSpecFromFile)
	require.NoError(t, err)
	ocrJobSpecFromFile.TransmitterAddress = &app.Key.Address
	jobID, _ := app.AddJobV2(context.Background(), ocrJobSpecFromFile, null.String{})

	var ereJobSpecFromFile services.EthRequestEventSpec
	tree, err = toml.LoadFile("testdata/eth-request-event-spec.toml")
	require.NoError(t, err)
	err = tree.Unmarshal(&ereJobSpecFromFile)
	require.NoError(t, err)
	jobID2, _ := app.AddJobV2(context.Background(), ereJobSpecFromFile, null.String{})

	return client, cleanup, ocrJobSpecFromFile, jobID, ereJobSpecFromFile, jobID2
}
