package main

import (
	"fmt"
	"os"
	"time"

	pbar "github.com/schollz/progressbar/v3"
)

func main() {
	bar := pbar.NewOptions(300,
		pbar.OptionSetDescription("任务进度"),
		pbar.OptionUseANSICodes(true),
		pbar.OptionSetWriter(os.Stderr),
		pbar.OptionSetWidth(10),
		pbar.OptionThrottle(65*time.Millisecond),
		pbar.OptionShowCount(),
		pbar.OptionShowIts(),
		pbar.OptionOnCompletion(func() {
			fmt.Printf("\n")
		}),
		pbar.OptionSpinnerType(14),
		pbar.OptionFullWidth(),
	)
	bar.RenderBlank()
	for i := 0; i <= 300; i++ {
		bar.Add(1)
		time.Sleep(time.Millisecond * 100)
	}
	bar.Clear()
}
