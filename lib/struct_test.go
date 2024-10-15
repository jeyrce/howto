package lib

import (
	"testing"
)

type Meta struct {
	Name string
	Num  int
}

type Info struct {
	X string
}

type MultipleStruct struct {
	Meta
	Info
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

func TestNilStruct(t *testing.T) {
	var (
		i    = new(Info)
		meta = Meta{
			Name: "jeyrcelu",
			Num:  20,
		}
		M = MultipleStruct{
			Meta: meta,
			Info: *i,
		}
	)
	t.Log(M)
}
