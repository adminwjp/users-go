package rpc_thrift_impl

import (
	"context"
	"github.com/adminwjp/users-go/services"
	t "github.com/adminwjp/users-go/rpcs/gen-go/users"
	util "github.com/adminwjp/infrastructure-go/utils"
)

//用户服务 thrift 接口 实现
type UserThriftServiceImpl struct {
	BaseUserThriftServiceImpl
   Service func() service.UserService
}

/**
根据手机号、邮箱、用户名登录
*/
func(user1 *UserThriftServiceImpl) Login(ctx context.Context, user *t.UserInputThrift) (_r *t.UserOuputThrift, _err error){
	userDto:=ToUserDto(user)
	service:=user1.Service()
	//service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	userModel,_:=service.Login(userDto)
	out := new(t.UserOuputThrift)
	if userModel==nil{return out,  nil}
	util.MapTo(userModel,out)
	out.ID=userModel.Id
	return out, nil
}

/**
根据手机号登录
*/
func(user1 *UserThriftServiceImpl) LoginByPhone(ctx context.Context, user *t.UserPhoneInputThrift) (_r *t.UserOuputThrift, _err error){
	userDto:=ToUserPhoneDto(user)
	service:=user1.Service()
	//service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	userModel,_:=service.LoginByPhone(userDto)
	out := new(t.UserOuputThrift)
	if userModel==nil{return out,  nil}
	util.MapTo(userModel,out)
	out.ID=userModel.Id
	return out, nil
}

/**
根据邮箱登录
*/
func(user1 *UserThriftServiceImpl) LoginByEmail(ctx context.Context, user *t.UserEmailInputThrift) (_r *t.UserOuputThrift, _err error){
	userDto:=ToUserEmailDto(user)
	service:=user1.Service()
	//service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	userModel,_:=service.LoginByEmail(userDto)
	out := new(t.UserOuputThrift)
	if userModel==nil{return out,  nil}
	util.MapTo(userModel,out)
	out.ID=userModel.Id
	return out, nil
}

/**
根据用户名登录
*/
func(user1 *UserThriftServiceImpl) LoginByUserName(ctx context.Context, user *t.UserUserNameInputThrift) (_r *t.UserOuputThrift, _err error){
	userDto:=ToUserUserNameDto(user)
	service:=user1.Service()
	//service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	userModel,_:=service.LoginByUserName(userDto)
	out := new(t.UserOuputThrift)
	if userModel==nil{return out,  nil}
	util.MapTo(userModel,out)
	out.ID=userModel.Id
	return out, nil
}

/*
	根据id查询用户信息
*/
func(user1 *UserThriftServiceImpl) Get(ctx context.Context, user *t.IdThrift) (_r *t.UserOuputThrift, _err error){
	service:=user1.Service()
	//service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	userModel,_:=service.Get(user.ID)
	out := new(t.UserOuputThrift)
	if userModel==nil{return out,  nil}
	util.MapTo(userModel,out)
	out.ID=userModel.Id
	return out, nil
}
func(user1 *UserThriftServiceImpl) GetAuthBasic(ctx context.Context, user *t.IdThrift) (_r *t.UpdateAuthBasicInputThrift, _err error){
	out := new(t.UpdateAuthBasicInputThrift)
	return out, nil
}

/*
	修改身份认证基本信息
*/
func(user1 *UserThriftServiceImpl) UpdateAuthBasic(ctx context.Context, user *t.UpdateAuthBasicInputThrift) (_r *t.ResultOuputThrift, _err error){
	service:=user1.Service()
	service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	res,_:=service.UpdateAuthBasic(ToUpdateUserAuthBasicDto(user))
	out := new(t.ResultOuputThrift)
	out.Result_=int32(res)
	return out, nil
}

