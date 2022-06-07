package main

import (
	"fmt"

	"github.com/spf13/cobra"
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
			fmt.Println("OK")
		},
	}
)

func init() {
	cli.AddCommand(ps)
}

func main() {
	_ = cli.Execute()
}
