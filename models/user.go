package models

type User struct {
	Id       int    `gorm:"column:FuiId;primary_key;AUTO_INCREMENT"json:"id"`
	Username string `gorm:"column:FstrUsername"json:"username"`
	Password string `gorm:"column:FstrPassword"json:"password"`
}

func CheckAuth(username string, password string) bool {
	var user User

	db.Select("FuiId").Where(User{Username: username, Password: password}).First(&user)
	if user.Id > 0 {
		return true
	}
	return false
}

func GetUser(username string, password string) User {
	user := User{}

	db.Where(User{Username: username, Password: password}).First(&user)
	return user
}
