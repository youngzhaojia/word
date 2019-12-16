package api

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"word/models"
	"word/pkg/app"
)

// 单词列表
func GetWordList(c *gin.Context) {
	appG := app.Gin{C: c}
	groupId, _ := com.StrTo(c.Query("groupId")).Int()

	group := models.GetGroupDetail(groupId)
	userId := c.GetInt("userId")
	// 验证是否是该分组用户
	if group.UserId != userId {
		appG.ResponseErrMsg("参数错误")
		return
	}

	data := make(map[string]interface{})
	groupList, err := models.GetWordListByGroupId(groupId)
	if err != nil {
		appG.ResponseErrMsg("数据出错")
		return
	}
	data["list"] = groupList
	data["count"] = models.GetWordCountByGroupId(groupId)

	appG.ResponseSuccess("ok", data)
}

// 单词新增
func AddWord(c *gin.Context) {
	appG := app.Gin{C: c}

	content := c.DefaultPostForm("content", "")
	translation := c.DefaultPostForm("translation", "")
	groupId, _ := com.StrTo(c.PostForm("groupId")).Int()
	if content == "" || translation == "" || groupId == 0 {
		appG.ResponseErrMsg("参数不能为空")
		return
	}

	group := models.GetGroupDetail(groupId)
	userId := c.GetInt("userId")
	// 验证是否是该分组用户
	if group.UserId != userId {
		appG.ResponseErrMsg("参数错误")
		return
	}

	wordId, err := models.AddWord(content, translation, userId, groupId)
	if err != nil {
		appG.ResponseErrMsg("新增单词出错")
		return
	}
	data := make(map[string]interface{})
	data["id"] = wordId
	appG.ResponseSuccess("ok", data)
}

// 单词删除
func DeleteWord(c *gin.Context) {
	appG := app.Gin{C: c}

	wordId, _ := com.StrTo(c.PostForm("wordId")).Int()
	word := models.GetWordDetail(wordId)
	userId := c.GetInt("userId")

	// 验证是否是该单词创建用户
	if word.UserId != userId {
		appG.ResponseErrMsg("没权限删除")
		return
	}

	err := models.DeleteWord(wordId)
	if err != nil {
		appG.ResponseErrMsg("删除失败")
		return
	}
	appG.ResponseSuccess("删除成功", nil)
}
