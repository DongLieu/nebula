package keeper_test

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	apptesting "github.com/nebula-labs/nebula/app/apptesting"
	"github.com/nebula-labs/nebula/x/ido/types"
	launchpadtypes "github.com/nebula-labs/nebula/x/launchpad/types"
)

//============ EnableIDO testing ============

func (suite *KeeperTestSuite) TestEnableIDOWithState() {
	tests := []struct {
		name                 string
		projectState         uint64
		tokenForDistribution sdk.Coin
		expectErr            bool
	}{
		{
			name:                 "enable IDO when project is just initialized",
			projectState:         launchpadtypes.PROJECT_INIT,
			tokenForDistribution: TokenForDistribution,
			expectErr:            false,
		},
		{
			name:                 "enable IDO when project is active",
			projectState:         launchpadtypes.PROJECT_ACTIVE,
			tokenForDistribution: TokenForDistribution,
			expectErr:            true,
		},
		{
			name:                 "enable IDO when project is ended",
			projectState:         launchpadtypes.PROJECT_ENDED,
			tokenForDistribution: TokenForDistribution,
			expectErr:            true,
		},
		{
			name:                 "enable IDO when there is not enough funds",
			projectState:         launchpadtypes.PROJECT_INIT,
			tokenForDistribution: sdk.NewCoin(apptesting.TokenDenom, sdk.NewInt(100000000000)),
			expectErr:            true,
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		project_owner := suite.TestAccs[0]
		project_id, err := suite.App.LaunchpadKeeper.CreateProject(suite.Ctx, project_owner, launchpadtypes.NewMsgCreateProjectRequest(
			project_owner.String(),
			"temp",
			"this is a temp project",
		))
		suite.Require().NoError(err, "test: %v", test.name)

		// change state of project
		project, err := suite.App.LaunchpadKeeper.GetProjectById(suite.Ctx, project_id)
		suite.Require().NoError(err, "test: %v", test.name)
		project.ProjectStatus = test.projectState
		suite.App.LaunchpadKeeper.SetProject(suite.Ctx, project)

		// acc initial coin
		accCoinBefore := suite.App.BankKeeper.GetBalance(suite.Ctx, project_owner, apptesting.TokenDenom)

		// create MsgEnableIDO
		err = suite.querier.Keeper.EnableIDO(suite.Ctx, project_owner, types.NewMsgEnableIDORequest(
			project_owner.String(),
			project_id,
			sdk.NewCoins(test.tokenForDistribution),
			sdk.NewCoins(TokenListingPrice),
			AllocationLimit,
			suite.Ctx.BlockTime().Add(time.Minute*30),
		))

		if test.expectErr {
			suite.Require().Error(err, "test: %v", test.name)
			continue
		} else {
			suite.Require().NoError(err, "test: %v", test.name)
		}

		// coin is added to project address correctly
		// coin is subtracted correctly from project owner
		projCoin := suite.App.BankKeeper.GetBalance(suite.Ctx, sdk.AccAddress(project.GetProjectAddress()), apptesting.TokenDenom)
		suite.Require().Equal(test.tokenForDistribution, projCoin, "test: %v", test.name)
		accCoinAfter := suite.App.BankKeeper.GetBalance(suite.Ctx, project_owner, apptesting.TokenDenom)
		suite.Require().Equal(accCoinBefore.Sub(test.tokenForDistribution), accCoinAfter, "test: %v", test.name)
	}
}

func (suite *KeeperTestSuite) TestEnableIDOWithSituation() {
	tests := []struct {
		name string
		fn   func(name string)
	}{
		{
			name: "enable IDO when there is no project",
			fn: func(name string) {
				project_owner := suite.TestAccs[0]
				project_id := uint64(1)

				// create MsgEnableIDO
				err := suite.querier.Keeper.EnableIDO(suite.Ctx, project_owner, types.NewMsgEnableIDORequest(
					project_owner.String(),
					project_id,
					sdk.NewCoins(TokenForDistribution),
					sdk.NewCoins(TokenListingPrice),
					AllocationLimit,
					suite.Ctx.BlockTime().Add(time.Minute*30),
				))

				suite.Require().Error(err, "test: %v", name)
			},
		},
		{
			name: "Enable IDO repeatedly",
			fn: func(name string) {
				projectOwner := suite.TestAccs[0]

				// create project and IDO once
				projectId := suite.setupProject_IDO(projectOwner)

				suite.Require().Equal(true, suite.querier.HasIDO(suite.Ctx, projectId), "test: %v", name)

				// enable IDO once more
				err := suite.querier.Keeper.EnableIDO(suite.Ctx, projectOwner, types.NewMsgEnableIDORequest(
					projectOwner.String(),
					projectId,
					sdk.NewCoins(TokenForDistribution),
					sdk.NewCoins(TokenListingPrice),
					AllocationLimit,
					suite.Ctx.BlockTime().Add(time.Minute*30),
				))

				suite.Require().Error(err, "test: %v", name)
			},
		},
	}

	for _, test := range tests {
		suite.SetupTest()

		test.fn(test.name)
	}
}

// ============ CommitParticipation testing ============
var (
	CommitTokens = sdk.NewCoins(sdk.NewCoin(apptesting.StableDenom, sdk.NewInt(50000000)))
)

func (suite *KeeperTestSuite) TestCommitParticipationWithState() {
	tests := []struct {
		name         string
		rmState      uint64
		commitTokens sdk.Coins
		expectErr    bool
	}{
		{
			name:         "Commit participation when RM is in INIT",
			rmState:      launchpadtypes.RM_INIT,
			commitTokens: CommitTokens,
			expectErr:    true,
		},
		{
			name:         "Commit participation when RM is in ACTIVE",
			rmState:      launchpadtypes.RM_ACTIVE,
			commitTokens: CommitTokens,
			expectErr:    false,
		},
		{
			name:         "Commit participation when participant commit is out of range",
			rmState:      launchpadtypes.RM_ACTIVE,
			commitTokens: sdk.NewCoins(sdk.NewCoin(apptesting.StableDenom, sdk.NewInt(1000))),
			expectErr:    true,
		},
		{
			name:         "Commit participation when RM is in ENDED",
			rmState:      launchpadtypes.RM_ENDED,
			commitTokens: CommitTokens,
			expectErr:    true,
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		projectOwner := suite.TestAccs[0]
		participant := suite.TestAccs[1]

		projectId := suite.setupProject_IDO(projectOwner)
		project, err := suite.App.LaunchpadKeeper.GetProjectById(suite.Ctx, projectId)
		suite.Require().NoError(err)

		// change ido state
		ido, err := suite.querier.GetIDOByID(suite.Ctx, projectId)
		suite.Require().NoError(err, "test: %v", test.name)
		ido.IdoStatus = test.rmState
		err = suite.querier.SetIDOAndRegisterLaunchpad(suite.Ctx, ido)
		suite.Require().NoError(err, "test: %v", test.name)

		// amount before participation
		participantTokensBefore := suite.App.BankKeeper.GetAllBalances(suite.Ctx, participant)
		projTokensBefore := suite.App.BankKeeper.GetAllBalances(suite.Ctx, sdk.AccAddress(project.GetProjectAddress()))
		tokenForDistributionBefore := ido.TokenForDistribution
		totalDistributedAmountBefore := ido.TotalDistributedAmount

		// commit
		err = suite.querier.CommitParticipation(suite.Ctx, participant, types.NewMsgCommitParticipationRequest(
			participant.String(),
			projectId,
			test.commitTokens,
		))

		if test.expectErr {
			suite.Require().Error(err, "test: %v", test.name)
			continue
		} else {
			suite.Require().NoError(err, "test: %v", test.name)
		}

		// check correct fund is moved
		ido, err = suite.querier.GetIDOByID(suite.Ctx, projectId)
		suite.Require().NoError(err, "test: %v", test.name)

		calculatedDistributionTokens := suite.querier.CalculateDistributionTokens(test.commitTokens, ido.TokenListingPrice, ido.TokenForDistribution.GetDenomByIndex(0))

		participantTokensAfter := suite.App.BankKeeper.GetAllBalances(suite.Ctx, participant)
		projTokensAfter := suite.App.BankKeeper.GetAllBalances(suite.Ctx, sdk.AccAddress(project.GetProjectAddress()))
		tokenForDistributionAfter := ido.TokenForDistribution
		totalDistributedAmountAfter := ido.TotalDistributedAmount

		fmt.Printf("participantTokensAfter = %v, projTokensAfter = %v \n", participantTokensAfter, projTokensAfter)

		participantExpected := participantTokensBefore.Add(calculatedDistributionTokens...)
		participantExpected = participantExpected.Sub(test.commitTokens)
		projExpected := projTokensBefore.Add(test.commitTokens...)
		projExpected = projExpected.Sub(calculatedDistributionTokens)
		suite.Require().Equal(participantExpected, participantTokensAfter, "test: %v", test.name)
		suite.Require().Equal(projExpected, projTokensAfter, "test: %v", test.name)

		// TODO: check ido.TokenForDistribution and ido.TotalDistributedAmount as well
		tokenForDistributionExpected := tokenForDistributionBefore.Sub(calculatedDistributionTokens)
		totalDistributedAmountExpected := totalDistributedAmountBefore.Add(calculatedDistributionTokens...)
		suite.Require().Equal(tokenForDistributionExpected, tokenForDistributionAfter, "test: %v", test.name)
		suite.Require().Equal(totalDistributedAmountExpected, totalDistributedAmountAfter, "test: %v", test.name)
	}
}

func (suite *KeeperTestSuite) TestCommitParticipationSituation() {

	tests := []struct {
		fn func(project_id uint64)
	}{
		// test if not enough delegation
		{fn: func(project_id uint64) {
			// create project ido
			participant := suite.TestAccs[1]

			// activate ido
			ido, err := suite.querier.GetIDOByID(suite.Ctx, project_id)
			suite.Require().NoError(err)
			ido.IdoStatus = launchpadtypes.RM_ACTIVE
			ido.TokenForDistribution = sdk.NewCoins(sdk.NewCoin(apptesting.TokenDenom, sdk.NewInt(10000000)))
			suite.querier.SetIDOAndRegisterLaunchpad(suite.Ctx, ido)

			// 50000000uusdt ~ 14285714unebula
			err = suite.querier.CommitParticipation(suite.Ctx, participant, types.NewMsgCommitParticipationRequest(
				participant.String(),
				project_id,
				CommitTokens,
			))

			suite.Require().EqualError(types.ErrNotEnoughIDOTokens, err.Error())
		}},
		// test a person fails to commit accumulated amount over allocation limit range
		{fn: func(project_id uint64) {
			// create project ido
			participant := suite.TestAccs[1]

			// activate ido
			ido, err := suite.querier.GetIDOByID(suite.Ctx, project_id)
			suite.Require().NoError(err)
			ido.IdoStatus = launchpadtypes.RM_ACTIVE
			suite.querier.SetIDOAndRegisterLaunchpad(suite.Ctx, ido)

			// 100000000uusdt means 100 nebula --> 100000000unebula
			err = suite.querier.CommitParticipation(suite.Ctx, participant, types.NewMsgCommitParticipationRequest(
				participant.String(),
				project_id,
				sdk.NewCoins(sdk.NewCoin(apptesting.StableDenom, sdk.NewInt(100000000))),
			))
			suite.Require().NoError(err)

			// check
			err = suite.querier.CommitParticipation(suite.Ctx, participant, types.NewMsgCommitParticipationRequest(
				participant.String(),
				project_id,
				sdk.NewCoins(sdk.NewCoin(apptesting.StableDenom, sdk.NewInt(100000000))),
			))
			suite.Require().EqualError(err, types.ErrOutOfBoundPurchase.Error())
		}},
		// not enough commit tokens
		{fn: func(project_id uint64) {
			// create project ido
			participant := apptesting.CreateRandomAccounts(1)[0]
			suite.FundAcc(participant, sdk.NewCoins(sdk.NewCoin(apptesting.TokenDenom, sdk.NewInt(1000000000)), sdk.NewCoin(apptesting.StableDenom, sdk.NewInt(15000000))))

			// activate ido
			ido, err := suite.querier.GetIDOByID(suite.Ctx, project_id)
			suite.Require().NoError(err)
			ido.IdoStatus = launchpadtypes.RM_ACTIVE
			suite.querier.SetIDOAndRegisterLaunchpad(suite.Ctx, ido)

			// 50000000uusdt ~ 14285714unebula
			err = suite.querier.CommitParticipation(suite.Ctx, participant, types.NewMsgCommitParticipationRequest(
				participant.String(),
				project_id,
				CommitTokens,
			))

			suite.Require().EqualError(types.ErrNotEnoughFunds, err.Error())
		}},
	}

	for _, test := range tests {
		suite.SetupTest()
		project_owner := suite.TestAccs[0]
		project_id := suite.setupProject_IDO(project_owner)

		test.fn(project_id)
	}
}
