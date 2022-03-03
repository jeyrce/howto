package lib

import (
	"testing"
)

func TestMap(t *testing.T) {
	m := map[string]string{"x": "author", "x yz dd": "clsfd"}
	t.Log(m)
}

func TestMapKey(t *testing.T) {
	var m = map[string]string{
		"x": "ccc",
		"y": "bbb",
		"z": "ddd",
	}
	t.Log(m["a"] == "true", m["z"] == "ddd", m["y"] == "ccc")
}
