package server

import "gopkg.in/macaron.v1"

// Index 主页
func Index(ctx *macaron.Context) {
	ctx.Data["Master"] = GetStrCache("master")
	ctx.Data["Secret"] = GetStrCache("secret")
	ctx.HTML(200, "index")
}
