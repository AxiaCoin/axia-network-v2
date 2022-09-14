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
	testGenesisConfigJSON = `
	{
		"networkID": 5,
		"allocations": [
			{
				"ethAddr": "0xb5f7c1eef626cc30f2c9fec3b63cd68a5b1fe716",
				"axcAddr": "Swap-test1rzzcqwf36aqdw5e6tuy48u5m9t990fpaprzdzk",
				"initialAmount": 17999994000000000000,
				"unlockSchedule": [
							{
							  "amount": 100000000000000,
							  "locktime": 1720170000
							}
						  ]
			},
			{
				"ethAddr": "0x0452a25e29673a8e8e5862bb6cbd496160b3e87e",
				"axcAddr": "Swap-test1d6dcmlus7zrxpkujdxsckaus0psfhpatumug78",
				"initialAmount": 17999994000000000000,
				"unlockSchedule": [
							{
							  "amount": 100000000000000,
							  "locktime": 1720170000
							}
						  ]
			},
			{
				"ethAddr": "0xdC8B26CF5dA71419F87b7fCb7652f368b9A4B095",
				"axcAddr": "Swap-test195urz4jext723lpdu0809r9uuchupamxghtmxs",
				"initialAmount": 17999994000000000000,
				"unlockSchedule": [
							{
							  "amount": 100000000000000,
							  "locktime": 1720170000
							}
						  ]
			},
			{
				"ethAddr": "0x91f3559Cad61d7fbd9983525515DA4A6c1793D7E",
				"axcAddr": "Swap-test1d8a4xklysgfgry4ltayvlnerjp6whsfp957tdf",
				"initialAmount": 17999994000000000000,
				"unlockSchedule": [
							{
							  "amount": 100000000000000,
							  "locktime": 1720170000
							}
						  ]
			},
			{
				"ethAddr": "0xAccf0E02593f0b6AaB1B2298a2Ce933c96C92623",
				"axcAddr": "Swap-test1addxfscvd28a5l89tmryjje9ax4nll7fmrqukz",
				"initialAmount": 17999994000000000000,
				"unlockSchedule": [
							{
							  "amount": 100000000000000,
							  "locktime": 1720170000
							}
						  ]
			},
			{
				"ethAddr": "0x8f90748dF95f0f69F4269A14265B8Bc966864912",
				"axcAddr": "Swap-test1qlvg8r7hkcs477twdmx0rch3m05q4plyag2lpx",
				"initialAmount": 17999994000000000000,
				"unlockSchedule": [
							{
							  "amount": 100000000000000,
							  "locktime": 1720170000
							}
						  ]
			},
			{
				"ethAddr": "0xBE4a848D12b09Ad476DbB3537cD7C8AEc6Aa95ba",
				"axcAddr": "Swap-test1dmd5t276mc5drfg25pkjnhk9n6q5ek3s5ezhy2",
				"initialAmount": 17999994000000000000,
				"unlockSchedule": [
							{
							  "amount": 100000000000000,
							  "locktime": 1720170000
							}
						  ]
			},
			{
				"ethAddr": "0xA7108D1F5231969060ca4d5bb151c0efaC8BF55C",
				"axcAddr": "Swap-test1exq6gsfuv0evfalfarth9p94w8l8nxslltufqy",
				"initialAmount": 17999994000000000000,
				"unlockSchedule": [
							{
							  "amount": 100000000000000,
							  "locktime": 1720170000
							}
						  ]
			},
			{
				"ethAddr": "0x3af8bf4CbD8303493D845EF7BBcF4A3B4e17BBE9",
				"axcAddr": "Swap-test1kl8ckmwdfp0je7jmdhnyy96y0ra0qr0sv0xn7u",
				"initialAmount": 17999994000000000000,
				"unlockSchedule": [
							{
							  "amount": 100000000000000,
							  "locktime": 1720170000
							}
						  ]
			},
			{
				"ethAddr": "0x15514038CEd579A705084Cd215C61DbC2b5B865f",
				"axcAddr": "Swap-test1wagx7284qfxf4ytpeq5l4rp6839ryls2gv0hep",
				"initialAmount": 17999994000000000000,
				"unlockSchedule": [
							{
							  "amount": 100000000000000,
							  "locktime": 1720170000
							}
						  ]
			}
		],
		"startTime": 1657179865,
		"initialStakeDuration": 31536000,
		"initialStakeDurationOffset": 5400,
		"initialStakedFunds": [
			"Swap-test1rzzcqwf36aqdw5e6tuy48u5m9t990fpaprzdzk",
			"Swap-test1d6dcmlus7zrxpkujdxsckaus0psfhpatumug78",
			"Swap-test195urz4jext723lpdu0809r9uuchupamxghtmxs",
			"Swap-test1d8a4xklysgfgry4ltayvlnerjp6whsfp957tdf",
			"Swap-test1addxfscvd28a5l89tmryjje9ax4nll7fmrqukz",
			"Swap-test1qlvg8r7hkcs477twdmx0rch3m05q4plyag2lpx"
		],
		"initialStakers": [
			{
				"nodeID": "NodeID-D8yKWBBXqpx8Qvi5jLZEUG1RURsciR8Au",
				"rewardAddress": "Swap-test1rzzcqwf36aqdw5e6tuy48u5m9t990fpaprzdzk",
				"nominationFee": 20000
			},
			{
				"nodeID": "NodeID-G7wmfRYiRcLzsir8nSvtaw2HTYJAUjoMc",
				"rewardAddress": "Swap-test1d6dcmlus7zrxpkujdxsckaus0psfhpatumug78",
				"nominationFee": 20000
			},
			{
				"nodeID": "NodeID-9h7tSWTDT1VgLNqtNcZi9f1Egm2xxjSM",
				"rewardAddress": "Swap-test195urz4jext723lpdu0809r9uuchupamxghtmxs",
				"nominationFee": 20000
			},
			{
				"nodeID": "NodeID-MMN1PTM26z8VGxk1vxEScX4dcsJhwfEjj",
				"rewardAddress": "Swap-test1d8a4xklysgfgry4ltayvlnerjp6whsfp957tdf",
				"nominationFee": 20000
			},
			{
				"nodeID": "NodeID-Fjks9n8sEKVKD7rdJ1BFBWzDA1tLD3jtN",
				"rewardAddress": "Swap-test1addxfscvd28a5l89tmryjje9ax4nll7fmrqukz",
				"nominationFee": 20000
			},
			{
				"nodeID": "NodeID-6Z4HDTEmdWFcrfDey9y9MwA1RGbYexYfE",
				"rewardAddress": "Swap-test1qlvg8r7hkcs477twdmx0rch3m05q4plyag2lpx",
				"nominationFee": 20000
			}
		],
		"axChainGenesis": "{\"config\":{\"chainId\":4000,\"homesteadBlock\":0,\"daoForkBlock\":0,\"daoForkSupport\":true,\"eip150Block\":0,\"eip150Hash\":\"0x2086799aeebeae135c246c65021c82b4e15a2c451340993aacfd2751886514f0\",\"eip155Block\":0,\"eip158Block\":0,\"byzantiumBlock\":0,\"constantinopleBlock\":0,\"petersburgBlock\":0,\"istanbulBlock\":0,\"muirGlacierBlock\":0},\"nonce\":\"0x0\",\"timestamp\":\"0x0\",\"extraData\":\"0x00\",\"gasLimit\":\"0x5f5e100\",\"difficulty\":\"0x0\",\"mixHash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\",\"coinbase\":\"0x0000000000000000000000000000000000000000\",\"alloc\":{\"0100000000000000000000000000000000000000\":{\"code\":\"0x7300000000000000000000000000000000000000003014608060405260043610603d5760003560e01c80631e010439146042578063b6510bb314606e575b600080fd5b605c60048036036020811015605657600080fd5b503560b1565b60408051918252519081900360200190f35b818015607957600080fd5b5060af60048036036080811015608e57600080fd5b506001600160a01b03813516906020810135906040810135906060013560b6565b005b30cd90565b836001600160a01b031681836108fc8690811502906040516000604051808303818888878c8acf9550505050505015801560f4573d6000803e3d6000fd5b505050505056fea26469706673582212201eebce970fe3f5cb96bf8ac6ba5f5c133fc2908ae3dcd51082cfee8f583429d064736f6c634300060a0033\",\"balance\":\"0x0\"}},\"number\":\"0x0\",\"gasUsed\":\"0x0\",\"parentHash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\"}",
		"message": "AXIA_System"
	  }`

	TestParams = Params{
		TxFeeConfig: TxFeeConfig{
			TxFee:                 units.MilliAxc,
			CreateAssetTxFee:      10 * units.MilliAxc,
			CreateAllychainTxFee:  1 * units.Axc,
			CreateBlockchainTxFee: 1 * units.Axc,
		},
		StakingConfig: StakingConfig{
			UptimeRequirement: .8, // 80%
			MinValidatorStake: 100 * units.KiloAxc,
			MaxValidatorStake: 3 * units.MegaAxc,
			MinNominatorStake: 20 * units.Axc,
			MinNominationFee:  20000, // 2%
			MinStakeDuration:  120 * 24 * time.Hour,
			MaxStakeDuration:  2 * 365 * 24 * time.Hour,
			RewardConfig: reward.Config{
				MaxConsumptionRate: .12 * reward.PercentDenominator,
				MinConsumptionRate: .10 * reward.PercentDenominator,
				MintingPeriod:      2 * 365 * 24 * time.Hour,
				SupplyCap:          uint128.Uint128{Hi: 180, Lo: 000000000000000000},
			},
		},
	}
)
