package api

import (
	"myTodoList/service"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	//创建一个task服务。
	var taskCreateService service.CreateTaskService
	token := c.GetHeader("Authorization")
	//绑定前端数据
	if err := c.ShouldBind(&taskCreateService); err != nil {
		c.JSON(400, err)
	} else {
		taskCreateService.CommonTaskService.Token = token
		res := taskCreateService.CreateTask()
		c.JSON(200, res)
	}
}

func ShowList(c *gin.Context) {
	var taskShowListService service.ShowListService
	token := c.GetHeader("Authorization")
	taskShowListService.CommonTaskService.Token = token
	res := taskShowListService.ShowList()
	c.JSON(200, res)
}

func UpdateTask(c *gin.Context) {
	var updateTaskService service.UpdateService
	token := c.GetHeader("Authorization")
	updateTaskService.CommonTaskService.Token = token
	if err := c.ShouldBind(&updateTaskService); err != nil {
		c.JSON(400, err)
	} else {
		res := updateTaskService.Update()
		c.JSON(200, res)
	}
}

func DeleteTask(c *gin.Context) {
	var deleteTaskService service.DeleteService
	token := c.GetHeader("Authorization")
	deleteTaskService.CommonTaskService.Token = token
	if err := c.ShouldBind(&deleteTaskService); err != nil {
		c.JSON(400, err)
	} else {
		res := deleteTaskService.DeleteTask()
		c.JSON(200, res)
	}

}
