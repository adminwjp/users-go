package web_http_controller

import (
	"github.com/adminwjp/infrastructure-go/webs/https"
	web "github.com/adminwjp/users-go/webs"
	"net/http"
)

type SmsCtl struct {
	web.SmsCtrl
}
func (ctrl *SmsCtl) M(writer http.ResponseWriter,request *http.Request) {
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.SmsCtrl.M(httpWeb)
}
/*
	添加
*/
func (ctrl *SmsCtl)Add(writer http.ResponseWriter,request *http.Request){
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.SmsCtrl.Add(httpWeb)
}

/*
	修改
*/
func (ctrl *SmsCtl)Update(writer http.ResponseWriter,request *http.Request){
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.SmsCtrl.Update(httpWeb)
}

/*
	删除
*/
func (ctrl *SmsCtl)Delete(writer http.ResponseWriter,request *http.Request){
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.SmsCtrl.Delete(httpWeb)
}

/*
	删除
*/
func (ctrl *SmsCtl)DeleteBatch(writer http.ResponseWriter,request *http.Request){
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.SmsCtrl.DeleteBatch(httpWeb)
}

func (ctrl *SmsCtl) List(writer http.ResponseWriter,request *http.Request){
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.SmsCtrl.List(httpWeb)
}