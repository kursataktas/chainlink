package evmcontracts

import (
	"context"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-testing-framework/seth"
	"github.com/smartcontractkit/chainlink/e2e/capabilities/components/evmcontracts/simple_ocr"
)

type Input struct {
	URL string  `toml:"url"`
	Out *Output `toml:"out"`
}

type Output struct {
	UseCache  bool             `toml:"use_cache"`
	Addresses []common.Address `toml:"addresses"`
}

func NewProductOnChainDeployment(in *Input) (*Output, error) {
	if in.Out.UseCache {
		return in.Out, nil
	}

	// deploy your contracts here, example

	c, err := seth.NewClientBuilder().
		WithRpcUrl(in.URL).
		WithPrivateKeys([]string{os.Getenv("PRIVATE_KEY")}).
		Build()
	if err != nil {
		return nil, err
	}

	addr, tx, _, err := simple_ocr.DeploySimpleOCR(
		c.NewTXOpts(),
		c.Client,
	)
	if err != nil {
		return nil, err
	}

	_, err = bind.WaitMined(context.Background(), c.Client, tx)
	if err != nil {
		return nil, err
	}

	out := &Output{
		UseCache: true,
		// save all the addresses to output, so it can be cached
		Addresses: []common.Address{addr},
	}
	in.Out = out
	return out, nil
}
