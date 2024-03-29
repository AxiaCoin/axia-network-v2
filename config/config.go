// Copyright (C) 2019-2021, Axia Systems, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package config

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/viper"

	"github.com/axiacoin/axia-network-v2/app/runner"
	"github.com/axiacoin/axia-network-v2/chains"
	"github.com/axiacoin/axia-network-v2/genesis"
	"github.com/axiacoin/axia-network-v2/ids"
	"github.com/axiacoin/axia-network-v2/ipcs"
	"github.com/axiacoin/axia-network-v2/nat"
	"github.com/axiacoin/axia-network-v2/network"
	"github.com/axiacoin/axia-network-v2/network/dialer"
	"github.com/axiacoin/axia-network-v2/network/throttling"
	"github.com/axiacoin/axia-network-v2/node"
	"github.com/axiacoin/axia-network-v2/snow/consensus/axia"
	"github.com/axiacoin/axia-network-v2/snow/consensus/snowball"
	"github.com/axiacoin/axia-network-v2/snow/networking/benchlist"
	"github.com/axiacoin/axia-network-v2/snow/networking/router"
	"github.com/axiacoin/axia-network-v2/snow/networking/sender"
	"github.com/axiacoin/axia-network-v2/staking"
	"github.com/axiacoin/axia-network-v2/utils"
	"github.com/axiacoin/axia-network-v2/utils/constants"
	"github.com/axiacoin/axia-network-v2/utils/dynamicip"
	"github.com/axiacoin/axia-network-v2/utils/logging"
	"github.com/axiacoin/axia-network-v2/utils/password"
	"github.com/axiacoin/axia-network-v2/utils/profiler"
	"github.com/axiacoin/axia-network-v2/utils/storage"
	"github.com/axiacoin/axia-network-v2/utils/timer"
	"github.com/axiacoin/axia-network-v2/utils/uint128"
	"github.com/axiacoin/axia-network-v2/vms"
)

const (
	pluginsDirName         = "plugins"
	chainConfigFileName    = "config"
	chainUpgradeFileName   = "upgrade"
	allychainConfigFileExt = ".json"
)

var (
	deprecatedKeys = map[string]string{}

	errInvalidStakerWeights          = errors.New("staking weights must be positive")
	errStakingDisableOnPublicNetwork = errors.New("staking disabled on public network")
	errAuthPasswordTooWeak           = errors.New("API auth password is not strong enough")
	errInvalidUptimeRequirement      = errors.New("uptime requirement must be in the range [0, 1]")
	errMinValidatorStakeAboveMax     = errors.New("minimum validator stake can't be greater than maximum validator stake")
	errInvalidNominationFee          = errors.New("nomination fee must be in the range [0, 1,000,000]")
	errInvalidMinStakeDuration       = errors.New("min stake duration must be > 0")
	errMinStakeDurationAboveMax      = errors.New("max stake duration can't be less than min stake duration")
	errStakeMaxConsumptionBelowMin   = errors.New("stake max consumption can't be less than min stake consumption")
	errStakeMintingPeriodBelowMin    = errors.New("stake minting period can't be less than max stake duration")
	errCannotWhitelistPrimaryNetwork = errors.New("cannot whitelist primary network")
	errStakingKeyContentUnset        = fmt.Errorf("%s key not set but %s set", StakingKeyContentKey, StakingCertContentKey)
	errStakingCertContentUnset       = fmt.Errorf("%s key set but %s not set", StakingKeyContentKey, StakingCertContentKey)
)

func GetRunnerConfig(v *viper.Viper) (runner.Config, error) {
	config := runner.Config{
		DisplayVersionAndExit: v.GetBool(VersionKey),
		BuildDir:              os.ExpandEnv(v.GetString(BuildDirKey)),
		PluginMode:            v.GetBool(PluginModeKey),
	}

	// Build directory should have this structure:
	//
	// build
	// ├── axia (the binary from compiling the app directory)
	// └── plugins
	//     └── evm
	validBuildDir := func(dir string) bool {
		info, err := os.Stat(dir)
		if err != nil || !info.IsDir() {
			return false
		}

		// make sure the expected subdirectory exists
		_, err = os.Stat(filepath.Join(dir, pluginsDirName))
		return err == nil
	}
	if validBuildDir(config.BuildDir) {
		return config, nil
	}

	foundBuildDir := false
	for _, dir := range defaultBuildDirs {
		if validBuildDir(dir) {
			config.BuildDir = dir
			foundBuildDir = true
			break
		}
	}
	if !foundBuildDir {
		return runner.Config{}, fmt.Errorf(
			"couldn't find valid build directory in any of the default locations: %s",
			defaultBuildDirs,
		)
	}
	return config, nil
}

func getConsensusConfig(v *viper.Viper) axia.Parameters {
	return axia.Parameters{
		Parameters: snowball.Parameters{
			K:                     v.GetInt(SnowSampleSizeKey),
			Alpha:                 v.GetInt(SnowQuorumSizeKey),
			BetaVirtuous:          v.GetInt(SnowVirtuousCommitThresholdKey),
			BetaRogue:             v.GetInt(SnowRogueCommitThresholdKey),
			ConcurrentRepolls:     v.GetInt(SnowConcurrentRepollsKey),
			OptimalProcessing:     v.GetInt(SnowOptimalProcessingKey),
			MaxOutstandingItems:   v.GetInt(SnowMaxProcessingKey),
			MaxItemProcessingTime: v.GetDuration(SnowMaxTimeProcessingKey),
		},
		BatchSize: v.GetInt(SnowAxiaBatchSizeKey),
		Parents:   v.GetInt(SnowAxiaNumParentsKey),
	}
}

func getLoggingConfig(v *viper.Viper) (logging.Config, error) {
	loggingConfig := logging.DefaultConfig
	if v.IsSet(LogsDirKey) {
		loggingConfig.Directory = os.ExpandEnv(v.GetString(LogsDirKey))
	}
	var err error
	loggingConfig.LogLevel, err = logging.ToLevel(v.GetString(LogLevelKey))
	if err != nil {
		return loggingConfig, err
	}
	logDisplayLevel := v.GetString(LogLevelKey)
	if v.IsSet(LogDisplayLevelKey) {
		logDisplayLevel = v.GetString(LogDisplayLevelKey)
	}
	loggingConfig.DisplayLevel, err = logging.ToLevel(logDisplayLevel)
	if err != nil {
		return loggingConfig, err
	}
	loggingConfig.DisplayHighlight, err = logging.ToHighlight(v.GetString(LogDisplayHighlightKey), os.Stdout.Fd())
	loggingConfig.DisableWriterDisplaying = v.GetBool(LogDisableDisplayPluginLogsKey)
	return loggingConfig, err
}

func getAPIAuthConfig(v *viper.Viper) (node.APIAuthConfig, error) {
	config := node.APIAuthConfig{
		APIRequireAuthToken: v.GetBool(APIAuthRequiredKey),
	}
	if !config.APIRequireAuthToken {
		return config, nil
	}

	if v.IsSet(APIAuthPasswordKey) {
		config.APIAuthPassword = v.GetString(APIAuthPasswordKey)
	} else {
		passwordFilePath := v.GetString(APIAuthPasswordFileKey) // picks flag value or default
		passwordBytes, err := os.ReadFile(passwordFilePath)
		if err != nil {
			return node.APIAuthConfig{}, fmt.Errorf("API auth password file %q failed to be read: %w", passwordFilePath, err)
		}
		config.APIAuthPassword = strings.TrimSpace(string(passwordBytes))
	}

	if !password.SufficientlyStrong(config.APIAuthPassword, password.OK) {
		return node.APIAuthConfig{}, errAuthPasswordTooWeak
	}
	return config, nil
}

func getIPCConfig(v *viper.Viper) node.IPCConfig {
	config := node.IPCConfig{
		IPCAPIEnabled: v.GetBool(IpcAPIEnabledKey),
		IPCPath:       ipcs.DefaultBaseURL,
	}
	if v.IsSet(IpcsChainIDsKey) {
		config.IPCDefaultChainIDs = strings.Split(v.GetString(IpcsChainIDsKey), ",")
	}
	if v.IsSet(IpcsPathKey) {
		config.IPCPath = os.ExpandEnv(v.GetString(IpcsPathKey))
	}
	return config
}

func getHTTPConfig(v *viper.Viper) (node.HTTPConfig, error) {
	var (
		httpsKey  []byte
		httpsCert []byte
		err       error
	)
	switch {
	case v.IsSet(HTTPSKeyContentKey):
		rawContent := v.GetString(HTTPSKeyContentKey)
		httpsKey, err = base64.StdEncoding.DecodeString(rawContent)
		if err != nil {
			return node.HTTPConfig{}, fmt.Errorf("unable to decode base64 content: %w", err)
		}
	case v.IsSet(HTTPSKeyFileKey):
		httpsKeyFilepath := os.ExpandEnv(v.GetString(HTTPSKeyFileKey))
		if httpsKey, err = os.ReadFile(filepath.Clean(httpsKeyFilepath)); err != nil {
			return node.HTTPConfig{}, err
		}
	}

	switch {
	case v.IsSet(HTTPSCertContentKey):
		rawContent := v.GetString(HTTPSCertContentKey)
		httpsCert, err = base64.StdEncoding.DecodeString(rawContent)
		if err != nil {
			return node.HTTPConfig{}, fmt.Errorf("unable to decode base64 content: %w", err)
		}
	case v.IsSet(HTTPSCertFileKey):
		httpsCertFilepath := os.ExpandEnv(v.GetString(HTTPSCertFileKey))
		if httpsCert, err = os.ReadFile(filepath.Clean(httpsCertFilepath)); err != nil {
			return node.HTTPConfig{}, err
		}
	}

	config := node.HTTPConfig{
		APIConfig: node.APIConfig{
			APIIndexerConfig: node.APIIndexerConfig{
				IndexAPIEnabled:      v.GetBool(IndexEnabledKey),
				IndexAllowIncomplete: v.GetBool(IndexAllowIncompleteKey),
			},
			AdminAPIEnabled:    v.GetBool(AdminAPIEnabledKey),
			InfoAPIEnabled:     v.GetBool(InfoAPIEnabledKey),
			KeystoreAPIEnabled: v.GetBool(KeystoreAPIEnabledKey),
			MetricsAPIEnabled:  v.GetBool(MetricsAPIEnabledKey),
			HealthAPIEnabled:   v.GetBool(HealthAPIEnabledKey),
		},
		HTTPHost:          v.GetString(HTTPHostKey),
		HTTPPort:          uint16(v.GetUint(HTTPPortKey)),
		HTTPSEnabled:      v.GetBool(HTTPSEnabledKey),
		HTTPSKey:          httpsKey,
		HTTPSCert:         httpsCert,
		APIAllowedOrigins: v.GetStringSlice(HTTPAllowedOrigins),

		ShutdownTimeout: v.GetDuration(HTTPShutdownTimeoutKey),
		ShutdownWait:    v.GetDuration(HTTPShutdownWaitKey),
	}

	config.APIAuthConfig, err = getAPIAuthConfig(v)
	if err != nil {
		return node.HTTPConfig{}, err
	}
	config.IPCConfig = getIPCConfig(v)
	return config, nil
}

func getRouterHealthConfig(v *viper.Viper, halflife time.Duration) (router.HealthConfig, error) {
	config := router.HealthConfig{
		MaxDropRate:            v.GetFloat64(RouterHealthMaxDropRateKey),
		MaxOutstandingRequests: int(v.GetUint(RouterHealthMaxOutstandingRequestsKey)),
		MaxOutstandingDuration: v.GetDuration(NetworkHealthMaxOutstandingDurationKey),
		MaxRunTimeRequests:     v.GetDuration(NetworkMaximumTimeoutKey),
		MaxDropRateHalflife:    halflife,
	}
	switch {
	case config.MaxDropRate < 0 || config.MaxDropRate > 1:
		return router.HealthConfig{}, fmt.Errorf("%q must be in [0,1]", RouterHealthMaxDropRateKey)
	case config.MaxOutstandingDuration <= 0:
		return router.HealthConfig{}, fmt.Errorf("%q must be positive", NetworkHealthMaxOutstandingDurationKey)
	case config.MaxRunTimeRequests <= 0:
		return router.HealthConfig{}, fmt.Errorf("%q must be positive", NetworkMaximumTimeoutKey)
	}
	return config, nil
}

func getAdaptiveTimeoutConfig(v *viper.Viper) (timer.AdaptiveTimeoutConfig, error) {
	config := timer.AdaptiveTimeoutConfig{
		InitialTimeout:     v.GetDuration(NetworkInitialTimeoutKey),
		MinimumTimeout:     v.GetDuration(NetworkMinimumTimeoutKey),
		MaximumTimeout:     v.GetDuration(NetworkMaximumTimeoutKey),
		TimeoutHalflife:    v.GetDuration(NetworkTimeoutHalflifeKey),
		TimeoutCoefficient: v.GetFloat64(NetworkTimeoutCoefficientKey),
	}
	switch {
	case config.MinimumTimeout < 1:
		return timer.AdaptiveTimeoutConfig{}, fmt.Errorf("%q must be positive", NetworkMinimumTimeoutKey)
	case config.MinimumTimeout > config.MaximumTimeout:
		return timer.AdaptiveTimeoutConfig{}, fmt.Errorf("%q must be >= %q", NetworkMaximumTimeoutKey, NetworkMinimumTimeoutKey)
	case config.InitialTimeout < config.MinimumTimeout || config.InitialTimeout > config.MaximumTimeout:
		return timer.AdaptiveTimeoutConfig{}, fmt.Errorf("%q must be in [%q, %q]", NetworkInitialTimeoutKey, NetworkMinimumTimeoutKey, NetworkMaximumTimeoutKey)
	case config.TimeoutHalflife <= 0:
		return timer.AdaptiveTimeoutConfig{}, fmt.Errorf("%q must > 0", NetworkTimeoutHalflifeKey)
	case config.TimeoutCoefficient < 1:
		return timer.AdaptiveTimeoutConfig{}, fmt.Errorf("%q must be >= 1", NetworkTimeoutCoefficientKey)
	}

	return config, nil
}

func getGossipConfig(v *viper.Viper) sender.GossipConfig {
	return sender.GossipConfig{
		AcceptedFrontierValidatorSize:    uint(v.GetUint32(ConsensusGossipAcceptedFrontierValidatorSizeKey)),
		AcceptedFrontierNonValidatorSize: uint(v.GetUint32(ConsensusGossipAcceptedFrontierNonValidatorSizeKey)),
		AcceptedFrontierPeerSize:         uint(v.GetUint32(ConsensusGossipAcceptedFrontierPeerSizeKey)),
		OnAcceptValidatorSize:            uint(v.GetUint32(ConsensusGossipOnAcceptValidatorSizeKey)),
		OnAcceptNonValidatorSize:         uint(v.GetUint32(ConsensusGossipOnAcceptNonValidatorSizeKey)),
		OnAcceptPeerSize:                 uint(v.GetUint32(ConsensusGossipOnAcceptPeerSizeKey)),
		AppGossipValidatorSize:           uint(v.GetUint32(AppGossipValidatorSizeKey)),
		AppGossipNonValidatorSize:        uint(v.GetUint32(AppGossipNonValidatorSizeKey)),
		AppGossipPeerSize:                uint(v.GetUint32(AppGossipPeerSizeKey)),
	}
}

func getNetworkConfig(v *viper.Viper, halflife time.Duration) (network.Config, error) {
	// Set the max number of recent inbound connections upgraded to be
	// equal to the max number of inbound connections per second.
	maxInboundConnsPerSec := v.GetFloat64(InboundThrottlerMaxConnsPerSecKey)
	upgradeCooldown := v.GetDuration(InboundConnUpgradeThrottlerCooldownKey)
	upgradeCooldownInSeconds := upgradeCooldown.Seconds()
	maxRecentConnsUpgraded := int(math.Ceil(maxInboundConnsPerSec * upgradeCooldownInSeconds))
	config := network.Config{
		// Throttling
		ThrottlerConfig: network.ThrottlerConfig{
			MaxInboundConnsPerSec: maxInboundConnsPerSec,
			InboundConnUpgradeThrottlerConfig: throttling.InboundConnUpgradeThrottlerConfig{
				UpgradeCooldown:        upgradeCooldown,
				MaxRecentConnsUpgraded: maxRecentConnsUpgraded,
			},

			InboundMsgThrottlerConfig: throttling.InboundMsgThrottlerConfig{
				MsgByteThrottlerConfig: throttling.MsgByteThrottlerConfig{
					AtLargeAllocSize:    v.GetUint64(InboundThrottlerAtLargeAllocSizeKey),
					VdrAllocSize:        v.GetUint64(InboundThrottlerVdrAllocSizeKey),
					NodeMaxAtLargeBytes: v.GetUint64(InboundThrottlerNodeMaxAtLargeBytesKey),
				},
				BandwidthThrottlerConfig: throttling.BandwidthThrottlerConfig{
					RefillRate:   v.GetUint64(InboundThrottlerBandwidthRefillRateKey),
					MaxBurstSize: v.GetUint64(InboundThrottlerBandwidthMaxBurstSizeKey),
				},
				MaxProcessingMsgsPerNode: v.GetUint64(InboundThrottlerMaxProcessingMsgsPerNodeKey),
			},

			OutboundMsgThrottlerConfig: throttling.MsgByteThrottlerConfig{
				AtLargeAllocSize:    v.GetUint64(OutboundThrottlerAtLargeAllocSizeKey),
				VdrAllocSize:        v.GetUint64(OutboundThrottlerVdrAllocSizeKey),
				NodeMaxAtLargeBytes: v.GetUint64(OutboundThrottlerNodeMaxAtLargeBytesKey),
			},
		},

		HealthConfig: network.HealthConfig{
			MaxTimeSinceMsgSent:          v.GetDuration(NetworkHealthMaxTimeSinceMsgSentKey),
			MaxTimeSinceMsgReceived:      v.GetDuration(NetworkHealthMaxTimeSinceMsgReceivedKey),
			MaxPortionSendQueueBytesFull: v.GetFloat64(NetworkHealthMaxPortionSendQueueFillKey),
			MinConnectedPeers:            v.GetUint(NetworkHealthMinPeersKey),
			MaxSendFailRate:              v.GetFloat64(NetworkHealthMaxSendFailRateKey),
			SendFailRateHalflife:         halflife,
		},

		DialerConfig: dialer.Config{
			ThrottleRps:       v.GetUint32(OutboundConnectionThrottlingRps),
			ConnectionTimeout: v.GetDuration(OutboundConnectionTimeout),
		},

		TimeoutConfig: network.TimeoutConfig{
			PingPongTimeout:      v.GetDuration(NetworkPingTimeoutKey),
			ReadHandshakeTimeout: v.GetDuration(NetworkReadHandshakeTimeoutKey),
		},

		PeerListGossipConfig: network.PeerListGossipConfig{
			PeerListNumValidatorIPs:        v.GetUint32(NetworkPeerListNumValidatorIPsKey),
			PeerListValidatorGossipSize:    v.GetUint32(NetworkPeerListValidatorGossipSizeKey),
			PeerListNonValidatorGossipSize: v.GetUint32(NetworkPeerListNonValidatorGossipSizeKey),
			PeerListPeersGossipSize:        v.GetUint32(NetworkPeerListPeersGossipSizeKey),
			PeerListGossipFreq:             v.GetDuration(NetworkPeerListGossipFreqKey),
		},

		DelayConfig: network.DelayConfig{
			MaxReconnectDelay:     v.GetDuration(NetworkMaxReconnectDelayKey),
			InitialReconnectDelay: v.GetDuration(NetworkInitialReconnectDelayKey),
		},

		MaxClockDifference:           v.GetDuration(NetworkMaxClockDifferenceKey),
		CompressionEnabled:           v.GetBool(NetworkCompressionEnabledKey),
		PingFrequency:                v.GetDuration(NetworkPingFrequencyKey),
		AllowPrivateIPs:              v.GetBool(NetworkAllowPrivateIPsKey),
		UptimeMetricFreq:             v.GetDuration(UptimeMetricFreqKey),
		MaximumInboundMessageTimeout: v.GetDuration(NetworkMaximumInboundTimeoutKey),

		RequireValidatorToConnect: v.GetBool(NetworkRequireValidatorToConnectKey),
		PeerReadBufferSize:        int(v.GetUint(NetworkPeerReadBufferSizeKey)),
		PeerWriteBufferSize:       int(v.GetUint(NetworkPeerWriteBufferSizeKey)),
	}

	switch {
	case config.HealthConfig.MaxTimeSinceMsgSent < 0:
		return network.Config{}, fmt.Errorf("%s must be >= 0", NetworkHealthMaxTimeSinceMsgSentKey)
	case config.HealthConfig.MaxTimeSinceMsgReceived < 0:
		return network.Config{}, fmt.Errorf("%s must be >= 0", NetworkHealthMaxTimeSinceMsgReceivedKey)
	case config.HealthConfig.MaxSendFailRate < 0 || config.HealthConfig.MaxSendFailRate > 1:
		return network.Config{}, fmt.Errorf("%s must be in [0,1]", NetworkHealthMaxSendFailRateKey)
	case config.HealthConfig.MaxPortionSendQueueBytesFull < 0 || config.HealthConfig.MaxPortionSendQueueBytesFull > 1:
		return network.Config{}, fmt.Errorf("%s must be in [0,1]", NetworkHealthMaxPortionSendQueueFillKey)
	case config.DialerConfig.ConnectionTimeout < 0:
		return network.Config{}, fmt.Errorf("%q must be >= 0", OutboundConnectionTimeout)
	case config.PeerListGossipFreq < 0:
		return network.Config{}, fmt.Errorf("%s must be >= 0", NetworkPeerListGossipFreqKey)
	case config.MaxReconnectDelay < 0:
		return network.Config{}, fmt.Errorf("%s must be >= 0", NetworkMaxReconnectDelayKey)
	case config.InitialReconnectDelay < 0:
		return network.Config{}, fmt.Errorf("%s must be >= 0", NetworkInitialReconnectDelayKey)
	case config.MaxReconnectDelay < config.InitialReconnectDelay:
		return network.Config{}, fmt.Errorf("%s must be >= %s", NetworkMaxReconnectDelayKey, NetworkInitialReconnectDelayKey)
	case config.PingPongTimeout < 0:
		return network.Config{}, fmt.Errorf("%s must be >= 0", NetworkPingTimeoutKey)
	case config.PingFrequency < 0:
		return network.Config{}, fmt.Errorf("%s must be >= 0", NetworkPingFrequencyKey)
	case config.PingPongTimeout <= config.PingFrequency:
		return network.Config{}, fmt.Errorf("%s must be > %s", NetworkPingTimeoutKey, NetworkPingFrequencyKey)
	case config.ReadHandshakeTimeout < 0:
		return network.Config{}, fmt.Errorf("%s must be >= 0", NetworkReadHandshakeTimeoutKey)
	case config.MaxClockDifference < 0:
		return network.Config{}, fmt.Errorf("%s must be >= 0", NetworkMaxClockDifferenceKey)
	}

	return config, nil
}

func getBenchlistConfig(v *viper.Viper, alpha, k int) (benchlist.Config, error) {
	config := benchlist.Config{
		Threshold:              v.GetInt(BenchlistFailThresholdKey),
		Duration:               v.GetDuration(BenchlistDurationKey),
		MinimumFailingDuration: v.GetDuration(BenchlistMinFailingDurationKey),
		MaxPortion:             (1.0 - (float64(alpha) / float64(k))) / 3.0,
	}
	switch {
	case config.Duration < 0:
		return benchlist.Config{}, fmt.Errorf("%q must be >= 0", BenchlistDurationKey)
	case config.MinimumFailingDuration < 0:
		return benchlist.Config{}, fmt.Errorf("%q must be >= 0", BenchlistMinFailingDurationKey)
	}
	return config, nil
}

func getBootstrapConfig(v *viper.Viper, networkID uint32) (node.BootstrapConfig, error) {
	config := node.BootstrapConfig{
		RetryBootstrap:                          v.GetBool(RetryBootstrapKey),
		RetryBootstrapWarnFrequency:             v.GetInt(RetryBootstrapWarnFrequencyKey),
		BootstrapBeaconConnectionTimeout:        v.GetDuration(BootstrapBeaconConnectionTimeoutKey),
		BootstrapMaxTimeGetAncestors:            v.GetDuration(BootstrapMaxTimeGetAncestorsKey),
		BootstrapAncestorsMaxContainersSent:     int(v.GetUint(BootstrapAncestorsMaxContainersSentKey)),
		BootstrapAncestorsMaxContainersReceived: int(v.GetUint(BootstrapAncestorsMaxContainersReceivedKey)),
	}

	ipsSet := v.IsSet(BootstrapIPsKey)
	idsSet := v.IsSet(BootstrapIDsKey)
	if ipsSet && !idsSet {
		return node.BootstrapConfig{}, fmt.Errorf("set %q but didn't set %q", BootstrapIPsKey, BootstrapIDsKey)
	}
	if !ipsSet && idsSet {
		return node.BootstrapConfig{}, fmt.Errorf("set %q but didn't set %q", BootstrapIDsKey, BootstrapIPsKey)
	}

	bootstrapIPs, bootstrapIDs := genesis.SampleBeacons(networkID, 5)
	if ipsSet {
		bootstrapIPs = strings.Split(v.GetString(BootstrapIPsKey), ",")
	}
	for _, ip := range bootstrapIPs {
		if ip == "" {
			continue
		}
		addr, err := utils.ToIPDesc(ip)
		if err != nil {
			return node.BootstrapConfig{}, fmt.Errorf("couldn't parse bootstrap ip %s: %w", ip, err)
		}
		config.BootstrapIPs = append(config.BootstrapIPs, addr)
	}

	if idsSet {
		bootstrapIDs = strings.Split(v.GetString(BootstrapIDsKey), ",")
	}
	for _, id := range bootstrapIDs {
		if id == "" {
			continue
		}
		nodeID, err := ids.ShortFromPrefixedString(id, constants.NodeIDPrefix)
		if err != nil {
			return node.BootstrapConfig{}, fmt.Errorf("couldn't parse bootstrap peer id: %w", err)
		}
		config.BootstrapIDs = append(config.BootstrapIDs, nodeID)
	}

	lenIPs := len(config.BootstrapIPs)
	lenIDs := len(config.BootstrapIDs)
	if lenIPs != lenIDs {
		return node.BootstrapConfig{}, fmt.Errorf("expected the number of bootstrapIPs (%d) to match the number of bootstrapIDs (%d)", lenIPs, lenIDs)
	}

	return config, nil
}

func getIPConfig(v *viper.Viper) (node.IPConfig, error) {
	config := node.IPConfig{}
	// Resolves our public IP, or does nothing
	config.DynamicPublicIPResolver = dynamicip.NewResolver(v.GetString(DynamicPublicIPResolverKey))
	config.DynamicUpdateDuration = v.GetDuration(DynamicUpdateDurationKey)
	if config.DynamicUpdateDuration < 0 {
		return node.IPConfig{}, fmt.Errorf("%q must be <= 0", DynamicUpdateDurationKey)
	}

	var (
		ip  net.IP
		err error
	)
	publicIP := v.GetString(PublicIPKey)
	switch {
	case config.DynamicPublicIPResolver.IsResolver():
		// User specified to use dynamic IP resolution; don't use NAT traversal
		config.Nat = nat.NewNoRouter()
		ip, err = dynamicip.FetchExternalIP(config.DynamicPublicIPResolver)
		if err != nil {
			return node.IPConfig{}, fmt.Errorf("dynamic ip address fetch failed: %w", err)
		}
	case publicIP == "":
		// User didn't specify a public IP to use; try with NAT traversal
		config.AttemptedNATTraversal = true
		config.Nat = nat.GetRouter()
		ip, err = config.Nat.ExternalIP()
		if err != nil {
			ip = net.IPv4zero // Couldn't get my IP...set to 0.0.0.0
		}
	default:
		// User specified a public IP to use; don't use NAT
		config.Nat = nat.NewNoRouter()
		ip = net.ParseIP(publicIP)
	}
	if ip == nil {
		return node.IPConfig{}, fmt.Errorf("invalid IP Address %s", publicIP)
	}

	stakingPort := uint16(v.GetUint(StakingPortKey))
	config.IP = utils.NewDynamicIPDesc(ip, stakingPort)
	return config, nil
}

func getProfilerConfig(v *viper.Viper) (profiler.Config, error) {
	config := profiler.Config{
		Dir:         os.ExpandEnv(v.GetString(ProfileDirKey)),
		Enabled:     v.GetBool(ProfileContinuousEnabledKey),
		Freq:        v.GetDuration(ProfileContinuousFreqKey),
		MaxNumFiles: v.GetInt(ProfileContinuousMaxFilesKey),
	}
	if config.Freq < 0 {
		return profiler.Config{}, fmt.Errorf("%s must be >= 0", ProfileContinuousFreqKey)
	}
	return config, nil
}

func getStakingTLSCertFromFlag(v *viper.Viper) (tls.Certificate, error) {
	stakingKeyRawContent := v.GetString(StakingKeyContentKey)
	stakingKeyContent, err := base64.StdEncoding.DecodeString(stakingKeyRawContent)
	if err != nil {
		return tls.Certificate{}, fmt.Errorf("unable to decode base64 content: %w", err)
	}

	stakingCertRawContent := v.GetString(StakingCertContentKey)
	stakingCertContent, err := base64.StdEncoding.DecodeString(stakingCertRawContent)
	if err != nil {
		return tls.Certificate{}, fmt.Errorf("unable to decode base64 content: %w", err)
	}

	cert, err := staking.LoadTLSCertFromBytes(stakingKeyContent, stakingCertContent)
	if err != nil {
		return tls.Certificate{}, fmt.Errorf("failed creating cert: %w", err)
	}

	return *cert, nil
}

func getStakingTLSCertFromFile(v *viper.Viper) (tls.Certificate, error) {
	// Parse the staking key/cert paths and expand environment variables
	stakingKeyPath := os.ExpandEnv(v.GetString(StakingKeyPathKey))
	stakingCertPath := os.ExpandEnv(v.GetString(StakingCertPathKey))

	// If staking key/cert locations are specified but not found, error
	if v.IsSet(StakingKeyPathKey) || v.IsSet(StakingCertPathKey) {
		if _, err := os.Stat(stakingKeyPath); os.IsNotExist(err) {
			return tls.Certificate{}, fmt.Errorf("couldn't find staking key at %s", stakingKeyPath)
		} else if _, err := os.Stat(stakingCertPath); os.IsNotExist(err) {
			return tls.Certificate{}, fmt.Errorf("couldn't find staking certificate at %s", stakingCertPath)
		}
	} else {
		// Create the staking key/cert if [stakingKeyPath] and [stakingCertPath] don't exist
		if err := staking.InitNodeStakingKeyPair(stakingKeyPath, stakingCertPath); err != nil {
			return tls.Certificate{}, fmt.Errorf("couldn't generate staking key/cert: %w", err)
		}
	}

	// Load and parse the staking key/cert
	cert, err := staking.LoadTLSCertFromFiles(stakingKeyPath, stakingCertPath)
	if err != nil {
		return tls.Certificate{}, fmt.Errorf("couldn't read staking certificate: %w", err)
	}
	return *cert, nil
}

func getStakingTLSCert(v *viper.Viper) (tls.Certificate, error) {
	if v.GetBool(StakingEphemeralCertEnabledKey) {
		// Use an ephemeral staking key/cert
		cert, err := staking.NewTLSCert()
		if err != nil {
			return tls.Certificate{}, fmt.Errorf("couldn't generate ephemeral staking key/cert: %w", err)
		}
		return *cert, nil
	}

	switch {
	case v.IsSet(StakingKeyContentKey) && !v.IsSet(StakingCertContentKey):
		return tls.Certificate{}, errStakingCertContentUnset
	case !v.IsSet(StakingKeyContentKey) && v.IsSet(StakingCertContentKey):
		return tls.Certificate{}, errStakingKeyContentUnset
	case v.IsSet(StakingKeyContentKey) && v.IsSet(StakingCertContentKey):
		return getStakingTLSCertFromFlag(v)
	default:
		return getStakingTLSCertFromFile(v)
	}
}

func getStakingConfig(v *viper.Viper, networkID uint32) (node.StakingConfig, error) {
	config := node.StakingConfig{
		EnableStaking:         v.GetBool(StakingEnabledKey),
		DisabledStakingWeight: v.GetUint64(StakingDisabledWeightKey),
		StakingKeyPath:        os.ExpandEnv(v.GetString(StakingKeyPathKey)),
		StakingCertPath:       os.ExpandEnv(v.GetString(StakingCertPathKey)),
	}
	if !config.EnableStaking && config.DisabledStakingWeight == 0 {
		return node.StakingConfig{}, errInvalidStakerWeights
	}

	if !config.EnableStaking && (networkID == constants.MainnetID || networkID == constants.TestID) {
		return node.StakingConfig{}, errStakingDisableOnPublicNetwork
	}

	var err error
	config.StakingTLSCert, err = getStakingTLSCert(v)
	if err != nil {
		return node.StakingConfig{}, err
	}
	if networkID != constants.MainnetID && networkID != constants.TestID {
		config.UptimeRequirement = v.GetFloat64(UptimeRequirementKey)
		config.MinValidatorStake = v.GetUint64(MinValidatorStakeKey)
		config.MaxValidatorStake = v.GetUint64(MaxValidatorStakeKey)
		config.MinNominatorStake = v.GetUint64(MinNominatorStakeKey)
		config.MinStakeDuration = v.GetDuration(MinStakeDurationKey)
		config.MaxStakeDuration = v.GetDuration(MaxStakeDurationKey)
		config.RewardConfig.MaxConsumptionRate = v.GetUint64(StakeMaxConsumptionRateKey)
		config.RewardConfig.MinConsumptionRate = v.GetUint64(StakeMinConsumptionRateKey)
		config.RewardConfig.MintingPeriod = v.GetDuration(StakeMintingPeriodKey)
		// config.RewardConfig.SupplyCap = v.GetUint64(StakeSupplyCapKey)
		config.RewardConfig.SupplyCap, _ = uint128.FromString(v.GetString(StakeSupplyCapKey))
		config.MinNominationFee = v.GetUint32(MinNominatorFeeKey)
		switch {
		case config.UptimeRequirement < 0 || config.UptimeRequirement > 1:
			return node.StakingConfig{}, errInvalidUptimeRequirement
		case config.MinValidatorStake > config.MaxValidatorStake:
			return node.StakingConfig{}, errMinValidatorStakeAboveMax
		case config.MinNominationFee > 1_000_000:
			return node.StakingConfig{}, errInvalidNominationFee
		case config.MinStakeDuration <= 0:
			return node.StakingConfig{}, errInvalidMinStakeDuration
		case config.MaxStakeDuration < config.MinStakeDuration:
			return node.StakingConfig{}, errMinStakeDurationAboveMax
		case config.RewardConfig.MaxConsumptionRate < config.RewardConfig.MinConsumptionRate:
			return node.StakingConfig{}, errStakeMaxConsumptionBelowMin
		case config.RewardConfig.MintingPeriod < config.MaxStakeDuration:
			return node.StakingConfig{}, errStakeMintingPeriodBelowMin
		}
	} else {
		config.StakingConfig = genesis.GetStakingConfig(networkID)
	}
	return config, nil
}

func getTxFeeConfig(v *viper.Viper, networkID uint32) genesis.TxFeeConfig {
	if networkID != constants.MainnetID && networkID != constants.TestID {
		return genesis.TxFeeConfig{
			TxFee:                 v.GetUint64(TxFeeKey),
			CreateAssetTxFee:      v.GetUint64(CreateAssetTxFeeKey),
			CreateAllychainTxFee:  v.GetUint64(CreateAllychainTxFeeKey),
			CreateBlockchainTxFee: v.GetUint64(CreateBlockchainTxFeeKey),
		}
	}
	return genesis.GetTxFeeConfig(networkID)
}

func getGenesisData(v *viper.Viper, networkID uint32) ([]byte, ids.ID, error) {
	// try first loading genesis content directly from flag/env-var
	if v.IsSet(GenesisConfigContentKey) {
		genesisData := v.GetString(GenesisConfigContentKey)
		return genesis.FromFlag(networkID, genesisData)
	}

	// if content is not specified go for the file
	if v.IsSet(GenesisConfigFileKey) {
		genesisFileName := os.ExpandEnv(v.GetString(GenesisConfigFileKey))
		return genesis.FromFile(networkID, genesisFileName)
	}

	// finally if file is not specified/readable go for the predefined config
	config := genesis.GetConfig(networkID)
	return genesis.FromConfig(config)
}

func getWhitelistedAllychains(v *viper.Viper) (ids.Set, error) {
	whitelistedAllychainIDs := ids.Set{}
	for _, allychain := range strings.Split(v.GetString(WhitelistedAllychainsKey), ",") {
		if allychain == "" {
			continue
		}
		allychainID, err := ids.FromString(allychain)
		if err != nil {
			return nil, fmt.Errorf("couldn't parse allychainID %q: %w", allychain, err)
		}
		if allychainID == constants.PrimaryNetworkID {
			return nil, errCannotWhitelistPrimaryNetwork
		}
		whitelistedAllychainIDs.Add(allychainID)
	}
	return whitelistedAllychainIDs, nil
}

func getDatabaseConfig(v *viper.Viper, networkID uint32) (node.DatabaseConfig, error) {
	var (
		configBytes []byte
		err         error
	)
	if v.IsSet(DBConfigContentKey) {
		dbConfigContent := v.GetString(DBConfigContentKey)
		configBytes, err = base64.StdEncoding.DecodeString(dbConfigContent)
		if err != nil {
			return node.DatabaseConfig{}, fmt.Errorf("unable to decode base64 content: %w", err)
		}
	} else if v.IsSet(DBConfigFileKey) {
		path := os.ExpandEnv(v.GetString(DBConfigFileKey))
		configBytes, err = os.ReadFile(path)
		if err != nil {
			return node.DatabaseConfig{}, err
		}
	}

	return node.DatabaseConfig{
		Name: v.GetString(DBTypeKey),
		Path: filepath.Join(
			os.ExpandEnv(v.GetString(DBPathKey)),
			constants.NetworkName(networkID),
		),
		Config: configBytes,
	}, nil
}

func getVMAliases(v *viper.Viper) (map[ids.ID][]string, error) {
	var fileBytes []byte
	if v.IsSet(VMAliasesContentKey) {
		var err error
		aliasFlagContent := v.GetString(VMAliasesContentKey)
		fileBytes, err = base64.StdEncoding.DecodeString(aliasFlagContent)
		if err != nil {
			return nil, fmt.Errorf("unable to decode base64 content: %w", err)
		}
	} else {
		aliasFilePath := filepath.Clean(v.GetString(VMAliasesFileKey))
		exists, err := storage.FileExists(aliasFilePath)
		if err != nil {
			return nil, err
		}

		if !exists {
			if v.IsSet(VMAliasesFileKey) {
				return nil, fmt.Errorf("vm alias file does not exist in %v", aliasFilePath)
			}
			return nil, nil
		}

		fileBytes, err = os.ReadFile(aliasFilePath)
		if err != nil {
			return nil, err
		}
	}

	vmAliasMap := make(map[ids.ID][]string)
	if err := json.Unmarshal(fileBytes, &vmAliasMap); err != nil {
		return nil, fmt.Errorf("problem unmarshaling vmAliases: %w", err)
	}
	return vmAliasMap, nil
}

func getVMManager(v *viper.Viper) (vms.Manager, error) {
	vmAliases, err := getVMAliases(v)
	if err != nil {
		return nil, err
	}

	manager := vms.NewManager()
	for vmID, aliases := range vmAliases {
		for _, alias := range aliases {
			if err := manager.Alias(vmID, alias); err != nil {
				return nil, err
			}
		}
	}
	return manager, nil
}

// getPathFromDirKey reads flag value from viper instance and then checks the folder existence
func getPathFromDirKey(v *viper.Viper, configKey string) (string, error) {
	configDir := os.ExpandEnv(v.GetString(configKey))
	cleanPath := filepath.Clean(configDir)
	ok, err := storage.FolderExists(cleanPath)
	if err != nil {
		return "", err
	}
	if ok {
		return cleanPath, nil
	}
	if v.IsSet(configKey) {
		// user specified a config dir explicitly, but dir does not exist.
		return "", fmt.Errorf("cannot read directory: %v", cleanPath)
	}
	return "", nil
}

func getChainConfigsFromFlag(v *viper.Viper) (map[string]chains.ChainConfig, error) {
	chainConfigContentB64 := v.GetString(ChainConfigContentKey)
	chainConfigContent, err := base64.StdEncoding.DecodeString(chainConfigContentB64)
	if err != nil {
		return nil, fmt.Errorf("unable to decode base64 content: %w", err)
	}

	chainConfigs := make(map[string]chains.ChainConfig)
	if err := json.Unmarshal(chainConfigContent, &chainConfigs); err != nil {
		return nil, fmt.Errorf("could not unmarshal JSON: %w", err)
	}
	return chainConfigs, nil
}

func getChainConfigsFromDir(v *viper.Viper) (map[string]chains.ChainConfig, error) {
	chainConfigPath, err := getPathFromDirKey(v, ChainConfigDirKey)
	if err != nil {
		return nil, err
	}

	if len(chainConfigPath) == 0 {
		return make(map[string]chains.ChainConfig), nil
	}

	chainConfigs, err := readChainConfigPath(chainConfigPath)
	if err != nil {
		return nil, fmt.Errorf("couldn't read chain configs: %w", err)
	}
	return chainConfigs, nil
}

// getChainConfigs reads & puts chainConfigs to node config
func getChainConfigs(v *viper.Viper) (map[string]chains.ChainConfig, error) {
	if v.IsSet(ChainConfigContentKey) {
		return getChainConfigsFromFlag(v)
	}
	return getChainConfigsFromDir(v)
}

// ReadsChainConfigs reads chain config files from static directories and returns map with contents,
// if successful.
func readChainConfigPath(chainConfigPath string) (map[string]chains.ChainConfig, error) {
	chainDirs, err := filepath.Glob(filepath.Join(chainConfigPath, "*"))
	if err != nil {
		return nil, err
	}
	chainConfigMap := make(map[string]chains.ChainConfig)
	for _, chainDir := range chainDirs {
		dirInfo, err := os.Stat(chainDir)
		if err != nil {
			return nil, err
		}

		if !dirInfo.IsDir() {
			continue
		}

		// chainconfigdir/chainId/config.*
		configData, err := storage.ReadFileWithName(chainDir, chainConfigFileName)
		if err != nil {
			return chainConfigMap, err
		}

		// chainconfigdir/chainId/upgrade.*
		upgradeData, err := storage.ReadFileWithName(chainDir, chainUpgradeFileName)
		if err != nil {
			return chainConfigMap, err
		}

		chainConfigMap[dirInfo.Name()] = chains.ChainConfig{
			Config:  configData,
			Upgrade: upgradeData,
		}
	}
	return chainConfigMap, nil
}

func getAllychainConfigsFromFlags(v *viper.Viper, allychainIDs []ids.ID) (map[ids.ID]chains.AllychainConfig, error) {
	allychainConfigContentB64 := v.GetString(AllychainConfigContentKey)
	allychainConfigContent, err := base64.StdEncoding.DecodeString(allychainConfigContentB64)
	if err != nil {
		return nil, fmt.Errorf("unable to decode base64 content: %w", err)
	}

	// partially parse configs to be filled by defaults later
	allychainConfigs := make(map[ids.ID]json.RawMessage, len(allychainIDs))
	if err := json.Unmarshal(allychainConfigContent, &allychainConfigs); err != nil {
		return nil, fmt.Errorf("could not unmarshal JSON: %w", err)
	}

	res := make(map[ids.ID]chains.AllychainConfig)
	for _, allychainID := range allychainIDs {
		if rawAllychainConfigBytes, ok := allychainConfigs[allychainID]; ok {
			allychainConfig := defaultAllychainConfig(v)
			if err := json.Unmarshal(rawAllychainConfigBytes, &allychainConfig); err != nil {
				return nil, err
			}
			if err := allychainConfig.ConsensusParameters.Valid(); err != nil {
				return nil, err
			}
			res[allychainID] = allychainConfig
		}
	}
	return res, nil
}

// getAllychainConfigs reads AllychainConfigs to node config map
func getAllychainConfigsFromDir(v *viper.Viper, allychainIDs []ids.ID) (map[ids.ID]chains.AllychainConfig, error) {
	allychainConfigPath, err := getPathFromDirKey(v, AllychainConfigDirKey)
	if err != nil {
		return nil, err
	}
	if len(allychainConfigPath) == 0 {
		// allychain config path does not exist but not explicitly specified, so ignore it
		return make(map[ids.ID]chains.AllychainConfig), nil
	}

	allychainConfigs, err := readAllychainConfigs(allychainConfigPath, allychainIDs, defaultAllychainConfig(v))
	if err != nil {
		return nil, fmt.Errorf("couldn't read allychain configs: %w", err)
	}
	return allychainConfigs, nil
}

func getAllychainConfigs(v *viper.Viper, allychainIDs []ids.ID) (map[ids.ID]chains.AllychainConfig, error) {
	if v.IsSet(AllychainConfigContentKey) {
		return getAllychainConfigsFromFlags(v, allychainIDs)
	}
	return getAllychainConfigsFromDir(v, allychainIDs)
}

// readAllychainConfigs reads allychain config files from a path and given allychainIDs and returns a map.
func readAllychainConfigs(allychainConfigPath string, allychainIDs []ids.ID, defaultAllychainConfig chains.AllychainConfig) (map[ids.ID]chains.AllychainConfig, error) {
	allychainConfigs := make(map[ids.ID]chains.AllychainConfig)
	for _, allychainID := range allychainIDs {
		filePath := filepath.Join(allychainConfigPath, allychainID.String()+allychainConfigFileExt)
		fileInfo, err := os.Stat(filePath)
		switch {
		case errors.Is(err, os.ErrNotExist):
			// this allychain config does not exist, move to the next one
			continue
		case err != nil:
			return nil, err
		case fileInfo.IsDir():
			return nil, fmt.Errorf("%q is a directory, expected a file", fileInfo.Name())
		}

		// allychainConfigDir/allychainID.json
		file, err := os.ReadFile(filePath)
		if err != nil {
			return nil, err
		}

		configData := defaultAllychainConfig
		if err := json.Unmarshal(file, &configData); err != nil {
			return nil, err
		}
		if err := configData.ConsensusParameters.Valid(); err != nil {
			return nil, err
		}
		allychainConfigs[allychainID] = configData
	}

	return allychainConfigs, nil
}

func defaultAllychainConfig(v *viper.Viper) chains.AllychainConfig {
	return chains.AllychainConfig{
		ConsensusParameters: getConsensusConfig(v),
		ValidatorOnly:       false,
		GossipConfig:        getGossipConfig(v),
	}
}

func GetNodeConfig(v *viper.Viper, buildDir string) (node.Config, error) {
	nodeConfig := node.Config{}

	// Plugin directory defaults to [buildDir]/[pluginsDirName]
	nodeConfig.PluginDir = filepath.Join(buildDir, pluginsDirName)

	// Consensus Parameters
	nodeConfig.ConsensusParams = getConsensusConfig(v)
	if err := nodeConfig.ConsensusParams.Valid(); err != nil {
		return node.Config{}, err
	}
	nodeConfig.ConsensusShutdownTimeout = v.GetDuration(ConsensusShutdownTimeoutKey)
	if nodeConfig.ConsensusShutdownTimeout < 0 {
		return node.Config{}, fmt.Errorf("%q must be >= 0", ConsensusShutdownTimeoutKey)
	}

	// Gossiping
	nodeConfig.ConsensusGossipFrequency = v.GetDuration(ConsensusGossipFrequencyKey)
	if nodeConfig.ConsensusGossipFrequency < 0 {
		return node.Config{}, fmt.Errorf("%s must be >= 0", ConsensusGossipFrequencyKey)
	}

	var err error
	// Logging
	nodeConfig.LoggingConfig, err = getLoggingConfig(v)
	if err != nil {
		return node.Config{}, err
	}

	// Network ID
	nodeConfig.NetworkID, err = constants.NetworkID(v.GetString(NetworkNameKey))

	if err != nil {
		return node.Config{}, err
	}

	// Database
	nodeConfig.DatabaseConfig, err = getDatabaseConfig(v, nodeConfig.NetworkID)
	if err != nil {
		return node.Config{}, err
	}

	// IP configuration
	nodeConfig.IPConfig, err = getIPConfig(v)
	if err != nil {
		return node.Config{}, err
	}

	// Staking
	nodeConfig.StakingConfig, err = getStakingConfig(v, nodeConfig.NetworkID)
	if err != nil {
		return node.Config{}, err
	}

	// Whitelisted Allychains
	nodeConfig.WhitelistedAllychains, err = getWhitelistedAllychains(v)
	if err != nil {
		return node.Config{}, err
	}

	// HTTP APIs
	nodeConfig.HTTPConfig, err = getHTTPConfig(v)
	if err != nil {
		return node.Config{}, err
	}

	// Health
	nodeConfig.HealthCheckFreq = v.GetDuration(HealthCheckFreqKey)
	if nodeConfig.HealthCheckFreq < 0 {
		return node.Config{}, fmt.Errorf("%s must be positive", HealthCheckFreqKey)
	}
	// Halflife of continuous averager used in health checks
	healthCheckAveragerHalflife := v.GetDuration(HealthCheckAveragerHalflifeKey)
	if healthCheckAveragerHalflife <= 0 {
		return node.Config{}, fmt.Errorf("%s must be positive", HealthCheckAveragerHalflifeKey)
	}

	// Router
	nodeConfig.ConsensusRouter = &router.ChainRouter{}
	nodeConfig.RouterHealthConfig, err = getRouterHealthConfig(v, healthCheckAveragerHalflife)
	if err != nil {
		return node.Config{}, err
	}

	// Metrics
	nodeConfig.MeterVMEnabled = v.GetBool(MeterVMsEnabledKey)

	// Adaptive Timeout Config
	nodeConfig.AdaptiveTimeoutConfig, err = getAdaptiveTimeoutConfig(v)
	if err != nil {
		return node.Config{}, err
	}

	// Network Config
	nodeConfig.NetworkConfig, err = getNetworkConfig(v, healthCheckAveragerHalflife)
	if err != nil {
		return node.Config{}, err
	}

	nodeConfig.GossipConfig = getGossipConfig(v)

	// Benchlist
	nodeConfig.BenchlistConfig, err = getBenchlistConfig(v, nodeConfig.ConsensusParams.Alpha, nodeConfig.ConsensusParams.K)
	if err != nil {
		return node.Config{}, err
	}

	// File Descriptor Limit
	nodeConfig.FdLimit = v.GetUint64(FdLimitKey)

	// Tx Fee
	nodeConfig.TxFeeConfig = getTxFeeConfig(v, nodeConfig.NetworkID)

	// Genesis Data
	nodeConfig.GenesisBytes, nodeConfig.AxcAssetID, err = getGenesisData(v, nodeConfig.NetworkID)
	if err != nil {
		return node.Config{}, fmt.Errorf("unable to load genesis file: %w", err)
	}

	// Assertions
	nodeConfig.EnableAssertions = v.GetBool(AssertionsEnabledKey)

	// Crypto
	nodeConfig.EnableCrypto = v.GetBool(SignatureVerificationEnabledKey)

	// Bootstrap Configs
	nodeConfig.BootstrapConfig, err = getBootstrapConfig(v, nodeConfig.NetworkID)
	if err != nil {
		return node.Config{}, err
	}

	// Allychain Configs
	allychainConfigs, err := getAllychainConfigs(v, nodeConfig.WhitelistedAllychains.List())
	if err != nil {
		return node.Config{}, err
	}
	nodeConfig.AllychainConfigs = allychainConfigs

	// Chain Configs
	nodeConfig.ChainConfigs, err = getChainConfigs(v)
	if err != nil {
		return node.Config{}, err
	}

	// Profiler
	nodeConfig.ProfilerConfig, err = getProfilerConfig(v)
	if err != nil {
		return node.Config{}, err
	}

	// VM Aliases
	nodeConfig.VMManager, err = getVMManager(v)
	if err != nil {
		return node.Config{}, err
	}

	// reset proposerVM height index
	nodeConfig.ResetProposerVMHeightIndex = v.GetBool(ResetProposerVMHeightIndexKey)

	return nodeConfig, nil
}
