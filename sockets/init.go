package sockets

import (
	"github.com/adminwjp/infrastructure-go/utils"
)

var AdminSInstance=&AdminS{}
var UserSInstance=&UserS{}
var SocketInstance=&utils.SocketUtil{}
func Init()  {

	AdminSInstance.user="admin"
	UserSInstance.user="user"
	funs:=make(map[string]func(msg string)string,100)
	AdminSInstance.funs=funs
	AdminSInstance.Socket=SocketInstance
	init1(AdminSInstance)
	AdminSInstance.UpdatePwdByOldPwd()
	AdminSInstance.Login()
	AdminSInstance.LoginByPhone()
	AdminSInstance.LoginByEmail()
	AdminSInstance.LoginByUserName()
	AdminSInstance.Get()

	UserSInstance.funs=funs
	UserSInstance.Socket=SocketInstance
	init1(UserSInstance)
	UserSInstance.UpdateAuthBasic()
	UserSInstance.Login()
	UserSInstance.LoginByPhone()
	UserSInstance.LoginByEmail()
	UserSInstance.LoginByUserName()
	UserSInstance.Get()

	//server listen client
	RegisterServer(SocketInstance,funs)
}

func init1(user IUserS)  {
	user.Register()
	user.RegisterByPhone()
	user.RegisterByEmail()
	user.RegisterByUserName()

	user.Exists()
	user.ExistsByPhone()
	user.ExistsByPhone()
	user.ExistsByUserName()

	user.UpdatePhone()

	user.UpdateEmailByPhone()
	user.UpdateEmailByEmail()

	user.UpdatePwdByPhone()
	user.UpdatePwdByEmail()
}
type IUserS interface {
	/**
	根据手机号、邮箱、用户名注册
	*/
	Register()


	/**
	根据手机号注册
	*/
	RegisterByPhone()

	/**
	根据邮箱注册
	*/
	RegisterByEmail()

	/**
	根据用户名注册
	*/
	RegisterByUserName()

	/**
	根据手机号、邮箱、用户名检测账号是否存在
	*/
	Exists()


	/**
	根据手机号检测账号是否存在
	*/
	ExistsByPhone()

	/**
	根据邮箱检测账号是否存在
	*/
	ExistsByEmail()

	/**
	根据用户名检测账号是否存在
	*/
	ExistsByUserName()

	/*
		根据手机号修改手机号
	*/
	UpdatePhone()


	/*
		根据手机号修改邮箱
	*/
	UpdateEmailByPhone()

	/*
		根据邮箱修改邮箱
	*/
	UpdateEmailByEmail()


	/*
		根据手机号修改密码
	*/
	UpdatePwdByPhone()

	/*
		根据邮箱修改密码
	*/
	UpdatePwdByEmail()
}