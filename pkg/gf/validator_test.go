package main

import (
	"context"
	"testing"

	"github.com/gogf/gf/v2/frame/g"
)

type User struct {
	Name string `json:"Name" v:"required|length:3,5#必须填写姓名"`
	Age  int64  `json:"Age" v:"int@required|min:1|max:140"`
}

type Param struct {
	Action string `v:"required|length:1,6"`
	Nested struct {
		Account User  `json:"Account"`
		Ext     g.Map `json:"Ext"`
	} `json:"Nested"`
}

func TestValidateNestedStruct(t *testing.T) {
	var (
		p1 = Param{
			Action: "test",
			Nested: struct {
				Account User  `json:"Account"`
				Ext     g.Map `json:"Ext"`
			}(struct {
				Account User
				Ext     g.Map
			}{Account: User{
				Name: "test",
				Age:  27,
			}, Ext: g.Map{}}),
		}
		p2 = Param{
			Action: "xxx",
			Nested: struct {
				Account User  `json:"Account"`
				Ext     g.Map `json:"Ext"`
			}(struct {
				Account User
				Ext     g.Map
			}{Account: User{
				Name: "jeyrcelu",
				Age:  0,
			}, Ext: g.Map{}}),
		}
	)
	t.Log(g.Validator().Data(p1).Run(context.TODO()))
	t.Log(g.Validator().Data(p2).Run(context.TODO()))
}

type testData struct {
	Name   string
	Age    int
	Nested User `json:"Nested"`
}
