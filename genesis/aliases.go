// Copyright (C) 2019-2021, Axia Systems, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package genesis

import (
	"github.com/axiacoin/axia-network-v2/ids"
	"github.com/axiacoin/axia-network-v2/utils/constants"
	"github.com/axiacoin/axia-network-v2/vms/nftfx"
	"github.com/axiacoin/axia-network-v2/vms/platformvm"
	"github.com/axiacoin/axia-network-v2/vms/propertyfx"
	"github.com/axiacoin/axia-network-v2/vms/secp256k1fx"
)

// Aliases returns the default aliases based on the network ID
func Aliases(genesisBytes []byte) (map[string][]string, map[ids.ID][]string, error) {
	apiAliases := map[string][]string{
		constants.ChainAliasPrefix + constants.PlatformChainID.String(): {
			"P",
			"platform",
			constants.ChainAliasPrefix + "P",
			constants.ChainAliasPrefix + "platform",
		},
	}
	chainAliases := map[ids.ID][]string{
		constants.PlatformChainID: {"P", "platform"},
	}
	genesis := &platformvm.Genesis{} // TODO let's not re-create genesis to do aliasing
	if _, err := platformvm.GenesisCodec.Unmarshal(genesisBytes, genesis); err != nil {
		return nil, nil, err
	}
	if err := genesis.Initialize(); err != nil {
		return nil, nil, err
	}

	for _, chain := range genesis.Chains {
		uChain := chain.UnsignedTx.(*platformvm.UnsignedCreateChainTx)
		switch uChain.VMID {
		case constants.AVMID:
			apiAliases[constants.ChainAliasPrefix+chain.ID().String()] = []string{"Swap", "avm", constants.ChainAliasPrefix + "Swap", constants.ChainAliasPrefix + "avm"}
			chainAliases[chain.ID()] = GetSwapChainAliases()
		case constants.EVMID:
			apiAliases[constants.ChainAliasPrefix+chain.ID().String()] = []string{"C", "evm", constants.ChainAliasPrefix + "C", constants.ChainAliasPrefix + "evm"}
			chainAliases[chain.ID()] = GetAXChainAliases()
		}
	}
	return apiAliases, chainAliases, nil
}

func GetAXChainAliases() []string {
	return []string{"C", "evm"}
}

func GetSwapChainAliases() []string {
	return []string{"Swap", "avm"}
}

func GetVMAliases() map[ids.ID][]string {
	return map[ids.ID][]string{
		constants.PlatformVMID: {"platform"},
		constants.AVMID:        {"avm"},
		constants.EVMID:        {"evm"},
		secp256k1fx.ID:         {"secp256k1fx"},
		nftfx.ID:               {"nftfx"},
		propertyfx.ID:          {"propertyfx"},
	}
}
