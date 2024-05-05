package explorerutil

import (
	tendermintCommon "github.com/tendermint/tmlibs/common"
)

func IsHex(s string) bool {
	if tendermintCommon.IsHex(s) {
		return true
	}

	if len(s) > 2 && s[0:2] != "0x" {
		var with0x = "0x" + s
		return tendermintCommon.IsHex(with0x)
	}

	return false
}
