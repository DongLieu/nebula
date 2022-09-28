package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	// ModuleName defines the module name
	ModuleName = "ido"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_ido"
)

var (
	KeyPrefixIDO = []byte{0x02}
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func GetKeyPrefixProject(projectId uint64) []byte {
	return append(KeyPrefixIDO, sdk.Uint64ToBigEndian(projectId)...)
}
