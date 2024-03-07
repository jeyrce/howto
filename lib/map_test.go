package lib

import (
	"encoding/json"
	"testing"

	"github.com/gogf/gf/encoding/gjson"
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

type Handler interface {
	Type() string
	Code() string
	Version() string
}

type TestHandler struct {
	t string
	c string
	v string
}

func (t TestHandler) Type() string {
	return t.t
}

func (t TestHandler) Code() string {
	return t.c
}

func (t TestHandler) Version() string {
	return t.v
}

func Registry(handlers ...Handler) {
	for _, h := range handlers {
		if _, ok := m[h.Type()]; !ok {
			m[h.Type()] = make(map[string]map[string]Handler)
		}
		if _, ok := m[h.Type()][h.Code()]; !ok {
			m[h.Type()][h.Code()] = make(map[string]Handler)
		}
		m[h.Type()][h.Code()][h.Version()] = h
	}
}

var m = make(map[string]map[string]map[string]Handler)

func TestMapRegistry(t *testing.T) {
	Registry(
		&TestHandler{
			t: "x",
			c: "x.cc",
			v: "1.0.0",
		},
		&TestHandler{
			t: "x",
			c: "x.cc",
			v: "1.0.2",
		},
		&TestHandler{
			t: "y",
			c: "y.cc",
			v: "1.0.0",
		},
		&TestHandler{
			t: "y",
			c: "y.dd",
			v: "1.0.0",
		},
	)
	t.Log(gjson.New(m).MustToJsonString())
}

func TestMapCount(t *testing.T) {
	var data = []struct {
		A string
		B int
	}{
		{A: "1", B: 1},
		{A: "1", B: 2},
		{A: "x", B: 1},
		{A: "x", B: 1},
		{A: "-", B: 1},
	}
	m := make(map[string]int)
	for _, d := range data {
		m[d.A]++
	}
	t.Log(m)
}
