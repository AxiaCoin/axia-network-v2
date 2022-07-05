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
			/* "52.0.104.155:9651",
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
			"16.170.156.224:9651", */
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
			"NodeID-FVJiVBMXg69Eys8v3K3fUZXM21fW5k7Ts",
			"NodeID-8ofNV8sb3J9rwaC3HuXcPFAkNEqF8oJzV",
			"NodeID-BLaJAbw7aYCmZtYoemCBpowijcBPq7yZs",
			"NodeID-FTirsWrt7FJVTTvQejEQpJERuJDKk9BW2",
			"NodeID-Gym2UoBo3b4tp3pukZzDjxJefNc5XE8Cw",
/*			"NodeID-AkZUoGaJUWrfGUBtk7Wrsyvu5DZKVkxKQ",
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
			"NodeID-5wFDMxCVVpvZip4oUrMyunwRde1aU9Vmb", */
		}
	case constants.TestID:
		return []string{
			"NodeID-2m38qc95mhHXtrhjyGbe7r2NhniqHHJRB",
			"NodeID-JjvzhxnLHLUQ5HjVRkvG827ivbLXPwA9u",
			"NodeID-LegbVf6qaMKcsXPnLStkdc1JVktmmiDxy",
			"NodeID-HGZ8ae74J3odT8ESreAdCtdnvWG1J4X5n",
			"NodeID-CYKruAjwH1BmV3m37sXNuprbr7dGQuJwG",
			"NodeID-4KXitMCoE9p2BHA6VzXtaTxLoEjNDo2Pt",
			"NodeID-LQwRLm4cbJ7T2kxcxp4uXCU5XD8DFrE1C",
			"NodeID-4CWTbdvgXHY1CLXqQNAp22nJDo5nAmts6",
			"NodeID-4QBwET5o8kUhvt9xArhir4d3R25CtmZho",
			"NodeID-JyE4P8f4cTryNV8DCz2M81bMtGhFFHexG",
			"NodeID-EDESh4DfZFC15i613pMtWniQ9arbBZRnL",
			"NodeID-BFa1padLXBj7VHa2JYvYGzcTBPQGjPhUy",
			"NodeID-CZmZ9xpCzkWqjAyS7L4htzh5Lg6kf1k18",
			"NodeID-FesGqwKq7z5nPFHa5iwZctHE5EZV9Lpdq",
			"NodeID-84KbQHSDnojroCVY7vQ7u9Tx7pUonPaS",
			"NodeID-CTtkcXvVdhpNp6f97LEUXPwsRD3A2ZHqP",
			"NodeID-hArafGhY2HFTbwaaVh1CSCUCUCiJ2Vfb",
			"NodeID-4B4rc5vdD1758JSBYL1xyvE5NHGzz6xzH",
			"NodeID-EzGaipqomyK9UKx9DBHV6Ky3y68hoknrF",
			"NodeID-NpagUxt6KQiwPch9Sd4osv8kD1TZnkjdk",
			"NodeID-3VWnZNViBP2b56QBY7pNJSLzN2rkTyqnK",
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
