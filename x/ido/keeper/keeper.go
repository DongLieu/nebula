package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/nebula-labs/nebula/x/ido/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   sdk.StoreKey
		memKey     sdk.StoreKey
		paramstore paramtypes.Subspace

		launchpadKeeper types.LaunchpadKeeper
		bankKeeper      types.BankKeeper
		accKeeper       types.AccountKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,
	launchpadKeeper types.LaunchpadKeeper,
	bankKeeper types.BankKeeper,
	accKeeper types.AccountKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,

		launchpadKeeper: launchpadKeeper,
		bankKeeper:      bankKeeper,
		accKeeper:       accKeeper,
	}
}

func (k Keeper) CalculateDistributionTokens(tokenCommit, tokenListingPrice sdk.Coins, distributionDenom string) sdk.Coins {
	amount := tokenCommit[0].Amount.ToDec().Quo(tokenListingPrice[0].Amount.ToDec()).Mul(sdk.NewIntWithDecimal(1, 6).ToDec())
	return sdk.NewCoins(sdk.NewCoin(distributionDenom, amount.RoundInt()))
}

func (k Keeper) CalculateCommitTokens(tokenDistribution, tokenListingPrice sdk.Coins, distributionDenom string) sdk.Coins {
	amount := tokenDistribution[0].Amount.ToDec().Quo(sdk.NewIntWithDecimal(1, 6).ToDec()).Mul(tokenListingPrice[0].Amount.ToDec())
	return sdk.NewCoins(sdk.NewCoin(distributionDenom, amount.RoundInt()))
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
