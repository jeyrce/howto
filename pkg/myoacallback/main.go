package main

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
)

func init() {
	g.SetDebug(true)
}

// Data 描述了单据变量
type Data struct {
	Key   string   `json:"key" dc:"业务系统字段标识"`
	Value []string `json:"value" dc:"字段值"`
}
type CallbackReq struct {
	ID               string            `json:"id" dc:"单据id"`
	BusinessKey      string            `json:"business_key" dc:"-"`
	Handler          string            `json:"handler" dc:"审批人id"`
	Category         string            `json:"category" dc:"业务领域"`
	ProcessName      string            `json:"process_name" dc:"流程名称"`
	ProcessInstId    string            `json:"process_inst_id" dc:"流程标识"`
	Activity         string            `json:"activity" dc:"单据节点"`
	SubmitSource     string            `json:"submit_source" dc:"审批者终端"`
	SubmitAction     string            `json:"submit_action" dc:"审批动作(agree、reject)"`
	SubmitActionName string            `json:"submit_action_name" dc:"审批动作名称(同意、驳回)"`
	SubmitOpinion    string            `json:"submit_opinion" dc:"审批意见"`
	SubmitForm       map[string]string `json:"submit_form" dc:"审批者自定义表单"`
	Data             []Data            `json:"data" dc:"业务系统传递的变量"`
	CreateTime       gtime.Time        `json:"create_time" dc:"单据创建时间"`
	SubmitTime       gtime.Time        `json:"submit_time" dc:"审批时间"`
}

func Callback() ghttp.HandlerFunc {
	return func(r *ghttp.Request) {
		var (
			param = new(CallbackReq)
			ctx   = r.GetCtx()
		)
		if err := r.GetStruct(&param); err != nil {
			g.Log().Errorf(ctx, "解析回调结果失败: %v", err)
			r.Response.WriteJsonExit(map[string]string{"err": err.Error()})
		}
		g.Log().Debugf(ctx, "回调结果: %s", gjson.New(param).MustToJsonString())
		// TODO: 根据情况更新本地单据
		r.Response.WriteJsonExit(map[string]string{"ok": "ok"})
	}
}

func main() {
	server := g.Server()
	server.SetAddr(":10086")
	server.BindHandler("/myoa/callback", Callback())
	server.Run()
}
