package server

import (
	"github.com/zhengxiaoyao0716/nssk-emulate/nssk"
	"gopkg.in/macaron.v1"
)

// PullAll 拉取全部数据
func PullAll(ctx *macaron.Context) (int, string) {
	body := map[string]interface{}{}
	body["logs"] = nssk.PullLog()
	body["connects"] = GetMapCache("connects")
	body["messages"] = GetMapCache("messages")
	return MakeResp(ctx, body)
}
