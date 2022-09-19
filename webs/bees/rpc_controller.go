package web_bee_controller

import (
	"github.com/adminwjp/infrastructure-go/webs/bees"
	web "github.com/adminwjp/users-go/webs"
	bee "github.com/beego/beego/v2/server/web"
)

type RpcCtl struct {
	web.RpcCtrl
	bee.Controller
}
func (ctrl *RpcCtl) M() {
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.RpcCtrl.M(httpWeb)
}
/*
	添加
*/
func (ctrl *RpcCtl)Add(){
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.RpcCtrl.Add(httpWeb)
}

/*
	修改
*/
func (ctrl *RpcCtl)Update(){
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.RpcCtrl.Update(httpWeb)
}

/*
	删除
*/
func (ctrl *RpcCtl)Delete(){
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.RpcCtrl.Delete(httpWeb)
}

/*
	删除
*/
func (ctrl *RpcCtl)DeleteBatch(){
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.RpcCtrl.DeleteBatch(httpWeb)
}

func (ctrl *RpcCtl) List(){
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.RpcCtrl.List(httpWeb)
}
