package service_mq_impl

import (
	mq "github.com/adminwjp/infrastructure-go/mqs"
	"github.com/adminwjp/users-go/daos"
	"github.com/adminwjp/users-go/dtos"
	service_impl "github.com/adminwjp/users-go/services/impls"
	"log"
	"strconv"
	"time"
)

//用户服务基础接口
type BaseUserServiceImpl struct {
	service_impl.BaseUserServiceImpl
	//Baseervice *service_impl.BaseUserServiceImpl

	/*BaseDao	daos.BaseUserDao
	BaseCache caches.BaseUserCache
	//锁
	Lock locks.Lock

	//重试
	Retry retries.Retry*/

	userDao daos.UserDao
	adminDao daos.AdminDao
	Mq mq.Mq

	//小写挎包 不继承
	//数据实现方式
	//gorm gorm jinzhus bee orm xorm  mong es
	dataWay string

	//缓存实现方式
	//empty local_memory file_memory redis
	cache1 string

	//锁实现方式
	//empty local_lock  redis consul zookeeper
	lock1 string

	//admin user
	user string

	//empty kafka rabbitmq
	mq1 string

	//mq topic
	reg string
	regPhone string
	regEmail string
	regUserName string
	updateEmil string
	updateEmilByPhone string
	updatePhone string
	updatePwd string
	updatePwdByEmil string
	updatePwdByPhone string
	updateLoginFailCount string
	resetLoginFailCount string
	batchCount int
}
func (service *BaseUserServiceImpl) Init(admin bool){
	service.batchCount=10
	service.user="admin"
	if !admin{
		service.user="user"
	}
	service.reg=service.user+"reg"
	service.regPhone=service.user+"regPhone"
	service.regEmail=service.user+"regEmail"
	service.regUserName=service.user+"regUserName"
	service.updateEmil=service.user+"updateEmil"
	service.updateEmilByPhone=service.user+"updateEmilByPhone"
	service.updatePhone=service.user+"updatePhone"
	service.updatePwd=service.user+"updatePwd"
	service.updatePwdByEmil=service.user+"updatePwdByEmil"
	service.updatePwdByPhone=service.user+"updatePwdByPhone"
	service.updateLoginFailCount=service.user+"updateLoginFailCount"
	service.resetLoginFailCount=service.user+"resetLoginFailCount"
}
func (service *BaseUserServiceImpl)GetTranction()daos.TranDao{
	return daos.EmptyTranInstance
}
/**
根据手机号、邮箱、用户名注册
*/
func (service *BaseUserServiceImpl)Register(input *dtos.UserInput)(int,error){
	if input.OperatorDate<1{
		input.OperatorDate=time.Now().Unix()
	}
	l:=service.OnLock(input.Account)
	if l{
		defer func() {
			service.OnUnLock(input.Account)
		}()
	}else{
		//log.Println("mq  register lock fail")
		return 0, nil
	}
	//mq
	_,err:=service.Mq.PublishString(service.reg,input.ToMq())
	if err!=nil{
		log.Printf("%s %s account reg fail:account %s ,ex err: %s ",service.user,service.mq1,
			input.Account,err.Error())
		return 0, err
	}else{
		log.Printf("%s %s account reg suc:account %s ",service.user,service.mq1,
			input.Account)
	}
	//cache
	_,err= service.RegisterCache(input)
	if err!=nil{
		log.Printf("%s %s account reg fail:account %s ,ex err: %s ",service.user,service.cache1,
			input.Account,err.Error())
	}else{
		log.Printf("%s %s account reg suc:account %s ",service.user,service.cache1,
			input.Account)
	}
	return 1, err
}


/**
根据手机号注册
*/
func (service *BaseUserServiceImpl)RegisterByPhone(input *dtos.UserPhoneInput)(int,error){
	if input.OperatorDate<1{
		input.OperatorDate=time.Now().Unix()
	}
	l:=service.OnLock(input.Phone)
	if l{
		defer func() {
			service.OnUnLock(input.Phone)
		}()
	}else{

		return 0, nil
	}
	//mq
	_,err:=service.Mq.PublishString(service.regPhone,input.ToMq())
	if err!=nil{
		log.Printf("%s %s phone reg fail:phone %s ,ex err: %s ",service.user,service.mq1,
			input.Phone,err.Error())
		return 0, err
	}else{
		log.Printf("%s %s phone reg suc:phone %s ",service.user,service.mq1,
			input.Phone)
	}
	_,err=  service.RegisterPhoneCache(input)
	if err!=nil{
		log.Printf("%s %s phone reg fail:phone %s ,ex err: %s ",
			service.user,service.cache1,input.Phone,err.Error())
	}else{
		log.Printf("%s %s phone reg suc:phone %s ",service.user,service.cache1,
			input.Phone)
	}

	return 1, err
}

/**
根据邮箱注册
*/
func (service *BaseUserServiceImpl)RegisterByEmail(input *dtos.UserEmailInput)(int,error){
	if input.OperatorDate<1{
		input.OperatorDate=time.Now().Unix()
	}
	l:=service.OnLock(input.Email)
	if l{
		defer func() {
			service.OnUnLock(input.Email)
		}()
	}else{
		return 0, nil
	}
	//mq
	_,err:=service.Mq.PublishString(service.regEmail,input.ToMq())
	if err!=nil{
		log.Printf("%s %s email reg fail:email %s ,ex err: %s ",service.user,service.mq1,
			input.Email,err.Error())
		return 0, err
	}else{
		log.Printf("%s %s  email suc:email %s ",service.user,service.mq1,
			input.Email)
	}
	_,err=   service.RegisterByEmail(input)
	if err!=nil{
		log.Printf("%s %s email reg fail:email %s ,ex err: %s ",
			service.user,service.cache1,input.Email,err.Error())
	}else{
		log.Printf("%s %s email reg suc:email %s ",service.user,service.cache1,
			input.Email)
	}
	return 1, err
}

/**
根据用户名注册
*/
func (service *BaseUserServiceImpl)RegisterByUserName(input *dtos.UserUserNameInput)(int,error){
	if input.OperatorDate<1{
		input.OperatorDate=time.Now().Unix()
	}
	l:=service.OnLock(input.UserName)
	if l{
		defer func() {
			service.OnUnLock(input.UserName)
		}()
	}else{
		return 0, nil
	}
	//mq
	_,err:=service.Mq.PublishString(service.regUserName,input.ToMq())
	if err!=nil{
		log.Printf("%s %s user_name reg fail:user_name %s ,ex err: %s ",service.user,service.mq1,
			input.UserName,err.Error())
		return 0, err
	}else{
		log.Printf("%s %s user_name  suc:user_name %s ",service.user,service.mq1,
			input.UserName)
	}
	_,err=   service.RegisterUserNameCache(input)
	if err!=nil{
		log.Printf("%s %s user_name reg fail:user_name %s ,ex err: %s ",
			service.user,service.cache1,input.UserName,err.Error())
	}else{
		log.Printf("%s %s user_name reg suc:user_name %s ",service.user,service.cache1,
			input.UserName)
	}
	return 1, err
}


/*
	根据手机号修改手机号
*/
func (service *BaseUserServiceImpl)UpdatePhone(input *dtos.UpdateUserPhoneInput)(int,error) {
	l:=service.OnLock(input.Phone)
	if l{
		defer func() {
			service.OnUnLock(input.Phone)
		}()
	}else{
		return 0, nil
	}
	//mq
	_,err:=service.Mq.PublishString(service.updatePhone,input.ToMq())
	if err!=nil{
		log.Printf("%s %s update phone fail:phone %s ,ex err: %s ",service.user,service.mq1,
			input.Phone,err.Error())
		return 0, err
	}else{
		log.Printf("%s %s update phone  suc:phone %s ",service.user,service.mq1,
			input.Phone)
	}
	_, err = service.BaseCache.UpdatePhone(input)
	if err!=nil{
		log.Printf("%s %s update phone  fail:phone %s ,ex err: %s ",
			service.user,service.cache1,input.Phone,err.Error())
	}else{
		log.Printf("%s %s update phone  suc:phone %s  ",
			service.user,service.cache1,input.Phone)
		service.BaseCache.UpdatePhone(input)
	}
	return 1, err
}


/*
	根据手机号修改邮箱
*/
func (service *BaseUserServiceImpl)UpdateEmailByPhone(input *dtos.UpdateUserEmailByPhoneInput)(int,error){
	l:=service.OnLock(input.Phone)
	if l{
		defer func() {
			service.OnUnLock(input.Phone)
		}()
	}else{
		return 0, nil
	}
	//mq
	_,err:=service.Mq.PublishString(service.updateEmilByPhone,input.ToMq())
	if err!=nil{
		log.Printf("%s %s update email fail:phone %s ,ex err: %s ",service.user,service.mq1,
			input.Phone,err.Error())
		return 0, err
	}else{
		log.Printf("%s %s update email  suc:phone %s ",service.user,service.mq1,
			input.Phone)
	}
	r, err :=  service.BaseCache.UpdateEmailByPhone(input)
	if err!=nil{
		log.Printf("%s %s update email  fail:phone %s ,ex err: %s ",
			service.user,service.cache1,input.Phone,err.Error())
	}else{
		log.Printf("%s %s update email  suc:phone %s  ",
			service.user,service.cache1,input.Phone)
	}
	return r, err
}

/*
	根据邮箱修改邮箱
*/
func (service *BaseUserServiceImpl)UpdateEmailByEmail(input *dtos.UpdateUserEmailInput)(int,error){
	l:=service.OnLock(input.Email)
	if l{
		defer func() {
			service.OnUnLock(input.Email)
		}()
	}else{
		return 0, nil
	}
	//mq
	_,err:=service.Mq.PublishString(service.updateEmil,input.ToMq())
	if err!=nil{
		log.Printf("%s %s update email fail:email %s ,ex err: %s ",service.user,service.mq1,
			input.Email,err.Error())
		return 0, err
	}else{
		log.Printf("%s %s update email  suc:email %s ",service.user,service.mq1,
			input.Email)
	}
	_, err = service.BaseCache.UpdateEmailByEmail(input)
	if err!=nil{
		log.Printf("%s %s update email  fail:email %s ,ex err: %s ",
			service.user,service.cache1,input.Email,err.Error())
	}else{
		log.Printf("%s %s update email  suc:email %s  ",
			service.user,service.cache1,input.Email)
	}
	return 1, err
}



/*
	根据手机号修改密码
*/
func (service *BaseUserServiceImpl)UpdatePwdByPhone(input *dtos.UpdateUserPwdByPhoneInput)(int,error){
	//mq
	_,err:=service.Mq.PublishString(service.updatePwdByPhone,input.ToMq())
	if err!=nil{
		log.Printf("%s %s update pwd fail:phone %s ,ex err: %s ",service.user,service.mq1,
			input.Phone,err.Error())
		return 0, err
	}else{
		log.Printf("%s %s update pwd  suc:phone %s ",service.user,service.mq1,
			input.Phone)
	}
	_, err = service.BaseCache.UpdatePwdByPhone(input)
	if err!=nil{
		log.Printf("%s %s update pwd  fail:phone %s ,ex err: %s ",
			service.user,service.cache1,input.Phone,err.Error())
	}else{
		log.Printf("%s %s update pwd  suc:phone %s  ",
			service.user,service.cache1,input.Phone)
	}
	return 1, err
}

/*
	根据邮箱修改密码
*/
func (service *BaseUserServiceImpl)UpdatePwdByEmail(input *dtos.UpdateUserPwdByEmailInput)(int,error){
	//mq
	_,err:=service.Mq.PublishString(service.updatePwdByEmil,input.ToMq())
	if err!=nil{
		log.Printf("%s %s update pwd fail:email %s ,ex err: %s ",service.user,service.mq1,
			input.Email,err.Error())
		return 0, err
	}else{
		log.Printf("%s %s update pwd  suc:email %s ",service.user,service.mq1,
			input.Email)
	}
	_, err = service.BaseCache.UpdatePwdByEmail(input)
	if err!=nil{
		log.Printf("%s %s update pwd  fail:email %s ,ex err: %s ",
			service.user,service.cache1,input.Email,err.Error())
	}else{
		log.Printf("%s %s update pwd  suc:email %s  ",
			service.user,service.cache1,input.Email)
	}
	return 1, err
}
/*
	修改登录次数
*/
func(service *BaseUserServiceImpl) UpdateLoginFailCount(id int64)(int,error){
	//mq
	_,err:=service.Mq.PublishString(service.updateLoginFailCount,strconv.FormatInt(id,10))
	if err!=nil{
		log.Printf("%s %s update login fail count fail:id %d ,ex err: %s ",service.user,service.mq1,
			id,err.Error())
		return 0, err
	}else{
		log.Printf("%s %s update login fail count  suc:id %d ",service.user,service.mq1,
			id)
	}
	_, err = service.BaseCache.UpdateLoginFailCount(id)
	if err!=nil{
		log.Printf("%s %s update login fail count  fail:id %d ,ex err: %s ",
			service.user,service.cache1,id,err.Error())
	}else{
		log.Printf("%s %s update login fail count  suc:id %d  ",
			service.user,service.cache1,id)
	}
	return 1, err
}


func(service *BaseUserServiceImpl) ResetLoginFailCount(id int64)(int,error){
	//mq
	_,err:=service.Mq.PublishString(service.resetLoginFailCount,strconv.FormatInt(id,10))
	if err!=nil{
		log.Printf("%s %s  restet login fail count  fail:id %d ,ex err: %s ",
			service.user,service.mq1,id,err.Error())
		return 0, err
	}else {
		log.Printf("%s %s restet login fail count  suc:id %d  ",
			service.user, service.mq1, id)
	}
	_, err = service.BaseCache.ResetLoginFailCount(id)
	if err!=nil{
		log.Printf("%s %s  restet login fail count  fail:id %d ,ex err: %s ",
			service.user,service.cache1,id,err.Error())
	}else{
		log.Printf("%s %s restet login fail count  suc:id %d  ",
			service.user,service.cache1,id)

	}
	return 1,err
}