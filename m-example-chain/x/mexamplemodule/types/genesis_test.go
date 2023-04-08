package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"m-example-chain/x/mexamplemodule/types"
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

				EntityNameList: []types.EntityName{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				EntityNameCount: 2,
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated entityName",
			genState: &types.GenesisState{
				EntityNameList: []types.EntityName{
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
			desc: "invalid entityName count",
			genState: &types.GenesisState{
				EntityNameList: []types.EntityName{
					{
						Id: 1,
					},
				},
				EntityNameCount: 0,
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
