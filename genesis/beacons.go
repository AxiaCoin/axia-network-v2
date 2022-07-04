// Copyright (C) 2019-2021, Axia Systems, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package genesis

import (
	"github.com/axiacoin/axia-network-v2/utils/constants"
	"github.com/axiacoin/axia-network-v2/utils/sampler"
)

// getIPs returns the beacon IPs for each network
func getIPs(networkID uint32) []string {
	switch networkID {
	case constants.MainnetID:
		return []string{
			"54.94.43.49:9651",
			"52.79.47.77:9651",
			"18.229.206.191:9651",
			"3.34.221.73:9651",
			"13.244.155.170:9651",
			"13.244.47.224:9651",
			"122.248.200.212:9651",
			"52.30.9.211:9651",
			"122.248.199.127:9651",
			"18.202.190.40:9651",
			"15.206.182.45:9651",
			"15.207.11.193:9651",
			"44.226.118.72:9651",
			"54.185.87.50:9651",
			"18.158.15.12:9651",
			"3.21.38.33:9651",
			"54.93.182.129:9651",
			"3.128.138.36:9651",
			"3.104.107.241:9651",
			"3.106.25.139:9651",
			"18.162.129.129:9651",
			"18.162.161.230:9651",
			"52.47.181.114:9651",
			"15.188.9.42:9651",
		}
	case constants.TestID:
		return []string{
			// "52.0.104.155:9651",
			// "54.85.164.125:9651",
			// "54.174.170.42:9651",
			// "34.198.56.39:9651",
			// "54.198.58.78:9651",
			// "44.199.134.13:9651",
			// "44.208.58.148:9651",
			// "34.202.117.69:9651",
			// "43.205.21.247:9651",
			// "43.204.135.131:9651",
			// "3.108.49.224:9651",
			// "43.205.43.85:9651",
			// "13.235.176.106:9651",
			// "43.205.43.152:9651",
			// "3.6.99.202:9651",
			// "3.109.25.43:9651",
			// "13.49.207.244:9651",
			// "13.51.42.243:9651",
			// "13.48.188.47:9651",
			// "16.16.62.131:9651",
			// "13.51.62.60:9651",
			// "13.48.246.168:9651",
			// "16.170.156.224:9651",
		}
	case constants.LocalID:
		return []string{}
	default:
		return nil
	}
}

// getNodeIDs returns the beacon node IDs for each network
func getNodeIDs(networkID uint32) []string {
	switch networkID {
	case constants.MainnetID:
		return []string{
			"NodeID-A6onFGyJjA37EZ7kYHANMR1PFRT8NmXrF",
			"NodeID-6SwnPJLH8cWfrJ162JjZekbmzaFpjPcf",
			"NodeID-GSgaA47umS1px2ohVjodW9621Ks63xDxD",
			"NodeID-BQEo5Fy1FRKLbX51ejqDd14cuSXJKArH2",
			"NodeID-Drv1Qh7iJvW3zGBBeRnYfCzk56VCRM2GQ",
			"NodeID-DAtCoXfLT6Y83dgJ7FmQg8eR53hz37J79",
			"NodeID-FGRoKnyYKFWYFMb6Xbocf4hKuyCBENgWM",
			"NodeID-Dw7tuwxpAmcpvVGp9JzaHAR3REPoJ8f2R",
			"NodeID-4kCLS16Wy73nt1Zm54jFZsL7Msrv3UCeJ",
			"NodeID-9T7NXBFpp8LWCyc58YdKNoowDipdVKAWz",
			"NodeID-6ghBh6yof5ouMCya2n9fHzhpWouiZFVVj",
			"NodeID-HiFv1DpKXkAAfJ1NHWVqQoojjznibZXHP",
			"NodeID-Fv3t2shrpkmvLnvNzcv1rqRKbDAYFnUor",
			"NodeID-AaxT2P4uuPAHb7vAD8mNvjQ3jgyaV7tu9",
			"NodeID-kZNuQMHhydefgnwjYX1fhHMpRNAs9my1",
			"NodeID-A7GwTSd47AcDVqpTVj7YtxtjHREM33EJw",
			"NodeID-Hr78Fy8uDYiRYocRYHXp4eLCYeb8x5UuM",
			"NodeID-9CkG9MBNavnw7EVSRsuFr7ws9gascDQy3",
			"NodeID-A8jypu63CWp76STwKdqP6e9hjL675kdiG",
			"NodeID-HsBEx3L71EHWSXaE6gvk2VsNntFEZsxqc",
			"NodeID-Nr584bLpGgbCUbZFSBaBz3Xum5wpca9Ym",
			"NodeID-QKGoUvqcgormCoMj6yPw9isY7DX9H4mdd",
			"NodeID-HCw7S2TVbFPDWNBo1GnFWqJ47f9rDJtt1",
			"NodeID-FYv1Lb29SqMpywYXH7yNkcFAzRF2jvm3K",
		}
	case constants.TestID:
		return []string{
			// "NodeID-FVJiVBMXg69Eys8v3K3fUZXM21fW5k7Ts",
			// "NodeID-8ofNV8sb3J9rwaC3HuXcPFAkNEqF8oJzV",
			// "NodeID-BLaJAbw7aYCmZtYoemCBpowijcBPq7yZs",
			// "NodeID-FTirsWrt7FJVTTvQejEQpJERuJDKk9BW2",
			// "NodeID-Gym2UoBo3b4tp3pukZzDjxJefNc5XE8Cw",
			// "NodeID-AkZUoGaJUWrfGUBtk7Wrsyvu5DZKVkxKQ",
			// "NodeID-FuzyvQMfFRhAWiiQLvcDVErbNoA2czuYp",
			// "NodeID-Ag6dycckf8nwo2EYrvfnuw4oAw8xTz6P3",
			// "NodeID-PUWVULUUXozYBvUQwoPxQFMXT13MqRBAK",
			// "NodeID-MbjtKYV1EEbRdaK1d396g54VrVnTXet5Q",
			// "NodeID-FeFgdAG3v8YJidF4xR8E3zSHUEuNRtBfB",
			// "NodeID-2eAGdTFKwdi63NDvADr99qoDANmY11uun",
			// "NodeID-pbDBA6Gfy9ShgscRm83kmJpVECqyh2EV",
			// "NodeID-KbQdsqnna4KwQfsRksNMah1ZbpQgvqtrX",
			// "NodeID-Ne4NyLV1WAyyXoRAzSSSTbc7pU8jbFTez",
			// "NodeID-Mx3reKpoVipQVUzH41qx5gJZmNURgvFJD",
			// "NodeID-CksSwGj4tSAyyRVrFavs6p9farwwtEC1b",
			// "NodeID-QHHAyHWvLEJsawMTYgZfAT66erXRVqtHu",
			// "NodeID-21zDZTDjnKcLE8yMy1SXsC5oYvJq1m1Ti",
			// "NodeID-Q1KAsFkpTVszwbAtANydsRTLj9jjs8PAP",
			// "NodeID-HPz5qwnbvsiz4oqmsf7zr8NeWsmbSnAep",
			// "NodeID-G3FjbZ7rSrBqL11BpZG89RJPXkP2cUaiL",
			// "NodeID-5wFDMxCVVpvZip4oUrMyunwRde1aU9Vmb",
		}
	case constants.LocalID:
		return []string{}
	default:
		return nil
	}
}

// SampleBeacons returns the some beacons this node should connect to
func SampleBeacons(networkID uint32, count int) ([]string, []string) {
	ips := getIPs(networkID)
	ids := getNodeIDs(networkID)

	if numIPs := len(ips); numIPs < count {
		count = numIPs
	}

	sampledIPs := make([]string, 0, count)
	sampledIDs := make([]string, 0, count)

	s := sampler.NewUniform()
	_ = s.Initialize(uint64(len(ips)))
	indices, _ := s.Sample(count)
	for _, index := range indices {
		sampledIPs = append(sampledIPs, ips[int(index)])
		sampledIDs = append(sampledIDs, ids[int(index)])
	}

	return sampledIPs, sampledIDs
}
