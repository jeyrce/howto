package lib

import (
	"fmt"
	"math"
	"testing"
)

// 速记结论：数值溢出问题 = 边界值 + 溢出部分 (最终值在范围内取模)
//			上界超出 = 下界 + 超出部分
//			下界超出 = 上界 - 超出部分
func TestInt8(t *testing.T) {
	var (
		a uint8 = 1
		b uint8 = 3
		d       = uint8(int(math.Pow(2, 8)))
	)
	// 无符号数，下溢出 = 2**位数 - 最终值
	fmt.Println(a - b)
	// 上溢出 = 最终值 - 2**位数
	fmt.Println(d + a)
	var (
		x int8 = 1
		y int8 = 5
		z      = int8(int(math.Pow(2, 7)))
		k      = -z
	)
	fmt.Println(x, y, z, k)
	fmt.Println(x - y)
	fmt.Println(z + y)
}
