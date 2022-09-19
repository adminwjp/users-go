package rpc_thrift_impl

import service "github.com/adminwjp/users-go/services"

type RpcThriftServiceImpl struct {
	Service func()service.RpcService
}

