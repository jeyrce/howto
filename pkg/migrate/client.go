package main

import (
	"database/sql"
	"fmt"
	"os"
	"sync"
	"time"
)

var (
	source *sql.DB
	target *sql.DB
	once   sync.Once
)

func Init() {
	passwd, ok := os.LookupEnv("MYSQL_ROOT_PASSWORD")
	if !ok {
		passwd = "123456"
	}
	sourceConn, err := sql.Open("mysql", fmt.Sprintf("root:%s@tcp(42.192.55.32:13306)/source", passwd))
	if err != nil {
		fmt.Println("error ->")
		panic(err)
	}
	sourceConn.SetConnMaxLifetime(time.Minute * 30)
	sourceConn.SetMaxOpenConns(10)
	sourceConn.SetMaxIdleConns(10)
	source = sourceConn

	targetConn, err := sql.Open("mysql", fmt.Sprintf("root:%s@tcp(42.192.55.32:13306)/target", passwd))
	if err != nil {
		panic(err)
	}
	targetConn.SetConnMaxLifetime(time.Minute * 30)
	targetConn.SetMaxOpenConns(10)
	targetConn.SetMaxIdleConns(10)
	target = targetConn
}

func init() {
	once.Do(Init)
}
