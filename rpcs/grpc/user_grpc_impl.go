package rpc_grpc_impl

import (
	context "context"
	 pb "github.com/adminwjp/users-go/rpcs/impl/service"
	"github.com/adminwjp/users-go/services"
	util "github.com/adminwjp/infrastructure-go/utils"
)

type UserGrpcServiceImpl struct {
	BaseUserGrpcServiceImpl
	Service func()service.UserService
}
func (user1 UserGrpcServiceImpl) Login(ctx context.Context, in *pb.UserRequest) (*pb.UserReply, error) {
	user:=ToUserDto(in)
	service:=user1.Service()
	//service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	userModel,_:=service.Login(user)
	out := new(pb.UserReply)
	if userModel!=nil{
		util.MapTo(userModel,out)
		//out.Id=userModel.Id
	}
	return out, nil
}
func (user1 UserGrpcServiceImpl) LoginByUserName(ctx context.Context, in *pb.UserUserNameRequest) (*pb.UserReply, error) {
	user:=ToUserUserNameDto(in)
	service:=user1.Service()
	//service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	userModel,_:=service.LoginByUserName(user)
	out := new(pb.UserReply)
	if user!=nil{
		util.MapTo(userModel,out)
	}
	return out, nil
}
func (user1 UserGrpcServiceImpl) LoginByPhone(ctx context.Context, in *pb.UserPhoneRequest) (*pb.UserReply, error) {
	user:=ToUserPhoneDto(in)
	service:=user1.Service()
	//service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	userModel,_:=service.LoginByPhone(user)
	out := new(pb.UserReply)
	if user!=nil{
		util.MapTo(userModel,out)
	}
	return out, nil
}
func (user1 UserGrpcServiceImpl) LoginByEmail(ctx context.Context, in *pb.UserEmailRequest) (*pb.UserReply, error) {
	user:=ToUserEmailDto(in)
	service:=user1.Service()
	//service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	userModel,_:=service.LoginByEmail(user)
	out := new(pb.UserReply)
	if user!=nil{
		util.MapTo(userModel,out)
	}
	return out, nil
}


func (user1 UserGrpcServiceImpl) Get(ctx context.Context, in *pb.IdGrpc) (*pb.UserReply, error) {
	service:=user1.Service()
	//service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	userModel,_:=service.Get(in.Id)
	out := new(pb.UserReply)
	if userModel!=nil{
		util.MapTo(userModel,out)
	}
	return out, nil
}
func (user1 UserGrpcServiceImpl) GetAuthBasic(ctx context.Context, in *pb.IdGrpc) (*pb.UpdateAuthBasicRequest, error) {
	out := new(pb.UpdateAuthBasicRequest)
	return out, nil
}
func (user1 UserGrpcServiceImpl) UpdateAuthBasic(ctx context.Context, in *pb.UpdateAuthBasicRequest) (*pb.ResultReply, error) {
	out := new(pb.ResultReply)
	service:=user1.Service()
	//service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	r,_:=service.UpdateAuthBasic(ToUpdateUserAuthBasicDto(in))
	out.Result=int32(r)
	return out, nil
}
