package keeper

import (
	"clms/x/lms/types"
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) ListAllTheLeaves(goCtx context.Context, req *types.ListAllTheLeavesRequest) (*types.ListAllTheLeavesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	k.GetAllLeaves(ctx, req)
	return &types.ListAllTheLeavesResponse{}, nil
}
func (k Keeper) ListAllTheStudent(goCtx context.Context, req *types.ListAllTheStudentRequest) (*types.ListAllTheStudentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	k.GetAllStudents(ctx, req)
	return &types.ListAllTheStudentResponse{}, nil
}
