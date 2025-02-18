package tendermint

import (
	errorsmod "cosmossdk.io/errors"
)

// IBC tendermint client sentinel errors
var (
	ErrInvalidChainID          = errorsmod.Register(ModuleName, 12, "invalid chain-id")
	ErrInvalidTrustingPeriod   = errorsmod.Register(ModuleName, 13, "invalid trusting period")
	ErrInvalidUnbondingPeriod  = errorsmod.Register(ModuleName, 14, "invalid unbonding period")
	ErrInvalidHeaderHeight     = errorsmod.Register(ModuleName, 15, "invalid header height")
	ErrInvalidHeader           = errorsmod.Register(ModuleName, 16, "invalid header")
	ErrInvalidMaxClockDrift    = errorsmod.Register(ModuleName, 17, "invalid max clock drift")
	ErrProcessedTimeNotFound   = errorsmod.Register(ModuleName, 18, "processed time not found")
	ErrProcessedHeightNotFound = errorsmod.Register(ModuleName, 19, "processed height not found")
	ErrDelayPeriodNotPassed    = errorsmod.Register(ModuleName, 110, "packet-specified delay period has not been reached")
	ErrTrustingPeriodExpired   = errorsmod.Register(ModuleName, 111, "time since latest trusted state has passed the trusting period")
	ErrUnbondingPeriodExpired  = errorsmod.Register(ModuleName, 112, "time since latest trusted state has passed the unbonding period")
	ErrInvalidProofSpecs       = errorsmod.Register(ModuleName, 113, "invalid proof specs")
	ErrInvalidValidatorSet     = errorsmod.Register(ModuleName, 114, "invalid validator set")
	ErrInvalidTrustLevel       = errorsmod.Register(ModuleName, 115, "invalid trust level")
)
