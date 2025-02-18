package types

import (
	errorsmod "cosmossdk.io/errors"
)

// SubModuleName is the error codespace
const SubModuleName string = "commitment"

// IBC connection sentinel errors
var (
	ErrInvalidProof       = errorsmod.Register(SubModuleName, 12, "invalid proof")
	ErrInvalidPrefix      = errorsmod.Register(SubModuleName, 13, "invalid prefix")
	ErrInvalidMerkleProof = errorsmod.Register(SubModuleName, 14, "invalid merkle proof")
)
