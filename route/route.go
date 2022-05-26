package route

import (
	"myTodoList/api"
	"myTodoList/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	//创建一个session中间件
	store := cookie.NewStore([]byte("bk"))
	r.Use(sessions.Sessions("mysession", store))
	v1 := r.Group("api/v1")
	{
		v1.POST("/user/register", api.UserRegister)
		v1.POST("/user/login", api.UserLogin)
		v2 := v1.Group("/")
		v2.Use(middleware.JWT())
		{
			v2.POST("/task", api.CreateTask)
			v2.GET("/list", api.ShowList)
			v2.POST("/update", api.UpdateTask)
			v2.POST("/delete", api.DeleteTask)
		}
	}

	return r
}
