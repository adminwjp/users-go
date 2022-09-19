package web_http_controller

import (
	"github.com/adminwjp/infrastructure-go/webs/https"
	web "github.com/adminwjp/users-go/webs"
	"net/http"
)

type PayCtl struct {
	web.PayCtrl
}

func (ctrl *PayCtl) M(writer http.ResponseWriter,request *http.Request) {
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.PayCtrl.M(httpWeb)
}
/*
	添加
*/
func (ctrl *PayCtl)Add(writer http.ResponseWriter,request *http.Request){
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.PayCtrl.Add(httpWeb)
}

/*
	修改
*/
func (ctrl *PayCtl)Update(writer http.ResponseWriter,request *http.Request){
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.PayCtrl.Update(httpWeb)
}

/*
	删除
*/
func (ctrl *PayCtl)Delete(writer http.ResponseWriter,request *http.Request){
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.PayCtrl.Delete(httpWeb)
}

/*
	删除
*/
func (ctrl *PayCtl)DeleteBatch(writer http.ResponseWriter,request *http.Request){
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.PayCtrl.DeleteBatch(httpWeb)
}

func (ctrl *PayCtl) List(writer http.ResponseWriter,request *http.Request){
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.PayCtrl.List(httpWeb)
}