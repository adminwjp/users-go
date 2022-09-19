package rpc_thrift_impl

import service "github.com/adminwjp/users-go/services"

type ConfigThriftServiceImpl struct {
	Service func()service.ConfigService
}
