package models

import (
	"GTODO/utils"
	"time"

	"gorm.io/gorm"
)

type User_Msged struct {
	gorm.Model
	UserID     int64         //处理者ID
	InfoID     int64         //消息ID
	CreateTime time.Time     //创建时间
	DoneTime   time.Time     //完结时间
	DoingTime  time.Duration //已用时间
}

func (umed *User_Msged) Tablename() string {
	return "user_msgeds"
}

func AddUMsged(umsged User_Msged) *gorm.DB {
	return utils.DB.Create(&umsged)
}
