package service

import (
	"github.com/adminwjp/users-go/models"
	"github.com/adminwjp/users-go/dtos"
)
//管理员服务接口
type AdminService interface {
	BaseUserService
	/**
	根据手机号、邮箱、用户名登录
	*/
	Login(user *dtos.UserInput)(*models.AdminModel,error)

	/**
	根据手机号登录
	*/
	LoginByPhone(user *dtos.UserPhoneInput)(*models.AdminModel,error)

	/**
	根据邮箱登录
	*/
	LoginByEmail(user *dtos.UserEmailInput)(*models.AdminModel,error)

	/**
	根据用户名登录
	*/
	LoginByUserName(user *dtos.UserUserNameInput)(*models.AdminModel,error)



	/*
		根据旧密码修改密码
	*/
	UpdatePwdByOldPwd(input *dtos.UpdateUserPwdByPwdInput)(int,error)

	/*
		根据条件查询用户信息
	*/
	List(user *dtos.GetAdminInput) ([]models.AdminModel,int64,error)

	/*
		根据id查询用户信息
	*/
	Get(id int64) (*models.AdminModel,error)

	/*
		根据条件查询用户日志信息
	*/
	ListLogs(user *dtos.GetUserLogInput) ([]models.AdminLogModel,int64,error)
}
