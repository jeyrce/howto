package main

import (
	"context"
	"os"
	"time"

	"github.com/spf13/cobra"
)

func init() {

}

type command cobra.Command

func (c *command) Execute() {

}

func main() {
	ctl := cobra.Command{
		Use:                        "",
		Aliases:                    nil,
		SuggestFor:                 nil,
		Short:                      "",
		Long:                       "",
		Example:                    "",
		ValidArgs:                  nil,
		ValidArgsFunction:          nil,
		Args:                       nil,
		ArgAliases:                 nil,
		BashCompletionFunction:     "",
		Deprecated:                 "",
		Annotations:                nil,
		Version:                    "",
		PersistentPreRun:           nil,
		PersistentPreRunE:          nil,
		PreRun:                     nil,
		PreRunE:                    nil,
		Run:                        nil,
		RunE:                       nil,
		PostRun:                    nil,
		PostRunE:                   nil,
		PersistentPostRun:          nil,
		PersistentPostRunE:         nil,
		FParseErrWhitelist:         cobra.FParseErrWhitelist{},
		CompletionOptions:          cobra.CompletionOptions{},
		TraverseChildren:           false,
		Hidden:                     false,
		SilenceErrors:              false,
		SilenceUsage:               false,
		DisableFlagParsing:         false,
		DisableAutoGenTag:          false,
		DisableFlagsInUseLine:      false,
		DisableSuggestions:         false,
		SuggestionsMinimumDistance: 0,
	}
	cobra.OnInitialize(func() {
		// do in initialize
	})
	// 自定义命令参数（而不是os.Args[1:]）
	ctl.SetArgs()

	// 自定义输出信息格式
	ctl.SetVersionTemplate()
	ctl.SetUsageTemplate()
	ctl.SetHelpTemplate()

	// 超时控制
	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Minute)
	defer cancelFunc()

	// 自定义输入输出
	ctl.SetOut(os.Stdout)
	ctl.SetErr(os.Stderr)

	ctl.GenBashCompletionFile("")

	ctl.ExecuteContext(timeout)
}
