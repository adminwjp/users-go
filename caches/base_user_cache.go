package caches

import (
	"github.com/adminwjp/users-go/dtos"
	dto "github.com/adminwjp/infrastructure-go/dtos"
)

//用户缓存接口
type BaseUserCache interface {

 	IsSupport()bool
	/*
		根据手机号、邮箱、用户名注册到缓存
	*/
	Register(account string,flag int)(int,error)


	/*
		根据手机号注册到缓存
	*/
	RegisterByPhone(phone string)(int,error)

	/*
		根据邮箱注册到缓存
	*/
	RegisterByEmail(email string)(int,error)

	/*
		根据用户名注册到缓存
	*/
	RegisterByUserName(userName string)(int,error)

	/*
		根据手机号、邮箱、用户名检测账号是否存在
	*/
	Exists(account string,flag dto.AccounType)(int,error)


	/*
		根据手机号测账号是否存在
	*/
	ExistsByPhone(phone string)(int,error)

	/*
		根据邮箱测账号是否存在
	*/
	ExistsByEmail(email string)(int,error)

	/*
		根据用户名测账号是否存在
	*/
	ExistsByUserName(userName string)(int,error)

	/*
		根据手机号修改手机号
	*/
	UpdatePhone(input *dtos.UpdateUserPhoneInput)(int,error)


	/*
		根据手机号修改邮箱
	*/
	UpdateEmailByPhone(input *dtos.UpdateUserEmailByPhoneInput)(int,error)

	/*
		根据邮箱修改邮箱
	*/
	UpdateEmailByEmail(input *dtos.UpdateUserEmailInput)(int,error)


	/*
		根据手机号修改密码
	*/
	UpdatePwdByPhone(input *dtos.UpdateUserPwdByPhoneInput)(int,error)

	/*
		根据邮箱修改密码
	*/
	UpdatePwdByEmail(input *dtos.UpdateUserPwdByEmailInput)(int,error)

	/*
		修改登录次数
	*/
	UpdateLoginFailCount(id int64)(int,error)

	GetLoginFailCount(id int64)(int,error)

	ResetLoginFailCount(id int64)(int,error)
}
