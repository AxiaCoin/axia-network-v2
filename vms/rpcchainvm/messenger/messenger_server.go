// Copyright (C) 2019-2021, Axia Systems, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package messenger

import (
	"context"
	"errors"

	"github.com/axiacoin/axia-network-v2/snow/engine/common"

	messengerpb "github.com/axiacoin/axia-network-v2/proto/pb/messenger"
)

var (
	errFullQueue = errors.New("full message queue")

	_ messengerpb.MessengerServer = &Server{}
)

// Server is a messenger that is managed over RPC.
type Server struct {
	messengerpb.UnimplementedMessengerServer
	messenger chan<- common.Message
}

// NewServer returns a messenger connected to a remote channel
func NewServer(messenger chan<- common.Message) *Server {
	return &Server{messenger: messenger}
}

func (s *Server) Notify(_ context.Context, req *messengerpb.NotifyRequest) (*messengerpb.NotifyResponse, error) {
	msg := common.Message(req.Message)
	select {
	case s.messenger <- msg:
		return &messengerpb.NotifyResponse{}, nil
	default:
		return nil, errFullQueue
	}
}
