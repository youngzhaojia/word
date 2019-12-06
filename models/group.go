package models

import (
	"time"
)

type Group struct {
	Id         int       `gorm:"column:FuiId;primary_key;AUTO_INCREMENT"json:"id"`
	Name       string    `gorm:"column:FstrName"json:"name"`
	UserId     int       `gorm:"column:FuiUserId"json:"userId"`
	CreateTime int64     `gorm:"column:FuiCreateTime"json:"createTime"`
	UpdateTime time.Time `gorm:"column:FuiUpdateTime"json:"updateTime"`
}

func GetGroupList(pageNum int, pageSize int, params interface{}) ([]Group, error) {
	var (
		groups []Group
		err    error
	)

	if pageNum >= 0 && pageSize > 0 {
		err = db.Where(params).Limit(pageSize).Offset(pageNum).Find(&groups).Error
	} else {
		err = db.Where(params).Find(&groups).Error
	}

	if err != nil {
		return nil, err
	}
	return groups, err
}

func GetGroupTotal(params interface{}) (count int) {
	db.Model(&Group{}).Where(params).Count(&count)
	return
}

func AddGroup(name string, userId int) (int, error) {
	group := Group{
		Name:       name,
		UserId:     userId,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now(),
	}
	if err := db.Create(&group).Error; err != nil {
		return 0, err
	}
	return group.Id, nil
}

func EditGroup(groupId int, params interface{}) error {
	err := db.Model(&Group{}).Where("FuiId = ?", groupId).Update(params).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteGroup(groupId int) error {
	err := db.Where("FuiId = ?", groupId).Delete(&Group{}).Error
	if err != nil {
		return err
	}
	return nil
}
