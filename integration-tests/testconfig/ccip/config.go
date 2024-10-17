package ccip

import (
	"fmt"
	"strconv"

	"github.com/AlekSi/pointer"
	chainselectors "github.com/smartcontractkit/chain-selectors"

	ctfconfig "github.com/smartcontractkit/chainlink-testing-framework/lib/config"

	"github.com/smartcontractkit/chainlink/integration-tests/client"
)

const (
	E2E_JD_IMAGE       = "E2E_JD_IMAGE"
	E2E_JD_VERSION     = "E2E_JD_VERSION"
	E2E_JD_GRPC        = "E2E_JD_GRPC"
	E2E_JD_WSRPC       = "E2E_JD_WSRPC"
	DEFAULT_DB_NAME    = "JD_DB"
	DEFAULT_DB_VERSION = "14.1"
)

var (
	ErrInvalidHomeChainSelector = fmt.Errorf("invalid home chain selector")
	ErrInvalidFeedChainSelector = fmt.Errorf("invalid feed chain selector")
)

type Config struct {
	PrivateEthereumNetworks map[string]*ctfconfig.EthereumNetworkConfig `toml:",omitempty"`
	CLNode                  *NodeConfig                                 `toml:",omitempty"`
	JobDistributorConfig    JDConfig                                    `toml:",omitempty"`
	HomeChainSelector       *string                                     `toml:",omitempty"`
	FeedChainSelector       *string                                     `toml:",omitempty"`
	RMNConfig               RMNConfig                                   `toml:",omitempty"`
}

type RMNConfig struct {
	NoOfNodes    *int    `toml:",omitempty"`
	ProxyImage   *string `toml:",omitempty"`
	ProxyVersion *string `toml:",omitempty"`
	AFNImage     *string `toml:",omitempty"`
	AFNVersion   *string `toml:",omitempty"`
}

type NodeConfig struct {
	NoOfPluginNodes *int                    `toml:",omitempty"`
	NoOfBootstraps  *int                    `toml:",omitempty"`
	ClientConfig    *client.ChainlinkConfig `toml:",omitempty"`
}

type JDConfig struct {
	Image     *string `toml:",omitempty"`
	Version   *string `toml:",omitempty"`
	DBName    *string `toml:",omitempty"`
	DBVersion *string `toml:",omitempty"`
	JDGRPC    *string `toml:",omitempty"`
	JDWSRPC   *string `toml:",omitempty"`
}

// TODO: include all JD specific input in generic secret handling
func (o *JDConfig) GetJDGRPC() string {
	grpc := pointer.GetString(o.JDGRPC)
	if grpc == "" {
		return ctfconfig.MustReadEnvVar_String(E2E_JD_GRPC)
	}
	return grpc
}

func (o *JDConfig) GetJDWSRPC() string {
	wsrpc := pointer.GetString(o.JDWSRPC)
	if wsrpc == "" {
		return ctfconfig.MustReadEnvVar_String(E2E_JD_WSRPC)
	}
	return wsrpc
}

func (o *JDConfig) GetJDImage() string {
	image := pointer.GetString(o.Image)
	if image == "" {
		return ctfconfig.MustReadEnvVar_String(E2E_JD_IMAGE)
	}
	return image
}

func (o *JDConfig) GetJDVersion() string {
	version := pointer.GetString(o.Version)
	if version == "" {
		return ctfconfig.MustReadEnvVar_String(E2E_JD_VERSION)
	}
	return version
}

func (o *JDConfig) GetJDDBName() string {
	dbname := pointer.GetString(o.DBName)
	if dbname == "" {
		return DEFAULT_DB_NAME
	}
	return dbname
}

func (o *JDConfig) GetJDDBVersion() string {
	dbversion := pointer.GetString(o.DBVersion)
	if dbversion == "" {
		return DEFAULT_DB_VERSION
	}
	return dbversion
}

func (o *Config) Validate() error {
	var chainIds []uint64
	for _, network := range o.PrivateEthereumNetworks {
		chainIds = append(chainIds, uint64(network.EthereumChainConfig.ChainID))
	}
	homeChainSelector, err := strconv.ParseUint(pointer.GetString(o.HomeChainSelector), 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse home chain selector: %w", err)
	}
	ok, err := IsSelectorValid(homeChainSelector, chainIds)
	if err != nil {
		return fmt.Errorf("failed to validate home chain selector %w", err)
	}
	if !ok {
		return ErrInvalidHomeChainSelector
	}
	feedChainSelector, err := strconv.ParseUint(pointer.GetString(o.FeedChainSelector), 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse feed chain selector: %w", err)
	}
	ok, err = IsSelectorValid(feedChainSelector, chainIds)
	if err != nil {
		return fmt.Errorf("failed to validate feed chain selector %w", err)
	}
	if !ok {
		return ErrInvalidFeedChainSelector
	}
	return nil
}

func IsSelectorValid(selector uint64, chainIds []uint64) (bool, error) {
	chainId, err := chainselectors.ChainIdFromSelector(selector)
	if err != nil {
		return false, err
	}
	for _, id := range chainIds {
		if id == chainId {
			return true, nil
		}
	}
	return false, nil
}
