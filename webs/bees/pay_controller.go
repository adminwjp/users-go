package web_bee_controller

import (
	"github.com/adminwjp/infrastructure-go/webs/bees"
	web "github.com/adminwjp/users-go/webs"
	bee "github.com/beego/beego/v2/server/web"
)

type PayCtl struct {
	bee.Controller
	web.PayCtrl
}

func (ctrl *PayCtl) M() {
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.PayCtrl.M(httpWeb)
}
/*
	添加
*/
func (ctrl *PayCtl)Add(){
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.PayCtrl.Add(httpWeb)
}

/*
	修改
*/
func (ctrl *PayCtl)Update(){
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.PayCtrl.Update(httpWeb)
}

/*
	删除
*/
func (ctrl *PayCtl)Delete(){
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.PayCtrl.Delete(httpWeb)
}

/*
	删除
*/
func (ctrl *PayCtl)DeleteBatch(){
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.PayCtrl.DeleteBatch(httpWeb)
}

func (ctrl *PayCtl) List(){
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.PayCtrl.List(httpWeb)
}