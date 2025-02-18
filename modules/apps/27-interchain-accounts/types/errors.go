package types

import (
	errorsmod "cosmossdk.io/errors"
)

var (
	ErrUnknownDataType             = errorsmod.Register(ModuleName, 112, "unknown data type")
	ErrAccountAlreadyExist         = errorsmod.Register(ModuleName, 113, "account already exist")
	ErrPortAlreadyBound            = errorsmod.Register(ModuleName, 114, "port is already bound")
	ErrInvalidChannelFlow          = errorsmod.Register(ModuleName, 115, "invalid message sent to channel end")
	ErrInvalidOutgoingData         = errorsmod.Register(ModuleName, 116, "invalid outgoing data")
	ErrInvalidRoute                = errorsmod.Register(ModuleName, 117, "invalid route")
	ErrInterchainAccountNotFound   = errorsmod.Register(ModuleName, 118, "interchain account not found")
	ErrInterchainAccountAlreadySet = errorsmod.Register(ModuleName, 119, "interchain account is already set")
	ErrActiveChannelAlreadySet     = errorsmod.Register(ModuleName, 1110, "active channel already set for this owner")
	ErrActiveChannelNotFound       = errorsmod.Register(ModuleName, 1111, "no active channel for this owner")
	ErrInvalidVersion              = errorsmod.Register(ModuleName, 1112, "invalid interchain accounts version")
	ErrInvalidAccountAddress       = errorsmod.Register(ModuleName, 1113, "invalid account address")
	ErrUnsupported                 = errorsmod.Register(ModuleName, 1114, "interchain account does not support this action")
	ErrInvalidControllerPort       = errorsmod.Register(ModuleName, 1115, "invalid controller port")
	ErrInvalidHostPort             = errorsmod.Register(ModuleName, 1116, "invalid host port")
	ErrInvalidTimeoutTimestamp     = errorsmod.Register(ModuleName, 1117, "timeout timestamp must be in the future")
	ErrInvalidCodec                = errorsmod.Register(ModuleName, 1118, "codec is not supported")
	ErrInvalidAccountReopening     = errorsmod.Register(ModuleName, 1119, "invalid account reopening")
)
