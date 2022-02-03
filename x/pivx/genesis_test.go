package pivx_test

import (
	"testing"

	keepertest "github.com/alice/pivx/testutil/keeper"
	"github.com/alice/pivx/testutil/nullify"
	"github.com/alice/pivx/x/pivx"
	"github.com/alice/pivx/x/pivx/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.PivxKeeper(t)
	pivx.InitGenesis(ctx, *k, genesisState)
	got := pivx.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
