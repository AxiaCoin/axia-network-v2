// Copyright (C) 2019-2021, Axia Systems, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package genesis

import (
	"time"

	"github.com/axiacoin/axia-network-v2/utils/constants"
	"github.com/axiacoin/axia-network-v2/vms/platformvm/reward"
)

type StakingConfig struct {
	// Staking uptime requirements
	UptimeRequirement float64 `json:"uptimeRequirement"`
	// Minimum stake, in nAXC, required to validate the primary network
	MinValidatorStake uint64 `json:"minValidatorStake"`
	// Maximum stake, in nAXC, allowed to be placed on a single validator in
	// the primary network
	MaxValidatorStake uint64 `json:"maxValidatorStake"`
	// Minimum stake, in nAXC, that can be nominated on the primary network
	MinNominatorStake uint64 `json:"minNominatorStake"`
	// Minimum nomination fee, in the range [0, 1000000], that can be charged
	// for nomination on the primary network.
	MinNominationFee uint32 `json:"minNominationFee"`
	// MinStakeDuration is the minimum amount of time a validator can validate
	// for in a single period.
	MinStakeDuration time.Duration `json:"minStakeDuration"`
	// MaxStakeDuration is the maximum amount of time a validator can validate
	// for in a single period.
	MaxStakeDuration time.Duration `json:"maxStakeDuration"`
	// RewardConfig is the config for the reward function.
	RewardConfig reward.Config `json:"rewardConfig"`
}

type TxFeeConfig struct {
	// Transaction fee
	TxFee uint64 `json:"txFee"`
	// Transaction fee for create asset transactions
	CreateAssetTxFee uint64 `json:"createAssetTxFee"`
	// Transaction fee for create allychain transactions
	CreateAllychainTxFee uint64 `json:"createAllychainTxFee"`
	// Transaction fee for create blockchain transactions
	CreateBlockchainTxFee uint64 `json:"createBlockchainTxFee"`
}

type Params struct {
	StakingConfig
	TxFeeConfig
}

func GetTxFeeConfig(networkID uint32) TxFeeConfig {
	switch networkID {
	case constants.MainnetID:
		return MainnetParams.TxFeeConfig
	case constants.TestID:
		return TestParams.TxFeeConfig
	case constants.LocalID:
		return LocalParams.TxFeeConfig
	default:
		return LocalParams.TxFeeConfig
	}
}

func GetStakingConfig(networkID uint32) StakingConfig {
	switch networkID {
	case constants.MainnetID:
		return MainnetParams.StakingConfig
	case constants.TestID:
		return TestParams.StakingConfig
	case constants.LocalID:
		return LocalParams.StakingConfig
	default:
		return LocalParams.StakingConfig
	}
}
