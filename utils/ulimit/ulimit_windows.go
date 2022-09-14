// Copyright (C) 2019-2021, Axia Systems, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

//go:build windows
// +build windows

package ulimit

import "github.com/axiacoin/axia-network-v2/utils/logging"

const DefaultFDLimit = 16384

// Set is a no-op for windows and will warn if the default is not used.
func Set(max uint64, log logging.Logger) error {
	if max != DefaultFDLimit {
		log.Warn("fd-limit is not supported for windows")
	}
	return nil
}
