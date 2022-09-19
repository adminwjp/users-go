package rpc_grpc_impl

import service "github.com/adminwjp/users-go/services"

type RpcGrpcServiceImpl struct {
	Service func()service.RpcService
}
