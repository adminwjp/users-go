package cache_impl

import (
	"github.com/adminwjp/users-go/dtos"
	"github.com/adminwjp/users-go/models"
)

var AdminCacheInstance=&AdminCacheImpl{}
//管理员接口 gorm jinzhus 实现
type AdminCacheImpl struct {
	BaseCacheImpl
}
/*
根据用户id获取用户信息
 */
func  (adminCache *AdminCacheImpl) GetByUserId(userId string) (*models.AdminModel,error)   {
	return  nil,nil
}
/*
根据手机号、邮箱、用户名获取用户信息
*/
func  (adminCache *AdminCacheImpl) GetByAccount(account string) (*models.AdminModel,error)   {
	return  nil,nil
}
/*
根据手机号获取用户信息
*/
func  (adminCache *AdminCacheImpl) GetByPhone(phone string) (*models.AdminModel,error)   {
	return  nil,nil
}
/*
根据邮箱获取用户信息
*/
func  (adminCache *AdminCacheImpl) GetByEmail(email string) (*models.AdminModel,error)   {
	return  nil,nil
}
/*
根据用户名获取用户信息
*/
func  (adminCache *AdminCacheImpl) GetByUserName(userName string) (*models.AdminModel,error)   {
	return  nil,nil
}
/*
	根据手机号、邮箱、用户名登录
*/
func(adminCache *AdminCacheImpl) Login(user *dtos.UserInput)(*models.AdminModel,error){
	return  nil,nil
}

/*
	根据手机号登录
*/
func(adminCache *AdminCacheImpl) LoginByPhone(user *dtos.UserPhoneInput)(*models.AdminModel,error){
	return  nil,nil
}
/*
	根据邮箱登录
*/
func (adminCache *AdminCacheImpl)LoginByEmail(user *dtos.UserEmailInput)(*models.AdminModel,error){
	return  nil,nil
}

/*
	根据用户名登录
*/
func (adminCache *AdminCacheImpl)LoginByUserName(user *dtos.UserUserNameInput)(*models.AdminModel,error){
	return  nil,nil
}


func (adminCache *AdminCacheImpl)UpdatePwdByOldPwd(input *dtos.UpdateUserPwdByPwdInput)(int,error){
	return 0, nil
}