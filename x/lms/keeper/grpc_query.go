package keeper

import (
	"clms/x/lms/types"
	"context"
	"fmt"

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
	res := k.GetAllStudents(ctx, req)
	res1 := types.ListAllTheStudentResponse{
		Students: res,
	}
	return &res1, nil
}
