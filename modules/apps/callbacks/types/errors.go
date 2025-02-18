package types

import (
	errorsmod "cosmossdk.io/errors"
)

var (
	ErrCannotUnmarshalPacketData = errorsmod.Register(ModuleName, 12, "cannot unmarshal packet data")
	ErrNotPacketDataProvider     = errorsmod.Register(ModuleName, 13, "packet is not a PacketDataProvider")
	ErrCallbackKeyNotFound       = errorsmod.Register(ModuleName, 14, "callback key not found in packet data")
	ErrCallbackAddressNotFound   = errorsmod.Register(ModuleName, 15, "callback address not found in packet data")
	ErrCallbackOutOfGas          = errorsmod.Register(ModuleName, 16, "callback out of gas")
	ErrCallbackPanic             = errorsmod.Register(ModuleName, 17, "callback panic")
)
