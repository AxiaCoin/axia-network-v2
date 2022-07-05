// Copyright (C) 2019-2021, Axia Systems, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package genesis

import (
	"encoding/hex"
	"errors"

	"github.com/axiacoin/axia-network-v2/ids"
	"github.com/axiacoin/axia-network-v2/utils/constants"
	"github.com/axiacoin/axia-network-v2/utils/formatting"
)

var errInvalidETHAddress = errors.New("invalid eth address")

type UnparsedAllocation struct {
	ETHAddr        string         `json:"ethAddr"`
	AXCAddr       string         `json:"axcAddr"`
	InitialAmount  uint64         `json:"initialAmount"`
	UnlockSchedule []LockedAmount `json:"unlockSchedule"`
}

func (ua UnparsedAllocation) Parse() (Allocation, error) {
	a := Allocation{
		InitialAmount:  ua.InitialAmount,
		UnlockSchedule: ua.UnlockSchedule,
	}

	if len(ua.ETHAddr) < 2 {
		return a, errInvalidETHAddress
	}

	ethAddrBytes, err := hex.DecodeString(ua.ETHAddr[2:])
	if err != nil {
		return a, err
	}
	ethAddr, err := ids.ToShortID(ethAddrBytes)
	if err != nil {
		return a, err
	}
	a.ETHAddr = ethAddr

	_, _, axcAddrBytes, err := formatting.ParseAddress(ua.AXCAddr)
	if err != nil {
		return a, err
	}
	axcAddr, err := ids.ToShortID(axcAddrBytes)
	if err != nil {
		return a, err
	}
	a.AXCAddr = axcAddr

	return a, nil
}

type UnparsedStaker struct {
	NodeID        string `json:"nodeID"`
	RewardAddress string `json:"rewardAddress"`
	NominationFee uint32 `json:"nominationFee"`
}

func (us UnparsedStaker) Parse() (Staker, error) {
	s := Staker{
		NominationFee: us.NominationFee,
	}

	nodeID, err := ids.ShortFromPrefixedString(us.NodeID, constants.NodeIDPrefix)
	if err != nil {
		return s, err
	}
	s.NodeID = nodeID

	_, _, axcAddrBytes, err := formatting.ParseAddress(us.RewardAddress)
	if err != nil {
		return s, err
	}
	axcAddr, err := ids.ToShortID(axcAddrBytes)
	if err != nil {
		return s, err
	}
	s.RewardAddress = axcAddr
	return s, nil
}

// UnparsedConfig contains the genesis addresses used to construct a genesis
type UnparsedConfig struct {
	NetworkID uint32 `json:"networkID"`

	Allocations []UnparsedAllocation `json:"allocations"`

	StartTime                  uint64           `json:"startTime"`
	InitialStakeDuration       uint64           `json:"initialStakeDuration"`
	InitialStakeDurationOffset uint64           `json:"initialStakeDurationOffset"`
	InitialStakedFunds         []string         `json:"initialStakedFunds"`
	InitialStakers             []UnparsedStaker `json:"initialStakers"`

	AXChainGenesis string `json:"axChainGenesis"`

	Message string `json:"message"`
}

func (uc UnparsedConfig) Parse() (Config, error) {
	c := Config{
		NetworkID:                  uc.NetworkID,
		Allocations:                make([]Allocation, len(uc.Allocations)),
		StartTime:                  uc.StartTime,
		InitialStakeDuration:       uc.InitialStakeDuration,
		InitialStakeDurationOffset: uc.InitialStakeDurationOffset,
		InitialStakedFunds:         make([]ids.ShortID, len(uc.InitialStakedFunds)),
		InitialStakers:             make([]Staker, len(uc.InitialStakers)),
		AXChainGenesis:              uc.AXChainGenesis,
		Message:                    uc.Message,
	}
	for i, ua := range uc.Allocations {
		a, err := ua.Parse()
		if err != nil {
			return c, err
		}
		c.Allocations[i] = a
	}
	for i, isa := range uc.InitialStakedFunds {
		_, _, axcAddrBytes, err := formatting.ParseAddress(isa)
		if err != nil {
			return c, err
		}
		axcAddr, err := ids.ToShortID(axcAddrBytes)
		if err != nil {
			return c, err
		}
		c.InitialStakedFunds[i] = axcAddr
	}
	for i, uis := range uc.InitialStakers {
		is, err := uis.Parse()
		if err != nil {
			return c, err
		}
		c.InitialStakers[i] = is
	}
	return c, nil
}
