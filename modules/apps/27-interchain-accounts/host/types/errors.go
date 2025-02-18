package types

import (
	errorsmod "cosmossdk.io/errors"
)

// ICA Host sentinel errors
var (
	ErrHostSubModuleDisabled = errorsmod.Register(SubModuleName, 112, "host submodule is disabled")
)
