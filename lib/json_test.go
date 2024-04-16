package lib

import (
	"encoding/json"
	"testing"
)

type Data struct {
	ID      string `json:"ID"`
	Name    string `json:"Name"`
	Ignore  int    `json:"-"`
	Empty   string `json:"Empty,omitempty"`
	TestInt int    `json:"TestInt,omitempty"`
}

func TestJsonMarshal(t *testing.T) {
	var d = Data{
		ID:     "xxx",
		Name:   "jeyrce",
		Ignore: 5,
		Empty:  "empty",
	}
	marshal, err := json.Marshal(d)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(marshal))
}

func TestJsonUnmarshal(t *testing.T) {
	var param = Data{
		ID:      "yyy",
		Name:    "jeyrce",
		Ignore:  7,
		Empty:   "empty",
		TestInt: 0,
	}
	var d = Data{}
	marshal, err := json.Marshal(param)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(marshal))
	if err = json.Unmarshal(marshal, &d); err != nil {
		t.Fatal(err)
	}
	t.Log(d)
}
