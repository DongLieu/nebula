package types

import (
	"encoding/json"
	time "time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
	launchpadtypes "github.com/nebula-labs/nebula/x/launchpad/types"
)

var _ launchpadtypes.ReleaseMechanismI = &IDO{}
var _ proto.Message = &IDO{}

func NewIDO(projectId uint64, tokenForDistribution, totalDistributedAmount, tokenListingPrice sdk.Coins, idoStatus uint64, allocationLimit []AllocationLimit, startTime time.Time, entries map[string]Entry) IDO {
	return IDO{
		ProjectId:              projectId,
		TokenForDistribution:   tokenForDistribution,
		TotalDistributedAmount: totalDistributedAmount,
		TokenListingPrice:      tokenListingPrice,
		IdoStatus:              idoStatus,
		AllocationLimit:        allocationLimit,
		StartTime:              startTime,
		Entries:                entries,
	}
}

func (ido IDO) GetId() uint64 {
	return ido.ProjectId
}

func (ido IDO) GetReleaseMechanismStatus() uint64 {
	return ido.IdoStatus
}

func (ido IDO) GetTokens() sdk.Coins {
	return ido.TokenForDistribution
}

func (ido IDO) GetType() string {
	return ModuleName
}

func (ido *IDO) String() string {
	out, err := json.Marshal(ido)
	if err != nil {
		panic(err)
	}
	return string(out)
}
