package ccipdeployment

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/smartcontractkit/ccip-owner-contracts/pkg/proposal/mcms"
	"github.com/smartcontractkit/ccip-owner-contracts/pkg/proposal/timelock"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink/integration-tests/deployment"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/offramp"

	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/fee_quoter"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/onramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/router"
)

func AddLane(e deployment.Environment, state CCIPOnChainState, from, to uint64) error {
	// TODO: Batch
	tx, err := state.Chains[from].Router.ApplyRampUpdates(e.Chains[from].DeployerKey, []router.RouterOnRamp{
		{
			DestChainSelector: to,
			OnRamp:            state.Chains[from].OnRamp.Address(),
		},
	}, []router.RouterOffRamp{}, []router.RouterOffRamp{})
	if _, err := deployment.ConfirmIfNoError(e.Chains[from], tx, err); err != nil {
		return err
	}
	tx, err = state.Chains[from].OnRamp.ApplyDestChainConfigUpdates(e.Chains[from].DeployerKey,
		[]onramp.OnRampDestChainConfigArgs{
			{
				DestChainSelector: to,
				Router:            state.Chains[from].Router.Address(),
			},
		})
	if _, err := deployment.ConfirmIfNoError(e.Chains[from], tx, err); err != nil {
		return err
	}

	_, err = state.Chains[from].FeeQuoter.UpdatePrices(
		e.Chains[from].DeployerKey, fee_quoter.InternalPriceUpdates{
			TokenPriceUpdates: []fee_quoter.InternalTokenPriceUpdate{
				{
					SourceToken: state.Chains[from].LinkToken.Address(),
					UsdPerToken: deployment.E18Mult(20),
				},
				{
					SourceToken: state.Chains[from].Weth9.Address(),
					UsdPerToken: deployment.E18Mult(4000),
				},
			},
			GasPriceUpdates: []fee_quoter.InternalGasPriceUpdate{
				{
					DestChainSelector: to,
					UsdPerUnitGas:     big.NewInt(2e12),
				},
			}})
	if _, err := deployment.ConfirmIfNoError(e.Chains[from], tx, err); err != nil {
		return err
	}

	// Enable dest in fee quoter
	tx, err = state.Chains[from].FeeQuoter.ApplyDestChainConfigUpdates(e.Chains[from].DeployerKey,
		[]fee_quoter.FeeQuoterDestChainConfigArgs{
			{
				DestChainSelector: to,
				DestChainConfig:   defaultFeeQuoterDestChainConfig(),
			},
		})
	if _, err := deployment.ConfirmIfNoError(e.Chains[from], tx, err); err != nil {
		return err
	}

	tx, err = state.Chains[to].OffRamp.ApplySourceChainConfigUpdates(e.Chains[to].DeployerKey,
		[]offramp.OffRampSourceChainConfigArgs{
			{
				Router:              state.Chains[to].Router.Address(),
				SourceChainSelector: from,
				IsEnabled:           true,
				OnRamp:              common.LeftPadBytes(state.Chains[from].OnRamp.Address().Bytes(), 32),
			},
		})
	if _, err := deployment.ConfirmIfNoError(e.Chains[to], tx, err); err != nil {
		return err
	}
	tx, err = state.Chains[to].Router.ApplyRampUpdates(e.Chains[to].DeployerKey, []router.RouterOnRamp{}, []router.RouterOffRamp{}, []router.RouterOffRamp{
		{
			SourceChainSelector: from,
			OffRamp:             state.Chains[to].OffRamp.Address(),
		},
	})
	_, err = deployment.ConfirmIfNoError(e.Chains[to], tx, err)
	return err
}

func defaultFeeQuoterDestChainConfig() fee_quoter.FeeQuoterDestChainConfig {
	// https://github.com/smartcontractkit/ccip/blob/c4856b64bd766f1ddbaf5d13b42d3c4b12efde3a/contracts/src/v0.8/ccip/libraries/Internal.sol#L337-L337
	/*
		```Solidity
			// bytes4(keccak256("CCIP ChainFamilySelector EVM"))
			bytes4 public constant CHAIN_FAMILY_SELECTOR_EVM = 0x2812d52c;
		```
	*/
	evmFamilySelector, _ := hex.DecodeString("2812d52c")
	return fee_quoter.FeeQuoterDestChainConfig{
		IsEnabled:                         true,
		MaxNumberOfTokensPerMsg:           10,
		MaxDataBytes:                      256,
		MaxPerMsgGasLimit:                 3_000_000,
		DestGasOverhead:                   50_000,
		DefaultTokenFeeUSDCents:           1,
		DestGasPerPayloadByte:             10,
		DestDataAvailabilityOverheadGas:   0,
		DestGasPerDataAvailabilityByte:    100,
		DestDataAvailabilityMultiplierBps: 1,
		DefaultTokenDestGasOverhead:       125_000,
		//DefaultTokenDestBytesOverhead:     32,
		DefaultTxGasLimit:      200_000,
		GasMultiplierWeiPerEth: 1,
		NetworkFeeUSDCents:     1,
		ChainFamilySelector:    [4]byte(evmFamilySelector),
	}
}

func generateAddLaneProposal(e deployment.Environment, state CCIPOnChainState, from, to uint64) (*timelock.MCMSWithTimelockProposal, error) {
	var batches []timelock.BatchChainOperation
	metaDataPerChain := make(map[mcms.ChainIdentifier]mcms.ChainMetadata)
	timelockAddresses := make(map[mcms.ChainIdentifier]common.Address)
	var sourceBatch McmsAccumulator
	var destBatch McmsAccumulator

	_, err := state.Chains[from].Router.ApplyRampUpdates(sourceBatch.TxnAppender(), []router.RouterOnRamp{
		{
			DestChainSelector: to,
			OnRamp:            state.Chains[from].OnRamp.Address(),
		},
	}, []router.RouterOffRamp{}, []router.RouterOffRamp{})
	if err != nil {
		return nil, err
	}

	_, err = state.Chains[from].OnRamp.ApplyDestChainConfigUpdates(sourceBatch.TxnAppender(),
		[]onramp.OnRampDestChainConfigArgs{
			{
				DestChainSelector: to,
				Router:            state.Chains[from].Router.Address(),
			},
		})
	if err != nil {
		return nil, err
	}

	_, err = state.Chains[from].FeeQuoter.UpdatePrices(
		sourceBatch.TxnAppender(), fee_quoter.InternalPriceUpdates{
			TokenPriceUpdates: []fee_quoter.InternalTokenPriceUpdate{
				{
					SourceToken: state.Chains[from].LinkToken.Address(),
					UsdPerToken: deployment.E18Mult(20),
				},
				{
					SourceToken: state.Chains[from].Weth9.Address(),
					UsdPerToken: deployment.E18Mult(4000),
				},
			},
			GasPriceUpdates: []fee_quoter.InternalGasPriceUpdate{
				{
					DestChainSelector: to,
					UsdPerUnitGas:     big.NewInt(2e12),
				},
			}})
	if err != nil {
		return nil, err
	}

	// Enable dest in fee quoter
	_, err = state.Chains[from].FeeQuoter.ApplyDestChainConfigUpdates(sourceBatch.TxnAppender(),
		[]fee_quoter.FeeQuoterDestChainConfigArgs{
			{
				DestChainSelector: to,
				DestChainConfig:   defaultFeeQuoterDestChainConfig(),
			},
		})
	if err != nil {
		return nil, err
	}

	_, err = state.Chains[to].OffRamp.ApplySourceChainConfigUpdates(destBatch.TxnAppender(),
		[]offramp.OffRampSourceChainConfigArgs{
			{
				Router:              state.Chains[to].Router.Address(),
				SourceChainSelector: from,
				IsEnabled:           true,
				OnRamp:              common.LeftPadBytes(state.Chains[from].OnRamp.Address().Bytes(), 32),
			},
		})
	if err != nil {
		return nil, err
	}
	_, err = state.Chains[to].Router.ApplyRampUpdates(destBatch.TxnAppender(), []router.RouterOnRamp{}, []router.RouterOffRamp{}, []router.RouterOffRamp{
		{
			SourceChainSelector: from,
			OffRamp:             state.Chains[to].OffRamp.Address(),
		},
	})
	if err != nil {
		return nil, err
	}

	batches = append(batches, timelock.BatchChainOperation{
		ChainIdentifier: mcms.ChainIdentifier(from),
		Batch:           sourceBatch.B,
	})
	batches = append(batches, timelock.BatchChainOperation{
		ChainIdentifier: mcms.ChainIdentifier(to),
		Batch:           destBatch.B,
	})

	return timelock.NewMCMSWithTimelockProposal(
		"1",
		2004259681, // TODO
		[]mcms.Signature{},
		false,
		metaDataPerChain,
		timelockAddresses,
		fmt.Sprintf("Add source (%d) and destination (%d)", from, to),
		batches,
		timelock.Schedule, "0s")
}

type McmsAccumulator struct {
	B []mcms.Operation
}

func (m *McmsAccumulator) TxnAppender() *bind.TransactOpts {
	return &bind.TransactOpts{
		Signer: func(address common.Address, transaction *types.Transaction) (*types.Transaction, error) {
			m.B = append(m.B, mcms.Operation{
				To:    *transaction.To(),
				Data:  transaction.Data(),
				Value: big.NewInt(0),
			})
			return transaction, nil
		},
		From:     common.HexToAddress("0x0"),
		NoSend:   true,
		GasLimit: 1_000_000,
	}
}
