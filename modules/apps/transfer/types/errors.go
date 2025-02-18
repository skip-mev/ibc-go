package types

import (
	errorsmod "cosmossdk.io/errors"
)

// IBC transfer sentinel errors
var (
	ErrInvalidPacketTimeout    = errorsmod.Register(ModuleName, 12, "invalid packet timeout")
	ErrInvalidDenomForTransfer = errorsmod.Register(ModuleName, 13, "invalid denomination for cross-chain transfer")
	ErrInvalidVersion          = errorsmod.Register(ModuleName, 14, "invalid ICS20 version")
	ErrInvalidAmount           = errorsmod.Register(ModuleName, 15, "invalid token amount")
	ErrDenomNotFound           = errorsmod.Register(ModuleName, 16, "denomination not found")
	ErrSendDisabled            = errorsmod.Register(ModuleName, 17, "fungible token transfers from this chain are disabled")
	ErrReceiveDisabled         = errorsmod.Register(ModuleName, 18, "fungible token transfers to this chain are disabled")
	ErrMaxTransferChannels     = errorsmod.Register(ModuleName, 19, "max transfer channels")
	ErrInvalidAuthorization    = errorsmod.Register(ModuleName, 110, "invalid transfer authorization")
	ErrInvalidMemo             = errorsmod.Register(ModuleName, 111, "invalid memo")
	ErrForwardedPacketTimedOut = errorsmod.Register(ModuleName, 112, "forwarded packet timed out")
	ErrForwardedPacketFailed   = errorsmod.Register(ModuleName, 113, "forwarded packet failed")
	ErrAbiEncoding             = errorsmod.Register(ModuleName, 114, "encoding abi failed")
	ErrAbiDecoding             = errorsmod.Register(ModuleName, 115, "decoding abi failed")
	ErrReceiveFailed           = errorsmod.Register(ModuleName, 116, "receive packet failed")
)
