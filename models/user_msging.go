package models

import (
	"GTODO/utils"
	"time"

	"gorm.io/gorm"
)

type User_Msging struct {
	gorm.Model
	UserID    int64     //处理者ID
	InfoID    int64     //消息ID
	DoingTime time.Time //已用时间
	Status    int       //状态 0-Doing;1-Buzy;2-finishing
}

func (uming *User_Msging) Tablename() string {
	return "user_msgings"
}

func AdddInfo(umsg User_Msging) *gorm.DB {
	return utils.DB.Create(&umsg)
}
