package main

import (
    "fmt"

    kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
    addr  = kingpin.Flag("listen.addr", "地址").Default(":9090").String()
    check = kingpin.Command("check", "检查配置")
    file  = check.Arg("filename", "文件名").String()
)

func init() {
    kingpin.HelpFlag.Short('h')
    kingpin.Version("0.1.0")
    kingpin.Parse()
}

func main() {
    switch kingpin.Parse() {
    case check.FullCommand():
        fmt.Println("检查成功", *file)
    }
}
