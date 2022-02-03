package keeper

import (
	"github.com/alice/pivx/x/pivx/types"
)

var _ types.QueryServer = Keeper{}
