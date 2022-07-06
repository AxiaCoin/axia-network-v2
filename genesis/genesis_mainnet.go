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
				"ethAddr": "0x777366a0b26fed1f8915dbfbbdbac5c7eb618fd9",
				"axcAddr": "Swap-axc1lgr4vjsuy5lcvg2660t5ufqp0yfv9rkv0yqkd9",
				"initialAmount": 7825086956521739000,
				"unlockSchedule": [
					{
						"amount": 10000000000000000,
						"locktime": 1633824000
					}
				]
			  },
			  {
				"ethAddr": "0x57bf5570117472719691dcc5dc1c33338affda72",
				"axcAddr": "Swap-axc194dsz5fxexgkx93jwj5a8zqjcw6pcnfzlhg95f",
				"initialAmount": 7825086956521739000,
				"unlockSchedule": [
					{
						"amount": 10000000000000000,
						"locktime": 1633824000
					}
				]
			  },
			  {
				"ethAddr": "0x0efd26e3ea1e7df85c2b13695b51d3e3a3e9ef42",
				"axcAddr": "Swap-axc1hz268ehr8ur46eu8gwzrczc5sxgz0ww7qt5tt3",
				"initialAmount": 7825086956521739000,
				"unlockSchedule": [
					{
						"amount": 10000000000000000,
						"locktime": 1633824000
					}
				]
			  },
			  {
				"ethAddr": "0x8fc3cdef753bc368d3efdc5cc55d64becfce980f",
				"axcAddr": "Swap-axc1rm97wskg54pljccf5tc6xykhgk76z582n9rg54",
				"initialAmount": 7825086956521739000,
				"unlockSchedule": [
					{
						"amount": 10000000000000000,
						"locktime": 1633824000
					}
				]
			  },
			  {
				"ethAddr": "0x6f202346e30847cd2ff558a17983801e7154111e",
				"axcAddr": "Swap-axc1lypz8hpz9dml0jzmxmqvg5kl49pc0kepk478x3",
				"initialAmount": 7825086956521739000,
				"unlockSchedule": [
					{
						"amount": 10000000000000000,
						"locktime": 1633824000
					}
				]
			  },
			  {
				"ethAddr": "0x770dcbf12c940286b3d560832935d2b8844b5541",
				"axcAddr": "Swap-axc195ul7h2c0e28lnmz3nv26dukk3pkqayf888jx2",
				"initialAmount": 7825086956521739000,
				"unlockSchedule": [
					{
						"amount": 10000000000000000,
						"locktime": 1633824000
					}
				]
			  },
			  {
				"ethAddr": "0xab2aa0c562cd632ad89e4f1c97db2f246877ba09",
				"axcAddr": "Swap-axc1rnta60c0vvdhxly9rgf9g2atsh803lvfa2y4z8",
				"initialAmount": 7825086956521739000,
				"unlockSchedule": [
					{
						"amount": 10000000000000000,
						"locktime": 1633824000
					}
				]
			  },
			  {
				"ethAddr": "0x324e4d7f5047d1c61dc947ecbdd3c1566d3eae52",
				"axcAddr": "Swap-axc1sx3hm9zjcx76ecd4fyej085nrkz2w4hfq42pa2",
				"initialAmount": 7825086956521739000,
				"unlockSchedule": [
					{
						"amount": 10000000000000000,
						"locktime": 1633824000
					}
				]
			  },
			  {
				"ethAddr": "0x17a13211ad58bb56b4c86e1a38126a1319e0efb9",
				"axcAddr": "Swap-axc1pp7cvggn2zflzxvur9vys3ywrwh5wt2xjqy483",
				"initialAmount": 7825086956521739000,
				"unlockSchedule": [
					{
						"amount": 10000000000000000,
						"locktime": 1633824000
					}
				]
			  },
			  {
				"ethAddr": "0x429e93cfc908fc9c74f83988f8e71d71eda7ff5a",
				"axcAddr": "Swap-axc1juwkcwtrulkkkf7p5zvlu4ruznxmlxu4vnz2je",
				"initialAmount": 7825086956521739000,
				"unlockSchedule": [
					{
						"amount": 10000000000000000,
						"locktime": 1633824000
					}
				]
			  },
			  {
				"ethAddr": "0xc74151f0b15b9d6f9ae83cf40b4182364622bcf0",
				"axcAddr": "Swap-axc1dn0f8tz9de7v9nrqjel4qxuadswftc8tmde8a4",
				"initialAmount": 7825086956521739000,
				"unlockSchedule": [
					{
						"amount": 10000000000000000,
						"locktime": 1633824000
					}
				]
			  },
			  {
				"ethAddr": "0xabce54f4176c419070b4f9254e1140a1dfcf29e4",
				"axcAddr": "Swap-axc14dlest7mg83yggqxa425te4lfx6hzpk4dpz2ad",
				"initialAmount": 7825086956521739000,
				"unlockSchedule": [
					{
						"amount": 10000000000000000,
						"locktime": 1633824000
					}
				]
			  },
			  {
				"ethAddr": "0x5ec05e442feff963bc0801a9ebe3879dd7dc22dc",
				"axcAddr": "Swap-axc1a2hnt973ar2wntw8mtn7nyczeqwzlqtzydj4kr",
				"initialAmount": 7825086956521739000,
				"unlockSchedule": [
					{
						"amount": 10000000000000000,
						"locktime": 1633824000
					}
				]
			  },
			  {
				"ethAddr": "0x9606c0379c8f15c0b711db7d376fd61080e271c1",
				"axcAddr": "Swap-axc1qttxg0d5qwzze4j45rfd9hrvqlxp8c8n2tkwar",
				"initialAmount": 7825086956521739000,
				"unlockSchedule": [
					{
						"amount": 10000000000000000,
						"locktime": 1633824000
					}
				]
			  },
			  {
				"ethAddr": "0xbfd3b345b3ce93e57f0d50181493af6de49d5602",
				"axcAddr": "Swap-axc1h25zmmwgazncpjrhchv9k3wnpv9u724end7wa0",
				"initialAmount": 7825086956521739000,
				"unlockSchedule": [
					{
						"amount": 10000000000000000,
						"locktime": 1633824000
					}
				]
			  },
			  {
				"ethAddr": "0x62725809c283537b2823cd769ed8ce87681d9532",
				"axcAddr": "Swap-axc14p4gscefdj6gy7kdp6t2gf7zfkfvhpgyp735zz",
				"initialAmount": 7825086956521739000,
				"unlockSchedule": [
					{
						"amount": 10000000000000000,
						"locktime": 1633824000
					}
				]
			  },
			  {
				"ethAddr": "0x7a8fd85207720f8d7f5b85bdab5ae98fb874e71a",
				"axcAddr": "Swap-axc1amjdxxunua7q84xptmhnlcvwuul9cv0ldqvspq",
				"initialAmount": 7825086956521739000,
				"unlockSchedule": [
					{
						"amount": 10000000000000000,
						"locktime": 1633824000
					}
				]
			  },
			  {
				"ethAddr": "0xd52161eddc0d06399c7a81416b60ab6e8c33a214",
				"axcAddr": "Swap-axc19z94hszzphc0vnvgu8vmzdu9zrtlckht2v9es2",
				"initialAmount": 7825086956521739000,
				"unlockSchedule": [
					{
						"amount": 10000000000000000,
						"locktime": 1633824000
					}
				]
			  },
			  {
				"ethAddr": "0x1949d58756d912a3adc51a67346ac74bc30fe448",
				"axcAddr": "Swap-axc1hq3y32q364mwpanjc6mw9v5wr9ky6fmc0q3x6h",
				"initialAmount": 7825086956521739000,
				"unlockSchedule": [
					{
						"amount": 10000000000000000,
						"locktime": 1633824000
					}
				]
			  },
			  {
				"ethAddr": "0xc0fe574812202cb697f6568e72ee1e2dae7ff504",
				"axcAddr": "Swap-axc1uzy97sancj3lfzwam4earj88jwfue2u0d3suah",
				"initialAmount": 7825086956521739000,
				"unlockSchedule": [
					{
						"amount": 10000000000000000,
						"locktime": 1633824000
					}
				]
			  },
			  {
				"ethAddr": "0xc10f293e5019b2c845111cb5989651a5f514afac",
				"axcAddr": "Swap-axc1ds5gsn6wjrqzytn0sgnfmg5fmn5nfphhdlvg7u",
				"initialAmount": 7825086956521739000,
				"unlockSchedule": [
					{
						"amount": 10000000000000000,
						"locktime": 1633824000
					}
				]
			  },
			  {
				"ethAddr": "0xe2fbccbc351cd37dd864d1464cac9631429a70eb",
				"axcAddr": "Swap-axc1dw58c8wq9hl9rjw0jqy748lcmq0uqem7ayv0ex",
				"initialAmount": 7825086956521739000,
				"unlockSchedule": [
					{
						"amount": 10000000000000000,
						"locktime": 1633824000
					}
				]
			  },
			  {
				"ethAddr": "0xa6914310373bd688356dc02741eaa4bca862496f",
				"axcAddr": "Swap-axc1ca83yv97gwdm8epc44keve49vjsm2w04qgd35s",
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
			"Swap-axc1lgr4vjsuy5lcvg2660t5ufqp0yfv9rkv0yqkd9",
			"Swap-axc194dsz5fxexgkx93jwj5a8zqjcw6pcnfzlhg95f",
			"Swap-axc1hz268ehr8ur46eu8gwzrczc5sxgz0ww7qt5tt3",
			"Swap-axc1rm97wskg54pljccf5tc6xykhgk76z582n9rg54",
			"Swap-axc1lypz8hpz9dml0jzmxmqvg5kl49pc0kepk478x3",
			"Swap-axc195ul7h2c0e28lnmz3nv26dukk3pkqayf888jx2",
			"Swap-axc1rnta60c0vvdhxly9rgf9g2atsh803lvfa2y4z8",
			"Swap-axc1sx3hm9zjcx76ecd4fyej085nrkz2w4hfq42pa2",
			"Swap-axc1pp7cvggn2zflzxvur9vys3ywrwh5wt2xjqy483",
			"Swap-axc1juwkcwtrulkkkf7p5zvlu4ruznxmlxu4vnz2je",
			"Swap-axc1dn0f8tz9de7v9nrqjel4qxuadswftc8tmde8a4",
			"Swap-axc14dlest7mg83yggqxa425te4lfx6hzpk4dpz2ad",
			"Swap-axc1a2hnt973ar2wntw8mtn7nyczeqwzlqtzydj4kr",
			"Swap-axc1qttxg0d5qwzze4j45rfd9hrvqlxp8c8n2tkwar",
			"Swap-axc1h25zmmwgazncpjrhchv9k3wnpv9u724end7wa0",
			"Swap-axc14p4gscefdj6gy7kdp6t2gf7zfkfvhpgyp735zz",
			"Swap-axc1amjdxxunua7q84xptmhnlcvwuul9cv0ldqvspq",
			"Swap-axc19z94hszzphc0vnvgu8vmzdu9zrtlckht2v9es2",
			"Swap-axc1hq3y32q364mwpanjc6mw9v5wr9ky6fmc0q3x6h",
			"Swap-axc1uzy97sancj3lfzwam4earj88jwfue2u0d3suah",
			"Swap-axc1ds5gsn6wjrqzytn0sgnfmg5fmn5nfphhdlvg7u",
			"Swap-axc1dw58c8wq9hl9rjw0jqy748lcmq0uqem7ayv0ex",
			"Swap-axc1ca83yv97gwdm8epc44keve49vjsm2w04qgd35s"
		],
		"initialStakers": [
			{
				"nodeID": "NodeID-FVJiVBMXg69Eys8v3K3fUZXM21fW5k7Ts",
				"rewardAddress": "Swap-axc1lgr4vjsuy5lcvg2660t5ufqp0yfv9rkv0yqkd9",
				"nominationFee": 200000
			  },
			  {
				"nodeID": "NodeID-8ofNV8sb3J9rwaC3HuXcPFAkNEqF8oJzV",
				"rewardAddress": "Swap-axc194dsz5fxexgkx93jwj5a8zqjcw6pcnfzlhg95f",
				"nominationFee": 200000
			  },
			  {
				"nodeID": "NodeID-BLaJAbw7aYCmZtYoemCBpowijcBPq7yZs",
				"rewardAddress": "Swap-axc1hz268ehr8ur46eu8gwzrczc5sxgz0ww7qt5tt3",
				"nominationFee": 200000
			  },
			  {
				"nodeID": "NodeID-FTirsWrt7FJVTTvQejEQpJERuJDKk9BW2",
				"rewardAddress": "Swap-axc1rm97wskg54pljccf5tc6xykhgk76z582n9rg54",
				"nominationFee": 200000
			  },
			  {
				"nodeID": "NodeID-Gym2UoBo3b4tp3pukZzDjxJefNc5XE8Cw",
				"rewardAddress": "Swap-axc1lypz8hpz9dml0jzmxmqvg5kl49pc0kepk478x3",
				"nominationFee": 200000
			  },
			  {
				"nodeID": "NodeID-AkZUoGaJUWrfGUBtk7Wrsyvu5DZKVkxKQ",
				"rewardAddress": "Swap-axc195ul7h2c0e28lnmz3nv26dukk3pkqayf888jx2",
				"nominationFee": 200000
			  },
			  {
				"nodeID": "NodeID-FuzyvQMfFRhAWiiQLvcDVErbNoA2czuYp",
				"rewardAddress": "Swap-axc1rnta60c0vvdhxly9rgf9g2atsh803lvfa2y4z8",
				"nominationFee": 200000
			  },
			  {
				"nodeID": "NodeID-Ag6dycckf8nwo2EYrvfnuw4oAw8xTz6P3",
				"rewardAddress": "Swap-axc1sx3hm9zjcx76ecd4fyej085nrkz2w4hfq42pa2",
				"nominationFee": 200000
			  },
			  {
				"nodeID": "NodeID-PUWVULUUXozYBvUQwoPxQFMXT13MqRBAK",
				"rewardAddress": "Swap-axc1pp7cvggn2zflzxvur9vys3ywrwh5wt2xjqy483",
				"nominationFee": 200000
			  },
			  {
				"nodeID": "NodeID-MbjtKYV1EEbRdaK1d396g54VrVnTXet5Q",
				"rewardAddress": "Swap-axc1juwkcwtrulkkkf7p5zvlu4ruznxmlxu4vnz2je",
				"nominationFee": 200000
			  },
			  {
				"nodeID": "NodeID-FeFgdAG3v8YJidF4xR8E3zSHUEuNRtBfB",
				"rewardAddress": "Swap-axc1dn0f8tz9de7v9nrqjel4qxuadswftc8tmde8a4",
				"nominationFee": 200000
			  },
			  {
				"nodeID": "NodeID-2eAGdTFKwdi63NDvADr99qoDANmY11uun",
				"rewardAddress": "Swap-axc14dlest7mg83yggqxa425te4lfx6hzpk4dpz2ad",
				"nominationFee": 200000
			  },
			  {
				"nodeID": "NodeID-pbDBA6Gfy9ShgscRm83kmJpVECqyh2EV",
				"rewardAddress": "Swap-axc1a2hnt973ar2wntw8mtn7nyczeqwzlqtzydj4kr",
				"nominationFee": 200000
			  },
			  {
				"nodeID": "NodeID-KbQdsqnna4KwQfsRksNMah1ZbpQgvqtrX",
				"rewardAddress": "Swap-axc1qttxg0d5qwzze4j45rfd9hrvqlxp8c8n2tkwar",
				"nominationFee": 200000
			  },
			  {
				"nodeID": "NodeID-Ne4NyLV1WAyyXoRAzSSSTbc7pU8jbFTez",
				"rewardAddress": "Swap-axc1h25zmmwgazncpjrhchv9k3wnpv9u724end7wa0",
				"nominationFee": 200000
			  },
			  {
				"nodeID": "NodeID-Mx3reKpoVipQVUzH41qx5gJZmNURgvFJD",
				"rewardAddress": "Swap-axc14p4gscefdj6gy7kdp6t2gf7zfkfvhpgyp735zz",
				"nominationFee": 200000
			  },
			  {
				"nodeID": "NodeID-CksSwGj4tSAyyRVrFavs6p9farwwtEC1b",
				"rewardAddress": "Swap-axc1amjdxxunua7q84xptmhnlcvwuul9cv0ldqvspq",
				"nominationFee": 200000
			  },
			  {
				"nodeID": "NodeID-QHHAyHWvLEJsawMTYgZfAT66erXRVqtHu",
				"rewardAddress": "Swap-axc19z94hszzphc0vnvgu8vmzdu9zrtlckht2v9es2",
				"nominationFee": 200000
			  },
			  {
				"nodeID": "NodeID-21zDZTDjnKcLE8yMy1SXsC5oYvJq1m1Ti",
				"rewardAddress": "Swap-axc1hq3y32q364mwpanjc6mw9v5wr9ky6fmc0q3x6h",
				"nominationFee": 200000
			  },
			  {
				"nodeID": "NodeID-Q1KAsFkpTVszwbAtANydsRTLj9jjs8PAP",
				"rewardAddress": "Swap-axc1uzy97sancj3lfzwam4earj88jwfue2u0d3suah",
				"nominationFee": 200000
			  },
			  {
				"nodeID": "NodeID-HPz5qwnbvsiz4oqmsf7zr8NeWsmbSnAep",
				"rewardAddress": "Swap-axc1ds5gsn6wjrqzytn0sgnfmg5fmn5nfphhdlvg7u",
				"nominationFee": 200000
			  },
			  {
				"nodeID": "NodeID-G3FjbZ7rSrBqL11BpZG89RJPXkP2cUaiL",
				"rewardAddress": "Swap-axc1dw58c8wq9hl9rjw0jqy748lcmq0uqem7ayv0ex",
				"nominationFee": 200000
			  },
			  {
				"nodeID": "NodeID-5wFDMxCVVpvZip4oUrMyunwRde1aU9Vmb",
				"rewardAddress": "Swap-axc1ca83yv97gwdm8epc44keve49vjsm2w04qgd35s",
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
