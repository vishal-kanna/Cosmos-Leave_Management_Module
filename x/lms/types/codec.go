package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&RegisterAdminRequest{},
		&AddStudentRequest{},
		&ApplyLeaveRequest{},
		&AcceptLeaveRequest{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	legacy.RegisterAminoMsg(cdc, &RegisterAdminRequest{}, "clms/RegisterAdminRequest")
	legacy.RegisterAminoMsg(cdc, &AddStudentRequest{}, "clms/AddStudentRequest")
	legacy.RegisterAminoMsg(cdc, &ApplyLeaveRequest{}, "clms/ApplyLeaveRequest")
	legacy.RegisterAminoMsg(cdc, &AcceptLeaveRequest{}, "clms/AcceptLeaveRequest")
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	sdk.RegisterLegacyAminoCodec(amino)
}
