package ccipdeployment

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"go.uber.org/zap/zapcore"

	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	jobv1 "github.com/smartcontractkit/chainlink-protos/job-distributor/v1/job"
	"github.com/smartcontractkit/chainlink/integration-tests/deployment"
	"github.com/smartcontractkit/chainlink/integration-tests/deployment/devenv"
	"github.com/smartcontractkit/chainlink/integration-tests/deployment/memory"
)

//
// CCIPEnvBuilder is used to configure and build a CCIPEnvironment for testing.
// CCIPEnvironment helps manage the environment (start, deploy, etc) for CCIP tests.
//

/*
// types:
// - Source
// - Destination
// - All
{
  "fRoleDON": 5,
  "numNodes": 16,
  "chains":
	"chain1": {
      "fChain": 1,
      "type": "Source",
      "nodes": [0, 1, 2, 3]
    },
    "chain2":{
      "fChain": 1,
      "type": "Source",
      "nodes": [4, 5, 6, 7]
    },
    "chain3": {
      "fChain": 1,
      "type": "Source",
      "nodes": [4, 5, 6, 7]
    },
    "destChain": {
      "type": "Destination",
      "fChain": 2,
      "nodes": [8, 9, 10, 11, 12, 13, 14, 15]
    },
  "nodes": {
    "0": {
      "some_default_override": 99
    }
  }
}
*/

type CCIPChain struct {
	Name   string
	Home   bool
	FChain int
	Nodes  []int
}

type CCIPEnvConfig struct {
	numNodes int
	fRoleDON int
	chains   map[string]CCIPChain
	// TODO: node overrides?
}

type CCIPEnvOption func(e *CCIPEnvironment)

func WithInMemoryBackend() CCIPEnvOption {
	return func(e *CCIPEnvironment) {
		e.deployType = "memory"
	}
}

func WithJSONConfig(jsonConfig string) CCIPEnvOption {
	return func(e *CCIPEnvironment) {
		// TODO: json.Unmarshal(jsonConfig, e)
	}
}

func WithNumNodes(n int) CCIPEnvOption {
	return func(e *CCIPEnvironment) {
		e.config.numNodes = n
	}
}

func WithFRoleDON(n int) CCIPEnvOption {
	return func(e *CCIPEnvironment) {
		e.config.fRoleDON = n
	}
}

func WithChain(name string, chain CCIPChain) CCIPEnvOption {
	return func(e *CCIPEnvironment) {
		if e.config.chains == nil {
			e.config.chains = make(map[string]CCIPChain)
		}
		e.config.chains[name] = chain
	}
}

func WithLogger(lggr logger.Logger) CCIPEnvOption {
	return func(e *CCIPEnvironment) {
		e.lggr = lggr
	}
}

// TODO: launch and configure the postgres DB automatically?
func WithPostgres(connectionString string) CCIPEnvOption {
	return func(e *CCIPEnvironment) {
		panic("not implemented")
	}
}

type ChangeSetHandler func(
	ab deployment.AddressBook, env deployment.Environment, c DeployCCIPContractConfig,
) (deployment.ChangesetOutput, error)

// WithChangeSet injects a CCIP config into the environment.
func WithChangeSet(changeset ChangeSetHandler) CCIPEnvOption {
	return func(e *CCIPEnvironment) {
		if e != nil {
			panic("multiple changesets are not supported")
		}
		e.changeset = changeset
	}
}

func NewCCIPEnvironment(opts ...CCIPEnvOption) *CCIPEnvironment {
	de := &CCIPEnvironment{}
	for _, opt := range opts {
		opt(de)
	}
	return de
}

type ccipDeployedEnv struct {
	DeployedEnv
	nodes map[string]memory.Node
}

type CCIPEnvironment struct {
	deployType      string // "memory", "local", ...
	config          CCIPEnvConfig
	lggr            logger.Logger
	ccipDeployedEnv ccipDeployedEnv

	// intermediate stuff
	changeset             ChangeSetHandler
	changesetOutput       deployment.ChangesetOutput
	replayBlocks          map[uint64]uint64
	homeChainSel, feedSel uint64
	shutdownHook          func() error
}

func makeDeployedMemoryEnv(
	ctx context.Context, lggr logger.Logger, t *testing.T, cfg CCIPEnvConfig,
) (func() error /* shutdownHook */, ccipDeployedEnv, error) {
	chains := memory.NewMemoryChains(t, len(cfg.chains))
	homeChainSel, feedSel := allocateCCIPChainSelectors(chains)
	replayBlocks, err := LatestBlocksByChain(ctx, chains)
	if err != nil {
		return nil, ccipDeployedEnv{}, fmt.Errorf("failed to get latest blocks: %w", err)
	}

	ab := deployment.NewMemoryAddressBook()
	feeTokenContracts, crConfig := DeployTestContracts(t, lggr, ab, homeChainSel, feedSel, chains)
	nodes := memory.NewNodes(t, zapcore.InfoLevel, chains, cfg.numNodes, 1, crConfig)
	for n, node := range nodes {
		if err = node.App.Start(ctx); err != nil {
			return nil, ccipDeployedEnv{}, fmt.Errorf("failed to start node [%s]: %w", n, err)
		}
	}

	env := memory.NewMemoryEnvironmentFromChainsNodes(t, lggr, chains, nodes)
	cde := ccipDeployedEnv{
		nodes: nodes,
		DeployedEnv: DeployedEnv{
			Ab:                ab,
			Env:               env,
			HomeChainSel:      homeChainSel,
			FeedChainSel:      feedSel,
			ReplayBlocks:      replayBlocks,
			FeeTokenContracts: feeTokenContracts,
		},
	}

	shutdownHook := func() error {
		var errs []error
		for n, node := range nodes {
			if err = node.App.Stop(); err != nil {
				errs = append(errs, fmt.Errorf("failed to stop node [%s]: %w", n, err))
			}
		}
		return errors.Join(errs...)
	}

	return shutdownHook, cde, nil
}

func runChangeset(
	changeset ChangeSetHandler, tenv DeployedEnv, mcmsConfig MCMSConfig,
) (deployment.ChangesetOutput, error) {
	state, err := LoadOnchainState(tenv.Env, tenv.Ab)
	if err != nil {
		return deployment.ChangesetOutput{}, fmt.Errorf("failed to load onchain state: %w", err)
	}
	if state.Chains[tenv.HomeChainSel].LinkToken == nil {
		return deployment.ChangesetOutput{}, fmt.Errorf("link token is missing")
	}

	feeds := state.Chains[tenv.FeedChainSel].USDFeeds
	tokenConfig := NewTokenConfig()
	tokenConfig.UpsertTokenInfo(LinkSymbol,
		pluginconfig.TokenInfo{
			AggregatorAddress: cciptypes.UnknownEncodedAddress(feeds[LinkSymbol].Address().String()),
			Decimals:          LinkDecimals,
			DeviationPPB:      cciptypes.NewBigIntFromInt64(1e9),
		},
	)
	tokenConfig.UpsertTokenInfo(WethSymbol,
		pluginconfig.TokenInfo{
			AggregatorAddress: cciptypes.UnknownEncodedAddress(feeds[WethSymbol].Address().String()),
			Decimals:          WethDecimals,
			DeviationPPB:      cciptypes.NewBigIntFromInt64(1e9),
		},
	)

	// TODO: deploy hooks
	output, err := changeset(tenv.Ab, tenv.Env, DeployCCIPContractConfig{
		HomeChainSel:       tenv.HomeChainSel,
		FeedChainSel:       tenv.FeedChainSel,
		ChainsToDeploy:     tenv.Env.AllChainSelectors(),
		TokenConfig:        tokenConfig,
		MCMSConfig:         mcmsConfig,
		CapabilityRegistry: state.Chains[tenv.HomeChainSel].CapabilityRegistry.Address(),
		FeeTokenContracts:  tenv.FeeTokenContracts,
		OCRSecrets:         deployment.XXXGenerateTestOCRSecrets(),
	})
	if err != nil {
		return deployment.ChangesetOutput{}, fmt.Errorf("failed to deploy CCIP contracts: %w", err)
	}

	return output, nil
}

func deployJobs(ctx context.Context, tenv DeployedEnv, output deployment.ChangesetOutput) error {
	// Ensure capreg logs are up to date.
	//ReplayLogs(t, tenv.Env.Offchain, tenv.ReplayBlocks)
	{
		var err error

		// TODO: use ReplayLogs once testing.T dependency is removed.
		switch oc := tenv.Env.Offchain.(type) {
		case *memory.JobClient:
			err = oc.ReplayLogs(tenv.ReplayBlocks)
		case *devenv.JobDistributor:
			err = oc.ReplayLogs(tenv.ReplayBlocks)
		default:
			err = fmt.Errorf("unsupported offchain client type %T", oc)
		}

		if err != nil {
			return fmt.Errorf("failed to replay logs: %w", err)
		}
	}

	// Apply the jobs.
	for nodeID, jobs := range output.JobSpecs {
		for _, job := range jobs {
			// Note these auto-accept
			_, err := tenv.Env.Offchain.ProposeJob(ctx,
				&jobv1.ProposeJobRequest{
					NodeId: nodeID,
					Spec:   job,
				})
			if err != nil {
				return fmt.Errorf("failed to propose job: %w", err)
			}
		}
	}

	return nil
}

func wireCustomTopology(tenv DeployedEnv, config CCIPEnvConfig) error {
	// Get new state after migration.
	state, err := LoadOnchainState(tenv.Env, tenv.Ab)
	if err != nil {
		return fmt.Errorf("failed to load onchain state: %w", err)
	}

	// Add all lanes
	err = AddLanesForAll(tenv.Env, state)
	if err != nil {
		return fmt.Errorf("failed to add lanes: %w", err)
	}

	return nil
}

// Start the environment. Returns a shutdown hook.
func (e *CCIPEnvironment) Start(ctx context.Context, t *testing.T) (func() error /* shutdown */, error) {
	numChains := len(e.config.chains)

	if e.lggr == nil {
		l := logger.Test(t)
		e.lggr = l
	}

	//
	// Validate configuration and initialize defaults.
	//
	if numChains < 2 {
		return nil, fmt.Errorf("numChains must be at least 2 for home and feed chains")
	}

	if e.config.numNodes < 4 {
		return nil, fmt.Errorf("numNodes must be at least 4")
	}

	if e.config.fRoleDON == 0 {
		e.lggr.Infow("fRoleDON not set, using default value of numNodes/3", "numNodes", e.config.numNodes)
		e.config.fRoleDON = e.config.numNodes / 3
	}

	//
	// Deploy and initialize a blank environment (chains and nodes).
	//
	{
		switch e.deployType {
		case "memory":
			shutdownHook, de, err := makeDeployedMemoryEnv(ctx, e.lggr, t, e.config)
			if err != nil {
				return nil, err
			}
			e.ccipDeployedEnv = de
			e.shutdownHook = shutdownHook
		default:
			return nil, fmt.Errorf("unsupported deployment type %s", e.deployType)
		}
	}

	//
	// Configure the environment
	//
	{
		mcmsConfig := NewTestMCMSConfig(t, e.ccipDeployedEnv.Env)
		output, err := runChangeset(e.changeset, e.ccipDeployedEnv.DeployedEnv, mcmsConfig)
		if err != nil {
			return nil, fmt.Errorf("failed to run changeset: %w", err)
		}
		e.changesetOutput = output
	}

	//
	// Deploy jobs to nodes.
	//
	{
		err := deployJobs(ctx, e.ccipDeployedEnv.DeployedEnv, e.changesetOutput)
		if err != nil {
			return nil, fmt.Errorf("failed to deploy and wire CCIP: %w", err)
		}
	}

	//
	// Configure CCIP topology.
	//
	{
		err := wireCustomTopology(e.ccipDeployedEnv.DeployedEnv, e.config)
		if err != nil {
			return nil, fmt.Errorf("failed to wire custom topology: %w", err)
		}
	}

	return e.shutdownHook, nil
}
