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

var _ types.MsgServer = msgServer{}

func (k msgServer) AddStudent(ctx context.Context, req *types.AddStudentRequest) (*types.AddStudentResponse, error) {
	return &types.AddStudentResponse{}, nil
}

func (k msgServer) RegisterAdmin(goctx context.Context, req *types.RegisterAdminRequest) (*types.RegisterAdminResponse, error) {
	ctx := sdk.UnwrapSDKContext(goctx)
	// err := Keeper.AdminRegister(ctx, &req)
	k.AdminRegister(ctx, req)

	return &types.RegisterAdminResponse{}, nil
}

func (k msgServer) ApplyLeave(goctx context.Context, req *types.ApplyLeaveRequest) (*types.ApplyLeaveResponse, error) {
	// k.AcceptLeave(ctx, req)
	ctx := sdk.UnwrapSDKContext(goctx)
	k.ApplyLeaves(ctx, req)
	return &types.ApplyLeaveResponse{}, nil
}
func (k msgServer) AcceptLeave(goctx context.Context, req *types.AcceptLeaveRequest) (*types.AcceptLeaveResponse, error) {
	ctx := sdk.UnwrapSDKContext(goctx)
	k.AcceptLeaves(ctx, req)
	return &types.AcceptLeaveResponse{}, nil
}
