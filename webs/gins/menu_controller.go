package web_gin_controller

import (
	"github.com/adminwjp/infrastructure-go/webs/gins"
	web "github.com/adminwjp/users-go/webs"
	"github.com/gin-gonic/gin"
)

type MenuCtl struct {
	web.MenuCtrl
}
func (ctrl *MenuCtl)Load(c *gin.Context) {
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.MenuCtrl.Load(httpWeb)
}