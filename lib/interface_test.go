package lib

import "testing"

type caller interface {
	Call(f func() error) error
}

type DefaultCaller struct {
	Version string `json:"Version"`
	Token   string `json:"Token"`
}

func (d *DefaultCaller) Call(f func() error) error {
	return f()
}

var _ caller = (*DefaultCaller)(nil)

func TestInterface(t *testing.T) {

}
