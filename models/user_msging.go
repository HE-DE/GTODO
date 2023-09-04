package models

import (
	"GTODO/utils"
	"time"

	"gorm.io/gorm"
)

type User_Msging struct {
	gorm.Model
	UserID    int64         //处理者ID
	InfoID    int64         //消息ID
	DoingTime time.Duration //已用时间
	Status    int           //状态 0-Doing;1-Buzy;2-finishing
}

func (uming *User_Msging) Tablename() string {
	return "user_msgings"
}

func AddInfo(umsg User_Msging) *gorm.DB {
	return utils.DB.Create(&umsg)
}

// 在办事项的状态的更新
func UpdateMsging(InfoId int64, status int) *gorm.DB {
	umsging := User_Msging{}
	utils.DB.Where("info_id = ?", InfoId).Find(&umsging)
	return utils.DB.Model(&umsging).Updates(map[string]interface{}{"status": status})
}

// 事项办结后删除表中的记录
func DeleteUserMsging(InfoId int64) *gorm.DB {
	umsging := User_Msging{}
	utils.DB.Where("info_id = ?", InfoId).Find(&umsging)
	return utils.DB.Delete(&umsging)
}
