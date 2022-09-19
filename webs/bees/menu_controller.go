package web_bee_controller

import (
	"github.com/adminwjp/infrastructure-go/webs/bees"
	web "github.com/adminwjp/users-go/webs"
	bee "github.com/beego/beego/v2/server/web"
)

type MenuCtl struct {
	web.MenuCtrl
	bee.Controller
}

func (ctrl *MenuCtl)Load()  {
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.MenuCtrl.Load(httpWeb)
}