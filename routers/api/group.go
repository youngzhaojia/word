package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"word/models"
	"word/pkg/app"
	"word/pkg/setting"
	"word/pkg/util"
)

// 分组列表
func GetGroupList(c *gin.Context) {
	appG := app.Gin{C: c}
	params := make(map[string]interface{})
	data := make(map[string]interface{})

	// 参数
	pageNum := util.GetPage(c)
	userId := c.GetInt("userId")
	params["FuiUserId"] = userId

	// 分组列表数据
	groupList, err := models.GetGroupList(pageNum, setting.AppSetting.PageSize, params)
	if err != nil {
		appG.ResponseErrMsg("数据查询出错")
		return
	}
	// 分组单词数量
	groupWordCount := models.GetGroupWordCountByUserId(userId)

	// 拼装数据
	list := make([]interface{}, len(groupList))
	for key, groupItem := range groupList {
		listItem := make(map[string]interface{})

		itemJson, _ := json.Marshal(groupItem)
		json.Unmarshal(itemJson, &listItem)
		listItem["wordNum"] = groupWordCount[groupItem.Id]
		list[key] = listItem
	}

	data["list"] = list
	data["total"] = models.GetGroupTotal(params)

	appG.ResponseSuccess("ok", data)
}

// 分组新增
func AddGroup(c *gin.Context) {
	appG := app.Gin{C: c}
	name := c.PostForm("name")
	if name == "" {
		name = "未命名"
	}

	userId := c.GetInt("userId")

	groupId, err := models.AddGroup(name, userId)
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

// 分组编辑
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

// 分组删除
func DeleteGroup(c *gin.Context) {
	appG := app.Gin{C: c}

	groupId := com.StrTo(c.PostForm("id")).MustInt()

	userId := c.GetInt("userId")
	group := models.GetGroupDetail(groupId)

	if group.UserId != userId {
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
