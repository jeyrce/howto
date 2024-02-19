package main

import (
	"fmt"
	"net/url"
	"reflect"
	"testing"
)

// StructToUrlValues 通过自定义的 query 标签将结构体转换为url中的查询字符串
func StructToUrlValues(param interface{}) (values url.Values) {
	var (
		t   = reflect.TypeOf(param)
		v   = reflect.ValueOf(param)
		num = t.NumField()
	)
	values = make(url.Values)
	for i := 0; i < num; i++ {
		field := t.Field(i)
		name := field.Tag.Get("query")
		if len(name) < 1 {
			name = field.Name
		}
		values.Set(name, fmt.Sprint(v.Field(i)))
	}
	return values
}

func TestStructToUrlValues(t *testing.T) {
	var params = struct {
		Name  string `query:"name"`
		Test  string
		Age   int      `query:"num"`
		Array []string `query:"array"`
	}{
		Name:  "jeyrcelu",
		Test:  "ttt",
		Age:   22,
		Array: []string{"1", "2", "3"},
	}
	v := StructToUrlValues(params)
	t.Log(v.Encode())
}
