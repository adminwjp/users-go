package web_gin_controller

import (
	"github.com/adminwjp/infrastructure-go/webs/gins"
	web "github.com/adminwjp/users-go/webs"
	"github.com/gin-gonic/gin"
)

type EmailCtl struct {
	web.EmailCtrl
}
func (ctrl *EmailCtl) M(c *gin.Context) {
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.EmailCtrl.M(httpWeb)
}
/*
	添加
*/
func (ctrl *EmailCtl)Add(c *gin.Context){
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.EmailCtrl.Add(httpWeb)
}

/*
	修改
*/
func (ctrl *EmailCtl)Update(c *gin.Context){
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.EmailCtrl.Update(httpWeb)
}

/*
	删除
*/
func (ctrl *EmailCtl)Delete(c *gin.Context){
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.EmailCtrl.Delete(httpWeb)
}

/*
	删除
*/
func (ctrl *EmailCtl)DeleteBatch(c *gin.Context){
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.EmailCtrl.DeleteBatch(httpWeb)
}

func (ctrl *EmailCtl) List(c *gin.Context){
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.EmailCtrl.List(httpWeb)
}