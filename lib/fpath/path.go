package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	var (
		x = "/root/x/"
		y = "../"
		z = "../test"
	)
	fmt.Println(filepath.Abs(x))
	fmt.Println(filepath.Abs(y))
	fmt.Println(filepath.Abs(z))

	var (
		ss     = []string{"x", "y", "z"}
		marked = false
	)
	for _, s := range ss {
		if marked {
			fmt.Println("跳过")
		} else {
			fmt.Println("--------")
		}
		fmt.Println(s)
		marked = true
	}

}
