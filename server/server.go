package server

import (
	"github.com/go-macaron/binding"
	"github.com/go-macaron/pongo2"
	"gopkg.in/macaron.v1"
)

func regHandlers(m *macaron.Macaron) {
	m.Get("/view", Index)
	m.Get("/view/index.html", Index)

	m.Group("/api", func() {
	})
	m.Group("/api-master", func() {
		m.Get("/user/list", ListUser)
		m.Post("/user/join", binding.Bind(JoinUserData{}), JoinUser)
		m.Get("", func(ctx *macaron.Context) (int, string) {
			return MakeResp(ctx, []map[string]string{
				map[string]string{"url": "./user/list", "method": "GET"},
				map[string]string{"url": "./user/join", "method": "POST"},
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
