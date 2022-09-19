package cache_impl

import (
	"github.com/adminwjp/users-go/dtos"
	"github.com/adminwjp/users-go/models"
)

var UserCacheInstance=&UserCacheImpl{}
//管理员接口 gorm jinzhus 实现
type UserCacheImpl struct {
	BaseCacheImpl
}
/*
根据用户id获取用户信息
*/
func  (userCache *UserCacheImpl) GetByUserId(userId string) (*models.UserModel,error)   {
	return  nil,nil
}
/*
根据手机号、邮箱、用户名获取用户信息
*/
func  (userCache *UserCacheImpl) GetByAccount(account string) (*models.UserModel,error)   {
	return  nil,nil
}
/*
根据手机号获取用户信息
*/
func  (userCache *UserCacheImpl) GetByPhone(phone string) (*models.UserModel,error)   {
	return  nil,nil
}
/*
根据邮箱获取用户信息
*/
func  (userCache *UserCacheImpl) GetByEmail(email string) (*models.UserModel,error)   {
	return  nil,nil
}
/*
根据用户名获取用户信息
*/
func  (userCache *UserCacheImpl) GetByUserName(userName string) (*models.UserModel,error)   {
	return  nil,nil
}
/*
	根据手机号、邮箱、用户名登录
*/
func(userCache *UserCacheImpl) Login(user *dtos.UserInput)(*models.UserModel,error){
	return  nil,nil
}

/*
	根据手机号登录
*/
func(userCache *UserCacheImpl) LoginByPhone(user *dtos.UserPhoneInput)(*models.UserModel,error){
	return  nil,nil
}
/*
	根据邮箱登录
*/
func (userCache *UserCacheImpl)LoginByEmail(user *dtos.UserEmailInput)(*models.UserModel,error){
	return  nil,nil
}

/*
	根据用户名登录
*/
func (userCache *UserCacheImpl)LoginByUserName(user *dtos.UserUserNameInput)(*models.UserModel,error){
	return  nil,nil
}




