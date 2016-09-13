package server

import "gopkg.in/macaron.v1"

// Index 主页
func Index(ctx *macaron.Context) {
	ctx.Data["master"] = GetStrCache("master")
	ctx.Data["address"] = GetStrCache("address")
	ctx.Data["isMaster"] = GetCache("secret") == nil
	ctx.HTML(200, "index")
}
