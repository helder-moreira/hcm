package keeper

import (
	"github.com/helder-moreira/hcm/x/hcm/types"
)

var _ types.QueryServer = Keeper{}
