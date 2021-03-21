package controllers

import (
	"github.com/go-macaron/renders"
	"gopkg.in/macaron.v1"
)

func IndexPageHandler(r renders.Render, ctx *macaron.Context) {
	ctx.Data["Title"] = "Home"
	r.HTML(200, "index.html", ctx.Data)
}
