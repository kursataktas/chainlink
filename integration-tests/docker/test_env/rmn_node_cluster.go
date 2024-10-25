package test_env

import (
	"context"
	"time"

	"golang.org/x/sync/errgroup"
)

type RmnCluster struct {
	Nodes []*RmnNode `json:"nodes"`
}

func (c *RmnCluster) Start() error {
	eg := &errgroup.Group{}
	nodes := c.Nodes

	for i := 0; i < len(nodes); i++ {
		nodeIndex := i
		eg.Go(func() error {
			reuse := false
			err := nodes[nodeIndex].StartContainer(reuse)
			if err != nil {
				return err
			}
			return nil
		})
	}

	return eg.Wait()
}

func (c *RmnCluster) Stop() error {
	var eg errgroup.Group
	nodes := c.Nodes
	timeout := time.Minute * 1

	for i := 0; i < len(nodes); i++ {
		nodeIndex := i
		eg.Go(func() error {
			if container := nodes[nodeIndex].Container; container != nil {
				return container.Stop(context.Background(), &timeout)
			}
			return nil
		})
	}

	return eg.Wait()
}
