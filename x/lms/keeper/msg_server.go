package keeper

// "clms/x/lms/keeper"

// "github.com/cosmos/cosmos-sdk/x/staking/types"
// v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
// "github.com/cosmos/cosmos-sdk/x/bank/types"
// "github.com/cosmos/cosmos-sdk/x/bank/types"

import (
	"clms/x/lms/types"
	"context"
	"fmt"
	"log"

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
	res := k.AddStudents(ctx, req)
	if res == true {
		return &types.AddStudentResponse{}, nil
	} else {
		return nil, types.ErrStudentAlreadyExist
	}
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
	if req.Studentid == "" {
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
	ans, err := k.ApplyLeaves(ctx, req)
	if err != nil {
		return &types.ApplyLeaveResponse{}, err
	}
	fmt.Println("the response is+++++++++++++++++++++++++++++++++", ans)
	return &ans, nil
	// return nil, nil

}
func (k Keeper) AcceptLeave(goctx context.Context, req *types.AcceptLeaveRequest) (*types.AcceptLeaveResponse, error) {
	ctx := sdk.UnwrapSDKContext(goctx)
	if req.Admin == "" {
		return nil, types.ErrAddress
	}
	if req.StudentId == "" {
		return nil, types.ErrStudentDetails
	}
	err := k.AcceptLeaves(ctx, req)
	if err != nil {
		log.Println("error is", err)
		return nil, err
	}
	return &types.AcceptLeaveResponse{}, nil
}
