package lib

import (
	"testing"
)

func TestMap(t *testing.T) {
	m := map[string]string{"x": "author", "x yz dd": "clsfd"}
	t.Log(m)
}
