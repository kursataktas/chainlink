package capabilities

import (
	"github.com/smartcontractkit/chainlink-testing-framework/framework/components/blockchain"
	"github.com/smartcontractkit/chainlink-testing-framework/framework/components/clnode"
)

type Input struct {
	Nodes []*clnode.Input `toml:"nodes" validate:"required"`
	Out   *Output         `toml:"out"`
}

type Output struct {
	UseCache bool             `toml:"use_cache"`
	Nodes    []*clnode.Output `toml:"node"`
}

func AddKVStore(in *Input, bcOut *blockchain.Output) (*Output, error) {

	nodeOuts := make([]*clnode.Output, 0)
	for _, n := range in.Nodes {
		net, err := clnode.NewNetworkCfgOneNetworkAllNodes(bcOut)
		if err != nil {
			return nil, err
		}
		n.Node.TestConfigOverrides = net
		n.DataProviderURL = "https://example.com" // TODO: Remove this, should be optional
		o, err := clnode.NewNode(n)
		if err != nil {
			return nil, err
		}
		nodeOuts = append(nodeOuts, o)
	}
	out := &Output{
		UseCache: true,
		Nodes:    nodeOuts,
	}
	in.Out = out
	return out, nil
}
