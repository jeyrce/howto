package lib

import (
	"math"
	"testing"
)

func TestFloat(t *testing.T) {
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
