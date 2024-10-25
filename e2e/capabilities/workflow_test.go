package capabilities_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/stretchr/testify/require"

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

func TestWorkflow(t *testing.T) {
	t.Run("smoke test", func(t *testing.T) {
		in, err := framework.Load[WorkflowTestConfig](t)
		require.NoError(t, err)

		// deploy docker test environment
		bc, err := blockchain.NewBlockchainNetwork(in.BlockchainA)
		require.NoError(t, err)

		out, err := ns.NewNodeSet(in.NodeSet, bc, "https://example.com") // TODO: Should not be a thing
		require.NoError(t, err)

		for i, n := range out.CLNodes {
			fmt.Printf("Node %d --> %s\n", i, n.Node.HostURL)
			fmt.Printf("Node P2P %d --> %s\n", i, n.Node.HostP2PURL)
		}

		// connect clients
		sc, err := seth.NewClientBuilder().
			WithRpcUrl(bc.Nodes[0].HostWSUrl).
			WithPrivateKeys([]string{os.Getenv("PRIVATE_KEY")}).
			Build()
		require.NoError(t, err)

		c, err := clclient.NewCLCDefaultlients(out.CLNodes, framework.L)
		require.NoError(t, err)

		fmt.Println("Setting up KV store capabilities...")

		simpleOCRAddress, tx, _, err := simple_ocr.DeploySimpleOCR(
			sc.NewTXOpts(),
			sc.Client,
		)
		require.NoError(t, err)
		fmt.Println("Deployed Simple OCR contract at", simpleOCRAddress.Hex())

		_, err = bind.WaitMined(context.Background(), sc.Client, tx)
		require.NoError(t, err)

		// Add bootstrap spec to the first node
		r, _, err := c[0].CreateJobRaw(fmt.Sprintf(`
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
		fmt.Println(r)

		p2pKeys, err := c[0].MustReadP2PKeys()
		require.NoError(t, err)

		fmt.Println("P2P keys", p2pKeys)

		// t.Fail()

		// for i, nodeClient := range c {
		// 	// First node is a bootstrap node, so we skip it
		// 	if i == 0 {
		// 		continue
		// 	}

		// 	response, _, err := nodeClient.CreateJobRaw(fmt.Sprintf(`
		// 		type = "standardcapabilities"
		// 		schemaVersion = 1
		// 		name = "%s-capabilities"
		// 		command="%s"

		// 		[oracle_factory]
		// 		enabled=true
		// 		bootstrap_peers = [
		// 			"%s@localhost:%s"
		// 		]
		// 		network="%s"
		// 		chain_id="%s"
		// 		ocr_contract_address="%s"`,
		// 		"kvstore",
		// 		"./capabilities/kvstore",
		// 		p2pKeys.Data[0].Attributes.PeerID,
		// 		"", // bootstrapNodeInfo.Ports.P2P,
		// 		"evm",
		// 		bc.ChainID,
		// 		simpleOCRAddress,
		// 	))
		// 	require.NoError(t, err)
		// 	require.Equal(t, len(response.Errors), 0)
		// 	fmt.Println(r)
		// }

		// Add bootstrap spec
		// 1. Fetch node keys
		// 2. Deploy KV store OCR contract
		// 3. Configure OCR contract
		// 4. Add boostrap job spec
		// 4. Add KV store capabilities (hardocded binaries for now)
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
