package web_gin_controller

import (
	"github.com/adminwjp/infrastructure-go/webs/gins"
	web "github.com/adminwjp/users-go/webs"
	"github.com/gin-gonic/gin"
)

type ConfigCtl struct {
	web.ConfigCtrl
}
func (ctrl *ConfigCtl) Get(c *gin.Context)  {
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.ConfigCtrl.Get(httpWeb)
}
func (ctrl *ConfigCtl) SetFlag(c *gin.Context)  {
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.ConfigCtrl.SetFlag(httpWeb)
}