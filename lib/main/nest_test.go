package main

import (
	"fmt"
	"testing"
)

type Father struct {
}

func (f *Father) Parse() {
	fmt.Println("Father's parse")
}

func (f *Father) Call() {
	fmt.Println("Father's call")
	f.Parse()
}

type Child struct {
	Father
}

func (c *Child) Parse() {
	fmt.Println("Child's parse")
}

func TestNest(t *testing.T) {
	var f, c = Father{}, Child{}
	f.Call()
	c.Call()
}
