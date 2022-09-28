package keeper_test

import (
	"github.com/nebula-labs/nebula/x/launchpad/keeper"
	"github.com/nebula-labs/nebula/x/launchpad/types"
)

func (suite *KeeperTestSuite) TestHandleSetProjectVerifiedProposal() {
	type Action struct {
		ProjectOwner string
		expectedErr  bool
	}

	testCase := []struct {
		name    string
		actions []Action
	}{
		{
			"correct one",
			[]Action{
				{
					string(suite.TestAccs[0]), false,
				},
			},
		},
	}

	for _, tc := range testCase {
		suite.Run(tc.name, func() {
			suite.SetupTest()
			project := suite.NewBaseProject()

			for _, action := range tc.actions {
				err := keeper.HandleSetProjectVerifiedProposal(suite.Ctx, suite.querier.Keeper, &types.SetProjectVerifiedProposal{
					Title:        "title",
					Description:  "description",
					ProjectOwner: suite.TestAccs[0].String(),
					ProjectId:    project.ProjectId,
				})

				if action.expectedErr {
					suite.Require().Error(err)
				} else {
					suite.Require().NoError(err)
				}

				//check state
				newProject, err := suite.querier.GetProjectById(suite.Ctx, project.ProjectId)
				suite.Require().NoError(err)
				suite.Require().Equal(true, newProject.ProjectVerified)
			}
		})
	}
}
