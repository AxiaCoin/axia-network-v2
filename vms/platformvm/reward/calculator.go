// Copyright (C) 2019-2021, Axia Systems, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package reward

import (
	"math/big"
	"time"

	"github.com/axiacoin/axia-network-v2/utils/uint128"
)

var _ Calculator = &calculator{}

type Calculator interface {
	Calculate(stakedDuration time.Duration, stakedAmount uint64, currentSupply uint128.Uint128) uint64
}

type calculator struct {
	maxSubMinConsumptionRate *big.Int
	minConsumptionRate       *big.Int
	mintingPeriod            *big.Int
	supplyCap                uint128.Uint128
}

func NewCalculator(c Config) Calculator {
	return &calculator{
		maxSubMinConsumptionRate: new(big.Int).SetUint64(c.MaxConsumptionRate - c.MinConsumptionRate),
		minConsumptionRate:       new(big.Int).SetUint64(c.MinConsumptionRate),
		mintingPeriod:            new(big.Int).SetUint64(uint64(c.MintingPeriod)),
		supplyCap:                c.SupplyCap,
	}
}

// Reward returns the amount of tokens to reward the staker with.
//
// RemainingSupply = SupplyCap - ExistingSupply
// PortionOfExistingSupply = StakedAmount / ExistingSupply
// PortionOfStakingDuration = StakingDuration / MaximumStakingDuration
// MintingRate = MinMintingRate + MaxSubMinMintingRate * PortionOfStakingDuration
// Reward = RemainingSupply * PortionOfExistingSupply * MintingRate * PortionOfStakingDuration
func (c *calculator) Calculate(stakedDuration time.Duration, stakedAmount uint64, currentSupply uint128.Uint128) uint64 {
	return 0
	// bigStakedDuration := new(big.Int).SetUint64(uint64(stakedDuration))
	// bigStakedAmount := new(big.Int).SetUint64(stakedAmount)
	// bigCurrentSupply, _ := new(big.Int).SetString(currentSupply.String(), 10)

	// adjustedConsumptionRateNumerator := new(big.Int).Mul(c.maxSubMinConsumptionRate, bigStakedDuration)
	// adjustedMinConsumptionRateNumerator := new(big.Int).Mul(c.minConsumptionRate, c.mintingPeriod)
	// adjustedConsumptionRateNumerator.Add(adjustedConsumptionRateNumerator, adjustedMinConsumptionRateNumerator)
	// adjustedConsumptionRateDenominator := new(big.Int).Mul(c.mintingPeriod, consumptionRateDenominator)

	// reward, _ := new(big.Int).SetString(c.supplyCap.Sub(currentSupply).String(), 10)
	// reward.Mul(reward, adjustedConsumptionRateNumerator)
	// reward.Mul(reward, bigStakedAmount)
	// reward.Mul(reward, bigStakedDuration)
	// reward.Div(reward, adjustedConsumptionRateDenominator)
	// reward.Div(reward, bigCurrentSupply)
	// reward.Div(reward, c.mintingPeriod)

	// return reward.Uint64()
}
