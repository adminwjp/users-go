package web_http_controller

import (
	"github.com/adminwjp/infrastructure-go/webs"
	"github.com/adminwjp/infrastructure-go/webs/https"
	web1 "github.com/adminwjp/users-go/webs"
	"github.com/dchest/captcha"
	"net/http"
	//"github.com/gin-contrib/sessions/cookhpie"
)

type CaptchaController struct {
	
}
func (c1 *CaptchaController)Captcha1(writer http.ResponseWriter,request *http.Request){
	c1.Captcha( writer, request,4)

}
func (c1 *CaptchaController)Captcha(writer http.ResponseWriter,request *http.Request,length ...int) {
	l := captcha.DefaultLen
	w, h := 107, 36
	if len(length) == 1 {
		l = length[0]
	}
	if len(length) == 2 {
		w = length[1]
	}
	if len(length) == 3 {
		h = length[2]
	}
	captchaId := captcha.NewLen(l)
	cookies:=append(request.Response.Cookies(),&http.Cookie{
		Name: "captcha",
		Value: captchaId,
	})
	var str=""
	for i := 0; i <len(cookies) ; i++ {
		str+=cookies[i].String()+";"
	}
	writer.Header().Set("Set-Cookie",str)
	_ = web1.CaptchaServe(writer, request, captchaId, ".png", "zh", false, w, h)
}
func (c1 *CaptchaController)CaptchaVerify1(writer http.ResponseWriter,request *http.Request){
	value := request.URL.Query().Get("value")
	if c1.CaptchaVerify( writer,request,value) {
		https.OutputXmlOrJson(writer,webs.Json,map[string]interface{}{"status": 0, "msg": "success"})
	} else {
		https.OutputXmlOrJson(writer,webs.Json,map[string]interface{}{"status": 1, "msg": "failed"})
	}
}
func (c1 *CaptchaController)CaptchaVerify(writer http.ResponseWriter,request *http.Request,code string) bool {
	cookies:=request.Response.Cookies()
	e,f:=false,false
	for i := 0; i <len(cookies) ; i++ {
		if  cookies[i].Name=="captcha" {
			f=true

			if captcha.VerifyString(cookies[i].Value, code) {
				e = true
			}
			cookies[i]=nil
			break
		}
	}
	if f{
		var str=""
		for i := 0; i <len(cookies) ; i++ {
			if cookies[i]==nil{
				continue
			}
			str+=cookies[i].String()+";"
		}
		writer.Header().Set("Set-Cookie",str)
	}
	return e
}
