package main

import (
	"testing"

	"github.com/gogf/gf/encoding/gjson"
)

type Data struct {
	Name string `json:"Name"`
	Age  int    `json:"Age"`
	X    struct {
		Info  string `json:"Info"`
		Score int    `json:"Score"`
	} `json:"X"`
}

func TestGJson(t *testing.T) {
	var n = Data{
		Name: "jeyrcelu",
		Age:  22,
		X: struct {
			Info  string `json:"Info"`
			Score int    `json:"Score"`
		}(struct {
			Info  string
			Score int
		}{
			Info:  "test",
			Score: 333,
		}),
	}
	json := gjson.New(n)
	t.Log(json.MustToJsonString())
}

func TestGJsonMarshal(t *testing.T) {
	j := gjson.New(Data{
		Name: "jeyrcelu",
		Age:  12,
		X: struct {
			Info  string `json:"Info"`
			Score int    `json:"Score"`
		}(struct {
			Info  string
			Score int
		}{Info: "cccc", Score: 66}),
	})
	t.Log(j.MustToJsonString())
}
