package web

import (
	dto "github.com/adminwjp/infrastructure-go/dtos"
	"github.com/adminwjp/infrastructure-go/webs"
	"github.com/adminwjp/users-go/dtos"
	"github.com/adminwjp/users-go/models"
	service "github.com/adminwjp/users-go/services"
)

func (ctrl *AdminCtrl)M(httpWeb webs.HttpWeb){
	res:=dto.ResponseDataDto{Status: true,Code:200,Msg:"success ",Data:&models.AdminModel{}}
	httpWeb.Response(200,res)
}

func (ctrl *AdminCtrl) Login(httpWeb webs.HttpWeb){

	var login=func(input *dtos.UserInput,service1 service.BaseUserService){
		if r,_:=service1.Exists(input.Account,input.Flag);r<1{
			ctrl.AccountNotExists(httpWeb)
			return
		}
		userModel,_:=service1.(service.AdminService).Login(input)
		ctrl.SetAdminLogin(httpWeb, userModel)
	}
	ctrl.Login1(httpWeb,login,true)

}


func (ctrl *AdminCtrl) LoginByEmail(httpWeb webs.HttpWeb,validate func(input *dtos.UserEmailInput)bool){
	var login= func(input *dtos.UserEmailInput,service1 service.BaseUserService) {
		if r,_:=service1.ExistsByEmail(input.Email);r<1{
			ctrl.EmailNotExists(httpWeb)
			return
		}
		userModel,_:=service1.(service.AdminService).LoginByEmail(input)
		ctrl.SetAdminLogin(httpWeb, userModel)
	}
	ctrl.LoginByEmail1(httpWeb,validate,login,true)
}



func (ctrl *AdminCtrl) LoginByUserName(httpWeb webs.HttpWeb,validate func(input *dtos.UserUserNameInput)bool){
	var login= func(input *dtos.UserUserNameInput,service1 service.BaseUserService) {
		if r,_:=service1.ExistsByUserName(input.UserName);r<1{
			ctrl.UserNameNotExists(httpWeb)
			return
		}
		userModel,_:=service1.(service.AdminService).LoginByUserName(input)
		ctrl.SetAdminLogin(httpWeb, userModel)
	}
	ctrl.LoginByUserName1(httpWeb,validate,login,true)
}




func (ctrl *AdminCtrl) LoginByPhone(httpWeb webs.HttpWeb,validate func(input *dtos.UserPhoneInput)bool){
	var login= func(input *dtos.UserPhoneInput,service1 service.BaseUserService) {
		if r,_:=service1.ExistsByPhone(input.Phone);r<1{
			ctrl.PhoneNotExists(httpWeb)
			return
		}
		userModel,_:=service1.(service.AdminService).LoginByPhone(input)
		ctrl.SetAdminLogin(httpWeb, userModel)
	}
	ctrl.LoginByPhone1(httpWeb,validate,login,true)

}
