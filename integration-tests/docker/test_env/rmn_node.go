package test_env

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/pelletier/go-toml/v2"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/smartcontractkit/chainlink-testing-framework/lib/docker"
	"github.com/smartcontractkit/chainlink-testing-framework/lib/docker/test_env"
	"github.com/smartcontractkit/chainlink-testing-framework/lib/logging"
	"github.com/smartcontractkit/chainlink-testing-framework/lib/logstream"
	tc "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/exec"
	tcwait "github.com/testcontainers/testcontainers-go/wait"
)

const (
	DefaultAFNPasphrase = "my-not-so-secret-passphrase"
	RMNKeyStore         = "keystore/afn2proxy-keystore.json"
)

type Chain struct {
	Name string   `toml:"name"`
	RPCS []string `toml:"rpcs"`
}

type SharedConfig struct {
	Chains []SharedChain `toml:"chains"`
	Lanes  []Lane        `toml:"lanes"`
}

func (s SharedConfig) afn2ProxySharedConfigFile() (string, error) {
	data, err := toml.Marshal(s)
	if err != nil {
		return "", fmt.Errorf("failed to marshal afn2Proxy shared config: %w", err)
	}
	return CreateTempFile(data, "afn2proxy_shared")
}

type SharedChain struct {
	Name                                                             string          `toml:"name"`
	MaxTaggedRootsPerVoteToBless                                     int             `toml:"max_tagged_roots_per_vote_to_bless"`
	AfnType                                                          string          `toml:"afn_type"`
	AfnContract                                                      string          `toml:"afn_contract"`
	InflightTime                                                     Duration        `toml:"inflight_time"`
	MaxFreshBlockAge                                                 Duration        `toml:"max_fresh_block_age"`
	UponFinalityViolationVoteToCurseOnOtherChainsWithLegacyContracts bool            `toml:"upon_finality_violation_vote_to_curse_on_other_chains_with_legacy_contracts"`
	Stability                                                        StabilityConfig `toml:"stability"`
	BlessFeeConfig                                                   FeeConfig       `toml:"bless_fee_config"`
	CurseFeeConfig                                                   FeeConfig       `toml:"curse_fee_config"`
}

type Duration struct {
	Seconds int `toml:"Seconds,omitempty"`
	Minutes int `toml:"Minutes,omitempty"`
}

type StabilityConfig struct {
	Type              string `toml:"type"`
	SoftConfirmations int    `toml:"soft_confirmations"`
}

type FeeConfig struct {
	Type                 string `toml:"type"`
	MaxFeePerGas         *Gwei  `toml:"max_fee_per_gas,omitempty"`
	MaxPriorityFeePerGas *Gwei  `toml:"max_priority_fee_per_gas,omitempty"`
	GasPrice             *Gwei  `toml:"gas_price,omitempty"`
}

type Gwei struct {
	Gwei int `toml:"Gwei"`
}

type Lane struct {
	Name                   string `toml:"name"`
	Type                   string `toml:"type"`
	SourceChainName        string `toml:"source_chain_name"`
	SourceStartBlockNumber int    `toml:"source_start_block_number"`
	DestChainName          string `toml:"dest_chain_name"`
	DestStartBlockNumber   int    `toml:"dest_start_block_number"`
	OnRamp                 string `toml:"on_ramp"`
	OffRamp                string `toml:"off_ramp"`
	CommitStore            string `toml:"commit_store"`
}

type LocalConfig struct {
	Chains []Chain `toml:"chains"`
}

func (l LocalConfig) afn2ProxyLocalConfigFile() (string, error) {
	data, err := toml.Marshal(l)
	if err != nil {
		return "", fmt.Errorf("failed to marshal afn2Proxy local config: %w", err)
	}
	return CreateTempFile(data, "afn2proxy_local")
}

func CreateTempFile(data []byte, pattern string) (string, error) {
	file, err := os.CreateTemp("", pattern)
	if err != nil {
		return "", fmt.Errorf("failed to create temp file for %s: %w", pattern, err)
	}
	_, err = file.Write(data)
	if err != nil {
		return "", fmt.Errorf("failed to write  %s: %w", pattern, err)
	}
	return file.Name(), nil
}

type RmnNode struct {
	test_env.EnvComponent
	AFNPassphrase  string
	Shared         SharedConfig
	Local          LocalConfig
	BlessCurseKeys map[string]BlessCurseKeys
	t              *testing.T
	l              zerolog.Logger
}

func NewRmnNode(
	networks []string,
	name,
	imageName,
	imageVersion string,
	shared SharedConfig,
	local LocalConfig,
	logStream *logstream.LogStream) (*RmnNode, error) {
	afnName := fmt.Sprintf("%s-%s", name, uuid.NewString()[0:8])
	rmn := &RmnNode{
		EnvComponent: test_env.EnvComponent{
			ContainerName:    afnName,
			ContainerImage:   imageName,
			ContainerVersion: imageVersion,
			Networks:         networks,
			LogStream:        logStream,
		},
		AFNPassphrase: DefaultAFNPasphrase,
		Shared:        shared,
		Local:         local,
	}

	return rmn, nil
}

func (rmn *RmnNode) StartContainer(reuse bool) error {
	localAFN2Proxy, err := rmn.Local.afn2ProxyLocalConfigFile()
	if err != nil {
		return err
	}
	sharedAFN2Proxy, err := rmn.Shared.afn2ProxySharedConfigFile()
	if err != nil {
		return err
	}

	l := tc.Logger
	if rmn.t != nil {
		l = logging.CustomT{
			T: rmn.t,
			L: rmn.l,
		}
	}
	container, err := docker.StartContainerWithRetry(rmn.l, tc.GenericContainerRequest{
		ContainerRequest: tc.ContainerRequest{
			Name:  rmn.ContainerName,
			Image: fmt.Sprintf("%s:%s", rmn.ContainerImage, rmn.ContainerVersion),
			Env: map[string]string{
				"AFN_PASSPHRASE": rmn.AFNPassphrase,
			},
			Networks: append(rmn.Networks, "tracing"),
			Files: []tc.ContainerFile{
				{
					HostFilePath:      sharedAFN2Proxy,
					ContainerFilePath: "/app/cfg/afn2proxy_shared.toml",
					FileMode:          0644,
				},
				{
					HostFilePath:      localAFN2Proxy,
					ContainerFilePath: "/app/cfg/afn2proxy_local.toml",
					FileMode:          0644,
				},
			},
			WaitingFor: tcwait.ForExec([]string{"cat", RMNKeyStore}),
			LifecycleHooks: []tc.ContainerLifecycleHooks{
				{
					PostStarts:    rmn.PostStartsHooks,
					PostStops:     rmn.PostStopsHooks,
					PreTerminates: rmn.PreTerminatesHooks,
				},
			},
		},
		Started: true,
		Reuse:   reuse,
		Logger:  l,
	})
	if err != nil {
		return err
	}
	_, reader, err := container.Exec(context.Background(), []string{
		"cat", RMNKeyStore}, exec.Multiplexed())
	if err != nil {
		return errors.Wrapf(err, "Unable to cat keystore")
	}
	b, err := io.ReadAll(reader)
	if err != nil {
		return err
	}
	blessCurseKeys, err := parseBlessCurseKeys(b)
	if err != nil {
		return errors.Wrapf(err, "Unable to extract peerID %s", string(b))
	}
	rmn.BlessCurseKeys = blessCurseKeys
	rmn.Container = container
	return nil
}

// Define the structure for BlessCurseKeys
type BlessCurseKeys struct {
	Bless string `json:"bless"`
	Curse string `json:"curse"`
}

// Define the structure for the JSON data
type Data struct {
	AssociatedData string `json:"associated_data"`
}

func parseBlessCurseKeys(jsonData []byte) (map[string]BlessCurseKeys, error) {
	var data Data
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		return nil, err
	}

	s := data.AssociatedData

	// Remove "BlessCurseKeysByChain(" and the last ")"
	prefix := "BlessCurseKeysByChain("
	suffix := ")"
	if strings.HasPrefix(s, prefix) && strings.HasSuffix(s, suffix) {
		s = s[len(prefix) : len(s)-len(suffix)]
	} else {
		return nil, fmt.Errorf("unexpected format")
	}

	// Remove "BlessCurseKeys"
	s = strings.ReplaceAll(s, "BlessCurseKeys", "")

	// Replace unquoted keys with quoted keys
	reKey := regexp.MustCompile(`(\w+):`)
	s = reKey.ReplaceAllString(s, `"$1":`)

	// Put double quotes around hex values
	reHex := regexp.MustCompile(`(0x[0-9a-fA-F]+)`)
	s = reHex.ReplaceAllString(s, `"$1"`)

	// Now, s should be valid JSON
	// Add outer braces if needed
	if !strings.HasPrefix(s, "{") {
		s = "{" + s + "}"
	}

	// Now unmarshal s into map[string]BlessCurseKeys
	var result map[string]BlessCurseKeys
	err = json.Unmarshal([]byte(s), &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
