package uuids

import (
	"testing"

	"github.com/segmentio/ksuid"
)

// 27 位
func TestKSUID(t *testing.T) {
	for i := 0; i < 999; i++ {
		k := ksuid.New()
		t.Log(k.String())
	}
}
