package server

import (
	"github.com/zhengxiaoyao0716/nssk-emulate/nssk"
	"gopkg.in/macaron.v1"
)

// PullLog 拉取日志
func PullLog(ctx *macaron.Context) (int, string) {
	return MakeResp(ctx, nssk.PullLog())
}

// BindServerData 请求数据
type BindServerData struct {
	Address string `binding:"Required"`
}

// BindServer 绑定服务器
func BindServer(data BindServerData, ctx *macaron.Context) (int, string) {
	PushCache("connects", map[string]interface{}{})
	PushCache("messages", map[string]interface{}{})
	PushCache("master", "")
	PushCache("address", data.Address)
	nssk.CleanUser()
	nssk.CleanLog()
	PushCache("secret", nssk.JoinUser(data.Address))
	return MakeResp(ctx, nssk.PullLog())
}

// PullAll 拉取全部数据
func PullAll(ctx *macaron.Context) (int, string) {
	body := map[string]interface{}{}
	body["logs"] = nssk.PullLog()
	body["connects"] = GetMapCache("connects")
	body["messages"] = GetMapCache("messages")
	return MakeResp(ctx, body)
}
