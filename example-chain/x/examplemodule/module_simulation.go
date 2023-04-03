package examplemodule

import (
	"math/rand"

	"example-chain/testutil/sample"
	examplemodulesimulation "example-chain/x/examplemodule/simulation"
	"example-chain/x/examplemodule/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = examplemodulesimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateEntityName = "op_weight_msg_entity_name"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateEntityName int = 100

	opWeightMsgUpdateEntityName = "op_weight_msg_entity_name"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateEntityName int = 100

	opWeightMsgDeleteEntityName = "op_weight_msg_entity_name"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteEntityName int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	examplemoduleGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		EntityNameList: []types.EntityName{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		EntityNameCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&examplemoduleGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateEntityName int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateEntityName, &weightMsgCreateEntityName, nil,
		func(_ *rand.Rand) {
			weightMsgCreateEntityName = defaultWeightMsgCreateEntityName
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateEntityName,
		examplemodulesimulation.SimulateMsgCreateEntityName(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateEntityName int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateEntityName, &weightMsgUpdateEntityName, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateEntityName = defaultWeightMsgUpdateEntityName
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateEntityName,
		examplemodulesimulation.SimulateMsgUpdateEntityName(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteEntityName int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteEntityName, &weightMsgDeleteEntityName, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteEntityName = defaultWeightMsgDeleteEntityName
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteEntityName,
		examplemodulesimulation.SimulateMsgDeleteEntityName(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
