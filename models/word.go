package models

import "time"

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
	db.Where("FuiGroupId = ?", groupId).Count(&count)
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
