package api

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"word/models"
	"word/pkg/app"
	"word/pkg/setting"
	"word/pkg/util"
)

func GetGroupList(c *gin.Context) {
	appG := app.Gin{C: c}
	params := make(map[string]interface{})
	data := make(map[string]interface{})

	pageNum := util.GetPage(c)

	// 参数
	params["FuiUserId"], _ = c.Get("userId")

	list, err := models.GetGroupList(pageNum, setting.PageSize, params)
	if err != nil {
		appG.ResponseErrMsg("数据查询出错")
		return
	}

	data["list"] = list
	data["total"] = models.GetGroupTotal(params)

	appG.ResponseSuccess("ok", data)
}

func AddGroup(c *gin.Context) {
	appG := app.Gin{C: c}
	name := c.PostForm("name")
	if name == "" {
		name = "未命名"
	}

	userId, _ := c.Get("userId")

	groupId, err := models.AddGroup(name, userId.(int))
	// 失败
	if err != nil {
		appG.ResponseErrMsg("新增失败")
		return
	}
	// 成功
	data := make(map[string]interface{})
	data["id"] = groupId

	appG.ResponseSuccess("ok", data)
}

func EditGroup(c *gin.Context) {
	appG := app.Gin{C: c}

	groupId := com.StrTo(c.PostForm("id")).MustInt()
	newName := c.PostForm("name")

	data := make(map[string]interface{})
	data["Name"] = newName

	err := models.EditGroup(groupId, data)
	if err != nil {
		appG.ResponseErrMsg("修改失败")
		return
	}
	appG.ResponseSuccess("ok", data)
}

func DeleteGroup(c *gin.Context) {
	appG := app.Gin{C: c}

	groupId := com.StrTo(c.PostForm("id")).MustInt()

	userId, _ := c.Get("userId")
	group := models.GetGroupDetail(groupId)

	if group.UserId != userId.(int) {
		appG.ResponseErrMsg("不是你的不能删除")
		return
	}

	err := models.DeleteGroup(groupId)
	if err != nil {
		appG.ResponseErrMsg("删除失败")
		return
	}
	appG.ResponseSuccess("ok", nil)
}
