package web

import (
	dto "github.com/adminwjp/infrastructure-go/dtos"
	util "github.com/adminwjp/infrastructure-go/utils"
	"github.com/adminwjp/infrastructure-go/webs"
	"github.com/adminwjp/users-go/dtos"
	"github.com/adminwjp/users-go/models"
	service "github.com/adminwjp/users-go/services"
	"reflect"
)
func (ctrl *UserCtrl)M(httpWeb webs.HttpWeb){
	res:=dto.ResponseDataDto{Status: true,Code:200,Msg:"success ",Data:&models.UserModel{}}
	httpWeb.Response(200,res)
}
func (ctrl *UserCtrl)BM(httpWeb webs.HttpWeb){
	res:=dto.ResponseDataDto{Status: true,Code:200,Msg:"success ",Data:&models.UserBasicModel{}}
	httpWeb.Response(200,res)
}
type AdminCtrl struct {
	Service func()service.AdminService
	BaseUserCtrl
}
type UserCtrl struct {
	BaseUserCtrl
	Service func()service.UserService
}
type BaseUserCtrl struct {
	Service func()service.BaseUserService

}
/*
	根据条件查询用户信息
*/
func (ctrl *AdminCtrl) List(httpWeb webs.HttpWeb)  {
	var input dtos.GetAdminInput
	_=httpWeb.ShouldBind(&input)

	/*	if err!=nil||reflect.DeepEqual(*&input,dtos.GetUserInput{}){
		res:=dto.ResponseDto{Status: false,Code:400,Msg:"fail"}
		httpWeb.Response(200,res)
		return
	}*/
	var page dto.PageDto
	_=httpWeb.ShouldBindUri(&page)

	service:=ctrl.Service()
	//service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	input.Page=page.Page
	input.Size=page.Size
	users,count,err1:=service.List(&input)
	if err1!=nil{
		res:=dto.ResponseDto{Status: false, Code: 400, Msg: "list fail"}
		httpWeb.Response(200,res)
		return
	}else{
		List(httpWeb,users,page.Page,page.Size,count)
	}
}

/*
	根据条件查询用户日志信息
*/
func (ctrl *AdminCtrl) ListLogs(httpWeb webs.HttpWeb)  {
	var input dtos.GetUserLogInput
	err:=httpWeb.ShouldBind(&input)

	if err!=nil||reflect.DeepEqual(*&input,dtos.GetUserLogInput{}){
		res:=dto.ResponseDto{Status: false,Code:400,Msg:"fail"}
		httpWeb.Response(200,res)
		return
	}
	var page dto.PageDto
	err=httpWeb.ShouldBindUri(&page)

	service:=ctrl.Service()
	//service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	input.Page=page.Page
	input.Size=page.Size
	users,count,err1:=service.ListLogs(&input)
	if err1!=nil{
		res:=dto.ResponseDto{Status: false, Code: 400, Msg: "list fail"}
		httpWeb.Response(200,res)
		return
	}else{
		List(httpWeb,users,page.Page,page.Size,count)
	}
}
/*
	根据条件查询用户信息
*/
func (ctrl *UserCtrl) List(httpWeb webs.HttpWeb)  {
	var input dtos.GetUserInput

	_=httpWeb.ShouldBind(&input)

/*	if err!=nil||reflect.DeepEqual(*&input,dtos.GetUserInput{}){
		res:=dto.ResponseDto{Status: false,Code:400,Msg:"fail"}
		httpWeb.Response(200,res)
		return
	}*/
	var page dto.PageDto
	_=httpWeb.ShouldBindUri(&page)

	service:=ctrl.Service()
	service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	input.Page=page.Page
	input.Size=page.Size
	if &input==nil{
		input=dtos.GetUserInput{}
	}
	users,count,err1:=service.List(&input)
	if err1!=nil{
		res:=dto.ResponseDto{Status: false, Code: 400, Msg: "list fail"}
		httpWeb.Response(200,res)
		return
	}else{
		List(httpWeb,users,page.Page,page.Size,count)
	}
}

/*
	根据条件查询用户日志信息
*/
func (ctrl *UserCtrl) ListLogs(httpWeb webs.HttpWeb)  {
	var input dtos.GetUserLogInput
	err:=httpWeb.ShouldBind(&input)

	if err!=nil||reflect.DeepEqual(*&input,dtos.GetUserLogInput{}){
		res:=dto.ResponseDto{Status: false,Code:400,Msg:"fail"}
		httpWeb.Response(200,res)
		return
	}
	var page dto.PageDto
	err=httpWeb.ShouldBindUri(&page)

	service:=ctrl.Service()
	//service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	input.Page=page.Page
	input.Size=page.Size
	users,count,err1:=service.ListLogs(&input)
	if err1!=nil{
		res:=dto.ResponseDto{Status: false, Code: 400, Msg: "list fail"}
		httpWeb.Response(200,res)
		return
	}else{
		List(httpWeb,users,page.Page,page.Size,count)
	}
}

func (ctrl *UserCtrl) UpdateInfo(httpWeb webs.HttpWeb)  {
	/*if userCtl.OnActionExecution(c){
		return
	}*/
	var input dtos.UpdateUserAuthBasicInput
	err:=httpWeb.ShouldBind(&input)
	if err!=nil||reflect.DeepEqual(*&input,dtos.UserPhoneInput{}){
		res:=dto.ResponseDto{Status: false,Code:400,Msg:"fail"}
		httpWeb.Json(200,res)
		return
	}
	if !ctrl.ValidatorUpdateUserAuthBasicInput(httpWeb,&input){
		return
	}
	ip:=httpWeb.ClientIP()
	*&input.OperatorStringIp=ip
	*&input.OperatorIp=int64(util.NetUtil.StringIpToInt(ip))
	service:=ctrl.Service()
	service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	var res dto.ResponseDto
	if r,_:=service.UpdateAuthBasic(&input);r<1{
		res=dto.ResponseDto{Status: false, Code: 405, Msg: "update basic fail"}

	}else{
		res=dto.ResponseDto{Status: true, Code: 200, Msg: "update basic success"}

	}
	httpWeb.Response(200,res)
}


func (ctrl *BaseUserCtrl) SetRegister(httpWeb webs.HttpWeb,rows int){
	var res dto.ResponseDto
	if rows<1{
		res= dto.ResponseDto{Status: false,Code:   406,Msg:    "register fail"}
	}else{
		res= dto.ResponseDto{Status: true,Code:   200,Msg:    "register success"}
	}
	httpWeb.Response(200,res)
}

func (ctrl *BaseUserCtrl) AccountNotExists(httpWeb webs.HttpWeb)  {
	res:= dto.ResponseDto{Status: false,Code:   404,Msg:    "account not exists"}
	httpWeb.Response(200,res)
}

func (ctrl *BaseUserCtrl) EmailNotExists(httpWeb webs.HttpWeb)  {
	res:= dto.ResponseDto{Status: false,Code:   404,Msg:    "email not exists"}
	httpWeb.Response(200,res)
}

func (ctrl *BaseUserCtrl) PhoneNotExists(httpWeb webs.HttpWeb)  {
	res:= dto.ResponseDto{Status: false,Code:   404,Msg:    "phone not exists"}
	httpWeb.Response(200,res)
}

func (ctrl *BaseUserCtrl) UserNameNotExists(httpWeb webs.HttpWeb)  {
	res:= dto.ResponseDto{Status: false,Code:   404,Msg:    "user_name not exists"}
	httpWeb.Response(200,res)
}

func (ctrl *BaseUserCtrl) AccountExists(httpWeb webs.HttpWeb)  {
	res:= dto.ResponseDto{Status: false,Code:   404,Msg:    "account  exists"}
	httpWeb.Response(200,res)
}

func (ctrl *BaseUserCtrl) EmailExists(httpWeb webs.HttpWeb)  {
	res:= dto.ResponseDto{Status: false,Code:   404,Msg:    "email  exists"}
	httpWeb.Response(200,res)
}

func (ctrl *BaseUserCtrl) PhoneExists(httpWeb webs.HttpWeb)  {
	res:= dto.ResponseDto{Status: false,Code:   404,Msg:    "phone  exists"}
	httpWeb.Response(200,res)
}

func (ctrl *BaseUserCtrl) UserNameExists(httpWeb webs.HttpWeb)  {
	res:= dto.ResponseDto{Status: false,Code:   404,Msg:    "user_name  exists"}
	httpWeb.Response(200,res)
}



func (ctrl *BaseUserCtrl) SetLogin(httpWeb webs.HttpWeb,m *models.UserModel)  {
	if m!=nil{
		ctrl.UserLoginSuccess(httpWeb,m)
	}else{
		ctrl.LoginFail(httpWeb)
	}
}
func (ctrl *BaseUserCtrl) SetAdminLogin(httpWeb webs.HttpWeb,m *models.AdminModel)  {
	if m!=nil{
		ctrl.AdminLoginSuccess(httpWeb,m)
	}else{
		ctrl.LoginFail(httpWeb)
	}
}
func (ctrl *BaseUserCtrl) LoginFail(httpWeb webs.HttpWeb)  {
	res:= dto.ResponseDto{Status: false,Code:   405,Msg:    "login fail"}
	httpWeb.Response(200,res)
}
func (ctrl *BaseUserCtrl) AdminLoginSuccess(httpWeb webs.HttpWeb,m *models.AdminModel)  {
	res:=AdminLogin(*m,httpWeb.ClientIP())
	httpWeb.Response(200,res)
}
func (ctrl *BaseUserCtrl) UserLoginSuccess(httpWeb webs.HttpWeb,m *models.UserModel)  {
	res:= UserLogin(*m,httpWeb.ClientIP())
	httpWeb.Response(200,res)
}