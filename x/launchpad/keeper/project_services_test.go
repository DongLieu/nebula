package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	idotypes "github.com/nebula-labs/nebula/x/ido/types"
	"github.com/nebula-labs/nebula/x/launchpad/types"
)

func (suite *KeeperTestSuite) TestCreateProject() {

	// try out calling CreateProject() to see if any error
	func() {
		keeper := suite.App.LaunchpadKeeper

		// try to create project
		// tomorrow, this project is supposed to be inactive
		msg := types.MsgCreateProjectRequest{
			Owner:              suite.TestAccs[0].String(),
			ProjectTitle:       "Project Title",
			ProjectInformation: "Project Information",
		}

		// first create means that projectId is 1
		projectId, err := keeper.CreateProject(suite.Ctx, suite.TestAccs[0], &msg)
		suite.Require().NoError(err)
		suite.Require().Equal(uint64(1), projectId)

		// second create means that projectId is 2
		projectId, err = keeper.CreateProject(suite.Ctx, suite.TestAccs[0], &msg)
		suite.Require().NoError(err)
		suite.Require().Equal(uint64(2), projectId)
	}()

	tests := []struct {
		fn func()
	}{
		// ideal case
		{fn: func() {
			keeper := suite.App.LaunchpadKeeper
			msg := types.MsgCreateProjectRequest{
				Owner:              suite.TestAccs[0].String(),
				ProjectTitle:       "Project Title",
				ProjectInformation: "Project Information",
			}

			projectId, err := keeper.CreateProject(suite.Ctx, suite.TestAccs[0], &msg)
			suite.Require().NoError(err)

			project, err := keeper.GetProjectById(suite.Ctx, projectId)
			suite.Require().NoError(err)
			suite.Require().NotEqual(project, types.Project{})

			// check if project can contain token and return correctly the amount
			projectAddress := sdk.AccAddress(project.GetProjectAddress())

			suite.FundAcc(projectAddress, defaultAcctFunds)
			projectBalance := suite.App.BankKeeper.GetAllBalances(suite.Ctx, projectAddress)
			suite.Require().Equal(projectBalance, defaultAcctFunds)
		}},
	}

	for _, test := range tests {
		suite.SetupTest()

		test.fn()
	}
}

//============ DeleteProject testing ============

func (suite *KeeperTestSuite) TestDeleteProjectWithState() {
	tests := []struct {
		name         string
		projectState uint64
		idoState     uint64
		expectErr    bool
	}{
		{
			name:         "Delete project when Project and all RM is in INIT",
			projectState: types.PROJECT_INIT,
			idoState:     types.RM_INIT,
			expectErr:    false,
		},
		{
			name:         "Delete project when Project is ACTIVE and a RM is INIT",
			projectState: types.PROJECT_ACTIVE,
			idoState:     types.RM_INIT,
			expectErr:    true,
		},
		{
			name:         "Delete project when Project is ACTIVE and a RM is ACTIVE",
			projectState: types.PROJECT_ACTIVE,
			idoState:     types.RM_ACTIVE,
			expectErr:    true,
		},
		{
			name:         "Delete project when Project is ACTIVE and a RM is ENDED",
			projectState: types.PROJECT_ACTIVE,
			idoState:     types.RM_ENDED,
			expectErr:    true,
		},
		{
			name:         "Delete project when Project and RM is ENDED",
			projectState: types.PROJECT_ENDED,
			idoState:     types.RM_ENDED,
			expectErr:    true,
		},
	}

	for _, test := range tests {
		suite.SetupTest()

		project := suite.NewBaseProject()
		ido := suite.NewIdo(project)

		project.ProjectStatus = test.projectState
		suite.querier.SetProject(suite.Ctx, project)
		ido.IdoStatus = test.idoState
		suite.App.IdoKeeper.SetIDOAndRegisterLaunchpad(suite.Ctx, ido)

		owner, err := sdk.AccAddressFromBech32(project.ProjectOwner)
		suite.Require().NoError(err, "test: %v", test.name)

		// pre state
		ownerBefore := suite.App.BankKeeper.GetAllBalances(suite.Ctx, owner)
		suite.Require().Equal(true, suite.App.IdoKeeper.HasIDO(suite.Ctx, project.ProjectId))

		err = suite.querier.DeleteProject(suite.Ctx, owner, &types.MsgDeleteProjectRequest{
			Owner:     project.ProjectOwner,
			ProjectId: project.ProjectId,
		})

		if test.expectErr {
			suite.Require().Error(err, "test: %v", test.name)
			continue
		} else {
			suite.Require().NoError(err, "test: %v", test.name)
		}

		// check balances correctly adjust
		ownerAfter := suite.App.BankKeeper.GetAllBalances(suite.Ctx, owner)
		ownerExpected := ownerBefore.Add(ido.TokenForDistribution...)
		suite.Require().Equal(ownerExpected, ownerAfter, "test: %v", test.name)

		// check project and IDO is deleted
		_, err = suite.querier.GetProjectById(suite.Ctx, project.ProjectId)
		suite.Require().Error(err, "test: %v", test.name)

		suite.Require().Equal(false, suite.App.IdoKeeper.HasIDO(suite.Ctx, project.ProjectId))
	}
}

func (suite *KeeperTestSuite) TestDeleteProjectWithSituation() {
	tests := []struct {
		fn func()
	}{
		// check not owner
		{
			fn: func() {
				project := suite.NewBaseProject()
				keeper := suite.App.LaunchpadKeeper
				msg := types.MsgDeleteProjectRequest{
					Owner:     suite.TestAccs[1].String(),
					ProjectId: project.ProjectId,
				}

				owner, err := sdk.AccAddressFromBech32(project.ProjectOwner)
				suite.Require().NoError(err)

				err = keeper.DeleteProject(suite.Ctx, owner, &msg)
				suite.Require().Error(err, types.ErrNotProjecOwner)
			},
		},
		// check if ID is still valid (if not existed before)
		{
			fn: func() {
				project := suite.NewBaseProject()
				msg := types.MsgDeleteProjectRequest{
					Owner:     suite.TestAccs[0].String(),
					ProjectId: uint64(project.ProjectId + 1),
				}

				owner, err := sdk.AccAddressFromBech32(project.ProjectOwner)
				suite.Require().NoError(err)

				err = suite.querier.DeleteProject(suite.Ctx, owner, &msg)
				suite.Require().Error(err)
			},
		},
		// Delete multiple times
		{
			fn: func() {
				project := suite.NewBaseProject()
				keeper := suite.App.LaunchpadKeeper
				msg := types.MsgDeleteProjectRequest{
					Owner:     suite.TestAccs[0].String(),
					ProjectId: project.ProjectId,
				}

				owner, err := sdk.AccAddressFromBech32(project.ProjectOwner)
				suite.Require().NoError(err)

				err = keeper.DeleteProject(suite.Ctx, owner, &msg)
				suite.Require().NoError(err)

				// Try delete again
				newErr := keeper.DeleteProject(suite.Ctx, owner, &msg)
				suite.Require().Error(newErr)
			},
		},
	}

	for _, test := range tests {
		suite.SetupTest()

		test.fn()
	}
}

//============ WithdrawTokens testing ============

func (suite *KeeperTestSuite) TestWithdrawTokensWithState() {
	tests := []struct {
		name         string
		projectState uint64
		idoState     uint64
		expectErr    bool
	}{
		{
			name:         "Withdraw tokens when Project and all RM is in INIT",
			projectState: types.PROJECT_INIT,
			idoState:     types.RM_INIT,
			expectErr:    true,
		},
		{
			name:         "Withdraw tokens when Project is ACTIVE and a RM is INIT",
			projectState: types.PROJECT_ACTIVE,
			idoState:     types.RM_INIT,
			expectErr:    true,
		},
		{
			name:         "Withdraw tokens when Project is ACTIVE and a RM is ACTIVE",
			projectState: types.PROJECT_ACTIVE,
			idoState:     types.RM_ACTIVE,
			expectErr:    true,
		},
		{
			name:         "Withdraw tokens when Project is ACTIVE and a RM is ENDED",
			projectState: types.PROJECT_ACTIVE,
			idoState:     types.RM_ENDED,
			expectErr:    true,
		},
		{
			name:         "Withdraw tokens when Project and RM is ENDED",
			projectState: types.PROJECT_ENDED,
			idoState:     types.RM_ENDED,
			expectErr:    false,
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		// setup project
		project := suite.NewBaseProject()
		ido := suite.NewIdo(project)
		owner, err := sdk.AccAddressFromBech32(project.ProjectOwner)
		suite.Require().NoError(err, "test: %v", test.name)

		// setup state
		project.ProjectStatus = test.projectState
		suite.querier.SetProject(suite.Ctx, project)
		ido.IdoStatus = test.idoState
		suite.App.IdoKeeper.SetIDOAndRegisterLaunchpad(suite.Ctx, ido)

		// pre state
		ownerBefore := suite.App.BankKeeper.GetAllBalances(suite.Ctx, owner)
		idoBefore := ido.TokenForDistribution

		err = suite.querier.WithdrawTokens(suite.Ctx, owner, types.NewMsgWithdrawAllTokensRequest(
			owner.String(),
			project.ProjectId,
		))

		if test.expectErr {
			suite.Require().Error(err, "test: %v", test.name)
			continue
		} else {
			suite.Require().NoError(err, "test: %v", test.name)
		}

		// after
		project, err = suite.querier.GetProjectById(suite.Ctx, project.ProjectId)
		suite.Require().NoError(err, "test: %v", test.name)
		ido, err = suite.App.IdoKeeper.GetIDOByID(suite.Ctx, project.ProjectId)
		suite.Require().NoError(err, "test: %v", test.name)

		// check balances correctly adjust
		ownerAfter := suite.App.BankKeeper.GetAllBalances(suite.Ctx, owner)
		ownerExpected := ownerBefore.Add(idoBefore...)
		suite.Require().Equal(ownerExpected, ownerAfter, "test: %v", test.name)
		projectAfter := suite.App.BankKeeper.GetAllBalances(suite.Ctx, sdk.AccAddress(project.ProjectAddress))
		suite.Require().Equal(0, len(projectAfter), "test: %v", test.name)
	}
}

func (suite *KeeperTestSuite) TestWithdrawTokensWithSituation() {
	tests := []struct {
		fn func(project types.Project, ido idotypes.IDO)
	}{
		// check not owner
		{
			fn: func(project types.Project, ido idotypes.IDO) {
				owner, err := sdk.AccAddressFromBech32(project.ProjectOwner)
				suite.Require().NoError(err)

				// setup state
				project.ProjectStatus = types.PROJECT_ENDED
				suite.querier.SetProject(suite.Ctx, project)
				ido.IdoStatus = types.RM_ENDED
				suite.App.IdoKeeper.SetIDOAndRegisterLaunchpad(suite.Ctx, ido)

				err = suite.querier.WithdrawTokens(suite.Ctx, owner, types.NewMsgWithdrawAllTokensRequest(
					suite.TestAccs[1].String(),
					project.ProjectId,
				))
				suite.Require().Error(err, types.ErrNotProjecOwner)
			},
		},
		// check if ID is still valid (if not existed before)
		{
			fn: func(project types.Project, ido idotypes.IDO) {
				owner, err := sdk.AccAddressFromBech32(project.ProjectOwner)
				suite.Require().NoError(err)

				// setup state
				project.ProjectStatus = types.PROJECT_ENDED
				suite.querier.SetProject(suite.Ctx, project)
				ido.IdoStatus = types.RM_ENDED
				suite.App.IdoKeeper.SetIDOAndRegisterLaunchpad(suite.Ctx, ido)

				err = suite.querier.WithdrawTokens(suite.Ctx, owner, types.NewMsgWithdrawAllTokensRequest(
					owner.String(),
					uint64(project.ProjectId+1),
				))
				suite.Require().Error(err)
			},
		},
		// Withdraw tokens multiple time
		{
			fn: func(project types.Project, ido idotypes.IDO) {
				owner, err := sdk.AccAddressFromBech32(project.ProjectOwner)
				suite.Require().NoError(err)

				ownerBefore := suite.App.BankKeeper.GetAllBalances(suite.Ctx, owner)

				// setup state
				project.ProjectStatus = types.PROJECT_ENDED
				suite.querier.SetProject(suite.Ctx, project)
				ido.IdoStatus = types.RM_ENDED
				suite.App.IdoKeeper.SetIDOAndRegisterLaunchpad(suite.Ctx, ido)

				err = suite.querier.WithdrawTokens(suite.Ctx, owner, types.NewMsgWithdrawAllTokensRequest(
					owner.String(),
					project.ProjectId,
				))
				suite.Require().NoError(err)

				// check balances correctly adjust
				ownerAfter := suite.App.BankKeeper.GetAllBalances(suite.Ctx, owner)
				ownerExpected := ownerBefore.Add(ido.TokenForDistribution...)
				suite.Require().Equal(ownerExpected, ownerAfter)

				err = suite.querier.WithdrawTokens(suite.Ctx, owner, types.NewMsgWithdrawAllTokensRequest(
					owner.String(),
					project.ProjectId,
				))
				suite.Require().NoError(err)

				// check balances correctly remain
				ownerAfterAfter := suite.App.BankKeeper.GetAllBalances(suite.Ctx, owner)
				suite.Require().Equal(ownerAfter, ownerAfterAfter)
			},
		},
	}

	for _, test := range tests {
		suite.SetupTest()

		// setup project
		project := suite.NewBaseProject()

		// setup an IDO to this project
		ido := suite.NewIdo(project)

		test.fn(project, ido)
	}
}
