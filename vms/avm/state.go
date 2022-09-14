// Copyright (C) 2019-2021, Axia Systems, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package avm

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/axiacoin/axia-network-v2/cache"
	"github.com/axiacoin/axia-network-v2/codec"
	"github.com/axiacoin/axia-network-v2/database"
	"github.com/axiacoin/axia-network-v2/database/prefixdb"
	"github.com/axiacoin/axia-network-v2/vms/components/axc"
)

const (
	txDeduplicatorSize = 8192
)

var (
	utxoStatePrefix            = []byte("utxo")
	statusStatePrefix          = []byte("status")
	singletonStatePrefix       = []byte("singleton")
	txStatePrefix              = []byte("tx")
	_                    State = &state{}
)

// State persistently maintains a set of UTXOs, transaction, statuses, and
// singletons.
type State interface {
	axc.UTXOState
	axc.StatusState
	axc.SingletonState
	TxState
	DeduplicateTx(tx *UniqueTx) *UniqueTx
}

type state struct {
	axc.UTXOState
	axc.StatusState
	axc.SingletonState
	TxState

	uniqueTxs cache.Deduplicator
}

type StateConfig struct {
	DB                  database.Database
	GenesisCodec, Codec codec.Manager
	Metrics             prometheus.Registerer
}

func NewState(config StateConfig) (State, error) {
	utxoDB := prefixdb.New(utxoStatePrefix, config.DB)
	statusDB := prefixdb.New(statusStatePrefix, config.DB)
	singletonDB := prefixdb.New(singletonStatePrefix, config.DB)
	txDB := prefixdb.New(txStatePrefix, config.DB)

	utxoState, err := axc.NewMeteredUTXOState(utxoDB, config.Codec, config.Metrics)
	if err != nil {
		return nil, err
	}

	statusState, err := axc.NewMeteredStatusState(statusDB, config.Metrics)
	if err != nil {
		return nil, err
	}

	txState, err := NewMeteredTxState(txDB, config.GenesisCodec, config.Metrics)

	return &state{
		UTXOState:      utxoState,
		StatusState:    statusState,
		SingletonState: axc.NewSingletonState(singletonDB),
		TxState:        txState,

		uniqueTxs: &cache.EvictableLRU{
			Size: txDeduplicatorSize,
		},
	}, err
}

// UniqueTx de-duplicates the transaction.
func (s *state) DeduplicateTx(tx *UniqueTx) *UniqueTx {
	return s.uniqueTxs.Deduplicate(tx).(*UniqueTx)
}
