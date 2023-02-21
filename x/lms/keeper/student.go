package keeper

import (
	"clms/x/lms/types"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	// "google.golang.org/genproto/googleapis/actions/sdk/v2"
)

func (k Keeper) AdminRegister(ctx sdk.Context, req *types.RegisterAdminRequest) error {
	if _, err := sdk.AccAddressFromBech32(req.Address); err != nil {
		panic(fmt.Errorf("invalid admin address %w", err))
	}
	// sdk.AccAddressFromBech32()
	store := ctx.KVStore(k.storeKey)
	bz, err := k.cdc.Marshal(req)
	if err != nil {
		fmt.Println("error ")
		return err
	} else {
		fmt.Println("Admin added successfullly")
		store.Set(types.AdminstoreId(req.Address), bz)
	}
	return nil
}

func (k Keeper) AcceptLeaves(ctx sdk.Context, req *types.AcceptLeaveRequest) error {
	if _, err := sdk.AccAddressFromBech32(req.Admin); err != nil {
		panic(fmt.Errorf("invalid bank authority address: %w", err))
	}
	store := ctx.KVStore(k.storeKey)

	req.Status = types.LeaveStatus_STATUS_ACCEPTED
	bz, err := k.cdc.Marshal(req)
	if err != nil {
		return err
	} else {
		store.Set(types.StudentStoreId(req.Admin), bz)
		fmt.Println("Admin is registered")
	}
	return nil
}
func (k Keeper) ApplyLeaves(ctx sdk.Context, req *types.ApplyLeaveRequest) error {
	return nil
}
func (k Keeper)GetData()
