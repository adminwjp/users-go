package web_http_controller

import (
	"github.com/adminwjp/infrastructure-go/webs/https"
	web "github.com/adminwjp/users-go/webs"
	"net/http"
)

type Config1Ctl struct {
	web.Config1Ctrl
}
func (ctrl *Config1Ctl) M(writer http.ResponseWriter,request *http.Request) {
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.Config1Ctrl.M(httpWeb)
}
/*
	添加
*/
func (ctrl *Config1Ctl)Add(writer http.ResponseWriter,request *http.Request){
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.Config1Ctrl.Add(httpWeb)
}

/*
	修改
*/
func (ctrl *Config1Ctl)Update(writer http.ResponseWriter,request *http.Request){
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.Config1Ctrl.Update(httpWeb)
}

/*
	删除
*/
func (ctrl *Config1Ctl)Delete(writer http.ResponseWriter,request *http.Request){
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.Config1Ctrl.Delete(httpWeb)
}

/*
	删除
*/
func (ctrl *Config1Ctl)DeleteBatch(writer http.ResponseWriter,request *http.Request){
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.Config1Ctrl.DeleteBatch(httpWeb)
}

func (ctrl *Config1Ctl) List(writer http.ResponseWriter,request *http.Request){
	var httpWeb =&https.HttpWeb{Writer: writer,Request: request}
	ctrl.Config1Ctrl.List(httpWeb)
}
