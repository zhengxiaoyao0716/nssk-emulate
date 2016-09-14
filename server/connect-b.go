package server

import (
	"github.com/zhengxiaoyao0716/nssk-emulate/nssk"
	"gopkg.in/macaron.v1"
)

// ReplyCreateConnectData 请求数据
type ReplyCreateConnectData struct {
	Cbs string `json:"cbs" binding:"Required"`
}

// ReplyCreateConnect 应答创建连接请求
func ReplyCreateConnect(data ReplyCreateConnectData, ctx *macaron.Context) (int, string) {
	address, err := nssk.ReplyCreateConnect(data.Cbs, GetStrCache("secret"))
	if err != nil {
		delete(GetMapCache("connects"), address)
		return MakeErr(ctx, 403, err)
	}
	SetState(address, Verifying)
	return MakeResp(ctx, "fin")
}

// VerifyConnectData 请求数据
type VerifyConnectData struct {
	Address string `json:"address" binding:"Required"` // 要验证的连接对象的通讯地址
}

// VerifyConnect 发起验证连接请求
func VerifyConnect(data VerifyConnectData, ctx *macaron.Context) (int, string) {
	kab, err := nssk.VerifyConnect(data.Address, GetStrCache("address"))
	if err != nil {
		delete(GetMapCache("connects"), data.Address)
		return MakeErr(ctx, 403, err)
	}
	PushCache(data.Address, kab)
	SetState(data.Address, Connected)
	GetMapCache("messages")[data.Address] = []interface{}{}
	return MakeResp(ctx, "fin")
}
