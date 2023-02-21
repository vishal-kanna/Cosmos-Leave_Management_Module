package keeper_test

import (
	"clms/x/lms/keeper"
	"clms/x/lms/types"
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	dbm "github.com/tendermint/tm-db"
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
func (s *TestSuite) SetupTest() {
	db := dbm.NewMemDB()

	cms := store.NewCommitMultiStore(db)
	encCfg := simapp.MakeTestEncodingConfig()
	lmsKey := sdk.NewKVStoreKey(types.StoreKey)
	ctx := testutil.DefaultContext(lmsKey, sdk.NewTransientStoreKey("transient_test"))
	keeper := keeper.NewKeeper(lmsKey, encCfg.Codec)
	cms.MountStoreWithDB(lmsKey, storetypes.StoreTypeIAVL, db)
	s.Require().NoError(cms.LoadLatestVersion())
	s.lmskeeper = keeper
	s.ctx = ctx
}

func (s *TestSuite) TestAdminRegister() {
	fmt.Println("T")
	acc := sdk.AccAddress("sadfasd")
	// s.Require().NoError(err)

	req := types.RegisterAdminRequest{
		Name:    "vishal",
		Address: acc.String(),
	}
	s.lmskeeper.AdminRegister(s.ctx, &req)
}
