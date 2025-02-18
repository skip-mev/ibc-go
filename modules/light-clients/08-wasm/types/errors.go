package types

import errorsmod "cosmossdk.io/errors"

var (
	ErrInvalid              = errorsmod.Register(ModuleName, 12, "invalid")
	ErrInvalidData          = errorsmod.Register(ModuleName, 13, "invalid data")
	ErrInvalidChecksum      = errorsmod.Register(ModuleName, 14, "invalid checksum")
	ErrInvalidClientMessage = errorsmod.Register(ModuleName, 15, "invalid client message")
	ErrRetrieveClientID     = errorsmod.Register(ModuleName, 16, "failed to retrieve client id")
	// Wasm specific
	ErrWasmEmptyCode                   = errorsmod.Register(ModuleName, 17, "empty wasm code")
	ErrWasmCodeTooLarge                = errorsmod.Register(ModuleName, 18, "wasm code too large")
	ErrWasmCodeExists                  = errorsmod.Register(ModuleName, 19, "wasm code already exists")
	ErrWasmChecksumNotFound            = errorsmod.Register(ModuleName, 110, "wasm checksum not found")
	ErrWasmSubMessagesNotAllowed       = errorsmod.Register(ModuleName, 111, "execution of sub messages is not allowed")
	ErrWasmEventsNotAllowed            = errorsmod.Register(ModuleName, 112, "returning events from a contract is not allowed")
	ErrWasmAttributesNotAllowed        = errorsmod.Register(ModuleName, 113, "returning attributes from a contract is not allowed")
	ErrWasmContractCallFailed          = errorsmod.Register(ModuleName, 114, "wasm contract call failed")
	ErrWasmInvalidResponseData         = errorsmod.Register(ModuleName, 115, "wasm contract returned invalid response data")
	ErrWasmInvalidContractModification = errorsmod.Register(ModuleName, 116, "wasm contract made invalid state modifications")
	ErrVMError                         = errorsmod.Register(ModuleName, 117, "wasm VM error")
)
