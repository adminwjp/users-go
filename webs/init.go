package web

import (
	"github.com/adminwjp/infrastructure-go/webs"
)

type IRouter interface {
	Start(port int)
}

type IMCtrl interface{
	M(httpWeb webs.HttpWeb)
}
type IUserCtrl interface {
	Login(httpWeb webs.HttpWeb)
	LoginByPhone(httpWeb webs.HttpWeb)
	LoginByEmail(httpWeb webs.HttpWeb)
	LoginByUserName(httpWeb webs.HttpWeb)

	Register(httpWeb webs.HttpWeb)
	RegisterByPhone(httpWeb webs.HttpWeb)
	RegisterByEmail(httpWeb webs.HttpWeb)
	RegisterByUserName(httpWeb webs.HttpWeb)

	Exists(httpWeb webs.HttpWeb)
	ExistsByPhone(httpWeb webs.HttpWeb)
	ExistsByEmail(httpWeb webs.HttpWeb)
	ExistsByUserName(httpWeb webs.HttpWeb)
}
