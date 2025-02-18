package types

import (
	errorsmod "cosmossdk.io/errors"
)

var (
	ErrInvalidPacket            = errorsmod.Register(SubModuleName, 112, "invalid packet")
	ErrInvalidPayload           = errorsmod.Register(SubModuleName, 113, "invalid payload")
	ErrSequenceSendNotFound     = errorsmod.Register(SubModuleName, 114, "sequence send not found")
	ErrInvalidAcknowledgement   = errorsmod.Register(SubModuleName, 115, "invalid acknowledgement")
	ErrPacketCommitmentNotFound = errorsmod.Register(SubModuleName, 116, "packet commitment not found")
	ErrAcknowledgementNotFound  = errorsmod.Register(SubModuleName, 117, "packet acknowledgement not found")
	ErrInvalidTimeout           = errorsmod.Register(SubModuleName, 118, "invalid packet timeout")
	ErrTimeoutElapsed           = errorsmod.Register(SubModuleName, 119, "timeout elapsed")
	ErrTimeoutNotReached        = errorsmod.Register(SubModuleName, 1110, "timeout not reached")
	ErrAcknowledgementExists    = errorsmod.Register(SubModuleName, 1111, "acknowledgement for packet already exists")
	ErrNoOpMsg                  = errorsmod.Register(SubModuleName, 1112, "message is redundant, no-op will be performed")
)
