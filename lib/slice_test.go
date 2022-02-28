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
