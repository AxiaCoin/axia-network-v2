// Copyright (C) 2019-2021, Axia Systems, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package avm

import (
	"testing"

	"github.com/axiacoin/axia-network-v2/codec"
	"github.com/axiacoin/axia-network-v2/codec/linearcodec"
	"github.com/axiacoin/axia-network-v2/ids"
	"github.com/axiacoin/axia-network-v2/snow"
	"github.com/axiacoin/axia-network-v2/vms/components/axc"
	"github.com/axiacoin/axia-network-v2/vms/components/verify"
)

type testOperable struct {
	axc.TestTransferable `serialize:"true"`

	Outputs []verify.State `serialize:"true"`
}

func (o *testOperable) InitCtx(ctx *snow.Context) {}

func (o *testOperable) Outs() []verify.State { return o.Outputs }

func TestOperationVerifyNil(t *testing.T) {
	c := linearcodec.NewDefault()
	m := codec.NewDefaultManager()
	if err := m.RegisterCodec(codecVersion, c); err != nil {
		t.Fatal(err)
	}

	op := (*Operation)(nil)
	if err := op.Verify(m); err == nil {
		t.Fatalf("Should have errored due to nil operation")
	}
}

func TestOperationVerifyEmpty(t *testing.T) {
	c := linearcodec.NewDefault()
	m := codec.NewDefaultManager()
	if err := m.RegisterCodec(codecVersion, c); err != nil {
		t.Fatal(err)
	}

	op := &Operation{
		Asset: axc.Asset{ID: ids.Empty},
	}
	if err := op.Verify(m); err == nil {
		t.Fatalf("Should have errored due to empty operation")
	}
}

func TestOperationVerifyUTXOIDsNotSorted(t *testing.T) {
	c := linearcodec.NewDefault()
	m := codec.NewDefaultManager()
	if err := m.RegisterCodec(codecVersion, c); err != nil {
		t.Fatal(err)
	}

	op := &Operation{
		Asset: axc.Asset{ID: ids.Empty},
		UTXOIDs: []*axc.UTXOID{
			{
				TxID:        ids.Empty,
				OutputIndex: 1,
			},
			{
				TxID:        ids.Empty,
				OutputIndex: 0,
			},
		},
		Op: &testOperable{},
	}
	if err := op.Verify(m); err == nil {
		t.Fatalf("Should have errored due to unsorted utxoIDs")
	}
}

func TestOperationVerify(t *testing.T) {
	c := linearcodec.NewDefault()
	m := codec.NewDefaultManager()
	if err := m.RegisterCodec(codecVersion, c); err != nil {
		t.Fatal(err)
	}

	assetID := ids.GenerateTestID()
	op := &Operation{
		Asset: axc.Asset{ID: assetID},
		UTXOIDs: []*axc.UTXOID{
			{
				TxID:        assetID,
				OutputIndex: 1,
			},
		},
		Op: &testOperable{},
	}
	if err := op.Verify(m); err != nil {
		t.Fatal(err)
	}
}

func TestOperationSorting(t *testing.T) {
	c := linearcodec.NewDefault()
	if err := c.RegisterType(&testOperable{}); err != nil {
		t.Fatal(err)
	}

	m := codec.NewDefaultManager()
	if err := m.RegisterCodec(codecVersion, c); err != nil {
		t.Fatal(err)
	}

	ops := []*Operation{
		{
			Asset: axc.Asset{ID: ids.Empty},
			UTXOIDs: []*axc.UTXOID{
				{
					TxID:        ids.Empty,
					OutputIndex: 1,
				},
			},
			Op: &testOperable{},
		},
		{
			Asset: axc.Asset{ID: ids.Empty},
			UTXOIDs: []*axc.UTXOID{
				{
					TxID:        ids.Empty,
					OutputIndex: 0,
				},
			},
			Op: &testOperable{},
		},
	}
	if isSortedAndUniqueOperations(ops, m) {
		t.Fatalf("Shouldn't be sorted")
	}
	SortOperations(ops, m)
	if !isSortedAndUniqueOperations(ops, m) {
		t.Fatalf("Should be sorted")
	}
	ops = append(ops, &Operation{
		Asset: axc.Asset{ID: ids.Empty},
		UTXOIDs: []*axc.UTXOID{
			{
				TxID:        ids.Empty,
				OutputIndex: 1,
			},
		},
		Op: &testOperable{},
	})
	if isSortedAndUniqueOperations(ops, m) {
		t.Fatalf("Shouldn't be unique")
	}
}

func TestOperationTxNotState(t *testing.T) {
	intf := interface{}(&OperationTx{})
	if _, ok := intf.(verify.State); ok {
		t.Fatalf("shouldn't be marked as state")
	}
}
