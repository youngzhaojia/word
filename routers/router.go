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
		apiGroup.GET("/groups", api.GetGroupList)
		apiGroup.POST("/group", api.AddGroup)
		apiGroup.PUT("/group/:id", api.EditGroup)
		apiGroup.DELETE("/group/:id", api.DeleteGroup)
	}



	return r
}
