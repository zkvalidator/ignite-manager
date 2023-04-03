package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"mychain/x/mymodule/types"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				EntityList: []types.Entity{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				EntityCount: 2,
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated entity",
			genState: &types.GenesisState{
				EntityList: []types.Entity{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid entity count",
			genState: &types.GenesisState{
				EntityList: []types.Entity{
					{
						Id: 1,
					},
				},
				EntityCount: 0,
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
