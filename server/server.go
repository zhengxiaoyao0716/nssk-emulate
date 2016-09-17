package server

import (
	"github.com/go-macaron/binding"
	"github.com/go-macaron/pongo2"
	"gopkg.in/macaron.v1"
)

func regHandlers(m *macaron.Macaron) {
	m.Get("", Index)
	m.Get("/index.html", Index)
	m.Get("/resource/app/download", DownloadApp)

	m.Group("/api", func() {
		m.Get("/log/pull", PullLog)
		m.Post("/server/bind", binding.Bind(BindServerData{}), BindServer)
		m.Get("/connect/list", ListConnect)
		m.Group("/message", func() {
			m.Post("/send", binding.Bind(SendMessageData{}), SendMessage)
			m.Post("/recive", binding.Bind(ReciveMessageData{}), ReciveMessage)
			m.Get("/pull", PullMessage)
			m.Get("", func(ctx *macaron.Context) (int, string) {
				return MakeResp(ctx, []map[string]string{
					map[string]string{"url": "./send", "method": "POST", "desc": "发送消息"},
					map[string]string{"url": "./recive", "method": "POST", "desc": "接受消息"},
					map[string]string{"url": "./pull", "method": "GET", "desc": "拉取消息"},
				})
			})
		})
		m.Get("/all/pull", PullAll)
		m.Get("", func(ctx *macaron.Context) (int, string) {
			return MakeResp(ctx, []map[string]string{
				map[string]string{"url": "./log/pull", "method": "GET", "desc": "拉取日志"},
				map[string]string{"url": "./server/bind", "method": "POST", "desc": "绑定服务器"},
				map[string]string{"url": "./connect/list", "method": "GET", "desc": "列出连接信息"},
				map[string]string{"url": "./message", "method": "GET", "desc": "消息接口相关帮助"},
				map[string]string{"url": "./all/pull", "method": "GET", "desc": "拉取全部数据"},
			})
		})
	})

	m.Group("/api/a/connect", func() {
		m.Post("/create", binding.Bind(CreateConnectData{}), CreateConnect)
		m.Post("/reply-verify", binding.Bind(ReplyVerifyConnectData{}), ReplyVerifyConnect)
		m.Get("", func(ctx *macaron.Context) (int, string) {
			return MakeResp(ctx, []map[string]string{
				map[string]string{"url": "./create", "method": "POST", "desc": "发起建立连接请求"},
				map[string]string{"url": "./verify/reply", "method": "POST", "desc": "应答验证连接请求"},
			})
		})
	})
	m.Group("/api/b/connect", func() {
		m.Post("/reply-create", binding.Bind(ReplyCreateConnectData{}), ReplyCreateConnect)
		m.Post("/verify", binding.Bind(VerifyConnectData{}), VerifyConnect)
		m.Get("", func(ctx *macaron.Context) (int, string) {
			return MakeResp(ctx, []map[string]string{
				map[string]string{"url": "./create/reply", "method": "POST", "desc": "应答创建连接请求"},
				map[string]string{"url": "./verify", "method": "POST", "desc": "发起验证连接请求"},
			})
		})
	})
	m.Group("/api/s/user", func() {
		m.Get("/list", ListUser)
		m.Post("/join", binding.Bind(JoinUserData{}), JoinUser)
		m.Post("/connect", binding.Bind(ConnectUserData{}), ConnectUser)
		m.Get("", func(ctx *macaron.Context) (int, string) {
			return MakeResp(ctx, []map[string]string{
				map[string]string{"url": "./list", "method": "GET", "desc": "列出已接入服务的用户"},
				map[string]string{"url": "./join", "method": "POST", "desc": "接入服务"},
				map[string]string{"url": "./connect", "method": "POST", "desc": "与其它用户建立通讯连接"},
			})
		})
	})
}

// Run 运行server
func Run(host string, port int) {
	m := macaron.Classic()
	m.Use(pongo2.Pongoer(pongo2.Options{
		Directory: "html",
	}))
	m.Use(macaron.Static("html/static", macaron.StaticOptions{Prefix: "static"}))

	regHandlers(m)

	m.Run(host, port)
}
