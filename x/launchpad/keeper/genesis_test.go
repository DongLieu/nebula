package keeper_test

import (
	"github.com/nebula-labs/nebula/testutil/nullify"
	"github.com/nebula-labs/nebula/x/launchpad/types"
)

func (suite *KeeperTestSuite) TestGenesis() {
	projects := []*types.Project{}
	project := suite.NewBaseProject()
	projects = append(projects, &project)

	genesisState := types.GenesisState{
		Params:               types.DefaultParams(),
		Projects:             projects,
		GlobalProjectIdStart: uint64(len(projects)),
	}

	suite.querier.InitGenesis(suite.Ctx, genesisState)
	got := suite.querier.ExportGenesis(suite.Ctx)
	suite.Require().NotNil(got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)
}
