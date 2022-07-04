package main

import (
	"testing"
	"time"

	"github.com/schollz/progressbar/v3"
)

func Show(total int64, ch chan int, timeout time.Duration) {
	bar := progressbar.Default(total, "任务进度:")
	for {
		if bar.IsFinished() {
			return
		}
		select {
		case num := <-ch:
			bar.Add(num)
		default:
		}
	}
}

func TestShow(t *testing.T) {
	ch := make(chan int, 0)
	go Show(300, ch, time.Minute*5)
	for i := 0; i <= 300; i++ {
		ch <- 1
	}
	time.Sleep(time.Second)
}
