package mymodule

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"mychain/testutil/sample"
	mymodulesimulation "mychain/x/mymodule/simulation"
	"mychain/x/mymodule/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = mymodulesimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateEntity = "op_weight_msg_entity"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateEntity int = 100

	opWeightMsgUpdateEntity = "op_weight_msg_entity"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateEntity int = 100

	opWeightMsgDeleteEntity = "op_weight_msg_entity"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteEntity int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	mymoduleGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		EntityList: []types.Entity{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		EntityCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&mymoduleGenesis)
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

	var weightMsgCreateEntity int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateEntity, &weightMsgCreateEntity, nil,
		func(_ *rand.Rand) {
			weightMsgCreateEntity = defaultWeightMsgCreateEntity
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateEntity,
		mymodulesimulation.SimulateMsgCreateEntity(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateEntity int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateEntity, &weightMsgUpdateEntity, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateEntity = defaultWeightMsgUpdateEntity
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateEntity,
		mymodulesimulation.SimulateMsgUpdateEntity(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteEntity int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteEntity, &weightMsgDeleteEntity, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteEntity = defaultWeightMsgDeleteEntity
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteEntity,
		mymodulesimulation.SimulateMsgDeleteEntity(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
