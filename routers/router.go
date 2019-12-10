package routers

import (
	"github.com/gin-gonic/gin"
	"word/middleware/jwt"
	"word/pkg/setting"
	"word/routers/api"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	// 登录
	r.POST("/user/login", api.Login)

	apiGroup := r.Group("/api")
	apiGroup.Use(jwt.JWT())
	// 分组
	{
		apiGroup.GET("/group/list", api.GetGroupList)
		apiGroup.POST("/group/add", api.AddGroup)
		apiGroup.POST("/group/update", api.EditGroup)
		apiGroup.POST("/group/delete", api.DeleteGroup)
	}
	// 单词
	{
		apiGroup.GET("/word/list", api.GetWordList)
		apiGroup.POST("/word/add", api.AddWord)
		apiGroup.POST("/word/delete", api.DeleteWord)
	}
	return r
}
