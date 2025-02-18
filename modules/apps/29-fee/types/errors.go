package types

import (
	errorsmod "cosmossdk.io/errors"
)

// 29-fee sentinel errors
var (
	ErrInvalidVersion                = errorsmod.Register(ModuleName, 12, "invalid ICS29 middleware version")
	ErrRefundAccNotFound             = errorsmod.Register(ModuleName, 13, "no account found for given refund address")
	ErrBalanceNotFound               = errorsmod.Register(ModuleName, 14, "balance not found for given account address")
	ErrFeeNotFound                   = errorsmod.Register(ModuleName, 15, "there is no fee escrowed for the given packetID")
	ErrRelayersNotEmpty              = errorsmod.Register(ModuleName, 16, "relayers must not be set. This feature is not supported")
	ErrCounterpartyPayeeEmpty        = errorsmod.Register(ModuleName, 17, "counterparty payee must not be empty")
	ErrForwardRelayerAddressNotFound = errorsmod.Register(ModuleName, 18, "forward relayer address not found")
	ErrFeeNotEnabled                 = errorsmod.Register(ModuleName, 19, "fee module is not enabled for this channel. If this error occurs after channel setup, fee module may not be enabled")
	ErrRelayerNotFoundForAsyncAck    = errorsmod.Register(ModuleName, 110, "relayer address must be stored for async WriteAcknowledgement")
	ErrFeeModuleLocked               = errorsmod.Register(ModuleName, 111, "the fee module is currently locked, a severe bug has been detected")
	ErrUnsupportedAction             = errorsmod.Register(ModuleName, 112, "unsupported action")
)
