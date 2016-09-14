package server

import (
	"github.com/zhengxiaoyao0716/nssk-emulate/nssk"
	"gopkg.in/macaron.v1"
)

// CreateConnectData 请求数据
type CreateConnectData struct {
	Address string `json:"address" binding:"Required"` // 要通讯的对象的通讯地址
}

// CreateConnect 发起建立连接请求
func CreateConnect(data CreateConnectData, ctx *macaron.Context) (int, string) {
	SetState(data.Address, Creating)
	master := GetStrCache("master")
	address := GetStrCache("address")
	if master == "" {
		master = address
	}
	if err := nssk.CreateConnect(
		master,
		address,
		data.Address,
		GetStrCache("secret"),
	); err != nil {
		return MakeErr(ctx, 403, err)
	}
	return MakeResp(ctx, "fin")
}

// ReplyVerifyConnectData 请求数据
type ReplyVerifyConnectData struct {
	Address string `json:"address" binding:"Required"` // 要通讯的对象的通讯地址
	Cab     string `binding:"Required"`
}

// ReplyVerifyConnect 应答验证连接请求
func ReplyVerifyConnect(data ReplyVerifyConnectData, ctx *macaron.Context) (int, string) {
	cab, kab, err := nssk.ReplyVerifyConnect(GetStrCache("address"), data.Address, data.Cab)
	if err != nil {
		delete(GetMapCache("connects"), data.Address)
		return MakeErr(ctx, 403, err)
	}
	PushCache(data.Address, kab)
	// SetState(data.Address, Verified)
	SetState(data.Address, Connected)
	GetMapCache("messages")[data.Address] = []interface{}{}
	return MakeResp(ctx, cab)
}
