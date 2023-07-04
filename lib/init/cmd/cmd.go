package main

import (
	"fmt"
	_ "github.com/jeyrce/howto/lib/init"
	v "github.com/jeyrce/howto/lib/init/var"
)

func init() {
	fmt.Println("这是main的init")
}

func main() {
	fmt.Println("main主函数")
	fmt.Println(v.TestVar)
}
