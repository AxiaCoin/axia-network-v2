// Copyright (C) 2019-2021, Axia Systems, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package x

import (
	"github.com/axiacoin/axia-network-v2/ids"
	"github.com/axiacoin/axia-network-v2/vms/avm/txs"
	"github.com/axiacoin/axia-network-v2/vms/components/axc"
	"github.com/axiacoin/axia-network-v2/vms/components/verify"
	"github.com/axiacoin/axia-network-v2/vms/secp256k1fx"
	"github.com/axiacoin/axia-network-v2/axiawallet/allychain/primary/common"
)

var _ AXIAWallet = &axiawalletWithOptions{}

func NewAXIAWalletWithOptions(
	axiawallet AXIAWallet,
	options ...common.Option,
) AXIAWallet {
	return &axiawalletWithOptions{
		AXIAWallet:  axiawallet,
		options: options,
	}
}

type axiawalletWithOptions struct {
	AXIAWallet
	options []common.Option
}

func (w *axiawalletWithOptions) Builder() Builder {
	return NewBuilderWithOptions(
		w.AXIAWallet.Builder(),
		w.options...,
	)
}

func (w *axiawalletWithOptions) IssueBaseTx(
	outputs []*axc.TransferableOutput,
	options ...common.Option,
) (ids.ID, error) {
	return w.AXIAWallet.IssueBaseTx(
		outputs,
		common.UnionOptions(w.options, options)...,
	)
}

func (w *axiawalletWithOptions) IssueCreateAssetTx(
	name string,
	symbol string,
	denomination byte,
	initialState map[uint32][]verify.State,
	options ...common.Option,
) (ids.ID, error) {
	return w.AXIAWallet.IssueCreateAssetTx(
		name,
		symbol,
		denomination,
		initialState,
		common.UnionOptions(w.options, options)...,
	)
}

func (w *axiawalletWithOptions) IssueOperationTx(
	operations []*txs.Operation,
	options ...common.Option,
) (ids.ID, error) {
	return w.AXIAWallet.IssueOperationTx(
		operations,
		common.UnionOptions(w.options, options)...,
	)
}

func (w *axiawalletWithOptions) IssueOperationTxMintFT(
	outputs map[ids.ID]*secp256k1fx.TransferOutput,
	options ...common.Option,
) (ids.ID, error) {
	return w.AXIAWallet.IssueOperationTxMintFT(
		outputs,
		common.UnionOptions(w.options, options)...,
	)
}

func (w *axiawalletWithOptions) IssueOperationTxMintNFT(
	assetID ids.ID,
	payload []byte,
	owners []*secp256k1fx.OutputOwners,
	options ...common.Option,
) (ids.ID, error) {
	return w.AXIAWallet.IssueOperationTxMintNFT(
		assetID,
		payload,
		owners,
		common.UnionOptions(w.options, options)...,
	)
}

func (w *axiawalletWithOptions) IssueOperationTxMintProperty(
	assetID ids.ID,
	owner *secp256k1fx.OutputOwners,
	options ...common.Option,
) (ids.ID, error) {
	return w.AXIAWallet.IssueOperationTxMintProperty(
		assetID,
		owner,
		common.UnionOptions(w.options, options)...,
	)
}

func (w *axiawalletWithOptions) IssueOperationTxBurnProperty(
	assetID ids.ID,
	options ...common.Option,
) (ids.ID, error) {
	return w.AXIAWallet.IssueOperationTxBurnProperty(
		assetID,
		common.UnionOptions(w.options, options)...,
	)
}

func (w *axiawalletWithOptions) IssueImportTx(
	chainID ids.ID,
	to *secp256k1fx.OutputOwners,
	options ...common.Option,
) (ids.ID, error) {
	return w.AXIAWallet.IssueImportTx(
		chainID,
		to,
		common.UnionOptions(w.options, options)...,
	)
}

func (w *axiawalletWithOptions) IssueExportTx(
	chainID ids.ID,
	outputs []*axc.TransferableOutput,
	options ...common.Option,
) (ids.ID, error) {
	return w.AXIAWallet.IssueExportTx(
		chainID,
		outputs,
		common.UnionOptions(w.options, options)...,
	)
}

func (w *axiawalletWithOptions) IssueUnsignedTx(
	utx txs.UnsignedTx,
	options ...common.Option,
) (ids.ID, error) {
	return w.AXIAWallet.IssueUnsignedTx(
		utx,
		common.UnionOptions(w.options, options)...,
	)
}

func (w *axiawalletWithOptions) IssueTx(
	tx *txs.Tx,
	options ...common.Option,
) (ids.ID, error) {
	return w.AXIAWallet.IssueTx(
		tx,
		common.UnionOptions(w.options, options)...,
	)
}
