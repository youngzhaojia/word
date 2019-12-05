package models

type Group struct {
	Model

	Id     int    `gorm:"column:FuiId"json:"id"`
	Name   string `gorm:"column:FstrName"json:"name"`
	UserId int    `gorm:"column:FuiUserId"json:"userId"`
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
