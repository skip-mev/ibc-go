package types

import (
	errorsmod "cosmossdk.io/errors"
)

// IBC port sentinel errors
var (
	ErrPortExists   = errorsmod.Register(SubModuleName, 112, "port is already binded")
	ErrPortNotFound = errorsmod.Register(SubModuleName, 113, "port not found")
	ErrInvalidPort  = errorsmod.Register(SubModuleName, 114, "invalid port")
	ErrInvalidRoute = errorsmod.Register(SubModuleName, 115, "route not found")
)
