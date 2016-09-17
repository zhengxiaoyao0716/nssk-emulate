package server

import "gopkg.in/macaron.v1"

// Index 主页
func Index(ctx *macaron.Context) {
	address := GetStrCache("address")
	ctx.Data["address"] = GetStrCache("address")
	master := GetStrCache("master")
	if master == "" {
		ctx.Data["isMaster"] = true
		ctx.Data["master"] = address
	} else {
		ctx.Data["isMaster"] = false
		ctx.Data["master"] = master
	}
	ctx.HTML(200, "index")
}
