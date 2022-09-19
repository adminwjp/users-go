package web_bee_controller

import (
	"github.com/adminwjp/infrastructure-go/webs/bees"
	dto "github.com/adminwjp/users-go/dtos"
	"github.com/adminwjp/users-go/services"
	web "github.com/adminwjp/users-go/webs"
)

type AdminCtl struct {
	BaseUserCtl
	Service func() service.AdminService
	UserCtrl *web.AdminCtrl
}
func (ctrl *AdminCtl) M() {
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.UserCtrl.M(httpWeb)
}

/**
根据手机号、邮箱、用户名登录
*/
func (ctrl *AdminCtl) Login() {
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.UserCtrl.Login(httpWeb)
}

/**
根据邮箱登录
*/
func (ctrl *AdminCtl) LoginByEmail() {
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.UserCtrl.LoginByEmail(httpWeb, func(input *dto.UserEmailInput) bool {
		return ctrl.UserCtrl.ValidatorUserEmailInput(httpWeb,input)
	})
}

/**
根据用户名登录
*/
func (ctrl *AdminCtl) LoginByUserName() {
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.UserCtrl.LoginByUserName(httpWeb, func(input *dto.UserUserNameInput) bool {
		return ctrl.UserCtrl.ValidatorUserUserNameInput(httpWeb,input)
	})
}

/**
根据手机号登录
*/
func (ctrl *AdminCtl) LoginByPhone() {
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.UserCtrl.LoginByPhone(httpWeb, func(input *dto.UserPhoneInput) bool {
		return ctrl.UserCtrl.ValidatorUserPhoneInput(httpWeb,input)
	})
}

/*
	根据旧密码修改密码
*/
func (ctrl *AdminCtl) UpdatePwdByOldPwd() {
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.UserCtrl.UpdatePwdByOldPwd(httpWeb)
}
/*
	根据条件查询用户信息
*/
func (ctrl *AdminCtl) List()  {
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.UserCtrl.List(httpWeb)
}

/*
	根据条件查询用户日志信息
*/
func (ctrl *AdminCtl) ListLogs()  {
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.UserCtrl.ListLogs(httpWeb)
}