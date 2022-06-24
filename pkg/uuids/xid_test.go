package uuids

import (
	"testing"
	"time"

	"github.com/rs/xid"
)

func TestXID(t *testing.T) {
	var i = 999
	for i > 0 {
		id := xid.New().String()
		t.Log(id, len(id))
		i--
	}
}

func TestXIDL(t *testing.T) {
	j := 999
	for j > 0 {
		withTime := xid.NewWithTime(time.Now())
		id := withTime.String()
		t.Log(id, len(id))
		j--
	}
}
