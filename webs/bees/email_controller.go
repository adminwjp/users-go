package web_bee_controller

import (
	"github.com/adminwjp/infrastructure-go/webs/bees"
	web "github.com/adminwjp/users-go/webs"
	bee "github.com/beego/beego/v2/server/web"
)

type EmailCtl struct {
	web.EmailCtrl
	bee.Controller
}
func (ctrl *EmailCtl) M() {
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.EmailCtrl.M(httpWeb)
}
/*
	添加
*/
func (ctrl *EmailCtl)Add(){
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.EmailCtrl.Add(httpWeb)
}

/*
	修改
*/
func (ctrl *EmailCtl)Update(){
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.EmailCtrl.Update(httpWeb)
}

/*
	删除
*/
func (ctrl *EmailCtl)Delete(){
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.EmailCtrl.Delete(httpWeb)
}

/*
	删除
*/
func (ctrl *EmailCtl)DeleteBatch(){
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.EmailCtrl.DeleteBatch(httpWeb)
}

func (ctrl *EmailCtl) List(){
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.EmailCtrl.List(httpWeb)
}