package rpc_grpc_impl

import (
	context "context"
	dto "github.com/adminwjp/infrastructure-go/dtos"
	pb "github.com/adminwjp/users-go/rpcs/impl/service"
	"github.com/adminwjp/users-go/services"
)

type BaseUserGrpcServiceImpl struct {
	BaseService func()service.BaseUserService
}

func (user1 BaseUserGrpcServiceImpl) Register(ctx context.Context, in *pb.UserRequest) (*pb.IdGrpc, error) {
	user:=ToUserDto(in)
	service:=user1.BaseService()
	service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	service.Register(user)
	out := new(pb.IdGrpc)
	out.Id=user.Id
	return out, nil
}
func (user1 BaseUserGrpcServiceImpl) RegisterByUserName(ctx context.Context, in *pb.UserUserNameRequest) (*pb.IdGrpc, error) {
	user:=ToUserUserNameDto(in)
	service:=user1.BaseService()
	service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	service.RegisterByUserName(user)
	out := new(pb.IdGrpc)
	out.Id=user.Id
	return out, nil
}
func (user1 BaseUserGrpcServiceImpl) RegisterByPhone(ctx context.Context, in *pb.UserPhoneRequest) (*pb.IdGrpc, error) {
	user:=ToUserPhoneDto(in)
	service:=user1.BaseService()
	service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	service.RegisterByPhone(user)
	out := new(pb.IdGrpc)
	out.Id=user.Id
	return out, nil
}
func (user1 BaseUserGrpcServiceImpl) RegisterByEmail(ctx context.Context, in *pb.UserEmailRequest) (*pb.IdGrpc, error) {
	user:=ToUserEmailDto(in)
	service:=user1.BaseService()
	service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	service.RegisterByEmail(user)
	out := new(pb.IdGrpc)
	out.Id=user.Id
	return out, nil
}


func (user1 BaseUserGrpcServiceImpl) UpdateEmail(ctx context.Context, in *pb.UpdateEmailRequest) (*pb.ResultReply, error) {
	out := new(pb.ResultReply)
	service:=user1.BaseService()
	service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	r,_:=service.UpdateEmailByEmail(ToUpdateUserEmailDto(in))
	out.Result=int32(r)
	return out, nil
}
func (user1 BaseUserGrpcServiceImpl) UpdateEmailByPhone(ctx context.Context, in *pb.UpdateEmailByPhoneRequest) (*pb.ResultReply, error) {
	out := new(pb.ResultReply)
	service:=user1.BaseService()
	service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	r,_:=service.UpdateEmailByPhone(ToUpdateUserEmailByPhoneDto(in))
	out.Result=int32(r)
	return out, nil
}
func (user1 BaseUserGrpcServiceImpl) UpdatePhone(ctx context.Context, in *pb.UpdatePhoneRequest) (*pb.ResultReply, error) {
	out := new(pb.ResultReply)
	service:=user1.BaseService()
	service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	r,_:=service.UpdatePhone(ToUpdateUserPhoneDto(in))
	out.Result=int32(r)
	return out, nil
}
func (user1 BaseUserGrpcServiceImpl) UpdatePwdByEmail(ctx context.Context, in *pb.UpdatePwdByEmailRequest) (*pb.ResultReply, error) {
	out := new(pb.ResultReply)
	service:=user1.BaseService()
	service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	r,_:=service.UpdatePwdByEmail(ToUpdateUserPwdByEmailDto(in))
	out.Result=int32(r)
	return out, nil
}
func (user1 BaseUserGrpcServiceImpl) UpdatePwdByPhone(ctx context.Context, in *pb.UpdatePwdByPhoneRequest) (*pb.ResultReply, error) {
	out := new(pb.ResultReply)
	service:=user1.BaseService()
	service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	r,_:=service.UpdatePwdByPhone(ToUpdateUserPwdByPhoneDto(in))
	out.Result=int32(r)
	return out, nil
}

func (user1 BaseUserGrpcServiceImpl) Exists(ctx context.Context, in *pb.AccountRequest) (*pb.ExistsReply, error) {
	out := new(pb.ExistsReply)
	service:=user1.BaseService()
	//service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	r,_:=service.Exists(in.Account,dto.AccounType(in.Flag))
	out.Exists=r>0
	return out, nil
}
func (user1 BaseUserGrpcServiceImpl) ExistsPhone(ctx context.Context, in *pb.PhoneRequest) (*pb.ExistsReply, error) {
	out := new(pb.ExistsReply)
	service:=user1.BaseService()
	//service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	r,_:=service.ExistsByPhone(in.Phone)
	out.Exists=r>0
	return out, nil
}
func (user1 BaseUserGrpcServiceImpl) ExistsEmail(ctx context.Context, in *pb.EmailRequest) (*pb.ExistsReply, error) {
	out := new(pb.ExistsReply)
	service:=user1.BaseService()
	//service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	r,_:=service.ExistsByEmail(in.Email)
	out.Exists=r>0
	return out, nil
}
func (user1 BaseUserGrpcServiceImpl) ExistsUserName(ctx context.Context, in *pb.UserNameRequest) (*pb.ExistsReply, error) {
	out := new(pb.ExistsReply)
	service:=user1.BaseService()
	//service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	r,_:=service.ExistsByUserName(in.UserName)
	out.Exists=r>0
	return out, nil
}
