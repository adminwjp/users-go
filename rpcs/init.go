package rpcs

import service "github.com/adminwjp/users-go/services"
var CreateServiceInstance CreateService
type CreateService interface {
	AdminService()  service.AdminService
	UserService () service.UserService
	RoleService () service.RoleService
	EmailService () service.EmailService
	PaySecrtService () service.PaySecrtService
	SmsService () service.SmsService
	ConfigService () service.ConfigService
	RpcService () service.RpcService
}
