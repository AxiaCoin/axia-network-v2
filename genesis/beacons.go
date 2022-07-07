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
			"52.0.104.155:9651",
			"54.85.164.125:9651",
			"54.174.170.42:9651",
			"34.198.56.39:9651",
			"54.198.58.78:9651",
			"44.199.134.13:9651",
			"44.208.58.148:9651",
			"34.202.117.69:9651",
			"43.205.21.247:9651",
			"43.204.135.131:9651",
			"3.108.49.224:9651",
			"43.205.43.85:9651",
			"13.235.176.106:9651",
			"43.205.43.152:9651",
			"3.6.99.202:9651",
			"3.109.25.43:9651",
			"13.49.207.244:9651",
			"13.51.42.243:9651",
			"13.48.188.47:9651",
			"16.16.62.131:9651",
			"13.51.62.60:9651",
			"13.48.246.168:9651",
			"16.170.156.224:9651",
		}
	case constants.TestID:
		return []string{
			"54.205.204.194",
			"44.207.193.221",
			"35.170.105.82",
			"44.208.120.143",
			"52.23.71.222",
			"18.214.242.123",
			"18.210.91.149",
			"15.207.251.109",
			"43.204.88.72",
			"43.205.54.170",
			"65.1.18.105",
			"15.207.217.99",
			"43.204.170.41",
			"13.126.202.19",
			"13.49.180.209",
			"13.49.114.49",
			"16.16.50.171",
			"13.50.45.159",
			"13.50.26.246",
			"13.51.87.164",
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
			"NodeID-FVJiVBMXg69Eys8v3K3fUZXM21fW5k7Ts",
			"NodeID-8ofNV8sb3J9rwaC3HuXcPFAkNEqF8oJzV",
			"NodeID-BLaJAbw7aYCmZtYoemCBpowijcBPq7yZs",
			"NodeID-FTirsWrt7FJVTTvQejEQpJERuJDKk9BW2",
			"NodeID-Gym2UoBo3b4tp3pukZzDjxJefNc5XE8Cw",
			"NodeID-AkZUoGaJUWrfGUBtk7Wrsyvu5DZKVkxKQ",
			"NodeID-FuzyvQMfFRhAWiiQLvcDVErbNoA2czuYp",
			"NodeID-Ag6dycckf8nwo2EYrvfnuw4oAw8xTz6P3",
			"NodeID-PUWVULUUXozYBvUQwoPxQFMXT13MqRBAK",
			"NodeID-MbjtKYV1EEbRdaK1d396g54VrVnTXet5Q",
			"NodeID-FeFgdAG3v8YJidF4xR8E3zSHUEuNRtBfB",
			"NodeID-2eAGdTFKwdi63NDvADr99qoDANmY11uun",
			"NodeID-pbDBA6Gfy9ShgscRm83kmJpVECqyh2EV",
			"NodeID-KbQdsqnna4KwQfsRksNMah1ZbpQgvqtrX",
			"NodeID-Ne4NyLV1WAyyXoRAzSSSTbc7pU8jbFTez",
			"NodeID-Mx3reKpoVipQVUzH41qx5gJZmNURgvFJD",
			"NodeID-CksSwGj4tSAyyRVrFavs6p9farwwtEC1b",
			"NodeID-QHHAyHWvLEJsawMTYgZfAT66erXRVqtHu",
			"NodeID-21zDZTDjnKcLE8yMy1SXsC5oYvJq1m1Ti",
			"NodeID-Q1KAsFkpTVszwbAtANydsRTLj9jjs8PAP",
			"NodeID-HPz5qwnbvsiz4oqmsf7zr8NeWsmbSnAep",
			"NodeID-G3FjbZ7rSrBqL11BpZG89RJPXkP2cUaiL",
			"NodeID-5wFDMxCVVpvZip4oUrMyunwRde1aU9Vmb",
		}
	case constants.TestID:
		return []string{
			"NodeID-3qztCoJtH2ZJmfo5Gz5TdQu9mte6tV55D",
			"NodeID-GDfPMF6z9bvXrUsheNPSF5SoKVvN8PKFM",
			"NodeID-KeVFurTahZLqjcM8SdSRKQQiHce13G7rR",
			"NodeID-6c33F4uHVbX5H8AhbDuU29KHeD9LoUrAG",
			"NodeID-G5R2YN4KAxAdAdDr6F63JkP77mZv18jki",
			"NodeID-D3aq1xWbyXmuNgG66Qv57Tfx1Yk347676",
			"NodeID-EfjvM16g2Rs5qUUA4up1Jwyie8C9k7B1K",
			"NodeID-JXDrdUhsar2CdYnwq4YZ7py1nmJGNTLFM",
			"NodeID-PrBAuMhdfwEVsjCoLxmyqCowrLFzC9s6e",
			"NodeID-DJkTyVjyecGxseQ82GeL9b2UL4WWDUJWt",
			"NodeID-9P6EXhWpNwKEpJnMb7rVuAqoCcUBve59y",
			"NodeID-NQjPxpZtYgAqYBGtvoXAekxtb1jNPKt6J",
			"NodeID-MvQtjin5gjDV8wxaFjhvTxyfpas71spGs",
			"NodeID-8w7kvFNpeX4vU5v1PxGs8U7GJLpNnFYzr",
			"NodeID-7W5zm4r77CTnm1AQ2RK8kV5aTQN1WrcJf",
			"NodeID-GgtrcdnMSSQXe54fxkyp9SjzvWsTzAVVX",
			"NodeID-JnRiWWWtYrwtjJNUNCpwQvqtqTHV9YYr8",
			"NodeID-821sSxnqx1MMiFyUrCmpxJoTU6rvxEZEv",
			"NodeID-21Z3egMvBQrXPnGvA7Qp2FgJ3UGUzEEco",
			"NodeID-9DXNudaTqHTCA5r7K1MHjHgpujtvzzkUM",
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
