package service_impl

import (
	data "github.com/adminwjp/infrastructure-go/datas"
	"github.com/adminwjp/users-go/daos"
	"github.com/adminwjp/users-go/datas"
	"github.com/adminwjp/users-go/services"
)
//win 支持各种 依赖 win build pass linux build fail
//gorm.io/ gorm or gorm jinzhus 2 选1 报错  默认加载的 gorm.io 但我使用的是 gorm jinzhus
// go src remove  gorm jinzhus
//不是这个错误 test pass 其它 报错
//绕来绕去不好排查啊 单独方法体 可以
//方法 忘调用了
var CacheConfig= &CacheConfigImpl{}
var RegisterConfig= &RegisterConfigImpl{}
var LockConfig= &LockConfigImpl{}
var RetryConfig= &RetryConfigImpl{}
var DaoServiceInstance=&DaoServiceImpl{Cruds: make(map[string]daos.CrudDao,10)}
type ServiceImpl struct {
	Dao daos.CrudDao
	TranManager daos.TranDao
}
var AdminServiceInstance  *AdminServiceImpl
var as=false
func GetAdminService()  service.AdminService{
	if datas.GlobalConfig.EnableMq{
		//return 	service_mq_impl.AdminMqInstance
	}
	service1:=AdminServiceInstance
	switch datas.GlobalConfig.DataFlag {
	case data.DataEs,data.DataMong:
		if !as{
			service1=&AdminServiceImpl{}
			AdminServiceInstance=service1
			as=true
		}
		break
	default:
		service1=&AdminServiceImpl{}
		break
	}

	UpdateAdmin(service1)
	return  service1
}
func  UpdateAdmin(service1 *AdminServiceImpl)  {
	CacheConfig.CacheByAdmin(service1)
	RetryConfig.RetryByAdmin(service1)
	LockConfig.LockByAdmin(service1)
	DaoServiceInstance.UpdateDataAdmin(service1)
}
var UserServiceInstance  *UserServiceImpl
var us=false
func GetUserService()  service.UserService{
	if datas.GlobalConfig.EnableMq{
		//return 	service_mq_impl.UserMqInstance
	}
	service1:=UserServiceInstance
	switch datas.GlobalConfig.DataFlag {
	case data.DataEs,data.DataMong:
		if !us{
			service1=&UserServiceImpl{}
			UserServiceInstance=service1
			us=true
		}
		break
	default:
		service1=&UserServiceImpl{}
		break
	}
	UpdateUser(service1)
	return  service1
}
func UpdateUser(service1 *UserServiceImpl)  {
	CacheConfig.CacheByUser(service1)
	RetryConfig.RetryByUser(service1)
	LockConfig.LockByUser(service1)
	DaoServiceInstance.UpdateDataUser(service1)
}
func GetAdminBaseUserService()service.BaseUserService  {
	return  GetAdminService()
}
func GetUserBaseUserService() service.BaseUserService {
	return  GetUserService()
}

func GetSmsService() service.SmsService {
	return DaoServiceInstance.GetSms()
}
func GetEmailService () service.EmailService {
	return DaoServiceInstance.GetEmail()
}
func GetRoleService() service.RoleService {
	return DaoServiceInstance.GetRole()
}
func GetPaySecrtService() service.PaySecrtService {
	return DaoServiceInstance.GetPay()
}
func GetRpcService() service.RpcService {
	return DaoServiceInstance.GetRpc()
}
func GetConfigService() service.ConfigService {
	return DaoServiceInstance.GetConfig()
}

//事务手动提交
//多线程 不能这样搞  不可控 静态 只能 new aop 只能反射 调用 放弃
//有写可以单列 有些 不能必须重新定义


//========================== start pay
func (serivce1 *PayServiceImpl)EmptyCache()  {


}
func (service1 *PayServiceImpl)RedisCache()  {

}



//========================== end  pay

//========================== start role
func (serivce1 *RoleServiceImpl)EmptyCache()  {


}
func (service1 *RoleServiceImpl)RedisCache()  {

}



//========================== end role


//========================== start sms

func (serivce1 *SmsServiceImpl)EmptyCache()  {


}
func (service1 *SmsServiceImpl)RedisCache()  {

}


func (serivce1 *SmsServiceImpl)UpdateMong()  {


}


//========================== end sms


