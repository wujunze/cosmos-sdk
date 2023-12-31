package types

import (
	"cosmossdk.io/collections"
	collcodec "cosmossdk.io/collections/codec"
	"cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName defines the module name
	ModuleName = "bank"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// GovModuleName duplicates the gov module's name to avoid a cyclic dependency with x/gov.
	// It should be synced with the gov module's name if it is ever changed.
	// See: https://github.com/cosmos/cosmos-sdk/blob/b62a28aac041829da5ded4aeacfcd7a42873d1c8/x/gov/types/keys.go#L9
	GovModuleName = "gov"

	// MintModuleName duplicates the mint module's name to avoid a cyclic dependency with x/mint.
	// It should be synced with the mint module's name if it is ever changed.
	// See: https://github.com/cosmos/cosmos-sdk/blob/0e34478eb7420b69869ed50f129fc274a97a9b06/x/mint/types/keys.go#L13
	MintModuleName = "mint"
)

// KVStore keys
var (
	SupplyKey           = collections.NewPrefix(0)
	DenomMetadataPrefix = collections.NewPrefix(1)
	// BalancesPrefix is the prefix for the account balances store. We use a byte
	// (instead of `[]byte("balances")` to save some disk space).
	BalancesPrefix     = collections.NewPrefix(2)
	DenomAddressPrefix = collections.NewPrefix(3)
	// SendEnabledPrefix is the prefix for the SendDisabled flags for a Denom.
	SendEnabledPrefix = collections.NewPrefix(4)

	// ParamsKey is the prefix for x/bank parameters
	ParamsKey = collections.NewPrefix(5)
)

// NewBalanceCompatValueCodec is a codec for encoding Balances in a backwards compatible way
// with respect to the old format.
func NewBalanceCompatValueCodec() collcodec.ValueCodec[math.Int] {
	return balanceCompatValueCodec{
		sdk.IntValue,
	}
}

type balanceCompatValueCodec struct {
	collcodec.ValueCodec[math.Int]
}

func (v balanceCompatValueCodec) Decode(b []byte) (math.Int, error) {
	i, err := v.ValueCodec.Decode(b)
	if err == nil {
		return i, nil
	}
	c := new(sdk.Coin)
	err = c.Unmarshal(b)
	if err != nil {
		return math.Int{}, err
	}
	return c.Amount, nil
}
