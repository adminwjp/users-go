package web_http_controller

import (
	"github.com/adminwjp/infrastructure-go/webs/gins"
	"github.com/adminwjp/infrastructure-go/webs/https"
	"github.com/adminwjp/users-go/dtos"
	"github.com/adminwjp/users-go/services"
	"github.com/adminwjp/users-go/webs"
	"net/http"
)

type BaseUserCtl struct {
	gins.GinFilter
	//BaseUserService func()service.BaseUserService
	BaseService func()service.BaseUserService
	UserCtrl *web.BaseUserCtrl
}






/**
根据手机号、邮箱、用户名注册
*/
func (ctrl *BaseUserCtl) Register(writer http.ResponseWriter,request *http.Request)  {
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.UserCtrl.Register(httpWeb)
}
/**
根据邮箱注册
*/
func (ctrl *BaseUserCtl) RegisterByEmail(writer http.ResponseWriter,request *http.Request){
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.UserCtrl.RegisterByEmail(httpWeb, func(input *dtos.UserEmailInput) bool {
		return ctrl.UserCtrl.ValidatorUserEmailInput(httpWeb,input)
	})
}

/**
根据用户名注册
*/
func (ctrl *BaseUserCtl) RegisterByUserName(writer http.ResponseWriter,request *http.Request){
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.UserCtrl.RegisterByUserName(httpWeb, func(input *dtos.UserUserNameInput) bool {
		return ctrl.UserCtrl.ValidatorUserUserNameInput(httpWeb,input)
	})
}

/**
根据手机号注册
*/
func (ctrl *BaseUserCtl) RegisterByPhone(writer http.ResponseWriter,request *http.Request){
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.UserCtrl.RegisterByPhone(httpWeb, func(input *dtos.UserPhoneInput) bool {
		return ctrl.UserCtrl.ValidatorUserPhoneInput(httpWeb,input)
	})
}

/*
	根据手机号修改手机号
*/
func (ctrl *BaseUserCtl) UpdatePhone(writer http.ResponseWriter,request *http.Request)  {
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.UserCtrl.UpdatePhone(httpWeb)

}

/*
	根据邮箱修改邮箱
*/
func (ctrl *BaseUserCtl) UpdateEmail(writer http.ResponseWriter,request *http.Request)  {
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.UserCtrl.UpdateEmail(httpWeb)
}

/*
	根据手机号修改邮箱
*/
func (ctrl *BaseUserCtl) UpdateEmailByPhone(writer http.ResponseWriter,request *http.Request)  {
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.UserCtrl.UpdateEmailByPhone(httpWeb)
}

/*
	根据邮箱修改密码
*/
func (ctrl *BaseUserCtl) UpdatePwdByEmail(writer http.ResponseWriter,request *http.Request)  {
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.UserCtrl.UpdatePwdByEmail(httpWeb)
}

/*
	根据手机号修改密码
*/
func (ctrl *BaseUserCtl) UpdatePwdByPhone(writer http.ResponseWriter,request *http.Request)  {
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.UserCtrl.UpdatePwdByPhone(httpWeb)
}