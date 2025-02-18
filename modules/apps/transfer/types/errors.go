package types

import (
	errorsmod "cosmossdk.io/errors"
)

// IBC transfer sentinel errors
var (
	ErrInvalidPacketTimeout    = errorsmod.Register(ModuleName, 112, "invalid packet timeout")
	ErrInvalidDenomForTransfer = errorsmod.Register(ModuleName, 113, "invalid denomination for cross-chain transfer")
	ErrInvalidVersion          = errorsmod.Register(ModuleName, 114, "invalid ICS20 version")
	ErrInvalidAmount           = errorsmod.Register(ModuleName, 115, "invalid token amount")
	ErrDenomNotFound           = errorsmod.Register(ModuleName, 116, "denomination not found")
	ErrSendDisabled            = errorsmod.Register(ModuleName, 117, "fungible token transfers from this chain are disabled")
	ErrReceiveDisabled         = errorsmod.Register(ModuleName, 118, "fungible token transfers to this chain are disabled")
	ErrMaxTransferChannels     = errorsmod.Register(ModuleName, 119, "max transfer channels")
	ErrInvalidAuthorization    = errorsmod.Register(ModuleName, 1110, "invalid transfer authorization")
	ErrInvalidMemo             = errorsmod.Register(ModuleName, 1111, "invalid memo")
	ErrForwardedPacketTimedOut = errorsmod.Register(ModuleName, 1112, "forwarded packet timed out")
	ErrForwardedPacketFailed   = errorsmod.Register(ModuleName, 1113, "forwarded packet failed")
	ErrAbiEncoding             = errorsmod.Register(ModuleName, 1114, "encoding abi failed")
	ErrAbiDecoding             = errorsmod.Register(ModuleName, 1115, "decoding abi failed")
	ErrReceiveFailed           = errorsmod.Register(ModuleName, 1116, "receive packet failed")
)
