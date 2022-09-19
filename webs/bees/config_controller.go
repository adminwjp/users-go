package web_bee_controller

import (
	"github.com/adminwjp/infrastructure-go/webs/bees"
	web "github.com/adminwjp/users-go/webs"
	bee "github.com/beego/beego/v2/server/web"
)

type ConfigCtl struct {
	bee.Controller
	web.ConfigCtrl
}
func (ctrl *ConfigCtl) Get()  {
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.ConfigCtrl.Get(httpWeb)
}
func (ctrl *ConfigCtl) SetFlag()  {
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.ConfigCtrl.SetFlag(httpWeb)
}