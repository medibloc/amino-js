package types

import (
	sdk "github.com/cosmos/amino-js/go/lib/panacea/panacea-sdk/types"
)

var (
	_ sdk.Msg = &MsgCreateDID{}
	_ sdk.Msg = &MsgUpdateDID{}
	_ sdk.Msg = &MsgDeactivateDID{}
)

type MsgCreateDID struct {
	DID          DID            `json:"did"`
	Document     DIDDocument    `json:"document"`
	VeriMethodID VeriMethodID   `json:"verification_method_id"`
	Signature    []byte         `json:"signature"`
	FromAddress  sdk.AccAddress `json:"from_address"`
}

type MsgUpdateDID struct {
	DID          DID            `json:"did"`
	Document     DIDDocument    `json:"document"`
	VeriMethodID VeriMethodID   `json:"verification_method_id"`
	Signature    []byte         `json:"signature"`
	FromAddress  sdk.AccAddress `json:"from_address"`
}

type MsgDeactivateDID struct {
	DID          DID            `json:"did"`
	VeriMethodID VeriMethodID   `json:"verification_method_id"`
	Signature    []byte         `json:"signature"`
	FromAddress  sdk.AccAddress `json:"from_address"`
}
