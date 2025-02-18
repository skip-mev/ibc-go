package types

import (
	errorsmod "cosmossdk.io/errors"
)

var (
	ErrCannotUnmarshalPacketData = errorsmod.Register(ModuleName, 112, "cannot unmarshal packet data")
	ErrNotPacketDataProvider     = errorsmod.Register(ModuleName, 113, "packet is not a PacketDataProvider")
	ErrCallbackKeyNotFound       = errorsmod.Register(ModuleName, 114, "callback key not found in packet data")
	ErrCallbackAddressNotFound   = errorsmod.Register(ModuleName, 115, "callback address not found in packet data")
	ErrCallbackOutOfGas          = errorsmod.Register(ModuleName, 116, "callback out of gas")
	ErrCallbackPanic             = errorsmod.Register(ModuleName, 117, "callback panic")
)
