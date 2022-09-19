package rpc_grpc_impl

import (
	 dto "github.com/adminwjp/infrastructure-go/dtos"
	"github.com/adminwjp/users-go/dtos"
	pb "github.com/adminwjp/users-go/rpcs/impl/service"
)




/*
//必须同一个包 使用
func (user *pb.UserRequest) ToDto()*dto.UserInput  {
	return &dto.UserInput{Account: user.}
}
*/

func  ToUserDto(user *pb.UserRequest)*dtos.UserInput  {
	return &dtos.UserInput{Account: user.Account,Pwd: user.Pwd,Flag: dto.AccounType(user.Flag)}
}

func  ToUserEmailDto(user *pb.UserEmailRequest)*dtos.UserEmailInput  {
	return &dtos.UserEmailInput{Email: user.Email,Pwd: user.Pwd}
}

func  ToUserPhoneDto(user *pb.UserPhoneRequest)*dtos.UserPhoneInput  {
	return &dtos.UserPhoneInput{Phone: user.Phone,Pwd: user.Pwd}
}

func  ToUserUserNameDto(user *pb.UserUserNameRequest)*dtos.UserUserNameInput  {
	return &dtos.UserUserNameInput{UserName: user.UserName,Pwd: user.Pwd}
}

func  ToUpdateUserEmailDto(user *pb.UpdateEmailRequest)*dtos.UpdateUserEmailInput  {
	return &dtos.UpdateUserEmailInput{Email: user.Email,NewEmail: user.NewEmail}
}

func  ToUpdateUserEmailByPhoneDto(user *pb.UpdateEmailByPhoneRequest)*dtos.UpdateUserEmailByPhoneInput  {
	return &dtos.UpdateUserEmailByPhoneInput{Email: user.Email,Phone: user.Phone}
}

func  ToUpdateUserPhoneDto(user *pb.UpdatePhoneRequest)*dtos.UpdateUserPhoneInput  {
	return &dtos.UpdateUserPhoneInput{Phone: user.Phone,NewPhone: user.NewPhone}
}

func  ToUpdateUserPwdByPwdDto(user *pb.UpdatePwdRequest)*dtos.UpdateUserPwdByPwdInput  {
	return &dtos.UpdateUserPwdByPwdInput{Pwd: user.Pwd,NewPwd: user.NewPwd}
}

func  ToUpdateUserPwdByPhoneDto(user *pb.UpdatePwdByPhoneRequest)*dtos.UpdateUserPwdByPhoneInput  {
	return &dtos.UpdateUserPwdByPhoneInput{Pwd: user.Pwd,Phone: user.Phone}
}
func  ToUpdateUserPwdByEmailDto(user *pb.UpdatePwdByEmailRequest)*dtos.UpdateUserPwdByEmailInput  {
	return &dtos.UpdateUserPwdByEmailInput{Pwd: user.Pwd,Email: user.Email}
}
func  ToUpdateUserAuthBasicDto(user *pb.UpdateAuthBasicRequest)*dtos.UpdateUserAuthBasicInput  {
	return &dtos.UpdateUserAuthBasicInput{Id: user.UserId,CardId: user.CardId,
		CardPhoto1: user.CardPhoto1,CardPhoto2: user.CardPhoto2,
		HandCardPhoto1: user.HandCardPhoto1,HandCardPhoto2: user.HandCardPhoto2,
	}
}