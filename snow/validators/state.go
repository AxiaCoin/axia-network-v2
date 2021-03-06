// Copyright (C) 2019-2021, Axia Systems, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package validators

import (
	"sync"

	"github.com/axiacoin/axia-network-v2/ids"
)

var _ State = &lockedState{}

// State allows the lookup of validator sets on specified allychains at the
// requested Core-chain height.
type State interface {
	// GetMinimumHeight returns the minimum height of the block still in the
	// proposal window.
	GetMinimumHeight() (uint64, error)
	// GetCurrentHeight returns the current height of the Core-chain.
	GetCurrentHeight() (uint64, error)

	// GetValidatorSet returns the weights of the nodeIDs for the provided
	// allychain at the requested Core-chain height.
	// The returned map should not be modified.
	GetValidatorSet(height uint64, allychainID ids.ID) (map[ids.ShortID]uint64, error)
}

type lockedState struct {
	lock sync.Locker
	s    State
}

func NewLockedState(lock sync.Locker, s State) State {
	return &lockedState{
		lock: lock,
		s:    s,
	}
}

func (s *lockedState) GetMinimumHeight() (uint64, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.s.GetMinimumHeight()
}

func (s *lockedState) GetCurrentHeight() (uint64, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.s.GetCurrentHeight()
}

func (s *lockedState) GetValidatorSet(height uint64, allychainID ids.ID) (map[ids.ShortID]uint64, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.s.GetValidatorSet(height, allychainID)
}

type noState struct{}

func NewNoState() State {
	return &noState{}
}

func (s *noState) GetMinimumHeight() (uint64, error) {
	return 0, nil
}

func (s *noState) GetCurrentHeight() (uint64, error) {
	return 0, nil
}

func (s *noState) GetValidatorSet(height uint64, allychainID ids.ID) (map[ids.ShortID]uint64, error) {
	return nil, nil
}
