package rollups

import (
	"context"
	"fmt"
	"math/big"
	"slices"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/smartcontractkit/chainlink-common/pkg/services"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets"
	evmconfig "github.com/smartcontractkit/chainlink/v2/core/chains/evm/config"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/config/chaintype"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/config/toml"
)

// L1Oracle provides interface for fetching L1-specific fee components if the chain is an L2.
// For example, on Optimistic Rollups, this oracle can return rollup-specific l1BaseFee
type L1Oracle interface {
	services.Service

	GasPrice(ctx context.Context) (*assets.Wei, error)
}

type l1OracleClient interface {
	CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error)
	BatchCallContext(ctx context.Context, b []rpc.BatchElem) error
}

// Client is interface of client connections for additional chains layers
type Client interface {
	SuggestGasPrice(ctx context.Context) (*big.Int, error)
	FeeHistory(ctx context.Context, blockCount uint64, rewardPercentiles []float64) (feeHistory *ethereum.FeeHistory, err error)
}

type priceEntry struct {
	price     *assets.Wei
	timestamp time.Time
}

const (
	// Interval at which to poll for L1BaseFee. A good starting point is the L1 block time.
	PollPeriod = 6 * time.Second
)

var supportedChainTypes = []chaintype.ChainType{chaintype.ChainArbitrum, chaintype.ChainOptimismBedrock, chaintype.ChainKroma, chaintype.ChainScroll, chaintype.ChainZkSync, chaintype.ChainMantle}

func IsRollupWithL1Support(chainType chaintype.ChainType) bool {
	return slices.Contains(supportedChainTypes, chainType)
}

func NewL1GasOracle(lggr logger.Logger, ethClient l1OracleClient, chainType chaintype.ChainType, daOracle evmconfig.DAOracle, clientsByChainID map[string]Client) (L1Oracle, error) {
	if !IsRollupWithL1Support(chainType) {
		return nil, nil
	}

	// TODO implementation to use the clientsByChainID should update the check accordingly, potentially return errors instead of logging
	if clientsByChainID != nil {
		if _, exist := clientsByChainID[daOracle.L1ChainID()]; !exist {
			lggr.Debugf("eth client for chainID %v should exist in clientsByChainID map", daOracle.L1ChainID())
		}
	} else {
		lggr.Debugf("clientsByChainID map is missing, expect L1 client with chainID %v to exist", daOracle.L1ChainID())
	}

	var l1Oracle L1Oracle
	var err error
	if daOracle != nil {
		switch daOracle.OracleType() {
		case toml.DAOracleOPStack:
			l1Oracle, err = NewOpStackL1GasOracle(lggr, ethClient, chainType, daOracle)
		case toml.DAOracleArbitrum:
			l1Oracle, err = NewArbitrumL1GasOracle(lggr, ethClient)
		case toml.DAOracleZKSync:
			l1Oracle = NewZkSyncL1GasOracle(lggr, ethClient)
		}
		if err != nil {
			return nil, fmt.Errorf("failed to initialize L1 oracle for chaintype %s: %w", chainType, err)
		}
		if l1Oracle != nil {
			return l1Oracle, nil
		}
	}

	// Going forward all configs should specify a DAOracle config. This is a fall back to maintain backwards compat.
	switch chainType {
	case chaintype.ChainArbitrum:
		l1Oracle, err = NewArbitrumL1GasOracle(lggr, ethClient)
	case chaintype.ChainZkSync:
		l1Oracle = NewZkSyncL1GasOracle(lggr, ethClient)
	default:
		return nil, fmt.Errorf("received unsupported chaintype %s", chainType)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to initialize L1 oracle for chaintype %s: %w", chainType, err)
	}
	return l1Oracle, nil
}
