package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateEntity = "create_entity"
	TypeMsgUpdateEntity = "update_entity"
	TypeMsgDeleteEntity = "delete_entity"
)

var _ sdk.Msg = &MsgCreateEntity{}

func NewMsgCreateEntity(creator string, field1 string, field2 int32) *MsgCreateEntity {
	return &MsgCreateEntity{
		Creator: creator,
		Field1:  field1,
		Field2:  field2,
	}
}

func (msg *MsgCreateEntity) Route() string {
	return RouterKey
}

func (msg *MsgCreateEntity) Type() string {
	return TypeMsgCreateEntity
}

func (msg *MsgCreateEntity) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateEntity) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateEntity) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateEntity{}

func NewMsgUpdateEntity(creator string, id uint64, field1 string, field2 int32) *MsgUpdateEntity {
	return &MsgUpdateEntity{
		Id:      id,
		Creator: creator,
		Field1:  field1,
		Field2:  field2,
	}
}

func (msg *MsgUpdateEntity) Route() string {
	return RouterKey
}

func (msg *MsgUpdateEntity) Type() string {
	return TypeMsgUpdateEntity
}

func (msg *MsgUpdateEntity) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateEntity) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateEntity) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteEntity{}

func NewMsgDeleteEntity(creator string, id uint64) *MsgDeleteEntity {
	return &MsgDeleteEntity{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteEntity) Route() string {
	return RouterKey
}

func (msg *MsgDeleteEntity) Type() string {
	return TypeMsgDeleteEntity
}

func (msg *MsgDeleteEntity) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteEntity) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteEntity) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
