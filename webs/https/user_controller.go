package web_http_controller

import (
	"github.com/adminwjp/infrastructure-go/webs/https"
	dto "github.com/adminwjp/users-go/dtos"
	"github.com/adminwjp/users-go/services"
	web "github.com/adminwjp/users-go/webs"
	"log"
	"net/http"
)

func (ctrl *UserCtl) M(writer http.ResponseWriter,request *http.Request) {
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.UserCtrl.M(httpWeb)
}
func (ctrl *UserCtl) BM(writer http.ResponseWriter,request *http.Request) {
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.UserCtrl.BM(httpWeb)
}
type UserCtl struct {
	BaseUserCtl
	Service func()service.UserService
	UserCtrl *web.UserCtrl
}
/**
根据手机号、邮箱、用户名登录
*/

func (ctrl *UserCtl) Login(writer http.ResponseWriter,request *http.Request){
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	log.Println("gin login start")
	ctrl.UserCtrl.Login(httpWeb)
	log.Println("gin login end")
}

/**
根据邮箱登录
*/
func (ctrl *UserCtl) LoginByEmail(writer http.ResponseWriter,request *http.Request){
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.UserCtrl.LoginByEmail(httpWeb, func(input *dto.UserEmailInput) bool {
		return ctrl.UserCtrl.ValidatorUserEmailInput(httpWeb,input)
	})
}

/**
根据用户名登录
*/
func (ctrl *UserCtl) LoginByUserName(writer http.ResponseWriter,request *http.Request){
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.UserCtrl.LoginByUserName(httpWeb, func(input *dto.UserUserNameInput) bool {
		return ctrl.UserCtrl.ValidatorUserUserNameInput(httpWeb,input)
	})
}

/**
根据手机号登录
*/
func (ctrl *UserCtl) LoginByPhone(writer http.ResponseWriter,request *http.Request){
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.UserCtrl.LoginByPhone(httpWeb, func(input *dto.UserPhoneInput) bool {
		return ctrl.UserCtrl.ValidatorUserPhoneInput(httpWeb,input)
	})
}

/**
根据手机号、邮箱、用户名注册
*/
func (ctrl *UserCtl) Register(writer http.ResponseWriter,request *http.Request)  {
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.UserCtrl.Register(httpWeb)
}



/*
	修改身份认证基本信息
*/
func (ctrl *UserCtl) UpdateInfo(writer http.ResponseWriter,request *http.Request)  {
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.UserCtrl.UpdateInfo(httpWeb)
}

/*
	根据条件查询用户信息
*/
func (ctrl *UserCtl) List(writer http.ResponseWriter,request *http.Request)  {
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.UserCtrl.List(httpWeb)
}

/*
	根据条件查询用户日志信息
*/
func (ctrl *UserCtl) ListLogs(writer http.ResponseWriter,request *http.Request)  {
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.UserCtrl.ListLogs(httpWeb)
}



