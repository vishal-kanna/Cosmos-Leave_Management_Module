package types

import (
	"clms/x/lms/keeper"
	"clms/x/lms/types"
	// "github.com/cosmos/cosmos-sdk/x/staking/types"
	// v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	// "github.com/cosmos/cosmos-sdk/x/bank/types"
	// "github.com/cosmos/cosmos-sdk/x/bank/types"
)

type msgServer struct {
	keeper.Keeper
}

var _ types.MsgServer = msgServer{}
