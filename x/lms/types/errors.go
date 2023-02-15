package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrStudentIdDoesNotExist = sdkerrors.Register(ModuleName, 1, "Student Id Does not exist")
	ErrAdminDoesNotExist     = sdkerrors.Register(ModuleName, 2, "Admin Does not Exist")
	ErrStudentDidNotLogin    = sdkerrors.Register(ModuleName, 3, "Student Did not Login")
	ErrAdminDidNotLogin      = sdkerrors.Register(ModuleName, 4, "Admin Did not Login")
)
