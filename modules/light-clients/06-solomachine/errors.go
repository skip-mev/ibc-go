package solomachine

import (
	errorsmod "cosmossdk.io/errors"
)

var (
	ErrInvalidHeader               = errorsmod.Register(ModuleName, 12, "invalid header")
	ErrInvalidSequence             = errorsmod.Register(ModuleName, 13, "invalid sequence")
	ErrInvalidSignatureAndData     = errorsmod.Register(ModuleName, 14, "invalid signature and data")
	ErrSignatureVerificationFailed = errorsmod.Register(ModuleName, 15, "signature verification failed")
	ErrInvalidProof                = errorsmod.Register(ModuleName, 16, "invalid solo machine proof")
)
