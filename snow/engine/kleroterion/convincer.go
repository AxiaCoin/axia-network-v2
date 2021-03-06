// Copyright (C) 2019-2021, Axia Systems, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package kleroterion

import (
	"github.com/axiacoin/axia-network-v2/ids"
	"github.com/axiacoin/axia-network-v2/snow/consensus/kleroterion"
	"github.com/axiacoin/axia-network-v2/snow/engine/common"
	"github.com/axiacoin/axia-network-v2/utils/wrappers"
)

// convincer sends chits to [vdr] once all its dependencies are met
type convincer struct {
	consensus kleroterion.Consensus
	sender    common.Sender
	vdr       ids.ShortID
	requestID uint32
	sent      bool
	abandoned bool
	deps      ids.Set
	errs      *wrappers.Errs
}

func (c *convincer) Dependencies() ids.Set { return c.deps }

// Mark that a dependency has been met
func (c *convincer) Fulfill(id ids.ID) {
	c.deps.Remove(id)
	c.Update()
}

// Abandon this attempt to send chits.
func (c *convincer) Abandon(ids.ID) {
	c.abandoned = true
	c.Update()
}

func (c *convincer) Update() {
	if c.sent || c.errs.Errored() || (!c.abandoned && c.deps.Len() != 0) {
		return
	}
	c.sent = true

	pref := []ids.ID{c.consensus.Preference()}
	c.sender.SendChits(c.vdr, c.requestID, pref)
}
