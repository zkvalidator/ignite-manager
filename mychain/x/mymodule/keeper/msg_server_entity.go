package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"mychain/x/mymodule/types"
)

func (k msgServer) CreateEntity(goCtx context.Context, msg *types.MsgCreateEntity) (*types.MsgCreateEntityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var entity = types.Entity{
		Creator: msg.Creator,
		Field1:  msg.Field1,
		Field2:  msg.Field2,
	}

	id := k.AppendEntity(
		ctx,
		entity,
	)

	return &types.MsgCreateEntityResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateEntity(goCtx context.Context, msg *types.MsgUpdateEntity) (*types.MsgUpdateEntityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var entity = types.Entity{
		Creator: msg.Creator,
		Id:      msg.Id,
		Field1:  msg.Field1,
		Field2:  msg.Field2,
	}

	// Checks that the element exists
	val, found := k.GetEntity(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetEntity(ctx, entity)

	return &types.MsgUpdateEntityResponse{}, nil
}

func (k msgServer) DeleteEntity(goCtx context.Context, msg *types.MsgDeleteEntity) (*types.MsgDeleteEntityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetEntity(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveEntity(ctx, msg.Id)

	return &types.MsgDeleteEntityResponse{}, nil
}
