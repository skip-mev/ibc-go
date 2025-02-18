package types

import (
	errorsmod "cosmossdk.io/errors"
)

// IBC port sentinel errors
var (
	ErrPortExists   = errorsmod.Register(SubModuleName, 12, "port is already binded")
	ErrPortNotFound = errorsmod.Register(SubModuleName, 13, "port not found")
	ErrInvalidPort  = errorsmod.Register(SubModuleName, 14, "invalid port")
	ErrInvalidRoute = errorsmod.Register(SubModuleName, 15, "route not found")
)
