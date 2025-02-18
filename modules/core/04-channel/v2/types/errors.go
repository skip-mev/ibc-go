package types

import (
	errorsmod "cosmossdk.io/errors"
)

var (
	ErrInvalidPacket            = errorsmod.Register(SubModuleName, 12, "invalid packet")
	ErrInvalidPayload           = errorsmod.Register(SubModuleName, 13, "invalid payload")
	ErrSequenceSendNotFound     = errorsmod.Register(SubModuleName, 14, "sequence send not found")
	ErrInvalidAcknowledgement   = errorsmod.Register(SubModuleName, 15, "invalid acknowledgement")
	ErrPacketCommitmentNotFound = errorsmod.Register(SubModuleName, 16, "packet commitment not found")
	ErrAcknowledgementNotFound  = errorsmod.Register(SubModuleName, 17, "packet acknowledgement not found")
	ErrInvalidTimeout           = errorsmod.Register(SubModuleName, 18, "invalid packet timeout")
	ErrTimeoutElapsed           = errorsmod.Register(SubModuleName, 19, "timeout elapsed")
	ErrTimeoutNotReached        = errorsmod.Register(SubModuleName, 110, "timeout not reached")
	ErrAcknowledgementExists    = errorsmod.Register(SubModuleName, 111, "acknowledgement for packet already exists")
	ErrNoOpMsg                  = errorsmod.Register(SubModuleName, 112, "message is redundant, no-op will be performed")
)
