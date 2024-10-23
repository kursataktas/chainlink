package capabilities_test

import (
	"fmt"
	"github.com/smartcontractkit/chainlink-testing-framework/framework"
	"github.com/smartcontractkit/chainlink-testing-framework/framework/clclient"
	"github.com/smartcontractkit/chainlink-testing-framework/framework/components/blockchain"
	"github.com/smartcontractkit/chainlink-testing-framework/framework/components/don"
	"github.com/smartcontractkit/chainlink-testing-framework/framework/components/fake"
	burn_mint_erc677 "github.com/smartcontractkit/chainlink/e2e/capabilities/components/gethwrappers"
	"github.com/smartcontractkit/chainlink/e2e/capabilities/components/onchain"
	"github.com/smartcontractkit/seth"
	"github.com/stretchr/testify/require"
	"testing"
)

type Config struct {
	BlockchainA        *blockchain.Input `toml:"blockchain_a" validate:"required"`
	Contracts          *onchain.Input    `toml:"contracts" validate:"required"`
	MockerDataProvider *fake.Input       `toml:"data_provider" validate:"required"`
	DONInput           *don.Input        `toml:"don" validate:"required"`
}

func TestDON(t *testing.T) {
	in, err := framework.Load[Config](t)
	require.NoError(t, err)

	// deploy docker test environment
	bc, err := blockchain.NewBlockchainNetwork(in.BlockchainA)
	require.NoError(t, err)
	dp, err := fake.NewFakeDataProvider(in.MockerDataProvider)
	require.NoError(t, err)
	out, err := don.NewBasicDON(in.DONInput, bc, dp.BaseURLDocker)
	require.NoError(t, err)
	for i, n := range out.Nodes {
		fmt.Printf("Node %d --> %s\n", i, n.Node.HostURL)
	}

	// deploy product contracts
	in.Contracts.URL = bc.Nodes[0].HostWSUrl
	contracts, err := onchain.NewProductOnChainDeployment(in.Contracts)

	// connect clients
	sc, err := seth.NewClientBuilder().
		WithRpcUrl(bc.Nodes[0].HostWSUrl).
		WithPrivateKeys([]string{blockchain.DefaultAnvilPrivateKey}).
		Build()
	require.NoError(t, err)

	c, err := clclient.NewCLCDefaultlients(out.Nodes, framework.L)
	require.NoError(t, err)

	t.Run("smoke test", func(t *testing.T) {
		// interact with contracts
		i, err := burn_mint_erc677.NewBurnMintERC677(contracts.Addresses[0], sc.Client)
		require.NoError(t, err)
		balance, err := i.BalanceOf(sc.NewCallOpts(), contracts.Addresses[0])
		require.NoError(t, err)
		fmt.Println(balance)

		// create jobs using deployed contracts data, this is just an example
		r, _, err := c[0].CreateJobRaw(`
		type            = "cron"
		schemaVersion   = 1
		schedule        = "CRON_TZ=UTC */10 * * * * *" # every 10 secs
		observationSource   = """
		   // data source 2
		   fetch         [type=http method=GET url="https://min-api.cryptocompare.com/data/pricemultifull?fsyms=ETH&tsyms=USD"];
		   parse       [type=jsonparse path="RAW,ETH,USD,PRICE"];
		   multiply    [type="multiply" input="$(parse)" times=100]
		   encode_tx   [type="ethabiencode"
		                abi="submit(uint256 value)"
		                data="{ \\"value\\": $(multiply) }"]
		   submit_tx   [type="ethtx" to="0x859AAa51961284C94d970B47E82b8771942F1980" data="$(encode_tx)"]
		
		   fetch -> parse -> multiply -> encode_tx -> submit_tx
		"""`)
		require.NoError(t, err)
		require.Equal(t, len(r.Errors), 0)
		fmt.Printf("error: %v", err)
		fmt.Println(r)
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
