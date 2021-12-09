package pkg

import (
	"testing"

	gouuid "github.com/satori/go.uuid"
)

func TestGoUUID(t *testing.T) {
	i := 999
	for i > 0 {
		v4 := gouuid.NewV4().String()
		t.Log(v4, len(v4))
		i--
	}
}
