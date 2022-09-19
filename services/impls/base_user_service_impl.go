package service_impl

import (
	dto "github.com/adminwjp/infrastructure-go/dtos"
	 "github.com/adminwjp/infrastructure-go/locks"
	 "github.com/adminwjp/infrastructure-go/retries"
	 "github.com/adminwjp/users-go/caches"
	"github.com/adminwjp/users-go/daos"
	"github.com/adminwjp/users-go/dtos"
	"log"
	"time"
)

//用户服务基础接口
type BaseUserServiceImpl struct {
	BaseDao	daos.BaseUserDao
	BaseCache caches.BaseUserCache
	//锁
	Lock locks.Lock

	//重试
	Retry retries.Retry

	//日志写入
	Log func(flag string,data dtos.OperatorLog)(int,error)

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
}
func (service *BaseUserServiceImpl)GetTranction()daos.TranDao{
	return service.BaseDao.GetTranction()
}
func (service *BaseUserServiceImpl) OnLock(account string)bool{
	l:=service.Retry.OnRetry(func(retryNum int,retryTime time.Duration) bool {
		l1,err:=service.Lock.LockDate(service.user,time.Millisecond*30)
		if err!=nil{
			log.Printf("%s %s  reg or update %s lock retry:account==> %s ,ex err: %s ",
				service.user,service.dataWay,service.lock1,
				account,err.Error())
			time.Sleep(retryTime)
		}else{
			log.Printf("%s %s  reg or update %s lock retry:account==> %s suc ",
				service.user,service.dataWay,service.lock1,
				account)
		}
		return  l1
	},time.Millisecond*10,3)
	return  l
}
func (service *BaseUserServiceImpl) OnUnLock(account string)bool{
	l:=service.Retry.OnRetry(func(retryNum int,retryTime time.Duration) bool {
		l1,err:=service.Lock.UnLock(service.user)
		if err!=nil{
			log.Printf("%s %s  reg or update   %s un lock retry:account==> %s ,ex err: %s ",
				service.user,service.dataWay,service.lock1,
				account,err.Error())
			time.Sleep(retryTime)
		}else{
			log.Printf("%s %s  reg or update   %s un lock retry:account==> %s  ",
				service.user,service.dataWay,service.lock1,
				account)
		}
		return  l1
	},time.Millisecond*10,3)
	return  l
}
/**
根据手机号、邮箱、用户名注册
*/
func (service *BaseUserServiceImpl)Register(input *dtos.UserInput)(int,error){
	log.Println("reg")
	if input.OperatorDate<1{
		input.OperatorDate=time.Now().Unix()
	}
	log.Println(input.OperatorDate)
	l:=service.OnLock(input.Account)
	if l{
		defer func() {
			service.OnUnLock(input.Account)
		}()
	}else{
		log.Println("reg lock fail")
		return 0, nil
	}
	input.Id=Worker.GetId()
	log.Println("reg data")
	r,err:= service.BaseDao.Register(input)
	if err!=nil{
		log.Printf("%s %s account reg fail:account %s ,ex err: %s ",service.user,service.dataWay,
			input.Account,err.Error())
	}else{
		log.Printf("%s %s account reg suc:account %s ",service.user,service.dataWay,
			input.Account)
		service.RegisterCache(input)
	}
	log.Println("reg data end ")
	return r, err
}
func (service *BaseUserServiceImpl)RegisterPhoneCache(input *dtos.UserPhoneInput)(int,error)  {
	r,err1:=service.BaseCache.RegisterByPhone(input.Phone)
	if err1!=nil{
		log.Printf("%s %s phone cache reg fail:phone %s ,ex err: %s ",service.user,service.cache1,
			input.Phone,err1.Error())
	}else{
		log.Printf("%s %s phone cache reg suc:phone %s ",service.user,service.cache1,
			input.Phone)
	}
	return r, err1
}
func (service *BaseUserServiceImpl)RegisterEmailCache(input *dtos.UserEmailInput)(int,error)  {
	r,err1:=service.BaseCache.RegisterByEmail(input.Email)
	if err1!=nil{
		log.Printf("%s %s email cache reg fail:email %s ,ex err: %s ",service.user,service.cache1,
			input.Email,err1.Error())
	}else{
		log.Printf("%s %s email cache reg suc:email %s ",service.user,service.cache1,
			input.Email)
	}
	return r, err1
}
func (service *BaseUserServiceImpl)RegisterUserNameCache(input *dtos.UserUserNameInput)(int,error)  {
	r,err1:=service.BaseCache.RegisterByUserName(input.UserName)
	if err1!=nil{
		log.Printf("%s %s user_name cache reg fail:user_name %s ,ex err: %s ",service.user,service.cache1,
			input.UserName,err1.Error())
	}else{
		log.Printf("%s %s user_name cache reg suc:user_name %s ",service.user,service.cache1,
			input.UserName)
	}
	return r, err1
}
func (service *BaseUserServiceImpl)RegisterCache(input *dtos.UserInput)(int,error)  {
	r,err1:=service.BaseCache.Register(input.Account,int(input.Flag))
	if err1!=nil{
		log.Printf("%s %s account cache reg fail:account %s ,ex err: %s ",service.user,service.cache1,
			input.Account,err1.Error())
	}else{
		log.Printf("%s %s account cache reg suc:account %s ",service.user,service.cache1,
			input.Account)
	}
	return r, err1
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
	input.Id=Worker.GetId()
	r,err:=  service.BaseDao.RegisterByPhone(input)
	if err!=nil{
		log.Printf("%s %s phone reg fail:phone %s ,ex err: %s ",
			service.user,service.dataWay,input.Phone,err.Error())
	}else{
		log.Printf("%s %s phone reg suc:phone %s ",service.user,service.dataWay,
			input.Phone)
		service.RegisterPhoneCache(input)
	}

	return r, err
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
	input.Id=Worker.GetId()
	r,err:=   service.BaseDao.RegisterByEmail(input)
	if err!=nil{
		log.Printf("%s %s email reg fail:email %s ,ex err: %s ",
			service.user,service.dataWay,input.Email,err.Error())
	}else{
		log.Printf("%s %s email reg suc:email %s ",service.user,service.dataWay,
			input.Email)
		service.RegisterEmailCache(input)
	}
	return r, err
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
	input.Id=Worker.GetId()
	r,err:=   service.BaseDao.RegisterByUserName(input)
	if err!=nil{
		log.Printf("%s %s user_name reg fail:user_name %s ,ex err: %s ",
			service.user,service.dataWay,input.UserName,err.Error())
	}else{
		log.Printf("%s %s user_name reg suc:user_name %s ",service.user,service.dataWay,
			input.UserName)
		service.RegisterUserNameCache(input)
	}
	return r, err
}

/**
根据手机号、邮箱、用户名检测账号是否存在
*/
func (service *BaseUserServiceImpl)Exists(account string,flag dto.AccounType)(int,error){

	if c,err:=service.BaseCache.Exists(account,flag);c>0{
		return 1,err
	}
	r,err:=  service.BaseDao.Exists(account,flag)
	if err!=nil{
		log.Printf("%s %s account check fail:account %s ,ex err: %s ",
			service.user,service.dataWay,account,err.Error())
	}
	return r, err
}


/**
根据手机号检测账号是否存在
*/
func (service *BaseUserServiceImpl)ExistsByPhone(phone string)(int,error){
	if c,err:=service.BaseCache.ExistsByPhone(phone);c>0{
		return 1,err
	}
	r,err:=  service.BaseDao.ExistsByPhone(phone)
	if err!=nil{
		log.Printf("%s %s phone check fail:phone %s ,ex err: %s ",
			service.user,service.dataWay,phone,err.Error())
	}
	return r, err
}

/**
根据邮箱检测账号是否存在
*/
func (service *BaseUserServiceImpl)ExistsByEmail(email string)(int,error){
	if c,err:=service.BaseCache.ExistsByEmail(email);c>0{
		return 1,err
	}
	r,err:=   service.BaseDao.ExistsByEmail(email)
	if err!=nil{
		log.Printf("%s %s fail:email %s ,ex err: %s ",
			service.user,service.dataWay,email,err.Error())
	}
	return r, err
}

/**
根据用户名检测账号是否存在
*/
func (service *BaseUserServiceImpl)ExistsByUserName(userName string)(int,error){
	if c,err:=service.BaseCache.ExistsByUserName(userName);c>0{
		return 1,err
	}
	r,err:=   service.BaseDao.ExistsByUserName(userName)
	if err!=nil{
		log.Printf("%s %s user_name check fail:user_name %s ,ex err: %s ",
			service.user,service.dataWay,userName,err.Error())
	}
	return r, err
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
	r, err := service.BaseDao.UpdatePhone(input.Phone, input.NewPhone)
	if err!=nil{
		log.Printf("%s %s update phone  fail:phone %s ,ex err: %s ",
			service.user,service.dataWay,input.Phone,err.Error())
	}else{
		log.Printf("%s %s update phone  suc:phone %s  ",
			service.user,service.dataWay,input.Phone)
		service.BaseCache.UpdatePhone(input)
	}
	return r, err
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
	r, err :=  service.BaseDao.UpdateEmailByPhone(input.Phone,input.Email)
	if err!=nil{
		log.Printf("%s %s update email  fail:phone %s ,ex err: %s ",
			service.user,service.dataWay,input.Phone,err.Error())
	}else{
		log.Printf("%s %s update email  suc:phone %s  ",
			service.user,service.dataWay,input.Phone)
		service.BaseCache.UpdateEmailByPhone(input)
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
	r, err := service.BaseDao.UpdateEmail(input.Email,input.NewEmail)
	if err!=nil{
		log.Printf("%s %s update email  fail:email %s ,ex err: %s ",
			service.user,service.dataWay,input.Email,err.Error())
	}else{
		log.Printf("%s %s update email  suc:email %s  ",
			service.user,service.dataWay,input.Email)
		service.BaseCache.UpdateEmailByEmail(input)
	}
	return r, err
}



/*
	根据手机号修改密码
*/
func (service *BaseUserServiceImpl)UpdatePwdByPhone(input *dtos.UpdateUserPwdByPhoneInput)(int,error){
	r, err := service.BaseDao.UpdatePwdByPhone(input.Phone,input.Pwd)
	if err!=nil{
		log.Printf("%s %s update pwd  fail:phone %s ,ex err: %s ",
			service.user,service.dataWay,input.Phone,err.Error())
	}else{
		log.Printf("%s %s update pwd  suc:phone %s  ",
			service.user,service.dataWay,input.Phone)
		service.BaseCache.UpdatePwdByPhone(input)
	}
	return r, err
}

/*
	根据邮箱修改密码
*/
func (service *BaseUserServiceImpl)UpdatePwdByEmail(input *dtos.UpdateUserPwdByEmailInput)(int,error){
	r, err := service.BaseDao.UpdatePwdByEmail(input.Email,input.Pwd)
	if err!=nil{
		log.Printf("%s %s update pwd  fail:email %s ,ex err: %s ",
			service.user,service.dataWay,input.Email,err.Error())
	}else{
		log.Printf("%s %s update pwd  suc:email %s  ",
			service.user,service.dataWay,input.Email)
		service.BaseCache.UpdatePwdByEmail(input)
	}
	return r, err
}
/*
	修改登录次数
*/
func(service *BaseUserServiceImpl) UpdateLoginFailCount(id int64)(int,error){
	r, err := service.BaseDao.UpdateLoginFailCount(id)
	if err!=nil{
		log.Printf("%s %s update login fail count  fail:id %d ,ex err: %s ",
			service.user,service.dataWay,id,err.Error())
	}else{
		log.Printf("%s %s update login fail count  suc:id %d  ",
			service.user,service.dataWay,id)
		service.BaseCache.UpdateLoginFailCount(id)
	}
	return r, err
}
func(service *BaseUserServiceImpl) GetLoginFailCount(id int64)(int,error){
	if c,err:=service.BaseCache.GetLoginFailCount(id);err!=nil{
		log.Printf("%s %s get login fail count  suc:id %d  ",
			service.user,service.dataWay,id)
		return c, err
	}
	r, err := service.BaseDao.GetLoginFailCount(id)
	if err!=nil{
		log.Printf("%s %s get login fail count  fail:id %d ,ex err: %s ",
			service.user,service.dataWay,id,err.Error())
	}
	return r,err
}

func(service *BaseUserServiceImpl) ResetLoginFailCount(id int64)(int,error){
	r, err := service.BaseDao.ResetLoginFailCount(id)
	if err!=nil{
		log.Printf("%s %s  restet login fail count  fail:id %d ,ex err: %s ",
			service.user,service.dataWay,id,err.Error())
	}else{
		log.Printf("%s %s restet login fail count  suc:id %d  ",
			service.user,service.dataWay,id)
		service.BaseCache.ResetLoginFailCount(id)
	}
	return r,err
}