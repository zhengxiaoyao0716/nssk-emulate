package server

import (
	"log"

	"gopkg.in/macaron.v1"
)

// MakeErr 构造失败响应
func MakeErr(ctx *macaron.Context, code int, reas interface{}) (int, string) {
	resp, err := ctx.JSONString(&map[string]interface{}{"flag": false, "reas": reas})
	if err != nil {
		log.Println(err)
	}
	return code, resp
}

// MakeResp 构造成功响应
func MakeResp(ctx *macaron.Context, data interface{}) (int, string) {
	resp, err := ctx.JSONString(&map[string]interface{}{"flag": true, "data": data})
	if err != nil {
		log.Println(err)
	}
	return 200, resp
}
