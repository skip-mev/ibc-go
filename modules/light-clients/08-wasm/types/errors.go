package types

import errorsmod "cosmossdk.io/errors"

var (
	ErrInvalid              = errorsmod.Register(ModuleName, 112, "invalid")
	ErrInvalidData          = errorsmod.Register(ModuleName, 113, "invalid data")
	ErrInvalidChecksum      = errorsmod.Register(ModuleName, 114, "invalid checksum")
	ErrInvalidClientMessage = errorsmod.Register(ModuleName, 115, "invalid client message")
	ErrRetrieveClientID     = errorsmod.Register(ModuleName, 116, "failed to retrieve client id")
	// Wasm specific
	ErrWasmEmptyCode                   = errorsmod.Register(ModuleName, 117, "empty wasm code")
	ErrWasmCodeTooLarge                = errorsmod.Register(ModuleName, 118, "wasm code too large")
	ErrWasmCodeExists                  = errorsmod.Register(ModuleName, 119, "wasm code already exists")
	ErrWasmChecksumNotFound            = errorsmod.Register(ModuleName, 1110, "wasm checksum not found")
	ErrWasmSubMessagesNotAllowed       = errorsmod.Register(ModuleName, 1111, "execution of sub messages is not allowed")
	ErrWasmEventsNotAllowed            = errorsmod.Register(ModuleName, 1112, "returning events from a contract is not allowed")
	ErrWasmAttributesNotAllowed        = errorsmod.Register(ModuleName, 1113, "returning attributes from a contract is not allowed")
	ErrWasmContractCallFailed          = errorsmod.Register(ModuleName, 1114, "wasm contract call failed")
	ErrWasmInvalidResponseData         = errorsmod.Register(ModuleName, 1115, "wasm contract returned invalid response data")
	ErrWasmInvalidContractModification = errorsmod.Register(ModuleName, 1116, "wasm contract made invalid state modifications")
	ErrVMError                         = errorsmod.Register(ModuleName, 1117, "wasm VM error")
)
