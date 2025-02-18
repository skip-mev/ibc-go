package types

import (
	errorsmod "cosmossdk.io/errors"
)

// IBC channel sentinel errors
var (
	ErrChannelExists           = errorsmod.Register(SubModuleName, 112, "channel already exists")
	ErrChannelNotFound         = errorsmod.Register(SubModuleName, 113, "channel not found")
	ErrInvalidChannel          = errorsmod.Register(SubModuleName, 114, "invalid channel")
	ErrInvalidChannelState     = errorsmod.Register(SubModuleName, 115, "invalid channel state")
	ErrInvalidChannelOrdering  = errorsmod.Register(SubModuleName, 116, "invalid channel ordering")
	ErrInvalidCounterparty     = errorsmod.Register(SubModuleName, 117, "invalid counterparty channel")
	ErrSequenceSendNotFound    = errorsmod.Register(SubModuleName, 1110, "sequence send not found")
	ErrSequenceReceiveNotFound = errorsmod.Register(SubModuleName, 1111, "sequence receive not found")
	ErrSequenceAckNotFound     = errorsmod.Register(SubModuleName, 1112, "sequence acknowledgement not found")
	ErrInvalidPacket           = errorsmod.Register(SubModuleName, 1113, "invalid packet")

	// Deprecated: ErrPacketTimeout is deprecated and will be removed in a future release.
	// Please use ErrTimeoutElapsed instead.
	ErrPacketTimeout = errorsmod.Register(SubModuleName, 1114, "packet timeout")

	ErrTooManyConnectionHops    = errorsmod.Register(SubModuleName, 1115, "too many connection hops")
	ErrInvalidAcknowledgement   = errorsmod.Register(SubModuleName, 1116, "invalid acknowledgement")
	ErrAcknowledgementExists    = errorsmod.Register(SubModuleName, 1117, "acknowledgement for packet already exists")
	ErrInvalidChannelIdentifier = errorsmod.Register(SubModuleName, 1118, "invalid channel identifier")

	// packets already relayed errors
	ErrPacketReceived           = errorsmod.Register(SubModuleName, 1119, "packet already received")
	ErrPacketCommitmentNotFound = errorsmod.Register(SubModuleName, 1120, "packet commitment not found") // may occur for already received acknowledgements or timeouts and in rare cases for packets never sent

	// ORDERED channel error
	ErrPacketSequenceOutOfOrder = errorsmod.Register(SubModuleName, 1121, "packet sequence is out of order")

	// Antehandler error
	ErrRedundantTx = errorsmod.Register(SubModuleName, 1122, "packet messages are redundant")

	// Perform a no-op on the current Msg
	ErrNoOpMsg = errorsmod.Register(SubModuleName, 1123, "message is redundant, no-op will be performed")

	ErrInvalidChannelVersion           = errorsmod.Register(SubModuleName, 1124, "invalid channel version")
	ErrPacketNotSent                   = errorsmod.Register(SubModuleName, 1125, "packet has not been sent")
	ErrInvalidTimeout                  = errorsmod.Register(SubModuleName, 1126, "invalid packet timeout")
	ErrUpgradeErrorNotFound            = errorsmod.Register(SubModuleName, 1127, "upgrade error receipt not found")
	ErrInvalidUpgrade                  = errorsmod.Register(SubModuleName, 1128, "invalid upgrade")
	ErrInvalidUpgradeSequence          = errorsmod.Register(SubModuleName, 1129, "invalid upgrade sequence")
	ErrUpgradeNotFound                 = errorsmod.Register(SubModuleName, 1130, "upgrade not found")
	ErrIncompatibleCounterpartyUpgrade = errorsmod.Register(SubModuleName, 1131, "incompatible counterparty upgrade")
	ErrInvalidUpgradeError             = errorsmod.Register(SubModuleName, 1132, "invalid upgrade error")
	ErrUpgradeRestoreFailed            = errorsmod.Register(SubModuleName, 1133, "restore failed")
	ErrUpgradeTimeout                  = errorsmod.Register(SubModuleName, 1134, "upgrade timed-out")
	ErrInvalidUpgradeTimeout           = errorsmod.Register(SubModuleName, 1135, "upgrade timeout is invalid")
	ErrPendingInflightPackets          = errorsmod.Register(SubModuleName, 1136, "pending inflight packets exist")
	ErrUpgradeTimeoutFailed            = errorsmod.Register(SubModuleName, 1137, "upgrade timeout failed")
	ErrInvalidPruningLimit             = errorsmod.Register(SubModuleName, 1138, "invalid pruning limit")
	ErrTimeoutNotReached               = errorsmod.Register(SubModuleName, 1139, "timeout not reached")
	ErrTimeoutElapsed                  = errorsmod.Register(SubModuleName, 1140, "timeout elapsed")
	ErrPruningSequenceStartNotFound    = errorsmod.Register(SubModuleName, 1141, "pruning sequence start not found")
	ErrRecvStartSequenceNotFound       = errorsmod.Register(SubModuleName, 1142, "recv start sequence not found")
	ErrInvalidCommitment               = errorsmod.Register(SubModuleName, 1143, "invalid commitment")
)
