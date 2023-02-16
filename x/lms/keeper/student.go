package keeper

import (
	"clms/x/lms/types"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	// "google.golang.org/genproto/googleapis/actions/sdk/v2"
)

type StudentKeeper interface {
	AdminRegister(ctx sdk.Context, req *types.RegisterAdminRequest) error
	AcceptLeaves(ctx sdk.Context, req *types.AcceptLeaveRequest) error
}

var _ StudentKeeper = (*BaseStudentKeeper)(nil)

type BaseStudentKeeper struct {
	cdc      codec.BinaryCodec
	storeKey storetypes.StoreKey
}

func NewStudentKeeper(
	cdc codec.BinaryCodec,
	storekey storetypes.StoreKey,
) BaseStudentKeeper {
	if _, err := sdk.AccAddressFromBech32("h"); err != nil {
		panic(fmt.Errorf("invalid bank authority address: %w", err))
	}
	return BaseStudentKeeper{
		cdc:      cdc,
		storeKey: storekey,
	}
}
func (k BaseStudentKeeper) AdminRegister(ctx sdk.Context, req *types.RegisterAdminRequest) error {
	if _, err := sdk.AccAddressFromBech32(req.Address); err != nil {
		panic(fmt.Errorf("invalid bank authority address: %w", err))
	}
	store := ctx.KVStore(k.storeKey)
	bz, err := k.cdc.Marshal(req)
	if err != nil {
		return err
	} else {
		store.Set(types.AdminKey, bz)
	}
	return nil
}

func (k BaseStudentKeeper) AcceptLeaves(ctx sdk.Context, req *types.AcceptLeaveRequest) error {
	if _, err := sdk.AccAddressFromBech32(req.Admin); err != nil {
		panic(fmt.Errorf("invalid bank authority address: %w", err))
	}
	store := ctx.KVStore(k.storeKey)
	
	req.Status = types.LeaveStatus_STATUS_ACCEPTED
	bz, err := k.cdc.Marshal(req)
	if err != nil {
		return err
	} else {
		store.Set(types.LeaveKey, bz)
	}
	return nil
}
