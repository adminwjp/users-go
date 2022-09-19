package service_impl

import (
	"github.com/adminwjp/users-go/caches"
	"github.com/adminwjp/users-go/daos"
	"log"

	"github.com/adminwjp/users-go/dtos"
	"github.com/adminwjp/users-go/models"
)

func (service *UserServiceImpl)Clean()  {
	//admin
	service.BaseCache=nil
	service.BaseDao=nil

	service.Cache=nil
	service.LogDao=nil
	service.Dao=nil

}
//用户服务接口
type UserServiceImpl struct {
	BaseUserServiceImpl
	Dao	daos.UserDao
	Cache caches.UserCache
	LogDao daos.UserLogDao

}
/**
根据手机号、邮箱、用户名登录
*/
func (service *UserServiceImpl)Login(input *dtos.UserInput)(*models.UserModel,error){
	model,err:= service.Cache.Login(input)
	if service.Cache.IsSupport(){
		if err!=nil{
			log.Printf("%s %s cache account login fail: account %s,err => %s",
				service.user,service.cache1,input.Account,err.Error())
			return nil, err
		}
		if model!=nil{
			return  model,err
		}
		if model==nil{
			log.Printf("%s %s cache account login fail(pwd not match): account %s",
				service.user,service.cache1,input.Account)
			return nil, err
		}
	}

	model,err= service.Dao.Login(input)
	if err!=nil{
		log.Printf("%s %s  account login fail: account %s,err => %s",
			service.user,service.dataWay,input.Account,err.Error())
		return nil, err
	}
	log.Printf("%s %s  account login suc: account %s",
		service.user,service.dataWay,input.Account)
	return  model,err
}

/**
根据手机号登录
*/
func (service *UserServiceImpl)LoginByPhone(input *dtos.UserPhoneInput)(*models.UserModel,error){
	model,err:= service.Cache.LoginByPhone(input)
	if service.Cache.IsSupport(){
		if err!=nil{
			log.Printf("%s %s cache phone login fail: phone %s,err => %s",
				service.user,service.cache1,input.Phone,err.Error())
			return nil, err
		}
		if model!=nil{
			return  model,err
		}
		if model==nil{
			log.Printf("%s %s cache phone login fail(pwd not match): phone %s",
				service.user,service.cache1,input.Phone)
			return nil, err
		}
	}

	model,err= service.Dao.LoginByPhone(input)
	if err!=nil{
		log.Printf("%s %s  phone login fail: phone %s,err => %s",
			service.user,service.dataWay,input.Phone,err.Error())
		return nil, err
	}
	log.Printf("%s %s  phone login suc: phone %s",
		service.user,service.dataWay,input.Phone)
	return  model,err
}

/**
根据邮箱登录
*/
func (service *UserServiceImpl)LoginByEmail(input *dtos.UserEmailInput)(*models.UserModel,error){
	model,err:= service.Cache.LoginByEmail(input)
	if service.Cache.IsSupport(){
		if err!=nil{
			log.Printf("%s %s cache email login fail: email %s,err => %s",
				service.user,service.cache1,input.Email,err.Error())
			return nil, err
		}
		if model!=nil{
			return  model,err
		}
		if model==nil{
			log.Printf("%s %s cache email login fail(pwd not match): email %s",
				service.user,service.cache1,input.Email)
			return nil, err
		}
	}

	model,err= service.Dao.LoginByEmail(input)
	if err!=nil{
		log.Printf("%s %s  email login fail: email %s,err => %s",
			service.user,service.dataWay,input.Email,err.Error())
		return nil, err
	}
	log.Printf("%s %s  email login suc: email %s",
		service.user,service.dataWay,input.Email)
	return  model,err
}

/**
根据用户名登录
*/
func (service *UserServiceImpl)LoginByUserName(input *dtos.UserUserNameInput)(*models.UserModel,error){
	model,err:= service.Cache.LoginByUserName(input)
	if service.Cache.IsSupport(){
		if err!=nil{
			log.Printf("%s %s cache user_name login fail: user_name %s,err => %s",
				service.user,service.cache1,input.UserName,err.Error())
			return nil, err
		}
		if model!=nil{
			return  model,err
		}
		if model==nil{
			log.Printf("%s %s cache user_name login fail(pwd not match): user_name %s",
				service.user,service.cache1,input.UserName)
			return nil, err
		}
	}
	model,err= service.Dao.LoginByUserName(input)
	if err!=nil{
		log.Printf("%s %s  user_name login fail: user_name %s,err => %s",
			service.user,service.dataWay,input.UserName,err.Error())
		return nil, err
	}
	log.Printf("%s %s  user_name login suc: user_name %s",
		service.user,service.dataWay,input.UserName)
	return  model,err
}




/*
	根据条件查询用户信息
*/
func (service *UserServiceImpl)List(user *dtos.GetUserInput) ([]models.UserModel,int64,error){
	return service.Dao.List(user)
}

/*
	根据id查询用户信息
*/
func (service *UserServiceImpl)Get(id int64) (*models.UserModel,error){
	return service.Dao.Get(id)
}

/*
	根据条件查询用户日志信息
*/
func (service *UserServiceImpl)ListLogs(user *dtos.GetUserLogInput) ([]models.UserLogModel,int64,error){
	return service.LogDao.List(user)
}
/*
	修改身份认证基本信息
*/
func(service *UserServiceImpl)UpdateAuthBasic(input *dtos.UpdateUserAuthBasicInput)(int,error){
	return service.Dao.UpdateAuthBasic(input)
}