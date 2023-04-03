package keeper

import (
	"context"

	"example-chain/x/examplemodule/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) EntityNameAll(goCtx context.Context, req *types.QueryAllEntityNameRequest) (*types.QueryAllEntityNameResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var entityNames []types.EntityName
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	entityNameStore := prefix.NewStore(store, types.KeyPrefix(types.EntityNameKey))

	pageRes, err := query.Paginate(entityNameStore, req.Pagination, func(key []byte, value []byte) error {
		var entityName types.EntityName
		if err := k.cdc.Unmarshal(value, &entityName); err != nil {
			return err
		}

		entityNames = append(entityNames, entityName)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllEntityNameResponse{EntityName: entityNames, Pagination: pageRes}, nil
}

func (k Keeper) EntityName(goCtx context.Context, req *types.QueryGetEntityNameRequest) (*types.QueryGetEntityNameResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	entityName, found := k.GetEntityName(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetEntityNameResponse{EntityName: entityName}, nil
}
