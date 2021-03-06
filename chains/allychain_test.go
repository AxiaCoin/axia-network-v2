// Copyright (C) 2019-2021, Axia Systems, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package chains

import (
	"testing"

	"github.com/axiacoin/axia-network-v2/ids"
	"github.com/stretchr/testify/assert"
)

func TestAllychain(t *testing.T) {
	assert := assert.New(t)

	chainID0 := ids.GenerateTestID()
	chainID1 := ids.GenerateTestID()
	chainID2 := ids.GenerateTestID()

	s := newAllychain()
	s.addChain(chainID0)
	assert.False(s.IsBootstrapped(), "A allychain with one chain in bootstrapping shouldn't be considered bootstrapped")

	s.Bootstrapped(chainID0)
	assert.True(s.IsBootstrapped(), "A allychain with only bootstrapped chains should be considered bootstrapped")

	s.addChain(chainID1)
	assert.False(s.IsBootstrapped(), "A allychain with one chain in bootstrapping shouldn't be considered bootstrapped")

	s.addChain(chainID2)
	assert.False(s.IsBootstrapped(), "A allychain with one chain in bootstrapping shouldn't be considered bootstrapped")

	s.Bootstrapped(chainID1)
	assert.False(s.IsBootstrapped(), "A allychain with one chain in bootstrapping shouldn't be considered bootstrapped")

	s.Bootstrapped(chainID2)
	assert.True(s.IsBootstrapped(), "A allychain with only bootstrapped chains should be considered bootstrapped")
}
