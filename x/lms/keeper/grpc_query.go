package keeper

import (
	"clms/x/lms/types"
	"context"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) ListAllTheLeaves(context.Context, *types.ListAllTheLeavesRequest) (*types.ListAllTheLeavesResponse, error) {
	return &types.ListAllTheLeavesResponse{}, nil
}
func (k Keeper) ListAllTheStudent(context.Context, *types.ListAllTheStudentRequest) (*types.ListAllTheStudentResponse, error) {
	return &types.ListAllTheStudentResponse{}, nil
}
