package keeper_test

import (
	"testing"

	testkeeper "github.com/alice/pivx/testutil/keeper"
	"github.com/alice/pivx/x/pivx/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.PivxKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
