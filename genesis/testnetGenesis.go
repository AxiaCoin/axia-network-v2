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
			"ethAddr": "0x790e5825b65ade90095ddfe30709f14f04372194",
			"axcAddr": "Swap-custom1n9ulqdw7t5atnscqakh8m8v050fl90masx5haf",
			"initialAmount": 3000000000000000000,
			"unlockSchedule": [
			  {
				"amount": 10000000000000000,
				"locktime": 1655936730
			  }
			]
		  },
		  {
			"ethAddr": "0x07464552eae4efb8b985d311b093d25de1d69f91",
			"axcAddr": "Swap-custom1n8fjuq8h72w424kt80svgw4vx2gpej62gtd4pc",
			"initialAmount": 3000000000000000000,
			"unlockSchedule": [
			 
			  {
				"amount": 1000000000000000,
				"locktime": 1655936730
			  }
			]
		  },
		  {
			"ethAddr": "0x6fac93cbe0263cbabf9b95ce415fcdc5829713c4",
			"axcAddr": "Swap-custom19dl9a50ctqea369c54egn9qatuvvzpaa2rpdg9",
			"initialAmount": 3000000000000000000,
			"unlockSchedule": [
			  {
				"amount": 1000000000000000,
				"locktime": 1655936730
			  }
			]
		  },
		  {
			"ethAddr": "0xd05b8c7da18b39ef013dd490181d337d968a9bb0",
			"axcAddr": "Swap-custom1wy5etvnjn64gpyms2wuhwtl2gflud5uw8uwav2",
			"initialAmount": 3000000000000000000,
			"unlockSchedule": [
			  {
				"amount": 1000000000000000,
				"locktime": 1655936730
			  }
			]
		  },
		  {
			"ethAddr": "0x0f9ebc77d4681c2c3242fdf687806a3a1aff3422",
			"axcAddr": "Swap-custom1nyy6vlxxvt5amv4xvsj9ywzy75sfshr2z2qnpq",
			"initialAmount": 3000000000000000000,
			"unlockSchedule": [
			  {
				"amount": 1000000000000000,
				"locktime": 1655936730
			  }
			]
		  },
		  {
			"ethAddr": "0x190d0360eedc8a42830cb86f52590b573cd3796b",
			"axcAddr": "Swap-custom15r8vw35knz7ahtu8nka2mtrqm46kvaa3dlsk00",
			"initialAmount": 3000000000000000000,
			"unlockSchedule": [
			  {
				"amount": 1000000000000000,
				"locktime": 1655936730
			  }
			]
		  }
		],
		"startTime": 1630987200,
		"initialStakeDuration": 31536000,
		"initialStakeDurationOffset": 5400,
		"initialStakedFunds": [
		  "Swap-custom1n9ulqdw7t5atnscqakh8m8v050fl90masx5haf"
		],
		"initialStakers": [
		  {
			"nodeID": "NodeID-5cu5HH5gBCWnPqh1mKJYbRXFtCpFq22hL",
			"rewardAddress": "Swap-custom1n9ulqdw7t5atnscqakh8m8v050fl90masx5haf",
			"nominationFee": 1000000
		  }, 
		  {
			"nodeID": "NodeID-7qAvF7Zdb7SGsPrYSCLTxqJHN5AZ6QQSs",
			"rewardAddress": "Swap-custom1n9ulqdw7t5atnscqakh8m8v050fl90masx5haf",
			"nominationFee": 500000
		  },
		  {
			"nodeID": "NodeID-NFBbbJ4qCmNaCzeW7sxErhvWqvEQMnYcN",
			"rewardAddress": "Swap-custom1n9ulqdw7t5atnscqakh8m8v050fl90masx5haf",
			"nominationFee": 250000
		  },
		  {
			"nodeID": "NodeID-5FEKyX6b86Kjvhqsbb8gmCnSB9toAUmjP",
			"rewardAddress": "Swap-custom1n9ulqdw7t5atnscqakh8m8v050fl90masx5haf",
			"nominationFee": 125000
		  }
		],
		"axChainGenesis": "{\"config\":{\"chainId\":4000,\"homesteadBlock\":0,\"daoForkBlock\":0,\"daoForkSupport\":true,\"eip150Block\":0,\"eip150Hash\":\"0x2086799aeebeae135c246c65021c82b4e15a2c451340993aacfd2751886514f0\",\"eip155Block\":0,\"eip158Block\":0,\"byzantiumBlock\":0,\"constantinopleBlock\":0,\"petersburgBlock\":0,\"istanbulBlock\":0,\"muirGlacierBlock\":0,\"apricotPhase1BlockTimestamp\":0,\"apricotPhase2BlockTimestamp\":0},\"nonce\":\"0x0\",\"timestamp\":\"0x0\",\"extraData\":\"0x00\",\"gasLimit\":\"0x7A1200\",\"minBaseFee\":1000000,\"difficulty\":\"0x0\",\"mixHash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\",\"coinbase\":\"0x0000000000000000000000000000000000000000\",\"alloc\":{\"790e5825b65ade90095ddfe30709f14f04372194\":{\"balance\":\"0xD3C21BCECCEDA1000000\"},\"07464552eae4efb8b985d311b093d25de1d69f91\":{\"balance\":\"0xD3C21BCECCEDA1000000\"},\"6fac93cbe0263cbabf9b95ce415fcdc5829713c4\":{\"balance\":\"0xD3C21BCECCEDA1000000\"},\"d05b8c7da18b39ef013dd490181d337d968a9bb0\":{\"balance\":\"0xD3C21BCECCEDA1000000\"},\"0f9ebc77d4681c2c3242fdf687806a3a1aff3422\":{\"balance\":\"0xD3C21BCECCEDA1000000\"}},\"number\":\"0x0\",\"gasUsed\":\"0x0\",\"parentHash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\"}",
		"message": "{{ fun_quote }}"
	  }`

	TestParams = Params{
		TxFeeConfig: TxFeeConfig{
			TxFee:                 units.MilliAxc,
			CreateAssetTxFee:      10 * units.MilliAxc,
			CreateAllychainTxFee:  100 * units.MilliAxc,
			CreateBlockchainTxFee: 100 * units.MilliAxc,
		},
		StakingConfig: StakingConfig{
			UptimeRequirement: .8, // 80%
			MinValidatorStake: 10 * units.KiloAxc,
			MaxValidatorStake: 3 * units.MegaAxc,
			MinNominatorStake: 1 * units.Axc,
			MinNominationFee:  20000, // 2%
			MinStakeDuration:  120 * 24 * time.Hour,
			MaxStakeDuration:  2 * 365 * 24 * time.Hour,
			RewardConfig: reward.Config{
				MaxConsumptionRate: .12 * reward.PercentDenominator,
				MinConsumptionRate: .10 * reward.PercentDenominator,
				MintingPeriod:      365 * 24 * time.Hour,
				SupplyCap:          uint128.Uint128{Hi: 180, Lo: 000000000000000000},
			},
		},
	}
)
