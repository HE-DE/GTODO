package service

import (
	"GTODO/models"
	"GTODO/utils"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// AddMessageByUser
// @Summary 添加消息(未指派)
// @Tags 消息模块
// @param UserId query string false "UserId"
// @param Content query string false "Content"
// @Router /addmsgu [get]
func AddMessageByUser(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("UserId"))
	content := c.Query("Content")
	//使用雪花算法创建ID
	if err := utils.SnowflakeInit("2023-08-30", 1); err != nil {
		fmt.Println("Init() failed, err = ", err)
		return
	}
	id := utils.GenID()
	msg := models.Message{
		InfoID:     id,
		UserId:     int64(userId),
		Content:    content,
		CreateTime: time.Now(),
		Status:     0,
	}
	models.AddMessage(msg)
	models.UpdateUserInfo(int64(msg.UserId))
	//向user_msging表中添加元素
	umsg := models.User_Msging{
		UserID:    int64(msg.UserId),
		InfoID:    msg.InfoID,
		DoingTime: time.Since(msg.CreateTime),
		Status:    0,
	}
	models.AddInfo(umsg)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
	})
}

// AddMessageByAdmin
// @Summary 添加消息(指派)
// @Tags 消息模块
// @param UserId query string false "UserId"
// @param Content query string false "Content"
// @param AdminId query string false "AdminId"
// @Router /addmsga [get]
func AddMessageByAdmin(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("UserId"))
	content := c.Query("Content")
	adminId, _ := strconv.Atoi(c.Query("AdminId"))
	//使用雪花算法创建ID
	if err := utils.SnowflakeInit("2023-08-30", 1); err != nil {
		fmt.Println("Init() failed, err = ", err)
		return
	}
	id := utils.GenID()
	msg := models.Message{
		InfoID:     id,
		UserId:     int64(userId),
		AdminId:    int64(adminId),
		Content:    content,
		CreateTime: time.Now(),
		Status:     0,
	}
	models.AddMessage(msg)
	models.UpdateUserInfo(int64(msg.UserId))
	//向user_msging表中添加元素
	umsg := models.User_Msging{
		UserID:    int64(msg.UserId),
		InfoID:    msg.InfoID,
		DoingTime: time.Since(msg.CreateTime),
		Status:    0,
	}
	models.AddInfo(umsg)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
	})
}

// UpdateInfoStatus
// @Summary 更新消息的状态
// @Tags 消息模块
// @param InfoId query string false "InfoId"
// @param Status query string false "Status"
// @param UserId query string false "UserId"
// @Router /updatemsg [get]
func UpdateInfoStatus(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("InfoId"))
	status, _ := strconv.Atoi(c.Query("Status"))
	userId, _ := strconv.Atoi(c.Query("UserId"))
	if status == 3 {
		donetime := time.Now()
		models.UpdateMessage(int64(id), status, donetime)
		models.DeleteUserMsging(int64(id))
		CreateTime := models.QueryMessage(int64(id))
		umsged := models.User_Msged{
			UserID:     int64(userId),
			InfoID:     int64(id),
			CreateTime: CreateTime,
			DoneTime:   donetime,
			DoingTime:  donetime.Sub(CreateTime),
		}
		models.AddUMsged(umsged)
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "处理成功",
		})
	} else {
		models.UpdateMessage(int64(id), status, time.Now())
		models.UpdateMsging(int64(id), status)

		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "消息状态更新成功",
		})
	}
}
