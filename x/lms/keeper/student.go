package keeper

import (
	"clms/x/lms/types"
	"fmt"
	"log"
	"strconv"

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
		return err
	}
	store.Set(types.AdminstoreId(req.Address), bz)

	return nil
}
func (k Keeper) ApplyLeaves(ctx sdk.Context, req *types.ApplyLeaveRequest) bool {
	store := ctx.KVStore(k.storeKey)
	//checking whether student is added by the admin or not
	//if student is added then we will get the byte array
	val := store.Get(types.StudentStoreId(req.Studentid))
	if val == nil {
		fmt.Println("Student did not added by the admin")
		return false
	}
	leavecounter := store.Get(types.LeaveCounterStoreId(req.Studentid))
	if leavecounter == nil {
		fmt.Println("leavecounter============", leavecounter)
		//if leavecounter nil means there is no counter in the store
		c := 1
		InttoString1 := strconv.Itoa(c)
		store.Set(types.LeaveCounterStoreId(req.Studentid), []byte(InttoString1))
	} else {
		fmt.Println("leavecounter============", leavecounter)

		leavecounterstring := string(leavecounter)
		fmt.Println("leavecounter============", leavecounterstring)

		// a, err := strconv.Atoi(ans)
		leavecounterint, err := strconv.Atoi(leavecounterstring)
		if err != nil {
			panic(err)
		}
		leavecounterint = leavecounterint + 1
		store.Set(types.LeaveCounterStoreId(req.Studentid), []byte(fmt.Sprint(leavecounterint)))
	}
	counter := store.Get(types.LeaveCounterStoreId(req.Studentid))
	counterint, err := strconv.Atoi(string(counter))
	fmt.Println("leavecounter============counterint", counterint)

	if err != nil {
		panic(err)
	}
	req1 := &types.LeaveRequest{
		Leaveid: int64(counterint),
		Address: req.Studentid,
		Reason:  req.Reason,
		Signer:  req.Signer,
		From:    req.From,
		To:      req.To,
		Status:  types.LeaveStatus_STATUS_UNDEFINED,
	}
	fmt.Println("the req is", req1)
	data, err := k.cdc.Marshal(req1)
	leavecounterstring := strconv.Itoa(counterint)
	store.Set(types.LeaveStorinKeyId(req.Studentid, leavecounterstring), data)

	return true
}

func (k Keeper) AcceptLeaves(ctx sdk.Context, req *types.AcceptLeaveRequest) error {
	if _, err := sdk.AccAddressFromBech32(req.Signer); err != nil {
		panic(fmt.Errorf("invalid bank authority address: %w", err))
	}
	store := ctx.KVStore(k.storeKey)
	adminpresent := store.Get(types.AdminstoreId(req.Admin))
	if adminpresent == nil {
		return types.ErrAdminDidNotLogin
	} else {
		marshaldata := store.Get(types.LeaveCounterStoreId(req.StudentId))
		if marshaldata == nil {
			fmt.Println("student did not request leave")
		} else {
			reqdata := store.Get(types.LeaveStorinKeyId(req.StudentId, req.LeaveId))
			var acceptleave types.LeaveRequest
			k.cdc.Unmarshal(reqdata, &acceptleave)
			acceptleave.Status = req.Status
			d, e := k.cdc.Marshal(&acceptleave)
			if e != nil {
				panic(e)
			}
			store.Set(types.LeaveStorinKeyId(req.StudentId, req.LeaveId), d)
			store.Set(types.AllLeavesStoreId(string(reqdata)), d)
		}
	}
	return nil
}

func (k Keeper) CheckStudent(ctx sdk.Context, req *types.ApplyLeaveRequest) bool {
	store := ctx.KVStore(k.storeKey)
	val := store.Get(types.StudentStoreId(req.Studentid))
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

func (k Keeper) AddStudents(ctx sdk.Context, req *types.AddStudentRequest) bool {
	store := ctx.KVStore(k.storeKey)
	val := store.Get(types.AdminstoreId(req.Admin))
	if val == nil {
		return false
	} else {
		for _, val1 := range req.Students {
			student := store.Get(types.StudentStoreId(val1.Id))
			if student == nil {
				studentmarshall, err := k.cdc.Marshal(val1)
				if err != nil {
					log.Println(err)
					return false
				}
				store.Set(types.StudentStoreId(val1.Id), studentmarshall)
			} else {
				return false
			}
		}
		return true
	}
}
func (k Keeper) GetAllStudents(ctx sdk.Context, req *types.ListAllTheStudentRequest) []*types.Student {
	store := ctx.KVStore(k.storeKey)

	var students []*types.Student
	itr := sdk.KVStorePrefixIterator(store, types.StudentKey)
	defer itr.Close()

	for ; itr.Valid(); itr.Next() {
		var student types.Student
		k.cdc.Unmarshal(itr.Value(), &student)
		students = append(students, &student)
	}
	return students
}
func (k Keeper) GetAllLeaves(ctx sdk.Context, req *types.ListAllTheLeavesRequest) []*types.LeaveRequest {
	store := ctx.KVStore(k.storeKey)
	var leaves []*types.LeaveRequest
	itr := sdk.KVStorePrefixIterator(store, types.LeavesStoringKey)
	for ; itr.Valid(); itr.Next() {
		var leave types.LeaveRequest
		k.cdc.Unmarshal(itr.Value(), &leave)
		leaves = append(leaves, &leave)
	}
	return leaves
}

func (k Keeper) GetleaveStatus(ctx sdk.Context, studentaddress string) types.AcceptLeaveRequest {
	store := ctx.KVStore(k.storeKey)
	var leave types.AcceptLeaveRequest
	res := store.Get(types.AllLeavesStoreId(studentaddress))
	if res == nil {
		fmt.Println("no results")
	} else {
		k.cdc.Unmarshal(res, &leave)
	}
	return leave
}
func (k Keeper) Getadmin(ctx sdk.Context, req *types.GetadminRequest) []*types.RegisterAdminRequest {
	store := ctx.KVStore(k.storeKey)
	var admin []*types.RegisterAdminRequest
	itr := sdk.KVStorePrefixIterator(store, types.AdminKey)
	for ; itr.Valid(); itr.Next() {
		var adminn types.RegisterAdminRequest
		k.cdc.Unmarshal(itr.Value(), &adminn)
		admin = append(admin, &adminn)
	}
	return admin
}
