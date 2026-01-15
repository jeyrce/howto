package main

import (
	"reflect"
	"testing"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/jeyrce/howto/pkg/etcd"
)

func TestGconv(t *testing.T) {
	convert := gconv.Convert("/dd", "int")
	t.Log(convert)
}

func TestName(t *testing.T) {
	t.Log(reflect.TypeOf(X{}).String())
}

type X etcd.Data

type StringStruct struct {
	ID string `json:"ID"`
	X  string `json:"X"`
}

type IntStruct struct {
	ID int    `json:"ID"`
	X  string `json:"X"`
}

func TestMapStringInt(t *testing.T) {
	var s = new(StringStruct)
	err := gjson.New(IntStruct{
		ID: 123,
		X:  "xxx",
	}).Scan(s)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(s)

}
