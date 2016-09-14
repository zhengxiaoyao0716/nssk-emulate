package server

import "gopkg.in/macaron.v1"

// Index 主页
func Index(ctx *macaron.Context) {
	master := GetStrCache("master")
	ctx.Data["master"] = master
	ctx.Data["address"] = GetStrCache("address")
	ctx.Data["isMaster"] = master == ""
	ctx.HTML(200, "index")
}
