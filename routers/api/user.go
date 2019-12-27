package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"word/models"
	"word/pkg/app"
	"word/pkg/e"
	"word/pkg/logging"
	"word/pkg/util"
)

// 用户登录
func Login(c *gin.Context) {
	appG := app.Gin{C: c}

	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == "" || password == "" {
		logging.Info("参数错误")
		appG.ResponseErr(e.ERROR, "参数错误")
		return
	}

	data := make(map[string]interface{})

	user := models.GetUser(username, password)
	if user.Id <= 0 {
		logging.Info("参数错误, username:" + username + "|password:" + password)
		appG.ResponseErrMsg("用户名或密码错误")
		return
	} else {
		token, err := util.GenerateToken(user.Id, user.Username)
		// token生成失败
		if err != nil {
			logging.Error("token生成失败:" + fmt.Sprintf("%s", err))
			appG.ResponseErrMsg(e.GetMsgLabel(e.ERROR_AUTH_TOKEN))
			return
		} else {
			data["token"] = token
		}
	}
	appG.ResponseSuccess("ok", data)
}
