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
			  "axcAddr": "Swap-axc1xxg50twr5f29jcveh5xz9cjmd8ppkcegudmjtk",
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
			  "axcAddr": "Swap-axc14gqjvde2edqft8u4tm3srmm5j3wdt4tehmhpxx",
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
			  "axcAddr": "Swap-axc14afxf464atkkn6kgluqnkpwxewhdq4d8t57qvj",
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
			  "axcAddr": "Swap-axc137nqqcqrvgfe5cwe6trw6zdvuepryfkemuyr6s",
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
			  "axcAddr": "Swap-axc1kfr2d56p0sgfvr7hftzm0f5k970c663alc5axe",
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
			  "axcAddr": "Swap-axc1tga6qwfw0yn8d5zd7p654terd5zqr3yls0zpjx",
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
			  "axcAddr": "Swap-axc1d2uftrhn354404eclk7kz9dsjnuj5744l8yrks",
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
			  "axcAddr": "Swap-axc14w6n0k3gzjwr7zw0r0pptptmwmx9apf936j26z",
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
			  "axcAddr": "Swap-axc1y522u6emc36acasthyzk3urw22e9gjuvvp98w5",
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
			  "axcAddr": "Swap-axc1admvsyerxmv5g3tw0ah3v4sxtgjcv9l5zuhcf2",
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
			  "axcAddr": "Swap-axc1d6nm0svnj4kj4d2afdgf826fydw8v3utsha7xn",
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
			  "axcAddr": "Swap-axc1fsfh6cl3dxs2su0rlf7glf99tqn9296jednvnr",
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
			  "axcAddr": "Swap-axc1kff7pskk9ymsnn3m6ulxyncyzluw9v25s47fv5",
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
			  "axcAddr": "Swap-axc13njegescvp2thdkl2lfv8xsv6kqyt6j0h4a0qk",
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
			  "axcAddr": "Swap-axc1fsfjjckzcjxew8n6uqsdjm63kuum4f6zhxasz9",
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
			  "axcAddr": "Swap-axc12fjdse5q0dwwvlf3ehqvgdmnswxmxqmze6l8w5",
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
			  "axcAddr": "Swap-axc19p3tcrftm3tadczl4d3njxx2expuf2wys2ycej",
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
			  "axcAddr": "Swap-axc1twdg7x3addafy9xkszqnx7k36n4hykylg652wd",
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
			  "axcAddr": "Swap-axc1hwu5wvkndpuugk98nnzlpjx9uxklzmuuqlgkp4",
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
			  "axcAddr": "Swap-axc1hwfc82ss4377966k8wam7z3g8yp94shezc9lae",
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
			  "axcAddr": "Swap-axc1ury8vqvkek8pccyaw5kyrd7le20nqwuxml37mp",
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
			  "axcAddr": "Swap-axc1zkwq49vpcu933zu877zvpcvtrmnwq8m0cjfj7c",
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
			  "axcAddr": "Swap-axc19kpz93fg77jkd4e3a0fnzy8279c58e9x5ujzj2",
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
			"Swap-axc1xxg50twr5f29jcveh5xz9cjmd8ppkcegudmjtk",
			"Swap-axc14gqjvde2edqft8u4tm3srmm5j3wdt4tehmhpxx",
			"Swap-axc14afxf464atkkn6kgluqnkpwxewhdq4d8t57qvj",
			"Swap-axc137nqqcqrvgfe5cwe6trw6zdvuepryfkemuyr6s",
			"Swap-axc1kfr2d56p0sgfvr7hftzm0f5k970c663alc5axe",
			"Swap-axc1tga6qwfw0yn8d5zd7p654terd5zqr3yls0zpjx",
			"Swap-axc1d2uftrhn354404eclk7kz9dsjnuj5744l8yrks",
			"Swap-axc14w6n0k3gzjwr7zw0r0pptptmwmx9apf936j26z",
			"Swap-axc1y522u6emc36acasthyzk3urw22e9gjuvvp98w5",
			"Swap-axc1admvsyerxmv5g3tw0ah3v4sxtgjcv9l5zuhcf2",
			"Swap-axc1d6nm0svnj4kj4d2afdgf826fydw8v3utsha7xn",
			"Swap-axc1fsfh6cl3dxs2su0rlf7glf99tqn9296jednvnr",
			"Swap-axc1kff7pskk9ymsnn3m6ulxyncyzluw9v25s47fv5",
			"Swap-axc13njegescvp2thdkl2lfv8xsv6kqyt6j0h4a0qk",
			"Swap-axc1fsfjjckzcjxew8n6uqsdjm63kuum4f6zhxasz9",
			"Swap-axc12fjdse5q0dwwvlf3ehqvgdmnswxmxqmze6l8w5",
			"Swap-axc19p3tcrftm3tadczl4d3njxx2expuf2wys2ycej",
			"Swap-axc1twdg7x3addafy9xkszqnx7k36n4hykylg652wd",
			"Swap-axc1hwu5wvkndpuugk98nnzlpjx9uxklzmuuqlgkp4",
			"Swap-axc1hwfc82ss4377966k8wam7z3g8yp94shezc9lae",
			"Swap-axc1ury8vqvkek8pccyaw5kyrd7le20nqwuxml37mp",
			"Swap-axc1zkwq49vpcu933zu877zvpcvtrmnwq8m0cjfj7c",
			"Swap-axc19kpz93fg77jkd4e3a0fnzy8279c58e9x5ujzj2"
		],
		"initialStakers": [
			{
			  "nodeID": "NodeID-FVJiVBMXg69Eys8v3K3fUZXM21fW5k7Ts",
			  "rewardAddress": "Swap-axc1xxg50twr5f29jcveh5xz9cjmd8ppkcegudmjtk",
			  "nominationFee": 20000
			},
			{
			  "nodeID": "NodeID-8ofNV8sb3J9rwaC3HuXcPFAkNEqF8oJzV",
			  "rewardAddress": "Swap-axc14gqjvde2edqft8u4tm3srmm5j3wdt4tehmhpxx",
			  "nominationFee": 20000
			},
			{
			  "nodeID": "NodeID-BLaJAbw7aYCmZtYoemCBpowijcBPq7yZs",
			  "rewardAddress": "Swap-axc14afxf464atkkn6kgluqnkpwxewhdq4d8t57qvj",
			  "nominationFee": 20000
			},
			{
			  "nodeID": "NodeID-FTirsWrt7FJVTTvQejEQpJERuJDKk9BW2",
			  "rewardAddress": "Swap-axc137nqqcqrvgfe5cwe6trw6zdvuepryfkemuyr6s",
			  "nominationFee": 20000
			},
			{
			  "nodeID": "NodeID-Gym2UoBo3b4tp3pukZzDjxJefNc5XE8Cw",
			  "rewardAddress": "Swap-axc1kfr2d56p0sgfvr7hftzm0f5k970c663alc5axe",
			  "nominationFee": 20000
			},
			{
			  "nodeID": "NodeID-AkZUoGaJUWrfGUBtk7Wrsyvu5DZKVkxKQ",
			  "rewardAddress": "Swap-axc1tga6qwfw0yn8d5zd7p654terd5zqr3yls0zpjx",
			  "nominationFee": 20000
			},
			{
			  "nodeID": "NodeID-FuzyvQMfFRhAWiiQLvcDVErbNoA2czuYp",
			  "rewardAddress": "Swap-axc1d2uftrhn354404eclk7kz9dsjnuj5744l8yrks",
			  "nominationFee": 20000
			},
			{
			  "nodeID": "NodeID-Ag6dycckf8nwo2EYrvfnuw4oAw8xTz6P3",
			  "rewardAddress": "Swap-axc14w6n0k3gzjwr7zw0r0pptptmwmx9apf936j26z",
			  "nominationFee": 20000
			},
			{
			  "nodeID": "NodeID-PUWVULUUXozYBvUQwoPxQFMXT13MqRBAK",
			  "rewardAddress": "Swap-axc1y522u6emc36acasthyzk3urw22e9gjuvvp98w5",
			  "nominationFee": 20000
			},
			{
			  "nodeID": "NodeID-MbjtKYV1EEbRdaK1d396g54VrVnTXet5Q",
			  "rewardAddress": "Swap-axc1admvsyerxmv5g3tw0ah3v4sxtgjcv9l5zuhcf2",
			  "nominationFee": 20000
			},
			{
			  "nodeID": "NodeID-FeFgdAG3v8YJidF4xR8E3zSHUEuNRtBfB",
			  "rewardAddress": "Swap-axc1d6nm0svnj4kj4d2afdgf826fydw8v3utsha7xn",
			  "nominationFee": 20000
			},
			{
			  "nodeID": "NodeID-2eAGdTFKwdi63NDvADr99qoDANmY11uun",
			  "rewardAddress": "Swap-axc1fsfh6cl3dxs2su0rlf7glf99tqn9296jednvnr",
			  "nominationFee": 20000
			},
			{
			  "nodeID": "NodeID-pbDBA6Gfy9ShgscRm83kmJpVECqyh2EV",
			  "rewardAddress": "Swap-axc1kff7pskk9ymsnn3m6ulxyncyzluw9v25s47fv5",
			  "nominationFee": 20000
			},
			{
			  "nodeID": "NodeID-KbQdsqnna4KwQfsRksNMah1ZbpQgvqtrX",
			  "rewardAddress": "Swap-axc13njegescvp2thdkl2lfv8xsv6kqyt6j0h4a0qk",
			  "nominationFee": 20000
			},
			{
			  "nodeID": "NodeID-Ne4NyLV1WAyyXoRAzSSSTbc7pU8jbFTez",
			  "rewardAddress": "Swap-axc1fsfjjckzcjxew8n6uqsdjm63kuum4f6zhxasz9",
			  "nominationFee": 20000
			},
			{
			  "nodeID": "NodeID-Mx3reKpoVipQVUzH41qx5gJZmNURgvFJD",
			  "rewardAddress": "Swap-axc12fjdse5q0dwwvlf3ehqvgdmnswxmxqmze6l8w5",
			  "nominationFee": 20000
			},
			{
			  "nodeID": "NodeID-CksSwGj4tSAyyRVrFavs6p9farwwtEC1b",
			  "rewardAddress": "Swap-axc19p3tcrftm3tadczl4d3njxx2expuf2wys2ycej",
			  "nominationFee": 20000
			},
			{
			  "nodeID": "NodeID-QHHAyHWvLEJsawMTYgZfAT66erXRVqtHu",
			  "rewardAddress": "Swap-axc1twdg7x3addafy9xkszqnx7k36n4hykylg652wd",
			  "nominationFee": 20000
			},
			{
			  "nodeID": "NodeID-21zDZTDjnKcLE8yMy1SXsC5oYvJq1m1Ti",
			  "rewardAddress": "Swap-axc1hwu5wvkndpuugk98nnzlpjx9uxklzmuuqlgkp4",
			  "nominationFee": 20000
			},
			{
			  "nodeID": "NodeID-Q1KAsFkpTVszwbAtANydsRTLj9jjs8PAP",
			  "rewardAddress": "Swap-axc1hwfc82ss4377966k8wam7z3g8yp94shezc9lae",
			  "nominationFee": 20000
			},
			{
			  "nodeID": "NodeID-HPz5qwnbvsiz4oqmsf7zr8NeWsmbSnAep",
			  "rewardAddress": "Swap-axc1ury8vqvkek8pccyaw5kyrd7le20nqwuxml37mp",
			  "nominationFee": 20000
			},
			{
			  "nodeID": "NodeID-G3FjbZ7rSrBqL11BpZG89RJPXkP2cUaiL",
			  "rewardAddress": "Swap-axc1zkwq49vpcu933zu877zvpcvtrmnwq8m0cjfj7c",
			  "nominationFee": 20000
			},
			{
			  "nodeID": "NodeID-5wFDMxCVVpvZip4oUrMyunwRde1aU9Vmb",
			  "rewardAddress": "Swap-axc19kpz93fg77jkd4e3a0fnzy8279c58e9x5ujzj2",
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
