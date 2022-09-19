package web_http_controller

import (
	"github.com/adminwjp/infrastructure-go/webs/https"
	web "github.com/adminwjp/users-go/webs"
	"net/http"
)

type RoleCtl struct {
	web.RoleCtrl
}
func (ctrl *RoleCtl) M(writer http.ResponseWriter,request *http.Request) {
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.RoleCtrl.M(httpWeb)
}
/*
	添加
*/
func (ctrl *RoleCtl)Add(writer http.ResponseWriter,request *http.Request){
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.RoleCtrl.Add(httpWeb)
}

/*
	修改
*/
func (ctrl *RoleCtl)Update(writer http.ResponseWriter,request *http.Request){
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.RoleCtrl.Update(httpWeb)
}

/*
	删除
*/
func (ctrl *RoleCtl)Delete(writer http.ResponseWriter,request *http.Request){
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.RoleCtrl.Delete(httpWeb)
}

/*
	删除
*/
func (ctrl *RoleCtl)DeleteBatch(writer http.ResponseWriter,request *http.Request){
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.RoleCtrl.DeleteBatch(httpWeb)
}

func (ctrl *RoleCtl) List(writer http.ResponseWriter,request *http.Request){
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.RoleCtrl.List(httpWeb)
}

func (ctrl *RoleCtl) Parent(writer http.ResponseWriter,request *http.Request){
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.RoleCtrl.Parent(httpWeb)
}