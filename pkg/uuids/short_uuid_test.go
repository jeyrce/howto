package uuids

import (
	"testing"

	"github.com/lithammer/shortuuid"
)

func TestShortUUID(t *testing.T) {
	i := 999
	for i > 0 {
		// 必须57字符
		u := shortuuid.NewWithAlphabet("1234567890abcdefghijklmnopqrstuvwxyz")
		t.Log(u, len(u))
		i--
	}
}
