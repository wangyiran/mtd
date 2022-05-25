package api

import (
	"myTodoList/service"

	"github.com/gin-gonic/gin"
)

//api接口，gin路由执行函数
//接口内初始化一个服务，解耦，具体注册逻辑在服务处实现，前端数据也由服务接收
func UserRegister(c *gin.Context) {
	var userService service.UserService
	if err := c.ShouldBind(&userService); err != nil {
		c.JSON(400, err)
	} else {
		//res是一个待序列化标准返回结构体
		res := userService.Register()
		c.JSON(200, res)
	}
}

func UserLogin(c *gin.Context) {
	var userService service.UserService
	if err := c.ShouldBind(&userService); err != nil {
		c.JSON(400, err)
	} else {
		res := userService.Login()
		c.JSON(200, res)
	}
}
