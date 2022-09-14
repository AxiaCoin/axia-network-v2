// Copyright (C) 2019-2021, Axia Systems, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package block

import (
	"errors"
	"testing"
	"time"

	"github.com/axiacoin/axia-network-v2/ids"
	"github.com/axiacoin/axia-network-v2/snow/consensus/kleroterion"
)

var (
	errGetAncestor       = errors.New("unexpectedly called GetAncestor")
	errBatchedParseBlock = errors.New("unexpectedly called BatchedParseBlock")

	_ BatchedChainVM = &TestBatchedVM{}
)

// TestBatchedVM is a BatchedVM that is useful for testing.
type TestBatchedVM struct {
	T *testing.T

	CantGetAncestors    bool
	CantBatchParseBlock bool

	GetAncestorsF func(
		blkID ids.ID,
		maxBlocksNum int,
		maxBlocksSize int,
		maxBlocksRetrivalTime time.Duration,
	) ([][]byte, error)

	BatchedParseBlockF func(blks [][]byte) ([]kleroterion.Block, error)
}

func (vm *TestBatchedVM) Default(cant bool) {
	vm.CantGetAncestors = cant
	vm.CantBatchParseBlock = cant
}

func (vm *TestBatchedVM) GetAncestors(
	blkID ids.ID,
	maxBlocksNum int,
	maxBlocksSize int,
	maxBlocksRetrivalTime time.Duration,
) ([][]byte, error) {
	if vm.GetAncestorsF != nil {
		return vm.GetAncestorsF(blkID, maxBlocksNum, maxBlocksSize, maxBlocksRetrivalTime)
	}
	if vm.CantGetAncestors && vm.T != nil {
		vm.T.Fatal(errGetAncestor)
	}
	return nil, errGetAncestor
}

func (vm *TestBatchedVM) BatchedParseBlock(blks [][]byte) ([]kleroterion.Block, error) {
	if vm.BatchedParseBlockF != nil {
		return vm.BatchedParseBlockF(blks)
	}
	if vm.CantBatchParseBlock && vm.T != nil {
		vm.T.Fatal(errBatchedParseBlock)
	}
	return nil, errBatchedParseBlock
}
