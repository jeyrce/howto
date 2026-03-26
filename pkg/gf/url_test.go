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

func TestURLPath(t *testing.T) {
	parse, err := url.Parse("http://mydbops-dev.woa.com/newstore/workflow/")
	if err != nil {
		t.Fatal(err)
	}
	parse.Path = ""
	t.Log(parse.Hostname(), parse.Port(), parse.Scheme, parse.Path)

	t.Log(parse.JoinPath("x").String())
	t.Log(parse.JoinPath("y").String())

	var clusterUrl = parse.JoinPath("/newstore/cluster/")
	clusterUrl.RawQuery = url.Values{
		"RegionKey": []string{"ncdb-80012"},
		"ZoneKey":   []string{"80012az"},
	}.Encode()

	t.Log(clusterUrl.String())
}

// Action 将一些公共的操作封装到这里
type Action struct {
	name string
}

// Name 返回任务名称
func (a *Action) Name() string {
	if a.name == "" {
		var t = reflect.TypeOf(a)
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}
		a.name = t.Name()
	}
	return a.name
}

type XAction struct {
	*Action
}

func TestActionName(t *testing.T) {
	var x = XAction{&Action{name: "x"}}
	t.Log(x.Name())
}

func TestUUID(t *testing.T) {
	var str = "replace_disk_7b0a4448-4614-4863-a32c-d706cdce5439.result"
	t.Log(str[13:49])
}
