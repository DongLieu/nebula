package keeper_test

import (
	"fmt"

	idotypes "github.com/nebula-labs/nebula/x/ido/types"
	"github.com/nebula-labs/nebula/x/launchpad/types"
)

func (suite *KeeperTestSuite) TestRegisterReleaseMechanismToProject() {

	tests := []struct {
		fn func(project types.Project, ido idotypes.IDO)
	}{
		{
			// check if registration is succeed
			fn: func(project types.Project, ido idotypes.IDO) {
				project, err := suite.querier.GetProjectById(suite.Ctx, project.ProjectId)
				suite.Require().NoError(err)

				suite.Require().Equal(1, len(project.RegisteredRm))

				// try changing ido
				ido.IdoStatus = types.RM_ENDED
				suite.App.IdoKeeper.SetIDO(suite.Ctx, ido)
				err = suite.querier.RegisterReleaseMechanismToProject(suite.Ctx, project.ProjectId, &ido)
				suite.Require().NoError(err)

				// fetch new project and test
				project, err = suite.querier.GetProjectById(suite.Ctx, project.ProjectId)
				suite.Require().NoError(err)
				rm, err := suite.querier.ParseAnyReleaseMechanism(project.RegisteredRm[0])
				suite.Require().NoError(err)
				fmt.Printf("rm = %v \n", rm)
				suite.Require().Equal(ido.IdoStatus, rm.GetReleaseMechanismStatus())
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
