// This file only used to generate mocks

package testutil

import (
	"context"

	bank "cosmossdk.io/x/bank/types"
	"cosmossdk.io/x/group"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// AccountKeeper extends `AccountKeeper` from expected_keepers.
type AccountKeeper interface {
	group.AccountKeeper
}

// BankKeeper extends bank `MsgServer` to mock `Send` and to register handlers in MsgServiceRouter
type BankKeeper interface {
	group.BankKeeper
	bank.MsgServer

	MintCoins(ctx context.Context, moduleName string, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx context.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	GetAllBalances(ctx context.Context, addr sdk.AccAddress) sdk.Coins
}
