package server

import (
	"gopkg.in/macaron.v1"

	"github.com/zhengxiaoyao0716/nssk-emulate/nssk"
)

// ListUser 列出已接入服务的用户
func ListUser(ctx *macaron.Context) (int, string) {
	ctx.Header().Add("Access-Control-Allow-Origin", "*")
	return MakeResp(ctx, nssk.ListUser())
}

// JoinUserData 请求数据
type JoinUserData struct {
	Address string `json:"address" binding:"Required"` // 发起接入请求者的通讯地址
}

// JoinUser 接入服务
func JoinUser(data JoinUserData, ctx *macaron.Context) (int, string) {
	secret := nssk.JoinUser(data.Address)
	return MakeResp(ctx, secret)
}

// ConnectUserData 请求数据
type ConnectUserData struct {
	From string `json:"from" binding:"Required"` // 发起通讯请求者的通讯地址
	To   string `binding:"Required"`             // 要通讯对象的通讯地址
	Na   string `binding:"Required"`             //随机数
}

// ConnectUser 与其它用户建立通讯连接
func ConnectUser(data ConnectUserData, ctx *macaron.Context) (int, string) {
	return MakeResp(ctx, nssk.ConnectUser(data.From, data.To, data.Na))
}
