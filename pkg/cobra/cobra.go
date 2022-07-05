package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	pbar "github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	Factory = map[string]func(){
		"1.0": func() {
			fmt.Println("v1.0")
		},
		"1.9.0-gfyh": func() {

		},
		"1.7U1": func() {

		},
	}
	cobrax = &cobra.Command{
		Use:     "cobra",
		Short:   "测试cobra",
		Version: "0.1.0",
	}
	ps = &cobra.Command{
		Use:     "ps",
		Short:   "查看服务状态",
		Aliases: []string{"status"},
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
	run = cobra.Command{
		Use:     "run",
		Aliases: []string{"exec"},
		Short:   "执行迁移",
		Example: "cobra run 1.0",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Println("请指定版本")
				return
			}
			if f, ok := Factory[args[0]]; ok {
				f()
				return
			}
			fmt.Println("不支持的版本:", args[0])
			return
		},
	}
	ls = cobra.Command{
		Use:     "ls",
		Aliases: []string{"list"},
		Short:   "列出所有支持的版本",
		Example: "cobra ls",
		Run: func(cmd *cobra.Command, args []string) {
			var versions = make([]string, 0)
			for v := range Factory {
				versions = append(versions, v)
			}
			fmt.Printf("%s\n", strings.Join(versions, ", "))
			return
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
	cobrax.AddCommand(ps)
	cobrax.AddCommand(&run)
	cobrax.AddCommand(&ls)
}

func main() {
	_ = cobrax.Execute()
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
