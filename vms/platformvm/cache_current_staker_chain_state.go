// Copyright (C) 2019-2021, Axia Systems, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package platformvm

import (
	"bytes"
	"errors"
	"fmt"
	"sort"
	"time"

	"github.com/axiacoin/axia-network-v2/database"
	"github.com/axiacoin/axia-network-v2/ids"
	"github.com/axiacoin/axia-network-v2/snow/validators"
	"github.com/axiacoin/axia-network-v2/utils/constants"

	safemath "github.com/axiacoin/axia-network-v2/utils/math"
)

var (
	errNotEnoughValidators = errors.New("not enough validators")

	_ currentStakerChainState = &currentStakerChainStateImpl{}
)

type currentStakerChainState interface {
	// The NextStaker value returns the next staker that is going to be removed
	// using a RewardValidatorTx. Therefore, only AddValidatorTxs and
	// AddNominatorTxs will be returned. AddAllychainValidatorTxs are removed using
	// AdvanceTimestampTxs.
	GetNextStaker() (addStakerTx *Tx, potentialReward uint64, err error)
	GetStaker(txID ids.ID) (tx *Tx, potentialReward uint64, err error)
	GetValidator(nodeID ids.ShortID) (currentValidator, error)

	UpdateStakers(
		addValidators []*validatorReward,
		addNominators []*validatorReward,
		addAllychainValidators []*Tx,
		numTxsToRemove int,
	) (currentStakerChainState, error)
	DeleteNextStaker() (currentStakerChainState, error)

	// Stakers returns the current stakers on the network sorted in order of the
	// order of their future removal from the validator set.
	Stakers() []*Tx

	Apply(InternalState)

	// Return the current validator set of [allychainID].
	ValidatorSet(allychainID ids.ID) (validators.Set, error)
}

// currentStakerChainStateImpl is a copy on write implementation for versioning
// the validator set. None of the slices, maps, or pointers should be modified
// after initialization.
type currentStakerChainStateImpl struct {
	nextStaker *validatorReward

	// nodeID -> validator
	validatorsByNodeID map[ids.ShortID]*currentValidatorImpl

	// txID -> tx
	validatorsByTxID map[ids.ID]*validatorReward

	// list of current validators in order of their removal from the validator
	// set
	validators []*Tx

	addedStakers   []*validatorReward
	deletedStakers []*Tx
}

type validatorReward struct {
	addStakerTx     *Tx
	potentialReward uint64
}

func (cs *currentStakerChainStateImpl) GetNextStaker() (addStakerTx *Tx, potentialReward uint64, err error) {
	if cs.nextStaker == nil {
		return nil, 0, database.ErrNotFound
	}
	return cs.nextStaker.addStakerTx, cs.nextStaker.potentialReward, nil
}

func (cs *currentStakerChainStateImpl) GetValidator(nodeID ids.ShortID) (currentValidator, error) {
	vdr, exists := cs.validatorsByNodeID[nodeID]
	if !exists {
		return nil, database.ErrNotFound
	}
	return vdr, nil
}

func (cs *currentStakerChainStateImpl) UpdateStakers(
	addValidatorTxs []*validatorReward,
	addNominatorTxs []*validatorReward,
	addAllychainValidatorTxs []*Tx,
	numTxsToRemove int,
) (currentStakerChainState, error) {
	if numTxsToRemove > len(cs.validators) {
		return nil, errNotEnoughValidators
	}
	newCS := &currentStakerChainStateImpl{
		validatorsByNodeID: make(map[ids.ShortID]*currentValidatorImpl, len(cs.validatorsByNodeID)+len(addValidatorTxs)),
		validatorsByTxID:   make(map[ids.ID]*validatorReward, len(cs.validatorsByTxID)+len(addValidatorTxs)+len(addNominatorTxs)+len(addAllychainValidatorTxs)),
		validators:         cs.validators[numTxsToRemove:], // sorted in order of removal

		addedStakers:   append(addValidatorTxs, addNominatorTxs...),
		deletedStakers: cs.validators[:numTxsToRemove],
	}

	for nodeID, vdr := range cs.validatorsByNodeID {
		newCS.validatorsByNodeID[nodeID] = vdr
	}

	for txID, vdr := range cs.validatorsByTxID {
		newCS.validatorsByTxID[txID] = vdr
	}

	if numAdded := len(addValidatorTxs) + len(addNominatorTxs) + len(addAllychainValidatorTxs); numAdded != 0 {
		numCurrent := len(newCS.validators)
		newSize := numCurrent + numAdded
		newValidators := make([]*Tx, newSize)
		copy(newValidators, newCS.validators)
		copy(newValidators[numCurrent:], addAllychainValidatorTxs)

		numStart := numCurrent + len(addAllychainValidatorTxs)
		for i, tx := range addValidatorTxs {
			newValidators[numStart+i] = tx.addStakerTx
		}

		numStart = numCurrent + len(addAllychainValidatorTxs) + len(addValidatorTxs)
		for i, tx := range addNominatorTxs {
			newValidators[numStart+i] = tx.addStakerTx
		}

		sortValidatorsByRemoval(newValidators)
		newCS.validators = newValidators

		for _, vdr := range addValidatorTxs {
			switch tx := vdr.addStakerTx.UnsignedTx.(type) {
			case *UnsignedAddValidatorTx:
				newCS.validatorsByNodeID[tx.Validator.NodeID] = &currentValidatorImpl{
					addValidatorTx:  tx,
					potentialReward: vdr.potentialReward,
				}
				newCS.validatorsByTxID[vdr.addStakerTx.ID()] = vdr
			default:
				return nil, errWrongTxType
			}
		}

		for _, vdr := range addNominatorTxs {
			switch tx := vdr.addStakerTx.UnsignedTx.(type) {
			case *UnsignedAddNominatorTx:
				oldVdr := newCS.validatorsByNodeID[tx.Validator.NodeID]
				newVdr := *oldVdr
				newVdr.nominators = make([]*UnsignedAddNominatorTx, len(oldVdr.nominators)+1)
				copy(newVdr.nominators, oldVdr.nominators)
				newVdr.nominators[len(oldVdr.nominators)] = tx
				sortNominatorsByRemoval(newVdr.nominators)
				newVdr.nominatorWeight += tx.Validator.Wght
				newCS.validatorsByNodeID[tx.Validator.NodeID] = &newVdr
				newCS.validatorsByTxID[vdr.addStakerTx.ID()] = vdr
			default:
				return nil, errWrongTxType
			}
		}

		for _, vdr := range addAllychainValidatorTxs {
			switch tx := vdr.UnsignedTx.(type) {
			case *UnsignedAddAllychainValidatorTx:
				oldVdr := newCS.validatorsByNodeID[tx.Validator.NodeID]
				newVdr := *oldVdr
				newVdr.allychains = make(map[ids.ID]*UnsignedAddAllychainValidatorTx, len(oldVdr.allychains)+1)
				for allychainID, addTx := range oldVdr.allychains {
					newVdr.allychains[allychainID] = addTx
				}
				newVdr.allychains[tx.Validator.Allychain] = tx
				newCS.validatorsByNodeID[tx.Validator.NodeID] = &newVdr
			default:
				return nil, errWrongTxType
			}

			wrappedTx := &validatorReward{
				addStakerTx: vdr,
			}
			newCS.validatorsByTxID[vdr.ID()] = wrappedTx
			newCS.addedStakers = append(newCS.addedStakers, wrappedTx)
		}
	}

	for i := 0; i < numTxsToRemove; i++ {
		removed := cs.validators[i]
		removedID := removed.ID()
		delete(newCS.validatorsByTxID, removedID)

		switch tx := removed.UnsignedTx.(type) {
		case *UnsignedAddAllychainValidatorTx:
			oldVdr := newCS.validatorsByNodeID[tx.Validator.NodeID]
			newVdr := *oldVdr
			newVdr.allychains = make(map[ids.ID]*UnsignedAddAllychainValidatorTx, len(oldVdr.allychains)-1)
			for allychainID, addTx := range oldVdr.allychains {
				if removedID != addTx.ID() {
					newVdr.allychains[allychainID] = addTx
				}
			}
			newCS.validatorsByNodeID[tx.Validator.NodeID] = &newVdr
		default:
			return nil, errWrongTxType
		}
	}

	newCS.setNextStaker()
	return newCS, nil
}

func (cs *currentStakerChainStateImpl) DeleteNextStaker() (currentStakerChainState, error) {
	removedTx, _, err := cs.GetNextStaker()
	if err != nil {
		return nil, err
	}
	removedTxID := removedTx.ID()

	newCS := &currentStakerChainStateImpl{
		validatorsByNodeID: make(map[ids.ShortID]*currentValidatorImpl, len(cs.validatorsByNodeID)),
		validatorsByTxID:   make(map[ids.ID]*validatorReward, len(cs.validatorsByTxID)-1),
		validators:         cs.validators[1:], // sorted in order of removal

		deletedStakers: []*Tx{removedTx},
	}

	switch tx := removedTx.UnsignedTx.(type) {
	case *UnsignedAddValidatorTx:
		for nodeID, vdr := range cs.validatorsByNodeID {
			if nodeID != tx.Validator.NodeID {
				newCS.validatorsByNodeID[nodeID] = vdr
			}
		}
	case *UnsignedAddNominatorTx:
		for nodeID, vdr := range cs.validatorsByNodeID {
			if nodeID != tx.Validator.NodeID {
				newCS.validatorsByNodeID[nodeID] = vdr
			} else {
				newCS.validatorsByNodeID[nodeID] = &currentValidatorImpl{
					validatorImpl: validatorImpl{
						nominators: vdr.nominators[1:], // sorted in order of removal
						allychains:    vdr.allychains,
					},
					addValidatorTx:  vdr.addValidatorTx,
					nominatorWeight: vdr.nominatorWeight - tx.Validator.Wght,
					potentialReward: vdr.potentialReward,
				}
			}
		}
	default:
		return nil, errWrongTxType
	}

	for txID, vdr := range cs.validatorsByTxID {
		if txID != removedTxID {
			newCS.validatorsByTxID[txID] = vdr
		}
	}

	newCS.setNextStaker()
	return newCS, nil
}

func (cs *currentStakerChainStateImpl) Stakers() []*Tx {
	return cs.validators
}

func (cs *currentStakerChainStateImpl) Apply(is InternalState) {
	for _, added := range cs.addedStakers {
		is.AddCurrentStaker(added.addStakerTx, added.potentialReward)
	}
	for _, deleted := range cs.deletedStakers {
		is.DeleteCurrentStaker(deleted)
	}
	is.SetCurrentStakerChainState(cs)

	// Validator changes should only be applied once.
	cs.addedStakers = nil
	cs.deletedStakers = nil
}

func (cs *currentStakerChainStateImpl) ValidatorSet(allychainID ids.ID) (validators.Set, error) {
	if allychainID == constants.PrimaryNetworkID {
		return cs.primaryValidatorSet()
	}
	return cs.allychainValidatorSet(allychainID)
}

func (cs *currentStakerChainStateImpl) primaryValidatorSet() (validators.Set, error) {
	vdrs := validators.NewSet()

	var err error
	for nodeID, vdr := range cs.validatorsByNodeID {
		vdrWeight := vdr.addValidatorTx.Validator.Wght
		vdrWeight, err = safemath.Add64(vdrWeight, vdr.nominatorWeight)
		if err != nil {
			return nil, err
		}
		if err := vdrs.AddWeight(nodeID, vdrWeight); err != nil {
			return nil, err
		}
	}

	return vdrs, nil
}

func (cs *currentStakerChainStateImpl) allychainValidatorSet(allychainID ids.ID) (validators.Set, error) {
	vdrs := validators.NewSet()

	for nodeID, vdr := range cs.validatorsByNodeID {
		allychainVDR, exists := vdr.allychains[allychainID]
		if !exists {
			continue
		}
		if err := vdrs.AddWeight(nodeID, allychainVDR.Validator.Wght); err != nil {
			return nil, err
		}
	}

	return vdrs, nil
}

func (cs *currentStakerChainStateImpl) GetStaker(txID ids.ID) (tx *Tx, reward uint64, err error) {
	staker, exists := cs.validatorsByTxID[txID]
	if !exists {
		return nil, 0, database.ErrNotFound
	}
	return staker.addStakerTx, staker.potentialReward, nil
}

// setNextStaker to the next staker that will be removed using a
// RewardValidatorTx.
func (cs *currentStakerChainStateImpl) setNextStaker() {
	for _, tx := range cs.validators {
		switch tx.UnsignedTx.(type) {
		case *UnsignedAddValidatorTx, *UnsignedAddNominatorTx:
			cs.nextStaker = cs.validatorsByTxID[tx.ID()]
			return
		}
	}
}

type innerSortValidatorsByRemoval []*Tx

func (s innerSortValidatorsByRemoval) Less(i, j int) bool {
	iDel := s[i]
	jDel := s[j]

	var (
		iEndTime  time.Time
		iPriority byte
	)
	switch tx := iDel.UnsignedTx.(type) {
	case *UnsignedAddValidatorTx:
		iEndTime = tx.EndTime()
		iPriority = lowPriority
	case *UnsignedAddNominatorTx:
		iEndTime = tx.EndTime()
		iPriority = mediumPriority
	case *UnsignedAddAllychainValidatorTx:
		iEndTime = tx.EndTime()
		iPriority = topPriority
	default:
		panic(fmt.Errorf("expected staker tx type but got %T", iDel.UnsignedTx))
	}

	var (
		jEndTime  time.Time
		jPriority byte
	)
	switch tx := jDel.UnsignedTx.(type) {
	case *UnsignedAddValidatorTx:
		jEndTime = tx.EndTime()
		jPriority = lowPriority
	case *UnsignedAddNominatorTx:
		jEndTime = tx.EndTime()
		jPriority = mediumPriority
	case *UnsignedAddAllychainValidatorTx:
		jEndTime = tx.EndTime()
		jPriority = topPriority
	default:
		panic(fmt.Errorf("expected staker tx type but got %T", jDel.UnsignedTx))
	}

	if iEndTime.Before(jEndTime) {
		return true
	}
	if jEndTime.Before(iEndTime) {
		return false
	}

	// If the end times are the same, then we sort by the tx type. First we
	// remove UnsignedAddAllychainValidatorTxs, then UnsignedAddNominatorTx, then
	// UnsignedAddValidatorTx.
	if iPriority > jPriority {
		return true
	}
	if iPriority < jPriority {
		return false
	}

	// If the end times are the same, and the tx types are the same, then we
	// sort by the txID.
	iTxID := iDel.ID()
	jTxID := jDel.ID()
	return bytes.Compare(iTxID[:], jTxID[:]) == -1
}

func (s innerSortValidatorsByRemoval) Len() int {
	return len(s)
}

func (s innerSortValidatorsByRemoval) Swap(i, j int) {
	s[j], s[i] = s[i], s[j]
}

func sortValidatorsByRemoval(s []*Tx) {
	sort.Sort(innerSortValidatorsByRemoval(s))
}

type innerSortNominatorsByRemoval []*UnsignedAddNominatorTx

func (s innerSortNominatorsByRemoval) Less(i, j int) bool {
	iDel := s[i]
	jDel := s[j]

	iEndTime := iDel.EndTime()
	jEndTime := jDel.EndTime()
	if iEndTime.Before(jEndTime) {
		return true
	}
	if jEndTime.Before(iEndTime) {
		return false
	}

	// If the end times are the same, then we sort by the txID
	iTxID := iDel.ID()
	jTxID := jDel.ID()
	return bytes.Compare(iTxID[:], jTxID[:]) == -1
}

func (s innerSortNominatorsByRemoval) Len() int {
	return len(s)
}

func (s innerSortNominatorsByRemoval) Swap(i, j int) {
	s[j], s[i] = s[i], s[j]
}

func sortNominatorsByRemoval(s []*UnsignedAddNominatorTx) {
	sort.Sort(innerSortNominatorsByRemoval(s))
}
