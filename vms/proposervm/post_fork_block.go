// Copyright (C) 2019-2021, Axia Systems, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package proposervm

import (
	"github.com/axiacoin/axia-network-v2/ids"
	"github.com/axiacoin/axia-network-v2/snow/choices"
	"github.com/axiacoin/axia-network-v2/snow/consensus/kleroterion"
	"github.com/axiacoin/axia-network-v2/vms/proposervm/block"
)

var _ Block = &postForkBlock{}

type postForkBlock struct {
	block.SignedBlock
	postForkCommonComponents
}

// Accept:
// 1) Sets this blocks status to Accepted.
// 2) Persists this block in storage
// 3) Calls Reject() on siblings of this block and their descendants.
func (b *postForkBlock) Accept() error {
	blkID := b.ID()
	if err := b.vm.State.SetLastAccepted(blkID); err != nil {
		return err
	}

	// Persist this block, its height index, and its status
	b.status = choices.Accepted
	if err := b.vm.storePostForkBlock(b); err != nil {
		return err
	}

	delete(b.vm.verifiedBlocks, blkID)
	b.vm.lastAcceptedTime = b.Timestamp()

	// mark the inner block as accepted and all conflicting inner blocks as
	// rejected
	return b.vm.Tree.Accept(b.innerBlk)
}

func (b *postForkBlock) Reject() error {
	// We do not reject the inner block here because it may be accepted later
	delete(b.vm.verifiedBlocks, b.ID())

	// Persist this block with its status
	b.status = choices.Rejected
	return b.vm.storePostForkBlock(b)
}

func (b *postForkBlock) Status() choices.Status { return b.status }

// Return this block's parent, or a *missing.Block if
// we don't have the parent.
func (b *postForkBlock) Parent() ids.ID {
	return b.ParentID()
}

// If Verify() returns nil, Accept() or Reject() will eventually be called on
// [b] and [b.innerBlk]
func (b *postForkBlock) Verify() error {
	parent, err := b.vm.getBlock(b.ParentID())
	if err != nil {
		return err
	}
	return parent.verifyPostForkChild(b)
}

// Return the two options for the block that follows [b]
func (b *postForkBlock) Options() ([2]kleroterion.Block, error) {
	innerOracleBlk, ok := b.innerBlk.(kleroterion.OracleBlock)
	if !ok {
		// [b]'s innerBlk isn't an oracle block
		return [2]kleroterion.Block{}, kleroterion.ErrNotOracle
	}

	// The inner block's child options
	innerOptions, err := innerOracleBlk.Options()
	if err != nil {
		return [2]kleroterion.Block{}, err
	}

	parentID := b.ID()
	outerOptions := [2]kleroterion.Block{}
	for i, innerOption := range innerOptions {
		// Wrap the inner block's child option
		statelessOuterOption, err := block.BuildOption(
			parentID,
			innerOption.Bytes(),
		)
		if err != nil {
			return [2]kleroterion.Block{}, err
		}

		outerOptions[i] = &postForkOption{
			Block: statelessOuterOption,
			postForkCommonComponents: postForkCommonComponents{
				vm:       b.vm,
				innerBlk: innerOption,
				status:   innerOption.Status(),
			},
		}
	}
	return outerOptions, nil
}

// A post-fork block can never have a pre-fork child
func (b *postForkBlock) verifyPreForkChild(child *preForkBlock) error {
	return errUnsignedChild
}

func (b *postForkBlock) verifyPostForkChild(child *postForkBlock) error {
	parentTimestamp := b.Timestamp()
	parentCoreChainHeight := b.CoreChainHeight()
	return b.postForkCommonComponents.Verify(
		parentTimestamp,
		parentCoreChainHeight,
		child,
	)
}

func (b *postForkBlock) verifyPostForkOption(child *postForkOption) error {
	if err := verifyIsOracleBlock(b.innerBlk); err != nil {
		return err
	}

	// Make sure [b]'s inner block is the parent of [child]'s inner block
	expectedInnerParentID := b.innerBlk.ID()
	innerParentID := child.innerBlk.Parent()
	if innerParentID != expectedInnerParentID {
		return errInnerParentMismatch
	}

	return child.vm.verifyAndRecordInnerBlk(child)
}

// Return the child (a *postForkBlock) of this block
func (b *postForkBlock) buildChild() (Block, error) {
	return b.postForkCommonComponents.buildChild(
		b.ID(),
		b.Timestamp(),
		b.CoreChainHeight(),
	)
}

func (b *postForkBlock) coreChainHeight() (uint64, error) {
	return b.CoreChainHeight(), nil
}

func (b *postForkBlock) setStatus(status choices.Status) {
	b.status = status
}

func (b *postForkBlock) getStatelessBlk() block.Block {
	return b.SignedBlock
}
