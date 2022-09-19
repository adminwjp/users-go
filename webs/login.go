package web

import (
	dto "github.com/adminwjp/infrastructure-go/dtos"
	util "github.com/adminwjp/infrastructure-go/utils"
	"github.com/adminwjp/infrastructure-go/webs"
	"github.com/adminwjp/users-go/dtos"
	service "github.com/adminwjp/users-go/services"
	"log"
	"reflect"
)

func (ctrl *BaseUserCtrl) Login1(httpWeb webs.HttpWeb,
	method func(*dtos.UserInput,service.BaseUserService),
	login bool){
	log.Println("web login or reg  start.....")
	var input dtos.UserInput
	err:=httpWeb.ShouldBind(&input)

	if err!=nil||reflect.DeepEqual(*&input,dtos.UserInput{}){
		res:=dto.ResponseDto{Status: false,Code:400,Msg:"fail"}
		log.Printf("login or reg bind model fail,err:%s",err.Error())
		httpWeb.Response(200,res)
		return
	}
	//0
	log.Printf("web Register flag,%d",input.Flag)
	log.Println("web login or reg  vali pass .....")
	if !ctrl.ValidatorUserInput(httpWeb,&input){
		return
	}
	if ctrl.Service==nil{
		log.Println("web login or reg  service if null break .....")
		return
	}
	service:=ctrl.Service()
	if service==nil{
		return
	}
	ip:=httpWeb.ClientIP()
	*&input.OperatorStringIp=ip
	*&input.OperatorIp=int64(util.NetUtil.StringIpToInt(ip))
	if ! login{
		service.GetTranction().Begin()
	}
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	if method!=nil{

	}
	log.Println("web login or reg data m invoke start")
	method(&input,service)
	log.Println("web login or reg data m invoke end")
}

func (ctrl *UserCtrl) Login(httpWeb webs.HttpWeb){
	log.Println("web login start")
	var login=func(input *dtos.UserInput,service1 service.BaseUserService){
		log.Println("web login data m start")
		if r,_:=service1.Exists(input.Account,input.Flag);r<1{
			ctrl.AccountNotExists(httpWeb)
			return
		}
		log.Println("web login data start")
		//log show not has
		s,s1:=service1.(service.UserService)
		if s1{
			log.Println("web login service parse  suc")
		}
		userModel,err:=s.Login(input)
		if err!=nil{
			log.Println("login fail,err:%s",err.Error())
		}
		log.Println("web login data end")
		ctrl.SetLogin(httpWeb, userModel)
	}
	ctrl.Login1(httpWeb,login,true)
	log.Println("web login end")

}

func (ctrl *BaseUserCtrl) LoginByEmail1(httpWeb webs.HttpWeb,
	validate func(input *dtos.UserEmailInput)bool,method func(*dtos.UserEmailInput,service.BaseUserService),
	login bool){
	var input dtos.UserEmailInput
	err:=httpWeb.ShouldBind(&input)

	if err!=nil||reflect.DeepEqual(*&input,dtos.UserEmailInput{}){
		res:=dto.ResponseDto{Status: false,Code:400,Msg:"fail"}
		httpWeb.Response(200,res)
		return
	}
	if validate!=nil&&!validate(&input){
		return
	}
	service:=ctrl.Service()
	ip:=httpWeb.ClientIP()
	*&input.OperatorStringIp=ip
	*&input.OperatorIp=int64(util.NetUtil.StringIpToInt(ip))
	if !login{
		service.GetTranction().Begin()
	}
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()

	method(&input,service)
}

func (ctrl *UserCtrl) LoginByEmail(httpWeb webs.HttpWeb,validate func(input *dtos.UserEmailInput)bool){
	var login= func(input *dtos.UserEmailInput,service1 service.BaseUserService) {
		if r,_:=service1.ExistsByEmail(input.Email);r<1{
			ctrl.EmailNotExists(httpWeb)
			return
		}
		userModel,_:=service1.(service.UserService).LoginByEmail(input)
		ctrl.SetLogin(httpWeb, userModel)
	}
	ctrl.LoginByEmail1(httpWeb,validate,login,true)
}

func (ctrl *BaseUserCtrl) LoginByUserName1(httpWeb webs.HttpWeb,
	validate func(input *dtos.UserUserNameInput)bool,method func(*dtos.UserUserNameInput,service.BaseUserService),
	login bool){
	var input dtos.UserUserNameInput
	err:=httpWeb.ShouldBind(&input)

	if err!=nil||reflect.DeepEqual(*&input,dtos.UserUserNameInput{}){
		res:=dto.ResponseDto{Status: false,Code:400,Msg:"fail"}
		httpWeb.Response(200,res)
		return
	}
	if validate!=nil&&!validate(&input){
		return
	}
	service:=ctrl.Service()
	ip:=httpWeb.ClientIP()
	*&input.OperatorStringIp=ip
	*&input.OperatorIp=int64(util.NetUtil.StringIpToInt(ip))
	if !login{
		service.GetTranction().Begin()
	}
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	if method!=nil{

	}
	method(&input,service)
}

func (ctrl *UserCtrl) LoginByUserName(httpWeb webs.HttpWeb,validate func(input *dtos.UserUserNameInput)bool){
	var login= func(input *dtos.UserUserNameInput,service1 service.BaseUserService) {
		if r,_:=service1.ExistsByUserName(input.UserName);r<1{
			ctrl.UserNameNotExists(httpWeb)
			return
		}
		userModel,_:=service1.(service.UserService).LoginByUserName(input)
		ctrl.SetLogin(httpWeb, userModel)
	}
	ctrl.LoginByUserName1(httpWeb,validate,login,true)
}


func (ctrl *BaseUserCtrl) LoginByPhone1(httpWeb webs.HttpWeb,
	validate func(input *dtos.UserPhoneInput)bool,method func(*dtos.UserPhoneInput,service.BaseUserService),
	login bool){
	var input dtos.UserPhoneInput
	err:=httpWeb.ShouldBind(&input)

	if err!=nil||reflect.DeepEqual(*&input,dtos.UserPhoneInput{}){
		res:=dto.ResponseDto{Status: false,Code:400,Msg:"fail"}
		httpWeb.Response(200,res)
		return
	}
	if validate!=nil&&!validate(&input){
		return
	}
	if ctrl==nil{
		log.Println("web login ctrl is nil")
	}
	if ctrl.Service==nil{
		log.Println("web login ctrl.Service is nil")
	}
	service:=ctrl.Service()
	ip:=httpWeb.ClientIP()
	*&input.OperatorStringIp=ip
	*&input.OperatorIp=int64(util.NetUtil.StringIpToInt(ip))
	if !login{
		service.GetTranction().Begin()
	}
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	if method!=nil{

	}
	method(&input,service)
}

func (ctrl *UserCtrl) LoginByPhone(httpWeb webs.HttpWeb,validate func(input *dtos.UserPhoneInput)bool){
	var login= func(input *dtos.UserPhoneInput,service1 service.BaseUserService) {
		if r,_:=service1.ExistsByPhone(input.Phone);r<1{
			ctrl.PhoneNotExists(httpWeb)
			return
		}
		userModel,err:=service1.(service.UserService).LoginByPhone(input)
		if err!=nil{
			log.Printf("phone login fail,err:%s",err.Error())
		}else if userModel==nil{
			log.Printf("phone login fail,err:%s","is nil ")
		}
		ctrl.SetLogin(httpWeb, userModel)
	}
	ctrl.LoginByPhone1(httpWeb,validate,login,true)

}
