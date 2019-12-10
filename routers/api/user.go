package api

import (
	"github.com/gin-gonic/gin"
	"word/models"
	"word/pkg/app"
	"word/pkg/e"
	"word/pkg/util"
)

// 用户登录
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
		appG.ResponseErrMsg("用户名或密码错误")
		return
	} else {
		token, err := util.GenerateToken(user.Id, user.Username)
		// token生成失败
		if err != nil {
			appG.ResponseErrMsg(e.GetMsgLabel(e.ERROR_AUTH_TOKEN))
			return
		} else {
			data["token"] = token
		}
	}
	appG.ResponseSuccess("ok", data)
}
