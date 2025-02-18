package solomachine

import (
	errorsmod "cosmossdk.io/errors"
)

var (
	ErrInvalidHeader               = errorsmod.Register(ModuleName, 112, "invalid header")
	ErrInvalidSequence             = errorsmod.Register(ModuleName, 113, "invalid sequence")
	ErrInvalidSignatureAndData     = errorsmod.Register(ModuleName, 114, "invalid signature and data")
	ErrSignatureVerificationFailed = errorsmod.Register(ModuleName, 115, "signature verification failed")
	ErrInvalidProof                = errorsmod.Register(ModuleName, 116, "invalid solo machine proof")
)
