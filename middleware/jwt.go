package middleware

import (
	"myTodoList/myutils"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if len(token) == 0 {
			//无授权
			//fmt.Println("no1")
			c.JSON(404, gin.H{
				"msg": "未授权",
			})
			c.Abort()
			return
		}
		_, err := myutils.ParseToken(token)
		if err != nil {
			//fmt.Println("no2")
			c.JSON(404, gin.H{
				"msg": err.Error(),
			})
			c.Abort()
			return
		}
		//fmt.Println("iok")
		c.Next()
	}
}
