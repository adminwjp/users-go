package cache_bee_impl

import (
	dto "github.com/adminwjp/infrastructure-go/dtos"
	"github.com/adminwjp/users-go/dtos"
	"github.com/adminwjp/users-go/services"
	"github.com/beego/beego/v2/adapter/cache"
)

type BaseUserCacheImpl struct {
	Cache cache.Cache
	HashService service.HashService
	user string
}
func (cache *BaseUserCacheImpl)IsSupport()bool{
	return  true
}
/*
根据用户id获取用户信息
*/
func  (cache *BaseUserCacheImpl) GetByUserId(userId string) (*dtos.GetUserOutput,error)   {
	var m,_=cache.Cache.Get(cache.user+"_"+userId).(*dtos.GetUserOutput)
	return m,nil
}
/*
根据手机号、邮箱、用户名获取用户信息
*/
func  (cache *BaseUserCacheImpl) GetByAccount(account string) (*dtos.GetUserOutput,error)   {
	userId,e:=cache.Cache.Get(cache.user+"_"+account).(string)
	if e{
		return  cache.GetByUserId(userId)
	}
	return nil,nil
}
/*
根据手机号获取用户信息
*/
func  (cache *BaseUserCacheImpl) GetByPhone(phone string) (*dtos.GetUserOutput,error)   {
	userId,e:=cache.Cache.Get(cache.user+"_"+phone).(string)
	if e{
		return  cache.GetByUserId(userId)
	}
	return nil,nil
}
/*
根据邮箱获取用户信息
*/
func  (cache *BaseUserCacheImpl) GetByEmail(email string) (*dtos.GetUserOutput,error)   {
	userId,e:=cache.Cache.Get(cache.user+"_"+email).(string)
	if e{
		return  cache.GetByUserId(userId)
	}
	return nil,nil
}
/*
根据用户名获取用户信息
*/
func  (cache *BaseUserCacheImpl) GetByUserName(userName string) (*dtos.GetUserOutput,error)   {
	userId,e:=cache.Cache.Get(cache.user+"_"+userName).(string)
	if e{
		return  cache.GetByUserId(userId)
	}
	return nil,nil
}
/*
	根据手机号、邮箱、用户名登录
*/
func(cache *BaseUserCacheImpl) Login(user *dtos.UserInput)(*dtos.GetUserOutput,error){
	m,err:=cache.GetByAccount(user.Account)
	if m!=nil&&m.Pwd==user.Pwd{
		return  nil,err
	}
	return  nil,err
}

/*
	根据手机号登录
*/
func(cache *BaseUserCacheImpl) LoginByPhone(user *dtos.UserPhoneInput)(*dtos.GetUserOutput,error){
	m,err:=cache.GetByPhone(user.Phone)
	if m!=nil&&m.Pwd==user.Pwd{
		return  nil,err
	}
	return  nil,err
}
/*
	根据邮箱登录
*/
func (cache *BaseUserCacheImpl)LoginByEmail(user *dtos.UserEmailInput)(*dtos.GetUserOutput,error){
	m ,err:=cache.GetByEmail(user.Email)
	if m!=nil&&m.Pwd==user.Pwd{
		return  nil,err
	}
	return  nil,err
}

/*
	根据用户名登录
*/
func (cache *BaseUserCacheImpl)LoginByUserName(user *dtos.UserUserNameInput)(*dtos.GetUserOutput,error){
	m ,err:=cache.GetByUserName(user.UserName)
	if m!=nil&&m.Pwd==user.Pwd{
		return  nil,err
	}
	return  nil,err
}





/*
	根据手机号、邮箱、用户名检测账号是否存在
*/
func (cache *BaseUserCacheImpl)Exists(account string,flag dto.AccounType)(int,error){
	e:=  cache.Cache.IsExist(cache.user+account)
	if e{
		return 1,nil
	}
	return 0, nil
}


/*
	根据手机号测账号是否存在
*/
func (cache *BaseUserCacheImpl)ExistsByPhone(phone string)(int,error){
	e:=  cache.Cache.IsExist(cache.user+phone)
	if e{
		return 1,nil
	}
	return 0, nil
}

/*
	根据邮箱测账号是否存在
*/
func (cache *BaseUserCacheImpl)ExistsByEmail(email string)(int,error){
	e:=  cache.Cache.IsExist(cache.user+email)
	if e{
		return 1,nil
	}
	return 0, nil
}

/*
	根据用户名测账号是否存在
*/
func (cache *BaseUserCacheImpl)ExistsByUserName(userName string)(int,error){
	e:=  cache.Cache.IsExist(cache.user+userName)
	if e{
		return 1,nil
	}
	return 0, nil
}
/*
	根据手机号修改手机号
*/
func (cache *BaseUserCacheImpl) UpdatePhone(input *dtos.UpdateUserPhoneInput)(int,error){
	return 0, nil
}


/*
	根据手机号修改邮箱
*/
func (cache *BaseUserCacheImpl)UpdateEmailByPhone(input *dtos.UpdateUserEmailByPhoneInput)(int,error){
	return 0, nil
}

/*
	根据邮箱修改邮箱
*/
func (cache *BaseUserCacheImpl)UpdateEmailByEmail(input *dtos.UpdateUserEmailInput)(int,error){
	return 0, nil
}


/*
	根据手机号修改密码
*/
func (cache *BaseUserCacheImpl)UpdatePwdByPhone(input *dtos.UpdateUserPwdByPhoneInput)(int,error){
	return 0, nil
}

/*
	根据邮箱修改密码
*/
func (cache *BaseUserCacheImpl)UpdatePwdByEmail(input *dtos.UpdateUserPwdByEmailInput)(int,error){
	return 0, nil
}

/*
	修改登录次数
*/
func (cache *BaseUserCacheImpl)UpdateLoginFailCount(id int64)(int,error){
	return 0, nil
}

func (cache *BaseUserCacheImpl)GetLoginFailCount(id int64)(int,error){
	return 0, nil
}

func (cache *BaseUserCacheImpl)ResetLoginFailCount(id int64)(int,error){
	return 0, nil
}