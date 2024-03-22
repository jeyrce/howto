package lib

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	a := [5]int{1, 2, 3, 4, 5}
	test := a[3:4:4]
	fmt.Println(test[0]) // 输出4, 考察slice底层数组存储知识
}

func TestSliceIdx(t *testing.T) {
	var (
		array = []string{
			"a",
			"b",
			"c",
			"d",
		}
		stayIdx int
		array2  = make([]string, 0, 3)
	)
	for idx, item := range array {
		if item == "d" {
			stayIdx = idx
			break
		}
	}
	t.Log(array[stayIdx : stayIdx+1])
	t.Log(array2[0:1])
}

func TestNewSlice(t *testing.T) {
	var ss []string
	ss = append(ss, "xxx", "yyy")
	t.Log(ss)
}
