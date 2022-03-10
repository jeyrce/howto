package main

import (
	"fmt"
	"runtime"
	"time"
)

func loop() {
	for {
		c := 10512 + 9999
		fmt.Sprintln(c)
	}
}

func main() {
	ticker := time.NewTicker(time.Second * 1)
	defer ticker.Stop()
	for i := 0; i < 999999; i++ {
		go loop()
	}
	for {
		select {
		case <-ticker.C:
			fmt.Println(runtime.NumGoroutine(), runtime.NumCPU(), runtime.NumCgoCall())
		}
	}
}
