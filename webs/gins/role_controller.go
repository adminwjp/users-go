package web_gin_controller

import (
	"github.com/adminwjp/infrastructure-go/webs/gins"
	web "github.com/adminwjp/users-go/webs"
	"github.com/gin-gonic/gin"
)

type RoleCtl struct {
	web.RoleCtrl
}
func (ctrl *RoleCtl) M(c *gin.Context) {
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.RoleCtrl.M(httpWeb)
}
/*
	添加
*/
func (ctrl *RoleCtl)Add(c *gin.Context){
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.RoleCtrl.Add(httpWeb)
}

/*
	修改
*/
func (ctrl *RoleCtl)Update(c *gin.Context){
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.RoleCtrl.Update(httpWeb)
}

/*
	删除
*/
func (ctrl *RoleCtl)Delete(c *gin.Context){
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.RoleCtrl.Delete(httpWeb)
}

/*
	删除
*/
func (ctrl *RoleCtl)DeleteBatch(c *gin.Context){
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.RoleCtrl.DeleteBatch(httpWeb)
}

func (ctrl *RoleCtl) List(c *gin.Context){
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.RoleCtrl.List(httpWeb)
}

func (ctrl *RoleCtl) Parent(c *gin.Context){
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.RoleCtrl.Parent(httpWeb)
}