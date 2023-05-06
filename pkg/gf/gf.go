package main

import (
	"fmt"
	"github.com/gogf/gf/encoding/gjson"
	"io"
	"os"
)

func main() {
	filename := "/Users/jeyrcelu/workspace/go/src/github.com/jeyrce/howto/pkg/gf/tdsql_hosts"
	//all := gcfg.New(filename).GetJson(".")
	//if all == nil {
	//	fmt.Println("不合法的文件")
	//}
	//hosts := gcfg.New("tdsql_hosts").GetJson(".")
	//fmt.Println(all.Map())
	//fmt.Println("---------")
	//fmt.Println(hosts.Map())
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	all, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	content, err := gjson.LoadContentType("ini", all)
	if err != nil {
		panic(err)
	}
	fmt.Println(content.GetJson("tdsql_allmacforcheck").GetString("tdsql_mac1 ansible_ssh_host"))
}
