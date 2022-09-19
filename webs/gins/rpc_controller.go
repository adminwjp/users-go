package web_gin_controller

import (
	"github.com/adminwjp/infrastructure-go/webs/gins"
	web "github.com/adminwjp/users-go/webs"
	"github.com/gin-gonic/gin"
)

type RpcCtl struct {
	web.RpcCtrl
}
func (ctrl *RpcCtl) M(c *gin.Context) {
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.RpcCtrl.M(httpWeb)
}
/*
	添加
*/
func (ctrl *RpcCtl)Add(c *gin.Context){
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.RpcCtrl.Add(httpWeb)
}

/*
	修改
*/
func (ctrl *RpcCtl)Update(c *gin.Context){
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.RpcCtrl.Update(httpWeb)
}

/*
	删除
*/
func (ctrl *RpcCtl)Delete(c *gin.Context){
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.RpcCtrl.Delete(httpWeb)
}

/*
	删除
*/
func (ctrl *RpcCtl)DeleteBatch(c *gin.Context){
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.RpcCtrl.DeleteBatch(httpWeb)
}

func (ctrl *RpcCtl) List(c *gin.Context){
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.RpcCtrl.List(httpWeb)
}
