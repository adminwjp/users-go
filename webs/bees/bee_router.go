package web_bee_controller

import (
	web1 "github.com/adminwjp/users-go/webs"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"strings"
)

type BeeRouter struct {
	
}
func(b *BeeRouter) Start(port int){
	Init()
	web.InsertFilterChain("/*", func(next1 web.FilterFunc) web.FilterFunc {
		return func(next *context.Context) {
			p := next.Request.URL.Path
			p = strings.Split(p, "?")[0]
			if strings.HasSuffix(p,".html")||
				strings.HasSuffix(p, "login") ||
				strings.HasSuffix(p, "register") ||
				strings.HasSuffix(p, "exists")||
				strings.HasSuffix(p,"test") {
				next1(next)
			} else {
				t := next.Input.Query("token")
				if t == "" {
					next.Output.JSON(map[string]interface{}{"status": false, "code": 400, "msg": "token is empty"}, false, false)
					//next.Abort(200)
					return
				}

				r, _, err := web1.JwtInstance.ParseToken(t)
				if err != nil || !r.Valid {
					next.Output.JSON(map[string]interface{}{"status": false, "code": 400, "msg": "token validate fail"}, false, false)
					//next.Abort(200)
				}
			}
	}})
	b.RegisterRouter()
	b.RegisterRouter1()
}
func(*BeeRouter) RegisterRouter()  {
	web.Get("/test1",func(ctx *context.Context){
		ctx.Output.Body([]byte("hello world"))
	})

}
func (b *BeeRouter)CaptchaGinRouter( )  {
	captchaController:=&CaptchaController{}
	web.Router("/captcha", captchaController,"Get:Captcha1")
	web.Router("/captcha/verify/:value", captchaController,"Get:CaptchaVerify1")
}
func (g *BeeRouter)RegisterRouter1() {
	meunCtrl:=&MenuCtl{MenuCtrl:web1.MenuCtrl{}}
	web.Router("/menu", meunCtrl,"Get:Load")





	g.UserRouter()
	g.mRouter("user",userCtrl,userCtrl)
	web.Router("/user/basic/m", userCtrl,"Get:BM")

	g.AdminRouter()
	web.Router("/admin1/m", adminCtrl,"Get:M")

	g.CrudRouter("role",roleCtrl,roleCtrl)
	web.Router("/role/parent",roleCtrl,"Get:Parent")
	g.mRouter("role",roleCtrl,roleCtrl)

	g.CrudRouter("sms",smsCtrl,smsCtrl)
	g.mRouter("sms",smsCtrl,smsCtrl)

	g.CrudRouter("email",emailCtrl,emailCtrl)
	g.mRouter("email",emailCtrl,emailCtrl)

	g.CrudRouter("pay",payCtrl,payCtrl)
	g.mRouter("pay",payCtrl,payCtrl)

	g.CrudRouter("rpc",rpcCtl,rpcCtl)
	g.mRouter("rpc",rpcCtl,rpcCtl)

	g.CrudRouter("config",config1Ctl,config1Ctl)
	g.mRouter("config",config1Ctl,config1Ctl)
	g.CaptchaGinRouter()




	cfgCtrl:=&ConfigCtl{}
	web.Router("/config/get", cfgCtrl,"Get:Get")
	web.Router("/config/set/:flag", cfgCtrl,"Get:SetFlag")

}

func (g *BeeRouter)UserRouter()  {
	g.userRouter("user",userCtrl,userCtrl)


	web.Router("/user/update/auth_basic", userCtrl,"Post:UpdateInfo")

	web.Router("/user/list/:page/:size", userCtrl,"Get:List")
	web.Router("/user/list/:page/:size", userCtrl,"Post:List")

}

func (*BeeRouter) ImgRouter()  {

	//test pic  start
	//down show file
	imgCtl:=&ImgController{}
	web.Router("/img/:img",imgCtl ,"Get:Get")
	//upload
	web.Router("/img/upload", imgCtl,"Post:Upload")
}

func ( g *BeeRouter) AdminRouter()  {


	g.userRouter("admin",adminCtrl,adminCtrl)


	web.Router("/admin/update_pwd", adminCtrl,"Post:UpdatePwdByOldPwd")



	web.Router("/admin/list/:page/:size", adminCtrl,"Post:List")
	//web.Router("/admin/list/:page/:size", adminCtrl,"Post:List")

}

func (*BeeRouter) userRouter(name string,ctrl1 web.ControllerInterface,ctrl IUserCtrl)  {


	web.Router("/"+name+"/login", ctrl1,"Post:Login")
	web.Router("/"+name+"/phone/login", ctrl1,"Post:LoginByPhone")
	web.Router("/"+name+"/email/login", ctrl1,"Post:LoginByEmail")
	web.Router("/"+name+"/user_name/login", ctrl1,"Post:LoginByUserName")

	web.Router("/"+name+"/register", ctrl1,"Post:Update")
	web.Router("/"+name+"/phone/register", ctrl1,"Post:Update")
	web.Router("/"+name+"/email/register", ctrl1,"Post:Update")
	web.Router("/"+name+"/user_name/register", ctrl1,"Post:Update")



	web.Router("/"+name+"/update/phone", ctrl1,"Post:UpdatePhone")

	web.Router("/"+name+"/update/email", ctrl1,"Post:UpdateEmail")
	web.Router("/"+name+"/update_email/phone", ctrl1,"Post:UpdateEmailByPhone")

	//r.POST("/"+name+"/update_pwd", ctrl.UpdatePwdByOldPwd)
	web.Router("/"+name+"/update_pwd/phone", ctrl1,"Post:UpdatePwdByPhone")
	web.Router("/"+name+"/update_pwd/email", ctrl1,"Post:UpdatePwdByEmail")



	//r.POST("/admin/list/:page/:size", adminCtrl.GetUsers)
	//r.POST("/admin/list/:page/:size", adminCtrl.GetUsers)

	web.Router("/"+name+"_log/list/:page/:size", ctrl1,"Post:ListLogs")
	web.Router("/"+name+"_log/list/:page/:size", ctrl1,"Get:ListLogs")
}
func (*BeeRouter)mRouter(name string,ctrl1 web.ControllerInterface,ctrl IMCtl){
	web.Router("/"+name+"/m", ctrl1,"Get:M")
}
func (*BeeRouter)CrudRouter(name string,ctrl1 web.ControllerInterface,ctrl ICrudCtrl)  {

	web.Router("/"+name+"/add", ctrl1,"Post:Add")
	web.Router("/"+name+"/update", ctrl1,"Post:Update")
	web.Router("/"+name+"/delete", ctrl1,"Get:Delete")
	web.Router("/"+name+"/remove_batch1", ctrl1,"Post:DeleteBatch")
	web.Router("/"+name+"/list/:page/:size", ctrl1,"Get:List")
}