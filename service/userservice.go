package service

import (
	"GTODO/models"
	"GTODO/utils"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Register
// @Summary 注册用户
// @Tags 用户模块
// @param name formData string false "name"
// @param password formData string false "password"
// @param phone formData string false "phone"
// @param sex formData string false "sex"
// @param email formData string false "email"
// @param indentity formData string false "indentity"
// @Router /register [post]
func Register(c *gin.Context) {
	Name := c.PostForm("name")
	password := c.PostForm("password")
	Password := utils.Md5Encode(password)
	Sex, err := strconv.Atoi(c.PostForm("sex"))
	if err != nil {
		c.JSON(200, gin.H{"status": "error", "message": "性别数据错误！"})
		return
	}
	Phone := c.PostForm("phone")
	Email := c.PostForm("email")
	Indentity, err := strconv.Atoi(c.PostForm("indentity"))
	if err != nil {
		c.JSON(200, gin.H{"status": "error", "message": "身份数据错误！"})
		return
	}
	data1 := models.FindUserByName(Name)
	//判断用户名是否存在
	if data1.Name != "" {
		c.JSON(200, gin.H{"status": "error", "message": "用户名已存在！"})
		return
	}
	// 创建用户

	//使用雪花算法创建ID
	if err := utils.SnowflakeInit("2023-08-30", 1); err != nil {
		fmt.Println("Init() failed, err = ", err)
		return
	}
	id := utils.GenID()
	user := models.User{
		Name:      Name,
		UserID:    id,
		Password:  Password,
		Sex:       Sex,
		Phone:     Phone,
		Email:     Email,
		Indentity: Indentity,
	}
	models.CreateUser(user)
	c.JSON(200, gin.H{"status": "success", "message": "创建用户成功！"})
}

// Login
// @Summary 登录
// @Tags 用户模块
// @param name formData string false "name"
// @param password formData string false "password"
// @Router /login [post]
func Login(c *gin.Context) {
	Name := c.PostForm("name")
	password := c.PostForm("password")
	fmt.Println("Name: ", Name)
	data1 := models.FindUserByName(Name)
	//判断用户名是否存在
	if data1.Password == "" {
		c.JSON(200, gin.H{"status": "error", "message": "用户名不存在", "data": ""})
		return
	}
	//判断解密后的密码是否相同
	flag := utils.VaildPassword(password, data1.Password)
	if !flag {
		c.JSON(200, gin.H{"status": "error", "message": "密码错误", "data": ""})
		return
	}
	c.JSON(200, gin.H{"status": "success", "message": "登录成功", "data": data1})
}

// GetUser
// @Summary 获取用户信息
// @Tags 用户模块
// @param name query string false "name"
// @Router /getuser [get]
func GetUser(c *gin.Context) {
	Name := c.Query("name")
	fmt.Println("Name: ", Name)
	data1 := models.FindUserByName(Name)
	if data1.Password == "" {
		c.JSON(200, gin.H{"status": "error", "message": "用户名不存在"})
		return
	}
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "获取用户成功",
		"data":    data1,
	})
}

// GetUserName
// @Summary 获取用户名的全部信息
// @Tags 用户模块
// @Router /getusername [get]
func GetUserName(c *gin.Context) {
	data1 := models.GetAllUserName()
	if len(data1) == 0 {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "暂无用户",
		})
		return
	}
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "获取用户名成功",
		"data":    data1,
	})
}

// DeleteUser
// @Summary 删除用户信息
// @Tags 用户模块
// @param name formData string false "name"
// @param password formData string false "password"
// @Router /deleteuser [post]
func DeleteUser(c *gin.Context) {
	Name := c.PostForm("name")
	data1 := models.FindUserByName(Name)
	password := c.PostForm("password")
	//判断用户名是否存在
	if data1.Password == "" {
		c.JSON(200, gin.H{"status": "error", "message": "用户名不存在"})
		return
	}
	//判断解密后的密码是否相同
	flag := utils.VaildPassword(password, data1.Password)
	if !flag {
		c.JSON(200, gin.H{"status": "error", "message": "密码错误"})
		return
	}
	models.DeleteUser(data1)
	c.JSON(200, gin.H{"status": "success", "message": "删除成功"})
}

// UpdateUser
// @Summary 更新用户
// @Tags 用户模块
// @param name formData string false "name"
// @param password formData string false "password"
// @param phone formData string false "phone"
// @param sex formData string false "sex"
// @param email formData string false "email"
// @Router /updateuser [post]
func UpdateUser(c *gin.Context) {
	Name := c.PostForm("name")
	data1 := models.FindUserByName(Name)
	if data1.Password == "" {
		c.JSON(200, gin.H{"status": "error", "message": "用户名不存在"})
		return
	}
	password := c.PostForm("password")
	Password := utils.Md5Encode(password)
	Sex, err := strconv.Atoi(c.PostForm("sex"))
	if err != nil {
		c.JSON(200, gin.H{"status": "error", "message": "性别数据错误！"})
		return
	}
	Phone := c.PostForm("phone")
	Email := c.PostForm("email")
	models.UpdateUser(data1, Password, Sex, Phone, Email)
	c.JSON(200, gin.H{"status": "success", "message": "修改成功"})
}

// GetUserList
// @Summary 获取用户的全部信息
// @Tags 用户模块
// @Router /getuserlist [get]
func GetUserList(c *gin.Context) {
	data1 := models.GetAllUser()
	if len(data1) == 0 {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "暂无用户",
		})
		return
	}
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "获取用户名成功",
		"data":    data1,
	})
}
