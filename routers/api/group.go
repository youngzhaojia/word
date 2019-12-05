package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"word/models"
	"word/pkg/e"
	"word/pkg/setting"
	"word/pkg/util"
)

func GetGroupList(c *gin.Context) {
	params := make(map[string]interface{})
	data := make(map[string]interface{})

	pageNum := util.GetPage(c)

	// 参数
	params["FuiUserId"] = 8

	list, err := models.GetGroupList(pageNum, setting.PageSize, params)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ret": e.ERROR,
			"msg": "数据查询出错",
		})
	}

	data["list"] = list
	data["total"] = models.GetGroupTotal(params)

	c.JSON(http.StatusOK, gin.H{
		"ret":  e.SUCCESS,
		"msg":  "success",
		"data": data,
	})
}

func AddGroup(c *gin.Context) {

}

func EditGroup(c *gin.Context) {

}

func DeleteGroup(c *gin.Context) {

}
