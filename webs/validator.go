package web

import (
	dto "github.com/adminwjp/infrastructure-go/dtos"
	util "github.com/adminwjp/infrastructure-go/utils"
	"github.com/adminwjp/infrastructure-go/webs"
	"github.com/adminwjp/users-go/dtos"
	"regexp"
)

func (ctrl *UserCtrl)ValidatorUpdateUserAuthBasicInput(httpWeb webs.HttpWeb,input *dtos.UpdateUserAuthBasicInput) bool {
	var res *dto.ResponseDto
	if input.Id<1{
		res= &dto.ResponseDto{Status: false,Code:401,Msg:"id error"}
	}
	if !util.RegexUtil.IsMatch(input.CardId,"\\d{16,18}[a-z|A-Z]{0,2}"){
		res= &dto.ResponseDto{Status: false,Code:401,Msg:"card_id error"}
	}
	if input.CardPhoto1==""||len(input.CardPhoto1)<5{
		res= &dto.ResponseDto{Status: false,Code:401,Msg:"card_photo1 error"}
	}
	if input.CardPhoto2==""||len(input.CardPhoto2)<5{
		res= &dto.ResponseDto{Status: false,Code:401,Msg:"card_photo2 error"}
	}
	if input.HandCardPhoto1==""||len(input.HandCardPhoto1)<5{
		res= &dto.ResponseDto{Status: false,Code:401,Msg:"hand_card_photo1 error"}
	}
	if input.HandCardPhoto2==""||len(input.HandCardPhoto2)<5{
		res= &dto.ResponseDto{Status: false,Code:401,Msg:"hand_card_photo2 error"}
	}
	if res!=nil{
		httpWeb.Response(200,res)
		return false
	}
	return true
}
func (ctrl *BaseUserCtrl)ValidatorUserPhoneInput(httpWeb webs.HttpWeb,input *dtos.UserPhoneInput) bool {
	if input.Phone=="" ||!util.RegexUtil.IsPhone(input.Phone){
		res:= dto.ResponseDto{Status: false,Code:   402,Msg:    "phone error"}
		httpWeb.Response(200,res)
		return false
	}
	if input.Pwd==""{
		res:= dto.ResponseDto{Status: false,Code:   403,Msg:    "pwd error"}
		httpWeb.Response(200,res)
		return false
	}
	return true
}

func (ctrl *BaseUserCtrl)ValidatorUserUserNameInput(httpWeb webs.HttpWeb,input *dtos.UserUserNameInput) bool {
	if input.UserName=="" {
		res:= dto.ResponseDto{Status: false,Code:402,Msg:"user_name error"}
		httpWeb.Response(200,res)
		return false
	}
	if input.Pwd==""{
		res:= dto.ResponseDto{Status: false,Code:403,Msg:"pwd error"}
		httpWeb.Response(200,res)
		return false
	}
	return true
}
func (ctrl *BaseUserCtrl)ValidatorUserEmailInput(httpWeb webs.HttpWeb,input *dtos.UserEmailInput) bool {
	if input.Email=="" ||!util.RegexUtil.IsEmail(input.Email){
		res:= dto.ResponseDto{Status: false,Code:402,Msg:"email error"}
		httpWeb.Response(200,res)
		return false
	}
	if input.Pwd==""{
		res:=dto.ResponseDto{Status: false,Code:403,Msg:"pwd error"}
		httpWeb.Response(200,res)
		return false
	}
	return true
}
func (ctrl *BaseUserCtrl) ValidatorUserInput(httpWeb webs.HttpWeb,input *dtos.UserInput) bool {

	/*if input.Account!="" {
		match, _ := regexp.Match("[13|15|17|18]\\d{9}", []byte(input.Account))
		if match {
			input.Flag=dto.AccounTypeByPhone
		}else if util.RegexUtil.IsEmail(input.Account){
			input.Flag=dto.AccounTypeByEamil
		}else{
			input.Flag=dto.AccounTypeByUsername
		}
	}*/
	if input.Account!="" {
		if match, _ := regexp.Match("[13|15|17|18]\\d{9}", []byte(input.Account));input.Flag==dto.AccounTypeByPhone&&!match {
			res:=dto.ResponseDto{Status: false,Code:402,Msg:"account error,not is phone"}
			httpWeb.Response(200,res)
			return false
		}
		if input.Flag==dto.AccounTypeByEamil&&!util.RegexUtil.IsEmail(input.Account){
			res:=dto.ResponseDto{Status: false,Code:402,Msg:"account error,not is email"}
			httpWeb.Response(200,res)
			return false
		}
		if match, _ := regexp.Match("[a-z|A-Z|\\-|0-9]{4,9}", []byte(input.Account));input.Flag==dto.AccounTypeByUsername&&!match{
			res:=dto.ResponseDto{Status: false,Code:402,Msg:"account error,user_name error "}
			httpWeb.Response(200,res)
			return false
		}


	}else{
		res:=dto.ResponseDto{Status: false,Code:402,Msg:"account error"}
		httpWeb.Response(200,res)
		return false
	}
	if input.Pwd==""{
		res:=dto.ResponseDto{Status: false,Code:403,Msg:"pwd error"}
		httpWeb.Response(200,res)
		return false
	}
	return true
}
