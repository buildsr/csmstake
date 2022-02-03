package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/alice/pivx/testutil/keeper"
	"github.com/alice/pivx/x/pivx/keeper"
	"github.com/alice/pivx/x/pivx/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.PivxKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
