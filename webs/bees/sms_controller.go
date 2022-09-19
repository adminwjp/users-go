package web_bee_controller

import (
	"github.com/adminwjp/infrastructure-go/webs/bees"
	web "github.com/adminwjp/users-go/webs"
	bee "github.com/beego/beego/v2/server/web"
)

type SmsCtl struct {
	bee.Controller
	web.SmsCtrl
}
func (ctrl *SmsCtl) M() {
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.SmsCtrl.M(httpWeb)
}
/*
	添加
*/
func (ctrl *SmsCtl)Add(){
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.SmsCtrl.Add(httpWeb)
}

/*
	修改
*/
func (ctrl *SmsCtl)Update(){
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.SmsCtrl.Update(httpWeb)
}

/*
	删除
*/
func (ctrl *SmsCtl)Delete(){
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.SmsCtrl.Delete(httpWeb)
}

/*
	删除
*/
func (ctrl *SmsCtl)DeleteBatch(){
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.SmsCtrl.DeleteBatch(httpWeb)
}

func (ctrl *SmsCtl) List(){
	var httpWeb =&bees.BeeHttpWeb{C: ctrl.Ctx}
	ctrl.SmsCtrl.List(httpWeb)
}