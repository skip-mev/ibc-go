package types

import errorsmod "cosmossdk.io/errors"

var (
	ErrInvalidCounterparty  = errorsmod.Register(SubModuleName, 1134, "invalid counterparty")
	ErrCounterpartyNotFound = errorsmod.Register(SubModuleName, 1135, "counterparty not found")
)
