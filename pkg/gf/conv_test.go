package main

import (
	"reflect"
	"testing"

	"github.com/gogf/gf/v2/util/gconv"

	"github.com/jeyrce/howto/pkg/etcd"
)

func TestGconv(t *testing.T) {
	convert := gconv.Convert("/dd", "int")
	t.Log(convert)
}

func TestName(t *testing.T) {
	t.Log(reflect.TypeOf(X{}).String())
}

type X etcd.Data
