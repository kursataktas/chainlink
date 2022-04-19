package terra

import (
	"github.com/smartcontractkit/sqlx"

	terradb "github.com/smartcontractkit/chainlink-terra/pkg/terra/db"

	"github.com/smartcontractkit/chainlink/core/chains"
	"github.com/smartcontractkit/chainlink/core/chains/terra/types"
	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/services/pg"
)

type orm struct {
	*chains.ChainsORM[string, terradb.ChainCfg, types.Chain]
	*chains.NodesORM[string, types.NewNode, terradb.Node]
}

var _ types.ORM = (*orm)(nil)

// NewORM returns an ORM backed by db.
func NewORM(db *sqlx.DB, lggr logger.Logger, cfg pg.LogConfig) types.ORM {
	q := pg.NewQ(db, lggr.Named("ORM"), cfg)
	return &orm{
		chains.NewChainsORM[string, terradb.ChainCfg, types.Chain](q, "terra"),
		chains.NewNodesORM[string, types.NewNode, terradb.Node](q, "terra", "tendermint_url"),
	}
}
