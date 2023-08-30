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
		UserId:     userId,
		Content:    content,
		CreateTime: time.Now(),
		Status:     0,
	}
	models.AddMessage(msg)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
	})
}
