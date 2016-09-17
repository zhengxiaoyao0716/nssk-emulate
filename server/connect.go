package server

import "gopkg.in/macaron.v1"

// ListConnect 列出连接信息
func ListConnect(ctx *macaron.Context) (int, string) {
	return MakeResp(ctx, GetMapCache("connects"))
}

// State 连接状态
type State int

const (
	// Connected 连接已建立
	Connected State = iota
	// Creating 创建连接中
	Creating
	// Verifying 验证连接中
	Verifying
)

// SetState 设置连接状态
func SetState(address string, state State) {
	GetMapCache("connects")[address] = state
}

func init() {
	PushCache("connects", map[string]interface{}{})
}
