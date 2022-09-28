package types

import (
	fmt "fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// string format here would be like: 10000000uusdt-100000000uusdt,10000000uusdt-100000000uusdt
func ParseAllocationLimitArrayFromString(allocationLimitStr string) ([]AllocationLimit, error) {
	allocationLimitStr = strings.TrimSpace(allocationLimitStr)
	if len(allocationLimitStr) == 0 {
		return nil, nil
	}

	allocationRanges := strings.Split(allocationLimitStr, ",")
	res := make([]AllocationLimit, len(allocationRanges))
	for i, allocationRange := range allocationRanges {
		allocationLimits := strings.Split(allocationRange, "-")
		lowerLimit, err := sdk.ParseCoinNormalized(allocationLimits[0])
		if err != nil {
			return nil, err
		}
		upperLimit, err := sdk.ParseCoinNormalized(allocationLimits[1])
		if err != nil {
			return nil, err
		}
		res[i] = NewAllocationLimit(lowerLimit.Denom, lowerLimit, upperLimit)
	}

	return res, nil
}

func ParseStringFromAllocationLimitArray(allocationLimit []*AllocationLimit) string {
	res := make([]string, len(allocationLimit))
	for i, limit := range allocationLimit {
		limitRange := fmt.Sprintf("%s-%s", limit.LowerLimit.String(), limit.UpperLimit.String())
		res[i] = limitRange
	}

	return strings.Join(res, ",")
}

func NewAllocationLimitArray(allocationLimit ...AllocationLimit) []AllocationLimit {
	return allocationLimit
}

func NewAllocationLimit(denom string, lowerLimit, upperLimit sdk.Coin) AllocationLimit {
	return AllocationLimit{
		Denom:      denom,
		LowerLimit: lowerLimit,
		UpperLimit: upperLimit,
	}
}
