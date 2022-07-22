package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 从原库读取
	rows, err := source.Query("SELECT * FROM test WHERE 1=1;")
	if err != nil {
		return
	}
	defer rows.Close()
	var object any
	for rows.Next() {
		err := rows.Scan(&object) // scan 变量数量必须和 字段数量保持一致
		if err != nil {
			panic(err)
		}
		fmt.Println(object)
	}
	// 目标端组合SQL
	target.Exec("INSERT INTO test VALUES ()")
}
