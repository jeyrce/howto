package main

import (
	"fmt"
	"os"
	"path"
)

func main() {
	root, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	sqlDir := path.Join(root, "testdata/qdm/")
	dirs, err := os.ReadDir(sqlDir)
	if err != nil {
		panic(err)
	}
	for _, dir := range dirs {
		fmt.Println(dir.Name())
	}
}
