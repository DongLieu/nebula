package keeper_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	apptesing "github.com/nebula-labs/nebula/app/apptesting"
	"github.com/nebula-labs/nebula/x/ido/keeper"
	"github.com/nebula-labs/nebula/x/ido/types"
	launchpad "github.com/nebula-labs/nebula/x/launchpad/types"
	"github.com/stretchr/testify/suite"
)

var (
	TokenForDistribution = sdk.NewCoin(apptesing.TokenDenom, sdk.NewInt(900000000))
	TokenListingPrice    = sdk.NewCoin(apptesing.StableDenom, sdk.NewInt(3500000))
	// limit: 10000000uusdt - 100000000uusdt
	AllocationLimit = types.NewAllocationLimitArray(
		types.NewAllocationLimit(
			apptesing.StableDenom,
			sdk.NewCoin(apptesing.StableDenom, sdk.NewInt(10000000)),
			sdk.NewCoin(apptesing.StableDenom, sdk.NewInt(100000000)),
		),
	)
)

type KeeperTestSuite struct {
	apptesing.KeeperTestHelper

	queryClient types.QueryClient
	querier     keeper.Querier
}

func (suite *KeeperTestSuite) SetupTest() {
	// setup KeeperTestSuite
	suite.SetupKeeperTestHelper()
	suite.queryClient = types.NewQueryClient(suite.QueryHelper)
	suite.querier = keeper.NewQuerier(suite.App.IdoKeeper)

	// set params of global_project_id
	suite.App.IdoKeeper.SetParams(suite.Ctx, types.Params{})
}

func (suite *KeeperTestSuite) setupProject_IDO(project_owner sdk.AccAddress) uint64 {
	project_id, err := suite.App.LaunchpadKeeper.CreateProject(suite.Ctx, project_owner, launchpad.NewMsgCreateProjectRequest(
		project_owner.String(),
		"temp",
		"this is a temp project",
	))
	suite.Require().NoError(err)

	startTime := suite.Ctx.BlockTime().Add(time.Minute * 1)

	err = suite.querier.Keeper.EnableIDO(suite.Ctx, project_owner, types.NewMsgEnableIDORequest(
		project_owner.String(),
		project_id,
		sdk.NewCoins(TokenForDistribution),
		sdk.NewCoins(TokenListingPrice),
		AllocationLimit,
		startTime,
	))
	suite.Require().NoError(err)

	return project_id
}

func (suite *KeeperTestSuite) TestCalculateTokens() {
	// 50000000uusdt ~ 14285714unebula
	token := suite.querier.CalculateDistributionTokens(CommitTokens, sdk.NewCoins(TokenListingPrice), apptesing.TokenDenom)
	suite.Require().Equal(sdk.NewInt(14285714), token.AmountOf(apptesing.TokenDenom))

	// 900000000unebula ~ 3150000000uusdt
	token = suite.querier.CalculateCommitTokens(sdk.NewCoins(TokenForDistribution), sdk.NewCoins(TokenListingPrice), apptesing.StableDenom)
	suite.Require().Equal(sdk.NewInt(3150000000), token.AmountOf(apptesing.StableDenom))
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
