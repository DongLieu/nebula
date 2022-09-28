package keeper_test

import (
	"testing"
	"time"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	apptesing "github.com/nebula-labs/nebula/app/apptesting"
	idotypes "github.com/nebula-labs/nebula/x/ido/types"
	"github.com/nebula-labs/nebula/x/launchpad/keeper"
	"github.com/nebula-labs/nebula/x/launchpad/types"
	"github.com/stretchr/testify/suite"
)

var (
	defaultAcctFunds sdk.Coins = sdk.NewCoins(
		sdk.NewCoin(apptesing.TokenDenom, sdk.NewInt(10000000)),
	)
	TokenForDistribution = sdk.NewCoin(apptesing.TokenDenom, sdk.NewInt(1000000000))
	TokenListingPrice    = sdk.NewCoin(apptesing.StableDenom, sdk.NewInt(1000000))
	// limit: 10000000uusdt - 100000000uusdt
	AllocationLimit = idotypes.NewAllocationLimitArray(
		idotypes.NewAllocationLimit(
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
	suite.querier = keeper.NewQuerier(suite.App.LaunchpadKeeper)

	// set params of global_project_id
	suite.App.LaunchpadKeeper.SetParams(suite.Ctx, types.Params{})
	suite.App.LaunchpadKeeper.SetNextProjectID(suite.Ctx, 0)
}

func (suite *KeeperTestSuite) NewBaseProject() types.Project {
	projectId := uint64(1)

	projectAddress := suite.querier.NewProjectAddress(projectId)

	project := types.Project{
		ProjectOwner:       suite.TestAccs[0].String(),
		ProjectTitle:       "temp",
		ProjectId:          projectId,
		ProjectAddress:     projectAddress.String(),
		ProjectInformation: "",
		ProjectStatus:      types.PROJECT_INIT,
		ProjectVerified:    false,
		RegisteredRm:       []*codectypes.Any{},
	}

	suite.querier.SetProject(suite.Ctx, project)

	return project
}

func (suite *KeeperTestSuite) NewIdo(project types.Project) idotypes.IDO {
	owner, err := sdk.AccAddressFromBech32(project.ProjectOwner)
	suite.Require().NoError(err)

	startTime := suite.Ctx.BlockTime().Add(time.Minute * 30)

	err = suite.App.IdoKeeper.EnableIDO(suite.Ctx, owner, idotypes.NewMsgEnableIDORequest(
		project.ProjectOwner,
		project.ProjectId,
		sdk.NewCoins(TokenForDistribution),
		sdk.NewCoins(TokenListingPrice),
		AllocationLimit,
		startTime,
	))
	suite.Require().NoError(err)

	ido, err := suite.App.IdoKeeper.GetIDOByID(suite.Ctx, project.ProjectId)
	suite.Require().NoError(err)

	return ido
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
