package host

import (
	errorsmod "cosmossdk.io/errors"
)

// SubModuleName defines the ICS 24 host
const SubModuleName = "host"

// IBC client sentinel errors
var (
	ErrInvalidID     = errorsmod.Register(SubModuleName, 12, "invalid identifier")
	ErrInvalidPath   = errorsmod.Register(SubModuleName, 13, "invalid path")
	ErrInvalidPacket = errorsmod.Register(SubModuleName, 14, "invalid packet")
)
