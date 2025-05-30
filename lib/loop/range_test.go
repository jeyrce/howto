package main

import (
	"testing"

	"github.com/gogf/gf/util/gconv"
)

type batch struct {
	Param param    `json:"Param"`
	Ips   []string `json:"Ips"`
}

type param struct {
	Name string
	Age  int
	Ip   string
}

func TestRangeEqu(t *testing.T) {
	b := batch{
		Param: param{
			Name: "test",
			Age:  20,
		},
		Ips: []string{"127.0.0.1", "127.0.0.2"},
	}
	sub := new(param)
	err := gconv.Struct(b.Param, sub)
	if err != nil {
		t.Fatal(err)
	}
	for _, ip := range b.Ips {
		sub.Ip = ip
		t.Log(sub)
	}
}
