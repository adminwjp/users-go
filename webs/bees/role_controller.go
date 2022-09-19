package web_bee_controller

import (
	"github.com/adminwjp/infrastructure-go/webs/bees"
	web "github.com/adminwjp/users-go/webs"
	bee "github.com/beego/beego/v2/server/web"
)

type RoleCtl struct {
	bee.Controller
	web.RoleCtrl
}
func (ctrl *RoleCtl) M() {
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.RoleCtrl.M(httpWeb)
}
/*
	添加
*/
func (ctrl *RoleCtl)Add(){
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.RoleCtrl.Add(httpWeb)
}

/*
	修改
*/
func (ctrl *RoleCtl)Update(){
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.RoleCtrl.Update(httpWeb)
}

/*
	删除
*/
func (ctrl *RoleCtl)Delete(){
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.RoleCtrl.Delete(httpWeb)
}

/*
	删除
*/
func (ctrl *RoleCtl)DeleteBatch(){
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.RoleCtrl.DeleteBatch(httpWeb)
}

func (ctrl *RoleCtl) List(){
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.RoleCtrl.List(httpWeb)
}

func (ctrl *RoleCtl) Parent(){
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.RoleCtrl.Parent(httpWeb)
}