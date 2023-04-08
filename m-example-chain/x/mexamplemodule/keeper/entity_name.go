package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"m-example-chain/x/mexamplemodule/types"
)

// GetEntityNameCount get the total number of entityName
func (k Keeper) GetEntityNameCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.EntityNameCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetEntityNameCount set the total number of entityName
func (k Keeper) SetEntityNameCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.EntityNameCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendEntityName appends a entityName in the store with a new id and update the count
func (k Keeper) AppendEntityName(
	ctx sdk.Context,
	entityName types.EntityName,
) uint64 {
	// Create the entityName
	count := k.GetEntityNameCount(ctx)

	// Set the ID of the appended value
	entityName.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EntityNameKey))
	appendedValue := k.cdc.MustMarshal(&entityName)
	store.Set(GetEntityNameIDBytes(entityName.Id), appendedValue)

	// Update entityName count
	k.SetEntityNameCount(ctx, count+1)

	return count
}

// SetEntityName set a specific entityName in the store
func (k Keeper) SetEntityName(ctx sdk.Context, entityName types.EntityName) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EntityNameKey))
	b := k.cdc.MustMarshal(&entityName)
	store.Set(GetEntityNameIDBytes(entityName.Id), b)
}

// GetEntityName returns a entityName from its id
func (k Keeper) GetEntityName(ctx sdk.Context, id uint64) (val types.EntityName, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EntityNameKey))
	b := store.Get(GetEntityNameIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveEntityName removes a entityName from the store
func (k Keeper) RemoveEntityName(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EntityNameKey))
	store.Delete(GetEntityNameIDBytes(id))
}

// GetAllEntityName returns all entityName
func (k Keeper) GetAllEntityName(ctx sdk.Context) (list []types.EntityName) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EntityNameKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.EntityName
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetEntityNameIDBytes returns the byte representation of the ID
func GetEntityNameIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetEntityNameIDFromBytes returns ID in uint64 format from a byte array
func GetEntityNameIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
