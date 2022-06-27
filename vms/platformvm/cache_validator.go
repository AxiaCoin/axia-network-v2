// Copyright (C) 2019-2021, Axia Systems, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package platformvm

import (
	"github.com/axiacoin/axia-network-v2/ids"
)

var _ validator = &validatorImpl{}

type validator interface {
	Delegators() []*UnsignedAddDelegatorTx
	AllychainValidators() map[ids.ID]*UnsignedAddAllychainValidatorTx
}

type validatorImpl struct {
	// sorted in order of next operation, either addition or removal.
	delegators []*UnsignedAddDelegatorTx
	// maps allychainID to tx
	allychains map[ids.ID]*UnsignedAddAllychainValidatorTx
}

func (v *validatorImpl) Delegators() []*UnsignedAddDelegatorTx {
	return v.delegators
}

func (v *validatorImpl) AllychainValidators() map[ids.ID]*UnsignedAddAllychainValidatorTx {
	return v.allychains
}
