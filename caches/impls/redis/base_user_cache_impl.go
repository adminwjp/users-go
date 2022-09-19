package cache_redis_impl

import (
	cache_redis "github.com/adminwjp/infrastructure-go/caches/redises"
	dto "github.com/adminwjp/infrastructure-go/dtos"
	util "github.com/adminwjp/infrastructure-go/utils"
	"github.com/adminwjp/users-go/dtos"
	"github.com/adminwjp/users-go/services"
	"github.com/go-redis/redis"
)

type BaseUserCacheImpl struct {
	cache_redis.RedisCache
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
	var m dtos.GetUserOutput
	str,err:=cache.Get(cache.user+"_"+userId)
	if err!=nil{
		return nil ,err
	}
	err=util.SeriablizeUtil.JsonDesriablizeObject(&m,[]byte(str))
	return  &m,err
}
/*
根据手机号、邮箱、用户名获取用户信息
*/
func  (cache *BaseUserCacheImpl) GetByAccount(account string) (*dtos.GetUserOutput,error)   {
	str,err:=cache.Get(cache.user+"_"+account)
	if err!=nil{
		return  nil,err
	}
	/*if str==""{
		log.Println("")
		return  nil,err
	}*/
	return  cache.GetByUserId(str)
}
/*
根据手机号获取用户信息
*/
func  (cache *BaseUserCacheImpl) GetByPhone(phone string) (*dtos.GetUserOutput,error)   {
	str,err:=cache.Get(cache.user+"_"+phone)
	if err!=nil{
		return  nil,err
	}
	return  cache.GetByUserId(str)
}
/*
根据邮箱获取用户信息
*/
func  (cache *BaseUserCacheImpl) GetByEmail(email string) (*dtos.GetUserOutput,error)   {
	str,err:=cache.Get(cache.user+"_"+email)
	if err!=nil{
		return  nil,err
	}
	return  cache.GetByUserId(str)
}
/*
根据用户名获取用户信息
*/
func  (cache *BaseUserCacheImpl) GetByUserName(userName string) (*dtos.GetUserOutput,error)   {
	str,err:=cache.Get(cache.user+"_"+userName)
	if err!=nil{
		return  nil,err
	}
	return  cache.GetByUserId(str)
}
/*
	根据手机号、邮箱、用户名登录
*/
func(cache *BaseUserCacheImpl) Login(user dtos.UserInput)(*dtos.GetUserOutput,error){
	m,err:=cache.GetByAccount(user.Account)
	if m!=nil&&m.Pwd==user.Pwd{
		return  nil,err
	}
	return  nil,err
}

/*
	根据手机号登录
*/
func(cache *BaseUserCacheImpl) LoginByPhone(user dtos.UserPhoneInput)(*dtos.GetUserOutput,error){
	m,err:=cache.GetByPhone(user.Phone)
	if m!=nil&&m.Pwd==user.Pwd{
		return  nil,err
	}
	return  nil,err
}
/*
	根据邮箱登录
*/
func (cache *BaseUserCacheImpl)LoginByEmail(user dtos.UserEmailInput)(*dtos.GetUserOutput,error){
	m ,err:=cache.GetByEmail(user.Email)
	if m!=nil&&m.Pwd==user.Pwd{
		return  nil,err
	}
	return  nil,err
}

/*
	根据用户名登录
*/
func (cache *BaseUserCacheImpl)LoginByUserName(user dtos.UserUserNameInput)(*dtos.GetUserOutput,error){
	m ,err:=cache.GetByUserName(user.UserName)
	if m!=nil&&m.Pwd==user.Pwd{
		return  nil,err
	}
	return  nil,err
}


/*
	根据手机号、邮箱、用户名注册到缓存
*/
func (cache *BaseUserCacheImpl)Register(account string,flag int)(int,error){
	return 0, nil
}


/*
	根据手机号注册到缓存
*/
func (cache *BaseUserCacheImpl)RegisterByPhone(phone string)(int,error){
	return 0, nil
}

/*
	根据邮箱注册到缓存
*/
func (cache *BaseUserCacheImpl)RegisterByEmail(email string)(int,error){
	return 0, nil
}

/*
	根据用户名注册到缓存
*/
func (cache *BaseUserCacheImpl)RegisterByUserName(userName string)(int,error){
	return 0, nil
}


/*
	根据手机号、邮箱、用户名检测账号是否存在
*/
func (cache *BaseUserCacheImpl)Exists(account string,flag dto.AccounType)(int,error){
	var hash=cache.HashService.Hash(account)
	var val *redis.IntCmd
	if flag==dto.AccounTypeByEamil{
		val=cache.Client.GetBit(cache.user+"_"+"email",hash)
	}else  if flag==dto.AccounTypeByPhone{
		val=cache.Client.GetBit(cache.user+"_phone",hash)
	}else {
		val=cache.Client.GetBit(cache.user+"_user_name",hash)
	}
	return intResult(val)
}


/*
	根据手机号测账号是否存在
*/
func (cache *BaseUserCacheImpl)ExistsByPhone(phone string)(int,error){
	var hash=cache.HashService.Hash(phone)
	val:=cache.Client.GetBit(cache.user+"_phone",hash)
	return intResult(val)
}

/*
	根据邮箱测账号是否存在
*/
func (cache *BaseUserCacheImpl)ExistsByEmail(email string)(int,error){
	var hash=cache.HashService.Hash(email)
	val:=cache.Client.GetBit(cache.user+"_email",hash)
	return intResult(val)
}

/*
	根据用户名测账号是否存在
*/
func (cache *BaseUserCacheImpl)ExistsByUserName(userName string)(int,error){
	var hash=cache.HashService.Hash(userName)
	val:=cache.Client.GetBit(cache.user+"_user_name",hash)
	return intResult(val)
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