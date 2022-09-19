package cache_bee_impl

import (
	util "github.com/adminwjp/infrastructure-go/utils"
	"github.com/adminwjp/users-go/dtos"
	"github.com/adminwjp/users-go/models"
)

type UserCacheImpl struct {
	BaseUserCacheImpl
}

/*
根据用户id获取用户信息
*/
func  (cache *UserCacheImpl) GetByUserId(userId string) (*models.UserModel,error)   {
	m1,err:=cache.BaseUserCacheImpl.GetByUserId(userId)
	if err!=nil{
		return nil, err
	}
	var m =models.UserModel{}
	util.MapTo(m1,&m)
	return  &m,err
}
/*
根据手机号、邮箱、用户名获取用户信息
*/
func  (cache *UserCacheImpl) GetByAccount(account string) (*models.UserModel,error)   {
	m1,err:=cache.BaseUserCacheImpl.GetByAccount(account)
	if err!=nil{
		return nil, err
	}
	var m =models.UserModel{}
	util.MapTo(m1,&m)
	return  &m,err
}
/*
根据手机号获取用户信息
*/
func  (cache *UserCacheImpl) GetByPhone(phone string) (*models.UserModel,error)   {
	str,e:=cache.Cache.Get(cache.user+"_"+phone).(string)
	if !e{
		return  nil,nil
	}
	return  cache.GetByUserId(str)
}
/*
根据邮箱获取用户信息
*/
func  (cache *UserCacheImpl) GetByEmail(email string) (*models.UserModel,error)   {
	str,e:=cache.Cache.Get(cache.user+"_"+email).(string)
	if !e{
		return  nil,nil
	}
	return  cache.GetByUserId(str)
}
/*
根据用户名获取用户信息
*/
func  (cache *UserCacheImpl) GetByUserName(userName string) (*models.UserModel,error)   {
	str,e:=cache.Cache.Get(cache.user+"_"+userName).(string)
	if !e{
		return  nil,nil
	}
	return  cache.GetByUserId(str)
}
/*
	根据手机号、邮箱、用户名登录
*/
func(cache *UserCacheImpl) Login(user *dtos.UserInput)(*models.UserModel,error){
	m,err:=cache.GetByAccount(user.Account)
	if m!=nil&&m.Pwd==user.Pwd{
		return  nil,err
	}
	return  nil,err
}

/*
	根据手机号登录
*/
func(cache *UserCacheImpl) LoginByPhone(user *dtos.UserPhoneInput)(*models.UserModel,error){
	m,err:=cache.GetByPhone(user.Phone)
	if m!=nil&&m.Pwd==user.Pwd{
		return  nil,err
	}
	return  nil,err
}
/*
	根据邮箱登录
*/
func (cache *UserCacheImpl)LoginByEmail(user *dtos.UserEmailInput)(*models.UserModel,error){
	m ,err:=cache.GetByEmail(user.Email)
	if m!=nil&&m.Pwd==user.Pwd{
		return  nil,err
	}
	return  nil,err
}

/*
	根据用户名登录
*/
func (cache *UserCacheImpl)LoginByUserName(user *dtos.UserUserNameInput)(*models.UserModel,error){
	m ,err:=cache.GetByUserName(user.UserName)
	if m!=nil&&m.Pwd==user.Pwd{
		return  nil,err
	}
	return  nil,err
}


