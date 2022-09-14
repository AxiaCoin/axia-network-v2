// Copyright (C) 2019-2021, Axia Systems, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package metervm

import (
	"github.com/axiacoin/axia-network-v2/ids"
	"github.com/axiacoin/axia-network-v2/snow/engine/kleroterion/block"
)

var _ block.HeightIndexedChainVM = &blockVM{}

func (vm *blockVM) VerifyHeightIndex() error {
	hVM, ok := vm.ChainVM.(block.HeightIndexedChainVM)
	if !ok {
		return block.ErrHeightIndexedVMNotImplemented
	}
	return hVM.VerifyHeightIndex()
}

func (vm *blockVM) GetBlockIDAtHeight(height uint64) (ids.ID, error) {
	hVM, ok := vm.ChainVM.(block.HeightIndexedChainVM)
	if !ok {
		return ids.Empty, block.ErrHeightIndexedVMNotImplemented
	}
	return hVM.GetBlockIDAtHeight(height)
}
