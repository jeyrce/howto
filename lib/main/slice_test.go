package main

import (
	"errors"
	"testing"
)

func TestEmptySlice(t *testing.T) {
	var ss = make([]error, 0)
	if errors.Join(ss...) != nil {
		t.Log("not nil")
	}
	t.Log("nil")
	ss = append(ss, errors.New("test"))
	ss = append(ss, errors.New("test"))
	if errors.Join(ss...) != nil {
		t.Log("error occurred")
	}
}
