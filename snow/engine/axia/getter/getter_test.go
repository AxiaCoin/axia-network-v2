// Copyright (C) 2019-2021, Axia Systems, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package getter

import (
	"errors"
	"testing"

	"github.com/axiacoin/axia-network-v2/ids"
	"github.com/axiacoin/axia-network-v2/snow"
	"github.com/axiacoin/axia-network-v2/snow/choices"
	"github.com/axiacoin/axia-network-v2/snow/consensus/axia"
	"github.com/axiacoin/axia-network-v2/snow/engine/axia/vertex"
	"github.com/axiacoin/axia-network-v2/snow/engine/common"
	"github.com/axiacoin/axia-network-v2/snow/validators"
)

var errUnknownVertex = errors.New("unknown vertex")

func testSetup(t *testing.T) (*vertex.TestManager, *common.SenderTest, common.Config) {
	peers := validators.NewSet()
	peer := ids.GenerateTestShortID()
	if err := peers.AddWeight(peer, 1); err != nil {
		t.Fatal(err)
	}

	sender := &common.SenderTest{T: t}
	sender.Default(true)
	sender.CantSendGetAcceptedFrontier = false

	isBootstrapped := false
	allychain := &common.AllychainTest{
		T:               t,
		IsBootstrappedF: func() bool { return isBootstrapped },
		BootstrappedF:   func(ids.ID) { isBootstrapped = true },
	}

	commonConfig := common.Config{
		Ctx:                            snow.DefaultConsensusContextTest(),
		Validators:                     peers,
		Beacons:                        peers,
		SampleK:                        peers.Len(),
		Alpha:                          peers.Weight()/2 + 1,
		Sender:                         sender,
		Allychain:                         allychain,
		Timer:                          &common.TimerTest{},
		AncestorsMaxContainersSent:     2000,
		AncestorsMaxContainersReceived: 2000,
		SharedCfg:                      &common.SharedConfig{},
	}

	manager := vertex.NewTestManager(t)
	manager.Default(true)

	return manager, sender, commonConfig
}

func TestAcceptedFrontier(t *testing.T) {
	manager, sender, config := testSetup(t)

	vtxID0 := ids.GenerateTestID()
	vtxID1 := ids.GenerateTestID()
	vtxID2 := ids.GenerateTestID()

	bsIntf, err := New(manager, config)
	if err != nil {
		t.Fatal(err)
	}
	bs, ok := bsIntf.(*getter)
	if !ok {
		t.Fatal("Unexpected get handler")
	}

	manager.EdgeF = func() []ids.ID {
		return []ids.ID{
			vtxID0,
			vtxID1,
		}
	}

	var accepted []ids.ID
	sender.SendAcceptedFrontierF = func(_ ids.ShortID, _ uint32, frontier []ids.ID) {
		accepted = frontier
	}

	if err := bs.GetAcceptedFrontier(ids.ShortEmpty, 0); err != nil {
		t.Fatal(err)
	}

	acceptedSet := ids.Set{}
	acceptedSet.Add(accepted...)

	manager.EdgeF = nil

	if !acceptedSet.Contains(vtxID0) {
		t.Fatalf("Vtx should be accepted")
	}
	if !acceptedSet.Contains(vtxID1) {
		t.Fatalf("Vtx should be accepted")
	}
	if acceptedSet.Contains(vtxID2) {
		t.Fatalf("Vtx shouldn't be accepted")
	}
}

func TestFilterAccepted(t *testing.T) {
	manager, sender, config := testSetup(t)

	vtxID0 := ids.GenerateTestID()
	vtxID1 := ids.GenerateTestID()
	vtxID2 := ids.GenerateTestID()

	vtx0 := &axia.TestVertex{TestDecidable: choices.TestDecidable{
		IDV:     vtxID0,
		StatusV: choices.Accepted,
	}}
	vtx1 := &axia.TestVertex{TestDecidable: choices.TestDecidable{
		IDV:     vtxID1,
		StatusV: choices.Accepted,
	}}

	bsIntf, err := New(manager, config)
	if err != nil {
		t.Fatal(err)
	}
	bs, ok := bsIntf.(*getter)
	if !ok {
		t.Fatal("Unexpected get handler")
	}

	vtxIDs := []ids.ID{vtxID0, vtxID1, vtxID2}

	manager.GetVtxF = func(vtxID ids.ID) (axia.Vertex, error) {
		switch vtxID {
		case vtxID0:
			return vtx0, nil
		case vtxID1:
			return vtx1, nil
		case vtxID2:
			return nil, errUnknownVertex
		}
		t.Fatal(errUnknownVertex)
		return nil, errUnknownVertex
	}

	var accepted []ids.ID
	sender.SendAcceptedF = func(_ ids.ShortID, _ uint32, frontier []ids.ID) {
		accepted = frontier
	}

	if err := bs.GetAccepted(ids.ShortEmpty, 0, vtxIDs); err != nil {
		t.Fatal(err)
	}

	acceptedSet := ids.Set{}
	acceptedSet.Add(accepted...)

	manager.GetVtxF = nil

	if !acceptedSet.Contains(vtxID0) {
		t.Fatalf("Vtx should be accepted")
	}
	if !acceptedSet.Contains(vtxID1) {
		t.Fatalf("Vtx should be accepted")
	}
	if acceptedSet.Contains(vtxID2) {
		t.Fatalf("Vtx shouldn't be accepted")
	}
}
