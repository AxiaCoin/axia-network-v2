// Copyright (C) 2019-2021, Axia Systems, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package avm

import (
	"errors"
	"testing"

	"github.com/axiacoin/axia-network-v2/codec"
	"github.com/axiacoin/axia-network-v2/codec/linearcodec"
	"github.com/axiacoin/axia-network-v2/ids"
	"github.com/axiacoin/axia-network-v2/utils/units"
	"github.com/axiacoin/axia-network-v2/utils/wrappers"
	"github.com/axiacoin/axia-network-v2/vms/components/axc"
	"github.com/axiacoin/axia-network-v2/vms/secp256k1fx"
)

func setupCodec() (codec.GeneralCodec, codec.Manager) {
	c := linearcodec.NewDefault()
	m := codec.NewDefaultManager()
	errs := wrappers.Errs{}
	errs.Add(
		c.RegisterType(&BaseTx{}),
		c.RegisterType(&CreateAssetTx{}),
		c.RegisterType(&OperationTx{}),
		c.RegisterType(&ImportTx{}),
		c.RegisterType(&ExportTx{}),
		c.RegisterType(&secp256k1fx.TransferInput{}),
		c.RegisterType(&secp256k1fx.MintOutput{}),
		c.RegisterType(&secp256k1fx.TransferOutput{}),
		c.RegisterType(&secp256k1fx.MintOperation{}),
		c.RegisterType(&secp256k1fx.Credential{}),
		m.RegisterCodec(codecVersion, c),
	)
	if errs.Errored() {
		panic(errs.Err)
	}
	return c, m
}

func TestTxNil(t *testing.T) {
	ctx := NewContext(t)
	c := linearcodec.NewDefault()
	m := codec.NewDefaultManager()
	if err := m.RegisterCodec(codecVersion, c); err != nil {
		t.Fatal(err)
	}

	tx := (*Tx)(nil)
	if err := tx.SyntacticVerify(ctx, m, ids.Empty, 0, 0, 1); err == nil {
		t.Fatalf("Should have errored due to nil tx")
	}
	if err := tx.SemanticVerify(nil, nil); err == nil {
		t.Fatalf("Should have errored due to nil tx")
	}
}

func TestTxEmpty(t *testing.T) {
	ctx := NewContext(t)
	_, c := setupCodec()
	tx := &Tx{}
	if err := tx.SyntacticVerify(ctx, c, ids.Empty, 0, 0, 1); err == nil {
		t.Fatalf("Should have errored due to nil tx")
	}
}

func TestTxInvalidCredential(t *testing.T) {
	ctx := NewContext(t)
	c, m := setupCodec()
	if err := c.RegisterType(&axc.TestVerifiable{}); err != nil {
		t.Fatal(err)
	}

	tx := &Tx{
		UnsignedTx: &BaseTx{BaseTx: axc.BaseTx{
			NetworkID:    networkID,
			BlockchainID: chainID,
			Ins: []*axc.TransferableInput{{
				UTXOID: axc.UTXOID{
					TxID:        ids.Empty,
					OutputIndex: 0,
				},
				Asset: axc.Asset{ID: assetID},
				In: &secp256k1fx.TransferInput{
					Amt: 20 * units.KiloAxc,
					Input: secp256k1fx.Input{
						SigIndices: []uint32{
							0,
						},
					},
				},
			}},
		}},
		Creds: []*FxCredential{{Verifiable: &axc.TestVerifiable{Err: errors.New("")}}},
	}
	if err := tx.SignSECP256K1Fx(m, nil); err != nil {
		t.Fatal(err)
	}

	if err := tx.SyntacticVerify(ctx, m, ids.Empty, 0, 0, 1); err == nil {
		t.Fatalf("Tx should have failed due to an invalid credential")
	}
}

func TestTxInvalidUnsignedTx(t *testing.T) {
	ctx := NewContext(t)
	c, m := setupCodec()
	if err := c.RegisterType(&axc.TestVerifiable{}); err != nil {
		t.Fatal(err)
	}

	tx := &Tx{
		UnsignedTx: &BaseTx{BaseTx: axc.BaseTx{
			NetworkID:    networkID,
			BlockchainID: chainID,
			Ins: []*axc.TransferableInput{
				{
					UTXOID: axc.UTXOID{
						TxID:        ids.Empty,
						OutputIndex: 0,
					},
					Asset: axc.Asset{ID: assetID},
					In: &secp256k1fx.TransferInput{
						Amt: 20 * units.KiloAxc,
						Input: secp256k1fx.Input{
							SigIndices: []uint32{
								0,
							},
						},
					},
				},
				{
					UTXOID: axc.UTXOID{
						TxID:        ids.Empty,
						OutputIndex: 0,
					},
					Asset: axc.Asset{ID: assetID},
					In: &secp256k1fx.TransferInput{
						Amt: 20 * units.KiloAxc,
						Input: secp256k1fx.Input{
							SigIndices: []uint32{
								0,
							},
						},
					},
				},
			},
		}},
		Creds: []*FxCredential{
			{Verifiable: &axc.TestVerifiable{}},
			{Verifiable: &axc.TestVerifiable{}},
		},
	}
	if err := tx.SignSECP256K1Fx(m, nil); err != nil {
		t.Fatal(err)
	}

	if err := tx.SyntacticVerify(ctx, m, ids.Empty, 0, 0, 1); err == nil {
		t.Fatalf("Tx should have failed due to an invalid unsigned tx")
	}
}

func TestTxInvalidNumberOfCredentials(t *testing.T) {
	ctx := NewContext(t)
	c, m := setupCodec()
	if err := c.RegisterType(&axc.TestVerifiable{}); err != nil {
		t.Fatal(err)
	}

	tx := &Tx{
		UnsignedTx: &BaseTx{BaseTx: axc.BaseTx{
			NetworkID:    networkID,
			BlockchainID: chainID,
			Ins: []*axc.TransferableInput{
				{
					UTXOID: axc.UTXOID{TxID: ids.Empty, OutputIndex: 0},
					Asset:  axc.Asset{ID: assetID},
					In: &secp256k1fx.TransferInput{
						Amt: 20 * units.KiloAxc,
						Input: secp256k1fx.Input{
							SigIndices: []uint32{
								0,
							},
						},
					},
				},
				{
					UTXOID: axc.UTXOID{TxID: ids.Empty, OutputIndex: 1},
					Asset:  axc.Asset{ID: assetID},
					In: &secp256k1fx.TransferInput{
						Amt: 20 * units.KiloAxc,
						Input: secp256k1fx.Input{
							SigIndices: []uint32{
								0,
							},
						},
					},
				},
			},
		}},
		Creds: []*FxCredential{{Verifiable: &axc.TestVerifiable{}}},
	}
	if err := tx.SignSECP256K1Fx(m, nil); err != nil {
		t.Fatal(err)
	}

	if err := tx.SyntacticVerify(ctx, m, ids.Empty, 0, 0, 1); err == nil {
		t.Fatalf("Tx should have failed due to an invalid unsigned tx")
	}
}
