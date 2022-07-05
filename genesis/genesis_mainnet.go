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
				"initialAmount": 240000000000,
				"unlockSchedule": [
				{
					"amount": 0,
					"locktime": 1664957273
					}
					]
			},
			{
				"ethAddr": "0xb4caAc709D045AD998B50aB2B2202F109f5Dcf76",
				"axcAddr": "Swap-custom173x6vjspg95nn4n68cerdhjnhwtehkc4ey6dlf",
				"initialAmount": 240000000000,
				"unlockSchedule": [
				{
					"amount": 0,
					"locktime": 1664957273
					}
					]
			},
			{
				"ethAddr": "0x725c7698d19F89C395789d58714c2Adb287242aD",
				"axcAddr": "Swap-custom197g3cfnq5k3lemp5h9l5dthe4plzdldge2kh0m",
				"initialAmount": 240000000000,
				"unlockSchedule": [
				{
					"amount": 0,
					"locktime": 1664957273
					}
					]
			},
			{
				"ethAddr": "0x22CB8Bbb1Da74c3B93c3C9442a01D01c98E85b82",
				"axcAddr": "Swap-custom1s2kquu7shue6aq9xg8zklm0jvjnwzha7qsv8p3",
				"initialAmount": 240000000000,
				"unlockSchedule": [
				{
					"amount": 0,
					"locktime": 1664957273
					}
					]
			},
			{
				"ethAddr": "0xb37A616Fd5d3ffabf8bf8AE21b975A732A4C45A2",
				"axcAddr": "Swap-custom15lm8nkxt0rz4nlkrx3emqj5rlkskwt545vgft3",
				"initialAmount": 240000000000,
				"unlockSchedule": [
				{
					"amount": 0,
					"locktime": 1664957273
					}
					]
			},
			{
				"ethAddr": "0xe462d667F98Bd7d0f167100Cb73170CC31943E5c",
				"axcAddr": "Swap-custom1mzw7ej4usa22wax93lkje68la4w65wa0zx2ayq",
				"initialAmount": 240000000000,
				"unlockSchedule": [
				{
					"amount": 0,
					"locktime": 1664957273
					}
					]
			},
			{
				"ethAddr": "0xbd8FBC8D65BDBC5F05a53ed247D7A881E6976777",
				"axcAddr": "Swap-custom102ffs9hx83kvuxqk09ffyd5dlzxvlugfpz735z",
				"initialAmount": 240000000000,
				"unlockSchedule": [
				{
					"amount": 0,
					"locktime": 1664957273
					}
					]
			},
			{
				"ethAddr": "0x8fa07a39f8f0BBe436F12c92e9c5F37D88dFcB71",
				"axcAddr": "Swap-custom1xp7wtcznw3fksdagcn3uhhh6w2vpyhwrar0acl",
				"initialAmount": 240000000000,
				"unlockSchedule": [
				{
					"amount": 0,
					"locktime": 1664957273
					}
					]
			},
			{
				"ethAddr": "0x69621fF5EA18DD7bFae4060B04A901c8290Add45",
				"axcAddr": "Swap-custom1k6mg398l2u6ptsqdfh3kp9qpdrd6yx4q8lggml",
				"initialAmount": 240000000000,
				"unlockSchedule": [
				{
					"amount": 0,
					"locktime": 1664957273
					}
					]
			},
			{
				"ethAddr": "0x413b668F25d91435c1f38441706D856ff65e7dbE",
				"axcAddr": "Swap-custom13ch9spukcunug2e2lcqsr5wrkc5f5a75hv8tne",
				"initialAmount": 240000000000,
				"unlockSchedule": [
				{
					"amount": 0,
					"locktime": 1664957273
					}
					]
			},
			{
				"ethAddr": "0x645285290C4B12A4dC567c2F7baF167823AE1FF9",
				"axcAddr": "Swap-custom13a43uyspwx8w5lqdfgu5vzplrm8ev0se6whwrq",
				"initialAmount": 240000000000,
				"unlockSchedule": [
				{
					"amount": 0,
					"locktime": 1664957273
					}
					]
			},
			{
				"ethAddr": "0x2A318b8bcAFe7fbE96Af043816248918F4c66208",
				"axcAddr": "Swap-custom1jf5zakcz9qupyj74vjvcasf5895kfc8kt3t0jp",
				"initialAmount": 240000000000,
				"unlockSchedule": [
				{
					"amount": 0,
					"locktime": 1664957273
					}
					]
			},
			{
				"ethAddr": "0x343668497653C33F979394264ad950B8aC9Aa091",
				"axcAddr": "Swap-custom1t3mrtggn0rmymq04nkdjaq0m0v6anseqlm5jce",
				"initialAmount": 240000000000,
				"unlockSchedule": [
				{
					"amount": 0,
					"locktime": 1664957273
					}
					]
			},
			{
				"ethAddr": "0x0cDC895C8D56C0B4FeFE3c9bcfAF43657341ecF9",
				"axcAddr": "Swap-custom10wdj6e67v9ng30rvvggtfp6glhalrrtm3khvr9",
				"initialAmount": 240000000000,
				"unlockSchedule": [
				{
					"amount": 0,
					"locktime": 1664957273
					}
					]
			},
			{
				"ethAddr": "0x3560864D2cb7645613c3634163041233fA30F5E3",
				"axcAddr": "Swap-custom1kgujpjy4j43pkaxxa8j4ryhldv3dqgxfvhw08g",
				"initialAmount": 240000000000,
				"unlockSchedule": [
				{
					"amount": 0,
					"locktime": 1664957273
					}
					]
			},
			{
				"ethAddr": "0x09458Fe1Fef7Ceb2f04f92d4b5FF33134fb7Cb01",
				"axcAddr": "Swap-custom1vtk37wvup7smfxne0759m99ejclqm2spj9h3y4",
				"initialAmount": 240000000000,
				"unlockSchedule": [
				{
					"amount": 0,
					"locktime": 1664957273
					}
					]
			},
			{
				"ethAddr": "0x4Cd3Fa7d7483c419692D504bC14d216406f8c769",
				"axcAddr": "Swap-custom1m63ygw4dfeeyzqdzpf973hg0n5j9tzz2q5gsun",
				"initialAmount": 240000000000,
				"unlockSchedule": [
				{
					"amount": 0,
					"locktime": 1664957273
					}
					]
			},
			{
				"ethAddr": "0x761fAeDA8d23704f42Cec06A0b82261cdfc05e3B",
				"axcAddr": "Swap-custom19g8merermr6t65k2xks3ufftf6phfsvwjck7gk",
				"initialAmount": 240000000000,
				"unlockSchedule": [
				{
					"amount": 0,
					"locktime": 1664957273
					}
					]
			},
			{
				"ethAddr": "0x7Dea7FDDD3384eF7E9134a55317bdb7cE9690cDe",
				"axcAddr": "Swap-custom1maflf8cytj0x3ke9wsagze2nhqnunjvf8s5d95",
				"initialAmount": 240000000000,
				"unlockSchedule": [
				{
					"amount": 0,
					"locktime": 1664957273
					}
					]
			},
			{
				"ethAddr": "0x1d04Ea52545BCbae78bDd2FDdaE34aEd9c380c19",
				"axcAddr": "Swap-custom1hnwdhnzxff8slvxkw6dh9ffqyrw3v8dwvu7j8h",
				"initialAmount": 240000000000,
				"unlockSchedule": [
				{
					"amount": 0,
					"locktime": 1664957273
					}
					]
			},
			{
				"ethAddr": "0x2fa7add662AE8152bbfE7208c006f0e06D17B5C5",
				"axcAddr": "Swap-custom148l3qjpp96d654xp09ymlsrda0fdz22se6jkxu",
				"initialAmount": 240000000000,
				"unlockSchedule": [
				{
					"amount": 0,
					"locktime": 1664957273
					}
					]
			},
			{
				"ethAddr": "0x129CAeF7E185EB9605Ab8305d4F5042C43D191d1",
				"axcAddr": "Swap-custom1fsqj0nj6jrucmfcn4z62n64fln28lznh28a2xn",
				"initialAmount": 240000000000,
				"unlockSchedule": [
				{
					"amount": 0,
					"locktime": 1664957273
					}
					]
			},
			{
				"ethAddr": "0x5f4Dc7ea4De0d9aBda8c5AADdcfaB8864f8c674A",
				"axcAddr": "Swap-custom1udkepu0rdv3zs8ywm6ndpxv7rfslv6ukj52gkp",
				"initialAmount": 240000000000,
				"unlockSchedule": [
				{
					"amount": 0,
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
			"Swap-custom15lm8nkxt0rz4nlkrx3emqj5rlkskwt545vgft3",
			"Swap-custom1mzw7ej4usa22wax93lkje68la4w65wa0zx2ayq",
			"Swap-custom102ffs9hx83kvuxqk09ffyd5dlzxvlugfpz735z",
			"Swap-custom1xp7wtcznw3fksdagcn3uhhh6w2vpyhwrar0acl",
			"Swap-custom1k6mg398l2u6ptsqdfh3kp9qpdrd6yx4q8lggml",
			"Swap-custom13ch9spukcunug2e2lcqsr5wrkc5f5a75hv8tne",
			"Swap-custom13a43uyspwx8w5lqdfgu5vzplrm8ev0se6whwrq",
			"Swap-custom1jf5zakcz9qupyj74vjvcasf5895kfc8kt3t0jp",
			"Swap-custom1t3mrtggn0rmymq04nkdjaq0m0v6anseqlm5jce",
			"Swap-custom10wdj6e67v9ng30rvvggtfp6glhalrrtm3khvr9",
			"Swap-custom1kgujpjy4j43pkaxxa8j4ryhldv3dqgxfvhw08g",
			"Swap-custom1vtk37wvup7smfxne0759m99ejclqm2spj9h3y4",
			"Swap-custom1m63ygw4dfeeyzqdzpf973hg0n5j9tzz2q5gsun",
			"Swap-custom19g8merermr6t65k2xks3ufftf6phfsvwjck7gk",
			"Swap-custom1maflf8cytj0x3ke9wsagze2nhqnunjvf8s5d95",
			"Swap-custom1hnwdhnzxff8slvxkw6dh9ffqyrw3v8dwvu7j8h",
			"Swap-custom148l3qjpp96d654xp09ymlsrda0fdz22se6jkxu",
			"Swap-custom1fsqj0nj6jrucmfcn4z62n64fln28lznh28a2xn",
			"Swap-custom1udkepu0rdv3zs8ywm6ndpxv7rfslv6ukj52gkp"
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
			},
			{
				"nodeID": "NodeID-AkZUoGaJUWrfGUBtk7Wrsyvu5DZKVkxKQ",
				"rewardAddress": "Swap-custom1mzw7ej4usa22wax93lkje68la4w65wa0zx2ayq",
				"delegationFee": 200000
			},
			{
				"nodeID": "NodeID-FuzyvQMfFRhAWiiQLvcDVErbNoA2czuYp",
				"rewardAddress": "Swap-custom102ffs9hx83kvuxqk09ffyd5dlzxvlugfpz735z",
				"delegationFee": 200000
			},
			{
				"nodeID": "NodeID-Ag6dycckf8nwo2EYrvfnuw4oAw8xTz6P3",
				"rewardAddress": "Swap-custom1xp7wtcznw3fksdagcn3uhhh6w2vpyhwrar0acl",
				"delegationFee": 200000
			},
			{
				"nodeID": "NodeID-PUWVULUUXozYBvUQwoPxQFMXT13MqRBAK",
				"rewardAddress": "Swap-custom1k6mg398l2u6ptsqdfh3kp9qpdrd6yx4q8lggml",
				"delegationFee": 200000
			},
			{
				"nodeID": "NodeID-MbjtKYV1EEbRdaK1d396g54VrVnTXet5Q",
				"rewardAddress": "Swap-custom13ch9spukcunug2e2lcqsr5wrkc5f5a75hv8tne",
				"delegationFee": 200000
			},
			{
				"nodeID": "NodeID-FeFgdAG3v8YJidF4xR8E3zSHUEuNRtBfB",
				"rewardAddress": "Swap-custom13a43uyspwx8w5lqdfgu5vzplrm8ev0se6whwrq",
				"delegationFee": 200000
			},
			{
				"nodeID": "NodeID-2eAGdTFKwdi63NDvADr99qoDANmY11uun",
				"rewardAddress": "Swap-custom1jf5zakcz9qupyj74vjvcasf5895kfc8kt3t0jp",
				"delegationFee": 200000
			},
			{
				"nodeID": "NodeID-pbDBA6Gfy9ShgscRm83kmJpVECqyh2EV",
				"rewardAddress": "Swap-custom1t3mrtggn0rmymq04nkdjaq0m0v6anseqlm5jce",
				"delegationFee": 200000
			},
			{
				"nodeID": "NodeID-KbQdsqnna4KwQfsRksNMah1ZbpQgvqtrX",
				"rewardAddress": "Swap-custom10wdj6e67v9ng30rvvggtfp6glhalrrtm3khvr9",
				"delegationFee": 200000
			},
			{
				"nodeID": "NodeID-Ne4NyLV1WAyyXoRAzSSSTbc7pU8jbFTez",
				"rewardAddress": "Swap-custom1kgujpjy4j43pkaxxa8j4ryhldv3dqgxfvhw08g",
				"delegationFee": 200000
			},
			{
				"nodeID": "NodeID-Mx3reKpoVipQVUzH41qx5gJZmNURgvFJD",
				"rewardAddress": "Swap-custom1vtk37wvup7smfxne0759m99ejclqm2spj9h3y4",
				"delegationFee": 200000
			},
			{
				"nodeID": "NodeID-CksSwGj4tSAyyRVrFavs6p9farwwtEC1b",
				"rewardAddress": "Swap-custom1m63ygw4dfeeyzqdzpf973hg0n5j9tzz2q5gsun",
				"delegationFee": 200000
			},
			{
				"nodeID": "NodeID-QHHAyHWvLEJsawMTYgZfAT66erXRVqtHu",
				"rewardAddress": "Swap-custom19g8merermr6t65k2xks3ufftf6phfsvwjck7gk",
				"delegationFee": 200000
			},
			{
				"nodeID": "NodeID-21zDZTDjnKcLE8yMy1SXsC5oYvJq1m1Ti",
				"rewardAddress": "Swap-custom1maflf8cytj0x3ke9wsagze2nhqnunjvf8s5d95",
				"delegationFee": 200000
			},
			{
				"nodeID": "NodeID-Q1KAsFkpTVszwbAtANydsRTLj9jjs8PAP",
				"rewardAddress": "Swap-custom1hnwdhnzxff8slvxkw6dh9ffqyrw3v8dwvu7j8h",
				"delegationFee": 200000
			},
			{
				"nodeID": "NodeID-HPz5qwnbvsiz4oqmsf7zr8NeWsmbSnAep",
				"rewardAddress": "Swap-custom148l3qjpp96d654xp09ymlsrda0fdz22se6jkxu",
				"delegationFee": 200000
			},
			{
				"nodeID": "NodeID-G3FjbZ7rSrBqL11BpZG89RJPXkP2cUaiL",
				"rewardAddress": "Swap-custom1fsqj0nj6jrucmfcn4z62n64fln28lznh28a2xn",
				"delegationFee": 200000
			},
			{
				"nodeID": "NodeID-5wFDMxCVVpvZip4oUrMyunwRde1aU9Vmb",
				"rewardAddress": "Swap-custom1udkepu0rdv3zs8ywm6ndpxv7rfslv6ukj52gkp",
				"delegationFee": 200000
			}
		],
		"axChainGenesis": "{\"config\":{\"chainId\":4001,\"homesteadBlock\":0,\"daoForkBlock\":0,\"daoForkSupport\":true,\"eip150Block\":0,\"eip150Hash\":\"0x2086799aeebeae135c246c65021c82b4e15a2c451340993aacfd2751886514f0\",\"eip155Block\":0,\"eip158Block\":0,\"byzantiumBlock\":0,\"constantinopleBlock\":0,\"petersburgBlock\":0,\"istanbulBlock\":0,\"muirGlacierBlock\":0},\"nonce\":\"0x0\",\"timestamp\":\"0x0\",\"extraData\":\"0x00\",\"gasLimit\":\"0x5f5e100\",\"difficulty\":\"0x0\",\"mixHash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\",\"coinbase\":\"0x0000000000000000000000000000000000000000\",\"alloc\":{\"000000000000000000000000\":{\"code\":\"0x7300000000000000000000000000000000000000003014608060405260043610603d5760003560e01c80631e010439146042578063b6510bb314606e575b600080fd5b605c60048036036020811015605657600080fd5b503560b1565b60408051918252519081900360200190f35b818015607957600080fd5b5060af60048036036080811015608e57600080fd5b506001600160a01b03813516906020810135906040810135906060013560b6565b005b30cd90565b836001600160a01b031681836108fc8690811502906040516000604051808303818888878c8acf9550505050505015801560f4573d6000803e3d6000fd5b505050505056fea26469706673582212201eebce970fe3f5cb96bf8ac6ba5f5c133fc2908ae3dcd51082cfee8f583429d064736f6c634300060a0033\",\"balance\":\"0x0\"}},\"number\":\"0x0\",\"gasUsed\":\"0x0\",\"parentHash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\"}",
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
