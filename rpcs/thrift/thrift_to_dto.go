package rpc_thrift_impl



import (
	dto  "github.com/adminwjp/infrastructure-go/dtos"
	"github.com/adminwjp/users-go/dtos"
	t "github.com/adminwjp/users-go/rpcs/gen-go/users"
)


/*
//必须同一个包 使用
func (user *t.UserInputThrift) ToDto()*dto.UserInput  {
	return &dto.UserInput{Account: user.}
}
*/

func  ToUserDto(user *t.UserInputThrift)*dtos.UserInput  {
	return &dtos.UserInput{Account: user.Account,Pwd: user.Pwd,Flag: dto.AccounType(user.Flag)}
}

func  ToUserEmailDto(user *t.UserEmailInputThrift)*dtos.UserEmailInput  {
	return &dtos.UserEmailInput{Email: user.Email,Pwd: user.Pwd}
}

func  ToUserPhoneDto(user *t.UserPhoneInputThrift)*dtos.UserPhoneInput  {
	return &dtos.UserPhoneInput{Phone: user.Phone,Pwd: user.Pwd}
}

func  ToUserUserNameDto(user *t.UserUserNameInputThrift)*dtos.UserUserNameInput  {
	return &dtos.UserUserNameInput{UserName: user.UserName,Pwd: user.Pwd}
}

func  ToUpdateUserEmailDto(user *t.UpdateEmailInputThrift)*dtos.UpdateUserEmailInput  {
	return &dtos.UpdateUserEmailInput{Email: user.Email,NewEmail: user.NewEmail_}
}

func  ToUpdateUserEmailByPhoneDto(user *t.UpdateEmailByPhoneInputThrift)*dtos.UpdateUserEmailByPhoneInput  {
	return &dtos.UpdateUserEmailByPhoneInput{Email: user.Email,Phone: user.Phone}
}

func  ToUpdateUserPhoneDto(user *t.UpdatePhoneInputThrift)*dtos.UpdateUserPhoneInput  {
	return &dtos.UpdateUserPhoneInput{Phone: user.Phone,NewPhone: user.NewPhone_}
}

func  ToUpdateUserPwdByPwdDto(user *t.UpdatePwdInputThrift)*dtos.UpdateUserPwdByPwdInput  {
	return &dtos.UpdateUserPwdByPwdInput{Pwd: user.Pwd,NewPwd: user.NewPwd_}
}

func  ToUpdateUserPwdByPhoneDto(user *t.UpdatePwdByPhoneInputThrift)*dtos.UpdateUserPwdByPhoneInput  {
	return &dtos.UpdateUserPwdByPhoneInput{Pwd: user.Pwd,Phone: user.Phone}
}
func  ToUpdateUserPwdByEmailDto(user *t.UpdatePwdByEmailInputThrift)*dtos.UpdateUserPwdByEmailInput  {
	return &dtos.UpdateUserPwdByEmailInput{Pwd: user.Pwd,Email: user.Email}
}
func  ToUpdateUserAuthBasicDto(user *t.UpdateAuthBasicInputThrift)*dtos.UpdateUserAuthBasicInput  {
	return &dtos.UpdateUserAuthBasicInput{Id: user.UserID,CardId: user.CardID,
		CardPhoto1: user.CardPhoto1,CardPhoto2: user.CardPhoto2,
		HandCardPhoto1: user.HandCardPhoto1,HandCardPhoto2: user.HandCardPhoto2,
	}
}