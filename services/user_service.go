package service

import (
	"github.com/adminwjp/users-go/models"
	"github.com/adminwjp/users-go/dtos"
)
//用户服务接口
type UserService interface {
	Clean()
	BaseUserService
	/**
	根据手机号、邮箱、用户名登录
	*/
	Login(user *dtos.UserInput)(*models.UserModel,error)

	/**
	根据手机号登录
	*/
	LoginByPhone(user *dtos.UserPhoneInput)(*models.UserModel,error)

	/**
	根据邮箱登录
	*/
	LoginByEmail(user *dtos.UserEmailInput)(*models.UserModel,error)

	/**
	根据用户名登录
	*/
	LoginByUserName(user *dtos.UserUserNameInput)(*models.UserModel,error)



	/*
		修改身份认证基本信息
	*/
	UpdateAuthBasic(input *dtos.UpdateUserAuthBasicInput)(int,error)

	/*
		根据条件查询用户信息
	*/
	List(user *dtos.GetUserInput) ([]models.UserModel,int64,error)

	/*
		根据id查询用户信息
	*/
	Get(id int64) (*models.UserModel,error)

	/*
	根据条件查询用户日志信息
	*/
	ListLogs(user *dtos.GetUserLogInput) ([]models.UserLogModel,int64,error)
}

