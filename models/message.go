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
func UpdateMessage(InfoId int64, status int, donetime time.Time) *gorm.DB {
	msg := Message{}
	utils.DB.Where("info_id = ?", InfoId).Find(&msg)
	if status == 3 {
		return utils.DB.Model(&msg).Updates(map[string]interface{}{"status": 3, "done_time": donetime})
	}
	return utils.DB.Model(&msg).Updates(map[string]interface{}{"status": status})
}

// 查询信息
func QueryMessage(InfoId int64) time.Time {
	msg := Message{}
	utils.DB.Where("info_id = ?", InfoId).Find(&msg)
	return msg.CreateTime
}

// 查询所有消息
func FindMsgByName(InfoId int64) []Message {
	var msg []Message
	utils.DB.Where("user_id = ?", InfoId).Find(&msg)
	return msg
}
