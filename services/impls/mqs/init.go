package service_mq_impl

import (
	mq "github.com/adminwjp/infrastructure-go/mqs"
	"github.com/adminwjp/users-go/daos"
	"github.com/adminwjp/users-go/dtos"
	"github.com/adminwjp/users-go/models"
	service "github.com/adminwjp/users-go/services"
	service_impl "github.com/adminwjp/users-go/services/impls"
)

//var UserMqInstance=&BaseUserServiceImpl{}
var UserMqInstance=&UserMq{}
var AdminMqInstance=&AdminMq{}
var MqInstance mq.Mq
func GetAdminService()  service.AdminService{
	if !AdminMqInstance.initAdmin{
		//service1:=&service_impl.AdminServiceImpl{}
		//AdminMqInstance.AdminServiceImpl=*service1
		AdminMqInstance.Service=service_impl.GetAdminService()
		AdminMqInstance.Init(true)
		AdminMqInstance.Mq=MqInstance
		AdminMqInstance.SubscriptRegister()
		AdminMqInstance.SubscriptRegisterByPhone()
		AdminMqInstance.SubscriptRegisterByEmail()
		AdminMqInstance.SubscriptRegisterByUserName()

		AdminMqInstance.SubscriptUpdateEmailByEmail()
		AdminMqInstance.SubscriptUpdateEmailByPhone()

		AdminMqInstance.SubscriptResetLoginFailCount()
		AdminMqInstance.SubscriptUpdateLoginFailCount()

		AdminMqInstance.SubscriptUpdatePhone()

		AdminMqInstance.SubscriptUpdatePwdByEmail()
		AdminMqInstance.SubscriptUpdatePwdByPhone()
	}
	//service1:=&AdminMqInstance.AdminServiceImpl


	service1:=AdminMqInstance.Service.(*service_impl.AdminServiceImpl)


	service_impl.UpdateAdmin(service1)
	//AdminMqInstance.BaseUserServiceImpl.BaseUserServiceImpl=service1.BaseUserServiceImpl
	AdminMqInstance.Lock=service1.Lock
	AdminMqInstance.Retry=service1.Retry
	AdminMqInstance.BaseDao=service1.BaseDao
	AdminMqInstance.BaseCache=service1.Cache

	//AdminMqInstance.BaseUserServiceImpl=BaseUserServiceImpl{

	//}

	//AdminMqInstance.BaseUserServiceImpl.BaseUserServiceImpl=service_impl.BaseUserServiceImpl{

	//}
	AdminMqInstance.BaseUserServiceImpl.adminDao=service1.Dao
	AdminMqInstance.BaseUserServiceImpl.BaseUserServiceImpl.BaseDao=service1.BaseDao
	AdminMqInstance.BaseUserServiceImpl.BaseUserServiceImpl.BaseCache=service1.BaseCache

	//ambiguous selector AdminMqInstance.Lock
	AdminMqInstance.Lock=service1.Lock

	AdminMqInstance.initAdmin=true


	return 	AdminMqInstance
}
func GetUserService()  service.UserService{
	if !UserMqInstance.initUser{
		UserMqInstance.Service=service_impl.GetUserService()
		UserMqInstance.Init(false)
		UserMqInstance.Mq=MqInstance
		UserMqInstance.SubscriptRegister()
		UserMqInstance.SubscriptRegisterByPhone()
		UserMqInstance.SubscriptRegisterByEmail()
		UserMqInstance.SubscriptRegisterByUserName()

		UserMqInstance.SubscriptUpdateEmailByEmail()
		UserMqInstance.SubscriptUpdateEmailByPhone()

		UserMqInstance.SubscriptResetLoginFailCount()
		UserMqInstance.SubscriptUpdateLoginFailCount()

		UserMqInstance.SubscriptUpdatePhone()

		UserMqInstance.SubscriptUpdatePwdByEmail()
		UserMqInstance.SubscriptUpdatePwdByPhone()
	}
	service1:=UserMqInstance.Service.(*service_impl.UserServiceImpl)
	service_impl.UpdateUser(service1)

	UserMqInstance.BaseUserServiceImpl.userDao=service1.Dao
	//继承 可读性太差了 需要一一赋值 1-n 层 重写
	//ex nil
	//UserMqInstance.BaseUserServiceImpl.BaseUserServiceImpl=service1.BaseUserServiceImpl
	UserMqInstance.Lock=service1.Lock
	UserMqInstance.Retry=service1.Retry

	UserMqInstance.BaseCache=service1.Cache

	//UserMqInstance.BaseUserServiceImpl=BaseUserServiceImpl{

	//}
	//AdminMqInstance.BaseUserServiceImpl.BaseUserServiceImpl=service_impl.BaseUserServiceImpl{

	//}
	UserMqInstance.BaseUserServiceImpl.BaseDao=service1.BaseDao
	UserMqInstance.BaseUserServiceImpl.userDao=service1.Dao
	UserMqInstance.BaseUserServiceImpl.BaseUserServiceImpl.BaseDao=service1.BaseDao
	//nil
	UserMqInstance.BaseUserServiceImpl.BaseUserServiceImpl.BaseCache=service1.BaseCache


	UserMqInstance.initUser=true
	return 	UserMqInstance
}

type AdminMq struct {
	BaseUserServiceImpl
	Service service.AdminService
	//service.AdminService
	// service_impl.AdminServiceImpl
	initAdmin bool
}

func (service1 *AdminMq)GetTranction()daos.TranDao{
	return service1.Service.GetTranction()
}
func (service1 *AdminMq)Clean()  {
	service1.Service.Clean()

}
/**
根据手机号、邮箱、用户名登录
*/
func (service1 *AdminMq)Login(input *dtos.UserInput)(*models.AdminModel,error){
	return service1.Service.Login(input)
}

/**
根据手机号登录
*/
func (service1 *AdminMq)LoginByPhone(input *dtos.UserPhoneInput)(*models.AdminModel,error){
	return service1.Service.LoginByPhone(input)
}

/**
根据邮箱登录
*/
func (service1 *AdminMq)LoginByEmail(input *dtos.UserEmailInput)(*models.AdminModel,error){
	return service1.Service.LoginByEmail(input)
}

/**
根据用户名登录
*/
func (service1 *AdminMq)LoginByUserName(input *dtos.UserUserNameInput)(*models.AdminModel,error){
	return service1.Service.LoginByUserName(input)
}

/*
	根据旧密码修改密码
*/
func (service1 *AdminMq)UpdatePwdByOldPwd(input *dtos.UpdateUserPwdByPwdInput)(int,error){
	return service1.Service.UpdatePwdByOldPwd(input)
}



/*
	根据条件查询用户信息
*/
func (service1 *AdminMq)List(user *dtos.GetAdminInput) ([]models.AdminModel,int64,error){
	return service1.Service.List(user)
}

/*
	根据id查询用户信息
*/
func (service1 *AdminMq)Get(id int64) (*models.AdminModel,error){
	return service1.Service.Get(id)
}

/*
	根据条件查询用户日志信息
*/
func (service1 *AdminMq)ListLogs(user *dtos.GetUserLogInput) ([]models.AdminLogModel,int64,error){
	return service1.Service.ListLogs(user)
}
type UserMq struct {
	BaseUserServiceImpl
	//service_impl.UserServiceImpl
	//service.UserService
	Service service.UserService

	initUser bool
}
/**
根据手机号、邮箱、用户名登录
*/
func (service1 *UserMq)Login(input *dtos.UserInput)(*models.UserModel,error){
	return service1.Service.Login(input)
}

/**
根据手机号登录
*/
func (service1 *UserMq)LoginByPhone(input *dtos.UserPhoneInput)(*models.UserModel,error){
	return service1.Service.LoginByPhone(input)
}

/**
根据邮箱登录
*/
func (service1 *UserMq)LoginByEmail(input *dtos.UserEmailInput)(*models.UserModel,error){
	return service1.Service.LoginByEmail(input)
}

/**
根据用户名登录
*/
func (service1 *UserMq)LoginByUserName(input *dtos.UserUserNameInput)(*models.UserModel,error){
	return service1.Service.LoginByUserName(input)
}




/*
	根据条件查询用户信息
*/
func (service1 *UserMq)List(user *dtos.GetUserInput) ([]models.UserModel,int64,error){
	return service1.Service.List(user)
}

/*
	根据id查询用户信息
*/
func (service1 *UserMq)Get(id int64) (*models.UserModel,error){
	return service1.Service.Get(id)
}

/*
	根据条件查询用户日志信息
*/
func (service1 *UserMq)ListLogs(user *dtos.GetUserLogInput) ([]models.UserLogModel,int64,error){
	return service1.Service.ListLogs(user)
}
/*
	修改身份认证基本信息
*/
func(service1 *UserMq)UpdateAuthBasic(input *dtos.UpdateUserAuthBasicInput)(int,error){
	return service1.Service.UpdateAuthBasic(input)
}
func (service1 *UserMq)GetTranction()daos.TranDao{
	return service1.Service.GetTranction()
}
func (service1 *UserMq)Clean()  {
	service1.Service.Clean()

}
