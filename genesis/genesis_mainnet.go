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
			  "ethAddr": "0x4D2216462cEf1348aC09d2bDe817473ffa8BBf70",
			  "axcAddr": "Swap-axc17tvxpcn4qfzekc9u0l6dama68x2a8tlx0q4ymk",
			  "initialAmount": 7826087000000000000,
			  "unlockSchedule": [
					{
					  "amount": 1000000000000000,
					  "locktime": 1688816151
					}
				  ]
			},
			{
			  "ethAddr": "0xce1c4815A10A0B4e3f4BCCeF52d8d7BC9aC38cf2",
			  "axcAddr": "Swap-axc1r20qucnpjn0ecx7m3lfwxuqfejzcrxy6246ul6",
			  "initialAmount": 7826087000000000000,
			  "unlockSchedule": [
					{
					  "amount": 1000000000000000,
					  "locktime": 1688816151
					}
				  ]
			},
			{
			  "ethAddr": "0xcAc8765E3a55dE8a7B293A8614DA528de69ca6a2",
			  "axcAddr": "Swap-axc19znen38lt5tu6g8kqkc4zye3v2yg6nhyfn3uvd",
			  "initialAmount": 7826087000000000000,
			  "unlockSchedule": [
					{
					  "amount": 1000000000000000,
					  "locktime": 1688816151
					}
				  ]
			},
			{
			  "ethAddr": "0x7A17513a080766B97a1D15570c386aAE3f1eED2F",
			  "axcAddr": "Swap-axc1s4j7xpjrjznkx66ksjda50rur8sur3cz6dl227",
			  "initialAmount": 7826087000000000000,
			  "unlockSchedule": [
					{
					  "amount": 1000000000000000,
					  "locktime": 1688816151
					}
				  ]
			},
			{
			  "ethAddr": "0xaED28B010eBB55465e86A95fB363Ea360B7572F4",
			  "axcAddr": "Swap-axc1xk8m4hkym5g6rvxh4rafyv59a4pn4v4m7g49nm",
			  "initialAmount": 7826087000000000000,
			  "unlockSchedule": [
					{
					  "amount": 1000000000000000,
					  "locktime": 1688816151
					}
				  ]
			},
			{
			  "ethAddr": "0x76db86467Fc30942A8d5D4B2f506fb5F81c7dd54",
			  "axcAddr": "Swap-axc19jc0tze73ed4fywhu6wvml083s5h6vnwqwy3l9",
			  "initialAmount": 7826087000000000000,
			  "unlockSchedule": [
					{
					  "amount": 1000000000000000,
					  "locktime": 1688816151
					}
				  ]
			},
			{
			  "ethAddr": "0xE3c9650cf02ec69c457d5bfaeD43c12A141Bf303",
			  "axcAddr": "Swap-axc15n3cxa2qkxecc6jc0kvql4fsjx94q8g7na24la",
			  "initialAmount": 7826087000000000000,
			  "unlockSchedule": [
					{
					  "amount": 1000000000000000,
					  "locktime": 1688816151
					}
				  ]
			},
			{
			  "ethAddr": "0xcaecbfdc5727b7f094D09e1D273c19478C90F5AC",
			  "axcAddr": "Swap-axc1tgmqymvp3ltl9zaedpkxvt3vejsn52jnlwxkrx",
			  "initialAmount": 7826087000000000000,
			  "unlockSchedule": [
					{
					  "amount": 1000000000000000,
					  "locktime": 1688816151
					}
				  ]
			},
			{
			  "ethAddr": "0xcDEFA8e65A19d2981af83f523f0b3Cdd4411edB1",
			  "axcAddr": "Swap-axc1dz4qhkz72mlurcq7tgf2xe9y3m083vwpzef42a",
			  "initialAmount": 7826087000000000000,
			  "unlockSchedule": [
					{
					  "amount": 1000000000000000,
					  "locktime": 1688816151
					}
				  ]
			},
			{
			  "ethAddr": "0x24e4a639d5a08c3Fe86410FDB86531395c09F251",
			  "axcAddr": "Swap-axc1afm07602785l3ln29yh2k527zrj7g7ktng22ry",
			  "initialAmount": 7826087000000000000,
			  "unlockSchedule": [
					{
					  "amount": 1000000000000000,
					  "locktime": 1688816151
					}
				  ]
			},
			{
			  "ethAddr": "0xF830cB93dEFFc2649a4d9873e41096D74da11d01",
			  "axcAddr": "Swap-axc1zxxf2q0sj7nx8z57gnncqd5h6ygw3v6fwcq8j7",
			  "initialAmount": 7826087000000000000,
			  "unlockSchedule": [
					{
					  "amount": 1000000000000000,
					  "locktime": 1688816151
					}
				  ]
			},
			{
			  "ethAddr": "0xB0A1F3331aA978C4Ae389939D41B227899f4Bb9F",
			  "axcAddr": "Swap-axc1c44f5usljdg6rlputzh5p4kqr33zp0r7ks7fhk",
			  "initialAmount": 7826087000000000000,
			  "unlockSchedule": [
					{
					  "amount": 1000000000000000,
					  "locktime": 1688816151
					}
				  ]
			},
			{
			  "ethAddr": "0x44E7Cea37f4d2a2FC980FA4a63fE346ea9D60B9A",
			  "axcAddr": "Swap-axc1nqsf0mr9mu55qp3e4j58y8w0jdk89mn4pzgn0z",
			  "initialAmount": 7826087000000000000,
			  "unlockSchedule": [
					{
					  "amount": 1000000000000000,
					  "locktime": 1688816151
					}
				  ]
			},
			{
			  "ethAddr": "0x58802CA234F8dB0170E9E77FD491BC9Ca1603F6D",
			  "axcAddr": "Swap-axc1sejdysh4rj5ns54v58kxuaq4j5mehr0t6lvz3k",
			  "initialAmount": 7826087000000000000,
			  "unlockSchedule": [
					{
					  "amount": 1000000000000000,
					  "locktime": 1688816151
					}
				  ]
			},
			{
			  "ethAddr": "0xc718cEaF716aBFcD8999c7B0AC067ad4e440faB5",
			  "axcAddr": "Swap-axc1njjxswayvp3cwjyy68durj3e844pua6zdlt64t",
			  "initialAmount": 7826087000000000000,
			  "unlockSchedule": [
					{
					  "amount": 1000000000000000,
					  "locktime": 1688816151
					}
				  ]
			},
			{
			  "ethAddr": "0x9B62A14B8d0dBda38CB274D938D6377383F72529",
			  "axcAddr": "Swap-axc1us3grqx44edxu0mgj6p4wmv0ux055r03xcpquq",
			  "initialAmount": 7826087000000000000,
			  "unlockSchedule": [
					{
					  "amount": 1000000000000000,
					  "locktime": 1688816151
					}
				  ]
			},
			{
			  "ethAddr": "0x99566c876d16687834B29AD3e15d3Afe1A578052",
			  "axcAddr": "Swap-axc1sewfx0j04swn2830ve8v8lkll935ng6n8apykw",
			  "initialAmount": 7826087000000000000,
			  "unlockSchedule": [
					{
					  "amount": 1000000000000000,
					  "locktime": 1688816151
					}
				  ]
			},
			{
			  "ethAddr": "0x16da52Ed03ecABf2244d89dbCEbcaCC44f9F0D49",
			  "axcAddr": "Swap-axc14hf9hpkw8l2m22jxlma697cewgm7359l3gq6m2",
			  "initialAmount": 7826087000000000000,
			  "unlockSchedule": [
					{
					  "amount": 1000000000000000,
					  "locktime": 1688816151
					}
				  ]
			},
			{
			  "ethAddr": "0x94FC7b98C44b193EbF66f79E0608f95c0b93363C",
			  "axcAddr": "Swap-axc1wna3hy27qs27yjpk6hlqrv4wa3a3jj2ym0tczj",
			  "initialAmount": 7826087000000000000,
			  "unlockSchedule": [
					{
					  "amount": 1000000000000000,
					  "locktime": 1688816151
					}
				  ]
			},
			{
			  "ethAddr": "0xf64AD39c3D40606ce319B824DbA3F972ca8fF585",
			  "axcAddr": "Swap-axc1vez04vmwe4xc06f7kff8903wlk39ty8xhjd4p8",
			  "initialAmount": 7826087000000000000,
			  "unlockSchedule": [
					{
					  "amount": 1000000000000000,
					  "locktime": 1688816151
					}
				  ]
			},
			{
			  "ethAddr": "0xEb5a5a786875439967e4eF26a49CC70131fb4b0f",
			  "axcAddr": "Swap-axc1pk3tuv3s6lheuy9qzhwy76m47ztl78z9alr4aq",
			  "initialAmount": 7826087000000000000,
			  "unlockSchedule": [
					{
					  "amount": 1000000000000000,
					  "locktime": 1688816151
					}
				  ]
			},
			{
			  "ethAddr": "0x78a80a315aE2c2EaFc938913109B2D89B166Ad03",
			  "axcAddr": "Swap-axc1gy28ymwhyfqgtdcannyjnsx5js9a5m466le8la",
			  "initialAmount": 7826087000000000000,
			  "unlockSchedule": [
					{
					  "amount": 1000000000000000,
					  "locktime": 1688816151
					}
				  ]
			},
			{
			  "ethAddr": "0x33EdB9CA5Cb7F9D57885E1dcC9091EE9359750DA",
			  "axcAddr": "Swap-axc1kqrwmmlwd6wtmp05qwsdlv7kfa7uxsxuvglhzz",
			  "initialAmount": 7826087000000000000,
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
			"Swap-axc17tvxpcn4qfzekc9u0l6dama68x2a8tlx0q4ymk",
			"Swap-axc1r20qucnpjn0ecx7m3lfwxuqfejzcrxy6246ul6",
			"Swap-axc19znen38lt5tu6g8kqkc4zye3v2yg6nhyfn3uvd",
			"Swap-axc1s4j7xpjrjznkx66ksjda50rur8sur3cz6dl227",
			"Swap-axc1xk8m4hkym5g6rvxh4rafyv59a4pn4v4m7g49nm",
			"Swap-axc19jc0tze73ed4fywhu6wvml083s5h6vnwqwy3l9",
			"Swap-axc15n3cxa2qkxecc6jc0kvql4fsjx94q8g7na24la",
			"Swap-axc1tgmqymvp3ltl9zaedpkxvt3vejsn52jnlwxkrx",
			"Swap-axc1dz4qhkz72mlurcq7tgf2xe9y3m083vwpzef42a",
			"Swap-axc1afm07602785l3ln29yh2k527zrj7g7ktng22ry",
			"Swap-axc1zxxf2q0sj7nx8z57gnncqd5h6ygw3v6fwcq8j7",
			"Swap-axc1c44f5usljdg6rlputzh5p4kqr33zp0r7ks7fhk",
			"Swap-axc1nqsf0mr9mu55qp3e4j58y8w0jdk89mn4pzgn0z",
			"Swap-axc1sejdysh4rj5ns54v58kxuaq4j5mehr0t6lvz3k",
			"Swap-axc1njjxswayvp3cwjyy68durj3e844pua6zdlt64t",
			"Swap-axc1us3grqx44edxu0mgj6p4wmv0ux055r03xcpquq",
			"Swap-axc1sewfx0j04swn2830ve8v8lkll935ng6n8apykw",
			"Swap-axc14hf9hpkw8l2m22jxlma697cewgm7359l3gq6m2",
			"Swap-axc1wna3hy27qs27yjpk6hlqrv4wa3a3jj2ym0tczj",
			"Swap-axc1vez04vmwe4xc06f7kff8903wlk39ty8xhjd4p8",
			"Swap-axc1pk3tuv3s6lheuy9qzhwy76m47ztl78z9alr4aq",
			"Swap-axc1gy28ymwhyfqgtdcannyjnsx5js9a5m466le8la",
			"Swap-axc1kqrwmmlwd6wtmp05qwsdlv7kfa7uxsxuvglhzz"
		],
		"initialStakers": [
			{
			  "nodeID": "NodeID-FVJiVBMXg69Eys8v3K3fUZXM21fW5k7Ts",
			  "rewardAddress": "Swap-axc17tvxpcn4qfzekc9u0l6dama68x2a8tlx0q4ymk",
			  "nominationFee": 20000
			},
			{
			  "nodeID": "NodeID-8ofNV8sb3J9rwaC3HuXcPFAkNEqF8oJzV",
			  "rewardAddress": "Swap-axc1r20qucnpjn0ecx7m3lfwxuqfejzcrxy6246ul6",
			  "nominationFee": 20000
			},
			{
			  "nodeID": "NodeID-BLaJAbw7aYCmZtYoemCBpowijcBPq7yZs",
			  "rewardAddress": "Swap-axc19znen38lt5tu6g8kqkc4zye3v2yg6nhyfn3uvd",
			  "nominationFee": 20000
			},
			{
			  "nodeID": "NodeID-FTirsWrt7FJVTTvQejEQpJERuJDKk9BW2",
			  "rewardAddress": "Swap-axc1s4j7xpjrjznkx66ksjda50rur8sur3cz6dl227",
			  "nominationFee": 20000
			},
			{
			  "nodeID": "NodeID-Gym2UoBo3b4tp3pukZzDjxJefNc5XE8Cw",
			  "rewardAddress": "Swap-axc1xk8m4hkym5g6rvxh4rafyv59a4pn4v4m7g49nm",
			  "nominationFee": 20000
			},
			{
			  "nodeID": "NodeID-AkZUoGaJUWrfGUBtk7Wrsyvu5DZKVkxKQ",
			  "rewardAddress": "Swap-axc19jc0tze73ed4fywhu6wvml083s5h6vnwqwy3l9",
			  "nominationFee": 20000
			},
			{
			  "nodeID": "NodeID-FuzyvQMfFRhAWiiQLvcDVErbNoA2czuYp",
			  "rewardAddress": "Swap-axc15n3cxa2qkxecc6jc0kvql4fsjx94q8g7na24la",
			  "nominationFee": 20000
			},
			{
			  "nodeID": "NodeID-Ag6dycckf8nwo2EYrvfnuw4oAw8xTz6P3",
			  "rewardAddress": "Swap-axc1tgmqymvp3ltl9zaedpkxvt3vejsn52jnlwxkrx",
			  "nominationFee": 20000
			},
			{
			  "nodeID": "NodeID-PUWVULUUXozYBvUQwoPxQFMXT13MqRBAK",
			  "rewardAddress": "Swap-axc1dz4qhkz72mlurcq7tgf2xe9y3m083vwpzef42a",
			  "nominationFee": 20000
			},
			{
			  "nodeID": "NodeID-MbjtKYV1EEbRdaK1d396g54VrVnTXet5Q",
			  "rewardAddress": "Swap-axc1afm07602785l3ln29yh2k527zrj7g7ktng22ry",
			  "nominationFee": 20000
			},
			{
			  "nodeID": "NodeID-FeFgdAG3v8YJidF4xR8E3zSHUEuNRtBfB",
			  "rewardAddress": "Swap-axc1zxxf2q0sj7nx8z57gnncqd5h6ygw3v6fwcq8j7",
			  "nominationFee": 20000
			},
			{
			  "nodeID": "NodeID-2eAGdTFKwdi63NDvADr99qoDANmY11uun",
			  "rewardAddress": "Swap-axc1c44f5usljdg6rlputzh5p4kqr33zp0r7ks7fhk",
			  "nominationFee": 20000
			},
			{
			  "nodeID": "NodeID-pbDBA6Gfy9ShgscRm83kmJpVECqyh2EV",
			  "rewardAddress": "Swap-axc1nqsf0mr9mu55qp3e4j58y8w0jdk89mn4pzgn0z",
			  "nominationFee": 20000
			},
			{
			  "nodeID": "NodeID-KbQdsqnna4KwQfsRksNMah1ZbpQgvqtrX",
			  "rewardAddress": "Swap-axc1sejdysh4rj5ns54v58kxuaq4j5mehr0t6lvz3k",
			  "nominationFee": 20000
			},
			{
			  "nodeID": "NodeID-Ne4NyLV1WAyyXoRAzSSSTbc7pU8jbFTez",
			  "rewardAddress": "Swap-axc1njjxswayvp3cwjyy68durj3e844pua6zdlt64t",
			  "nominationFee": 20000
			},
			{
			  "nodeID": "NodeID-Mx3reKpoVipQVUzH41qx5gJZmNURgvFJD",
			  "rewardAddress": "Swap-axc1us3grqx44edxu0mgj6p4wmv0ux055r03xcpquq",
			  "nominationFee": 20000
			},
			{
			  "nodeID": "NodeID-CksSwGj4tSAyyRVrFavs6p9farwwtEC1b",
			  "rewardAddress": "Swap-axc1sewfx0j04swn2830ve8v8lkll935ng6n8apykw",
			  "nominationFee": 20000
			},
			{
			  "nodeID": "NodeID-QHHAyHWvLEJsawMTYgZfAT66erXRVqtHu",
			  "rewardAddress": "Swap-axc14hf9hpkw8l2m22jxlma697cewgm7359l3gq6m2",
			  "nominationFee": 20000
			},
			{
			  "nodeID": "NodeID-21zDZTDjnKcLE8yMy1SXsC5oYvJq1m1Ti",
			  "rewardAddress": "Swap-axc1wna3hy27qs27yjpk6hlqrv4wa3a3jj2ym0tczj",
			  "nominationFee": 20000
			},
			{
			  "nodeID": "NodeID-Q1KAsFkpTVszwbAtANydsRTLj9jjs8PAP",
			  "rewardAddress": "Swap-axc1vez04vmwe4xc06f7kff8903wlk39ty8xhjd4p8",
			  "nominationFee": 20000
			},
			{
			  "nodeID": "NodeID-HPz5qwnbvsiz4oqmsf7zr8NeWsmbSnAep",
			  "rewardAddress": "Swap-axc1pk3tuv3s6lheuy9qzhwy76m47ztl78z9alr4aq",
			  "nominationFee": 20000
			},
			{
			  "nodeID": "NodeID-G3FjbZ7rSrBqL11BpZG89RJPXkP2cUaiL",
			  "rewardAddress": "Swap-axc1gy28ymwhyfqgtdcannyjnsx5js9a5m466le8la",
			  "nominationFee": 20000
			},
			{
			  "nodeID": "NodeID-5wFDMxCVVpvZip4oUrMyunwRde1aU9Vmb",
			  "rewardAddress": "Swap-axc1kqrwmmlwd6wtmp05qwsdlv7kfa7uxsxuvglhzz",
			  "nominationFee": 20000
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
