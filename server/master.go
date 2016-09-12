package server

import (
	"gopkg.in/macaron.v1"

	"github.com/zhengxiaoyao0716/nssk-emulate/nssk"
)

// ListUser 列出已连接的用户
func ListUser(ctx *macaron.Context) (int, string) {
	return MakeResp(ctx, nssk.ListUser())
}

// JoinUserData 请求数据
type JoinUserData struct {
	Address string `json:"address" binding:"Required"` // 中间服务器请求地址
}

// JoinUser 加入节点
func JoinUser(data JoinUserData, ctx *macaron.Context) (int, string) {
	secret := nssk.JoinUser(data.Address)
	return MakeResp(ctx, secret)
}
