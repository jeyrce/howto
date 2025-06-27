package main

import (
	"testing"

	"github.com/gogf/gf/util/gconv"
)

type S1 struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	SS   []string
}

type S2 struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	SS   []string
	Ext  string
}

func TestS12S2(t *testing.T) {
	var (
		s1 = S1{
			Name: "jeyrce",
			Age:  19,
			SS: []string{
				"god",
				"bless",
				"you",
			},
		}
		s2 = new(S2)
	)

	err := gconv.Struct(s1, s2)

	s2.Ext = "test copy struct"
	if err != nil {
		t.Fatal(err)
	}

	t.Log(s2)

}
