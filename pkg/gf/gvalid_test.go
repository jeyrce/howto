package main

import (
	"context"
	"testing"

	"github.com/gogf/gf/v2/util/gvalid"
)

type Req struct {
	Action    string `json:"Action" v:"required|length:1,64#必须携带操作码"`
	ApiModule string `json:"ApiModule" v:"required|length:1,64#必须携带模块代码"`
	RequestId string `json:"RequestId" v:"required|length:1,64#必须携带请求ID"`
	AppId     int64  `json:"AppId"`
	Uin       int64  `json:"Uin"`
}

func TestGValidStruct(t *testing.T) {
	var reqs = []Req{
		{
			Action:    "cxcx",
			ApiModule: "1",
			RequestId: "33",
			AppId:     0,
			Uin:       0,
		},
		{
			Action:    "",
			ApiModule: "c",
			RequestId: "1",
			AppId:     0,
			Uin:       0,
		},
	}
	for _, r := range reqs {
		if err := gvalid.New().Data(r).Run(context.TODO()); err != nil {
			t.Fatal(err)
		}
	}
}
