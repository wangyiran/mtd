package service

import (
	"myTodoList/model"
	"myTodoList/myutils"
	"myTodoList/serializer"

	"github.com/jinzhu/gorm"
)

//服务处接收前端数据
type UserService struct {
	UserName     string `form:"user_name" json:"user_name"`
	UserPassword string `form:"password" json:"password"`
}

//注册逻辑
func (u *UserService) Register() serializer.Response {
	//声明一个record，用于mysql插入数据
	var user model.User
	var count int64
	model.DB.Where("user_name=?", u.UserName).First(&user).Count(&count)
	if count != 0 {
		//已经存在
		return serializer.Response{
			Status: "400",
			Msg:    "用户已经存在，请直接登录",
		}
	}
	//不存在，生成用户
	user.UserName = u.UserName
	user.PasswordDigest = user.CreateDigestPassword(u.UserPassword)

	if err := model.DB.Create(&user).Error; err != nil {
		panic("注册用户入库失败")
	}

	return serializer.Response{
		Status: "200",
		Msg:    "注册用户成功！",
	}
}

func (u *UserService) Login() serializer.Response {
	var user model.User
	err := model.DB.Where("user_name=?", u.UserName).First(&user).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return serializer.Response{
				Status: "400",
				Msg:    "用户不存在 请注册",
			}
		} else {
			return serializer.Response{
				Status: "400",
				Msg:    "登录数据库错误" + err.Error(),
			}
		}
	}
	//用户存在,验证密码
	if user.CheckPassword(u.UserPassword) == false {
		return serializer.Response{
			Status: "400",
			Msg:    "密码错误",
		}
	}
	//登录成功 分发token
	token, err := myutils.GenerateToken(user.ID, user.UserName, u.UserPassword)
	if err != nil {
		return serializer.Response{
			Status: "400",
			Msg:    "Token分发错误",
		}
	}
	return serializer.Response{
		Status: "200",
		Data:   myutils.BuildTokenData(user.UserName, u.UserPassword, token),
	}
}
