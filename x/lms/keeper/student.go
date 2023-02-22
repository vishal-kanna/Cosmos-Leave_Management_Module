package keeper

import (
	"clms/x/lms/types"
	"fmt"
	"log"

	sdk "github.com/cosmos/cosmos-sdk/types"
	// "google.golang.org/genproto/googleapis/actions/sdk/v2"
)

//AdminRegister function take the Admin register request and
// marshalls the data and then set the based on the address
//by using this address which we used to set data ,we can also get the data
func (k Keeper) AdminRegister(ctx sdk.Context, req *types.RegisterAdminRequest) error {
	if _, err := sdk.AccAddressFromBech32(req.Address); err != nil {
		panic(fmt.Errorf("invalid admin address %w", err))
	}
	store := ctx.KVStore(k.storeKey)
	bz, err := k.cdc.Marshal(req)
	if err != nil {
		fmt.Println("error ")
		return err
	} else {
		// fmt.Println("Admin added successfullly")
		store.Set(types.AdminstoreId(req.Address), bz)
	}
	return nil
}

func (k Keeper) AcceptLeaves(ctx sdk.Context, req *types.AcceptLeaveRequest) error {
	if _, err := sdk.AccAddressFromBech32(req.Admin); err != nil {
		panic(fmt.Errorf("invalid bank authority address: %w", err))
	}
	store := ctx.KVStore(k.storeKey)

	req.Status = types.LeaveStatus_STATUS_ACCEPTED
	bz, err := k.cdc.Marshal(req)
	if err != nil {
		return err
	} else {
		store.Set(types.StudentStoreId(req.Admin), bz)
		fmt.Println("Admin is registered")
	}
	return nil
}
func (k Keeper) ApplyLeaves(ctx sdk.Context, req *types.ApplyLeaveRequest) bool {
	store := ctx.KVStore(k.storeKey)
	val := store.Get(types.StudentStoreId(req.Address))
	if val == nil {
		fmt.Println("Student did not added by the admin")
		return false
	} else {
		applyleavemarshalldata, err := k.cdc.Marshal(req)
		if err != nil {
			log.Println(err)
		} else {
			type leaveid struct {
				Id             int
				StudentAddress string
			}
			studentleave := leaveid{
				Id:             0,
				StudentAddress: req.Address,
			}
			val := store.Get(types.LeaveKeyStoreId(studentleave.Id))
			if val == nil {
				// k.cdc.Unmarshal(val)
				
			} else {

			}
			store.Set(types.LeaveStoreId(req.Address), applyleavemarshalldata)
		}
	}
}
func (k Keeper) CheckStudent(ctx sdk.Context, req *types.ApplyLeaveRequest) bool {
	store := ctx.KVStore(k.storeKey)
	val := store.Get(types.StudentStoreId(req.Address))
	if val == nil {
		fmt.Println("Student did not added by the admin")
		return false
	} else {
		return true
	}
}

// CheckAdminRegister Function Get the data from the prefixes (storeid)
//we need to pass the string with which we have added the add in store.set
//it will return true if the admin is register
//else it will return false
func (k Keeper) CheckAdminRegister(ctx sdk.Context, adminid string) bool {
	store := ctx.KVStore(k.storeKey)
	val := store.Get(types.AdminstoreId(adminid))
	// fmt.Println("the store get value is", val)
	if val == nil {
		fmt.Println("amdin didnot register")
		return false
	} else {
		res := types.RegisterAdminRequest{}
		k.cdc.Unmarshal(val, &res)
		// fmt.Println("the Response is", res)
		return true
	}
}

// func (k Keeper) CheckStudent(ctx sdk.Context, studentid string) {
// 	store := ctx.KVStore(k.storeKey)
// 	val:=
// }
func (k Keeper) AddStudents(ctx sdk.Context, req *types.AddStudentRequest) bool {
	store := ctx.KVStore(k.storeKey)
	val := store.Get(types.AdminstoreId(req.Admin))
	if val == nil {
		fmt.Println("Admin did not register ")
		return false
	} else {
		// store.Set(types.StudentStoreId(req.Studen))
		for _, val := range req.Students {
			student := store.Get(types.StudentStoreId(val.Address))
			if student == nil {
				// fmt.Println("studn")
				studentmarshall, err := k.cdc.Marshal(val)
				if err != nil {
					log.Println(err)
					return false
				}
				store.Set(types.StudentStoreId(val.Address), studentmarshall)
				fmt.Println("Student added successfully")
				return true
			} else {
				return false
			}
		}
		return true
	}
}
