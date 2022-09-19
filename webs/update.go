package web

import (
	dto "github.com/adminwjp/infrastructure-go/dtos"
	util "github.com/adminwjp/infrastructure-go/utils"
	"github.com/adminwjp/infrastructure-go/webs"
	"github.com/adminwjp/users-go/dtos"
	service "github.com/adminwjp/users-go/services"
	"net/http"
	"reflect"
)

/*
	根据手机号修改手机号
*/
func (ctrl *BaseUserCtrl) UpdatePhone(httpWeb webs.HttpWeb)  {
	var input dtos.UpdateUserPhoneInput
	err:=httpWeb.ShouldBind(&input)

	if err!=nil||reflect.DeepEqual(*&input,dtos.UpdateUserPhoneInput{}){
		res:=dto.ResponseDto{Status: false,Code:400,Msg:"fail"}
		httpWeb.Response(200,res)
		return
	}
	if input.Phone=="" ||!util.RegexUtil.IsPhone(input.Phone){
		res:=dto.ResponseDto{Status: false,Code:402,Msg:"phone error"}
		httpWeb.Response(200,res)
		return
	}
	if input.NewPhone=="" ||!util.RegexUtil.IsPhone(input.NewPhone){
		res:=dto.ResponseDto{Status: false,Code:402,Msg:"new_phone error"}
		httpWeb.Response(200,res)
		return
	}
	input.OperatorStringIp=httpWeb.ClientIP()
	input.OperatorIp=int64(util.NetUtil.StringIpToInt(httpWeb.ClientIP()))
	service:=ctrl.Service()
	service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	if r,_:=service.ExistsByPhone(input.Phone);r<1{
		ctrl.PhoneNotExists(httpWeb)
		return
	}else if r,_:=service.ExistsByPhone(input.NewPhone);r>0{
		res1:= dto.ResponseDto{Status: false,Code:   404,Msg:    "new_phone  exists"}
		httpWeb.Response(200,res1)
		return
	}
	if  r,_:=service.UpdatePhone(&input);r<1{
		res1:= dto.ResponseDto{Status: false,Code:   405,Msg:    "update phone fail"}
		httpWeb.Response(200,res1)
	}else{
		res1:= dto.ResponseDto{Status: false,Code:   200,Msg:    "update phone success"}
		httpWeb.Response(200,res1)
	}

}
func (ctrl *BaseUserCtrl) SetEmail(httpWeb webs.HttpWeb,suc bool)  {
	if  !suc{
		res1:= dto.ResponseDto{Status: false,Code:   405,Msg:    "update email fail"}
		httpWeb.Response(http.StatusOK,res1)
	}else{
		res1:= dto.ResponseDto{Status: false,Code:   200,Msg:    "update email success"}
		httpWeb.Response(http.StatusOK,res1)
	}
}
/*
	根据邮箱修改邮箱
*/
func (ctrl *BaseUserCtrl) UpdateEmail(httpWeb webs.HttpWeb)  {
	var input dtos.UpdateUserEmailInput
	err:=httpWeb.ShouldBind(&input)

	if err!=nil||reflect.DeepEqual(*&input,dtos.UpdateUserEmailInput{}){
		res:=dto.ResponseDto{Status: false,Code:400,Msg:"fail"}
		httpWeb.Response(200,res)
		return
	}
	if input.Email=="" ||!util.RegexUtil.IsEmail(input.Email){
		res:=dto.ResponseDto{Status: false,Code:402,Msg:"email error"}
		httpWeb.Response(200,res)
		return
	}
	if input.NewEmail=="" ||!util.RegexUtil.IsEmail(input.NewEmail){
		res:=dto.ResponseDto{Status: false,Code:402,Msg:"new_email error"}
		httpWeb.Response(200,res)
		return
	}
	input.OperatorStringIp=httpWeb.ClientIP()
	input.OperatorIp=int64(util.NetUtil.StringIpToInt(httpWeb.ClientIP()))
	service:=ctrl.Service()
	service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	if r,_:=service.ExistsByEmail(input.Email);r<1{
		ctrl.EmailNotExists(httpWeb)
		return
	}else if r,_:=service.ExistsByEmail(input.NewEmail);r>0{
		res1:= dto.ResponseDto{Status: false,Code:   404,Msg:    "new_email  exists"}
		httpWeb.Response(200,res1)
		return
	}
	r1,_:=service.UpdateEmailByEmail(&input)
	ctrl.SetEmail(httpWeb,r1>0)
}

/*
	根据手机号修改邮箱
*/
func (ctrl *BaseUserCtrl) UpdateEmailByPhone(httpWeb webs.HttpWeb)  {
	var input dtos.UpdateUserEmailByPhoneInput
	err:=httpWeb.ShouldBind(&input)

	if err!=nil||reflect.DeepEqual(*&input,dtos.UpdateUserEmailByPhoneInput{}){
		res:=dto.ResponseDto{Status: false,Code:400,Msg:"fail"}
		httpWeb.Response(200,res)
		return
	}
	if input.Email=="" ||!util.RegexUtil.IsEmail(input.Email){
		res:=dto.ResponseDto{Status: false,Code:402,Msg:"email error"}
		httpWeb.Response(200,res)
		return
	}
	if input.Phone=="" ||!util.RegexUtil.IsPhone(input.Phone){
		res:=dto.ResponseDto{Status: false,Code:402,Msg:"phone error"}
		httpWeb.Response(200,res)
		return
	}
	input.OperatorStringIp=httpWeb.ClientIP()
	input.OperatorIp=int64(util.NetUtil.StringIpToInt(httpWeb.ClientIP()))
	service:=ctrl.Service()
	service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	if r,_:=service.ExistsByEmail(input.Email);r<1{
		ctrl.EmailNotExists(httpWeb)
		return
	}else if r,_:=service.ExistsByPhone(input.Phone);r>0{
		res1:= dto.ResponseDto{Status: false,Code:   404,Msg:    "phone  exists"}
		httpWeb.Response(200,res1)
		return
	}
	r1,_:=service.UpdateEmailByPhone(&input)
	ctrl.SetEmail(httpWeb,r1>0)
}
func (ctrl *BaseUserCtrl) SetPwd(httpWeb webs.HttpWeb,suc bool){
	if  !suc{
		res1:= dto.ResponseDto{Status: false,Code:   405,Msg:    "update pwd fail"}
		httpWeb.Response(200,res1)
	}else{
		res1:= dto.ResponseDto{Status: false,Code:   200,Msg:    "update pwd success"}
		httpWeb.Response(200,res1)
	}
}
/*
	根据邮箱修改密码
*/
func (ctrl *BaseUserCtrl) UpdatePwdByEmail(httpWeb webs.HttpWeb)  {
	var input dtos.UpdateUserPwdByEmailInput
	err:=httpWeb.ShouldBind(&input)

	if err!=nil||reflect.DeepEqual(*&input,dtos.UpdateUserPwdByEmailInput{}){
		res:=dto.ResponseDto{Status: false,Code:400,Msg:"fail"}
		httpWeb.Response(200,res)
		return
	}
	if input.Email=="" ||!util.RegexUtil.IsEmail(input.Email){
		res:=dto.ResponseDto{Status: false,Code:402,Msg:"email error"}
		httpWeb.Response(200,res)
		return
	}
	if input.Pwd=="" {
		res:=dto.ResponseDto{Status: false,Code:402,Msg:"pwd error"}
		httpWeb.Response(200,res)
		return
	}
	input.OperatorStringIp=httpWeb.ClientIP()
	input.OperatorIp=int64(util.NetUtil.StringIpToInt(httpWeb.ClientIP()))
	service:=ctrl.Service()
	service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	if r,_:=service.ExistsByEmail(input.Email);r<1{
		ctrl.EmailNotExists(httpWeb)
		return
	}
	r1,_:=service.UpdatePwdByEmail(&input)
	ctrl.SetPwd(httpWeb,r1>0)
}

/*
	根据手机号修改密码
*/
func (ctrl *BaseUserCtrl) UpdatePwdByPhone(httpWeb webs.HttpWeb)  {
	var input dtos.UpdateUserPwdByPhoneInput
	err:=httpWeb.ShouldBind(&input)

	if err!=nil||reflect.DeepEqual(*&input,dtos.UpdateUserPwdByPhoneInput{}){
		res:=dto.ResponseDto{Status: false,Code:400,Msg:"fail"}
		httpWeb.Response(200,res)
		return
	}
	if input.Phone=="" ||!util.RegexUtil.IsPhone(input.Phone){
		res:= dto.ResponseDto{Status: false,Code:402,Msg:"phone error"}
		httpWeb.Response(200,res)
		return
	}
	if input.Pwd=="" {
		res:= dto.ResponseDto{Status: false,Code:402,Msg:"pwd error"}
		httpWeb.Response(200,res)
		return
	}
	input.OperatorStringIp=httpWeb.ClientIP()
	input.OperatorIp=int64(util.NetUtil.StringIpToInt(httpWeb.ClientIP()))
	service:=ctrl.Service()
	service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	if r,_:=service.ExistsByPhone(input.Phone);r<1{
		ctrl.PhoneNotExists(httpWeb)
		return
	}
	r1,_:=service.UpdatePwdByPhone(&input)
	ctrl.SetPwd(httpWeb,r1>0)
}

/*
	根据旧密码修改密码
*/
func (ctrl *BaseUserCtrl) UpdatePwdByOldPwd(httpWeb webs.HttpWeb) {
	var input dtos.UpdateUserPwdByPwdInput
	err:=httpWeb.ShouldBind(&input)

	if err!=nil||reflect.DeepEqual(*&input,dtos.UpdateUserPwdByPhoneInput{}){
		res:=dto.ResponseDto{Status: false,Code:400,Msg:"fail"}
		httpWeb.Response(200,res)
		return
	}
	if input.Pwd=="" {
		res:=dto.ResponseDto{Status: false,Code:402,Msg:"pwd error"}
		httpWeb.Response(200,res)
		return
	}
	if input.NewPwd=="" {
		res:=dto.ResponseDto{Status: false,Code:402,Msg:"new_pwd error"}
		httpWeb.Response(200,res)
		return
	}
	if input.EnterNewPwd=="" {
		res:=dto.ResponseDto{Status: false,Code:402,Msg:"enter_new_pwd error"}
		httpWeb.Response(200,res)
		return
	}
	if input.EnterNewPwd!=input.NewPwd {
		res:=dto.ResponseDto{Status: false,Code:402,Msg:"new_pwd eq enter_new_pwd not match"}
		httpWeb.Response(200,res)
		return
	}
	input.OperatorStringIp=httpWeb.ClientIP()
	input.OperatorIp=int64(util.NetUtil.StringIpToInt(httpWeb.ClientIP()))
	service1 := ctrl.Service()
	service1.GetTranction().Begin()
	defer func() {
		service1.GetTranction().Commit()
		service1.Clean()
		service1 = nil
	}()
	r,_:=service1.(service.AdminService).UpdatePwdByOldPwd(&input)
	ctrl.SetPwd(httpWeb, r > 0)
}