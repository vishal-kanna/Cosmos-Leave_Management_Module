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
	suite.Run(t, new(TestSuite))
}
func (s *TestSuite) SetupTest() {
	db := dbm.NewMemDB()
	i := 0
	fmt.Println("the i value is ", i)
	i++
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
	// fmt.Println("1")
	passcases := []struct {
		Name    string
		Address string
	}{
		{"vishal", "123"},
		{"kanna", "456"},
	}
	for _, test := range passcases {
		add := sdk.AccAddress(test.Address)
		req := types.RegisterAdminRequest{
			Name:    test.Name,
			Address: add.String(),
		}
		s.lmskeeper.AdminRegister(s.ctx, &req)
		check := s.lmskeeper.CheckAdminRegister(s.ctx, req.Address)
		if check == true {
			fmt.Println("Admin Registered")
		} else {
			fmt.Println("Admin didnot Registered")
		}
	}
	failcases := []struct {
		Name    string
		Address string
	}{
		{"vishal", "123"},
		{"kanna", "456"},
		{"saiteja", "766"},
	}
	for _, test := range failcases {
		add := sdk.AccAddress(test.Address)
		req := types.RegisterAdminRequest{
			Name:    test.Name,
			Address: add.String(),
		}
		// s.lmskeeper.AdminRegister(s.ctx, &req)
		check := s.lmskeeper.CheckAdminRegister(s.ctx, req.Address)
		if check == true {
			fmt.Println("Admin Registered")
		} else {
			fmt.Println("Admin didnot Registered")
		}
	}
}
func (s *TestSuite) TestAddStudent() {
	// fmt.Println("2")
	students := []*types.Student{
		{Address: "123", Name: "vishal", Id: "172118"},
		{Address: "456", Name: "chandini", Id: "12345"},
	}
	adminadd := sdk.AccAddress("123").String()
	req := types.AddStudentRequest{
		Admin:    adminadd,
		Students: students,
	}
	req1 := types.RegisterAdminRequest{
		Name:    "kanna",
		Address: adminadd,
	}
	s.lmskeeper.AdminRegister(s.ctx, &req1)
	s.lmskeeper.AddStudents(s.ctx, &req)
	// s.True(false)
}
func (s *TestSuite) TestApplyLeaves() {
	student := []*types.Student{
		{Address: "123", Name: "vishal", Id: "172118"},
	}
	adminadd := sdk.AccAddress("123").String()
	req := types.AddStudentRequest{
		Admin:    adminadd,
		Students: student,
	}
	req1 := types.RegisterAdminRequest{
		Name:    "kanna",
		Address: adminadd,
	}
	s.lmskeeper.AdminRegister(s.ctx, &req1)
	s.lmskeeper.AddStudents(s.ctx, &req)

}

// func (s *TestSuite) TestAdminRe() {
// 	fmt.Println("TestAdminre")
// 	tests := []struct {
// 		Name    string
// 		Address string
// 	}{
// 		{"vishal", "123"},
// 		{"kanna", "456"},
// 		{"saiteja", "766"},
// 	}
// 	for _, test := range tests {
// 		add := sdk.AccAddress(test.Address)
// 		req := types.RegisterAdminRequest{
// 			Name:    test.Name,
// 			Address: add.String(),
// 		}
// 		// s.lmskeeper.AdminRegister(s.ctx, &req)
// 		check := s.lmskeeper.CheckAdminRegister(s.ctx, req.Address)
// 		if check == true {
// 			fmt.Println("Admin Registered")
// 		} else {
// 			fmt.Println("Admin didnot Registered")
// 		}
// 	}
// }
