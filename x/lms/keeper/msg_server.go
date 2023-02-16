package keeper

// "clms/x/lms/keeper"

// "github.com/cosmos/cosmos-sdk/x/staking/types"
// v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
// "github.com/cosmos/cosmos-sdk/x/bank/types"
// "github.com/cosmos/cosmos-sdk/x/bank/types"

import (
	"clms/x/lms/types"
	"context"
)

type msgServer struct {
	Keeper
	StudentKeeper
}

// var _ types.MsgServer = msgServer{}

var _ types.MsgServer = msgServer{}

func (k msgServer) AddStudent(ctx context.Context, req *types.AddStudentRequest) (*types.AddStudentResponse, error) {
	return &types.AddStudentResponse{}, nil
}

func (k msgServer) RegisterAdmin(ctx context.Context, req *types.RegisterAdminRequest) (*types.RegisterAdminResponse, error) {
	return &types.RegisterAdminResponse{}, nil
}
func (k msgServer) ApplyLeave(ctx context.Context, req *types.ApplyLeaveRequest) (*types.ApplyLeaveResponse, error) {
	// k.AcceptLeave(ctx, req)
	StudentKeeper.AcceptLeaves(ctx, req)
	return &types.ApplyLeaveResponse{}, nil
}
func (k msgServer) AcceptLeave(ctx context.Context, req *types.AcceptLeaveRequest) (*types.AcceptLeaveResponse, error) {
	return &types.AcceptLeaveResponse{}, nil
}
