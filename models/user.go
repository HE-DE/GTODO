package models

import (
	"GTODO/utils"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID    int64  //用户ID
	Name      string //姓名
	Password  string //密码
	Sex       int    //0-man  1-woman
	Phone     string //电话
	Email     string //邮箱
	Indentity int    //身份 0-admin 1-user
	Msging    int    //在办消息数
	Msged     int    //办结消息数
}

func (user *User) Tablename() string {
	return "users"
}

// 按用户名查找
func FindUserByName(name string) User {
	user := User{}
	utils.DB.Where("name = ?", name).Find(&user)
	return user
}

// 创建用户
func CreateUser(user User) *gorm.DB {
	return utils.DB.Create(&user)
}

func DeleteUser(user User) *gorm.DB {
	return utils.DB.Delete(&user)
}

// 更新用户信息
func UpdateUser(user User, password string, sex int, phone string, email string) *gorm.DB {
	return utils.DB.Model(&user).Updates(map[string]interface{}{"password": password, "sex": sex, "phone": phone, "email": email})
}

// 更新用户信息（信息添加）
func UpdateUserInfo(userId int64) *gorm.DB {
	user := User{}
	utils.DB.Where("user_id = ?", userId).Find(&user)
	return utils.DB.Model(&user).Updates(map[string]interface{}{"msging": user.Msging + 1})
}

// 更新用户信息（信息办结）
func UpdateUserInfoDone(userId int64) *gorm.DB {
	user := User{}
	utils.DB.Where("user_id = ?", userId).Find(&user)
	utils.DB.Model(&user).Updates(map[string]interface{}{"msging": user.Msging - 1})
	return utils.DB.Model(&user).Updates(map[string]interface{}{"msged": user.Msged + 1})
}

// 获取所有用户的用户名
func GetAllUserName() []string {
	var users []User
	utils.DB.Find(&users)
	var userNames []string
	for _, user := range users {
		userNames = append(userNames, user.Name)
	}
	return userNames
}

// 获取所有用户信息
func GetAllUser() []User {
	var users []User
	utils.DB.Find(&users)
	return users
}
