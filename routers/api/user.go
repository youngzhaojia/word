package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"word/models"
	"word/pkg/e"
	"word/pkg/util"
)

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	code := e.SUCCESS
	data := make(map[string]interface{})

	isExist := models.CheckAuth(username, password)
	if !isExist {
		code = e.ERROR_AUTH
	} else {
		token, err := util.GenerateToken(username, password)
		if err != nil {
			log.Println("ERROR_AUTH_TOKEN:", err)
			code = e.ERROR_AUTH_TOKEN
		} else {
			data["token"] = token
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"ret":  code,
		"msg":  e.GetMsgLabel(code),
		"data": data,
	})
}
