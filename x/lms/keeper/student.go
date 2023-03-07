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
		// fmt.Println("error ")
		return err
	}
	// } else {
	// fmt.Println("Admin added successfullly")
	store.Set(types.AdminstoreId(req.Address), bz)

	return nil
}

func (k Keeper) AcceptLeaves(ctx sdk.Context, req *types.AcceptLeaveRequest) error {
	if _, err := sdk.AccAddressFromBech32(req.Signer); err != nil {
		panic(fmt.Errorf("invalid bank authority address: %w", err))
	}
	store := ctx.KVStore(k.storeKey)
	adminpresent := store.Get(types.AdminstoreId(req.Admin))
	if adminpresent == nil {
		req.Status = types.LeaveStatus_STATUS_UNDEFINED
		_, err := k.cdc.Marshal(req)
		if err != nil {
			log.Println(err)
		} else {
			//we need to check whether the student is requested leave or not
			//as well as we need to check whether the student is added by the register admin
			val := store.Get(types.StudentStoreId(req.StudentId))
			if val == nil {
				return types.ErrStudentDidNotLogin
				// fmt.Println("store not present in the student store")
			}
			leaveval := store.Get(types.LeaveStoreId(req.StudentId))
			if leaveval == nil {
				fmt.Println("student did not request the leave")
			}
		}
		// store.Set(types.StudentStoreId())
	} else {
		// panic("called")
		//if admin is present then we need to check the student is present in student store or not
		//if the student is present in store then we need create a new response of accepting the leave
		//and storing in the allleavestore
		req.Status = types.LeaveStatus_STATUS_ACCEPTED
		marshalaccepteddata, err := k.cdc.Marshal(req)
		if err != nil {
			log.Println(err)
		}
		marshaldata := store.Get(types.LeaveStoreId(req.StudentId))
		if marshaldata == nil {
			fmt.Println("student did not request leave")
		} else {
			store.Set(types.AllLeavesStoreId(req.StudentId), marshalaccepteddata)
		}
	}
	return nil
}
func (k Keeper) ApplyLeaves(ctx sdk.Context, req *types.ApplyLeaveRequest) bool {
	store := ctx.KVStore(k.storeKey)
	//checking whether student is added by the admin or not
	//if student is added then we will get the byte array
	val := store.Get(types.StudentStoreId(req.Address))
	if val == nil {
		fmt.Println("Student did not added by the admin")
		return false
	} else {
		applyleavemarshalldata, err := k.cdc.Marshal(req)
		if err != nil {
			log.Println(err)
		} else {
			var counter int
			counter = 0
			//val stores the byte array if the corressponding key is present in the prefix LeaveKeyId
			val := store.Get(types.LeaveKeyStoreId(req.Address))
			if val == nil {
				counter++
				//marshall the counter
				InttoString := strconv.Itoa(counter)
				// marhshallcounter, err := k.cdc.Marshal(InttoString)

				//to better understand LeaveStoreId store the respective studnet leave data
				store.Set(types.LeaveStoreId(req.Address), applyleavemarshalldata)
				//Need to change this because here Im manually converting the string into the
				//to better understand LeaveKeyStoreId store the respective studnet counter

				store.Set(types.LeaveKeyStoreId(req.Address), []byte(InttoString))
			} else {
				//if we do not get the nil means the student address already present in the array
				//so we need to get the student id from the leaveKeyStoreId and increment the counter
				val := store.Get(types.LeaveKeyStoreId(req.Address))
				ans := string(val)
				a, err := strconv.Atoi(ans)
				if err != nil {
					log.Println(err)
				} else {
					a++
					//convert it to string
					s := strconv.Itoa(a)
					store.Set(types.LeaveKeyStoreId(req.Address), []byte(s))
					store.Set(types.LeaveStoreId(req.Address), applyleavemarshalldata)
				}
				// k.cdc.Unmarshal(val, &res)
			}
			store.Set(types.LeaveStoreId(req.Address), applyleavemarshalldata)
		}
	}
	return true
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

func (k Keeper) AddStudents(ctx sdk.Context, req *types.AddStudentRequest) bool {
	store := ctx.KVStore(k.storeKey)
	val := store.Get(types.AdminstoreId(req.Admin))
	if val == nil {
		fmt.Println("Admin did not register ")
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
				// panic("called2")
				fmt.Println("Student added successfully")
				fmt.Println("Student added successfully")

				fmt.Println("Student added successfully")
				fmt.Println("Student added successfully")

				return true
			} else {
				return false
			}
		}
		return true
	}
}
func (k Keeper) GetAllStudents(ctx sdk.Context, req *types.ListAllTheStudentRequest) []*types.Student {
	// panic("called1")
	// fmt.Println("hey im in getallstudents keeper methods ")
	store := ctx.KVStore(k.storeKey)
	// fmt.Println("hey im in getallstudents keeper methods 5")

	var students []*types.Student
	fmt.Println("hey im in getallstudents keeper methods 6")
	itr := sdk.KVStorePrefixIterator(store, types.StudentKey)
	defer itr.Close()
	fmt.Println("iterator function ", itr)
	fmt.Println("hey im in getallstudents keeper methods 1")

	for ; itr.Valid(); itr.Next() {
		var student types.Student
		fmt.Println("hey im in getallstudents keeper methods 2")
		k.cdc.Unmarshal(itr.Value(), &student)
		students = append(students, &student)
	}
	fmt.Println("the students details are ==========================================", students)

	return students
}
func (k Keeper) GetAllLeaves(ctx sdk.Context, req *types.ListAllTheLeavesRequest) []*types.AcceptLeaveRequest {
	store := ctx.KVStore(k.storeKey)
	var leaves []*types.AcceptLeaveRequest
	itr := sdk.KVStorePrefixIterator(store, types.AllLeavesKey)
	for ; itr.Valid(); itr.Next() {
		var leave types.AcceptLeaveRequest
		k.cdc.Unmarshal(itr.Value(), &leave)
		leaves = append(leaves, &leave)
	}
	return leaves
}

func (k Keeper) GetleaveStatus(ctx sdk.Context, studentaddress string) types.AcceptLeaveRequest {
	store := ctx.KVStore(k.storeKey)
	var leave types.AcceptLeaveRequest
	res := store.Get(types.AllLeavesStoreId(studentaddress))
	// itr := sdk.KVStorePrefixIterator(store, types.StudentKey)
	if res == nil {
		fmt.Println("no results")
	} else {
		k.cdc.Unmarshal(res, &leave)
	}
	return leave
}
func (k Keeper) Getadmin(ctx sdk.Context, adminid string) types.RegisterAdminRequest {
	store := ctx.KVStore(k.storeKey)
	val := store.Get(types.AdminstoreId(adminid))
	if val == nil {
		return types.RegisterAdminRequest{}
	} else {
		res := types.RegisterAdminRequest{}
		k.cdc.Unmarshal(val, &res)
		return res
	}
}
