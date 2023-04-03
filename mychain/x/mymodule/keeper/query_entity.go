package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"mychain/x/mymodule/types"
)

func (k Keeper) EntityAll(goCtx context.Context, req *types.QueryAllEntityRequest) (*types.QueryAllEntityResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var entitys []types.Entity
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	entityStore := prefix.NewStore(store, types.KeyPrefix(types.EntityKey))

	pageRes, err := query.Paginate(entityStore, req.Pagination, func(key []byte, value []byte) error {
		var entity types.Entity
		if err := k.cdc.Unmarshal(value, &entity); err != nil {
			return err
		}

		entitys = append(entitys, entity)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllEntityResponse{Entity: entitys, Pagination: pageRes}, nil
}

func (k Keeper) Entity(goCtx context.Context, req *types.QueryGetEntityRequest) (*types.QueryGetEntityResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	entity, found := k.GetEntity(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetEntityResponse{Entity: entity}, nil
}
