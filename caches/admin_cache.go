package caches

import (
	"github.com/adminwjp/users-go/dtos"
	"github.com/adminwjp/users-go/models"
)

//管理员缓存接口
type AdminCache interface {
	BaseUserCache
	/*
		根据手机号、邮箱、用户名登录
	*/
	Login(user *dtos.UserInput)(*models.AdminModel,error)

	/*
		根据手机号登录
	*/
	LoginByPhone(user *dtos.UserPhoneInput)(*models.AdminModel,error)

	/*
		根据邮箱登录
	*/
	LoginByEmail(user *dtos.UserEmailInput)(*models.AdminModel,error)

	/*
		根据用户名登录
	*/
	LoginByUserName(user *dtos.UserUserNameInput)(*models.AdminModel,error)


	/*
	根据用户id获取用户信息
	*/
	 GetByUserId(userId string) (*models.AdminModel,error)

	/*
	根据手机号、邮箱、用户名获取用户信息
	*/
	 GetByAccount(account string) (*models.AdminModel,error)

	/*
	根据手机号获取用户信息
	*/
	GetByPhone(phone string) (*models.AdminModel,error)

	/*
	根据邮箱获取用户信息
	*/
	GetByEmail(email string) (*models.AdminModel,error)

	/*
	根据用户名获取用户信息
	*/
	GetByUserName(userName string) (*models.AdminModel,error)
	UpdatePwdByOldPwd(input *dtos.UpdateUserPwdByPwdInput)(int,error)
}

