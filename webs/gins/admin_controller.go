package web_gin_controller

import (
	"github.com/adminwjp/infrastructure-go/webs/gins"
	dto "github.com/adminwjp/users-go/dtos"
	"github.com/adminwjp/users-go/services"
	web "github.com/adminwjp/users-go/webs"
	"github.com/gin-gonic/gin"
)

type AdminCtl struct {
	BaseUserCtl
	Service func() service.AdminService
	UserCtrl *web.AdminCtrl
}
func (ctrl *AdminCtl) M(c *gin.Context) {
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.UserCtrl.M(httpWeb)
}

/**
根据手机号、邮箱、用户名登录
*/
func (ctrl *AdminCtl) Login(c *gin.Context) {
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.UserCtrl.Login(httpWeb)
}

/**
根据邮箱登录
*/
func (ctrl *AdminCtl) LoginByEmail(c *gin.Context) {
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.UserCtrl.LoginByEmail(httpWeb, func(input *dto.UserEmailInput) bool {
		return ctrl.UserCtrl.ValidatorUserEmailInput(httpWeb,input)
	})
}

/**
根据用户名登录
*/
func (ctrl *AdminCtl) LoginByUserName(c *gin.Context) {
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.UserCtrl.LoginByUserName(httpWeb, func(input *dto.UserUserNameInput) bool {
		return ctrl.UserCtrl.ValidatorUserUserNameInput(httpWeb,input)
	})
}

/**
根据手机号登录
*/
func (ctrl *AdminCtl) LoginByPhone(c *gin.Context) {
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.UserCtrl.LoginByPhone(httpWeb, func(input *dto.UserPhoneInput) bool {
		return ctrl.UserCtrl.ValidatorUserPhoneInput(httpWeb,input)
	})
}

/*
	根据旧密码修改密码
*/
func (ctrl *AdminCtl) UpdatePwdByOldPwd(c *gin.Context) {
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.UserCtrl.UpdatePwdByOldPwd(httpWeb)
}
/*
	根据条件查询用户信息
*/
func (ctrl *AdminCtl) List(c *gin.Context)  {
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.UserCtrl.List(httpWeb)
}

/*
	根据条件查询用户日志信息
*/
func (ctrl *AdminCtl) ListLogs(c *gin.Context)  {
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.UserCtrl.ListLogs(httpWeb)
}