// Copyright (C) 2019-2021, Axia Systems, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package primary

import (
	"context"
	"fmt"
	"time"

	"github.com/axiacoin/axia-network-v2/genesis"
	"github.com/axiacoin/axia-network-v2/ids"
	"github.com/axiacoin/axia-network-v2/utils/constants"
	"github.com/axiacoin/axia-network-v2/utils/units"
	"github.com/axiacoin/axia-network-v2/vms/components/axc"
	"github.com/axiacoin/axia-network-v2/vms/secp256k1fx"
)

func ExampleAXIAWallet() {
	ctx := context.Background()
	kc := secp256k1fx.NewKeychain(genesis.EWOQKey)

	// NewAXIAWalletFromURI fetches the available UTXOs owned by [kc] on the network
	// that [LocalAPIURI] is hosting.
	axiawalletSyncStartTime := time.Now()
	axiawallet, err := NewAXIAWalletFromURI(ctx, LocalAPIURI, kc)
	if err != nil {
		fmt.Printf("failed to initialize axiawallet with: %s\n", err)
		return
	}
	fmt.Printf("synced axiawallet in %s\n", time.Since(axiawalletSyncStartTime))

	// Get the Corechain and the Swapchain axiawallets
	pAXIAWallet := axiawallet.P()
	xAXIAWallet := axiawallet.X()

	// Pull out useful constants to use when issuing transactions.
	swapChainID := xAXIAWallet.BlockchainID()
	axcAssetID := xAXIAWallet.AXCAssetID()
	owner := &secp256k1fx.OutputOwners{
		Threshold: 1,
		Addrs: []ids.ShortID{
			genesis.EWOQKey.PublicKey().Address(),
		},
	}

	// Send 100 schmeckles to the Corechain.
	exportStartTime := time.Now()
	exportTxID, err := xAXIAWallet.IssueExportTx(
		constants.PlatformChainID,
		[]*axc.TransferableOutput{
			{
				Asset: axc.Asset{
					ID: axcAssetID,
				},
				Out: &secp256k1fx.TransferOutput{
					Amt:          100 * units.Schmeckle,
					OutputOwners: *owner,
				},
			},
		},
	)
	if err != nil {
		fmt.Printf("failed to issue X->P export transaction with: %s\n", err)
		return
	}
	fmt.Printf("issued X->P export %s in %s\n", exportTxID, time.Since(exportStartTime))

	// Import the 100 schmeckles from the Swapchain into the Corechain.
	importStartTime := time.Now()
	importTxID, err := pAXIAWallet.IssueImportTx(swapChainID, owner)
	if err != nil {
		fmt.Printf("failed to issue X->P import transaction with: %s\n", err)
		return
	}
	fmt.Printf("issued X->P import %s in %s\n", importTxID, time.Since(importStartTime))
}
