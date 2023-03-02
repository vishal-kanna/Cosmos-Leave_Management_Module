package keeper

// "clms/x/lms/keeper"

// "github.com/cosmos/cosmos-sdk/x/staking/types"
// v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
// "github.com/cosmos/cosmos-sdk/x/bank/types"
// "github.com/cosmos/cosmos-sdk/x/bank/types"

import (
	"clms/x/lms/types"
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgServer struct {
	Keeper
}

// var _ types.MsgServer = msgServer{}

var _ types.MsgServer = Keeper{}

func (k Keeper) AddStudent(goctx context.Context, req *types.AddStudentRequest) (*types.AddStudentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goctx)

	if req.Admin == "" {
		return nil, types.ErrAddress
	}
	k.AddStudents(ctx, req)
	return &types.AddStudentResponse{}, nil
}

func (k Keeper) RegisterAdmin(goctx context.Context, req *types.RegisterAdminRequest) (*types.RegisterAdminResponse, error) {
	ctx := sdk.UnwrapSDKContext(goctx)
	// err := Keeper.AdminRegister(ctx, &req)
	if req.Address == "" {
		return nil, types.ErrAddress
	}
	if req.Name == "" {
		return nil, types.ErrNameCantBeEmpty
	}
	err := k.AdminRegister(ctx, req)
	if err != nil {
		return nil, err
	}
	return &types.RegisterAdminResponse{}, nil
}

func (k Keeper) ApplyLeave(goctx context.Context, req *types.ApplyLeaveRequest) (*types.ApplyLeaveResponse, error) {
	// k.AcceptLeave(ctx, req)
	ctx := sdk.UnwrapSDKContext(goctx)
	if req.Address == "" {
		return nil, types.ErrAddress
	}
	if req.From == nil {
		return nil, types.ErrDate
	}
	if req.To == nil {
		return nil, types.ErrDate
	}
	if req.Reason == "" {
		return nil, types.ErrEmptyReason
	}
	ans := k.ApplyLeaves(ctx, req)
	if ans == true {
		return &types.ApplyLeaveResponse{}, nil
	}
	// return nil, false
	return nil, nil

}
func (k Keeper) AcceptLeave(goctx context.Context, req *types.AcceptLeaveRequest) (*types.AcceptLeaveResponse, error) {
	ctx := sdk.UnwrapSDKContext(goctx)
	if req.Admin == "" {
		return nil, types.ErrAddress
	}
	if req.StudentId == "" {
		return nil, types.ErrStudentDetails
	}
	k.AcceptLeaves(ctx, req)
	return &types.AcceptLeaveResponse{}, nil
}
