package oraclecreator

import (
	"context"
	"encoding/hex"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	ragep2ptypes "github.com/smartcontractkit/libocr/ragep2p/types"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/libocr/commontypes"
	libocr3 "github.com/smartcontractkit/libocr/offchainreporting2plus"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"

	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	ccipreaderpkg "github.com/smartcontractkit/chainlink-ccip/pkg/reader"

	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn"
	"github.com/smartcontractkit/chainlink-common/pkg/types"

	"github.com/smartcontractkit/chainlink/v2/core/capabilities/ccip/ocrimpls"
	cctypes "github.com/smartcontractkit/chainlink/v2/core/capabilities/ccip/types"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/chaintype"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocrcommon"
	"github.com/smartcontractkit/chainlink/v2/core/services/synchronization"
	"github.com/smartcontractkit/chainlink/v2/core/services/telemetry"
)

var _ cctypes.OracleCreator = &bootstrapOracleCreator{}

type bootstrapOracleCreator struct {
	peerWrapper           *ocrcommon.SingletonPeerWrapper
	bootstrapperLocators  []commontypes.BootstrapperLocator
	db                    ocr3types.Database
	monitoringEndpointGen telemetry.MonitoringEndpointGenerator
	lggr                  logger.Logger
	contractReader        types.ContractReader
}

func NewBootstrapOracleCreator(
	peerWrapper *ocrcommon.SingletonPeerWrapper,
	bootstrapperLocators []commontypes.BootstrapperLocator,
	db ocr3types.Database,
	monitoringEndpointGen telemetry.MonitoringEndpointGenerator,
	lggr logger.Logger,
	contractReader types.ContractReader,
) cctypes.OracleCreator {
	return &bootstrapOracleCreator{
		peerWrapper:           peerWrapper,
		bootstrapperLocators:  bootstrapperLocators,
		db:                    db,
		monitoringEndpointGen: monitoringEndpointGen,
		lggr:                  lggr,
		contractReader:        contractReader,
	}
}

// Type implements types.OracleCreator.
func (i *bootstrapOracleCreator) Type() cctypes.OracleType {
	return cctypes.OracleTypeBootstrap
}

// Create implements types.OracleCreator.
func (i *bootstrapOracleCreator) Create(_ uint32, config cctypes.OCR3ConfigWithMeta) (cctypes.CCIPOracle, error) {
	// Assuming that the chain selector is referring to an evm chain for now.
	// TODO: add an api that returns chain family.
	// NOTE: this doesn't really matter for the bootstrap node, it doesn't do anything on-chain.
	// Its for the monitoring endpoint generation below.
	chainID, err := chainsel.ChainIdFromSelector(uint64(config.Config.ChainSelector))
	if err != nil {
		return nil, fmt.Errorf("failed to get chain ID from selector: %w", err)
	}

	ctx := context.Background()
	rmnHomeReader, err := i.getRmnHomeReader(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("failed to get RMNHome reader: %w", err)
	}
	if err = rmnHomeReader.Start(ctx); err != nil {
		return nil, fmt.Errorf("failed to start RMNHome reader: %w", err)
	}

	destChainFamily := chaintype.EVM
	destRelayID := types.NewRelayID(string(destChainFamily), fmt.Sprintf("%d", chainID))

	oraclePeerIDs := make([]ragep2ptypes.PeerID, 0, len(config.Config.Nodes))
	for _, n := range config.Config.Nodes {
		oraclePeerIDs = append(oraclePeerIDs, n.P2pID)
	}

	pgd := newPeerGroupDialer(
		i.lggr,
		i.peerWrapper.PeerGroupFactory,
		i.bootstrapperLocators,
		oraclePeerIDs,
		config.ConfigDigest,
		/* todo: should also provide rmn home reader */
	)
	pgd.Start()

	bootstrapperArgs := libocr3.BootstrapperArgs{
		BootstrapperFactory:   i.peerWrapper.Peer2,
		V2Bootstrappers:       i.bootstrapperLocators,
		ContractConfigTracker: ocrimpls.NewConfigTracker(config),
		Database:              i.db,
		LocalConfig:           defaultLocalConfig(),
		Logger: ocrcommon.NewOCRWrapper(
			i.lggr.
				Named("CCIPBootstrap").
				Named(destRelayID.String()).
				Named(config.Config.ChainSelector.String()).
				Named(hexutil.Encode(config.Config.OfframpAddress)),
			false, /* traceLogging */
			func(ctx context.Context, msg string) {}),
		MonitoringEndpoint: i.monitoringEndpointGen.GenMonitoringEndpoint(
			string(destChainFamily),
			destRelayID.ChainID,
			hexutil.Encode(config.Config.OfframpAddress),
			synchronization.OCR3CCIPBootstrap,
		),
		OffchainConfigDigester: ocrimpls.NewConfigDigester(config.ConfigDigest),
	}
	bootstrapper, err := libocr3.NewBootstrapper(bootstrapperArgs)
	if err != nil {
		return nil, err
	}

	bootstrapperWithCustomClose := newWrappedOracle(
		bootstrapper,
		[]io.Closer{pgd},
	)

	return bootstrapperWithCustomClose, nil
}

func (i *bootstrapOracleCreator) getRmnHomeReader(ctx context.Context, config cctypes.OCR3ConfigWithMeta) (ccipreaderpkg.RMNHome, error) {
	rmnHomeBoundContract := types.BoundContract{
		Address: "0x" + hex.EncodeToString(config.Config.RmnHomeAddress),
		Name:    consts.ContractNameRMNHome,
	}

	if err1 := i.contractReader.Bind(ctx, []types.BoundContract{rmnHomeBoundContract}); err1 != nil {
		return nil, fmt.Errorf("failed to bind RMNHome contract: %w", err1)
	}
	rmnHomeReader := ccipreaderpkg.NewRMNHomePoller(
		i.contractReader,
		rmnHomeBoundContract,
		i.lggr,
		5*time.Second,
	)
	return rmnHomeReader, nil
}

// peerGroupDialer keeps watching for config changes and calls NewPeerGroup when needed.
// Required for managing RMN related peer group connections.
type peerGroupDialer struct {
	lggr logger.Logger

	peerGroupFactory rmn.PeerGroupFactory

	// common oracle config
	bootstrapLocators  []commontypes.BootstrapperLocator
	oraclePeerIDs      []ragep2ptypes.PeerID
	commitConfigDigest [32]byte

	activePeerGroups []rmn.PeerGroup

	syncInterval time.Duration

	mu *sync.Mutex
}

func newPeerGroupDialer(
	lggr logger.Logger,
	peerGroupFactory rmn.PeerGroupFactory,
	bootstrapLocators []commontypes.BootstrapperLocator,
	oraclePeerIDs []ragep2ptypes.PeerID,
	commitConfigDigest [32]byte,
) *peerGroupDialer {
	return &peerGroupDialer{
		lggr: lggr,

		peerGroupFactory: peerGroupFactory,

		bootstrapLocators:  bootstrapLocators,
		oraclePeerIDs:      oraclePeerIDs,
		commitConfigDigest: commitConfigDigest,

		activePeerGroups: []rmn.PeerGroup{},

		syncInterval: time.Minute, // todo: make it configurable

		mu: &sync.Mutex{},
	}
}

func (d *peerGroupDialer) Start() {
	go func() {
		d.sync()

		syncTicker := time.NewTicker(d.syncInterval)
		for {
			select {
			case <-syncTicker.C:
				d.sync()
			}
		}
	}()
}

func (d *peerGroupDialer) Close() error {
	d.mu.Lock()
	defer d.mu.Unlock()

	d.closeExistingPeerGroups()
	return nil
}

func (d *peerGroupDialer) sync() {
	d.mu.Lock()
	defer d.mu.Unlock()

	if !d.shouldSync() {
		return
	}

	d.closeExistingPeerGroups()
	d.createNewPeerGroups()
}

func (d *peerGroupDialer) shouldSync() bool {
	if len(d.activePeerGroups) == 0 {
		return true
	}

	// todo: if config has changed return true

	return false
}

func (d *peerGroupDialer) closeExistingPeerGroups() {
	for _, pg := range d.activePeerGroups {
		if err := pg.Close(); err != nil {
			d.lggr.Warnw("failed to close peer group", "err", err)
			continue
		}
	}

	d.activePeerGroups = []rmn.PeerGroup{}
}

func (d *peerGroupDialer) createNewPeerGroups() {
	/*
		Requires:
		- commit config digest - ok
		- rmn home config digest (maximum 2) - we need reader.RMNHome which should be updated with a new method
		- oracle peer ids - ok
		- rmn peer ids - we need reader.RMNHome
		- bootstrappers - ok
	*/

	// make calls to get rmn home config / etc...
	// d.peerGroupFactory.NewPeerGroup(...)
}
