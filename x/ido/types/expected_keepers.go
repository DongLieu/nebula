package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	launchpadtypes "github.com/nebula-labs/nebula/x/launchpad/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error

	GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	HasBalance(ctx sdk.Context, addr sdk.AccAddress, amt sdk.Coin) bool
}

type LaunchpadKeeper interface {
	GetProjectById(ctx sdk.Context, projectID uint64) (launchpadtypes.Project, error)
	RegisterReleaseMechanismToProject(ctx sdk.Context, projectId uint64, rm launchpadtypes.ReleaseMechanismI) error
	SetProjectActive(ctx sdk.Context, projectId uint64) error
	SetProjectEndable(ctx sdk.Context, projectId uint64) error
}
