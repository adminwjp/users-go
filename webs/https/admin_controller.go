package web_http_controller

import (
	"github.com/adminwjp/infrastructure-go/webs/https"
	dto "github.com/adminwjp/users-go/dtos"
	"github.com/adminwjp/users-go/services"
	web "github.com/adminwjp/users-go/webs"
	"net/http"
)

type AdminCtl struct {
	BaseUserCtl
	Service func() service.AdminService
	UserCtrl *web.AdminCtrl
}
func (ctrl *AdminCtl) M(writer http.ResponseWriter,request *http.Request) {
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.UserCtrl.M(httpWeb)
}

/**
根据手机号、邮箱、用户名登录
*/
func (ctrl *AdminCtl) Login(writer http.ResponseWriter,request *http.Request) {
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.UserCtrl.Login(httpWeb)
}

/**
根据邮箱登录
*/
func (ctrl *AdminCtl) LoginByEmail(writer http.ResponseWriter,request *http.Request) {
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.UserCtrl.LoginByEmail(httpWeb, func(input *dto.UserEmailInput) bool {
		return ctrl.UserCtrl.ValidatorUserEmailInput(httpWeb,input)
	})
}

/**
根据用户名登录
*/
func (ctrl *AdminCtl) LoginByUserName(writer http.ResponseWriter,request *http.Request) {
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.UserCtrl.LoginByUserName(httpWeb, func(input *dto.UserUserNameInput) bool {
		return ctrl.UserCtrl.ValidatorUserUserNameInput(httpWeb,input)
	})
}

/**
根据手机号登录
*/
func (ctrl *AdminCtl) LoginByPhone(writer http.ResponseWriter,request *http.Request) {
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.UserCtrl.LoginByPhone(httpWeb, func(input *dto.UserPhoneInput) bool {
		return ctrl.UserCtrl.ValidatorUserPhoneInput(httpWeb,input)
	})
}

/*
	根据旧密码修改密码
*/
func (ctrl *AdminCtl) UpdatePwdByOldPwd(writer http.ResponseWriter,request *http.Request) {
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.UserCtrl.UpdatePwdByOldPwd(httpWeb)
}
/*
	根据条件查询用户信息
*/
func (ctrl *AdminCtl) List(writer http.ResponseWriter,request *http.Request)  {
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.UserCtrl.List(httpWeb)
}

/*
	根据条件查询用户日志信息
*/
func (ctrl *AdminCtl) ListLogs(writer http.ResponseWriter,request *http.Request)  {
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.UserCtrl.ListLogs(httpWeb)
}