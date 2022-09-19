package cache_bee_impl

import (
	util "github.com/adminwjp/infrastructure-go/utils"
	"github.com/adminwjp/users-go/dtos"
	"github.com/adminwjp/users-go/models"
)

//管理员接口 gorm jinzhus 实现
type AdminCacheImpl struct {
	BaseUserCacheImpl
}
/*
根据用户id获取用户信息
 */
func  (cache *AdminCacheImpl) GetByUserId(userId string) (*models.AdminModel,error)   {
	m1,e:=cache.Cache.Get(cache.user+"_"+userId).(*dtos.GetUserOutput)
	if !e{
			return nil ,nil
	}
	var m =models.AdminModel{}
	util.MapTo(m1,&m)
	return  &m,nil
}
/*
根据手机号、邮箱、用户名获取用户信息
*/
func  (cache *AdminCacheImpl) GetByAccount(account string) (*models.AdminModel,error)   {
	str,e:=cache.Cache.Get(cache.user+"_"+account).(string)
	if !e{
		return  nil,nil
	}
	return  cache.GetByUserId(str)
}
/*
根据手机号获取用户信息
*/
func  (cache *AdminCacheImpl) GetByPhone(phone string) (*models.AdminModel,error)   {
	str,e:=cache.Cache.Get(cache.user+"_"+phone).(string)
	if !e{
		return  nil,nil
	}
	return  cache.GetByUserId(str)
}
/*
根据邮箱获取用户信息
*/
func  (cache *AdminCacheImpl) GetByEmail(email string) (*models.AdminModel,error)   {
	str,e:=cache.Cache.Get(cache.user+"_"+email).(string)
	if !e{
		return  nil,nil
	}
	return  cache.GetByUserId(str)
}
/*
根据用户名获取用户信息
*/
func  (cache *AdminCacheImpl) GetByUserName(userName string) (*models.AdminModel,error)   {
	str,e:=cache.Cache.Get(cache.user+"_"+userName).(string)
	if !e{
		return  nil,nil
	}
	return  cache.GetByUserId(str)
}
/*
	根据手机号、邮箱、用户名登录
*/
func(cache *AdminCacheImpl) Login(user *dtos.UserInput)(*models.AdminModel,error){
	 m,err:=cache.GetByAccount(user.Account)
	if m!=nil&&m.Pwd==user.Pwd{
		return  nil,err
	}
	return  nil,err
}

/*
	根据手机号登录
*/
func(cache *AdminCacheImpl) LoginByPhone(user *dtos.UserPhoneInput)(*models.AdminModel,error){
	 m,err:=cache.GetByPhone(user.Phone)
	if m!=nil&&m.Pwd==user.Pwd{
		return  nil,err
	}
	return  nil,err
}
/*
	根据邮箱登录
*/
func (cache *AdminCacheImpl)LoginByEmail(user *dtos.UserEmailInput)(*models.AdminModel,error){
	 m ,err:=cache.GetByEmail(user.Email)
	if m!=nil&&m.Pwd==user.Pwd{
		return  nil,err
	}
	return  nil,err
}

/*
	根据用户名登录
*/
func (cache *AdminCacheImpl)LoginByUserName(user *dtos.UserUserNameInput)(*models.AdminModel,error){
	m ,err:=cache.GetByUserName(user.UserName)
	if m!=nil&&m.Pwd==user.Pwd{
		return  nil,err
	}
	return  nil,err
}

