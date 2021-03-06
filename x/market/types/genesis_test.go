package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/txlabs/blockless-chain/x/market/types"
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

				OrderList: []types.Order{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				CompletedOrderList: []types.CompletedOrder{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				ClaimedOrderList: []types.ClaimedOrder{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated order",
			genState: &types.GenesisState{
				OrderList: []types.Order{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated completedOrder",
			genState: &types.GenesisState{
				CompletedOrderList: []types.CompletedOrder{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated claimedOrder",
			genState: &types.GenesisState{
				ClaimedOrderList: []types.ClaimedOrder{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
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
