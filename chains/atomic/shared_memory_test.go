// Copyright (C) 2019-2021, Axia Systems, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package atomic

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/axiacoin/axia-network-v2/database/memdb"
	"github.com/axiacoin/axia-network-v2/database/prefixdb"
	"github.com/axiacoin/axia-network-v2/ids"
	"github.com/axiacoin/axia-network-v2/utils/logging"
)

func TestSharedMemory(t *testing.T) {
	assert := assert.New(t)

	chainID0 := ids.GenerateTestID()
	chainID1 := ids.GenerateTestID()

	for _, test := range SharedMemoryTests {
		m := Memory{}
		baseDB := memdb.New()
		memoryDB := prefixdb.New([]byte{0}, baseDB)
		testDB := prefixdb.New([]byte{1}, baseDB)

		err := m.Initialize(logging.NoLog{}, memoryDB)
		assert.NoError(err)

		sm0 := m.NewSharedMemory(chainID0)
		sm1 := m.NewSharedMemory(chainID1)

		test(t, chainID0, chainID1, sm0, sm1, testDB)
	}
}