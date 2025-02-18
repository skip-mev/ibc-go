package types

import (
	errorsmod "cosmossdk.io/errors"
)

// ICA Controller sentinel errors
var (
	ErrControllerSubModuleDisabled = errorsmod.Register(SubModuleName, 112, "controller submodule is disabled")
)
