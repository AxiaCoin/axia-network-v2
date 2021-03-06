// Copyright (C) 2019-2021, Axia Systems, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package rpcchainvm

import (
	"github.com/axiacoin/axia-network-v2/database"
	"github.com/axiacoin/axia-network-v2/snow/engine/kleroterion/block"
)

var (
	errCodeToError = map[uint32]error{
		1: database.ErrClosed,
		2: database.ErrNotFound,
		3: block.ErrHeightIndexedVMNotImplemented,
		4: block.ErrIndexIncomplete,
	}
	errorToErrCode = map[error]uint32{
		database.ErrClosed:                     1,
		database.ErrNotFound:                   2,
		block.ErrHeightIndexedVMNotImplemented: 3,
		block.ErrIndexIncomplete:               4,
	}
)

func errorToRPCError(err error) error {
	if _, ok := errorToErrCode[err]; ok {
		return nil
	}
	return err
}
