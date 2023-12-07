package ref

import (
	"testing"
)

type handler interface {
	Code() string
	Desc() string
	Handle() func() string
}

type base struct {
	code string
	desc string
}

func (b base) Code() string {
	return b.code
}

func (b base) Desc() string {
	return b.desc
}

func (b base) Handle() func() {
	return nil
}

type Do struct {
	base
}

func (d Do) Handle() func() string {
	return func() string {
		return "do.Handle"
	}
}

func TestRefNestedStruct(t *testing.T) {
	do := Do{base{code: "1", desc: "ddd"}}
	t.Log(do.Code(), do.Desc(), do.Handle()())
}
