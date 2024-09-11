package keystone_test

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/test-go/testify/require"
	"go.uber.org/zap/zapcore"

	chainsel "github.com/smartcontractkit/chain-selectors"

	"github.com/smartcontractkit/chainlink/integration-tests/deployment"
	"github.com/smartcontractkit/chainlink/integration-tests/deployment/clo"
	"github.com/smartcontractkit/chainlink/integration-tests/deployment/clo/models"
	"github.com/smartcontractkit/chainlink/integration-tests/deployment/keystone"
	"github.com/smartcontractkit/chainlink/integration-tests/deployment/memory"
	kcr "github.com/smartcontractkit/chainlink/v2/core/gethwrappers/keystone/generated/capabilities_registry"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
)

func TestDeploy(t *testing.T) {
	lggr := logger.TestLogger(t)
	t.Run("memory environment", func(t *testing.T) {
		var (
			testNops = []kcr.CapabilitiesRegistryNodeOperator{
				{
					Admin: common.HexToAddress("0x6CdfBF967A8ec4C29Fe26aF2a33Eb485d02f22D6"),
					Name:  "NOP_00",
				},
				{
					Admin: common.HexToAddress("0x6CdfBF967A8ec4C29Fe26aF2a33Eb485d02f2200"),
					Name:  "NOP_01",
				},
				{
					Admin: common.HexToAddress("0x11dfBF967A8ec4C29Fe26aF2a33Eb485d02f22D6"),
					Name:  "NOP_02",
				},
				{
					Admin: common.HexToAddress("0x6CdfBF967A8ec4C29Fe26aF2a33Eb485d02f2222"),
					Name:  "NOP_03",
				},
			}
			ocr3Config = keystone.OracleConfigSource{
				MaxFaultyOracles: 1,
			}
		)

		multDonCfg := memory.MemoryEnvironmentMultiDonConfig{
			Configs: make(map[string]memory.MemoryEnvironmentConfig),
		}
		wfEnvCfg := memory.MemoryEnvironmentConfig{
			Bootstraps: 1,
			Chains:     1,
			Nodes:      4,
		}
		multDonCfg.Configs[keystone.WFDonName] = wfEnvCfg

		targetEnvCfg := memory.MemoryEnvironmentConfig{
			Bootstraps: 1,
			Chains:     4,
			Nodes:      4,
		}
		multDonCfg.Configs[keystone.TargetDonName] = targetEnvCfg

		e := memory.NewMultiDonMemoryEnvironment(t, lggr, zapcore.InfoLevel, multDonCfg)

		var nodeToNop = make(map[string]kcr.CapabilitiesRegistryNodeOperator) //node -> nop
		// assign nops to nodes
		for _, env := range e.DonToEnv {
			for i, nodeID := range env.NodeIDs {
				idx := i % len(testNops)
				nop := testNops[idx]
				nodeToNop[nodeID] = nop
			}
		}

		var donsToDeploy = map[string][]kcr.CapabilitiesRegistryCapability{
			keystone.WFDonName:     []kcr.CapabilitiesRegistryCapability{keystone.OCR3Cap},
			keystone.TargetDonName: []kcr.CapabilitiesRegistryCapability{keystone.WriteChainCap},
		}

		ctx := context.Background()
		// Deploy all the Keystone contracts.
		homeChainSel := e.Get(keystone.WFDonName).AllChainSelectors()[0]
		deployReq := keystone.DeployRequest{
			RegistryChainSel:  homeChainSel,
			Menv:              e,
			DonToCapabilities: donsToDeploy,
			NodeIDToNop:       nodeToNop,
			OCR3Config:        &ocr3Config,
		}

		deployResp, err := keystone.Deploy(ctx, lggr, deployReq)
		require.NoError(t, err)
		ad := deployResp.Changeset.AddressBook
		addrs, err := ad.Addresses()
		require.NoError(t, err)
		lggr.Infow("Deployed Keystone contracts", "address book", addrs)

		// all contracts on home chain
		homeChainAddrs, err := ad.AddressesForChain(homeChainSel)
		require.NoError(t, err)
		require.Len(t, homeChainAddrs, 3)
		// only forwarder on non-home chain
		for _, chain := range e.Get(keystone.TargetDonName).AllChainSelectors() {
			chainAddrs, err := ad.AddressesForChain(chain)
			require.NoError(t, err)
			if chain != homeChainSel {
				require.Len(t, chainAddrs, 1)
			} else {
				require.Len(t, chainAddrs, 3)
			}
			containsForwarder := false
			for _, tv := range chainAddrs {
				if tv.Type == keystone.KeystoneForwarder {
					containsForwarder = true
					break
				}
			}
			require.True(t, containsForwarder, "no forwarder found in %v on chain %d for target don", chainAddrs, chain)
		}

		req := &keystone.GetContractSetsRequest{
			Chains:      e.Chains(),
			AddressBook: ad,
		}

		contractSetsResp, err := keystone.GetContractSets(lggr, req)
		require.NoError(t, err)
		require.Len(t, contractSetsResp.ContractSets, 4)
		// check the registry
		regChainContracts, ok := contractSetsResp.ContractSets[homeChainSel]
		require.True(t, ok)
		gotRegistry := regChainContracts.CapabilitiesRegistry
		require.NotNil(t, gotRegistry)
		// contract reads
		gotDons, err := gotRegistry.GetDONs(&bind.CallOpts{})
		if err != nil {
			err = keystone.DecodeErr(kcr.CapabilitiesRegistryABI, err)
			require.Fail(t, fmt.Sprintf("failed to get Dons from registry at %s: %s", gotRegistry.Address().String(), err))
		}
		require.NoError(t, err)
		assert.Len(t, gotDons, len(e.DonToEnv))
		var donIds []uint32
		for n, info := range deployResp.DonInfos {
			donIds = append(donIds, info.Id)
			found := false
			for _, gdon := range gotDons {
				if gdon.Id == info.Id {
					found = true
					assert.EqualValues(t, info, gdon)
					break
				}
			}
			require.True(t, found, "don %s not found in registry", n)
		}
		// check the forwarder
		for _, cs := range contractSetsResp.ContractSets {
			forwarder := cs.Forwarder
			require.NotNil(t, forwarder)
			// TODO expand this test; there is no get method on the forwarder so unclear how to test it
		}
		// check the ocr3 contract
		for chainSel, cs := range contractSetsResp.ContractSets {
			if chainSel != homeChainSel {
				require.Nil(t, cs.OCR3)
				continue
			}
			require.NotNil(t, cs.OCR3)
		}
	})

	t.Run("memory chains clo offchain", func(t *testing.T) {
		wfNops := loadTestNops(t, "../clo/testdata/workflow_nodes.json")
		cwNops := loadTestNops(t, "../clo/testdata/chain_writer_nodes.json")
		var allNops []*models.NodeOperator
		allNops = append(allNops, wfNops...)
		allNops = append(allNops, cwNops...)

		registryChainIdStr := "11155111"
		var nodeToNop = make(map[string]kcr.CapabilitiesRegistryNodeOperator) //node -> nop
		for _, nop := range allNops {
			for _, node := range nop.Nodes {
				// admin address is on the registry chain
				found := false
				for _, chain := range node.ChainConfigs {
					if chain.Network.ChainID != registryChainIdStr {
						continue
					}

					nodeToNop[node.ID] = kcr.CapabilitiesRegistryNodeOperator{
						Admin: adminAddr(chain.AdminAddress),
						Name:  nop.Name}
					found = true
				}
				require.True(t, found, "no registry chain found for node %s", node.ID)
			}
		}

		wfDon := clo.NewDonEnvWithMemoryChains(t, clo.DonEnvConfig{
			DonName: keystone.WFDonName,
			Nops:    wfNops,
			Logger:  lggr,
		})

		cwDon := clo.NewDonEnvWithMemoryChains(t, clo.DonEnvConfig{
			DonName: keystone.TargetDonName,
			Nops:    cwNops,
			Logger:  lggr,
		})

		var donToEnv = map[string]*deployment.Environment{
			keystone.WFDonName:     wfDon,
			keystone.TargetDonName: cwDon,
		}

		menv := clo.NewMultiDonEnvironment(lggr, donToEnv)

		var donsToDeploy = map[string][]kcr.CapabilitiesRegistryCapability{
			keystone.WFDonName:     []kcr.CapabilitiesRegistryCapability{keystone.OCR3Cap},
			keystone.TargetDonName: []kcr.CapabilitiesRegistryCapability{keystone.WriteChainCap},
		}

		var ocr3Config = keystone.OracleConfigSource{
			MaxFaultyOracles: len(wfNops) / 3,
		}

		ctx := context.Background()

		// sepolia
		homeChainSel, err := chainsel.SelectorFromChainId(11155111)
		require.NoError(t, err)
		deployReq := keystone.DeployRequest{
			RegistryChainSel:  homeChainSel,
			Menv:              menv,
			DonToCapabilities: donsToDeploy,
			NodeIDToNop:       nodeToNop,
			OCR3Config:        &ocr3Config,
		}

		deployResp, err := keystone.Deploy(ctx, lggr, deployReq)
		require.NoError(t, err)
		ad := deployResp.Changeset.AddressBook
		addrs, err := ad.Addresses()
		require.NoError(t, err)
		lggr.Infow("Deployed Keystone contracts", "address book", addrs)

	})
}

func loadTestNops(t *testing.T, pth string) []*models.NodeOperator {
	f, err := os.ReadFile(pth)
	require.NoError(t, err)
	var nops []*models.NodeOperator
	require.NoError(t, json.Unmarshal(f, &nops))
	return nops
}

var emptyAddr = "0x0000000000000000000000000000000000000000"

// compute the admin address from the string. If the address is empty, replaces the 0s with fs
// contract registry disallows 0x0 as an admin address, but our test net nops use it
func adminAddr(addr string) common.Address {
	needsFixing := addr == emptyAddr
	addr = strings.TrimPrefix(addr, "0x")
	if needsFixing {
		addr = strings.ReplaceAll(addr, "0", "f")
	}
	return common.HexToAddress(strings.TrimPrefix(addr, "0x"))
}
