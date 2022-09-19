package web_http_controller

import (
	"github.com/adminwjp/infrastructure-go/webs/https"
	web "github.com/adminwjp/users-go/webs"
	"net/http"
)

type ConfigCtl struct {
	web.ConfigCtrl
}
func (ctrl *ConfigCtl) Get(writer http.ResponseWriter,request *http.Request)  {
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.ConfigCtrl.Get(httpWeb)
}
func (ctrl *ConfigCtl) SetFlag(writer http.ResponseWriter,request *http.Request)  {
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.ConfigCtrl.SetFlag(httpWeb)
}