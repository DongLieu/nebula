package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	apptesting "github.com/nebula-labs/nebula/app/apptesting"
	"github.com/nebula-labs/nebula/x/ido/types"
	launchpadtypes "github.com/nebula-labs/nebula/x/launchpad/types"
)

func (suite KeeperTestSuite) TestIDOStatusChange() {
	initialBlockTime := time.Now()

	tests := []struct {
		fn func()
	}{
		// test if ido can be set active
		{fn: func() {
			suite.Ctx = suite.Ctx.WithBlockHeight(1).WithBlockTime(initialBlockTime)

			// create a project
			projectOwner := suite.TestAccs[0]
			projectId := suite.setupProject_IDO(projectOwner)

			// simulate block change
			suite.querier.BeginBlocker(suite.Ctx)
			for i := 1; i <= 5; i++ {
				suite.Ctx = suite.Ctx.WithBlockHeight(suite.Ctx.BlockHeight() + int64(1)).WithBlockTime(suite.Ctx.BlockTime().Add(time.Second * 20))
				suite.querier.BeginBlocker(suite.Ctx)
			}

			// status change
			ido, err := suite.querier.GetIDOByID(suite.Ctx, projectId)
			suite.Require().NoError(err)
			suite.Require().Equal(launchpadtypes.RM_ACTIVE, ido.IdoStatus)

			// check project status change as well
			project, err := suite.App.LaunchpadKeeper.GetProjectById(suite.Ctx, projectId)
			suite.Require().NoError(err)
			suite.Require().Equal(launchpadtypes.PROJECT_ACTIVE, project.ProjectStatus)
		}},
		// test if ido can end
		{fn: func() {
			suite.Ctx = suite.Ctx.WithBlockHeight(1).WithBlockTime(initialBlockTime)

			// create a project
			projectOwner := suite.TestAccs[0]
			projectId := suite.setupProject_IDO(projectOwner)

			// change ido smaller
			ido, err := suite.querier.GetIDOByID(suite.Ctx, projectId)
			suite.Require().NoError(err)
			ido.TokenForDistribution = sdk.NewCoins(sdk.NewCoin(apptesting.TokenDenom, sdk.NewInt(10000000)))
			suite.querier.SetIDOAndRegisterLaunchpad(suite.Ctx, ido)

			// simulate block change
			suite.querier.BeginBlocker(suite.Ctx)
			for i := 1; i <= 5; i++ {
				suite.Ctx = suite.Ctx.WithBlockHeight(suite.Ctx.BlockHeight() + int64(1)).WithBlockTime(suite.Ctx.BlockTime().Add(time.Second * 20))
				suite.querier.BeginBlocker(suite.Ctx)
			}

			// simulate all token is distributed
			commitTokens := suite.querier.CalculateCommitTokens(ido.TokenForDistribution, sdk.NewCoins(TokenListingPrice), apptesting.StableDenom)

			err = suite.querier.CommitParticipation(suite.Ctx, suite.TestAccs[1], types.NewMsgCommitParticipationRequest(
				suite.TestAccs[1].String(),
				projectId,
				commitTokens,
			))
			suite.Require().NoError(err)

			// simulate block change
			for i := 1; i <= 5; i++ {
				suite.Ctx = suite.Ctx.WithBlockHeight(suite.Ctx.BlockHeight() + int64(1)).WithBlockTime(suite.Ctx.BlockTime().Add(time.Second * 20))
				suite.querier.BeginBlocker(suite.Ctx)
			}

			// status change
			ido, err = suite.querier.GetIDOByID(suite.Ctx, projectId)
			suite.Require().NoError(err)
			suite.Require().Equal(launchpadtypes.RM_ENDED, ido.IdoStatus)

			// check project status change as well
			project, err := suite.App.LaunchpadKeeper.GetProjectById(suite.Ctx, projectId)
			suite.Require().NoError(err)
			suite.Require().Equal(launchpadtypes.PROJECT_ENDED, project.ProjectStatus)
		}},
	}

	for _, test := range tests {
		suite.SetupTest()

		test.fn()
	}
}
