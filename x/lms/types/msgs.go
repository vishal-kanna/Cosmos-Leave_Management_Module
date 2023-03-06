package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	// "go.starlark.net/lib/time"
	// "x/lms/types"
)

var (
	_ sdk.Msg = &RegisterAdminRequest{}
	// _ sdk.Msg = &RegisterAdminResponse{}
	_ sdk.Msg = &AddStudentRequest{}
	_ sdk.Msg = &ApplyLeaveRequest{}
	_ sdk.Msg = &AcceptLeaveRequest{}
)

func NewRegisterAdminRequest(signer string, adminadd string, name string) *RegisterAdminRequest {
	return &RegisterAdminRequest{Signer: signer, Address: adminadd, Name: name}
}

func (msg RegisterAdminRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Address); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid from address: %s", err)
	}
	if msg.Name == "" {
		return ErrStudentNameDoesNotExist
	}
	return nil
}

func (msg RegisterAdminRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}
func (msg RegisterAdminRequest) GetSigners() []sdk.AccAddress {
	fromAddress, _ := sdk.AccAddressFromBech32(msg.Signer)
	return []sdk.AccAddress{fromAddress}
}
func NewAddStudentRequest(signer string, admin string, students []*Student) *AddStudentRequest {
	return &AddStudentRequest{Signer: signer, Admin: admin, Students: students}
}
func (msg AddStudentRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Signer); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("Invalid Address")
	}
	for i := 0; i < len(msg.Students); i++ {
		if msg.Students[i].Name == "" && msg.Students[i].Id == "" {
			return ErrStudentDetails
		}
	}
	return nil
}
func (msg AddStudentRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}
func (msg AddStudentRequest) GetSigners() []sdk.AccAddress {
	fromAddress, _ := sdk.AccAddressFromBech32(msg.Signer)
	return []sdk.AccAddress{fromAddress}
}

func NewApplyLeaveRequest(signer string, address string, reason string, from time.Time, to time.Time) *ApplyLeaveRequest {
	return &ApplyLeaveRequest{Signer: signer, Address: address, Reason: reason, From: &from, To: &to}
}
func (msg ApplyLeaveRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Signer); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("Invalid Address")
	}
	if len(msg.Reason) == 0 {
		return ErrEmptyReason
	}
	return nil
}
func (msg ApplyLeaveRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}
func (msg ApplyLeaveRequest) GetSigners() []sdk.AccAddress {
	fromAddress, _ := sdk.AccAddressFromBech32(msg.Signer)
	return []sdk.AccAddress{fromAddress}
}
func NewAcceptLeaveRequest(signer string, adminaddress string, studentid string) *AcceptLeaveRequest {
	return &AcceptLeaveRequest{Signer: signer, Admin: adminaddress, StudentId: studentid}
}

func (msg AcceptLeaveRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}
func (msg AcceptLeaveRequest) GetSigners() []sdk.AccAddress {
	fromAddress, _ := sdk.AccAddressFromBech32(msg.Signer)
	return []sdk.AccAddress{fromAddress}
}
func (msg AcceptLeaveRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Signer); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("Invalid admin Address")
	}
	if _, err := sdk.AccAddressFromBech32(msg.Signer); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("Invalid student  Address")
	}
	return nil
}

// func (msg )
