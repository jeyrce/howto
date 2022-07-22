package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestInsert(t *testing.T) {
	for i := 1; i <= 99999; i++ {
		sql := fmt.Sprintf("INSERT INTO test VALUES (NULL,name='%s',x='ccc',y=5, NULL);", strconv.Itoa(i)+"xxx")
		fmt.Println(sql)
		exec, err := source.Exec(sql)
		if err != nil {
			panic(err)
		}
		fmt.Println(exec.RowsAffected())
	}
}
