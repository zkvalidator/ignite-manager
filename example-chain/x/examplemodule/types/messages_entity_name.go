package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateEntityName = "create_entity_name"
	TypeMsgUpdateEntityName = "update_entity_name"
	TypeMsgDeleteEntityName = "delete_entity_name"
)

var _ sdk.Msg = &MsgCreateEntityName{}

func NewMsgCreateEntityName(creator string, field1 string, field2 int32) *MsgCreateEntityName {
	return &MsgCreateEntityName{
		Creator: creator,
		Field1:  field1,
		Field2:  field2,
	}
}

func (msg *MsgCreateEntityName) Route() string {
	return RouterKey
}

func (msg *MsgCreateEntityName) Type() string {
	return TypeMsgCreateEntityName
}

func (msg *MsgCreateEntityName) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateEntityName) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateEntityName) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateEntityName{}

func NewMsgUpdateEntityName(creator string, id uint64, field1 string, field2 int32) *MsgUpdateEntityName {
	return &MsgUpdateEntityName{
		Id:      id,
		Creator: creator,
		Field1:  field1,
		Field2:  field2,
	}
}

func (msg *MsgUpdateEntityName) Route() string {
	return RouterKey
}

func (msg *MsgUpdateEntityName) Type() string {
	return TypeMsgUpdateEntityName
}

func (msg *MsgUpdateEntityName) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateEntityName) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateEntityName) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteEntityName{}

func NewMsgDeleteEntityName(creator string, id uint64) *MsgDeleteEntityName {
	return &MsgDeleteEntityName{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteEntityName) Route() string {
	return RouterKey
}

func (msg *MsgDeleteEntityName) Type() string {
	return TypeMsgDeleteEntityName
}

func (msg *MsgDeleteEntityName) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteEntityName) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteEntityName) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
