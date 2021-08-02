package keeper

import (
	"github.com/toschdev/blog/x/blog/types"
)

var _ types.QueryServer = Keeper{}
