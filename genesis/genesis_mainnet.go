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
				"ethAddr": "0xb5f7c1eef626cc30f2c9fec3b63cd68a5b1fe716",
				"axcAddr": "Swap-axc1rzzcqwf36aqdw5e6tuy48u5m9t990fpaqz7nyu",
				"initialAmount": 14999900000000000000,
				"unlockSchedule": [
					  {
						"amount": 100000000000000,
						"locktime": 1720170000
					  }
					]
			  },
			  {
				"ethAddr": "0x0452a25e29673a8e8e5862bb6cbd496160b3e87e",
				"axcAddr": "Swap-axc1d6dcmlus7zrxpkujdxsckaus0psfhpata6qkcd",
				"initialAmount": 7499950000000000000,
				"unlockSchedule": [
					  {
						"amount": 100000000000000,
						"locktime": 1720170000
					  }
					]
			  },
			  {
				"ethAddr": "0xdC8B26CF5dA71419F87b7fCb7652f368b9A4B095",
				"axcAddr": "Swap-axc195urz4jext723lpdu0809r9uuchupamxfkh9q6",
				"initialAmount": 7499950000000000000,
				"unlockSchedule": [
					  {
						"amount": 100000000000000,
						"locktime": 1720170000
					  }
					]
			  },
			  {
				"ethAddr": "0x91f3559Cad61d7fbd9983525515DA4A6c1793D7E",
				"axcAddr": "Swap-axc1d8a4xklysgfgry4ltayvlnerjp6whsfpy4z4tr",
				"initialAmount": 7499950000000000000,
				"unlockSchedule": [
					  {
						"amount": 100000000000000,
						"locktime": 1720170000
					  }
					]
			  },
			  {
				"ethAddr": "0xAccf0E02593f0b6AaB1B2298a2Ce933c96C92623",
				"axcAddr": "Swap-axc1addxfscvd28a5l89tmryjje9ax4nll7f6zuzsg",
				"initialAmount": 7499950000000000000,
				"unlockSchedule": [
					  {
						"amount": 100000000000000,
						"locktime": 1720170000
					  }
					]
			  },
			  {
				"ethAddr": "0x8f90748dF95f0f69F4269A14265B8Bc966864912",
				"axcAddr": "Swap-axc1qlvg8r7hkcs477twdmx0rch3m05q4plyufkp8v",
				"initialAmount": 7499950000000000000,
				"unlockSchedule": [
					  {
						"amount": 100000000000000,
						"locktime": 1720170000
					  }
					]
			  },
			  {
				"ethAddr": "0xBE4a848D12b09Ad476DbB3537cD7C8AEc6Aa95ba",
				"axcAddr": "Swap-axc1dmd5t276mc5drfg25pkjnhk9n6q5ek3s4c7fzq",
				"initialAmount": 7499950000000000000,
				"unlockSchedule": [
					  {
						"amount": 100000000000000,
						"locktime": 1720170000
					  }
					]
			  },
			  {
				"ethAddr": "0xA7108D1F5231969060ca4d5bb151c0efaC8BF55C",
				"axcAddr": "Swap-axc1exq6gsfuv0evfalfarth9p94w8l8nxsl72qhxw",
				"initialAmount": 7499950000000000000,
				"unlockSchedule": [
					  {
						"amount": 100000000000000,
						"locktime": 1720170000
					  }
					]
			  },
			  {
				"ethAddr": "0x3af8bf4CbD8303493D845EF7BBcF4A3B4e17BBE9",
				"axcAddr": "Swap-axc1kl8ckmwdfp0je7jmdhnyy96y0ra0qr0sdw6dck",
				"initialAmount": 7499950000000000000,
				"unlockSchedule": [
					  {
						"amount": 100000000000000,
						"locktime": 1720170000
					  }
					]
			  },
			  {
				"ethAddr": "0x15514038CEd579A705084Cd215C61DbC2b5B865f",
				"axcAddr": "Swap-axc1wagx7284qfxf4ytpeq5l4rp6839ryls2fdnflt",
				"initialAmount": 7499950000000000000,
				"unlockSchedule": [
					  {
						"amount": 100000000000000,
						"locktime": 1720170000
					  }
					]
			  },
			  {
				"ethAddr": "0x2D3f09163065264CE2625a28EB5D6cF1e272c2cc",
				"axcAddr": "Swap-axc1n4xpwteadxpsgpwwugcm8uq6x720985kxqe4ex",
				"initialAmount": 7499950000000000000,
				"unlockSchedule": [
					  {
						"amount": 100000000000000,
						"locktime": 1720170000
					  }
					]
			  },
			  {
				"ethAddr": "0xcC691B5761D25aa072C9629F3b6d16b5A8806a81",
				"axcAddr": "Swap-axc1n4av78e2xnrugzy46xw8vz9cycc8t6pmsvtv5z",
				"initialAmount": 7499950000000000000,
				"unlockSchedule": [
					  {
						"amount": 100000000000000,
						"locktime": 1720170000
					  }
					]
			  },
			  {
				"ethAddr": "0xA1c7E7E7E35e05125cF1a50f89d35fA5159697bd",
				"axcAddr": "Swap-axc169hga7cqgt4ml3nn5a7fceeaaudd9nfxutp0kz",
				"initialAmount": 7499950000000000000,
				"unlockSchedule": []
			  },
			  {
				"ethAddr": "0x9dB15182550344543da9aed945bBCeC255365C78",
				"axcAddr": "Swap-axc1rjwsufwm2s9w8mkmgkjasqutq96x09mthre2zx",
				"initialAmount": 7499950000000000000,
				"unlockSchedule": []
			  },
			  {
				"ethAddr": "0x8d149952a42a18BD60f18DbDEf7deb3B59AaC4cB",
				"axcAddr": "Swap-axc1qzfyl3kpjg2a3lx9lqad7nuv7aj5sjm0l43tj4",
				"initialAmount": 7499950000000000000,
				"unlockSchedule": []
			  },
			  {
				"ethAddr": "0x713157810308767ec1263c4C84f79bE5f7c0D664",
				"axcAddr": "Swap-axc12kff3jtr5mrggrcsz8cw250532wxxe4uguunvy",
				"initialAmount": 7499950000000000000,
				"unlockSchedule": []
			  },
			  {
				"ethAddr": "0x9d95dfa99E56E04839c4B00324d5e3A4F87558C2",
				"axcAddr": "Swap-axc1n90hg5gtp7jma4nmz69jeny0ff2f6y3vfgrc9a",
				"initialAmount": 7499950000000000000,
				"unlockSchedule": []
			  },
			  {
				"ethAddr": "0x59Ea67A9D1054f390da043C833B8f88e6A32A683",
				"axcAddr": "Swap-axc1zdr4x6hvswjh9h0nkgzt9wqae57mk7ydgapc2d",
				"initialAmount": 7499950000000000000,
				"unlockSchedule": []
			  },
			  {
				"ethAddr": "0x448f7544A18C822c418BdBc48571E260EEd4DD1B",
				"axcAddr": "Swap-axc1ev58795mqtjuy429vrc92l0rer7fqf09rawlmy",
				"initialAmount": 7499950000000000000,
				"unlockSchedule": []
			  },
			  {
				"ethAddr": "0x74b1C859c9801F5cD42CCcdeF32c4c9988e2a2eC",
				"axcAddr": "Swap-axc1z2rvegtxzwrdgvsnuq3fql2kwdrh7nly9q53tm",
				"initialAmount": 7499950000000000000,
				"unlockSchedule": []
			  },
			  {
				"ethAddr": "0x7089faC055ba1ad9eA0945Dd6c167D572259E314",
				"axcAddr": "Swap-axc18wh7jmc7uyz0jqe8x2mjkrpdsawxdlhu9qrwqx",
				"initialAmount": 7499950000000000000,
				"unlockSchedule": []
			  },
			  {
				"ethAddr": "0xe61537Be48d9700B2Bb0498B33307b1C17FA0F9b",
				"axcAddr": "Swap-axc145jfmzhptpp3gur8579kyt4eyvqz4wwf38k4u0",
				"initialAmount": 7499950000000000000,
				"unlockSchedule": []
			  },
			  {
				"ethAddr": "0x4d45f5Ad0bb94D1C3E3a22FED3110eC9837e4D99",
				"axcAddr": "Swap-axc1n86uy54m2s4hjy245lacnaxrkgua076amt6vjq",
				"initialAmount": 7499950000000000000,
				"unlockSchedule": []
			  }
		],
		"startTime": 1657011600,
		"initialStakeDuration": 63072000,
		"initialStakeDurationOffset": 5400,
		"initialStakedFunds": [
			"Swap-axc1rzzcqwf36aqdw5e6tuy48u5m9t990fpaqz7nyu",
			"Swap-axc1d6dcmlus7zrxpkujdxsckaus0psfhpata6qkcd",
			"Swap-axc195urz4jext723lpdu0809r9uuchupamxfkh9q6",
			"Swap-axc1d8a4xklysgfgry4ltayvlnerjp6whsfpy4z4tr",
			"Swap-axc1addxfscvd28a5l89tmryjje9ax4nll7f6zuzsg",
			"Swap-axc1qlvg8r7hkcs477twdmx0rch3m05q4plyufkp8v",
			"Swap-axc1dmd5t276mc5drfg25pkjnhk9n6q5ek3s4c7fzq",
			"Swap-axc1exq6gsfuv0evfalfarth9p94w8l8nxsl72qhxw",
			"Swap-axc1kl8ckmwdfp0je7jmdhnyy96y0ra0qr0sdw6dck",
			"Swap-axc1wagx7284qfxf4ytpeq5l4rp6839ryls2fdnflt",
			"Swap-axc1n4xpwteadxpsgpwwugcm8uq6x720985kxqe4ex",
			"Swap-axc1n4av78e2xnrugzy46xw8vz9cycc8t6pmsvtv5z"
		],
		"initialStakers": [
			{
				"nodeID": "NodeID-H8Yey5QX1fKTrNAdVFPPwVJjixr4pnxYE",
				"rewardAddress": "Swap-axc1rzzcqwf36aqdw5e6tuy48u5m9t990fpaqz7nyu",
				"nominationFee": 20000
			  },
			  {
				"nodeID": "NodeID-J8ce9EyXVBSzkE7D2G4RcM9bhJLDiaj4C",
				"rewardAddress": "Swap-axc1d6dcmlus7zrxpkujdxsckaus0psfhpata6qkcd",
				"nominationFee": 20000
			  },
			  {
				"nodeID": "NodeID-N4GJrrsRHyAyJ6eyhbVLbe5hQDZm6Z2xD",
				"rewardAddress": "Swap-axc195urz4jext723lpdu0809r9uuchupamxfkh9q6",
				"nominationFee": 20000
			  },
			  {
				"nodeID": "NodeID-7ihRQ8QX98HYEguBxA8yfzfhgZabMmWnP",
				"rewardAddress": "Swap-axc1d8a4xklysgfgry4ltayvlnerjp6whsfpy4z4tr",
				"nominationFee": 20000
			  },
			  {
				"nodeID": "NodeID-JCrs3zduYwg7EDXu431bN5L4vM2kTvx8B",
				"rewardAddress": "Swap-axc1addxfscvd28a5l89tmryjje9ax4nll7f6zuzsg",
				"nominationFee": 20000
			  },
			  {
				"nodeID": "NodeID-JMmBfGTTuDoPzGzdfdH9jxZDwNJ2CajBQ",
				"rewardAddress": "Swap-axc1qlvg8r7hkcs477twdmx0rch3m05q4plyufkp8v",
				"nominationFee": 20000
			  },
			  {
				"nodeID": "NodeID-NeSHHK4SSSmqSWwvs45oqUxK5tnDEAdVE",
				"rewardAddress": "Swap-axc1dmd5t276mc5drfg25pkjnhk9n6q5ek3s4c7fzq",
				"nominationFee": 20000
			  },
			  {
				"nodeID": "NodeID-Aod3799i2D3manPEpy3NMfx9DZosWq6wQ",
				"rewardAddress": "Swap-axc1exq6gsfuv0evfalfarth9p94w8l8nxsl72qhxw",
				"nominationFee": 20000
			  },
			  {
				"nodeID": "NodeID-KvLN8t2XnruHhpJPQroNjnuF28howzKLE",
				"rewardAddress": "Swap-axc1kl8ckmwdfp0je7jmdhnyy96y0ra0qr0sdw6dck",
				"nominationFee": 20000
			  },
			  {
				"nodeID": "NodeID-LxXuxjktNJ4vDZgH2YiBbw2cibpcdwNCC",
				"rewardAddress": "Swap-axc1wagx7284qfxf4ytpeq5l4rp6839ryls2fdnflt",
				"nominationFee": 20000
			  },
			  {
				"nodeID": "NodeID-PoEdto9WVkNetcGAqYeWc4x8pbzTWRSuL",
				"rewardAddress": "Swap-axc1n4xpwteadxpsgpwwugcm8uq6x720985kxqe4ex",
				"nominationFee": 20000
			  },
			  {
				"nodeID": "NodeID-9ZGpMrJ6xriryxDfsQcqSFpEvqe7EBHrG",
				"rewardAddress": "Swap-axc1n4av78e2xnrugzy46xw8vz9cycc8t6pmsvtv5z",
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
