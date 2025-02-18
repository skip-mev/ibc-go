package tendermint

import (
	errorsmod "cosmossdk.io/errors"
)

// IBC tendermint client sentinel errors
var (
	ErrInvalidChainID          = errorsmod.Register(ModuleName, 112, "invalid chain-id")
	ErrInvalidTrustingPeriod   = errorsmod.Register(ModuleName, 113, "invalid trusting period")
	ErrInvalidUnbondingPeriod  = errorsmod.Register(ModuleName, 114, "invalid unbonding period")
	ErrInvalidHeaderHeight     = errorsmod.Register(ModuleName, 115, "invalid header height")
	ErrInvalidHeader           = errorsmod.Register(ModuleName, 116, "invalid header")
	ErrInvalidMaxClockDrift    = errorsmod.Register(ModuleName, 117, "invalid max clock drift")
	ErrProcessedTimeNotFound   = errorsmod.Register(ModuleName, 118, "processed time not found")
	ErrProcessedHeightNotFound = errorsmod.Register(ModuleName, 119, "processed height not found")
	ErrDelayPeriodNotPassed    = errorsmod.Register(ModuleName, 1110, "packet-specified delay period has not been reached")
	ErrTrustingPeriodExpired   = errorsmod.Register(ModuleName, 1111, "time since latest trusted state has passed the trusting period")
	ErrUnbondingPeriodExpired  = errorsmod.Register(ModuleName, 1112, "time since latest trusted state has passed the unbonding period")
	ErrInvalidProofSpecs       = errorsmod.Register(ModuleName, 1113, "invalid proof specs")
	ErrInvalidValidatorSet     = errorsmod.Register(ModuleName, 1114, "invalid validator set")
	ErrInvalidTrustLevel       = errorsmod.Register(ModuleName, 1115, "invalid trust level")
)
