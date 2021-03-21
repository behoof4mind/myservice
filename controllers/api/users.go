package api

import (
	"encoding/json"
	"github.com/go-macaron/renders"
	"github.com/sirupsen/logrus"
	"gopkg.in/macaron.v1"
	"myservice/models"
)

func GetUser(r renders.Render, ctx *macaron.Context) {
	q := ctx.QueryStrings("name")
	if len(q) > 0 {
		logrus.Info("Get user '" + q[0] + "' request")
	} else {
		logrus.Error("Wrong HTML query")
		r.JSON(500, map[string]interface{}{
			"Error": "Wrong HTML query",
		})
		return
	}
	u, err := models.GetUserByName(q[0])
	if u == nil {
		logrus.Info("User '" + q[0] + "' not found")
		r.JSON(404, map[string]interface{}{
			"Error": "User not found"})
		return
	}
	if err != nil {
		logrus.Error(err)
		r.JSON(500, map[string]interface{}{
			"Error": "Internal problem",
		})
	}
	b, _ := json.Marshal(u)
	r.Data(200, b)
}

func AddUser(r renders.Render, ctx *macaron.Context) {
	b, err := ctx.Req.Body().Bytes()
	if err != nil {
		logrus.Error(err)
		r.JSON(500, map[string]interface{}{
			"Error": err,
		})
		return
	}
	var u models.User
	err = json.Unmarshal(b, &u)
	if err != nil {
		logrus.Error(err)
		r.JSON(500, map[string]interface{}{
			"Error": err,
		})
		return
	}
	err = models.AddUserNew(u)
	if err != nil {
		r.JSON(500, map[string]interface{}{
			"Error": err,
		})
		logrus.Error(err)
	}
}
