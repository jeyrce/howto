package main

import (
	"fmt"
	"testing"
)

func TestChan(t *testing.T) {
	var ch = make(chan struct{})
	ch <- struct{}{}
	select {
	case <-ch:
		fmt.Println("rcv")
	default:
		//
	}
}
