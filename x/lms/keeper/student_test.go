package keeper

import (
	"clms/x/lms/types"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// type k struct {
// 	Keeper
// }
// func NewKeeper(key storetypes.StoreKey, cdc codec.Codec) Keeper {
// 	return Keeper{
// 		cdc:      cdc,
// 		storeKey: key,
// 	}
// }
func (k Keeper) TestAdminRegister(t *testing.T) {

	req := types.RegisterAdminRequest{
		Name:    "visal",
		Address: "bns",
	}
	ctx := sdk.Context{}
	k.AdminRegister(ctx, &req)
}

// func (k Keeper)TestXxx(t *testing.T) {

// }

// func TestSome(t *testing.T) {
// 	i := 0
// 	if i == 0 {
// 		fmt.Println("passed")
// 	}
// }
