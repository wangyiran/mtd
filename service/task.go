package service

import (
	"fmt"
	"myTodoList/model"
	"myTodoList/myutils"
	"myTodoList/serializer"
	"time"
)

type CommonTaskService struct {
	UserID   string
	UserName string `json:"user_name" form:"user_name"`
	Token    string
}

type CreateTaskService struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Status  string `json:"status" form:"status"`
	CommonTaskService
}

type ShowListService struct {
	CommonTaskService
}

type UpdateService struct {
	CommonTaskService
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Status  string `json:"status" form:"status"`
}

type DeleteService struct {
	CommonTaskService
	Title string `json:"title" form:"title"`
}

func (s *CreateTaskService) CreateTask() serializer.Response {
	claim, err := myutils.ParseToken(s.CommonTaskService.Token)
	if err != nil {
		return serializer.Response{
			Status: "400",
			Msg:    "创建todolist时令牌错误",
		}
	}
	username := claim.UserName
	var user model.User
	model.DB.Where("user_name=?", username).First(&user)
	var task model.Task
	var count int64
	model.DB.Table("users").Joins("left join tasks on users.id = tasks.user_id").Where("title=? and user_id=?", s.Title, user.ID).First(&user).Count(&count)
	fmt.Println(count)
	if count > 0 {
		return serializer.Response{
			Status: "400",
			Msg:    "已经存在",
		}
	}

	task = model.Task{
		User:      user,
		UserId:    user.ID,
		Title:     s.Title,
		Status:    s.Status,
		Content:   s.Content,
		StartTime: time.Now().Unix(),
	}
	//插入
	if err := model.DB.Create(&task).Error; err != nil {
		return serializer.Response{
			Status: "400",
			Msg:    "新建备忘录错误" + err.Error(),
		}
	} else {
		return serializer.Response{
			Status: "200",
			Msg:    "创建备忘录\"" + s.Title + "\"成功",
		}
	}
}

func (s *ShowListService) ShowList() serializer.TaskList {
	claim, _ := myutils.ParseToken(s.CommonTaskService.Token)
	userName := claim.UserName
	var user model.User
	model.DB.Where("user_name=?", userName).First(&user)
	var list []serializer.Task
	var count int64
	model.DB.Table("users").Joins("left join tasks on users.id=tasks.user_id").Select("user_name,title,status,start_time,end_time,content").Where("users.id=?", user.ID).Scan(&list).Count(&count)
	//model.DB.Raw("select user_name,title,status,start_time,end_time,content from users,tasks where users.id=tasks.user_id and users.id=1").Scan(&list).Count(&count)
	//model.DB.Table("users").Joins("left join tasks on users.id=tasks.user_id").Where("users.id=?", user.ID).Select("user_name,title,status,start_time,end_time,content").Scan(list).Count(&count)
	return serializer.TaskList{
		Status: "200",
		Count:  count,
		List:   list,
	}
}

func (s *UpdateService) Update() serializer.Response {
	claim, err := myutils.ParseToken(s.CommonTaskService.Token)
	if err != nil {
		return serializer.Response{
			Status: "400",
			Msg:    "创建todolist时令牌错误",
		}
	}
	username := claim.UserName
	var user model.User
	model.DB.Where("user_name=?", username).First(&user)
	var task model.Task
	var count int64
	model.DB.Table("users").Joins("left join tasks on users.id = tasks.user_id").Where("title=? and user_id=?", s.Title, user.ID).First(&user).Count(&count)
	fmt.Println(count)
	if count == 0 {
		return serializer.Response{
			Status: "400",
			Msg:    "不存在",
		}
	}

	model.DB.Where("title=?", s.Title).First(&task)

	task = model.Task{
		User:      user,
		UserId:    user.ID,
		Title:     s.Title,
		Status:    s.Status,
		Content:   s.Content,
		StartTime: time.Now().Unix(),
	}
	//插入
	if err := model.DB.Where("title=? and user_id=?", s.Title, user.ID).Save(&task).Error; err != nil {
		return serializer.Response{
			Status: "400",
			Msg:    "更新备忘录错误" + err.Error(),
		}
	} else {
		return serializer.Response{
			Status: "200",
			Msg:    "更新备忘录\"" + s.Title + "\"成功",
		}
	}
}

func (s *DeleteService) DeleteTask() serializer.Response {
	claim, err := myutils.ParseToken(s.CommonTaskService.Token)
	if err != nil {
		return serializer.Response{
			Status: "400",
			Msg:    "创建todolist时令牌错误",
		}
	}
	username := claim.UserName
	var user model.User
	model.DB.Where("user_name=?", username).First(&user)
	var task model.Task
	var count int64
	model.DB.Table("users").Joins("left join tasks on users.id = tasks.user_id").Where("title=? and user_id=?", s.Title, user.ID).First(&user).Count(&count)
	fmt.Println(count)
	if count == 0 {
		return serializer.Response{
			Status: "400",
			Msg:    "不存在,无法删除",
		}
	}
	if err := model.DB.Where("title=? and user_id=?", s.Title, user.ID).Delete(&task).Error; err != nil {
		return serializer.Response{
			Status: "400",
			Msg:    "删除备忘录错误" + err.Error(),
		}
	} else {
		return serializer.Response{
			Status: "200",
			Msg:    "删除备忘录\"" + s.Title + "\"成功",
		}
	}
}
