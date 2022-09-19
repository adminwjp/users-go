package rpc_grpc_impl

import service "github.com/adminwjp/users-go/services"

type ConfigGrpcServiceImpl struct {
	Service func()service.ConfigService
}

