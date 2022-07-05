// Copyright (C) 2019-2021, Axia Systems, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package genesis

import (
	"time"

	"github.com/axiacoin/axia-network-v2/utils/uint128"
	"github.com/axiacoin/axia-network-v2/utils/units"
	"github.com/axiacoin/axia-network-v2/vms/platformvm/reward"
)

var (
	mainnetGenesisConfigJSON = `{
		"networkID": 1,
		"allocations": [
			{
				"ethAddr": "0xF57b05E680EB007f7bB06931EAF4974151A08e51",
				"axcAddr": "Swap-custom1qu460602uxetjyhewsxkae5cp7j68zy6gz6cvj",
				"initialAmount": 7825086956521739130,
				"unlockSchedule": [
				{
					"amount": 1000000000000000,
					"locktime": 1664957273
					}
					]
			},
			{
				"ethAddr": "0xb4caAc709D045AD998B50aB2B2202F109f5Dcf76",
				"axcAddr": "Swap-custom173x6vjspg95nn4n68cerdhjnhwtehkc4ey6dlf",
				"initialAmount": 7825086956521739130,
				"unlockSchedule": [
				{
					"amount": 1000000000000000,
					"locktime": 1664957273
					}
					]
			},
			{
				"ethAddr": "0x725c7698d19F89C395789d58714c2Adb287242aD",
				"axcAddr": "Swap-custom197g3cfnq5k3lemp5h9l5dthe4plzdldge2kh0m",
				"initialAmount": 7825086956521739130,
				"unlockSchedule": [
				{
					"amount": 1000000000000000,
					"locktime": 1664957273
					}
					]
			},
			{
				"ethAddr": "0x22CB8Bbb1Da74c3B93c3C9442a01D01c98E85b82",
				"axcAddr": "Swap-custom1s2kquu7shue6aq9xg8zklm0jvjnwzha7qsv8p3",
				"initialAmount": 7825086956521739130,
				"unlockSchedule": [
				{
					"amount": 1000000000000000,
					"locktime": 1664957273
					}
					]
			},
			{
				"ethAddr": "0xb37A616Fd5d3ffabf8bf8AE21b975A732A4C45A2",
				"axcAddr": "Swap-custom15lm8nkxt0rz4nlkrx3emqj5rlkskwt545vgft3",
				"initialAmount": 7825086956521739130,
				"unlockSchedule": [
				{
					"amount": 1000000000000000,
					"locktime": 1664957273
					}
					]
			}
		],
		"startTime": 1656937964,
		"initialStakeDuration": 7776000,
		"initialStakeDurationOffset": 5400,
		"initialStakedFunds": [
			"Swap-custom1qu460602uxetjyhewsxkae5cp7j68zy6gz6cvj",
			"Swap-custom173x6vjspg95nn4n68cerdhjnhwtehkc4ey6dlf",
			"Swap-custom197g3cfnq5k3lemp5h9l5dthe4plzdldge2kh0m",
			"Swap-custom1s2kquu7shue6aq9xg8zklm0jvjnwzha7qsv8p3",
			"Swap-custom15lm8nkxt0rz4nlkrx3emqj5rlkskwt545vgft3"
		],
		"initialStakers": [
			{
				"nodeID": "NodeID-FVJiVBMXg69Eys8v3K3fUZXM21fW5k7Ts",
				"rewardAddress": "Swap-custom1qu460602uxetjyhewsxkae5cp7j68zy6gz6cvj",
				"delegationFee": 200000
			},
			{
				"nodeID": "NodeID-8ofNV8sb3J9rwaC3HuXcPFAkNEqF8oJzV",
				"rewardAddress": "Swap-custom173x6vjspg95nn4n68cerdhjnhwtehkc4ey6dlf",
				"delegationFee": 200000
			},
			{
				"nodeID": "NodeID-BLaJAbw7aYCmZtYoemCBpowijcBPq7yZs",
				"rewardAddress": "Swap-custom197g3cfnq5k3lemp5h9l5dthe4plzdldge2kh0m",
				"delegationFee": 200000
			},
			{
				"nodeID": "NodeID-FTirsWrt7FJVTTvQejEQpJERuJDKk9BW2",
				"rewardAddress": "Swap-custom1s2kquu7shue6aq9xg8zklm0jvjnwzha7qsv8p3",
				"delegationFee": 200000
			},
			{
				"nodeID": "NodeID-Gym2UoBo3b4tp3pukZzDjxJefNc5XE8Cw",
				"rewardAddress": "Swap-custom15lm8nkxt0rz4nlkrx3emqj5rlkskwt545vgft3",
				"delegationFee": 200000
			}
		],
		"axChainGenesis": "{\"config\":{\"chainId\":4001,\"homesteadBlock\":0,\"daoForkBlock\":0,\"daoForkSupport\":true,\"eip150Block\":0,\"eip150Hash\":\"0x2086799aeebeae135c246c65021c82b4e15a2c451340993aacfd2751886514f0\",\"eip155Block\":0,\"eip158Block\":0,\"byzantiumBlock\":0,\"constantinopleBlock\":0,\"petersburgBlock\":0,\"istanbulBlock\":0,\"muirGlacierBlock\":0},\"nonce\":\"0x0\",\"timestamp\":\"0x0\",\"extraData\":\"0x00\",\"gasLimit\":\"0x5f5e100\",\"difficulty\":\"0x0\",\"mixHash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\",\"coinbase\":\"0x0000000000000000000000000000000000000000\",\"alloc\":{\"0100000000000000000000000000000000000000\":{\"code\":\"0x7300000000000000000000000000000000000000003014608060405260043610603d5760003560e01c80631e010439146042578063b6510bb314606e575b600080fd5b605c60048036036020811015605657600080fd5b503560b1565b60408051918252519081900360200190f35b818015607957600080fd5b5060af60048036036080811015608e57600080fd5b506001600160a01b03813516906020810135906040810135906060013560b6565b005b30cd90565b836001600160a01b031681836108fc8690811502906040516000604051808303818888878c8acf9550505050505015801560f4573d6000803e3d6000fd5b505050505056fea26469706673582212201eebce970fe3f5cb96bf8ac6ba5f5c133fc2908ae3dcd51082cfee8f583429d064736f6c634300060a0033\",\"balance\":\"0x0\"}},\"number\":\"0x0\",\"gasUsed\":\"0x0\",\"parentHash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\"}",
		"message": "AXIA_System"
	}`

	// MainnetParams are the params used for mainnet
	MainnetParams = Params{
		TxFeeConfig: TxFeeConfig{
			TxFee:                 units.MilliAxc,
			CreateAssetTxFee:      10 * units.MilliAxc,
			CreateAllychainTxFee:  1 * units.Axc,
			CreateBlockchainTxFee: 1 * units.Axc,
		},
		StakingConfig: StakingConfig{
			UptimeRequirement: .8, // 80%
			MinValidatorStake: 2 * units.KiloAxc,
			MaxValidatorStake: 3 * units.MegaAxc,
			MinNominatorStake: 25 * units.Axc,
			MinDelegationFee:  20000, // 2%
			MinStakeDuration:  2 * 7 * 24 * time.Hour,
			MaxStakeDuration:  365 * 24 * time.Hour,
			RewardConfig: reward.Config{
				MaxConsumptionRate: .12 * reward.PercentDenominator,
				MinConsumptionRate: .10 * reward.PercentDenominator,
				MintingPeriod:      365 * 24 * time.Hour,
				SupplyCap:          uint128.Zero.Add64(720 * units.MegaAxc),
			},
		},
	}
)
