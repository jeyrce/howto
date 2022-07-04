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
			bar := NewBar(300, "1任务进度:")
			for i := 0; i <= 300; i++ {
				bar.Add(2)
				time.Sleep(time.Millisecond * 100)
			}
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
	bar.Clear()
	return bar
}
