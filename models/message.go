package models

import (
	"GTODO/utils"
	"time"

	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	InfoID     int64         //消息ID
	AdminId    int64         //指派者的ID
	UserId     int64         //处理者ID
	CreateTime time.Time     //创建时间
	DoneTime   time.Time     //完结时间
	DoingTime  time.Duration //处理时间
	Status     int           //状态 0-Doing;1-Buzy;2-finishing;3-Done
	Content    string        //消息内容
}

func (message *Message) Tablename() string {
	return "messages"
}

// 添加信息
func AddMessage(Msg Message) *gorm.DB {
	return utils.DB.Create(&Msg)
}

// 更新信息
func UpdateMessage(InfoId int64, status int, donetime time.Time) *gorm.DB {
	msg := Message{}
	utils.DB.Where("info_id = ?", InfoId).Find(&msg)
	if status == 3 {
		return utils.DB.Model(&msg).Updates(map[string]interface{}{"status": 3, "done_time": donetime, "doing_time": time.Since(msg.CreateTime)})
	}
	return utils.DB.Model(&msg).Updates(map[string]interface{}{"status": status, "doing_time": time.Since(msg.CreateTime)})
}

// 查询当前用户所有消息
func FindMsgByUser(UserId int64) []Message {
	var msg []Message
	utils.DB.Where("user_id = ?", UserId).Find(&msg)
	return msg
}

// 查询一条消息
func FindMsgById(InfoId int64) Message {
	var msg Message
	utils.DB.Where("info_id = ?", InfoId).Find(&msg)
	return msg
}
