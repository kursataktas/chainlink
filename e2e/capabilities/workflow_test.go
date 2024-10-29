package capabilities_test

import (
	"context"
	"encoding/hex"
	"fmt"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/libocr/offchainreporting2/confighelper"
	"github.com/smartcontractkit/libocr/offchainreporting2/types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3confighelper"

	"github.com/smartcontractkit/chainlink-testing-framework/framework"
	"github.com/smartcontractkit/chainlink-testing-framework/framework/clclient"
	"github.com/smartcontractkit/chainlink-testing-framework/framework/components/blockchain"
	"github.com/smartcontractkit/chainlink-testing-framework/framework/components/fake"
	ns "github.com/smartcontractkit/chainlink-testing-framework/framework/components/simple_node_set"
	"github.com/smartcontractkit/chainlink-testing-framework/seth"
	"github.com/smartcontractkit/chainlink/e2e/capabilities/components/evmcontracts/simple_ocr"
)

type WorkflowTestConfig struct {
	BlockchainA        *blockchain.Input `toml:"blockchain_a" validate:"required"`
	MockerDataProvider *fake.Input       `toml:"data_provider" validate:"required"`
	NodeSet            *ns.Input         `toml:"nodeset" validate:"required"`
}

type OCR3Config struct {
	Signers               []types.OnchainPublicKey
	Transmitters          []types.Account
	F                     uint8
	OnchainConfig         []byte
	OffchainConfigVersion uint64
	OffchainConfig        []byte
}

func generateOCR3Config(t *testing.T, nodes []*clclient.ChainlinkClient) (*OCR3Config, error) {
	// type OracleIdentity struct {
	// 	OnChainSigningAddress types.OnChainSigningAddress
	// 	TransmitAddress       common.Address
	// 	OffchainPublicKey     types.OffchainPublicKey
	// 	PeerID                string
	// }
	oracleIdentities := []confighelper.OracleIdentityExtra{}
	transmissionSchedule := []int{}

	for _, node := range nodes {
		transmissionSchedule = append(transmissionSchedule, 0)
		oracleIdentity := confighelper.OracleIdentityExtra{}
		// ocr2
		ocr2Keys, err := node.MustReadOCR2Keys()
		require.NoError(t, err)

		firstOCR2Key := ocr2Keys.Data[0].Attributes

		offchainPublicKeyBytes, err := hex.DecodeString(firstOCR2Key.OffChainPublicKey)
		require.NoError(t, err)
		var offchainPublicKey [32]byte
		copy(offchainPublicKey[:], offchainPublicKeyBytes)
		oracleIdentity.OffchainPublicKey = offchainPublicKey

		onchainPubkey, err := hex.DecodeString(firstOCR2Key.OnChainPublicKey)
		require.NoError(t, err)
		oracleIdentity.OnchainPublicKey = onchainPubkey

		sharedSecretEncryptionPublicKeyBytes, err := hex.DecodeString(firstOCR2Key.ConfigPublicKey)
		require.NoError(t, err)
		var sharedSecretEncryptionPublicKey [32]byte
		copy(sharedSecretEncryptionPublicKey[:], sharedSecretEncryptionPublicKeyBytes)
		oracleIdentity.ConfigEncryptionPublicKey = sharedSecretEncryptionPublicKey

		// p2p
		p2pKeys, err := node.MustReadP2PKeys()
		require.NoError(t, err)
		oracleIdentity.PeerID = p2pKeys.Data[0].Attributes.PeerID

		// eth
		ethKeys, err := node.MustReadETHKeys()
		require.NoError(t, err)
		oracleIdentity.TransmitAccount = types.Account(ethKeys.Data[0].Attributes.Address)

		oracleIdentities = append(oracleIdentities, oracleIdentity)
	}

	maxDurationInitialization := 10 * time.Second

	// Generate OCR3 configuration arguments for testing
	signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig, err := ocr3confighelper.ContractSetConfigArgsForTests(
		5*time.Second,              // DeltaProgress: Time between rounds
		5*time.Second,              // DeltaResend: Time between resending unconfirmed transmissions
		5*time.Second,              // DeltaInitial: Initial delay before starting the first round
		2*time.Second,              // DeltaRound: Time between rounds within an epoch
		500*time.Millisecond,       // DeltaGrace: Grace period for delayed transmissions
		1*time.Second,              // DeltaCertifiedCommitRequest: Time between certified commit requests
		30*time.Second,             // DeltaStage: Time between stages of the protocol
		uint64(10),                 // MaxRoundsPerEpoch: Maximum number of rounds per epoch
		transmissionSchedule,       // TransmissionSchedule: Transmission schedule
		oracleIdentities,           // Oracle identities with their public keys
		nil,                        // Plugin config (empty for now)
		&maxDurationInitialization, // MaxDurationInitialization: ???
		1*time.Second,              // MaxDurationQuery: Maximum duration for querying
		1*time.Second,              // MaxDurationObservation: Maximum duration for observation
		1*time.Second,              // MaxDurationAccept: Maximum duration for acceptance
		1*time.Second,              // MaxDurationTransmit: Maximum duration for transmission
		1,                          // F: Maximum number of faulty oracles
		nil,                        // OnChain config (empty for now)
	)
	require.NoError(t, err)

	// maxDurationInitialization *time.Duration,
	// maxDurationQuery time.Duration,
	// maxDurationObservation time.Duration,
	// maxDurationShouldAcceptAttestedReport time.Duration,
	// maxDurationShouldTransmitAcceptedReport time.Duration,

	return &OCR3Config{
		Signers:               signers,
		Transmitters:          transmitters,
		F:                     f,
		OnchainConfig:         onchainConfig,
		OffchainConfigVersion: offchainConfigVersion,
		OffchainConfig:        offchainConfig,
	}, nil
}

func TestWorkflow(t *testing.T) {
	t.Run("smoke test", func(t *testing.T) {
		in, err := framework.Load[WorkflowTestConfig](t)
		require.NoError(t, err)

		// deploy docker test environment
		bc, err := blockchain.NewBlockchainNetwork(in.BlockchainA)
		require.NoError(t, err)

		nodeset, err := ns.NewNodeSet(in.NodeSet, bc, "https://example.com") // TODO: Should not be a thing
		require.NoError(t, err)

		for i, n := range nodeset.CLNodes {
			fmt.Printf("Node %d --> %s\n", i, n.Node.HostURL)
			fmt.Printf("Node P2P %d --> %s\n", i, n.Node.HostP2PURL)
		}

		// connect clients
		sc, err := seth.NewClientBuilder().
			WithRpcUrl(bc.Nodes[0].HostWSUrl).
			WithPrivateKeys([]string{os.Getenv("PRIVATE_KEY")}).
			Build()
		require.NoError(t, err)

		nodeClients, err := clclient.NewCLCDefaultlients(nodeset.CLNodes, framework.L)
		require.NoError(t, err)

		fmt.Println("Setting up KV store capabilities...")

		simpleOCRAddress, tx, _ /* KVStoreOCRContract */, err := simple_ocr.DeploySimpleOCR(
			sc.NewTXOpts(),
			sc.Client,
		)
		require.NoError(t, err)
		fmt.Println("Deployed Simple OCR contract at", simpleOCRAddress.Hex())

		_, err = bind.WaitMined(context.Background(), sc.Client, tx)
		require.NoError(t, err)

		// Add bootstrap spec to the first node
		bootstrapNode := nodeClients[0]
		p2pKeys, err := bootstrapNode.MustReadP2PKeys()
		require.NoError(t, err)
		fmt.Println("P2P keys fetched")
		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
			defer wg.Done()
			r, _, err := bootstrapNode.CreateJobRaw(fmt.Sprintf(`
				type = "bootstrap"
				schemaVersion = 1
				name = "Botostrap"
				contractID = "%s"
				contractConfigTrackerPollInterval = "1s"
				contractConfigConfirmations = 1
				relay = "evm"
		
				[relayConfig]
				chainID = %s
			`, simpleOCRAddress, bc.ChainID))
			require.NoError(t, err)
			require.Equal(t, len(r.Errors), 0)
			fmt.Printf("Response from bootstrap node: %x\n", r)
		}()

		for i, nodeClient := range nodeClients {
			// First node is a bootstrap node, so we skip it
			if i == 0 {
				continue
			}

			wg.Add(1)
			go func() {
				defer wg.Done()
				response, _, err := nodeClient.CreateJobRaw(fmt.Sprintf(`
					type = "standardcapabilities"
					schemaVersion = 1
					name = "%s-capabilities"
					command="%s"

					[oracle_factory]
					enabled=true
					bootstrap_peers = [
						"%s@%s"
					]
					network="%s"
					chain_id="%s"
					ocr_contract_address="%s"`,
					"kvstore",
					"./capabilities/kvstore",
					p2pKeys.Data[0].Attributes.PeerID,
					nodeset.CLNodes[0].Node.HostP2PURL,
					"evm",
					bc.ChainID,
					simpleOCRAddress,
				))
				require.NoError(t, err)
				require.Equal(t, len(response.Errors), 0)
				fmt.Printf("Response from node %d: %x\n", i+1, response)
			}()
		}

		ocr3Config, err := generateOCR3Config(t, nodeClients)
		require.NoError(t, err)

		fmt.Println("ocr3Config", ocr3Config)

		wg.Wait()

		// Configure KV store OCR contract
		// KVStoreOCRContract.SetConfig(
		// 	sc,
		// 	// _signers []common.Address,
		// 	// _transmitters []common.Address,
		// 	// _f uint8,
		// 	// _onchainConfig []byte,
		// 	// _offchainConfigVersion uint64,
		// 	// _offchainConfig []byte
		// )

		// Add bootstrap spec
		// ✅ 2. Deploy KV store OCR contract
		// ✅ 4. Add boostrap job spec
		// ✅ 4. Add KV store capabilities (hardocded binaries for now)
		// 1. Fetch node keys
		// 3. Configure OCR contract
		// 4.1. Add CRON capabilities
		// 4.2. EVM target capabilities
		// 5. TODOs: Have a workflow running and tested

		// interact with contracts
		// i, err := burn_mint_erc677.NewBurnMintERC677(contracts.Addresses[0], sc.Client)
		// require.NoError(t, err)
		// balance, err := i.BalanceOf(sc.NewCallOpts(), contracts.Addresses[0])
		// require.NoError(t, err)
		// fmt.Println(balance)

		// // create jobs using deployed contracts data, this is just an example
		// r, _, err := c[0].CreateJobRaw(`
		// type            = "cron"
		// schemaVersion   = 1
		// schedule        = "CRON_TZ=UTC */10 * * * * *" # every 10 secs
		// observationSource   = """
		//    // data source 2
		//    fetch         [type=http method=GET url="https://min-api.cryptocompare.com/data/pricemultifull?fsyms=ETH&tsyms=USD"];
		//    parse       [type=jsonparse path="RAW,ETH,USD,PRICE"];
		//    multiply    [type="multiply" input="$(parse)" times=100]
		//    encode_tx   [type="ethabiencode"
		//                 abi="submit(uint256 value)"
		//                 data="{ \\"value\\": $(multiply) }"]
		//    submit_tx   [type="ethtx" to="0x859AAa51961284C94d970B47E82b8771942F1980" data="$(encode_tx)"]

		//    fetch -> parse -> multiply -> encode_tx -> submit_tx
		// """`)
		// require.NoError(t, err)
		// require.Equal(t, len(r.Errors), 0)
		// fmt.Printf("error: %v", err)
		// fmt.Println(r)
	})

	//t.Run("can access mockserver", func(t *testing.T) {
	//	// on the host, locally
	//	client := resty.New()
	//	_, err := client.R().
	//		Get(fmt.Sprintf("%s/mock1", dp.BaseURLHost))
	//	require.NoError(t, err)
	//	// other components can access inside docker like this
	//	err = components.NewDockerFakeTester(fmt.Sprintf("%s/mock1", dp.BaseURLDocker))
	//	require.NoError(t, err)
	//})
}
