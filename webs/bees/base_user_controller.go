package web_bee_controller

import (
	"github.com/adminwjp/infrastructure-go/webs/bees"
	"github.com/adminwjp/users-go/dtos"
	"github.com/adminwjp/users-go/services"
	"github.com/adminwjp/users-go/webs"
	bee "github.com/beego/beego/v2/server/web"
)

type BaseUserCtl struct {
	bee.Controller
	//BaseUserService func()service.BaseUserService
	BaseService func()service.BaseUserService
	UserCtrl *web.BaseUserCtrl
}






/**
根据手机号、邮箱、用户名注册
*/
func (ctrl *BaseUserCtl) Register()  {
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.UserCtrl.Register(httpWeb)
}
/**
根据邮箱注册
*/
func (ctrl *BaseUserCtl) RegisterByEmail(){
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.UserCtrl.RegisterByEmail(httpWeb, func(input *dtos.UserEmailInput) bool {
		return ctrl.UserCtrl.ValidatorUserEmailInput(httpWeb,input)
	})
}

/**
根据用户名注册
*/
func (ctrl *BaseUserCtl) RegisterByUserName(){
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.UserCtrl.RegisterByUserName(httpWeb, func(input *dtos.UserUserNameInput) bool {
		return ctrl.UserCtrl.ValidatorUserUserNameInput(httpWeb,input)
	})
}

/**
根据手机号注册
*/
func (ctrl *BaseUserCtl) RegisterByPhone(){
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.UserCtrl.RegisterByPhone(httpWeb, func(input *dtos.UserPhoneInput) bool {
		return ctrl.UserCtrl.ValidatorUserPhoneInput(httpWeb,input)
	})
}

/*
	根据手机号修改手机号
*/
func (ctrl *BaseUserCtl) UpdatePhone()  {
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.UserCtrl.UpdatePhone(httpWeb)

}

/*
	根据邮箱修改邮箱
*/
func (ctrl *BaseUserCtl) UpdateEmail()  {
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.UserCtrl.UpdateEmail(httpWeb)
}

/*
	根据手机号修改邮箱
*/
func (ctrl *BaseUserCtl) UpdateEmailByPhone()  {
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.UserCtrl.UpdateEmailByPhone(httpWeb)
}

/*
	根据邮箱修改密码
*/
func (ctrl *BaseUserCtl) UpdatePwdByEmail()  {
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.UserCtrl.UpdatePwdByEmail(httpWeb)
}

/*
	根据手机号修改密码
*/
func (ctrl *BaseUserCtl) UpdatePwdByPhone()  {
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.UserCtrl.UpdatePwdByPhone(httpWeb)
}