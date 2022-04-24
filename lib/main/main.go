package main

import (
	"fmt"
)

func test(sss []int) {
	for idx, _ := range sss {
		sss[idx] = 1
	}
}

func testMap(m map[string]string) {
	m["name"] = "jeyrce"
	m["mail"] = "jeyrce@gmail.com"
}

// map 和 slice 以及 chan 是指针类型
func main() {
	var (
		ss = make([]int, 5, 10)
		mm = map[string]string{
			"name": "xxx",
			"mail": "xxx@xxx.com",
		}
	)
	test(ss)
	testMap(mm)
	fmt.Println(ss, mm)
}
