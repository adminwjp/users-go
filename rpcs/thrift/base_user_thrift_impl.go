package rpc_thrift_impl

import (
	"context"
	dto "github.com/adminwjp/infrastructure-go/dtos"
	t "github.com/adminwjp/users-go/rpcs/gen-go/users"
	"github.com/adminwjp/users-go/services"
)

//用户服务 thrift 基类 接口 实现
type BaseUserThriftServiceImpl struct {
	BaseService func() service.BaseUserService
}

/**
根据手机号、邮箱、用户名注册
*/
func(user1 *BaseUserThriftServiceImpl) Register1(ctx context.Context, user *t.UserInputThrift) (_r *t.IdThrift, _err error){
	userDto:=ToUserDto(user)
	service:=user1.BaseService()
	service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	service.Register(userDto)
	out := new(t.IdThrift)
	out.ID=userDto.Id
	return out, nil
}

/**
根据手机号注册
*/
func(user1 *BaseUserThriftServiceImpl) RegisterByPhone(ctx context.Context, user *t.UserPhoneInputThrift) (_r *t.IdThrift, _err error){
	userDto:=ToUserPhoneDto(user)
	service:=user1.BaseService()
	service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	service.RegisterByPhone(userDto)
	out := new(t.IdThrift)
	out.ID=userDto.Id
	return out, nil
}

/**
根据邮箱注册
*/
func(user1 *BaseUserThriftServiceImpl) RegisterByEmail(ctx context.Context, user *t.UserEmailInputThrift) (_r *t.IdThrift, _err error){
	userDto:=ToUserEmailDto(user)
	service:=user1.BaseService()
	service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	service.RegisterByEmail(userDto)
	out := new(t.IdThrift)
	out.ID=userDto.Id
	return out, nil
}

/**
根据用户名注册
*/
func(user1 *BaseUserThriftServiceImpl) RegisterByUserName(ctx context.Context, user *t.UserUserNameInputThrift) (_r *t.IdThrift, _err error){
	userDto:=ToUserUserNameDto(user)
	service:=user1.BaseService()
	service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	service.RegisterByUserName(userDto)
	out := new(t.IdThrift)
	out.ID=userDto.Id
	return out, nil
}



/**
根据手机号、邮箱、用户名检测账号是否存在
*/
func(user1 *BaseUserThriftServiceImpl) Exists(ctx context.Context, user *t.AccountInputThrift) (_r *t.ExistsOuputThrift, _err error){
	service:=user1.BaseService()
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

/**
根据手机号检测账号是否存在
*/
func(user1 *BaseUserThriftServiceImpl) ExistsPhone(ctx context.Context, user *t.PhoneInputThrift) (_r *t.ExistsOuputThrift, _err error){
	service:=user1.BaseService()
	//service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	res,_:=service.ExistsByPhone(user.Phone)
	out := new(t.ExistsOuputThrift)
	out.Exists=res>0
	return out, nil
}

/**
根据邮箱检测账号是否存在
*/
func(user1 *BaseUserThriftServiceImpl) ExistsEmail(ctx context.Context, user *t.EmailInputThrift) (_r *t.ExistsOuputThrift, _err error){
	service:=user1.BaseService()
	//service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	res,_:=service.ExistsByEmail(user.Email)
	out := new(t.ExistsOuputThrift)
	out.Exists=res>0
	return out, nil
}

/**
根据用户名检测账号是否存在
*/
func(user1 *BaseUserThriftServiceImpl) ExistsUserName(ctx context.Context, user *t.UserNameInputThrift) (_r *t.ExistsOuputThrift, _err error){
	service:=user1.BaseService()
	//service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	res,_:=service.ExistsByUserName(user.UserName)
	out := new(t.ExistsOuputThrift)
	out.Exists=res>0
	return out, nil
}

/*
	根据邮箱修改邮箱
*/
func(user1 *BaseUserThriftServiceImpl) UpdateEmail(ctx context.Context, user *t.UpdateEmailInputThrift) (_r *t.ResultOuputThrift, _err error){
	service:=user1.BaseService()
	service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	res,_:=service.UpdateEmailByEmail(ToUpdateUserEmailDto(user))
	out := new(t.ResultOuputThrift)
	out.Result_=int32(res)
	return out, nil
}

/*
	根据手机号修改邮箱
*/
func(user1 *BaseUserThriftServiceImpl) UpdateEmailByPhone(ctx context.Context, user *t.UpdateEmailByPhoneInputThrift) (_r *t.ResultOuputThrift, _err error){
	service:=user1.BaseService()
	service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	res,_:=service.UpdateEmailByPhone(ToUpdateUserEmailByPhoneDto(user))
	out := new(t.ResultOuputThrift)
	out.Result_=int32(res)
	return out, nil
}

/*
	根据手机号修改手机号
*/
func(user1 *BaseUserThriftServiceImpl) UpdatePhone(ctx context.Context, user *t.UpdatePhoneInputThrift) (_r *t.ResultOuputThrift, _err error){
	service:=user1.BaseService()
	service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	res,_:=service.UpdatePhone(ToUpdateUserPhoneDto(user))
	out := new(t.ResultOuputThrift)
	out.Result_=int32(res)
	return out, nil
}

/*
	根据手机号修改密码
*/
func(user1 *BaseUserThriftServiceImpl) UpdatePwdByPhone(ctx context.Context, user *t.UpdatePwdByPhoneInputThrift) (_r *t.ResultOuputThrift, _err error){
	service:=user1.BaseService()
	service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	res,_:=service.UpdatePwdByPhone(ToUpdateUserPwdByPhoneDto(user))
	out := new(t.ResultOuputThrift)
	out.Result_=int32(res)
	return out, nil
}

/*
	根据邮箱修改密码
*/
func(user1 *BaseUserThriftServiceImpl) UpdatePwdByEmail(ctx context.Context, user *t.UpdatePwdByEmailInputThrift) (_r *t.ResultOuputThrift, _err error){
	service:=user1.BaseService()
	service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	res,_:=service.UpdatePwdByEmail(ToUpdateUserPwdByEmailDto(user))
	out := new(t.ResultOuputThrift)
	out.Result_=int32(res)
	return out, nil
}



