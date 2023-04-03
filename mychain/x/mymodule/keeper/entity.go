package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"mychain/x/mymodule/types"
)

// GetEntityCount get the total number of entity
func (k Keeper) GetEntityCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.EntityCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetEntityCount set the total number of entity
func (k Keeper) SetEntityCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.EntityCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendEntity appends a entity in the store with a new id and update the count
func (k Keeper) AppendEntity(
	ctx sdk.Context,
	entity types.Entity,
) uint64 {
	// Create the entity
	count := k.GetEntityCount(ctx)

	// Set the ID of the appended value
	entity.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EntityKey))
	appendedValue := k.cdc.MustMarshal(&entity)
	store.Set(GetEntityIDBytes(entity.Id), appendedValue)

	// Update entity count
	k.SetEntityCount(ctx, count+1)

	return count
}

// SetEntity set a specific entity in the store
func (k Keeper) SetEntity(ctx sdk.Context, entity types.Entity) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EntityKey))
	b := k.cdc.MustMarshal(&entity)
	store.Set(GetEntityIDBytes(entity.Id), b)
}

// GetEntity returns a entity from its id
func (k Keeper) GetEntity(ctx sdk.Context, id uint64) (val types.Entity, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EntityKey))
	b := store.Get(GetEntityIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveEntity removes a entity from the store
func (k Keeper) RemoveEntity(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EntityKey))
	store.Delete(GetEntityIDBytes(id))
}

// GetAllEntity returns all entity
func (k Keeper) GetAllEntity(ctx sdk.Context) (list []types.Entity) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EntityKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Entity
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetEntityIDBytes returns the byte representation of the ID
func GetEntityIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetEntityIDFromBytes returns ID in uint64 format from a byte array
func GetEntityIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
