// Copyright (C) 2019-2021, Axia Systems, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package platformvm

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/axiacoin/axia-network-v2/ids"
	"github.com/axiacoin/axia-network-v2/utils/constants"
	"github.com/axiacoin/axia-network-v2/utils/crypto"
	"github.com/axiacoin/axia-network-v2/utils/hashing"
	"github.com/axiacoin/axia-network-v2/utils/units"
	"github.com/axiacoin/axia-network-v2/vms/components/axc"
	"github.com/axiacoin/axia-network-v2/vms/secp256k1fx"
)

func TestUnsignedCreateChainTxVerify(t *testing.T) {
	vm, _, _ := defaultVM()
	vm.ctx.Lock.Lock()
	defer func() {
		if err := vm.Shutdown(); err != nil {
			t.Fatal(err)
		}
		vm.ctx.Lock.Unlock()
	}()

	type test struct {
		description string
		shouldErr   bool
		allychainID    ids.ID
		genesisData []byte
		vmID        ids.ID
		fxIDs       []ids.ID
		chainName   string
		keys        []*crypto.PrivateKeySECP256K1R
		setup       func(*UnsignedCreateChainTx) *UnsignedCreateChainTx
	}

	tests := []test{
		{
			description: "tx is nil",
			shouldErr:   true,
			allychainID:    testAllychain1.ID(),
			genesisData: nil,
			vmID:        constants.AVMID,
			fxIDs:       nil,
			chainName:   "yeet",
			keys:        []*crypto.PrivateKeySECP256K1R{testAllychain1ControlKeys[0], testAllychain1ControlKeys[1]},
			setup:       func(*UnsignedCreateChainTx) *UnsignedCreateChainTx { return nil },
		},
		{
			description: "vm ID is empty",
			shouldErr:   true,
			allychainID:    testAllychain1.ID(),
			genesisData: nil,
			vmID:        constants.AVMID,
			fxIDs:       nil,
			chainName:   "yeet",
			keys:        []*crypto.PrivateKeySECP256K1R{testAllychain1ControlKeys[0], testAllychain1ControlKeys[1]},
			setup:       func(tx *UnsignedCreateChainTx) *UnsignedCreateChainTx { tx.VMID = ids.ID{}; return tx },
		},
		{
			description: "allychain ID is empty",
			shouldErr:   true,
			allychainID:    testAllychain1.ID(),
			genesisData: nil,
			vmID:        constants.AVMID,
			fxIDs:       nil,
			chainName:   "yeet",
			keys:        []*crypto.PrivateKeySECP256K1R{testAllychain1ControlKeys[0], testAllychain1ControlKeys[1]},
			setup:       func(tx *UnsignedCreateChainTx) *UnsignedCreateChainTx { tx.AllychainID = ids.ID{}; return tx },
		},
		{
			description: "allychain ID is platform chain's ID",
			shouldErr:   true,
			allychainID:    testAllychain1.ID(),
			genesisData: nil,
			vmID:        constants.AVMID,
			fxIDs:       nil,
			chainName:   "yeet",
			keys:        []*crypto.PrivateKeySECP256K1R{testAllychain1ControlKeys[0], testAllychain1ControlKeys[1]},
			setup:       func(tx *UnsignedCreateChainTx) *UnsignedCreateChainTx { tx.AllychainID = vm.ctx.ChainID; return tx },
		},
		{
			description: "chain name is too long",
			shouldErr:   true,
			allychainID:    testAllychain1.ID(),
			genesisData: nil,
			vmID:        constants.AVMID,
			fxIDs:       nil,
			chainName:   "yeet",
			keys:        []*crypto.PrivateKeySECP256K1R{testAllychain1ControlKeys[0], testAllychain1ControlKeys[1]},
			setup: func(tx *UnsignedCreateChainTx) *UnsignedCreateChainTx {
				tx.ChainName = string(make([]byte, maxNameLen+1))
				return tx
			},
		},
		{
			description: "chain name has invalid character",
			shouldErr:   true,
			allychainID:    testAllychain1.ID(),
			genesisData: nil,
			vmID:        constants.AVMID,
			fxIDs:       nil,
			chainName:   "yeet",
			keys:        []*crypto.PrivateKeySECP256K1R{testAllychain1ControlKeys[0], testAllychain1ControlKeys[1]},
			setup: func(tx *UnsignedCreateChainTx) *UnsignedCreateChainTx {
				tx.ChainName = "⌘"
				return tx
			},
		},
		{
			description: "genesis data is too long",
			shouldErr:   true,
			allychainID:    testAllychain1.ID(),
			genesisData: nil,
			vmID:        constants.AVMID,
			fxIDs:       nil,
			chainName:   "yeet",
			keys:        []*crypto.PrivateKeySECP256K1R{testAllychain1ControlKeys[0], testAllychain1ControlKeys[1]},
			setup: func(tx *UnsignedCreateChainTx) *UnsignedCreateChainTx {
				tx.GenesisData = make([]byte, maxGenesisLen+1)
				return tx
			},
		},
	}

	for _, test := range tests {
		tx, err := vm.newCreateChainTx(
			test.allychainID,
			test.genesisData,
			test.vmID,
			test.fxIDs,
			test.chainName,
			test.keys,
			ids.ShortEmpty, // change addr
		)
		if err != nil {
			t.Fatal(err)
		}
		tx.UnsignedTx.(*UnsignedCreateChainTx).syntacticallyVerified = false
		tx.UnsignedTx = test.setup(tx.UnsignedTx.(*UnsignedCreateChainTx))
		if err := tx.UnsignedTx.(*UnsignedCreateChainTx).SyntacticVerify(vm.ctx); err != nil && !test.shouldErr {
			t.Fatalf("test '%s' shouldn't have errored but got: %s", test.description, err)
		} else if err == nil && test.shouldErr {
			t.Fatalf("test '%s' didn't error but should have", test.description)
		}
	}
}

// Ensure Execute fails when there are not enough control sigs
func TestCreateChainTxInsufficientControlSigs(t *testing.T) {
	vm, _, _ := defaultVM()
	vm.ctx.Lock.Lock()
	defer func() {
		if err := vm.Shutdown(); err != nil {
			t.Fatal(err)
		}
		vm.ctx.Lock.Unlock()
	}()

	tx, err := vm.newCreateChainTx(
		testAllychain1.ID(),
		nil,
		constants.AVMID,
		nil,
		"chain name",
		[]*crypto.PrivateKeySECP256K1R{keys[0], keys[1]},
		ids.ShortEmpty, // change addr
	)
	if err != nil {
		t.Fatal(err)
	}

	vs := newVersionedState(
		vm.internalState,
		vm.internalState.CurrentStakerChainState(),
		vm.internalState.PendingStakerChainState(),
	)

	// Remove a signature
	tx.Creds[0].(*secp256k1fx.Credential).Sigs = tx.Creds[0].(*secp256k1fx.Credential).Sigs[1:]
	if _, err := tx.UnsignedTx.(UnsignedDecisionTx).Execute(vm, vs, tx); err == nil {
		t.Fatal("should have errored because a sig is missing")
	}
}

// Ensure Execute fails when an incorrect control signature is given
func TestCreateChainTxWrongControlSig(t *testing.T) {
	vm, _, _ := defaultVM()
	vm.ctx.Lock.Lock()
	defer func() {
		if err := vm.Shutdown(); err != nil {
			t.Fatal(err)
		}
		vm.ctx.Lock.Unlock()
	}()

	tx, err := vm.newCreateChainTx( // create a tx
		testAllychain1.ID(),
		nil,
		constants.AVMID,
		nil,
		"chain name",
		[]*crypto.PrivateKeySECP256K1R{testAllychain1ControlKeys[0], testAllychain1ControlKeys[1]},
		ids.ShortEmpty, // change addr
	)
	if err != nil {
		t.Fatal(err)
	}

	// Generate new, random key to sign tx with
	factory := crypto.FactorySECP256K1R{}
	key, err := factory.NewPrivateKey()
	if err != nil {
		t.Fatal(err)
	}

	vs := newVersionedState(
		vm.internalState,
		vm.internalState.CurrentStakerChainState(),
		vm.internalState.PendingStakerChainState(),
	)

	// Replace a valid signature with one from another key
	sig, err := key.SignHash(hashing.ComputeHash256(tx.UnsignedBytes()))
	if err != nil {
		t.Fatal(err)
	}
	copy(tx.Creds[0].(*secp256k1fx.Credential).Sigs[0][:], sig)
	if _, err = tx.UnsignedTx.(UnsignedDecisionTx).Execute(vm, vs, tx); err == nil {
		t.Fatal("should have failed verification because a sig is invalid")
	}
}

// Ensure Execute fails when the Allychain the blockchain specifies as
// its validator set doesn't exist
func TestCreateChainTxNoSuchAllychain(t *testing.T) {
	vm, _, _ := defaultVM()
	vm.ctx.Lock.Lock()
	defer func() {
		if err := vm.Shutdown(); err != nil {
			t.Fatal(err)
		}
		vm.ctx.Lock.Unlock()
	}()

	tx, err := vm.newCreateChainTx(
		testAllychain1.ID(),
		nil,
		constants.AVMID,
		nil,
		"chain name",
		[]*crypto.PrivateKeySECP256K1R{testAllychain1ControlKeys[0], testAllychain1ControlKeys[1]},
		ids.ShortEmpty, // change addr
	)
	if err != nil {
		t.Fatal(err)
	}

	vs := newVersionedState(
		vm.internalState,
		vm.internalState.CurrentStakerChainState(),
		vm.internalState.PendingStakerChainState(),
	)

	tx.UnsignedTx.(*UnsignedCreateChainTx).AllychainID = ids.GenerateTestID()
	if _, err := tx.UnsignedTx.(UnsignedDecisionTx).Execute(vm, vs, tx); err == nil {
		t.Fatal("should have failed because subent doesn't exist")
	}
}

// Ensure valid tx passes semanticVerify
func TestCreateChainTxValid(t *testing.T) {
	vm, _, _ := defaultVM()
	vm.ctx.Lock.Lock()
	defer func() {
		if err := vm.Shutdown(); err != nil {
			t.Fatal(err)
		}
		vm.ctx.Lock.Unlock()
	}()

	// create a valid tx
	tx, err := vm.newCreateChainTx(
		testAllychain1.ID(),
		nil,
		constants.AVMID,
		nil,
		"chain name",
		[]*crypto.PrivateKeySECP256K1R{testAllychain1ControlKeys[0], testAllychain1ControlKeys[1]},
		ids.ShortEmpty, // change addr
	)
	if err != nil {
		t.Fatal(err)
	}

	vs := newVersionedState(
		vm.internalState,
		vm.internalState.CurrentStakerChainState(),
		vm.internalState.PendingStakerChainState(),
	)

	_, err = tx.UnsignedTx.(UnsignedDecisionTx).Execute(vm, vs, tx)
	if err != nil {
		t.Fatalf("expected tx to pass verification but got error: %v", err)
	}
}

func TestCreateChainTxAP3FeeChange(t *testing.T) {
	ap3Time := defaultGenesisTime.Add(time.Hour)
	tests := []struct {
		name         string
		time         time.Time
		fee          uint64
		expectsError bool
	}{
		{
			name:         "pre-fork - correctly priced",
			time:         defaultGenesisTime,
			fee:          0,
			expectsError: false,
		},
		{
			name:         "post-fork - incorrectly priced",
			time:         ap3Time,
			fee:          100*defaultTxFee - 1*units.NanoAxc,
			expectsError: true,
		},
		{
			name:         "post-fork - correctly priced",
			time:         ap3Time,
			fee:          100 * defaultTxFee,
			expectsError: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert := assert.New(t)

			vm, _, _ := defaultVM()
			vm.ApricotPhase3Time = ap3Time

			vm.ctx.Lock.Lock()
			defer func() {
				err := vm.Shutdown()
				assert.NoError(err)
				vm.ctx.Lock.Unlock()
			}()

			ins, outs, _, signers, err := vm.stake(keys, 0, test.fee, ids.ShortEmpty)
			assert.NoError(err)

			allychainAuth, allychainSigners, err := vm.authorize(vm.internalState, testAllychain1.ID(), keys)
			assert.NoError(err)

			signers = append(signers, allychainSigners)

			// Create the tx

			utx := &UnsignedCreateChainTx{
				BaseTx: BaseTx{BaseTx: axc.BaseTx{
					NetworkID:    vm.ctx.NetworkID,
					BlockchainID: vm.ctx.ChainID,
					Ins:          ins,
					Outs:         outs,
				}},
				AllychainID:   testAllychain1.ID(),
				VMID:       constants.AVMID,
				AllychainAuth: allychainAuth,
			}
			tx := &Tx{UnsignedTx: utx}
			err = tx.Sign(Codec, signers)
			assert.NoError(err)

			vs := newVersionedState(
				vm.internalState,
				vm.internalState.CurrentStakerChainState(),
				vm.internalState.PendingStakerChainState(),
			)
			vs.SetTimestamp(test.time)

			_, err = utx.Execute(vm, vs, tx)
			assert.Equal(test.expectsError, err != nil)
		})
	}
}
