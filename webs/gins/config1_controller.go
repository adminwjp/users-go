package web_gin_controller

import (
	"github.com/adminwjp/infrastructure-go/webs/gins"
	web "github.com/adminwjp/users-go/webs"
	"github.com/gin-gonic/gin"
)

type Config1Ctl struct {
	web.Config1Ctrl
}
func (ctrl *Config1Ctl) M(c *gin.Context) {
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.Config1Ctrl.M(httpWeb)
}
/*
	添加
*/
func (ctrl *Config1Ctl)Add(c *gin.Context){
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.Config1Ctrl.Add(httpWeb)
}

/*
	修改
*/
func (ctrl *Config1Ctl)Update(c *gin.Context){
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.Config1Ctrl.Update(httpWeb)
}

/*
	删除
*/
func (ctrl *Config1Ctl)Delete(c *gin.Context){
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.Config1Ctrl.Delete(httpWeb)
}

/*
	删除
*/
func (ctrl *Config1Ctl)DeleteBatch(c *gin.Context){
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.Config1Ctrl.DeleteBatch(httpWeb)
}

func (ctrl *Config1Ctl) List(c *gin.Context){
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.Config1Ctrl.List(httpWeb)
}

