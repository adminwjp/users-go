package daos

import (
	"github.com/adminwjp/users-go/dtos"
	"github.com/adminwjp/users-go/models"
)

//管理员接口 mong 实现
type UserDao interface {
	BaseUserDao

	/*
		根据手机号、邮箱、用户名登录
	*/
	Login(user *dtos.UserInput)(*models.UserModel,error)

	/*
		根据手机号登录
	*/
	LoginByPhone(user *dtos.UserPhoneInput)(*models.UserModel,error)
	/*
		根据邮箱登录
	*/
	LoginByEmail(user *dtos.UserEmailInput)(*models.UserModel,error)

	/*
		根据用户名登录
	*/
	LoginByUserName(user *dtos.UserUserNameInput)(*models.UserModel,error)

	Get(id int64)(*models.UserModel,error)
	/*
		根据条件查询用户信息
	*/
	List(input *dtos.GetUserInput)([]models.UserModel,int64,error)

	/*
		修改身份认证基本信息
	*/
	UpdateAuthBasic(input *dtos.UpdateUserAuthBasicInput)(int,error)
}
