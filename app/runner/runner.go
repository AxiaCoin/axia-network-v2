// Copyright (C) 2019-2021, Axia Systems, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package runner

import (
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"

	"github.com/axiacoin/axia-network-v2/app"
	"github.com/axiacoin/axia-network-v2/app/process"
	"github.com/axiacoin/axia-network-v2/node"
	"github.com/axiacoin/axia-network-v2/vms/rpcchainvm/grpcutils"

	appplugin "github.com/axiacoin/axia-network-v2/app/plugin"
)

// Run an Axia node.
// If specified in the config, serves a hashicorp plugin that can be consumed by
// the daemon (see axia/main).
func Run(config Config, nodeConfig node.Config) {
	nodeApp := process.NewApp(nodeConfig) // Create node wrapper
	if config.PluginMode {                // Serve as a plugin
		plugin.Serve(&plugin.ServeConfig{
			HandshakeConfig: appplugin.Handshake,
			Plugins: map[string]plugin.Plugin{
				appplugin.Name: appplugin.New(nodeApp),
			},
			GRPCServer: grpcutils.NewDefaultServer, // A non-nil value here enables gRPC serving for this plugin
			Logger: hclog.New(&hclog.LoggerOptions{
				Level: hclog.Error,
			}),
		})
		return
	}


	exitCode := app.Run(nodeApp)
	os.Exit(exitCode)
}
