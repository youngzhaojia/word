package models

import (
	"time"
)

type Word struct {
	FuiId       int    `gorm:"column:FuiId;primary_key;AUTO_INCREMENT"json:"id"`
	Content     string `gorm:"column:FstrContent"json:"content"`
	Translation string `gorm:"column:FstrTranslation"json:"translation"`
	UserId      int    `gorm:"column:FuiUserId"json:"userId"`
	GroupId     int    `gorm:"column:FuiGroupId"json:"groupId"`
	CreateTime  int64  `gorm:"column:FuiCreateTime"json:"createTime"`
}

func GetWordListByGroupId(groupId int) ([]Word, error) {
	var (
		wordList []Word
		err      error
	)

	err = db.Where("FuiGroupId = ?", groupId).Find(&wordList).Error
	if err != nil {
		return nil, err
	}
	return wordList, nil
}

func GetWordCountByGroupId(groupId int) int {
	count := 0
	db.Model(&Word{}).Where("FuiGroupId = ?", groupId).Count(&count)
	return count
}

func GetWordDetail(wordId int) Word {
	var word Word
	db.Where("FuiId = ?", wordId).First(&word)
	return word
}

func AddWord(content string, translation string, userId int, groupId int) (int, error) {
	word := Word{
		Content:     content,
		Translation: translation,
		UserId:      userId,
		GroupId:     groupId,
		CreateTime:  time.Now().Unix(),
	}

	if err := db.Create(&word).Error; err != nil {
		return 0, err
	}
	return word.FuiId, nil
}

func DeleteWord(wordId int) error {
	return db.Where("FuiId = ?", wordId).Delete(&Word{}).Error
}

// 根据用户id获取每个分组的单词数量
type GroupWordCount struct {
	GroupId int
	Num     int
}

// 根据用户id获取每个分组的单词数量
func GetGroupWordCountByUserId(userId int) map[int]int {
	groupWordCount := make(map[int]int)

	var data []GroupWordCount
	db.Table("t_word").Select(
		"FuiGroupId as group_id, count(FuiId) as num").Where(
		"FuiUserId = ?", userId).Group("FuiGroupId").Scan(&data)
	for _, item := range data {
		groupWordCount[item.GroupId] = item.Num
	}

	return groupWordCount
}

// 根据userId获取用户所有的单词
func GetAllWordListByUserId(userId int) ([]Word, error){
	var (
		wordList []Word
		err      error
	)

	err = db.Where("FuiUserId = ?", userId).Find(&wordList).Error
	if err != nil {
		return nil, err
	}
	return wordList, nil
}
