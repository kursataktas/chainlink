package versiontest

import (
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets"
	kcr_1_0_1 "github.com/smartcontractkit/chainlink/v2/core/gethwrappers/keystone/generated/capabilities_registry_1_1_0"
	forwarder_1_0_0 "github.com/smartcontractkit/chainlink/v2/core/gethwrappers/keystone/generated/forwarder_1_0_0"
	ocr3_1_0_0 "github.com/smartcontractkit/chainlink/v2/core/gethwrappers/keystone/generated/ocr3_capability_1_0_0"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
)

func TestVersion(t *testing.T) {
	owner := testutils.MustNewSimTransactor(t)
	chain := backends.NewSimulatedBackend(core.GenesisAlloc{
		owner.From: {Balance: assets.Ether(100).ToInt()},
	}, 30e6)

	t.Run("capabilities registry v1_0_1", func(t *testing.T) {
		_, _, capreg, err := kcr_1_0_1.DeployCapabilitiesRegistry(owner, chain)
		require.NoError(t, err)
		chain.Commit()
		tv, err := capreg.TypeAndVersion(nil)
		require.NoError(t, err)
		require.Equal(t, "CapabilitiesRegistry 1.0.1", tv)
	})

	t.Run("ocr3 capability v1_0_0", func(t *testing.T) {
		_, _, ocr3, err := ocr3_1_0_0.DeployOCR3Capability(owner, chain)
		require.NoError(t, err)
		chain.Commit()
		tv, err := ocr3.TypeAndVersion(nil)
		require.NoError(t, err)
		require.Equal(t, "Keystone 1.0.0", tv)
	})

	t.Run("forwarder v1_0_0", func(t *testing.T) {
		_, _, forwarder, err := forwarder_1_0_0.DeployKeystoneForwarder(owner, chain)
		require.NoError(t, err)
		chain.Commit()
		tv, err := forwarder.TypeAndVersion(nil)
		require.NoError(t, err)
		require.Equal(t, "Forwarder and Router 1.0.0", tv)
	})
}
