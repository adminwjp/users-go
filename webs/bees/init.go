package web_bee_controller

import (
	"github.com/adminwjp/users-go/inits"
	service "github.com/adminwjp/users-go/services"
	web "github.com/adminwjp/users-go/webs"
)

var BeeRouterImpl =&BeeRouter{}
type IMCtl interface{
	M()
}
var userCtrl *UserCtl
var adminCtrl *AdminCtl
var roleCtrl =&RoleCtl{}
var smsCtrl =&SmsCtl{}
var emailCtrl =&EmailCtl{}
var payCtrl =&PayCtl{}
var rpcCtl =&RpcCtl{}
var config1Ctl =&Config1Ctl{}
//多线程 无法控制 只能 new 要么变复制
func Init()  {
	userCtrl=&UserCtl{}
	adminCtrl=&AdminCtl{}

	//最好每个赋值 不然坑嗲 组合继承
	userCtrl.Service= inits.CreateServiceInstance.UserService
	//gin web
	userCtrl.BaseUserCtl.UserCtrl=&web.BaseUserCtrl{}
	userCtrl.BaseUserCtl.UserCtrl.Service=func() service.BaseUserService {
		return  inits.CreateServiceInstance.UserService()
		//return userCtrl.Service()
	}
	userCtrl.UserCtrl=&web.UserCtrl{Service: inits.CreateServiceInstance.UserService}
	userCtrl.UserCtrl.BaseUserCtrl=*userCtrl.BaseUserCtl.UserCtrl


	adminCtrl.Service= inits.CreateServiceInstance.AdminService
	adminCtrl.BaseUserCtl.UserCtrl=&web.BaseUserCtrl{}
	adminCtrl.BaseUserCtl.UserCtrl.Service=func() service.BaseUserService {
		return  inits.CreateServiceInstance.AdminService()
		//return adminCtrl.Service()
	}
	adminCtrl.UserCtrl=&web.AdminCtrl{Service: inits.CreateServiceInstance.AdminService}
	adminCtrl.UserCtrl.BaseUserCtrl=*adminCtrl.BaseUserCtl.UserCtrl


	userCtrl.BaseService= func() service.BaseUserService {
		return  inits.CreateServiceInstance.UserService()
		//return userCtrl.Service()
	}
	adminCtrl.BaseService= func() service.BaseUserService {
		return  inits.CreateServiceInstance.AdminService()
		//return adminCtrl.Service()
	}

	roleCtrl.Service= inits.CreateServiceInstance.RoleService

	smsCtrl.Service= inits.CreateServiceInstance.SmsService

	emailCtrl.Service= inits.CreateServiceInstance.EmailService

	payCtrl.Service= inits.CreateServiceInstance.PaySecrtService

	rpcCtl.Service= inits.CreateServiceInstance.RpcService

	config1Ctl.Service= inits.CreateServiceInstance.ConfigService
}

func init()  {

}
type IUserCtrl interface {
	Login()
	LoginByPhone()
	LoginByEmail()
	LoginByUserName()

	Register()
	RegisterByPhone()
	RegisterByEmail()
	RegisterByUserName()

	UpdatePhone()

	UpdateEmail()
	UpdateEmailByPhone()

	UpdatePwdByPhone()
	UpdatePwdByEmail()

	//List()
	ListLogs()

	/*Exists()
	ExistsByPhone()
	ExistsByEmail()
	ExistsByUserName()*/


}

type ICrudCtrl interface {
	Add()
	Update()
	Delete()
	DeleteBatch()
	List()
}
