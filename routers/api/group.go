package api

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
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
	name := c.DefaultPostForm("name", "未命名")
	if name == "" {
		name = "未命名"
	}

	userId := 8

	groupId, err := models.AddGroup(name, userId)
	// 失败
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ret": e.ERROR,
			"msg": err,
		})
		return
	}
	// 成功
	data := make(map[string]interface{})
	data["id"] = groupId

	c.JSON(http.StatusOK, gin.H{
		"ret":  e.SUCCESS,
		"msg":  "success",
		"data": data,
	})
}

func EditGroup(c *gin.Context) {
	groupId := com.StrTo(c.PostForm("id")).MustInt()
	newName := c.PostForm("name")

	data := make(map[string]interface{})
	data["Name"] = newName

	err := models.EditGroup(groupId, data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ret": e.ERROR,
			"msg": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"ret": e.SUCCESS,
		"msg": "success",
	})
}

func DeleteGroup(c *gin.Context) {
	groupId := com.StrTo(c.PostForm("id")).MustInt()
	err := models.DeleteGroup(groupId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ret": e.ERROR,
			"msg": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"ret": e.SUCCESS,
		"msg": "success",
	})
}
