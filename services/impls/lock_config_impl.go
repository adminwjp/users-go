package service_impl

import (
	 "github.com/adminwjp/infrastructure-go/locks"
)

type LockConfigImpl struct {

}
var LockInstance locks.Lock
func (config *LockConfigImpl)LockByAdmin(service1 *AdminServiceImpl)  {
	service1.Lock=LockInstance


}

func (config *LockConfigImpl)LockByUser(service1 *UserServiceImpl)  {
	service1.Lock=LockInstance
}