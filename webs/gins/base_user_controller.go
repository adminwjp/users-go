package web_gin_controller

import (
	"github.com/adminwjp/infrastructure-go/webs/gins"
	"github.com/adminwjp/users-go/dtos"
	"github.com/adminwjp/users-go/services"
	"github.com/adminwjp/users-go/webs"
	"github.com/gin-gonic/gin"
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
func (ctrl *BaseUserCtl) Register(c *gin.Context)  {
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.UserCtrl.Register(httpWeb)
}
/**
根据邮箱注册
*/
func (ctrl *BaseUserCtl) RegisterByEmail(c *gin.Context){
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.UserCtrl.RegisterByEmail(httpWeb, func(input *dtos.UserEmailInput) bool {
		return ctrl.UserCtrl.ValidatorUserEmailInput(httpWeb,input)
	})
}

/**
根据用户名注册
*/
func (ctrl *BaseUserCtl) RegisterByUserName(c *gin.Context){
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.UserCtrl.RegisterByUserName(httpWeb, func(input *dtos.UserUserNameInput) bool {
		return ctrl.UserCtrl.ValidatorUserUserNameInput(httpWeb,input)
	})
}

/**
根据手机号注册
*/
func (ctrl *BaseUserCtl) RegisterByPhone(c *gin.Context){
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.UserCtrl.RegisterByPhone(httpWeb, func(input *dtos.UserPhoneInput) bool {
		return ctrl.UserCtrl.ValidatorUserPhoneInput(httpWeb,input)
	})
}

/*
	根据手机号修改手机号
*/
func (ctrl *BaseUserCtl) UpdatePhone(c *gin.Context)  {
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.UserCtrl.UpdatePhone(httpWeb)

}

/*
	根据邮箱修改邮箱
*/
func (ctrl *BaseUserCtl) UpdateEmail(c *gin.Context)  {
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.UserCtrl.UpdateEmail(httpWeb)
}

/*
	根据手机号修改邮箱
*/
func (ctrl *BaseUserCtl) UpdateEmailByPhone(c *gin.Context)  {
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.UserCtrl.UpdateEmailByPhone(httpWeb)
}

/*
	根据邮箱修改密码
*/
func (ctrl *BaseUserCtl) UpdatePwdByEmail(c *gin.Context)  {
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.UserCtrl.UpdatePwdByEmail(httpWeb)
}

/*
	根据手机号修改密码
*/
func (ctrl *BaseUserCtl) UpdatePwdByPhone(c *gin.Context)  {
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.UserCtrl.UpdatePwdByPhone(httpWeb)
}