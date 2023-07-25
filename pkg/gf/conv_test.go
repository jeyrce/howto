package main

import (
	"github.com/gogf/gf/util/gconv"
	"testing"
)

func TestGconv(t *testing.T) {
	convert := gconv.Convert("/dd", "int")
	t.Log(convert)
}
