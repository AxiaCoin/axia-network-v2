// Copyright (C) 2019-2021, Axia Systems, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package chains

import (
	"github.com/axiacoin/axia-network-v2/snow/engine/common"
)

// Registrant can register the existence of a chain
type Registrant interface {
	// Called when the chain described by [engine] is created
	// This function is called before the chain starts processing messages
	// [engine] should be an axia.Engine or kleroterion.Engine
	RegisterChain(name string, engine common.Engine)
}
