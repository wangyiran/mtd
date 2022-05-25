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
		author := v1.Group("/")
		author.Use(middleware.JWT())
		{
			author.GET("test", func(ctx *gin.Context) {
				ctx.JSON(200, gin.H{
					"msg": "ok",
				})
			})
		}
	}

	return r
}
