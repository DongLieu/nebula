package apptesing

import (
	"time"

	"github.com/cosmos/cosmos-sdk/baseapp"
	cosmossimapp "github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/nebula-labs/nebula/app"
	"github.com/nebula-labs/nebula/testutil/simapp"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/crypto/ed25519"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

var (
	TokenDenom  = "unebula"
	StableDenom = "uusdt"
)

type KeeperTestHelper struct {
	suite.Suite

	Ctx         sdk.Context
	App         *app.App
	QueryHelper *baseapp.QueryServiceTestHelper
	TestAccs    []sdk.AccAddress
}

func (s *KeeperTestHelper) SetupKeeperTestHelper() {
	// setup app
	s.App = simapp.New(false)

	// setup ctx
	s.Ctx = s.App.BaseApp.NewContext(false, tmproto.Header{Height: 1, ChainID: "nebula-1", Time: time.Now().UTC()})

	// setup query helper
	s.QueryHelper = &baseapp.QueryServiceTestHelper{
		GRPCQueryRouter: s.App.GRPCQueryRouter(),
		Ctx:             s.Ctx,
	}

	// setup test accs
	s.TestAccs = CreateRandomAccounts(3)
	for _, testAcc := range s.TestAccs {
		s.FundAcc(testAcc, sdk.NewCoins(sdk.NewCoin(TokenDenom, sdk.NewInt(10000000000)), sdk.NewCoin(StableDenom, sdk.NewInt(10000000000))))
	}
}

// CreateRandomAccounts is a function return a list of randomly generated AccAddresses
func CreateRandomAccounts(numAccts int) []sdk.AccAddress {
	testAddrs := make([]sdk.AccAddress, numAccts)
	for i := 0; i < numAccts; i++ {
		pk := ed25519.GenPrivKey().PubKey()
		testAddrs[i] = sdk.AccAddress(pk.Address())
	}

	return testAddrs
}

func (s *KeeperTestHelper) FundAcc(acc sdk.AccAddress, amounts sdk.Coins) {
	err := cosmossimapp.FundAccount(s.App.BankKeeper, s.Ctx, acc, amounts)
	s.Require().NoError(err)
}
