package types

import (
	errorsmod "cosmossdk.io/errors"
)

// IBC connection sentinel errors
var (
	ErrConnectionExists              = errorsmod.Register(SubModuleName, 12, "connection already exists")
	ErrConnectionNotFound            = errorsmod.Register(SubModuleName, 13, "connection not found")
	ErrClientConnectionPathsNotFound = errorsmod.Register(SubModuleName, 14, "light client connection paths not found")
	ErrConnectionPath                = errorsmod.Register(SubModuleName, 15, "connection path is not associated to the given light client")
	ErrInvalidConnectionState        = errorsmod.Register(SubModuleName, 16, "invalid connection state")
	ErrInvalidCounterparty           = errorsmod.Register(SubModuleName, 17, "invalid counterparty connection")
	ErrInvalidConnection             = errorsmod.Register(SubModuleName, 18, "invalid connection")
	ErrInvalidVersion                = errorsmod.Register(SubModuleName, 19, "invalid connection version")
	ErrVersionNegotiationFailed      = errorsmod.Register(SubModuleName, 110, "connection version negotiation failed")
	ErrInvalidConnectionIdentifier   = errorsmod.Register(SubModuleName, 111, "invalid connection identifier")
)
