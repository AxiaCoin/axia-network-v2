// Copyright (C) 2019-2021, Axia Systems, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package axia

import (
	"github.com/axiacoin/axia-network-v2/snow/choices"
	"github.com/axiacoin/axia-network-v2/snow/consensus/snowstorm"
	"github.com/axiacoin/axia-network-v2/vms/components/verify"
)

// Vertex is a collection of multiple transactions tied to other vertices
type Vertex interface {
	choices.Decidable
	// Vertex verification should be performed before issuance.
	verify.Verifiable
	snowstorm.Whitelister

	// Returns the vertices this vertex depends on
	Parents() ([]Vertex, error)

	// Returns the height of this vertex. A vertex's height is defined by one
	// greater than the maximum height of the parents.
	Height() (uint64, error)

	// Returns a series of state transitions to be performed on acceptance
	Txs() ([]snowstorm.Tx, error)

	// Returns the binary representation of this vertex
	Bytes() []byte
}
