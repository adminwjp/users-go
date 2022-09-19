package rpc_thrift_impl

import (
	rpc_thrift "github.com/adminwjp/infrastructure-go/rpcs/thrifts"
	util "github.com/adminwjp/infrastructure-go/utils"
	"github.com/adminwjp/users-go/rpcs"
	t "github.com/adminwjp/users-go/rpcs/gen-go/users"
	service "github.com/adminwjp/users-go/services"
	"github.com/apache/thrift/lib/go/thrift"
)

var adminThrift=&AdminThriftServiceImpl{}
var userThrift=&UserThriftServiceImpl{}
//init
func Init()  {

	userThrift.Service= rpcs.CreateServiceInstance.UserService
	adminThrift.Service= rpcs.CreateServiceInstance.AdminService

	userThrift.BaseService= func() service.BaseUserService {
		return  userThrift.Service()
	}
	adminThrift.BaseService=func() service.BaseUserService {
		return  adminThrift.Service()
	}
}
func StartThrift(port string,ip string)  {
	if rpc_thrift.ThiriftInstance.Server!=nil{
		rpc_thrift.ThiriftInstance.Server.Stop()
		rpc_thrift.ThiriftInstance.Server=nil
	}
	if ip==""{
		ip=util.NetUtil.GetLocalIP()
	}
	go rpc_thrift.ThiriftInstance.StartThriftServer(ip+port, RegisterThriftServer)
}

func RegisterThriftServer(multiplexedProcessor *thrift.TMultiplexedProcessor)  {
	processor:=t.NewUserThriftServiceProcessor(userThrift)
	multiplexedProcessor.RegisterProcessor("user",processor)

	processor1:=t.NewAdminThriftServiceProcessor(adminThrift)
	multiplexedProcessor.RegisterProcessor("admin",processor1)
}
func RegisterThriftClient(){
	transport,protocolFactory,err:=rpc_thrift.ThiriftInstance.StartThriftClient("127.0.0.1:3000")
	if err!=nil{return}
	client:=t.NewUserThriftServiceClientFactory(transport,protocolFactory)
	if client!=nil{
		client.Login(nil,nil)
	}
}