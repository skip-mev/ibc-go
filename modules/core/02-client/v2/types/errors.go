package types

import errorsmod "cosmossdk.io/errors"

var (
	ErrInvalidCounterparty  = errorsmod.Register(SubModuleName, 134, "invalid counterparty")
	ErrCounterpartyNotFound = errorsmod.Register(SubModuleName, 135, "counterparty not found")
)
