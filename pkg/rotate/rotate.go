package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

var (
	pwd string
)

func main() {
	_, file, _, _ := runtime.Caller(0)
	pwd = filepath.Dir(file)
	fmt.Println(pwd)
	fmt.Println(os.Getwd())
}
