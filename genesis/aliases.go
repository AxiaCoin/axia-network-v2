// Copyright (C) 2019-2021, Axia Systems, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package genesis

import (
	"path"

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
		path.Join(constants.ChainAliasPrefix, constants.CoreChainID.String()): {
			"Core",
			"core",
			path.Join(constants.ChainAliasPrefix, "Core"),
			path.Join(constants.ChainAliasPrefix, "core"),
		},
	}
	chainAliases := map[ids.ID][]string{
		constants.CoreChainID: {"Core", "core"},
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
		chainID := chain.ID()
		endpoint := path.Join(constants.ChainAliasPrefix, chainID.String())
		switch uChain.VMID {
		case constants.AVMID:
			apiAliases[endpoint] = []string{
				"Swap",
				"avm",
				path.Join(constants.ChainAliasPrefix, "Swap"),
				path.Join(constants.ChainAliasPrefix, "avm"),
			}
			chainAliases[chainID] = GetSwapChainAliases()
		case constants.EVMID:
			apiAliases[endpoint] = []string{
				"A",
				"evm",
				path.Join(constants.ChainAliasPrefix, "A"),
				path.Join(constants.ChainAliasPrefix, "evm"),
			}
			chainAliases[chainID] = GetAXChainAliases()
		}
	}
	return apiAliases, chainAliases, nil
}

func GetAXChainAliases() []string {
	return []string{"A", "evm"}
}

func GetSwapChainAliases() []string {
	return []string{"Swap", "avm"}
}

func GetVMAliases() map[ids.ID][]string {
	return map[ids.ID][]string{
		constants.PlatformVMID: {"core"},
		constants.AVMID:        {"avm"},
		constants.EVMID:        {"evm"},
		secp256k1fx.ID:         {"secp256k1fx"},
		nftfx.ID:               {"nftfx"},
		propertyfx.ID:          {"propertyfx"},
	}
}
