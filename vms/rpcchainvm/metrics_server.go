// Copyright (C) 2019-2021, Axia Systems, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package rpcchainvm

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	vmpb "github.com/axiacoin/axia-network-v2/proto/pb/vm"
)

func (vm *VMServer) Gather(context.Context, *emptypb.Empty) (*vmpb.GatherResponse, error) {
	mfs, err := vm.ctx.Metrics.Gather()
	return &vmpb.GatherResponse{MetricFamilies: mfs}, err
}
