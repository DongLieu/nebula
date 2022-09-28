package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	apptesing "github.com/nebula-labs/nebula/app/apptesting"
	"github.com/nebula-labs/nebula/x/launchpad/types"
)

func (suite *KeeperTestSuite) TestTotalProjectQuery() {
	suite.SetupTest()

	// add project
	proj := suite.NewBaseProject()
	suite.querier.SetProject(suite.Ctx, proj)

	pageReq := &query.PageRequest{
		Key:        nil,
		Limit:      1,
		CountTotal: false,
	}

	// query for project
	resp, err := suite.querier.TotalProject(sdk.WrapSDKContext(suite.Ctx), &types.TotalProjectRequest{
		Pagination: pageReq,
	})
	suite.Require().NoError(err)
	suite.Require().Equal(1, len(resp.Projects))
}

func (suite *KeeperTestSuite) TestProjectBalances() {
	suite.SetupTest()

	// add project
	proj := suite.NewBaseProject()
	err := suite.querier.SetProject(suite.Ctx, proj)
	suite.Require().NoError(err)

	req := &types.ProjectBalancesRequest{
		ProjectId: proj.ProjectId,
	}

	// add money to project
	suite.App.BankKeeper.SendCoins(suite.Ctx, suite.TestAccs[0], sdk.AccAddress(proj.ProjectAddress), sdk.NewCoins(sdk.NewCoin(apptesing.TokenDenom, sdk.NewInt(100000000))))
	suite.App.BankKeeper.SendCoins(suite.Ctx, suite.TestAccs[0], sdk.AccAddress(proj.ProjectAddress), sdk.NewCoins(sdk.NewCoin(apptesing.StableDenom, sdk.NewInt(100000000))))

	// query for project
	resp, err := suite.querier.ProjectBalances(sdk.WrapSDKContext(suite.Ctx), req)
	suite.Require().NoError(err)
	suite.Require().Equal(2, len(resp.Balances))
	suite.Require().Equal(sdk.NewInt(100000000), resp.Balances.AmountOf(apptesing.TokenDenom))
	suite.Require().Equal(sdk.NewInt(100000000), resp.Balances.AmountOf(apptesing.StableDenom))
}
