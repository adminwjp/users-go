package web_http_controller

import (
	"github.com/adminwjp/infrastructure-go/webs/https"
	web "github.com/adminwjp/users-go/webs"
	"net/http"
)

type EmailCtl struct {
	web.EmailCtrl
}
func (ctrl *EmailCtl) M(writer http.ResponseWriter,request *http.Request) {
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.EmailCtrl.M(httpWeb)
}
/*
	添加
*/
func (ctrl *EmailCtl)Add(writer http.ResponseWriter,request *http.Request){
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.EmailCtrl.Add(httpWeb)
}

/*
	修改
*/
func (ctrl *EmailCtl)Update(writer http.ResponseWriter,request *http.Request){
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.EmailCtrl.Update(httpWeb)
}

/*
	删除
*/
func (ctrl *EmailCtl)Delete(writer http.ResponseWriter,request *http.Request){
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.EmailCtrl.Delete(httpWeb)
}

/*
	删除
*/
func (ctrl *EmailCtl)DeleteBatch(writer http.ResponseWriter,request *http.Request){
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.EmailCtrl.DeleteBatch(httpWeb)
}

func (ctrl *EmailCtl) List(writer http.ResponseWriter,request *http.Request){
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.EmailCtrl.List(httpWeb)
}