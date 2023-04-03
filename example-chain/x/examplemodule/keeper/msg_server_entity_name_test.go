package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"example-chain/x/examplemodule/types"
)

func TestEntityNameMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateEntityName(ctx, &types.MsgCreateEntityName{Creator: creator})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

func TestEntityNameMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateEntityName
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateEntityName{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateEntityName{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateEntityName{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.CreateEntityName(ctx, &types.MsgCreateEntityName{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdateEntityName(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestEntityNameMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteEntityName
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteEntityName{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteEntityName{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteEntityName{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)

			_, err := srv.CreateEntityName(ctx, &types.MsgCreateEntityName{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteEntityName(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
