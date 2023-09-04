package models

import (
	"GTODO/utils"
	"time"

	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	InfoID     int64     //消息ID
	AdminId    int64     //指派者的ID
	UserId     int64     //处理者ID
	CreateTime time.Time //创建时间
	DoneTime   time.Time //完结时间
	Status     int       //状态 0-Doing;1-Buzy;2-finishing;3-Done
	Content    string    //消息内容
}

func (message *Message) Tablename() string {
	return "messages"
}

// 添加信息
func AddMessage(Msg Message) *gorm.DB {
	return utils.DB.Create(&Msg)
}

// 删除信息
func DeleteMessage(Msg Message) *gorm.DB {
	return utils.DB.Delete(&Msg)
}

// 更新信息
func UpdateMessage(Msg Message) *gorm.DB {
	return utils.DB.Model(&Msg).Updates(Msg)
}

// 查询信息
func QueryMessage(userId int64, AdminId int64, createTime time.Time) Message {
	var message Message
	utils.DB.Where("user_id = ? and admin_id = ? and create_time = ?", userId, AdminId, createTime).First(&message)
	return message
}
