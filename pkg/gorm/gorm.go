package main

import (
    "fmt"

    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

type Proxy struct {
    Name    string
    Role    string
    Cid     string
    ProxyID string
}

func (p Proxy) TableName() string {
    return "proxy"
}

type Instance struct {
    Name  string
    Role  string
    Price float64
    Cid   string
}

func (i Instance) TableName() string {
    return "instance"
}

type Result struct {
    Name string
    Role string
    Cid  string
}

func main() {
    db, err := gorm.Open(sqlite.Open("/tmp/test.db"), &gorm.Config{})
    if err != nil {
        panic(err)
    }
    if err := db.AutoMigrate(&Instance{}, &Proxy{}); err != nil {
        panic(err)
    }
    db = db.Debug()
    insts := []Instance{
        {Cid: "111", Name: "aaa", Role: "dev", Price: 123.5},
        {Cid: "111", Name: "bbb", Role: "dev", Price: 123.5},
        {Cid: "111", Name: "ccc", Role: "dev", Price: 123.5},
    }
    for _, inst := range insts {
        if err := db.Create(&inst).Error; err != nil {
            panic(err)
        }
    }
    ps := []Proxy{
        {Cid: "111", Name: "777", Role: "dev", ProxyID: "098"},
        {Cid: "111", Name: "888", Role: "dev", ProxyID: "098"},
        {Cid: "111", Name: "999", Role: "dev", ProxyID: "098"},
    }
    for _, p := range ps {
        if err := db.Create(&p).Error; err != nil {
            panic(err)
        }
    }
    db = db.Model(&Result{}).Where("cid = ?", "111").Table("instance").Joins(
        "JOIN proxy ON instance.cid = proxy.cid",
    )
    var result []Result
    if err := db.Find(&result).Error; err != nil {
        fmt.Println(err.Error())
    }
    for _, r := range result {
        fmt.Println(r)
    }
}
