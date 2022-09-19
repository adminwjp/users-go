package web_http_controller

import (
	"github.com/adminwjp/infrastructure-go/webs/https"
	web "github.com/adminwjp/users-go/webs"
	"net/http"
)

type RpcCtl struct {
	web.RpcCtrl
}
func (ctrl *RpcCtl) M(writer http.ResponseWriter,request *http.Request) {
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.RpcCtrl.M(httpWeb)
}
/*
	添加
*/
func (ctrl *RpcCtl)Add(writer http.ResponseWriter,request *http.Request){
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.RpcCtrl.Add(httpWeb)
}

/*
	修改
*/
func (ctrl *RpcCtl)Update(writer http.ResponseWriter,request *http.Request){
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.RpcCtrl.Update(httpWeb)
}

/*
	删除
*/
func (ctrl *RpcCtl)Delete(writer http.ResponseWriter,request *http.Request){
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.RpcCtrl.Delete(httpWeb)
}

/*
	删除
*/
func (ctrl *RpcCtl)DeleteBatch(writer http.ResponseWriter,request *http.Request){
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.RpcCtrl.DeleteBatch(httpWeb)
}

func (ctrl *RpcCtl) List(writer http.ResponseWriter,request *http.Request){
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.RpcCtrl.List(httpWeb)
}