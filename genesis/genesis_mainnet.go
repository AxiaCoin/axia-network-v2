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
			  "ethAddr": "0x39b0a456f4a4157a61225602edbd895c6ee8b81f",
			  "axcAddr": "Swap-axc1gu9u289hz84ta9w5ejv3ukj8jm4umey4vsyctz",
			  "initialAmount": 7825086956521739000,
			  "unlockSchedule": [
					{
					  "amount": 10000000000000000,
					  "locktime": 1633824000
					}
				  ]
			},
			{
			  "ethAddr": "0x7c10871e4f78847694c170505ffd587368fe10e7",
			  "axcAddr": "Swap-axc1cl0jl6y5kx46jn9agavrfdg0rfnjcjwf6s24rm",
			  "initialAmount": 7825086956521739000,
			  "unlockSchedule": [
					{
					  "amount": 10000000000000000,
					  "locktime": 1633824000
					}
				  ]
			},
			{
			  "ethAddr": "0x23d9c532e04340a25c86f33e482c784aead159c7",
			  "axcAddr": "Swap-axc1xed5x26szd49wrf29vnmtnggtypvn3het22lkw",
			  "initialAmount": 7825086956521739000,
			  "unlockSchedule": [
					{
					  "amount": 10000000000000000,
					  "locktime": 1633824000
					}
				  ]
			},
			{
			  "ethAddr": "0x79b0c1ad7ec288a8c85ef729f85369a7a8e262e1",
			  "axcAddr": "Swap-axc1uufzezk96esdc4skdwkakzt6h0pklx5zsvajrc",
			  "initialAmount": 7825086956521739000,
			  "unlockSchedule": [
					{
					  "amount": 10000000000000000,
					  "locktime": 1633824000
					}
				  ]
			},
			{
			  "ethAddr": "0x29d2fc3015c8a44ca1507da6622fb157e833c3bf",
			  "axcAddr": "Swap-axc1dqagdh3ka8v6frqhs4y7lk6e83mdr4gawu3h3l",
			  "initialAmount": 7825086956521739000,
			  "unlockSchedule": [
					{
					  "amount": 10000000000000000,
					  "locktime": 1633824000
					}
				  ]
			},
			{
			  "ethAddr": "0x262bc126fc93fc5e4b3e82e9456f1105d352ae22",
			  "axcAddr": "Swap-axc1wjhtfqj4cvhvn9y34vdrwmmyht926659qu3ttz",
			  "initialAmount": 7825086956521739000,
			  "unlockSchedule": [
					{
					  "amount": 10000000000000000,
					  "locktime": 1633824000
					}
				  ]
			},
			{
			  "ethAddr": "0x9d94c99ca6a955ead6b5379aec9282e4d0cefd4e",
			  "axcAddr": "Swap-axc1cqu0d2zzyqfar54uzzrg86gzm3rqcchqj2uk6k",
			  "initialAmount": 7825086956521739000,
			  "unlockSchedule": [
					{
					  "amount": 10000000000000000,
					  "locktime": 1633824000
					}
				  ]
			},
			{
			  "ethAddr": "0xdf83080a3e7b6c61a5bff6eefb5a465fae2fd0ae",
			  "axcAddr": "Swap-axc1pjgm7v4k7g3stncu2yplesxmqvlckv35hnjl0y",
			  "initialAmount": 7825086956521739000,
			  "unlockSchedule": [
					{
					  "amount": 10000000000000000,
					  "locktime": 1633824000
					}
				  ]
			},
			{
			  "ethAddr": "0x0596814d0ac1ae058ff6a7e11349d98baf161184",
			  "axcAddr": "Swap-axc155kszdq5z7qkxfpt6wpaf4g0307za5myvsmgq6",
			  "initialAmount": 7825086956521739000,
			  "unlockSchedule": [
					{
					  "amount": 10000000000000000,
					  "locktime": 1633824000
					}
				  ]
			},
			{
			  "ethAddr": "0xbe14dabba65079a5ea37fb86f53ec0df0980f29f",
			  "axcAddr": "Swap-axc18zpxze0wtyteugl3pv4xz84xvs6jf657ex08yx",
			  "initialAmount": 7825086956521739000,
			  "unlockSchedule": [
					{
					  "amount": 10000000000000000,
					  "locktime": 1633824000
					}
				  ]
			},
			{
			  "ethAddr": "0xf22b8b13c5eb92830380da03c8a66062a7c3e85e",
			  "axcAddr": "Swap-axc1v3mekf7nc2pg8zvuxzqn854lcdxcs7pfrap9ml",
			  "initialAmount": 7825086956521739000,
			  "unlockSchedule": [
					{
					  "amount": 10000000000000000,
					  "locktime": 1633824000
					}
				  ]
			},
			{
			  "ethAddr": "0x344ec6a59fdbe424aff85a2e9e78956f6a04d3ea",
			  "axcAddr": "Swap-axc1hxdyuqnfssk7nrja25gt03delk5sn60z6yuy5n",
			  "initialAmount": 7825086956521739000,
			  "unlockSchedule": [
					{
					  "amount": 10000000000000000,
					  "locktime": 1633824000
					}
				  ]
			},
			{
			  "ethAddr": "0x8fb478d5a6cc14e35f41a0c766847b740d693725",
			  "axcAddr": "Swap-axc1ewfk7lja4zw55vh47fmxesx53vvydwm9jun5yw",
			  "initialAmount": 7825086956521739000,
			  "unlockSchedule": [
					{
					  "amount": 10000000000000000,
					  "locktime": 1633824000
					}
				  ]
			},
			{
			  "ethAddr": "0xc2b419523552245d07f654d229256a40fdff365e",
			  "axcAddr": "Swap-axc1j9szgk56z5kyfsac0m4ducdzhdlumjgq4xjgu6",
			  "initialAmount": 7825086956521739000,
			  "unlockSchedule": [
					{
					  "amount": 10000000000000000,
					  "locktime": 1633824000
					}
				  ]
			},
			{
			  "ethAddr": "0xb868ca1806bfa04113df65d41d607b00bf6450e7",
			  "axcAddr": "Swap-axc1wxlw2enmxn5cp9dm4ayp4wge80vma64qt9fjr6",
			  "initialAmount": 7825086956521739000,
			  "unlockSchedule": [
					{
					  "amount": 10000000000000000,
					  "locktime": 1633824000
					}
				  ]
			},
			{
			  "ethAddr": "0xf7a0309395569d455e5f18dd52ffa0ecf58353b0",
			  "axcAddr": "Swap-axc15wezxfm0f40l78g5yks5x3w2g2srmw8qxy0dhh",
			  "initialAmount": 7825086956521739000,
			  "unlockSchedule": [
					{
					  "amount": 10000000000000000,
					  "locktime": 1633824000
					}
				  ]
			},
			{
			  "ethAddr": "0x16e12f38b120cb9e3254878f3f05932053e4c8db",
			  "axcAddr": "Swap-axc1laxqd6uqyvl4afj2hw09ug2j4map8smtn3frtm",
			  "initialAmount": 7825086956521739000,
			  "unlockSchedule": [
					{
					  "amount": 10000000000000000,
					  "locktime": 1633824000
					}
				  ]
			},
			{
			  "ethAddr": "0x117cd4b0acada0caa278bcd0b1e0efefa462e279",
			  "axcAddr": "Swap-axc1kpzn855mts2s63vvf9jg4az4mwcx6spjz85wr3",
			  "initialAmount": 7825086956521739000,
			  "unlockSchedule": [
					{
					  "amount": 10000000000000000,
					  "locktime": 1633824000
					}
				  ]
			},
			{
			  "ethAddr": "0x2b87c2f0e7b3d41ce8a01012dc6b87fbf3c91d80",
			  "axcAddr": "Swap-axc1q6097uts3jlzpxn733xzxmgjl6cdf7k2jv4v5q",
			  "initialAmount": 7825086956521739000,
			  "unlockSchedule": [
					{
					  "amount": 10000000000000000,
					  "locktime": 1633824000
					}
				  ]
			},
			{
			  "ethAddr": "0xbdf541411f7b4b2ffbf3adf364d70e783197c9fb",
			  "axcAddr": "Swap-axc1k6dr7lg6rc7pj7ku9m2zl42fd8ejkact6y4zzj",
			  "initialAmount": 7825086956521739000,
			  "unlockSchedule": [
					{
					  "amount": 10000000000000000,
					  "locktime": 1633824000
					}
				  ]
			},
			{
			  "ethAddr": "0x53db297c3106e12f87e051e4b7709d070f269bf4",
			  "axcAddr": "Swap-axc1extn0qcv3a30kym9h0c5nt6awejrapqk8hp925",
			  "initialAmount": 7825086956521739000,
			  "unlockSchedule": [
					{
					  "amount": 10000000000000000,
					  "locktime": 1633824000
					}
				  ]
			},
			{
			  "ethAddr": "0x938f92351a0559be8014f65c9927560442b11221",
			  "axcAddr": "Swap-axc1m4lu70uxmwtdq4wfnts2d3y409s9cs4he0r96p",
			  "initialAmount": 7825086956521739000,
			  "unlockSchedule": [
					{
					  "amount": 10000000000000000,
					  "locktime": 1633824000
					}
				  ]
			},
			{
			  "ethAddr": "0xb61d1932a6b51f3ab731caf4f9a32fae156785f4",
			  "axcAddr": "Swap-axc1xeafvyxcrlcf43cxanp0eltgglse07364qurqn",
			  "initialAmount": 7825086956521739000,
			  "unlockSchedule": [
					{
					  "amount": 10000000000000000,
					  "locktime": 1633824000
					}
				  ]
			}
		],
		"startTime": 1656937964,
		"initialStakeDuration": 7776000,
		"initialStakeDurationOffset": 5400,
		"initialStakedFunds": [
			"Swap-axc1gu9u289hz84ta9w5ejv3ukj8jm4umey4vsyctz",
			"Swap-axc1cl0jl6y5kx46jn9agavrfdg0rfnjcjwf6s24rm",
			"Swap-axc1xed5x26szd49wrf29vnmtnggtypvn3het22lkw",
			"Swap-axc1uufzezk96esdc4skdwkakzt6h0pklx5zsvajrc",
			"Swap-axc1dqagdh3ka8v6frqhs4y7lk6e83mdr4gawu3h3l",
			"Swap-axc1wjhtfqj4cvhvn9y34vdrwmmyht926659qu3ttz",
			"Swap-axc1cqu0d2zzyqfar54uzzrg86gzm3rqcchqj2uk6k",
			"Swap-axc1pjgm7v4k7g3stncu2yplesxmqvlckv35hnjl0y",
			"Swap-axc155kszdq5z7qkxfpt6wpaf4g0307za5myvsmgq6",
			"Swap-axc18zpxze0wtyteugl3pv4xz84xvs6jf657ex08yx",
			"Swap-axc1v3mekf7nc2pg8zvuxzqn854lcdxcs7pfrap9ml",
			"Swap-axc1hxdyuqnfssk7nrja25gt03delk5sn60z6yuy5n",
			"Swap-axc1ewfk7lja4zw55vh47fmxesx53vvydwm9jun5yw",
			"Swap-axc1j9szgk56z5kyfsac0m4ducdzhdlumjgq4xjgu6",
			"Swap-axc1wxlw2enmxn5cp9dm4ayp4wge80vma64qt9fjr6",
			"Swap-axc15wezxfm0f40l78g5yks5x3w2g2srmw8qxy0dhh",
			"Swap-axc1laxqd6uqyvl4afj2hw09ug2j4map8smtn3frtm",
			"Swap-axc1kpzn855mts2s63vvf9jg4az4mwcx6spjz85wr3",
			"Swap-axc1q6097uts3jlzpxn733xzxmgjl6cdf7k2jv4v5q",
			"Swap-axc1k6dr7lg6rc7pj7ku9m2zl42fd8ejkact6y4zzj",
			"Swap-axc1extn0qcv3a30kym9h0c5nt6awejrapqk8hp925",
			"Swap-axc1m4lu70uxmwtdq4wfnts2d3y409s9cs4he0r96p",
			"Swap-axc1xeafvyxcrlcf43cxanp0eltgglse07364qurqn"
		],
		"initialStakers": [
			{
			  "nodeID": "NodeID-FVJiVBMXg69Eys8v3K3fUZXM21fW5k7Ts",
			  "rewardAddress": "Swap-axc1gu9u289hz84ta9w5ejv3ukj8jm4umey4vsyctz",
			  "nominationFee": 200000
			},
			{
			  "nodeID": "NodeID-8ofNV8sb3J9rwaC3HuXcPFAkNEqF8oJzV",
			  "rewardAddress": "Swap-axc1cl0jl6y5kx46jn9agavrfdg0rfnjcjwf6s24rm",
			  "nominationFee": 200000
			},
			{
			  "nodeID": "NodeID-BLaJAbw7aYCmZtYoemCBpowijcBPq7yZs",
			  "rewardAddress": "Swap-axc1xed5x26szd49wrf29vnmtnggtypvn3het22lkw",
			  "nominationFee": 200000
			},
			{
			  "nodeID": "NodeID-FTirsWrt7FJVTTvQejEQpJERuJDKk9BW2",
			  "rewardAddress": "Swap-axc1uufzezk96esdc4skdwkakzt6h0pklx5zsvajrc",
			  "nominationFee": 200000
			},
			{
			  "nodeID": "NodeID-Gym2UoBo3b4tp3pukZzDjxJefNc5XE8Cw",
			  "rewardAddress": "Swap-axc1dqagdh3ka8v6frqhs4y7lk6e83mdr4gawu3h3l",
			  "nominationFee": 200000
			},
			{
			  "nodeID": "NodeID-AkZUoGaJUWrfGUBtk7Wrsyvu5DZKVkxKQ",
			  "rewardAddress": "Swap-axc1wjhtfqj4cvhvn9y34vdrwmmyht926659qu3ttz",
			  "nominationFee": 200000
			},
			{
			  "nodeID": "NodeID-FuzyvQMfFRhAWiiQLvcDVErbNoA2czuYp",
			  "rewardAddress": "Swap-axc1cqu0d2zzyqfar54uzzrg86gzm3rqcchqj2uk6k",
			  "nominationFee": 200000
			},
			{
			  "nodeID": "NodeID-Ag6dycckf8nwo2EYrvfnuw4oAw8xTz6P3",
			  "rewardAddress": "Swap-axc1pjgm7v4k7g3stncu2yplesxmqvlckv35hnjl0y",
			  "nominationFee": 200000
			},
			{
			  "nodeID": "NodeID-PUWVULUUXozYBvUQwoPxQFMXT13MqRBAK",
			  "rewardAddress": "Swap-axc155kszdq5z7qkxfpt6wpaf4g0307za5myvsmgq6",
			  "nominationFee": 200000
			},
			{
			  "nodeID": "NodeID-MbjtKYV1EEbRdaK1d396g54VrVnTXet5Q",
			  "rewardAddress": "Swap-axc18zpxze0wtyteugl3pv4xz84xvs6jf657ex08yx",
			  "nominationFee": 200000
			},
			{
			  "nodeID": "NodeID-FeFgdAG3v8YJidF4xR8E3zSHUEuNRtBfB",
			  "rewardAddress": "Swap-axc1v3mekf7nc2pg8zvuxzqn854lcdxcs7pfrap9ml",
			  "nominationFee": 200000
			},
			{
			  "nodeID": "NodeID-2eAGdTFKwdi63NDvADr99qoDANmY11uun",
			  "rewardAddress": "Swap-axc1hxdyuqnfssk7nrja25gt03delk5sn60z6yuy5n",
			  "nominationFee": 200000
			},
			{
			  "nodeID": "NodeID-pbDBA6Gfy9ShgscRm83kmJpVECqyh2EV",
			  "rewardAddress": "Swap-axc1ewfk7lja4zw55vh47fmxesx53vvydwm9jun5yw",
			  "nominationFee": 200000
			},
			{
			  "nodeID": "NodeID-KbQdsqnna4KwQfsRksNMah1ZbpQgvqtrX",
			  "rewardAddress": "Swap-axc1j9szgk56z5kyfsac0m4ducdzhdlumjgq4xjgu6",
			  "nominationFee": 200000
			},
			{
			  "nodeID": "NodeID-Ne4NyLV1WAyyXoRAzSSSTbc7pU8jbFTez",
			  "rewardAddress": "Swap-axc1wxlw2enmxn5cp9dm4ayp4wge80vma64qt9fjr6",
			  "nominationFee": 200000
			},
			{
			  "nodeID": "NodeID-Mx3reKpoVipQVUzH41qx5gJZmNURgvFJD",
			  "rewardAddress": "Swap-axc15wezxfm0f40l78g5yks5x3w2g2srmw8qxy0dhh",
			  "nominationFee": 200000
			},
			{
			  "nodeID": "NodeID-CksSwGj4tSAyyRVrFavs6p9farwwtEC1b",
			  "rewardAddress": "Swap-axc1laxqd6uqyvl4afj2hw09ug2j4map8smtn3frtm",
			  "nominationFee": 200000
			},
			{
			  "nodeID": "NodeID-QHHAyHWvLEJsawMTYgZfAT66erXRVqtHu",
			  "rewardAddress": "Swap-axc1kpzn855mts2s63vvf9jg4az4mwcx6spjz85wr3",
			  "nominationFee": 200000
			},
			{
			  "nodeID": "NodeID-21zDZTDjnKcLE8yMy1SXsC5oYvJq1m1Ti",
			  "rewardAddress": "Swap-axc1q6097uts3jlzpxn733xzxmgjl6cdf7k2jv4v5q",
			  "nominationFee": 200000
			},
			{
			  "nodeID": "NodeID-Q1KAsFkpTVszwbAtANydsRTLj9jjs8PAP",
			  "rewardAddress": "Swap-axc1k6dr7lg6rc7pj7ku9m2zl42fd8ejkact6y4zzj",
			  "nominationFee": 200000
			},
			{
			  "nodeID": "NodeID-HPz5qwnbvsiz4oqmsf7zr8NeWsmbSnAep",
			  "rewardAddress": "Swap-axc1extn0qcv3a30kym9h0c5nt6awejrapqk8hp925",
			  "nominationFee": 200000
			},
			{
			  "nodeID": "NodeID-G3FjbZ7rSrBqL11BpZG89RJPXkP2cUaiL",
			  "rewardAddress": "Swap-axc1m4lu70uxmwtdq4wfnts2d3y409s9cs4he0r96p",
			  "nominationFee": 200000
			},
			{
			  "nodeID": "NodeID-5wFDMxCVVpvZip4oUrMyunwRde1aU9Vmb",
			  "rewardAddress": "Swap-axc1xeafvyxcrlcf43cxanp0eltgglse07364qurqn",
			  "nominationFee": 200000
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
			MinValidatorStake: 1 * units.MegaAxc,
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
