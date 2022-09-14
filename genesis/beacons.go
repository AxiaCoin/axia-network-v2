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
			"35.174.143.117:9651",
			"44.210.54.97:9651",
			"54.151.80.204:9651",
			"52.52.48.100:9651",
			"52.215.67.182:9651",
			"52.208.4.104:9651",
			"3.75.43.198:9651",
			"3.74.60.114:9651",
			"35.75.246.44:9651",
			"35.79.183.34:9651",
			"18.143.44.37:9651",
			"18.141.14.202:9651",
		}
	case constants.TestID:
		return []string{
			"35.170.190.190:9651",
			"52.8.149.181:9651",
			"63.34.5.196:9651",
			"3.125.111.163:9651",
			"3.113.71.59:9651",
			"18.142.52.36:9651",
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
			"NodeID-H8Yey5QX1fKTrNAdVFPPwVJjixr4pnxYE",
			"NodeID-J8ce9EyXVBSzkE7D2G4RcM9bhJLDiaj4C",
			"NodeID-N4GJrrsRHyAyJ6eyhbVLbe5hQDZm6Z2xD",
			"NodeID-7ihRQ8QX98HYEguBxA8yfzfhgZabMmWnP",
			"NodeID-JCrs3zduYwg7EDXu431bN5L4vM2kTvx8B",
			"NodeID-JMmBfGTTuDoPzGzdfdH9jxZDwNJ2CajBQ",
			"NodeID-NeSHHK4SSSmqSWwvs45oqUxK5tnDEAdVE",
			"NodeID-Aod3799i2D3manPEpy3NMfx9DZosWq6wQ",
			"NodeID-KvLN8t2XnruHhpJPQroNjnuF28howzKLE",
			"NodeID-LxXuxjktNJ4vDZgH2YiBbw2cibpcdwNCC",
			"NodeID-PoEdto9WVkNetcGAqYeWc4x8pbzTWRSuL",
			"NodeID-9ZGpMrJ6xriryxDfsQcqSFpEvqe7EBHrG",
		}
	case constants.TestID:
		return []string{
			"NodeID-D8yKWBBXqpx8Qvi5jLZEUG1RURsciR8Au",
			"NodeID-G7wmfRYiRcLzsir8nSvtaw2HTYJAUjoMc",
			"NodeID-9h7tSWTDT1VgLNqtNcZi9f1Egm2xxjSM",
			"NodeID-MMN1PTM26z8VGxk1vxEScX4dcsJhwfEjj",
			"NodeID-Fjks9n8sEKVKD7rdJ1BFBWzDA1tLD3jtN",
			"NodeID-6Z4HDTEmdWFcrfDey9y9MwA1RGbYexYfE",
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
