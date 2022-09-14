// Copyright (C) 2019-2021, Axia Systems, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package platformvm

import (
	"time"

	"github.com/axiacoin/axia-network-v2/chains"
	"github.com/axiacoin/axia-network-v2/ids"
	"github.com/axiacoin/axia-network-v2/snow"
	"github.com/axiacoin/axia-network-v2/snow/uptime"
	"github.com/axiacoin/axia-network-v2/snow/validators"
	"github.com/axiacoin/axia-network-v2/vms/platformvm/reward"
)

// Factory can create new instances of the Platform Chain
type Factory struct {
	// The node's chain manager
	Chains chains.Manager

	// Node's validator set maps allychainID -> validators of the allychain
	Validators validators.Manager

	// Provides access to the uptime manager as a thread safe data structure
	UptimeLockedCalculator uptime.LockedCalculator

	// True if the node is being run with staking enabled
	StakingEnabled bool

	// Set of allychains that this node is validating
	WhitelistedAllychains ids.Set

	// Fee that must be burned by every create staker transaction
	AddStakerTxFee uint64

	// Fee that is burned by every non-state creating transaction
	TxFee uint64

	// Fee that must be burned by every state creating transaction before AP3
	CreateAssetTxFee uint64

	// Fee that must be burned by every allychain creating transaction after AP3
	CreateAllychainTxFee uint64

	// Fee that must be burned by every blockchain creating transaction after AP3
	CreateBlockchainTxFee uint64

	// The minimum amount of tokens one must bond to be a validator
	MinValidatorStake uint64

	// The maximum amount of tokens that can be bonded on a validator
	MaxValidatorStake uint64

	// Minimum stake, in nAXC, that can be nominated on the primary network
	MinNominatorStake uint64

	// Minimum fee that can be charged for nomination
	MinNominationFee uint32

	// UptimePercentage is the minimum uptime required to be rewarded for staking
	UptimePercentage float64

	// Minimum amount of time to allow a staker to stake
	MinStakeDuration time.Duration

	// Maximum amount of time to allow a staker to stake
	MaxStakeDuration time.Duration

	// Config for the minting function
	RewardConfig reward.Config

	// Time of the AP3 network upgrade
	ApricotPhase3Time time.Time

	// Time of the AP4 network upgrade
	ApricotPhase4Time time.Time

	// Time of the AP5 network upgrade
	ApricotPhase5Time time.Time
}

// New returns a new instance of the Platform Chain
func (f *Factory) New(*snow.Context) (interface{}, error) {
	return &VM{Factory: *f}, nil
}
