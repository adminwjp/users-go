package web_http_controller

import (
	"github.com/adminwjp/infrastructure-go/webs/https"
	web "github.com/adminwjp/users-go/webs"
	"net/http"
)

type MenuCtl struct {
	web.MenuCtrl
}
func (ctrl *MenuCtl)Load(writer http.ResponseWriter,request *http.Request)  {
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.MenuCtrl.Load(httpWeb)
}