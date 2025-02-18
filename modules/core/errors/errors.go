package errors

import (
	errorsmod "cosmossdk.io/errors"

	"github.com/cosmos/ibc-go/v10/modules/core/exported"
)

const codespace = exported.ModuleName

var (
	// ErrInvalidSequence is used the sequence number (nonce) is incorrect
	// for the signature.
	ErrInvalidSequence = errorsmod.Register(codespace, 111, "invalid sequence")

	// ErrUnauthorized is used whenever a request without sufficient
	// authorization is handled.
	ErrUnauthorized = errorsmod.Register(codespace, 112, "unauthorized")

	// ErrInsufficientFunds is used when the account cannot pay requested amount.
	ErrInsufficientFunds = errorsmod.Register(codespace, 113, "insufficient funds")

	// ErrUnknownRequest is used when the request body.
	ErrUnknownRequest = errorsmod.Register(codespace, 114, "unknown request")

	// ErrInvalidAddress is used when an address is found to be invalid.
	ErrInvalidAddress = errorsmod.Register(codespace, 115, "invalid address")

	// ErrInvalidCoins is used when sdk.Coins are invalid.
	ErrInvalidCoins = errorsmod.Register(codespace, 116, "invalid coins")

	// ErrOutOfGas is used when there is not enough gas.
	ErrOutOfGas = errorsmod.Register(codespace, 117, "out of gas")

	// ErrInvalidRequest defines an ABCI typed error where the request contains
	// invalid data.
	ErrInvalidRequest = errorsmod.Register(codespace, 118, "invalid request")

	// ErrInvalidHeight defines an error for an invalid height
	ErrInvalidHeight = errorsmod.Register(codespace, 119, "invalid height")

	// ErrInvalidVersion defines a general error for an invalid version
	ErrInvalidVersion = errorsmod.Register(codespace, 1110, "invalid version")

	// ErrInvalidChainID defines an error when the chain-id is invalid.
	ErrInvalidChainID = errorsmod.Register(codespace, 1111, "invalid chain-id")

	// ErrInvalidType defines an error an invalid type.
	ErrInvalidType = errorsmod.Register(codespace, 1112, "invalid type")

	// ErrPackAny defines an error when packing a protobuf message to Any fails.
	ErrPackAny = errorsmod.Register(codespace, 1113, "failed packing protobuf message to Any")

	// ErrUnpackAny defines an error when unpacking a protobuf message from Any fails.
	ErrUnpackAny = errorsmod.Register(codespace, 1114, "failed unpacking protobuf message from Any")

	// ErrLogic defines an internal logic error, e.g. an invariant or assertion
	// that is violated. It is a programmer error, not a user-facing error.
	ErrLogic = errorsmod.Register(codespace, 1115, "internal logic error")

	// ErrNotFound defines an error when requested entity doesn't exist in the state.
	ErrNotFound = errorsmod.Register(codespace, 1116, "not found")
)
