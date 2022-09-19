package rpc_grpc_impl

import (
	rpc_grpc "github.com/adminwjp/infrastructure-go/rpcs/grpcs"
	util "github.com/adminwjp/infrastructure-go/utils"
	"github.com/adminwjp/users-go/rpcs"
	pb "github.com/adminwjp/users-go/rpcs/impl/service"
	service "github.com/adminwjp/users-go/services"
	"google.golang.org/grpc"
)

var adminGrpc=&AdminGrpcServiceImpl{}
var userGrpc=&UserGrpcServiceImpl{}
//init
func Init()  {


	userGrpc.Service= rpcs.CreateServiceInstance.UserService
	adminGrpc.Service= rpcs.CreateServiceInstance.AdminService

	userGrpc.BaseService= func() service.BaseUserService {
		return  userGrpc.Service()
	}
	adminGrpc.BaseService=func() service.BaseUserService {
		return  adminGrpc.Service()
	}
}

func StartGRpc(port string,ip string)*grpc.Server  {
	if rpc_grpc.GrpcInstance.Server!=nil{
		rpc_grpc.GrpcInstance.Server.Stop()
		rpc_grpc.GrpcInstance.Server=nil
	}
	var server *grpc.Server
	go rpc_grpc.GrpcInstance.StartGrpcServer(port, func(s *grpc.Server) {
		server=s
		RegisterGrpcServer(s)
	})
	if ip==""{
		ip=util.NetUtil.GetLocalIP()
	}
	return  server
}

func RegisterGrpcServer(server *grpc.Server )  {
	pb.RegisterUserGrpcServiceServer(server, userGrpc)
	pb.RegisterAdminGrpcServiceServer(server, adminGrpc)
}