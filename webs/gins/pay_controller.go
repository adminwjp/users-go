package web_gin_controller

import (
	"github.com/adminwjp/infrastructure-go/webs/gins"
	web "github.com/adminwjp/users-go/webs"
	"github.com/gin-gonic/gin"
)

type PayCtl struct {
	web.PayCtrl
}

func (ctrl *PayCtl) M(c *gin.Context) {
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.PayCtrl.M(httpWeb)
}
/*
	添加
*/
func (ctrl *PayCtl)Add(c *gin.Context){
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.PayCtrl.Add(httpWeb)
}

/*
	修改
*/
func (ctrl *PayCtl)Update(c *gin.Context){
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.PayCtrl.Update(httpWeb)
}

/*
	删除
*/
func (ctrl *PayCtl)Delete(c *gin.Context){
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.PayCtrl.Delete(httpWeb)
}

/*
	删除
*/
func (ctrl *PayCtl)DeleteBatch(c *gin.Context){
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.PayCtrl.DeleteBatch(httpWeb)
}

func (ctrl *PayCtl) List(c *gin.Context){
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.PayCtrl.List(httpWeb)
}