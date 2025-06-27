package main

import (
	"context"

	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gctx"
)

func init() {

}

func main() {
	var (
		jeyrcelu = gcmd.Command{
			Name:          "jeyrcelu",
			Usage:         "jeyrcelu [command]",
			Brief:         "jeyrce's private cli tool",
			Description:   "",
			Arguments:     nil,
			Func:          nil,
			FuncWithValue: nil,
			HelpFunc:      nil,
			Examples:      "",
			Additional:    "",
			Strict:        false,
			CaseSensitive: false,
			Config:        "",
		}
		show = gcmd.Command{
			Name:        "show",
			Usage:       "show",
			Brief:       "show my information",
			Description: "",
			Arguments:   nil,
			Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
				parser.GetArg()
			},
			FuncWithValue: nil,
			HelpFunc:      nil,
			Examples:      "",
			Additional:    "",
			Strict:        false,
			CaseSensitive: false,
			Config:        "",
		}
	)

	jeyrcelu.Run(gctx.GetInitCtx())
}
