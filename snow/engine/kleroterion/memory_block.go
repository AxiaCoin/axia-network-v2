// Copyright (C) 2019-2021, Axia Systems, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package kleroterion

import (
	"github.com/axiacoin/axia-network-v2/snow/consensus/kleroterion"
)

// memoryBlock wraps a kleroterion Block to manage non-verified blocks
type memoryBlock struct {
	kleroterion.Block

	tree    AncestorTree
	metrics *metrics
}

// Accept accepts the underlying block & removes sibling subtrees
func (mb *memoryBlock) Accept() error {
	mb.tree.RemoveSubtree(mb.Parent())
	mb.metrics.numNonVerifieds.Set(float64(mb.tree.Len()))
	return mb.Block.Accept()
}

// Reject rejects the underlying block & removes child subtrees
func (mb *memoryBlock) Reject() error {
	mb.tree.RemoveSubtree(mb.ID())
	mb.metrics.numNonVerifieds.Set(float64(mb.tree.Len()))
	return mb.Block.Reject()
}
