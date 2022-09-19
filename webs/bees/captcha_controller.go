package web_bee_controller

import (
	web1 "github.com/adminwjp/users-go/webs"
	bee "github.com/beego/beego/v2/server/web"
	"github.com/dchest/captcha"
	//"github.com/gin-contrib/sessions/cookhpie"
)

type CaptchaController struct {
	bee.Controller
}

func (c1 *CaptchaController)Captcha1(){
	c1.Captcha( 4)

}
func (c1 *CaptchaController)Captcha(length ...int) {
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
	c1.SetSession("captcha", captchaId)
	_ = web1.CaptchaServe(c1.Ctx.ResponseWriter, c1.Ctx.Request, captchaId, ".png", "zh", false, w, h)
}
func (c1 *CaptchaController)CaptchaVerify1(){
	value := c1.GetString("value")
	if c1.CaptchaVerify( value) {
		c1.Data["json"]=map[string]interface{}{"status": 0, "msg": "success"}
		c1.ServeJSON()
	} else {
		c1.Data["json"]=map[string]interface{}{"status": 1, "msg": "failed"}
		c1.ServeJSON()
	}
}
func (c1 *CaptchaController)CaptchaVerify(code string) bool {
	if captchaId := c1.GetSession("captcha"); captchaId != nil {
		c1.DelSession("captcha")
		if captcha.VerifyString(captchaId.(string), code) {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}
