package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/briandowns/spinner"
)

func Show(total float64, index chan float64, timeout time.Duration) {
	// point := time.Now()
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	s := spinner.New(spinner.CharSets[36], 200*time.Millisecond)
	s.Start()
	defer s.Stop()
	for {
		select {
		case <-ticker.C:
			// fmt.Println(time.Since(point).String())
		case cur := <-index:
			if total <= cur {
				return
			}
			// fmt.Println(cur / total)
		case <-time.After(timeout):
			fmt.Println("timeout...", timeout.String())
			return
		default:
			// 
		}
	}
}

func TestShow(t *testing.T) {
	ch := make(chan float64, 1)
	go Show(100, ch, time.Second*60)
	for i := 0; i <= 99; i++ {
		ch <- float64(i + 1)
		time.Sleep(time.Millisecond * 300)
	}
}
