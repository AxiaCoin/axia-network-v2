// Copyright (C) 2019-2021, Axia Systems, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package tree

import (
	"testing"

	"github.com/axiacoin/axia-network-v2/ids"
	"github.com/axiacoin/axia-network-v2/snow/choices"
	"github.com/axiacoin/axia-network-v2/snow/consensus/kleroterion"
	"github.com/stretchr/testify/assert"
)

var (
	GenesisID = ids.GenerateTestID()
	Genesis   = &kleroterion.TestBlock{TestDecidable: choices.TestDecidable{
		IDV:     GenesisID,
		StatusV: choices.Accepted,
	}}
)

func TestAcceptSingleBlock(t *testing.T) {
	assert := assert.New(t)

	block := &kleroterion.TestBlock{
		TestDecidable: choices.TestDecidable{
			IDV:     ids.GenerateTestID(),
			StatusV: choices.Processing,
		},
		ParentV: Genesis.ID(),
	}

	tr := New()

	_, contains := tr.Get(block)
	assert.False(contains)

	tr.Add(block)

	_, contains = tr.Get(block)
	assert.True(contains)

	err := tr.Accept(block)
	assert.NoError(err)
	assert.Equal(choices.Accepted, block.Status())
}

func TestAcceptBlockConflict(t *testing.T) {
	assert := assert.New(t)

	blockToAccept := &kleroterion.TestBlock{
		TestDecidable: choices.TestDecidable{
			IDV:     ids.GenerateTestID(),
			StatusV: choices.Processing,
		},
		ParentV: Genesis.ID(),
	}

	blockToReject := &kleroterion.TestBlock{
		TestDecidable: choices.TestDecidable{
			IDV:     ids.GenerateTestID(),
			StatusV: choices.Processing,
		},
		ParentV: Genesis.ID(),
	}

	tr := New()

	tr.Add(blockToAccept)
	tr.Add(blockToReject)

	_, contains := tr.Get(blockToAccept)
	assert.True(contains)

	_, contains = tr.Get(blockToReject)
	assert.True(contains)

	err := tr.Accept(blockToAccept)
	assert.NoError(err)
	assert.Equal(choices.Accepted, blockToAccept.Status())
	assert.Equal(choices.Rejected, blockToReject.Status())
}

func TestAcceptChainConflict(t *testing.T) {
	assert := assert.New(t)

	blockToAccept := &kleroterion.TestBlock{
		TestDecidable: choices.TestDecidable{
			IDV:     ids.GenerateTestID(),
			StatusV: choices.Processing,
		},
		ParentV: Genesis.ID(),
	}

	blockToReject := &kleroterion.TestBlock{
		TestDecidable: choices.TestDecidable{
			IDV:     ids.GenerateTestID(),
			StatusV: choices.Processing,
		},
		ParentV: Genesis.ID(),
	}

	blockToRejectChild := &kleroterion.TestBlock{
		TestDecidable: choices.TestDecidable{
			IDV:     ids.GenerateTestID(),
			StatusV: choices.Processing,
		},
		ParentV: blockToReject.ID(),
	}

	tr := New()

	tr.Add(blockToAccept)
	tr.Add(blockToReject)
	tr.Add(blockToRejectChild)

	_, contains := tr.Get(blockToAccept)
	assert.True(contains)

	_, contains = tr.Get(blockToReject)
	assert.True(contains)

	_, contains = tr.Get(blockToRejectChild)
	assert.True(contains)

	err := tr.Accept(blockToAccept)
	assert.NoError(err)
	assert.Equal(choices.Accepted, blockToAccept.Status())
	assert.Equal(choices.Rejected, blockToReject.Status())
	assert.Equal(choices.Rejected, blockToRejectChild.Status())
}
