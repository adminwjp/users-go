package service_impl

import (
	retry "github.com/adminwjp/infrastructure-go/retries"
	"github.com/adminwjp/users-go/datas"
)

type RetryConfigImpl struct {

}
var EmptyRetryInstance=& retry.EmptyRetry{}
func (config *RetryConfigImpl)RetryByAdmin(service1 *AdminServiceImpl)  {
	switch datas.GlobalConfig.RetryFlag {
		case datas.RetryEmpty:service1.Retry=EmptyRetryInstance
		break
		case datas.RetryLocal,datas.RetryRemote:
		default:
			break

	}
}
func (config *LockConfigImpl)LocalRetryByAdmin(serivce *AdminServiceImpl)  {


}
func (config *LockConfigImpl)RedisRetryAdmin(serivce *AdminServiceImpl)  {


}
func (config *LockConfigImpl)ConsulRetryAdmin(serivce *AdminServiceImpl)  {


}

func (config *LockConfigImpl)ZookeeperRetryAdmin(serivce *AdminServiceImpl)  {


}
func (config *RetryConfigImpl)RetryByUser(service1 *UserServiceImpl)  {

	switch datas.GlobalConfig.RetryFlag {
	case datas.RetryEmpty:service1.Retry=EmptyRetryInstance
		break
	case datas.RetryLocal,datas.RetryRemote:
	default:
		service1.Retry=EmptyRetryInstance
		break

	}

}