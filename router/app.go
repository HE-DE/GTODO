package router

import (
	"GTODO/docs"
	"GTODO/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})
	r.POST("/login", service.Login)
	r.POST("/register", service.Register)
	r.GET("/getuser", service.GetUser)
	r.POST("/deleteuser", service.DeleteUser)
	r.POST("/updateuser", service.UpdateUser)
	r.GET("/addmsgu", service.AddMessageByUser)
	r.GET("/addmsga", service.AddMessageByAdmin)
	r.GET("/updatemsg", service.UpdateInfoStatus)
	r.GET("/getmsg", service.GetMessage)
	r.GET("/getdoing", service.Getdoing)
	r.GET("/getusername", service.GetUserName)
	r.GET("/getuserlist", service.GetUserList)
	r.GET("/updatedoing", service.UpdateDoing)
	return r
}
