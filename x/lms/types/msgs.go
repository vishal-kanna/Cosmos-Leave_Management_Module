package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"
)

var (
	_ sdk.Msg = &Red
	_ sdk.Msg = &MsgMultiSend{}
	_ sdk.Msg = &MsgUpdateParams{}

	_ legacytx.LegacyMsg = &MsgSend{}
	_ legacytx.LegacyMsg = &MsgMultiSend{}
	_ legacytx.LegacyMsg = &MsgUpdateParams{}
)
