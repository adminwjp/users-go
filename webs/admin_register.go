package web

import (
	"github.com/adminwjp/infrastructure-go/webs"
	"github.com/adminwjp/users-go/dtos"
	service "github.com/adminwjp/users-go/services"
)

func (ctrl *AdminCtrl) Register(httpWeb webs.HttpWeb){
	var register=func(input *dtos.UserInput,service1 service.BaseUserService){
		if r,_:=service1.Exists(input.Account,input.Flag);r>0{

			ctrl.AccountExists(httpWeb)
			return
		}
		rows,_:=service1.Register(input)
		ctrl.SetRegister(httpWeb, rows)
	}
	ctrl.Login1(httpWeb,register,false)
}

func (ctrl *AdminCtrl) RegisterByEmail(httpWeb webs.HttpWeb,validate func(input *dtos.UserEmailInput)bool){
	var login= func(input *dtos.UserEmailInput,service service.BaseUserService) {
		if r,_:=service.ExistsByEmail(input.Email);r>0{
			ctrl.EmailExists(httpWeb)
			return
		}
		rows,_:=service.RegisterByEmail(input)
		ctrl.SetRegister(httpWeb, rows)
	}
	ctrl.LoginByEmail1(httpWeb,validate,login,false)

}

func (ctrl *AdminCtrl) RegisterByUserName(httpWeb webs.HttpWeb,validate func(input *dtos.UserUserNameInput)bool){
	var login= func(input *dtos.UserUserNameInput,service service.BaseUserService) {
		if r,_:=service.ExistsByUserName(input.UserName);r>0{
			ctrl.UserNameExists(httpWeb)
			return
		}
		rows,_:=service.RegisterByUserName(input)
		ctrl.SetRegister(httpWeb, rows)
	}
	ctrl.LoginByUserName1(httpWeb,validate,login,false)

}

func (ctrl *AdminCtrl) RegisterByPhone(httpWeb webs.HttpWeb,validate func(input *dtos.UserPhoneInput)bool){
	var login= func(input *dtos.UserPhoneInput,service service.BaseUserService) {
		if r,_:=service.ExistsByPhone(input.Phone);r>0{
			ctrl.PhoneExists(httpWeb)
			return
		}
		rows,_:=service.RegisterByPhone(input)
		ctrl.SetRegister(httpWeb, rows)
	}
	ctrl.LoginByPhone1(httpWeb,validate,login,false)

}


