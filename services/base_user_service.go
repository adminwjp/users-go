package service

import (
	dto "github.com/adminwjp/infrastructure-go/dtos"
	"github.com/adminwjp/users-go/daos"
	"github.com/adminwjp/users-go/dtos"
)
//用户服务基础接口
type BaseUserService interface {
	GetTranction()daos.TranDao
	Clean()
	/**
	根据手机号、邮箱、用户名注册
	*/
	Register(user *dtos.UserInput)(int,error)


	/**
	根据手机号注册
	*/
	RegisterByPhone(user *dtos.UserPhoneInput)(int,error)

	/**
	根据邮箱注册
	*/
	RegisterByEmail(user *dtos.UserEmailInput)(int,error)

	/**
	根据用户名注册
	*/
	RegisterByUserName(user *dtos.UserUserNameInput)(int,error)

	/**
	根据手机号、邮箱、用户名检测账号是否存在
	*/
	Exists(account string,flag dto.AccounType)(int,error)


	/**
	根据手机号检测账号是否存在
	*/
	ExistsByPhone(phone string)(int,error)

	/**
	根据邮箱检测账号是否存在
	*/
	ExistsByEmail(email string)(int,error)

	/**
	根据用户名检测账号是否存在
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

}

