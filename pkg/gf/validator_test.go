package main

import (
	"context"
	"testing"

	"github.com/gogf/gf/v2/encoding/gjson"
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

type s struct {
	SS []string `v:"array|max-length:3"`
}

func TestArrayLengthValidate(t *testing.T) {
	var data = []s{
		{
			SS: []string{"a", "b"},
		},
		{
			SS: []string{"a", "b", "c", "1"},
		},
	}

	for _, v := range data {
		if err := g.Validator().Data(v).Run(context.Background()); err != nil {
			t.Error(err)
		} else {
			t.Log(v)
		}
	}
}

// FreeDisk 设备上的空闲磁盘列表
type FreeDisk struct {
	Serial      string   `dc:"序列号" example:"PHAX410100TP7P6DGN"`
	Path        string   `dc:"盘符" example:"/dev/nvme23n1"`
	Size        string   `dc:"容量" example:"7T"`
	Tran        string   `dc:"协议" example:"nvme"`
	MountPoints []string `dc:"挂载点"`
	MountPoint  string   `dc:"挂载点"`
	Mountpoint  string

	// 转换后的字段
	SN string `dc:"磁盘SN" example:"/dev/disk/by-id/nvme-sn-PHAX410100TP7P6DGN"`
}

func TestGJsonStr(t *testing.T) {
	var ss = `{
   "blockdevices": [
      {
         "serial": "PHAX410303FM7P6DGN",
         "path": "/dev/nvme0n1",
         "size": "7T",
         "tran": "nvme",
         "mountpoint": null
      },{
         "serial": "PHAX410303DN7P6DGN",
         "path": "/dev/nvme1n1",
         "size": "7T",
         "tran": "nvme",
         "mountpoint": null
      },{
         "serial": "PHAX410303DR7P6DGN",
         "path": "/dev/nvme2n1",
         "size": "7T",
         "tran": "nvme",
         "mountpoint": null
      },{
         "serial": "PHAX410303LF7P6DGN",
         "path": "/dev/nvme3n1",
         "size": "7T",
         "tran": "nvme",
         "mountpoint": null
      },{
         "serial": "PHAX410302XS7P6DGN",
         "path": "/dev/nvme4n1",
         "size": "7T",
         "tran": "nvme",
         "mountpoint": null
      },{
         "serial": "PHAX4103059X7P6DGN",
         "path": "/dev/nvme5n1",
         "size": "7T",
         "tran": "nvme",
         "mountpoint": null
      },{
         "serial": "PHAX410305F17P6DGN",
         "path": "/dev/nvme6n1",
         "size": "7T",
         "tran": "nvme",
         "mountpoint": null
      },{
         "serial": "PHAX4103044G7P6DGN",
         "path": "/dev/nvme7n1",
         "size": "7T",
         "tran": "nvme",
         "mountpoint": null
      },{
         "serial": "PHAX410609RR7P6DGN",
         "path": "/dev/nvme8n1",
         "size": "7T",
         "tran": "nvme",
         "mountpoint": null
      },{
         "serial": "PHAX411007NE7P6DGN",
         "path": "/dev/nvme9n1",
         "size": "7T",
         "tran": "nvme",
         "mountpoint": null
      },{
         "serial": "PHAX411007SY7P6DGN",
         "path": "/dev/nvme10n1",
         "size": "7T",
         "tran": "nvme",
         "mountpoint": null
      },{
         "serial": "PHAX410304087P6DGN",
         "path": "/dev/nvme11n1",
         "size": "7T",
         "tran": "nvme",
         "mountpoint": null
      },{
         "serial": "PHAX411006X67P6DGN",
         "path": "/dev/nvme12n1",
         "size": "7T",
         "tran": "nvme",
         "mountpoint": null
      },{
         "serial": "PHAX411003TD7P6DGN",
         "path": "/dev/nvme13n1",
         "size": "7T",
         "tran": "nvme",
         "mountpoint": null
      },{
         "serial": "PHAX4111008G7P6DGN",
         "path": "/dev/nvme14n1",
         "size": "7T",
         "tran": "nvme",
         "mountpoint": null
      },{
         "serial": "PHAX411007TG7P6DGN",
         "path": "/dev/nvme15n1",
         "size": "7T",
         "tran": "nvme",
         "mountpoint": null
      },{
         "serial": "PHAX4111008U7P6DGN",
         "path": "/dev/nvme16n1",
         "size": "7T",
         "tran": "nvme",
         "mountpoint": null
      },{
         "serial": "PHAX410303M47P6DGN",
         "path": "/dev/nvme17n1",
         "size": "7T",
         "tran": "nvme",
         "mountpoint": null
      },{
         "serial": "PHAX411000L87P6DGN",
         "path": "/dev/nvme18n1",
         "size": "7T",
         "tran": "nvme",
         "mountpoint": null
      },{
         "serial": "PHAX4111003E7P6DGN",
         "path": "/dev/nvme19n1",
         "size": "7T",
         "tran": "nvme",
         "mountpoint": null
      },{
         "serial": "PHAX410302X67P6DGN",
         "path": "/dev/nvme20n1",
         "size": "7T",
         "tran": "nvme",
         "mountpoint": null
      },{
         "serial": "PHAX4103042G7P6DGN",
         "path": "/dev/nvme21n1",
         "size": "7T",
         "tran": "nvme",
         "mountpoint": null
      },{
         "serial": "PHAX4100035T7P6DGN",
         "path": "/dev/nvme22n1",
         "size": "7T",
         "tran": "nvme",
         "mountpoint": "/mnt"
      },{
         "serial": "PHAX410100TP7P6DGN",
         "path": "/dev/nvme23n1",
         "size": "7T",
         "tran": "nvme",
         "mountpoint": null
      }
   ]
}`
	var (
		j    = gjson.New(ss)
		data = make([]*FreeDisk, 0)
	)
	err := j.Get("blockdevices").Scan(&data)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(j.MustToJsonIndentString())

	for _, d := range data {
		t.Log(d.Path, d.MountPoint, d.Mountpoint)
	}
}
