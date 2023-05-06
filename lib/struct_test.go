package lib

import (
	"testing"
)

type Meta struct {
	Name string
	Num  int
}

func x() Meta {
	var y = struct {
		Meta
		X string
	}{}
	y.X = "x"
	y.Num = 5
	y.Name = "y"
	return y.Meta
}

func TestStruct(t *testing.T) {
	t.Log(x())
}
