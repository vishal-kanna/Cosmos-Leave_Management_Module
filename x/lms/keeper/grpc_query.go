package keeper

import (
	"clms/x/lms/types"
	"context"
	"fmt"
	"log"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) ListAllTheLeaves(goCtx context.Context, req *types.ListAllTheLeavesRequest) (*types.ListAllTheLeavesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	res := k.GetAllLeaves(ctx, req)
	fmt.Println("the leaves array is", res)

	// res1 := &types.ListAllTheLeavesResponse{}
	// // res1.Leaves = res
	return &types.ListAllTheLeavesResponse{
		Leaves: res,
	}, nil
}
func (k Keeper) ListAllTheStudent(goCtx context.Context, req *types.ListAllTheStudentRequest) (*types.ListAllTheStudentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	log.Println("======================")
	res := k.GetAllStudents(ctx, req)
	fmt.Println("the res is++++++++++++++++++++++++", res)
	res1 := types.ListAllTheStudentResponse{
		Students: res,
	}
	return &res1, nil
}
func (k Keeper) GetLeaveStatus(goctx context.Context, req *types.GetLeaveStatusRequest) (*types.GetLeaveStatusResponse, error) {
	ctx := sdk.UnwrapSDKContext(goctx)
	res := k.GetleaveStatus(ctx, req.Studentaddress)
	res1 := types.GetLeaveStatusResponse{
		Status: &res,
	}
	return &res1, nil
}
func (k Keeper) GetAdmin(goctx context.Context, req *types.GetadminRequest) (*types.GetadminResponse, error) {
	ctx := sdk.UnwrapSDKContext(goctx)
	res := k.Getadmin(ctx, req)
	res1 := types.GetadminResponse{
		Admins: res,
	}
	return &res1, nil
}
