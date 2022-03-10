package lib

import (
	"testing"
	"time"
)

func loop() {
	time.Sleep(time.Millisecond * 50)
}

func TestLoop(t *testing.T) {
	for {
		go loop()
	}
}
