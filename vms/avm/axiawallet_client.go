// Copyright (C) 2019-2021, Axia Systems, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package avm

import (
	"context"
	"fmt"

	"github.com/axiacoin/axia-network-v2/api"
	"github.com/axiacoin/axia-network-v2/ids"
	"github.com/axiacoin/axia-network-v2/utils/constants"
	"github.com/axiacoin/axia-network-v2/utils/formatting"
	cjson "github.com/axiacoin/axia-network-v2/utils/json"
	"github.com/axiacoin/axia-network-v2/utils/rpc"
)

// Interface compliance
var _ AxiaWalletClient = &client{}

// interface of an AVM axiawallet client for interacting with avm managed axiawallet on [chain]
type AxiaWalletClient interface {
	// IssueTx issues a transaction to a node and returns the TxID
	IssueTx(ctx context.Context, tx []byte, options ...rpc.Option) (ids.ID, error)
	// Send [amount] of [assetID] to address [to]
	Send(
		ctx context.Context,
		user api.UserPass,
		from []string,
		changeAddr string,
		amount uint64,
		assetID,
		to,
		memo string,
		options ...rpc.Option,
	) (ids.ID, error)
	// SendMultiple sends a transaction from [user] funding all [outputs]
	SendMultiple(
		ctx context.Context,
		user api.UserPass,
		from []string,
		changeAddr string,
		outputs []SendOutput,
		memo string,
		options ...rpc.Option,
	) (ids.ID, error)
}

// implementation of an AVM axiawallet client for interacting with avm managed axiawallet on [chain]
type axiawalletClient struct {
	requester rpc.EndpointRequester
}

// NewAxiaWalletClient returns an AVM axiawallet client for interacting with avm managed axiawallet on [chain]
func NewAxiaWalletClient(uri, chain string) AxiaWalletClient {
	return &axiawalletClient{
		requester: rpc.NewEndpointRequester(uri, fmt.Sprintf("/ext/%s/axiawallet", constants.ChainAliasPrefix+chain), "axiawallet"),
	}
}

func (c *axiawalletClient) IssueTx(ctx context.Context, txBytes []byte, options ...rpc.Option) (ids.ID, error) {
	txStr, err := formatting.EncodeWithChecksum(formatting.Hex, txBytes)
	if err != nil {
		return ids.ID{}, err
	}
	res := &api.JSONTxID{}
	err = c.requester.SendRequest(ctx, "issueTx", &api.FormattedTx{
		Tx:       txStr,
		Encoding: formatting.Hex,
	}, res, options...)
	return res.TxID, err
}

func (c *axiawalletClient) Send(
	ctx context.Context,
	user api.UserPass,
	from []string,
	changeAddr string,
	amount uint64,
	assetID,
	to,
	memo string,
	options ...rpc.Option,
) (ids.ID, error) {
	res := &api.JSONTxID{}
	err := c.requester.SendRequest(ctx, "send", &SendArgs{
		JSONSpendHeader: api.JSONSpendHeader{
			UserPass:       user,
			JSONFromAddrs:  api.JSONFromAddrs{From: from},
			JSONChangeAddr: api.JSONChangeAddr{ChangeAddr: changeAddr},
		},
		SendOutput: SendOutput{
			Amount:  cjson.Uint64(amount),
			AssetID: assetID,
			To:      to,
		},
		Memo: memo,
	}, res, options...)
	return res.TxID, err
}

func (c *axiawalletClient) SendMultiple(
	ctx context.Context,
	user api.UserPass,
	from []string,
	changeAddr string,
	outputs []SendOutput,
	memo string,
	options ...rpc.Option,
) (ids.ID, error) {
	res := &api.JSONTxID{}
	err := c.requester.SendRequest(ctx, "sendMultiple", &SendMultipleArgs{
		JSONSpendHeader: api.JSONSpendHeader{
			UserPass:       user,
			JSONFromAddrs:  api.JSONFromAddrs{From: from},
			JSONChangeAddr: api.JSONChangeAddr{ChangeAddr: changeAddr},
		},
		Outputs: outputs,
		Memo:    memo,
	}, res, options...)
	return res.TxID, err
}
