package main

import (
	"time"

	"github.com/briandowns/spinner"
)

func main() {
	i := len(spinner.CharSets)
	for j := 0; j <= i; j++ {
		s := spinner.New(spinner.CharSets[j], 100*time.Millisecond)
		s.Start()
		time.Sleep(2 * time.Second)
		s.Stop()
	}
}
