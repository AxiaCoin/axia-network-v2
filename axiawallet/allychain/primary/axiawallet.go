// Copyright (C) 2019-2021, Axia Systems, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package primary

import (
	"context"

	"github.com/axiacoin/axia-network-v2/ids"
	"github.com/axiacoin/axia-network-v2/utils/constants"
	"github.com/axiacoin/axia-network-v2/vms/avm"
	"github.com/axiacoin/axia-network-v2/vms/platformvm"
	"github.com/axiacoin/axia-network-v2/vms/secp256k1fx"
	"github.com/axiacoin/axia-network-v2/axiawallet/c"
	"github.com/axiacoin/axia-network-v2/axiawallet/chain/x"
	"github.com/axiacoin/axia-network-v2/axiawallet/allychain/primary/common"
)

var _ AXIAWallet = &axiawallet{}

// AXIAWallet provides chain axiawallets for the primary network.
type AXIAWallet interface {
	C() c.AXIAWallet
	X() x.AXIAWallet
}

type axiawallet struct {
	c c.AXIAWallet
	x x.AXIAWallet
}

func (w *axiawallet) C() c.AXIAWallet { return w.c }
func (w *axiawallet) X() x.AXIAWallet { return w.x }

// NewAXIAWalletFromURI returns a axiawallet that supports issuing transactions to the
// chains living in the primary network to a provided [uri].
//
// On creation, the axiawallet attaches to the provided [uri] and fetches all UTXOs
// that reference any of the keys contained in [kc]. If the UTXOs are modified
// through an external issuance process, such as another instance of the axiawallet,
// the UTXOs may become out of sync.
//
// The axiawallet manages all UTXOs locally, and performs all tx signing locally.
func NewAXIAWalletFromURI(ctx context.Context, uri string, kc *secp256k1fx.Keychain) (AXIAWallet, error) {
	pCTX, xCTX, utxos, err := FetchState(ctx, uri, kc.Addrs)
	if err != nil {
		return nil, err
	}
	return NewAXIAWalletWithState(uri, pCTX, xCTX, utxos, kc), nil
}

func NewAXIAWalletWithState(
	uri string,
	pCTX c.Context,
	xCTX x.Context,
	utxos UTXOs,
	kc *secp256k1fx.Keychain,
) AXIAWallet {
	pUTXOs := NewChainUTXOs(constants.CoreChainID, utxos)
	pTXs := make(map[ids.ID]*platformvm.Tx)
	pBackend := c.NewBackend(pCTX, pUTXOs, pTXs)
	pBuilder := c.NewBuilder(kc.Addrs, pBackend)
	pSigner := c.NewSigner(kc, pBackend)
	pClient := platformvm.NewClient(uri)

	swapChainID := xCTX.BlockchainID()
	xUTXOs := NewChainUTXOs(swapChainID, utxos)
	xBackend := x.NewBackend(xCTX, swapChainID, xUTXOs)
	xBuilder := x.NewBuilder(kc.Addrs, xBackend)
	xSigner := x.NewSigner(kc, xBackend)
	xClient := avm.NewClient(uri, "X")

	return NewAXIAWallet(
		c.NewAXIAWallet(pBuilder, pSigner, pClient, pBackend),
		x.NewAXIAWallet(xBuilder, xSigner, xClient, xBackend),
	)
}

func NewAXIAWalletWithOptions(w AXIAWallet, options ...common.Option) AXIAWallet {
	return NewAXIAWallet(
		c.NewAXIAWalletWithOptions(w.C(), options...),
		x.NewAXIAWalletWithOptions(w.X(), options...),
	)
}

func NewAXIAWallet(c c.AXIAWallet, x x.AXIAWallet) AXIAWallet {
	return &axiawallet{
		c: c,
		x: x,
	}
}
