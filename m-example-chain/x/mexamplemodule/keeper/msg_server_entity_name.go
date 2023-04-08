package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"m-example-chain/x/mexamplemodule/types"
)

func (k msgServer) CreateEntityName(goCtx context.Context, msg *types.MsgCreateEntityName) (*types.MsgCreateEntityNameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var entityName = types.EntityName{
		Creator: msg.Creator,
		Field1:  msg.Field1,
		Field2:  msg.Field2,
	}

	id := k.AppendEntityName(
		ctx,
		entityName,
	)

	return &types.MsgCreateEntityNameResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateEntityName(goCtx context.Context, msg *types.MsgUpdateEntityName) (*types.MsgUpdateEntityNameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var entityName = types.EntityName{
		Creator: msg.Creator,
		Id:      msg.Id,
		Field1:  msg.Field1,
		Field2:  msg.Field2,
	}

	// Checks that the element exists
	val, found := k.GetEntityName(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetEntityName(ctx, entityName)

	return &types.MsgUpdateEntityNameResponse{}, nil
}

func (k msgServer) DeleteEntityName(goCtx context.Context, msg *types.MsgDeleteEntityName) (*types.MsgDeleteEntityNameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetEntityName(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveEntityName(ctx, msg.Id)

	return &types.MsgDeleteEntityNameResponse{}, nil
}
