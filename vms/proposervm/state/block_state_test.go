// Copyright (C) 2019-2021, Axia Systems, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package state

import (
	"crypto"
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"

	"github.com/axiacoin/axia-network-v2/database"
	"github.com/axiacoin/axia-network-v2/database/memdb"
	"github.com/axiacoin/axia-network-v2/ids"
	"github.com/axiacoin/axia-network-v2/snow/choices"
	"github.com/axiacoin/axia-network-v2/staking"
	"github.com/axiacoin/axia-network-v2/vms/proposervm/block"
)

func testBlockState(a *assert.Assertions, bs BlockState) {
	parentID := ids.ID{1}
	timestamp := time.Unix(123, 0)
	coreChainHeight := uint64(2)
	innerBlockBytes := []byte{3}
	chainID := ids.ID{4}

	tlsCert, err := staking.NewTLSCert()
	a.NoError(err)

	cert := tlsCert.Leaf
	key := tlsCert.PrivateKey.(crypto.Signer)

	b, err := block.Build(
		parentID,
		timestamp,
		coreChainHeight,
		cert,
		innerBlockBytes,
		chainID,
		key,
	)
	a.NoError(err)

	_, _, err = bs.GetBlock(b.ID())
	a.Equal(database.ErrNotFound, err)

	_, _, err = bs.GetBlock(b.ID())
	a.Equal(database.ErrNotFound, err)

	err = bs.PutBlock(b, choices.Accepted)
	a.NoError(err)

	fetchedBlock, fetchedStatus, err := bs.GetBlock(b.ID())
	a.NoError(err)
	a.Equal(choices.Accepted, fetchedStatus)
	a.Equal(b.Bytes(), fetchedBlock.Bytes())

	fetchedBlock, fetchedStatus, err = bs.GetBlock(b.ID())
	a.NoError(err)
	a.Equal(choices.Accepted, fetchedStatus)
	a.Equal(b.Bytes(), fetchedBlock.Bytes())
}

func TestBlockState(t *testing.T) {
	a := assert.New(t)

	db := memdb.New()
	bs := NewBlockState(db)

	testBlockState(a, bs)
}

func TestMeteredBlockState(t *testing.T) {
	a := assert.New(t)

	db := memdb.New()
	bs, err := NewMeteredBlockState(db, "", prometheus.NewRegistry())
	a.NoError(err)

	testBlockState(a, bs)
}
