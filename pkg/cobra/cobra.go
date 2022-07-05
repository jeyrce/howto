package main

import (
	"fmt"
	"os"
	"time"

	pbar "github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cli = &cobra.Command{
		Use:     "cli",
		Short:   "测试cobra",
		Version: "0.1.0",
	}
	ps = &cobra.Command{
		Use:     "ps",
		Short:   "查看服务状态",
		Aliases: []string{"status", "ls"},
		Run: func(cmd *cobra.Command, args []string) {
			bar := NewBar(200, "任务1进度:")
			for i := 0; i <= 100; i++ {
				bar.Add(3)
				if bar.IsFinished() {
					break
				}
				time.Sleep(time.Millisecond * 100)
			}
			bar.Close()
			time.Sleep(time.Millisecond * 200)
			bar2 := NewBar(200, "任务2进度:")
			for j := 0; j <= 200; j++ {
				bar2.Add(2)
				if bar2.IsFinished() {
					break
				}
				time.Sleep(time.Millisecond * 100)
			}
			bar2.Close()
			time.Sleep(time.Millisecond * 200)
			bar3 := NewBar(200, "任务3进度:")
			for j := 0; j <= 200; j++ {
				bar3.Add(4)
				if bar3.IsFinished() {
					break
				}
				time.Sleep(time.Millisecond * 100)
			}
			bar3.Close()
			fmt.Println("所有任务都执行完毕...")
		},
	}
)

func init() {
	viper.SetConfigFile("test.toml")
	viper.SetConfigType("toml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	cli.AddCommand(ps)
}

func main() {
	_ = cli.Execute()
}

func NewBar(max int, desc string) *pbar.ProgressBar {
	bar := pbar.NewOptions(max,
		pbar.OptionSetDescription(desc),
		pbar.OptionUseANSICodes(true),
		pbar.OptionSetWriter(os.Stderr),
		pbar.OptionFullWidth(),
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
	return bar
}
