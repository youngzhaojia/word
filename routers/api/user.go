package api

import (
	"github.com/gin-gonic/gin"
	"word/models"
	"word/pkg/app"
	"word/pkg/e"
	"word/pkg/util"
)

func Login(c *gin.Context) {
	appG := app.Gin{C: c}

	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == "" || password == "" {
		appG.ResponseErr(e.ERROR, "参数错误")
		return
	}

	data := make(map[string]interface{})

	user := models.GetUser(username, password)
	if user.Id <= 0 {
		appG.ResponseErrMsg(e.GetMsgLabel(e.ERROR_AUTH))
		return
	} else {
		token, err := util.GenerateToken(user.Id, user.Username)
		if err != nil {
			appG.ResponseErrMsg(e.GetMsgLabel(e.ERROR_AUTH_TOKEN))
			return
		} else {
			data["token"] = token
		}
	}
	appG.ResponseSuccess("ok", data)
}
