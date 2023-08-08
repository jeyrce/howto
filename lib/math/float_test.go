package math

import (
	"fmt"
	"github.com/gogf/gf/util/gconv"
	"math"
	"strconv"
	"testing"
)

func TestFloat(t *testing.T) {
	var (
		nums = []string{
			"1", "2", "3.1415", "5.67", "242412.45353453",
		}
	)
	for _, n := range nums {
		f := gconv.Convert(n, "float64")
		fmt.Println(strconv.ParseFloat(fmt.Sprintf("%.2f", f), 64))
	}
}

func TestFloatNan(t *testing.T) {
	var (
		count    = 0
		pageSize = 0 // 当0作为分母时将会计算出一个 Nan
	)
	x := float64(count)
	y := float64(pageSize)
	// Special cases are:
	//	Ceil(±0) = ±0
	//	Ceil(±Inf) = ±Inf
	//	Ceil(NaN) = NaN
	z := math.Ceil(x / y) // Nan
	pageTotal := int(z)
	t.Log(x, y, z, pageTotal)
	t.Log(int(math.NaN()))
}
