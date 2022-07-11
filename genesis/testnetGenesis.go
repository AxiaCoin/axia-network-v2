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
				"axcAddr": "Swap-test1dssj9m0xtvpacypu2u536rjs39wte799zpmgk7",
				"initialAmount": 499000000000000000,
				"unlockSchedule": [
							{
							  "amount": 1000000000000000,
							  "locktime": 1688816151
							}
						  ]
			},
			{
				"ethAddr": "0xf28723a02b06a64d196ea2432f17b96b61c1129a",
				"axcAddr": "Swap-test10cgn6cknsyvp0n6afwrzmvp93maesf4c9fdn8z",
				"initialAmount": 499000000000000000,
				"unlockSchedule": [
							{
							  "amount": 1000000000000000,
							  "locktime": 1688816151
							}
						  ]
			},
			{
				"ethAddr": "0xbfe147a3df764a8be03e5994030ea90e6ba481ce",
				"axcAddr": "Swap-test102u6y4juspq0e0wgqy62pdgd3579asn4nfm8zh",
				"initialAmount": 499000000000000000,
				"unlockSchedule": [
							{
							  "amount": 1000000000000000,
							  "locktime": 1688816151
							}
						  ]
			},
			{
				"ethAddr": "0x2484d46bafd67ffb0d8a9111cc4c5456c7831c33",
				"axcAddr": "Swap-test16wy9v23ay2w070c0ppgm35gue755hewsnv4npf",
				"initialAmount": 499000000000000000,
				"unlockSchedule": [
							{
							  "amount": 1000000000000000,
							  "locktime": 1688816151
							}
						  ]
			},
			{
				"ethAddr": "0xf511a6d4d3f7f130d3b06ad21f3e69f1aad9a7bf",
				"axcAddr": "Swap-test156uyw8uy3c4cldv4gfmy6c5p0j79q3z0762jl6",
				"initialAmount": 499000000000000000,
				"unlockSchedule": [
							{
							  "amount": 1000000000000000,
							  "locktime": 1688816151
							}
						  ]
			},
			{
				"ethAddr": "0xdd1595ac09d1aa4127ad1a34a74b830f8655a5df",
				"axcAddr": "Swap-test1es639wwwxpzectuvdwlstc473qa6m3qrzuf3rk",
				"initialAmount": 499000000000000000,
				"unlockSchedule": [
							{
							  "amount": 1000000000000000,
							  "locktime": 1688816151
							}
						  ]
			},
			{
				"ethAddr": "0xff898744b24f63bafd9828a4a9fb151035eddd08",
				"axcAddr": "Swap-test1jzc63t9hxcesljtl23efqv4dq6ehny5x4nwm0w",
				"initialAmount": 499000000000000000,
				"unlockSchedule": [
							{
							  "amount": 1000000000000000,
							  "locktime": 1688816151
							}
						  ]
			},
			{
				"ethAddr": "0xac19d1fbb9ad3868bc9b244d4a7e2ff2ec160322",
				"axcAddr": "Swap-test16uvswz0rdj2c5ecpq6s9zeccgpq0w8lwcw4xda",
				"initialAmount": 499000000000000000,
				"unlockSchedule": [
							{
							  "amount": 1000000000000000,
							  "locktime": 1688816151
							}
						  ]
			},
			{
				"ethAddr": "0x539a62ad0f8c6174fd304582fa8fbab063bda03e",
				"axcAddr": "Swap-test1s5n5xq0uamylwldrky9ln4fu0cus8mt6mq4ayy",
				"initialAmount": 499000000000000000,
				"unlockSchedule": [
							{
							  "amount": 1000000000000000,
							  "locktime": 1688816151
							}
						  ]
			},
			{
				"ethAddr": "0x6a64052efac875715404a2c4c1ff045f4f894128",
				"axcAddr": "Swap-test1asg2ukxyhr45yyc7d4vzrwun9xfsl0zza30nlu",
				"initialAmount": 499000000000000000,
				"unlockSchedule": [
							{
							  "amount": 1000000000000000,
							  "locktime": 1688816151
							}
						  ]
			},
			{
				"ethAddr": "0x074d22b601aeedca6b918912c22ec0324e929e31",
				"axcAddr": "Swap-test1tezw3pagmmjfjz53k8kt0ymmp24zq9ewvc4ey8",
				"initialAmount": 499000000000000000,
				"unlockSchedule": [
							{
							  "amount": 1000000000000000,
							  "locktime": 1688816151
							}
						  ]
			},
			{
				"ethAddr": "0xa953dcc9b2f4c510861101435a5f71f8ebcc9dbf",
				"axcAddr": "Swap-test1mmdwn3n2a4ft524ggzw8jnvva92jd6wjxnjg4m",
				"initialAmount": 499000000000000000,
				"unlockSchedule": [
							{
							  "amount": 1000000000000000,
							  "locktime": 1688816151
							}
						  ]
			},
			{
				"ethAddr": "0x97554d4a9be8a170030137062532b40a8a1c85b4",
				"axcAddr": "Swap-test1nqvz6agrfxwuyj6skg0j6zwxcnqrun6mz8zvkj",
				"initialAmount": 499000000000000000,
				"unlockSchedule": [
							{
							  "amount": 1000000000000000,
							  "locktime": 1688816151
							}
						  ]
			},
			{
				"ethAddr": "0xf3644354d3736f7dc9e346e9a49ca9aac8259f3b",
				"axcAddr": "Swap-test1fmrn33tlxpvmrxcppddvh4tdvr7560w0c8lcry",
				"initialAmount": 499000000000000000,
				"unlockSchedule": [
							{
							  "amount": 1000000000000000,
							  "locktime": 1688816151
							}
						  ]
			},
			{
				"ethAddr": "0x5e1073ad14b87e1690b677105e3017ac1d224c97",
				"axcAddr": "Swap-test1q6xuesh0qe08eeadk262hge2xrg6e0pxelzuw8",
				"initialAmount": 499000000000000000,
				"unlockSchedule": [
							{
							  "amount": 1000000000000000,
							  "locktime": 1688816151
							}
						  ]
			},
			{
				"ethAddr": "0x8a07c956b599d1554434d9c0568a03d2b3318c97",
				"axcAddr": "Swap-test1kkepg8m6f49a6k66vqrs84cqretj6pvctnw7f3",
				"initialAmount": 499000000000000000,
				"unlockSchedule": [
							{
							  "amount": 1000000000000000,
							  "locktime": 1688816151
							}
						  ]
			},
			{
				"ethAddr": "0x78c7ae1be01a8e5068de8cc201af3f22e386129a",
				"axcAddr": "Swap-test1fslvqs9rduee2wjzyrhkfztff36kxmjt6a9urc",
				"initialAmount": 499000000000000000,
				"unlockSchedule": [
							{
							  "amount": 1000000000000000,
							  "locktime": 1688816151
							}
						  ]
			},
			{
				"ethAddr": "0x208c6a0f0a682a4ed44e43320f5d8feebd74ec2e",
				"axcAddr": "Swap-test17n9p4egxfjz7heu065ezvnxckdd3fdt6cfpy6d",
				"initialAmount": 499000000000000000,
				"unlockSchedule": [
							{
							  "amount": 1000000000000000,
							  "locktime": 1688816151
							}
						  ]
			},
			{
				"ethAddr": "0x18ba6590cdf76f26951197530b37433d95ce60b1",
				"axcAddr": "Swap-test1356g6fcq7r5v96ld3mqpntm0su0j80ppdm4r80",
				"initialAmount": 499000000000000000,
				"unlockSchedule": [
							{
							  "amount": 1000000000000000,
							  "locktime": 1688816151
							}
						  ]
			},
			{
				"ethAddr": "0xb43ab7e2b4dad9a4f43cb274d9cbee8a88319e24",
				"axcAddr": "Swap-test1vvlcdwcwlu5xs9qs0scn2t8gu3rsp0vtz06kjg",
				"initialAmount": 499000000000000000,
				"unlockSchedule": [
							{
							  "amount": 1000000000000000,
							  "locktime": 1688816151
							}
						  ]
			}
		],
		"startTime": 1657179865,
		"initialStakeDuration": 31536000,
		"initialStakeDurationOffset": 5400,
		"initialStakedFunds": [
			"Swap-test1dssj9m0xtvpacypu2u536rjs39wte799zpmgk7",
			"Swap-test10cgn6cknsyvp0n6afwrzmvp93maesf4c9fdn8z",
			"Swap-test102u6y4juspq0e0wgqy62pdgd3579asn4nfm8zh",
			"Swap-test16wy9v23ay2w070c0ppgm35gue755hewsnv4npf",
			"Swap-test156uyw8uy3c4cldv4gfmy6c5p0j79q3z0762jl6",
			"Swap-test1es639wwwxpzectuvdwlstc473qa6m3qrzuf3rk",
			"Swap-test1jzc63t9hxcesljtl23efqv4dq6ehny5x4nwm0w",
			"Swap-test16uvswz0rdj2c5ecpq6s9zeccgpq0w8lwcw4xda",
			"Swap-test1s5n5xq0uamylwldrky9ln4fu0cus8mt6mq4ayy",
			"Swap-test1asg2ukxyhr45yyc7d4vzrwun9xfsl0zza30nlu",
			"Swap-test1tezw3pagmmjfjz53k8kt0ymmp24zq9ewvc4ey8",
			"Swap-test1mmdwn3n2a4ft524ggzw8jnvva92jd6wjxnjg4m",
			"Swap-test1nqvz6agrfxwuyj6skg0j6zwxcnqrun6mz8zvkj",
			"Swap-test1fmrn33tlxpvmrxcppddvh4tdvr7560w0c8lcry",
			"Swap-test1q6xuesh0qe08eeadk262hge2xrg6e0pxelzuw8",
			"Swap-test1kkepg8m6f49a6k66vqrs84cqretj6pvctnw7f3",
			"Swap-test1fslvqs9rduee2wjzyrhkfztff36kxmjt6a9urc",
			"Swap-test17n9p4egxfjz7heu065ezvnxckdd3fdt6cfpy6d",
			"Swap-test1356g6fcq7r5v96ld3mqpntm0su0j80ppdm4r80",
			"Swap-test1vvlcdwcwlu5xs9qs0scn2t8gu3rsp0vtz06kjg"
		],
		"initialStakers": [
			{
				"nodeID": "NodeID-3qztCoJtH2ZJmfo5Gz5TdQu9mte6tV55D",
				"rewardAddress": "Swap-test1dssj9m0xtvpacypu2u536rjs39wte799zpmgk7",
				"nominationFee": 20000
			},
			{
				"nodeID": "NodeID-GDfPMF6z9bvXrUsheNPSF5SoKVvN8PKFM",
				"rewardAddress": "Swap-test10cgn6cknsyvp0n6afwrzmvp93maesf4c9fdn8z",
				"nominationFee": 20000
			},
			{
				"nodeID": "NodeID-KeVFurTahZLqjcM8SdSRKQQiHce13G7rR",
				"rewardAddress": "Swap-test102u6y4juspq0e0wgqy62pdgd3579asn4nfm8zh",
				"nominationFee": 20000
			},
			{
				"nodeID": "NodeID-6c33F4uHVbX5H8AhbDuU29KHeD9LoUrAG",
				"rewardAddress": "Swap-test16wy9v23ay2w070c0ppgm35gue755hewsnv4npf",
				"nominationFee": 20000
			},
			{
				"nodeID": "NodeID-G5R2YN4KAxAdAdDr6F63JkP77mZv18jki",
				"rewardAddress": "Swap-test156uyw8uy3c4cldv4gfmy6c5p0j79q3z0762jl6",
				"nominationFee": 20000
			},
			{
				"nodeID": "NodeID-D3aq1xWbyXmuNgG66Qv57Tfx1Yk347676",
				"rewardAddress": "Swap-test1es639wwwxpzectuvdwlstc473qa6m3qrzuf3rk",
				"nominationFee": 20000
			},
			{
				"nodeID": "NodeID-EfjvM16g2Rs5qUUA4up1Jwyie8C9k7B1K",
				"rewardAddress": "Swap-test1jzc63t9hxcesljtl23efqv4dq6ehny5x4nwm0w",
				"nominationFee": 20000
			},
			{
				"nodeID": "NodeID-JXDrdUhsar2CdYnwq4YZ7py1nmJGNTLFM",
				"rewardAddress": "Swap-test16uvswz0rdj2c5ecpq6s9zeccgpq0w8lwcw4xda",
				"nominationFee": 20000
			},
			{
				"nodeID": "NodeID-PrBAuMhdfwEVsjCoLxmyqCowrLFzC9s6e",
				"rewardAddress": "Swap-test1s5n5xq0uamylwldrky9ln4fu0cus8mt6mq4ayy",
				"nominationFee": 20000
			},
			{
				"nodeID": "NodeID-DJkTyVjyecGxseQ82GeL9b2UL4WWDUJWt",
				"rewardAddress": "Swap-test1asg2ukxyhr45yyc7d4vzrwun9xfsl0zza30nlu",
				"nominationFee": 20000
			},
			{
				"nodeID": "NodeID-9P6EXhWpNwKEpJnMb7rVuAqoCcUBve59y",
				"rewardAddress": "Swap-test1tezw3pagmmjfjz53k8kt0ymmp24zq9ewvc4ey8",
				"nominationFee": 20000
			},
			{
				"nodeID": "NodeID-NQjPxpZtYgAqYBGtvoXAekxtb1jNPKt6J",
				"rewardAddress": "Swap-test1mmdwn3n2a4ft524ggzw8jnvva92jd6wjxnjg4m",
				"nominationFee": 20000
			},
			{
				"nodeID": "NodeID-MvQtjin5gjDV8wxaFjhvTxyfpas71spGs",
				"rewardAddress": "Swap-test1nqvz6agrfxwuyj6skg0j6zwxcnqrun6mz8zvkj",
				"nominationFee": 20000
			},
			{
				"nodeID": "NodeID-8w7kvFNpeX4vU5v1PxGs8U7GJLpNnFYzr",
				"rewardAddress": "Swap-test1fmrn33tlxpvmrxcppddvh4tdvr7560w0c8lcry",
				"nominationFee": 20000
			},
			{
				"nodeID": "NodeID-7W5zm4r77CTnm1AQ2RK8kV5aTQN1WrcJf",
				"rewardAddress": "Swap-test1q6xuesh0qe08eeadk262hge2xrg6e0pxelzuw8",
				"nominationFee": 20000
			},
			{
				"nodeID": "NodeID-GgtrcdnMSSQXe54fxkyp9SjzvWsTzAVVX",
				"rewardAddress": "Swap-test1kkepg8m6f49a6k66vqrs84cqretj6pvctnw7f3",
				"nominationFee": 20000
			},
			{
				"nodeID": "NodeID-JnRiWWWtYrwtjJNUNCpwQvqtqTHV9YYr8",
				"rewardAddress": "Swap-test1fslvqs9rduee2wjzyrhkfztff36kxmjt6a9urc",
				"nominationFee": 20000
			},
			{
				"nodeID": "NodeID-821sSxnqx1MMiFyUrCmpxJoTU6rvxEZEv",
				"rewardAddress": "Swap-test17n9p4egxfjz7heu065ezvnxckdd3fdt6cfpy6d",
				"nominationFee": 20000
			},
			{
				"nodeID": "NodeID-21Z3egMvBQrXPnGvA7Qp2FgJ3UGUzEEco",
				"rewardAddress": "Swap-test1356g6fcq7r5v96ld3mqpntm0su0j80ppdm4r80",
				"nominationFee": 20000
			},
			{
				"nodeID": "NodeID-9DXNudaTqHTCA5r7K1MHjHgpujtvzzkUM",
				"rewardAddress": "Swap-test1vvlcdwcwlu5xs9qs0scn2t8gu3rsp0vtz06kjg",
				"nominationFee": 20000
			}
		],
		"axChainGenesis": "{\"config\":{\"chainId\":4000,\"homesteadBlock\":0,\"daoForkBlock\":0,\"daoForkSupport\":true,\"eip150Block\":0,\"eip150Hash\":\"0x2086799aeebeae135c246c65021c82b4e15a2c451340993aacfd2751886514f0\",\"eip155Block\":0,\"eip158Block\":0,\"byzantiumBlock\":0,\"constantinopleBlock\":0,\"petersburgBlock\":0,\"istanbulBlock\":0,\"muirGlacierBlock\":0,\"apricotPhase1BlockTimestamp\":0,\"apricotPhase2BlockTimestamp\":0},\"nonce\":\"0x0\",\"timestamp\":\"0x0\",\"extraData\":\"0x00\",\"gasLimit\":\"0x7A1200\",\"minBaseFee\":1000000,\"difficulty\":\"0x0\",\"mixHash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\",\"coinbase\":\"0x0000000000000000000000000000000000000000\",\"alloc\":{\"790e5825b65ade90095ddfe30709f14f04372194\":{\"balance\":\"0xD3C21BCECCEDA1000000\"},\"07464552eae4efb8b985d311b093d25de1d69f91\":{\"balance\":\"0xD3C21BCECCEDA1000000\"},\"6fac93cbe0263cbabf9b95ce415fcdc5829713c4\":{\"balance\":\"0xD3C21BCECCEDA1000000\"},\"d05b8c7da18b39ef013dd490181d337d968a9bb0\":{\"balance\":\"0xD3C21BCECCEDA1000000\"},\"0f9ebc77d4681c2c3242fdf687806a3a1aff3422\":{\"balance\":\"0xD3C21BCECCEDA1000000\"}},\"number\":\"0x0\",\"gasUsed\":\"0x0\",\"parentHash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\"}",
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
				SupplyCap:          uint128.Uint128{Hi: 10, Lo: 000000000000000000},
			},
		},
	}
)
