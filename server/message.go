package server

import (
	"errors"

	"github.com/zhengxiaoyao0716/nssk-emulate/nssk"
	"github.com/zhengxiaoyao0716/util/requests"
	"gopkg.in/macaron.v1"
)

// SendMessageData 请求数据
type SendMessageData struct {
	Address string `json:"address" binding:"Required"` // 发信对象的通讯地址
	Message string `binding:"Required"`                // 消息内容
}

// SendMessage 发送消息
func SendMessage(data SendMessageData, ctx *macaron.Context) (int, string) {
	kab := GetStrCache(data.Address)
	if kab == "" {
		return MakeErr(ctx, 403, errors.New("发送失败，连接未建立"))
	}
	body := map[string]interface{}{}
	body["address"] = GetStrCache("address")
	body["message"] = nssk.Encrypt(data.Message, GetStrCache(data.Address))
	_, err := requests.Post(data.Address+"/api/message/recive", body)
	if err != nil {
		return MakeErr(ctx, 403, err)
	}
	return MakeResp(ctx, "fin")
}

// ReciveMessageData 请求数据
type ReciveMessageData struct {
	Address string `json:"address" binding:"Required"` // 发信对象的通讯地址
	Message string `binding:"Required"`                // 消息内容（加密后的）
}

// ReciveMessage 接收消息
func ReciveMessage(data ReciveMessageData, ctx *macaron.Context) (int, string) {
	message, err := nssk.Decrypt(data.Message, GetStrCache(data.Address))
	if err != nil {
		return MakeErr(ctx, 403, err)
	}
	messageMap := GetMapCache("messages")
	messages := messageMap[data.Address].([]interface{})
	messages = append(messages, message)
	messageMap[data.Address] = messages
	return MakeResp(ctx, "fin")
}

// PullMessage 拉取消息
func PullMessage(ctx *macaron.Context) (int, string) {
	return MakeResp(ctx, GetMapCache("messages"))
}

func init() {
	PushCache("messages", map[string]interface{}{})
}
