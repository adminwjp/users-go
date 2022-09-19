package web_gin_controller

import (
	"github.com/adminwjp/users-go/inits"
	service "github.com/adminwjp/users-go/services"
	web "github.com/adminwjp/users-go/webs"
	"github.com/gin-gonic/gin"
	"log"
)

type IMCtl interface{
	M(c *gin.Context)
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
//手动调用 bug 不知道什么时候加载的 坑
func Init()  {
	userCtrl=&UserCtl{}
	adminCtrl=&AdminCtl{}

	//最好每个赋值 不然坑嗲 组合继承
	if inits.CreateServiceInstance.UserService==nil{
		log.Println("Service  UserServiceis nil")
	}
	if inits.CreateServiceInstance.AdminService==nil{
		log.Println("Service  AdminService nil")
	}
	userCtrl.Service= inits.CreateServiceInstance.UserService
	//gin web 之前可以的现在 不行了 毛线什么情况 难道 不能私有？ 需要全局才行 自动回收对象
	userCtrl.BaseUserCtl.UserCtrl=&web.BaseUserCtrl{}
	userCtrl.BaseUserCtl.UserCtrl.Service=func() service.BaseUserService {
		if userCtrl.Service==nil{
			log.Println("gin web  ctrl.Service is nil")
			return  inits.CreateServiceInstance.UserService()
		}
		return userCtrl.Service()
	}
	//userCtrl.UserCtrl=&web.UserCtrl{Service: userCtrl.Service}
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

	roleCtrl.Service=inits.CreateServiceInstance.RoleService

	smsCtrl.Service=inits.CreateServiceInstance.SmsService

	emailCtrl.Service= inits.CreateServiceInstance.EmailService

	payCtrl.Service=inits.CreateServiceInstance.PaySecrtService

	rpcCtl.Service= inits.CreateServiceInstance.RpcService

	config1Ctl.Service= inits.CreateServiceInstance.ConfigService
}

var GinRouterImpl *GinRouter
func init()  {
	GinRouterImpl=&GinRouter{}
}
type IUserCtrl interface {
	Login(c *gin.Context)
	LoginByPhone(c *gin.Context)
	LoginByEmail(c *gin.Context)
	LoginByUserName(c *gin.Context)

	Register(c *gin.Context)
	RegisterByPhone(c *gin.Context)
	RegisterByEmail(c *gin.Context)
	RegisterByUserName(c *gin.Context)

	UpdatePhone(c *gin.Context)

	UpdateEmail(c *gin.Context)
	UpdateEmailByPhone(c *gin.Context)

	UpdatePwdByPhone(c *gin.Context)
	UpdatePwdByEmail(c *gin.Context)

	//List(c *gin.Context)
	ListLogs(c *gin.Context)

	/*Exists(c *gin.Context)
	ExistsByPhone(c *gin.Context)
	ExistsByEmail(c *gin.Context)
	ExistsByUserName(c *gin.Context)*/


}

type ICrudCtrl interface {
	Add(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	DeleteBatch(c *gin.Context)
	List(c *gin.Context)
}