package keeper_test

import (
	"clms/x/lms/keeper"
	"clms/x/lms/types"
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	// "github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
)

// Keeper
type TestSuite struct {
	suite.Suite
	ctx       sdk.Context
	lmskeeper keeper.Keeper
	cdc       codec.Codec
}

func TestKeeperTestSuite(t *testing.T) {
	fmt.Println("hello1")
	suite.Run(t, new(TestSuite))
	fmt.Println("hello2")
}
func (s *TestSuite) SetUpTest() {

	fmt.Println("setup function called")
	s.ctx = sdk.Context{}
	s.cdc = &codec.ProtoCodec{}
	s.lmskeeper = keeper.NewKeeper(
		sdk.NewKVStoreKey(types.StoreKey),
		s.cdc,
	)
}

func (s *TestSuite) TestAdminRegister(t *testing.T) {
	fmt.Println("T")
	req := types.RegisterAdminRequest{
		Name:    "vishal",
		Address: "123",
	}
	ctx := sdk.Context{}
	s.lmskeeper.AdminRegister(ctx, &req)
}
