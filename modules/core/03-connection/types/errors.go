package types

import (
	errorsmod "cosmossdk.io/errors"
)

// IBC connection sentinel errors
var (
	ErrConnectionExists              = errorsmod.Register(SubModuleName, 112, "connection already exists")
	ErrConnectionNotFound            = errorsmod.Register(SubModuleName, 113, "connection not found")
	ErrClientConnectionPathsNotFound = errorsmod.Register(SubModuleName, 114, "light client connection paths not found")
	ErrConnectionPath                = errorsmod.Register(SubModuleName, 115, "connection path is not associated to the given light client")
	ErrInvalidConnectionState        = errorsmod.Register(SubModuleName, 116, "invalid connection state")
	ErrInvalidCounterparty           = errorsmod.Register(SubModuleName, 117, "invalid counterparty connection")
	ErrInvalidConnection             = errorsmod.Register(SubModuleName, 118, "invalid connection")
	ErrInvalidVersion                = errorsmod.Register(SubModuleName, 119, "invalid connection version")
	ErrVersionNegotiationFailed      = errorsmod.Register(SubModuleName, 1110, "connection version negotiation failed")
	ErrInvalidConnectionIdentifier   = errorsmod.Register(SubModuleName, 1111, "invalid connection identifier")
)
