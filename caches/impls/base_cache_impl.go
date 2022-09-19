package cache_impl

import (
	//不能同名 不然报错
	dto "github.com/adminwjp/infrastructure-go/dtos"
	"github.com/adminwjp/users-go/dtos"
)

//管理员接口 gorm jinzhus 实现
type BaseCacheImpl struct {
	phones map[string]bool
	emails map[string]bool
	userNames map[string]bool
}
func (cache *BaseCacheImpl)IsSupport()bool{
	return  false
}
/*
	根据手机号、邮箱、用户名注册到缓存
*/
func (cache *BaseCacheImpl)Register(account string,flag int)(int,error){
	return 0, nil
}


/*
	根据手机号注册到缓存
*/
func (cache *BaseCacheImpl)RegisterByPhone(phone string)(int,error){
	return 0, nil
}

/*
	根据邮箱注册到缓存
*/
func (cache *BaseCacheImpl)RegisterByEmail(email string)(int,error){
	return 0, nil
}

/*
	根据用户名注册到缓存
*/
func (cache *BaseCacheImpl)RegisterByUserName(userName string)(int,error){
	return 0, nil
}

func (cache *BaseCacheImpl) get(maps map[string]bool,key string)(int,error){
	_,e:=maps[key]
	if e{
		return 1, nil
	}
	return 0, nil
}
/*
	根据手机号、邮箱、用户名检测账号是否存在
*/
func  (cache *BaseCacheImpl) Exists(account string,flag dto.AccounType)(int,error) {
	switch flag {
	case dto.AccounTypeByEamil:
		return cache.get(cache.emails,account)
	case dto.AccounTypeByPhone:
		return cache.get(cache.phones,account)
	case dto.AccounTypeByUsername:
		return cache.get(cache.userNames,account)
	default:
		return  0,nil


	}


}

/*
	根据手机号测账号是否存在
*/
func  (cache *BaseCacheImpl)ExistsByPhone(phone string)(int,error){
	return cache.get(cache.phones,phone)
}

/*
	根据邮箱测账号是否存在
*/
func  (cache *BaseCacheImpl)ExistsByEmail(email string)(int,error){
	return cache.get(cache.emails,email)
}

/*
	根据用户名测账号是否存在
*/
func  (cache *BaseCacheImpl)ExistsByUserName(userName string)(int,error){
	return cache.get(cache.userNames,userName)
}

/*
	根据手机号修改手机号
*/
func  (cache *BaseCacheImpl)UpdatePhone(input *dtos.UpdateUserPhoneInput)(int,error){
	return  0,nil
}


/*
	根据手机号修改邮箱
*/
func  (cache *BaseCacheImpl)UpdateEmailByPhone(input *dtos.UpdateUserEmailByPhoneInput)(int,error){
	return  0,nil
}

/*
	根据邮箱修改邮箱
*/
func  (cache *BaseCacheImpl)UpdateEmailByEmail(input *dtos.UpdateUserEmailInput)(int,error){
	return  0,nil
}


/*
	根据手机号修改密码
*/
func  (cache *BaseCacheImpl)UpdatePwdByPhone(input *dtos.UpdateUserPwdByPhoneInput)(int,error){
	return  0,nil
}

/*
	根据邮箱修改密码
*/
func  (cache *BaseCacheImpl)UpdatePwdByEmail(input *dtos.UpdateUserPwdByEmailInput)(int,error){
	return  0,nil
}
/*
	修改登录次数
*/
func  (cache *BaseCacheImpl)UpdateLoginFailCount(id int64)(int,error){
	return  0,nil
}

func  (cache *BaseCacheImpl)GetLoginFailCount(id int64)(int,error){
	return  0,nil
}

func  (cache *BaseCacheImpl)ResetLoginFailCount(id int64)(int,error){
	return  0,nil
}