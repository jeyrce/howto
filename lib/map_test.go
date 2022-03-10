package lib

import (
	"encoding/json"
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

func TestMapNew(t *testing.T) {
	var m = make(map[string]string)
	m["x"] = "xx"
	t.Log(m["x"])
}

type xx struct {
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	Age  uint   `json:"age,omitempty" yaml:"age,omitempty"`
}

type yy struct {
	Name string `json:"name"`
	Age  uint   `json:"age"`
}

func TestOmitempty(t *testing.T) {
	var (
		emptyX = xx{
			Name: "",
			Age:  0,
		}
		emptyY = yy{
			Name: "",
			Age:  0,
		}
	)
	marshal, err := json.Marshal(emptyX)
	if err != nil {
		return
	}
	bytes, err := json.Marshal(emptyY)
	if err != nil {
		return
	}
	t.Log(string(marshal)) // {}
	t.Log(string(bytes))   // {"name":"","age":0}
}
