package web_gin_controller

import (
	web1 "github.com/adminwjp/users-go/webs"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/dchest/captcha"
	"github.com/gin-contrib/sessions"
	//"github.com/gin-contrib/sessions/cookhpie"
)

type CaptchaController struct {
	
}
// 中间件，处理session
func (c *CaptchaController)Session(keyPairs string) gin.HandlerFunc {
	store := c.SessionConfig()
	return sessions.Sessions(keyPairs, store)
}
func (*CaptchaController)SessionConfig() sessions.Store {
	sessionMaxAge := 3600
	sessionSecret := "topgoer"
	var store sessions.Store
	store = cookie.NewStore([]byte(sessionSecret))
	store.Options(sessions.Options{
		MaxAge: sessionMaxAge, //seconds
		Path:   "/",
	})
	return store
}

func (c1 *CaptchaController)Captcha(c *gin.Context, length ...int) {
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
	session := sessions.Default(c)
	session.Set("captcha", captchaId)
	_ = session.Save()
	_ = web1.CaptchaServe(c.Writer, c.Request, captchaId, ".png", "zh", false, w, h)
}
func (c1 *CaptchaController)CaptchaVerify(c *gin.Context, code string) bool {
	session := sessions.Default(c)
	if captchaId := session.Get("captcha"); captchaId != nil {
		session.Delete("captcha")
		_ = session.Save()
		if captcha.VerifyString(captchaId.(string), code) {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}
