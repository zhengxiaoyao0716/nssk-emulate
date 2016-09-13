package server

import (
	"errors"

	"gopkg.in/macaron.v1"
)

// SendMessageData 请求数据
type SendMessageData struct {
	Address     string `json:"address" binding:"Required"` // 发信对象的通讯地址
	Message     string `binding:"Required"`                // 消息内容
	AutoConnect bool   // 自动创建连接
}

// SendMessage 发送消息
func SendMessage(data SendMessageData, ctx *macaron.Context) (int, string) {
	kab := GetStrCache(data.Address)
	if kab == "" {
		return MakeErr(ctx, 403, errors.New("发送失败，连接未建立"))
	}
	return MakeResp(ctx, "fin")
}

// ReciveMessageData 请求数据
type ReciveMessageData struct {
	Address string `json:"address" binding:"Required"` // 发信对象的通讯地址
	Message string `binding:"Required"`                // 消息内容
}

// ReciveMessage 接收消息
func ReciveMessage(data ReciveMessageData, ctx *macaron.Context) (int, string) {
	messageMap := GetMapCache("messages")
	messages := messageMap[data.Address].([]string)
	messages = append(messages, data.Message)
	// messageMap[data.Address] = messages
	// PushCache("messages", messageMap)
	return MakeResp(ctx, "fin")
}

// PullMessage 拉取消息
func PullMessage(ctx *macaron.Context) (int, string) {
	return MakeResp(ctx, GetMapCache("messages"))
}

func init() {
	PushCache("messages", map[string]interface{}{})
}
