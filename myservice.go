package main

import (
	"github.com/go-macaron/renders"
	"github.com/sirupsen/logrus"
	_ "github.com/sirupsen/logrus"
	"gopkg.in/macaron.v1"
	"myservice/controllers"
	"myservice/controllers/api"
	"myservice/models"
)

func main() {

	uu, err := models.GetUserByName("test1")
	if err != nil {
		logrus.Error(err)
	}
	logrus.Println(uu[0].Name)

	m := macaron.New()
	m.Use(macaron.Recovery())
	m.Use(macaron.Static("public"))
	m.Use(func(ctx *macaron.Context) {
		ctx.Data["AppName"] = "Myservice"
		ctx.Data["AppVer"] = "0.0.1"
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
