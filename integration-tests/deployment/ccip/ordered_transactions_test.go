package ccipdeployment

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/ccip-owner-contracts/pkg/proposal/timelock"
	"github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/router"
	"math/big"
	"testing"

	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"

	"github.com/smartcontractkit/chainlink/integration-tests/deployment"

	jobv1 "github.com/smartcontractkit/chainlink-protos/job-distributor/v1/job"

	"github.com/smartcontractkit/chainlink-testing-framework/lib/utils/testcontext"

	"github.com/smartcontractkit/chainlink/v2/core/logger"

	"github.com/stretchr/testify/require"
)

func TestTransactionOrdering(t *testing.T) {
	lggr := logger.TestLogger(t)
	ctx := testcontext.Get(t)
	tenv := NewMemoryEnvironment(t, lggr, 2, 4)
	e := tenv.Env
	ab := tenv.Ab

	state, err := LoadOnchainState(tenv.Env, tenv.Ab)
	require.NoError(t, err)
	require.NotNil(t, state.Chains[tenv.HomeChainSel].LinkToken)

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

	err = DeployCCIPContracts(e, ab, DeployCCIPContractConfig{
		HomeChainSel:       tenv.HomeChainSel,
		FeedChainSel:       tenv.FeedChainSel,
		ChainsToDeploy:     tenv.Env.AllChainSelectors(),
		TokenConfig:        tokenConfig,
		MCMSConfig:         NewTestMCMSConfig(t, e),
		CapabilityRegistry: state.Chains[tenv.HomeChainSel].CapabilityRegistry.Address(),
		FeeTokenContracts:  tenv.FeeTokenContracts,
		OCRSecrets:         deployment.XXXGenerateTestOCRSecrets(),
	})
	require.NoError(t, err)
	js, err := NewCCIPJobSpecs(e.NodeIDs, e.Offchain)
	require.NoError(t, err)
	output, err := deployment.ChangesetOutput{
		Proposals:   []timelock.MCMSWithTimelockProposal{},
		AddressBook: ab,
		// Mapping of which nodes get which jobs.
		JobSpecs: js,
	}, nil
	require.NoError(t, err)
	// Get new state after migration.
	state, err = LoadOnchainState(e, tenv.Ab)
	require.NoError(t, err)

	tokenPoolData, err := DeployAndRegisterTokenPools(e, tenv.Env.AllChainSelectors(), state)
	require.NoError(t, err)

	state, err = LoadOnchainState(e, tenv.Ab)
	require.NoError(t, err)

	// Ensure capreg logs are up to date.
	ReplayLogs(t, e.Offchain, tenv.ReplayBlocks)

	// Apply the jobs.
	for nodeID, jobs := range output.JobSpecs {
		for _, job := range jobs {
			// Note these auto-accept
			_, err := e.Offchain.ProposeJob(ctx,
				&jobv1.ProposeJobRequest{
					NodeId: nodeID,
					Spec:   job,
				})
			require.NoError(t, err)
		}
	}

	// Add lane
	require.NoError(t, AddLanesForAll(e, state))

	src := e.Chains[tenv.HomeChainSel]
	dst := e.Chains[tenv.FeedChainSel]
	// Try out multiple requests, with combinations of messages and tokens
	requests := []router.ClientEVM2AnyMessage{
		{
			Receiver:     common.LeftPadBytes(state.Chains[dst.Selector].Receiver.Address().Bytes(), 32),
			Data:         []byte("Hello Chain"),
			TokenAmounts: nil,
			FeeToken:     common.HexToAddress("0x0"),
			ExtraArgs:    nil,
		},
		{
			Receiver: common.LeftPadBytes(state.Chains[dst.Selector].Receiver.Address().Bytes(), 32),
			Data:     []byte("Hello Chain, again"),
			TokenAmounts: []router.ClientEVMTokenAmount{
				{
					tokenPoolData[src.Selector].Token.Address(),
					big.NewInt(10),
				},
			},
			FeeToken:  common.HexToAddress("0x0"),
			ExtraArgs: nil,
		},
		{
			Receiver: common.LeftPadBytes(state.Chains[dst.Selector].Receiver.Address().Bytes(), 32),
			Data:     nil,
			TokenAmounts: []router.ClientEVMTokenAmount{
				{
					tokenPoolData[src.Selector].Token.Address(),
					big.NewInt(100),
				},
			},
			FeeToken:  common.HexToAddress("0x0"),
			ExtraArgs: nil,
		},
	}

	latesthdr, err := dst.Client.HeaderByNumber(testcontext.Get(t), nil)
	require.NoError(t, err)
	block := latesthdr.Number.Uint64()

	seqNum := SendMessage(t, e, state, src.Selector, dst.Selector, false, requests[0])
	seqNum2 := SendMessage(t, e, state, src.Selector, dst.Selector, false, requests[1])
	seqNum3 := SendMessage(t, e, state, src.Selector, dst.Selector, false, requests[2])

	//var seqNums []uint64
	//for range len(requests) {
	//	seqNums = append(seqNums, SendRequest(t, e, state, src.Selector, dst.Selector, false))
	//}

	// expect that operations arrive in the correct order
	//for _, seqNum := range seqNums {
	// Wait for all commit reports to land.
	//ConfirmCommitForAllWithExpectedSeqNums(t, e, state, map[uint64]uint64{dst.Selector: seqNum}, map[uint64]*uint64{dst.Selector: &startBlock})
	////
	//// Wait for all exec reports to land
	//ConfirmExecWithSeqNrForAll(t, e, state, map[uint64]uint64{dst.Selector: seqNum}, map[uint64]*uint64{dst.Selector: &startBlock})
	//}
	// Wait for all commit reports to land.
	//ConfirmCommitForAllWithExpectedSeqNums(t, e, state, expectedSeqNum, startBlocks)

	err = ConfirmCommitWithExpectedSeqNumRange(t, src, dst, state.Chains[dst.Selector].OffRamp, &block, ccipocr3.NewSeqNumRange(ccipocr3.SeqNum(seqNum), ccipocr3.SeqNum(seqNum3)))
	require.NoError(t, err)

	//err = ConfirmCommitWithExpectedSeqNumRange(t, src, dst, state.Chains[dst.Selector].OffRamp, &block, ccipocr3.NewSeqNumRange(ccipocr3.SeqNum(seqNum2), ccipocr3.SeqNum(seqNum2)))
	//require.NoError(t, err)

	// Wait for all exec reports to land
	//ConfirmExecWithSeqNrForAll(t, e, state, expectedSeqNum, startBlocks)
	err = ConfirmExecWithSeqNr(t, src, dst, state.Chains[dst.Selector].OffRamp, &block, seqNum)
	require.NoError(t, err)

	latesthdr, err = dst.Client.HeaderByNumber(testcontext.Get(t), nil)
	require.NoError(t, err)
	block = latesthdr.Number.Uint64()

	err = ConfirmExecWithSeqNr(t, src, dst, state.Chains[dst.Selector].OffRamp, &block, seqNum2)
	require.NoError(t, err)

	latesthdr, err = dst.Client.HeaderByNumber(testcontext.Get(t), nil)
	require.NoError(t, err)
	block = latesthdr.Number.Uint64()

	err = ConfirmExecWithSeqNr(t, src, dst, state.Chains[dst.Selector].OffRamp, &block, seqNum3)
	require.NoError(t, err)
}
