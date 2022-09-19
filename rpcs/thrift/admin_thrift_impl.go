package rpc_thrift_impl

import (
	"context"
	dto "github.com/adminwjp/infrastructure-go/dtos"
	"github.com/adminwjp/users-go/services"
	t "github.com/adminwjp/users-go/rpcs/gen-go/users"
	util "github.com/adminwjp/infrastructure-go/utils"
)

//管理员服务 thrift 接口 实现
type AdminThriftServiceImpl struct {
	BaseUserThriftServiceImpl
	Service func()service.AdminService
}

/**
根据手机号、邮箱、用户名登录
*/
func(admin *AdminThriftServiceImpl) Login(ctx context.Context, user *t.UserInputThrift) (_r *t.AdminOuputThrift, _err error){
	userDto:=ToUserDto(user)
	service:=admin.Service()
	//service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	adminModel,_:=service.Login(userDto)
	out := new(t.AdminOuputThrift)
	if adminModel==nil{return out,  nil}
	util.MapTo(adminModel,out)
	out.ID=adminModel.Id
	return out, nil
}

/**
根据手机号登录
*/
func(admin *AdminThriftServiceImpl) LoginByPhone(ctx context.Context, user *t.UserPhoneInputThrift) (_r *t.AdminOuputThrift, _err error){
	userDto:=ToUserPhoneDto(user)
	service:=admin.Service()
	//service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	adminModel,_:=service.LoginByPhone(userDto)
	out := new(t.AdminOuputThrift)
	if adminModel==nil{return out,  nil}
	util.MapTo(adminModel,out)
	out.ID=adminModel.Id
	return out, nil
}

/**
根据邮箱登录
*/
func(admin *AdminThriftServiceImpl) LoginByEmail(ctx context.Context, user *t.UserEmailInputThrift) (_r *t.AdminOuputThrift, _err error){
	userDto:=ToUserEmailDto(user)
	service:=admin.Service()
	//service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	adminModel,_:=service.LoginByEmail(userDto)
	out := new(t.AdminOuputThrift)
	if adminModel==nil{return out,  nil}
	util.MapTo(adminModel,out)
	out.ID=adminModel.Id
	return out, nil
}

/**
根据用户名登录
*/
func(admin *AdminThriftServiceImpl) LoginByUserName(ctx context.Context, user *t.UserUserNameInputThrift) (_r *t.AdminOuputThrift, _err error){
	userDto:=ToUserUserNameDto(user)
	service:=admin.Service()
	//service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	adminModel,_:=service.LoginByUserName(userDto)
	out := new(t.AdminOuputThrift)
	if adminModel==nil{return out,  nil}
	util.MapTo(adminModel,out)
	out.ID=adminModel.Id
	return out, nil
}

/*
	根据id查询用户信息
*/
func(admin *AdminThriftServiceImpl) Get(ctx context.Context, user *t.IdThrift) (_r *t.AdminOuputThrift, _err error){
	service:=admin.Service()
	//service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	userModel,_:=service.Get(user.ID)
	out := new(t.AdminOuputThrift)
	if userModel==nil{return out,  nil}
	util.MapTo(userModel,out)
	out.ID=userModel.Id
	return out, nil
}

/**
根据手机号、邮箱、用户名检测账号是否存在
*/
func(admin *AdminThriftServiceImpl) Exists(ctx context.Context, user *t.AccountInputThrift) (_r *t.ExistsOuputThrift, _err error){
	service:=admin.Service()
	//service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	res,_:=service.Exists(user.Account,dto.AccounType(user.Flag))
	out := new(t.ExistsOuputThrift)
	out.Exists=res>0
	return out, nil
}


/*
	根据旧密码修改密码
*/
func(admin *AdminThriftServiceImpl) UpdatePwd(ctx context.Context, user *t.UpdatePwdInputThrift) (_r *t.ResultOuputThrift, _err error){
	service:=admin.Service()
	service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	res,_:=service.UpdatePwdByOldPwd(ToUpdateUserPwdByPwdDto(user))
	out := new(t.ResultOuputThrift)
	out.Result_=int32(res)
	return out, nil
}


