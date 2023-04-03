package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"mychain/x/mymodule/types"
)

func TestEntityMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateEntity(ctx, &types.MsgCreateEntity{Creator: creator})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

func TestEntityMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateEntity
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateEntity{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateEntity{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateEntity{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.CreateEntity(ctx, &types.MsgCreateEntity{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdateEntity(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestEntityMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteEntity
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteEntity{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteEntity{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteEntity{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)

			_, err := srv.CreateEntity(ctx, &types.MsgCreateEntity{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteEntity(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
