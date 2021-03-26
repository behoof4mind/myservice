package main

import (
	"github.com/go-macaron/renders"
	_ "github.com/sirupsen/logrus"
	"gopkg.in/macaron.v1"
	"myservice/controllers"
	"myservice/controllers/api"
)

func main() {

	m := macaron.New()
	m.Use(macaron.Recovery())
	m.Use(macaron.Static("public"))
	m.Use(func(ctx *macaron.Context) {
		ctx.Data["AppName"] = "Myservice"
		ctx.Data["AppVer"] = "0.1.0"
	})
	m.Use(renders.Renderer(
		renders.Options{
			Directory:       "templates",
			Extensions:      []string{".tmpl", ".html"},
			Charset:         "UTF-8",
			IndentJSON:      true,
			IndentXML:       true,
			HTMLContentType: "text/html",
		}))
	m.Get("/", controllers.IndexPageHandler)
	m.Group("/api/v1/users", func() {
		m.Combo("").
			Get(api.GetUser).
			Post(api.AddUser)
	})
	m.Run()
}
