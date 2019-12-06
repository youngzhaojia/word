package routers

import (
	"github.com/gin-gonic/gin"
	"word/pkg/setting"
	"word/routers/api"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	apiGroup := r.Group("/api")
	{
		apiGroup.GET("/group/list", api.GetGroupList)
		apiGroup.POST("/group/add", api.AddGroup)
		apiGroup.POST("/group/update", api.EditGroup)
		apiGroup.POST("/group/delete", api.DeleteGroup)
	}
	return r
}
