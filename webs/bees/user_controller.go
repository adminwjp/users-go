package web_bee_controller

import (
	"github.com/adminwjp/infrastructure-go/webs/bees"
	dto "github.com/adminwjp/users-go/dtos"
	"github.com/adminwjp/users-go/services"
	web "github.com/adminwjp/users-go/webs"
	bee "github.com/beego/beego/v2/server/web"
	"log"
)

func (ctrl *UserCtl) M() {
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.UserCtrl.M(httpWeb)
}
func (ctrl *UserCtl) BM() {
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.UserCtrl.BM(httpWeb)
}
type UserCtl struct {
	bee.Controller
	BaseUserCtl
	Service func()service.UserService
	UserCtrl *web.UserCtrl
}
/**
根据手机号、邮箱、用户名登录
*/
func (ctrl *UserCtl) Login(){
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	log.Println("gin login start")
	ctrl.UserCtrl.Login(httpWeb)
	log.Println("gin login end")
}

/**
根据邮箱登录
*/
func (ctrl *UserCtl) LoginByEmail(){
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.UserCtrl.LoginByEmail(httpWeb, func(input *dto.UserEmailInput) bool {
		return ctrl.UserCtrl.ValidatorUserEmailInput(httpWeb,input)
	})
}

/**
根据用户名登录
*/
func (ctrl *UserCtl) LoginByUserName(){
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.UserCtrl.LoginByUserName(httpWeb, func(input *dto.UserUserNameInput) bool {
		return ctrl.UserCtrl.ValidatorUserUserNameInput(httpWeb,input)
	})
}

/**
根据手机号登录
*/
func (ctrl *UserCtl) LoginByPhone(){
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.UserCtrl.LoginByPhone(httpWeb, func(input *dto.UserPhoneInput) bool {
		return ctrl.UserCtrl.ValidatorUserPhoneInput(httpWeb,input)
	})
}

/**
根据手机号、邮箱、用户名注册
*/
func (ctrl *UserCtl) Register()  {
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.UserCtrl.Register(httpWeb)
}



/*
	修改身份认证基本信息
*/
func (ctrl *UserCtl) UpdateInfo()  {
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.UserCtrl.UpdateInfo(httpWeb)
}

/*
	根据条件查询用户信息
*/
func (ctrl *UserCtl) List()  {
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.UserCtrl.List(httpWeb)
}

/*
	根据条件查询用户日志信息
*/
func (ctrl *UserCtl) ListLogs()  {
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.UserCtrl.ListLogs(httpWeb)
}



