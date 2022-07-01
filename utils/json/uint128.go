// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package json

import (
	"github.com/axiacoin/axia-network-v2/utils/uint128"
)

type Uint128 uint128.Uint128

func (u Uint128) MarshalJSON() ([]byte, error) {
	return []byte("\"" + uint128.Uint128{Lo: u.Lo, Hi: u.Hi}.String() + "\""), nil
}

func (u *Uint128) UnmarshalJSON(b []byte) error {
	str := string(b)
	if str == Null {
		return nil
	}
	if len(str) >= 2 {
		if lastIndex := len(str) - 1; str[0] == '"' && str[lastIndex] == '"' {
			str = str[1:lastIndex]
		}
	}
	val, err := uint128.FromString(str)
	*u = Uint128(val)
	return err
}
