package web_gin_controller

import (
	"github.com/adminwjp/infrastructure-go/webs/gins"
	web "github.com/adminwjp/users-go/webs"
	"github.com/gin-gonic/gin"
)

type SmsCtl struct {
	web.SmsCtrl
}
func (ctrl *SmsCtl) M(c *gin.Context) {
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.SmsCtrl.M(httpWeb)
}
/*
	添加
*/
func (ctrl *SmsCtl)Add(c *gin.Context){
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.SmsCtrl.Add(httpWeb)
}

/*
	修改
*/
func (ctrl *SmsCtl)Update(c *gin.Context){
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.SmsCtrl.Update(httpWeb)
}

/*
	删除
*/
func (ctrl *SmsCtl)Delete(c *gin.Context){
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.SmsCtrl.Delete(httpWeb)
}

/*
	删除
*/
func (ctrl *SmsCtl)DeleteBatch(c *gin.Context){
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.SmsCtrl.DeleteBatch(httpWeb)
}

func (ctrl *SmsCtl) List(c *gin.Context){
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.SmsCtrl.List(httpWeb)
}