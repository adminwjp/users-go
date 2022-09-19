package web_bee_controller

import (
	"github.com/adminwjp/infrastructure-go/webs/bees"
	web "github.com/adminwjp/users-go/webs"
	bee "github.com/beego/beego/v2/server/web"
)

type Config1Ctl struct {
	web.Config1Ctrl
	bee.Controller
}
func (ctrl *Config1Ctl) M() {
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.Config1Ctrl.M(httpWeb)
}
/*
	添加
*/
func (ctrl *Config1Ctl)Add(){
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.Config1Ctrl.Add(httpWeb)
}

/*
	修改
*/
func (ctrl *Config1Ctl)Update(){
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.Config1Ctrl.Update(httpWeb)
}

/*
	删除
*/
func (ctrl *Config1Ctl)Delete(){
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.Config1Ctrl.Delete(httpWeb)
}

/*
	删除
*/
func (ctrl *Config1Ctl)DeleteBatch(){
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.Config1Ctrl.DeleteBatch(httpWeb)
}

func (ctrl *Config1Ctl) List(){
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.Config1Ctrl.List(httpWeb)
}