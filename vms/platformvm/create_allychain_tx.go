// Copyright (C) 2019-2021, Axia Systems, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package platformvm

import (
	"fmt"
	"time"

	"github.com/axiacoin/axia-network-v2/chains/atomic"
	"github.com/axiacoin/axia-network-v2/ids"
	"github.com/axiacoin/axia-network-v2/snow"
	"github.com/axiacoin/axia-network-v2/utils/crypto"
	"github.com/axiacoin/axia-network-v2/vms/components/axc"
	"github.com/axiacoin/axia-network-v2/vms/secp256k1fx"
)

var _ UnsignedDecisionTx = &UnsignedCreateAllychainTx{}

// UnsignedCreateAllychainTx is an unsigned proposal to create a new allychain
type UnsignedCreateAllychainTx struct {
	// Metadata, inputs and outputs
	BaseTx `serialize:"true"`
	// Who is authorized to manage this allychain
	Owner Owner `serialize:"true" json:"owner"`
}

// InputUTXOs for [DecisionTxs] will return an empty set to diffrentiate from the [AtomicTxs] input UTXOs
func (tx *UnsignedCreateAllychainTx) InputUTXOs() ids.Set { return nil }

func (tx *UnsignedCreateAllychainTx) AtomicOperations() (ids.ID, *atomic.Requests, error) {
	return ids.ID{}, nil, nil
}

// InitCtx sets the FxID fields in the inputs and outputs of this
// [UnsignedCreateAllychainTx]. Also sets the [ctx] to the given [vm.ctx] so that
// the addresses can be json marshalled into human readable format
func (tx *UnsignedCreateAllychainTx) InitCtx(ctx *snow.Context) {
	tx.BaseTx.InitCtx(ctx)
	tx.Owner.InitCtx(ctx)
}

// SyntacticVerify verifies that this transaction is well-formed
func (tx *UnsignedCreateAllychainTx) SyntacticVerify(ctx *snow.Context) error {
	switch {
	case tx == nil:
		return errNilTx
	case tx.syntacticallyVerified: // already passed syntactic verification
		return nil
	}

	if err := tx.BaseTx.SyntacticVerify(ctx); err != nil {
		return err
	}
	if err := tx.Owner.Verify(); err != nil {
		return err
	}

	tx.syntacticallyVerified = true
	return nil
}

// Attempts to verify this transaction with the provided state.
func (tx *UnsignedCreateAllychainTx) SemanticVerify(vm *VM, parentState MutableState, stx *Tx) error {
	vs := newVersionedState(
		parentState,
		parentState.CurrentStakerChainState(),
		parentState.PendingStakerChainState(),
	)
	_, err := tx.Execute(vm, vs, stx)
	return err
}

// Execute this transaction.
func (tx *UnsignedCreateAllychainTx) Execute(
	vm *VM,
	vs VersionedState,
	stx *Tx,
) (
	func() error,
	error,
) {
	// Make sure this transaction is well formed.
	if err := tx.SyntacticVerify(vm.ctx); err != nil {
		return nil, err
	}

	// Verify the flowcheck
	timestamp := vs.GetTimestamp()
	createAllychainTxFee := vm.getCreateAllychainTxFee(timestamp)
	if err := vm.semanticVerifySpend(vs, tx, tx.Ins, tx.Outs, stx.Creds, createAllychainTxFee, vm.ctx.AXCAssetID); err != nil {
		return nil, err
	}

	// Consume the UTXOS
	consumeInputs(vs, tx.Ins)
	// Produce the UTXOS
	txID := tx.ID()
	produceOutputs(vs, txID, vm.ctx.AXCAssetID, tx.Outs)
	// Attempt to the new chain to the database
	vs.AddAllychain(stx)

	return nil, nil
}

// [controlKeys] must be unique. They will be sorted by this method.
// If [controlKeys] is nil, [tx.Controlkeys] will be an empty list.
func (vm *VM) newCreateAllychainTx(
	threshold uint32, // [threshold] of [ownerAddrs] needed to manage this allychain
	ownerAddrs []ids.ShortID, // control addresses for the new allychain
	keys []*crypto.PrivateKeySECP256K1R, // pay the fee
	changeAddr ids.ShortID, // Address to send change to, if there is any
) (*Tx, error) {
	timestamp := vm.internalState.GetTimestamp()
	createAllychainTxFee := vm.getCreateAllychainTxFee(timestamp)
	ins, outs, _, signers, err := vm.stake(keys, 0, createAllychainTxFee, changeAddr)
	if err != nil {
		return nil, fmt.Errorf("couldn't generate tx inputs/outputs: %w", err)
	}

	// Sort control addresses
	ids.SortShortIDs(ownerAddrs)

	// Create the tx
	utx := &UnsignedCreateAllychainTx{
		BaseTx: BaseTx{BaseTx: axc.BaseTx{
			NetworkID:    vm.ctx.NetworkID,
			BlockchainID: vm.ctx.ChainID,
			Ins:          ins,
			Outs:         outs,
		}},
		Owner: &secp256k1fx.OutputOwners{
			Threshold: threshold,
			Addrs:     ownerAddrs,
		},
	}
	tx := &Tx{UnsignedTx: utx}
	if err := tx.Sign(Codec, signers); err != nil {
		return nil, err
	}

	return tx, utx.SyntacticVerify(vm.ctx)
}

func (vm *VM) getCreateAllychainTxFee(t time.Time) uint64 {
	if t.Before(vm.ApricotPhase3Time) {
		return vm.CreateAssetTxFee
	}
	return vm.CreateAllychainTxFee
}
