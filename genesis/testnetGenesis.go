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
				"ethAddr": "0xc185c8478f60acf3f9b3a9ed1b83284c7498e751",
				"axcAddr": "Swap-test182z7vwl3kqme3rng90uls66gtnu95fswuw40r8",
				"initialAmount": 7825086956521739000,
				"unlockSchedule": [
							{
							  "amount": 10000000000000000,
							  "locktime": 1633824000
							}
						  ]
			},
			{
				"ethAddr": "0xf28723a02b06a64d196ea2432f17b96b61c1129a",
				"axcAddr": "Swap-test12p6ccwcda4nma7h6yks88fewa05sqwv5j9wy08",
				"initialAmount": 7825086956521739000,
				"unlockSchedule": [
							{
							  "amount": 10000000000000000,
							  "locktime": 1633824000
							}
						  ]
			},
			{
				"ethAddr": "0xbfe147a3df764a8be03e5994030ea90e6ba481ce",
				"axcAddr": "Swap-test1guw3ghckdursxpatkv820xzwklmtdw23djj9sv",
				"initialAmount": 7825086956521739000,
				"unlockSchedule": [
							{
							  "amount": 10000000000000000,
							  "locktime": 1633824000
							}
						  ]
			},
			{
				"ethAddr": "0x2484d46bafd67ffb0d8a9111cc4c5456c7831c33",
				"axcAddr": "Swap-test1dwae37gzt9epvfqxtw5z7tyhe5p9pjscj8zwvf",
				"initialAmount": 7825086956521739000,
				"unlockSchedule": [
							{
							  "amount": 10000000000000000,
							  "locktime": 1633824000
							}
						  ]
			},
			{
				"ethAddr": "0xf511a6d4d3f7f130d3b06ad21f3e69f1aad9a7bf",
				"axcAddr": "Swap-test1z765wgqr4nhmw8sl0wlyvurkefwzj2p2ggvz9d",
				"initialAmount": 7825086956521739000,
				"unlockSchedule": [
							{
							  "amount": 10000000000000000,
							  "locktime": 1633824000
							}
						  ]
			},
			{
				"ethAddr": "0xdd1595ac09d1aa4127ad1a34a74b830f8655a5df",
				"axcAddr": "Swap-test17w7usncg6wm0s6vlwhcft8mzd359eh0je4g4dz",
				"initialAmount": 7825086956521739000,
				"unlockSchedule": [
							{
							  "amount": 10000000000000000,
							  "locktime": 1633824000
							}
						  ]
			},
			{
				"ethAddr": "0xff898744b24f63bafd9828a4a9fb151035eddd08",
				"axcAddr": "Swap-test1d8jls2k2qmceektvrq2xn6j4spjkg9k99wf7gg",
				"initialAmount": 7825086956521739000,
				"unlockSchedule": [
							{
							  "amount": 10000000000000000,
							  "locktime": 1633824000
							}
						  ]
			},
			{
				"ethAddr": "0xac19d1fbb9ad3868bc9b244d4a7e2ff2ec160322",
				"axcAddr": "Swap-test1syxx9njyc7yc3h8npqny9sq4h9a7fc3q9puwyk",
				"initialAmount": 7825086956521739000,
				"unlockSchedule": [
							{
							  "amount": 10000000000000000,
							  "locktime": 1633824000
							}
						  ]
			},
			{
				"ethAddr": "0x539a62ad0f8c6174fd304582fa8fbab063bda03e",
				"axcAddr": "Swap-test102ucvg4gmee23ncg2lygfvtmvza208pn3cy7tv",
				"initialAmount": 7825086956521739000,
				"unlockSchedule": [
							{
							  "amount": 10000000000000000,
							  "locktime": 1633824000
							}
						  ]
			},
			{
				"ethAddr": "0x6a64052efac875715404a2c4c1ff045f4f894128",
				"axcAddr": "Swap-test1uuyutp87sev9svh7jr6u6gc0tsfp5qa7k6yf7c",
				"initialAmount": 7825086956521739000,
				"unlockSchedule": [
							{
							  "amount": 10000000000000000,
							  "locktime": 1633824000
							}
						  ]
			},
			{
				"ethAddr": "0x074d22b601aeedca6b918912c22ec0324e929e31",
				"axcAddr": "Swap-test1t03a66tq3su2qrgpskgfac3xec80srllta7w7x",
				"initialAmount": 7825086956521739000,
				"unlockSchedule": [
							{
							  "amount": 10000000000000000,
							  "locktime": 1633824000
							}
						  ]
			},
			{
				"ethAddr": "0xa953dcc9b2f4c510861101435a5f71f8ebcc9dbf",
				"axcAddr": "Swap-test1ac627c52gvgz6v9rh5v7uju20rtz42qw4y3fkr",
				"initialAmount": 7825086956521739000,
				"unlockSchedule": [
							{
							  "amount": 10000000000000000,
							  "locktime": 1633824000
							}
						  ]
			},
			{
				"ethAddr": "0x97554d4a9be8a170030137062532b40a8a1c85b4",
				"axcAddr": "Swap-test16rtd7g2eeftkrr09dywfhpt72k46tfy93hux3n",
				"initialAmount": 7825086956521739000,
				"unlockSchedule": [
							{
							  "amount": 10000000000000000,
							  "locktime": 1633824000
							}
						  ]
			},
			{
				"ethAddr": "0xf3644354d3736f7dc9e346e9a49ca9aac8259f3b",
				"axcAddr": "Swap-test1yfeu6mmuje9l7fnfapr7v50xwehhyr250ludts",
				"initialAmount": 7825086956521739000,
				"unlockSchedule": [
							{
							  "amount": 10000000000000000,
							  "locktime": 1633824000
							}
						  ]
			},
			{
				"ethAddr": "0x5e1073ad14b87e1690b677105e3017ac1d224c97",
				"axcAddr": "Swap-test1nyykwtyrj3p3sq7lpwqu9prwthflts78rrw8ms",
				"initialAmount": 7825086956521739000,
				"unlockSchedule": [
							{
							  "amount": 10000000000000000,
							  "locktime": 1633824000
							}
						  ]
			},
			{
				"ethAddr": "0x8a07c956b599d1554434d9c0568a03d2b3318c97",
				"axcAddr": "Swap-test1c5x9krk47awchtz4kxpgrhq4w55498gjkhw688",
				"initialAmount": 7825086956521739000,
				"unlockSchedule": [
							{
							  "amount": 10000000000000000,
							  "locktime": 1633824000
							}
						  ]
			},
			{
				"ethAddr": "0x78c7ae1be01a8e5068de8cc201af3f22e386129a",
				"axcAddr": "Swap-test1l23wjsxmcw83wsj72xshu2cc3lvesl3vpyf433",
				"initialAmount": 7825086956521739000,
				"unlockSchedule": [
							{
							  "amount": 10000000000000000,
							  "locktime": 1633824000
							}
						  ]
			},
			{
				"ethAddr": "0x208c6a0f0a682a4ed44e43320f5d8feebd74ec2e",
				"axcAddr": "Swap-test1cpqhcmvdcsykhgncstlwsqxsnjgpkmn6nqls52",
				"initialAmount": 7825086956521739000,
				"unlockSchedule": [
							{
							  "amount": 10000000000000000,
							  "locktime": 1633824000
							}
						  ]
			},
			{
				"ethAddr": "0x18ba6590cdf76f26951197530b37433d95ce60b1",
				"axcAddr": "Swap-test1vv6u5e6jyk0y5t3wh52tn0qmyk9egye09wud68",
				"initialAmount": 7825086956521739000,
				"unlockSchedule": [
							{
							  "amount": 10000000000000000,
							  "locktime": 1633824000
							}
						  ]
			},
			{
				"ethAddr": "0xb43ab7e2b4dad9a4f43cb274d9cbee8a88319e24",
				"axcAddr": "Swap-test1c77dv0fuc7e4wmq9a4a3lav3elpe6ns4unzc50",
				"initialAmount": 7825086956521739000,
				"unlockSchedule": [
							{
							  "amount": 10000000000000000,
							  "locktime": 1633824000
							}
						  ]
			}
		],
		"startTime": 1630987200,
		"initialStakeDuration": 31536000,
		"initialStakeDurationOffset": 5400,
		"initialStakedFunds": [
			"Swap-test182z7vwl3kqme3rng90uls66gtnu95fswuw40r8",
			"Swap-test12p6ccwcda4nma7h6yks88fewa05sqwv5j9wy08",
			"Swap-test1guw3ghckdursxpatkv820xzwklmtdw23djj9sv",
			"Swap-test1dwae37gzt9epvfqxtw5z7tyhe5p9pjscj8zwvf",
			"Swap-test1z765wgqr4nhmw8sl0wlyvurkefwzj2p2ggvz9d",
			"Swap-test17w7usncg6wm0s6vlwhcft8mzd359eh0je4g4dz",
			"Swap-test1d8jls2k2qmceektvrq2xn6j4spjkg9k99wf7gg",
			"Swap-test1syxx9njyc7yc3h8npqny9sq4h9a7fc3q9puwyk",
			"Swap-test102ucvg4gmee23ncg2lygfvtmvza208pn3cy7tv",
			"Swap-test1uuyutp87sev9svh7jr6u6gc0tsfp5qa7k6yf7c",
			"Swap-test1t03a66tq3su2qrgpskgfac3xec80srllta7w7x",
			"Swap-test1ac627c52gvgz6v9rh5v7uju20rtz42qw4y3fkr",
			"Swap-test16rtd7g2eeftkrr09dywfhpt72k46tfy93hux3n",
			"Swap-test1yfeu6mmuje9l7fnfapr7v50xwehhyr250ludts",
			"Swap-test1nyykwtyrj3p3sq7lpwqu9prwthflts78rrw8ms",
			"Swap-test1c5x9krk47awchtz4kxpgrhq4w55498gjkhw688",
			"Swap-test1l23wjsxmcw83wsj72xshu2cc3lvesl3vpyf433",
			"Swap-test1cpqhcmvdcsykhgncstlwsqxsnjgpkmn6nqls52",
			"Swap-test1vv6u5e6jyk0y5t3wh52tn0qmyk9egye09wud68",
			"Swap-test1c77dv0fuc7e4wmq9a4a3lav3elpe6ns4unzc50"
		],
		"initialStakers": [
			{
				"nodeID": "NodeID-FVJiVBMXg69Eys8v3K3fUZXM21fW5k7Ts",
				"rewardAddress": "Swap-test182z7vwl3kqme3rng90uls66gtnu95fswuw40r8",
				"nominationFee": 200000
			},
			{
				"nodeID": "NodeID-8ofNV8sb3J9rwaC3HuXcPFAkNEqF8oJzV",
				"rewardAddress": "Swap-test12p6ccwcda4nma7h6yks88fewa05sqwv5j9wy08",
				"nominationFee": 200000
			},
			{
				"nodeID": "NodeID-BLaJAbw7aYCmZtYoemCBpowijcBPq7yZs",
				"rewardAddress": "Swap-test1guw3ghckdursxpatkv820xzwklmtdw23djj9sv",
				"nominationFee": 200000
			},
			{
				"nodeID": "NodeID-FTirsWrt7FJVTTvQejEQpJERuJDKk9BW2",
				"rewardAddress": "Swap-test1dwae37gzt9epvfqxtw5z7tyhe5p9pjscj8zwvf",
				"nominationFee": 200000
			},
			{
				"nodeID": "NodeID-Gym2UoBo3b4tp3pukZzDjxJefNc5XE8Cw",
				"rewardAddress": "Swap-test1z765wgqr4nhmw8sl0wlyvurkefwzj2p2ggvz9d",
				"nominationFee": 200000
			},
			{
				"nodeID": "NodeID-AkZUoGaJUWrfGUBtk7Wrsyvu5DZKVkxKQ",
				"rewardAddress": "Swap-test17w7usncg6wm0s6vlwhcft8mzd359eh0je4g4dz",
				"nominationFee": 200000
			},
			{
				"nodeID": "NodeID-FuzyvQMfFRhAWiiQLvcDVErbNoA2czuYp",
				"rewardAddress": "Swap-test1d8jls2k2qmceektvrq2xn6j4spjkg9k99wf7gg",
				"nominationFee": 200000
			},
			{
				"nodeID": "NodeID-Ag6dycckf8nwo2EYrvfnuw4oAw8xTz6P3",
				"rewardAddress": "Swap-test1syxx9njyc7yc3h8npqny9sq4h9a7fc3q9puwyk",
				"nominationFee": 200000
			},
			{
				"nodeID": "NodeID-PUWVULUUXozYBvUQwoPxQFMXT13MqRBAK",
				"rewardAddress": "Swap-test102ucvg4gmee23ncg2lygfvtmvza208pn3cy7tv",
				"nominationFee": 200000
			},
			{
				"nodeID": "NodeID-MbjtKYV1EEbRdaK1d396g54VrVnTXet5Q",
				"rewardAddress": "Swap-test1uuyutp87sev9svh7jr6u6gc0tsfp5qa7k6yf7c",
				"nominationFee": 200000
			},
			{
				"nodeID": "NodeID-FeFgdAG3v8YJidF4xR8E3zSHUEuNRtBfB",
				"rewardAddress": "Swap-test1t03a66tq3su2qrgpskgfac3xec80srllta7w7x",
				"nominationFee": 200000
			},
			{
				"nodeID": "NodeID-2eAGdTFKwdi63NDvADr99qoDANmY11uun",
				"rewardAddress": "Swap-test1ac627c52gvgz6v9rh5v7uju20rtz42qw4y3fkr",
				"nominationFee": 200000
			},
			{
				"nodeID": "NodeID-pbDBA6Gfy9ShgscRm83kmJpVECqyh2EV",
				"rewardAddress": "Swap-test16rtd7g2eeftkrr09dywfhpt72k46tfy93hux3n",
				"nominationFee": 200000
			},
			{
				"nodeID": "NodeID-KbQdsqnna4KwQfsRksNMah1ZbpQgvqtrX",
				"rewardAddress": "Swap-test1yfeu6mmuje9l7fnfapr7v50xwehhyr250ludts",
				"nominationFee": 200000
			},
			{
				"nodeID": "NodeID-Ne4NyLV1WAyyXoRAzSSSTbc7pU8jbFTez",
				"rewardAddress": "Swap-test1nyykwtyrj3p3sq7lpwqu9prwthflts78rrw8ms",
				"nominationFee": 200000
			},
			{
				"nodeID": "NodeID-Mx3reKpoVipQVUzH41qx5gJZmNURgvFJD",
				"rewardAddress": "Swap-test1c5x9krk47awchtz4kxpgrhq4w55498gjkhw688",
				"nominationFee": 200000
			},
			{
				"nodeID": "NodeID-CksSwGj4tSAyyRVrFavs6p9farwwtEC1b",
				"rewardAddress": "Swap-test1l23wjsxmcw83wsj72xshu2cc3lvesl3vpyf433",
				"nominationFee": 200000
			},
			{
				"nodeID": "NodeID-QHHAyHWvLEJsawMTYgZfAT66erXRVqtHu",
				"rewardAddress": "Swap-test1cpqhcmvdcsykhgncstlwsqxsnjgpkmn6nqls52",
				"nominationFee": 200000
			},
			{
				"nodeID": "NodeID-21zDZTDjnKcLE8yMy1SXsC5oYvJq1m1Ti",
				"rewardAddress": "Swap-test1vv6u5e6jyk0y5t3wh52tn0qmyk9egye09wud68",
				"nominationFee": 200000
			},
			{
				"nodeID": "NodeID-Q1KAsFkpTVszwbAtANydsRTLj9jjs8PAP",
				"rewardAddress": "Swap-test1c77dv0fuc7e4wmq9a4a3lav3elpe6ns4unzc50",
				"nominationFee": 200000
			}
		],
		"axChainGenesis": "{\"config\":{\"chainId\":4000,\"homesteadBlock\":0,\"daoForkBlock\":0,\"daoForkSupport\":true,\"eip150Block\":0,\"eip150Hash\":\"0x2086799aeebeae135c246c65021c82b4e15a2c451340993aacfd2751886514f0\",\"eip155Block\":0,\"eip158Block\":0,\"byzantiumBlock\":0,\"constantinopleBlock\":0,\"petersburgBlock\":0,\"istanbulBlock\":0,\"muirGlacierBlock\":0,\"apricotPhase1BlockTimestamp\":0,\"apricotPhase2BlockTimestamp\":0},\"nonce\":\"0x0\",\"timestamp\":\"0x0\",\"extraData\":\"0x00\",\"gasLimit\":\"0x7A1200\",\"minBaseFee\":1000000,\"difficulty\":\"0x0\",\"mixHash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\",\"coinbase\":\"0x0000000000000000000000000000000000000000\",\"alloc\":{\"790e5825b65ade90095ddfe30709f14f04372194\":{\"balance\":\"0xD3C21BCECCEDA1000000\"},\"07464552eae4efb8b985d311b093d25de1d69f91\":{\"balance\":\"0xD3C21BCECCEDA1000000\"},\"6fac93cbe0263cbabf9b95ce415fcdc5829713c4\":{\"balance\":\"0xD3C21BCECCEDA1000000\"},\"d05b8c7da18b39ef013dd490181d337d968a9bb0\":{\"balance\":\"0xD3C21BCECCEDA1000000\"},\"0f9ebc77d4681c2c3242fdf687806a3a1aff3422\":{\"balance\":\"0xD3C21BCECCEDA1000000\"}},\"number\":\"0x0\",\"gasUsed\":\"0x0\",\"parentHash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\"}",
		"message": "{{ fun_quote }}"
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
