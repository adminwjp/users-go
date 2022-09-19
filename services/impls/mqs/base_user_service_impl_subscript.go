package service_mq_impl

import (
	"github.com/adminwjp/users-go/dtos"
	"log"
	"strconv"
)


/**
根据手机号、邮箱、用户名注册
*/
func (service *BaseUserServiceImpl)SubscriptRegister()(int,error){

	//var dao1 daos.BaseUserDao =service.userDao//nil

	//dao1  =service.userDao.(daos.BaseUserDao)//interface conversion: interface is nil, not daos.BaseUserDao
	//if service.user=="admin"{
	//	dao1=service.adminDao//nil
		//dao1  =service.adminDao.(daos.BaseUserDao)
	//}
	var dao1=service.BaseDao
	//count:=0
	s1,e:=service.Mq.SubscriptString(service.reg, func(s string) bool {
		var input =&dtos.UserInput{}
		input.ParseMq(s)

		dao1.GetTranction().Begin()
		r,err:=dao1.Register(input)
		if err!=nil{
			log.Printf("%s %s account reg fail:account %s ,ex err: %s ",service.user,service.dataWay,
				input.Account,err.Error())
			dao1.GetTranction().Rollback()
			return false
		}
		dao1.GetTranction().Commit()
		return  r>0
	},nil)
	if s1{
		return 1, e
	}
	return 0, e
}


/**
根据手机号注册
*/
func (service *BaseUserServiceImpl)SubscriptRegisterByPhone()(int,error){
	var dao1=service.BaseDao
	//count:=0
	s1,e:=service.Mq.SubscriptString(service.regPhone, func(s string) bool {
		var input =&dtos.UserPhoneInput{}
		input.ParseMq(s)

		dao1.GetTranction().Begin()
		r,err:=dao1.RegisterByPhone(input)
		if err!=nil{
			log.Printf("%s %s phone reg fail:account %s ,ex err: %s ",service.user,service.dataWay,
				input.Phone,err.Error())
			dao1.GetTranction().Rollback()
			return false
		}
		dao1.GetTranction().Commit()
		return  r>0
	},nil)
	if s1{
		return 1, e
	}
	return 0, e
}

/**
根据邮箱注册
*/
func (service *BaseUserServiceImpl)SubscriptRegisterByEmail()(int,error){
	var dao1=service.BaseDao
	//count:=0
	s1,e:=service.Mq.SubscriptString(service.regEmail, func(s string) bool {
		var input =&dtos.UserEmailInput{}
		input.ParseMq(s)

		dao1.GetTranction().Begin()
		r,err:=dao1.RegisterByEmail(input)
		if err!=nil{
			log.Printf("%s %s email reg fail:account %s ,ex err: %s ",service.user,service.dataWay,
				input.Email,err.Error())
			dao1.GetTranction().Rollback()
			return false
		}
		dao1.GetTranction().Commit()
		return  r>0
	},nil)
	if s1{
		return 1, e
	}
	return 0, e
}

/**
根据用户名注册
*/
func (service *BaseUserServiceImpl)SubscriptRegisterByUserName()(int,error){
	var dao1=service.BaseDao
	//count:=0
	s1,e:=service.Mq.SubscriptString(service.regUserName, func(s string) bool {
		var input =&dtos.UserUserNameInput{}
		input.ParseMq(s)

		dao1.GetTranction().Begin()
		r,err:=dao1.RegisterByUserName(input)
		if err!=nil{
			log.Printf("%s %s user_name reg fail:account %s ,ex err: %s ",service.user,service.dataWay,
				input.UserName,err.Error())
			dao1.GetTranction().Rollback()
			return false
		}
		dao1.GetTranction().Commit()
		return  r>0
	},nil)
	if s1{
		return 1, e
	}
	return 0, e
}


/*
	根据手机号修改手机号
*/
func (service *BaseUserServiceImpl)SubscriptUpdatePhone()(int,error) {
	var dao1=service.BaseDao
	//count:=0
	s1,e:=service.Mq.SubscriptString(service.updatePhone, func(s string) bool {
		var input =&dtos.UpdateUserPhoneInput{}
		input.ParseMq(s)

		dao1.GetTranction().Begin()
		r,err:=dao1.UpdatePhone(input.Phone,input.NewPhone)
		if err!=nil{
			log.Printf("%s %s update phone fail:account %s ,ex err: %s ",service.user,service.dataWay,
				input.Phone,err.Error())
			dao1.GetTranction().Rollback()
			return false
		}
		dao1.GetTranction().Commit()
		return  r>0
	},nil)
	if s1{
		return 1, e
	}
	return 0, e
}


/*
	根据手机号修改邮箱
*/
func (service *BaseUserServiceImpl)SubscriptUpdateEmailByPhone()(int,error){
	var dao1=service.BaseDao
	//count:=0
	s1,e:=service.Mq.SubscriptString(service.updateEmilByPhone, func(s string) bool {
		var input =&dtos.UpdateUserEmailByPhoneInput{}
		input.ParseMq(s)

		dao1.GetTranction().Begin()
		r,err:=dao1.UpdateEmailByPhone(input.Phone,input.Email)
		if err!=nil{
			log.Printf("%s %s update email by phone fail:account %s ,ex err: %s ",service.user,service.dataWay,
				input.Phone,err.Error())
			dao1.GetTranction().Rollback()
			return false
		}
		dao1.GetTranction().Commit()
		return  r>0
	},nil)
	if s1{
		return 1, e
	}
	return 0, e
}

/*
	根据邮箱修改邮箱
*/
func (service *BaseUserServiceImpl)SubscriptUpdateEmailByEmail()(int,error){
	var dao1=service.BaseDao
	//count:=0
	s1,e:=service.Mq.SubscriptString(service.updateEmil, func(s string) bool {
		var input =&dtos.UpdateUserEmailInput{}
		input.ParseMq(s)

		dao1.GetTranction().Begin()
		r,err:=dao1.UpdateEmail(input.Email,input.NewEmail)
		if err!=nil{
			log.Printf("%s %s update email fail:account %s ,ex err: %s ",service.user,service.dataWay,
				input.Email,err.Error())
			dao1.GetTranction().Rollback()
			return false
		}
		dao1.GetTranction().Commit()
		return  r>0
	},nil)
	if s1{
		return 1, e
	}
	return 0, e
}



/*
	根据手机号修改密码
*/
func (service *BaseUserServiceImpl)SubscriptUpdatePwdByPhone()(int,error){
	var dao1=service.BaseDao
	//count:=0
	s1,e:=service.Mq.SubscriptString(service.updatePwdByPhone, func(s string) bool {
		var input =&dtos.UpdateUserPwdByPhoneInput{}
		input.ParseMq(s)

		dao1.GetTranction().Begin()
		r,err:=dao1.UpdatePwdByPhone(input.Phone,input.Pwd)
		if err!=nil{
			log.Printf("%s %s update pwwd by phone fail:account %s ,ex err: %s ",service.user,service.dataWay,
				input.Phone,err.Error())
			dao1.GetTranction().Rollback()
			return false
		}
		dao1.GetTranction().Commit()
		return  r>0
	},nil)
	if s1{
		return 1, e
	}
	return 0, e
}

/*
	根据邮箱修改密码
*/
func (service *BaseUserServiceImpl)SubscriptUpdatePwdByEmail()(int,error){
	var dao1=service.BaseDao
	//count:=0
	s1,e:=service.Mq.SubscriptString(service.updatePwdByEmil, func(s string) bool {
		var input =&dtos.UpdateUserPwdByEmailInput{}
		input.ParseMq(s)

		dao1.GetTranction().Begin()
		r,err:=dao1.UpdatePwdByPhone(input.Email,input.Pwd)
		if err!=nil{
			log.Printf("%s %s update pwd by email fail:account %s ,ex err: %s ",service.user,service.dataWay,
				input.Email,err.Error())
			dao1.GetTranction().Rollback()
			return false
		}
		dao1.GetTranction().Commit()
		return  r>0
	},nil)
	if s1{
		return 1, e
	}
	return 0, e
}
/*
	修改登录次数
*/
func(service *BaseUserServiceImpl) SubscriptUpdateLoginFailCount()(int,error){
	var dao1=service.BaseDao
	//count:=0
	s1,e:=service.Mq.SubscriptString(service.updateLoginFailCount, func(s string) bool {
		userId,_:=strconv.ParseInt(s,10,64)

		dao1.GetTranction().Begin()
		r,err:=dao1.UpdateLoginFailCount(userId)
		if err!=nil{
			log.Printf("%s %s update login fail count fail:userId %d ,ex err: %s ",service.user,service.dataWay,
				userId,err.Error())
			dao1.GetTranction().Rollback()
			return false
		}
		dao1.GetTranction().Commit()
		return  r>0
	},nil)
	if s1{
		return 1, e
	}
	return 0, e
}


func(service *BaseUserServiceImpl) SubscriptResetLoginFailCount()(int,error){
	var dao1=service.BaseDao
	//count:=0
	s1,e:=service.Mq.SubscriptString(service.resetLoginFailCount, func(s string) bool {
		userId,_:=strconv.ParseInt(s,10,64)

		dao1.GetTranction().Begin()
		r,err:=dao1.ResetLoginFailCount(userId)
		if err!=nil{
			log.Printf("%s %s reset login fail count fail:userId %d ,ex err: %s ",service.user,service.dataWay,
				userId,err.Error())
			dao1.GetTranction().Rollback()
			return false
		}
		dao1.GetTranction().Commit()
		return  r>0
	},nil)
	if s1{
		return 1, e
	}
	return 0, e
}