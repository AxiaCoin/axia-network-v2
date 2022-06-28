// Copyright (C) 2019-2021, Axia Systems, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package peer

import (
	"time"

	"github.com/axiacoin/axia-network-v2/ids"
	"github.com/axiacoin/axia-network-v2/message"
	"github.com/axiacoin/axia-network-v2/network/throttling"
	"github.com/axiacoin/axia-network-v2/snow/networking/router"
	"github.com/axiacoin/axia-network-v2/snow/validators"
	"github.com/axiacoin/axia-network-v2/utils/logging"
	"github.com/axiacoin/axia-network-v2/utils/timer/mockable"
	"github.com/axiacoin/axia-network-v2/version"
)

type Config struct {
	// Size, in bytes, of the buffer this peer reads messages into
	ReadBufferSize int
	// Size, in bytes, of the buffer this peer writes messages into
	WriteBufferSize      int
	Clock                mockable.Clock
	Metrics              *Metrics
	MessageCreator       message.Creator
	Log                  logging.Logger
	InboundMsgThrottler  throttling.InboundMsgThrottler
	OutboundMsgThrottler throttling.OutboundMsgThrottler
	Network              Network
	Router               router.InboundHandler
	VersionCompatibility version.Compatibility
	VersionParser        version.ApplicationParser
	MyAllychains            ids.Set
	Beacons              validators.Set
	NetworkID            uint32
	PingFrequency        time.Duration
	PongTimeout          time.Duration
	MaxClockDifference   time.Duration

	// Unix time of the last message sent and received respectively
	// Must only be accessed atomically
	LastSent, LastReceived int64
}
