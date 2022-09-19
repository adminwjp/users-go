package rpc_grpc_impl

import (
	context "context"
	 pb "github.com/adminwjp/users-go/rpcs/impl/service"
	"github.com/adminwjp/users-go/services"
	util "github.com/adminwjp/infrastructure-go/utils"
)

type AdminGrpcServiceImpl struct {
	BaseUserGrpcServiceImpl
	Service func()service.AdminService
}
func (admin AdminGrpcServiceImpl) Login(ctx context.Context, in *pb.UserRequest) (*pb.AdminReply, error) {
	user:=ToUserDto(in)
	service:=admin.Service()
	//service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	adminModel,_:=service.Login(user)
	out := new(pb.AdminReply)
	if adminModel!=nil{
		util.MapTo(adminModel,out)
	}
	return out, nil
}
func (admin AdminGrpcServiceImpl) LoginByUserName(ctx context.Context, in *pb.UserUserNameRequest) (*pb.AdminReply, error) {
	user:=ToUserUserNameDto(in)
	service:=admin.Service()
	//service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	adminModel,_:=service.LoginByUserName(user)
	out := new(pb.AdminReply)
	if adminModel!=nil{
		util.MapTo(adminModel,out)
	}
	return out, nil
}
func (admin AdminGrpcServiceImpl) LoginByPhone(ctx context.Context, in *pb.UserPhoneRequest) (*pb.AdminReply, error) {
	user:=ToUserPhoneDto(in)
	service:=admin.Service()
	//service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	adminModel,_:=service.LoginByPhone(user)
	out := new(pb.AdminReply)
	if adminModel!=nil{
		util.MapTo(adminModel,out)
	}
	return out, nil
}
func (admin AdminGrpcServiceImpl) LoginByEmail(ctx context.Context, in *pb.UserEmailRequest) (*pb.AdminReply, error) {
	user:=ToUserEmailDto(in)
	service:=admin.Service()
	//service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	adminModel,_:=service.LoginByEmail(user)
	out := new(pb.AdminReply)
	if adminModel!=nil{
		util.MapTo(adminModel,out)
	}
	return out, nil
}

func (admin AdminGrpcServiceImpl) Get(ctx context.Context, in *pb.IdGrpc) (*pb.UserReply, error) {
	service:=admin.Service()
	//service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	adminModel,_:=service.Get(in.Id)
	out := new(pb.UserReply)
	if adminModel!=nil{
		util.MapTo(adminModel,out)
	}
	return out, nil
}

func (admin AdminGrpcServiceImpl) UpdatePwd(ctx context.Context, in *pb.UpdatePwdRequest) (*pb.ResultReply, error) {
	out := new(pb.ResultReply)
	service:=admin.Service()
	service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	r,_:=service.UpdatePwdByOldPwd(ToUpdateUserPwdByPwdDto(in))
	out.Result=int32(r)
	return out, nil
}
