package web_http_controller

import (
	_ "github.com/adminwjp/infrastructure-go/utils"
	"github.com/adminwjp/infrastructure-go/webs/https"
	web "github.com/adminwjp/users-go/webs"
	"net/http"
	"strconv"
)

type HttpRouter struct {

}
func(b *HttpRouter) Start(port int){
	Init()
	b.RegisterRouter(port)
}
func (g *HttpRouter)CaptchaRouter()  {
	captchaController:=&CaptchaController{}
	http.HandleFunc("/captcha",captchaController.Captcha1)
	http.HandleFunc("/captcha/verify/:value",captchaController.CaptchaVerify1)
}
func (g *HttpRouter)RegisterRouter(port int)  {
	https.HttpServerInstance.HTTPHandlerFuncInterceptor(func(writer http.ResponseWriter, request *http.Request) {

	})
	//http.Serve()
	https.HttpServerInstance.Start(func() {

	},strconv.Itoa(port),nil,nil)

	meunCtrl:=&MenuCtl{MenuCtrl:web.MenuCtrl{}}
	http.HandleFunc("/menu", meunCtrl.Load)


	g.UserRouter()
	g.mRouter("user",userCtrl)
	http.HandleFunc("/user/basic/m", userCtrl.BM)

	g.AdminRouter()
	http.HandleFunc("/admin1/m", adminCtrl.M)

	g.CrudRouter("role",roleCtrl)
	http.HandleFunc("/role/parent",roleCtrl.Parent)
	g.mRouter("role",roleCtrl)

	g.CrudRouter("sms",smsCtrl)
	g.mRouter("sms",smsCtrl)

	g.CrudRouter("email",emailCtrl)
	g.mRouter("email",emailCtrl)

	g.CrudRouter("pay",payCtrl)
	g.mRouter("pay",payCtrl)

	g.CrudRouter("rpc",rpcCtl)
	g.mRouter("rpc",rpcCtl)

	g.CrudRouter("config",config1Ctl)
	g.mRouter("config",config1Ctl)

	g.CaptchaRouter()



	cfgCtrl:=&ConfigCtl{}
	http.HandleFunc("/config/get", cfgCtrl.Get)
	http.HandleFunc("/config/set/:flag", cfgCtrl.SetFlag)

	//什么玩意封装 太 厉害 怎么访问 不了  每个对应写 映射
	http.Handle("/lib",http.FileServer(http.Dir("./static/lib")))
	http.Handle("/admin",http.FileServer(http.Dir("./static/admin")))
	http.Handle("/swagger/*any", http.FileServer(http.Dir("./docs")))


}


func (g *HttpRouter)UserRouter()  {
	g.userRouter("user",userCtrl)


	http.HandleFunc("/user/update/auth_basic", userCtrl.UpdateInfo)

	http.HandleFunc("/user/list/:page/:size", userCtrl.List)
	http.HandleFunc("/user/list/:page/:size", userCtrl.List)

}

func (*HttpRouter) ImgRouter()  {
	//test pic  start
	//down show file
	imgCtl:=&ImgController{}
	http.HandleFunc("/img/:img",imgCtl.Get )
	//upload
	http.HandleFunc("/img/upload", imgCtl.Upload)
}

func ( g *HttpRouter) AdminRouter()  {
	g.userRouter("admin",adminCtrl)


	http.HandleFunc("/admin/update_pwd", adminCtrl.UpdatePwdByOldPwd)



	http.HandleFunc("/admin/list/:page/:size", adminCtrl.List)
	//http.HandleFunc("/admin/list/:page/:size", adminCtrl.List)

}

func (*HttpRouter) userRouter(name string,ctrl IUserCtrl)  {
	http.HandleFunc("/"+name+"/login", ctrl.Login)
	http.HandleFunc("/"+name+"/phone/login", ctrl.LoginByPhone)
	http.HandleFunc("/"+name+"/email/login", ctrl.LoginByEmail)
	http.HandleFunc("/"+name+"/user_name/login", ctrl.LoginByUserName)

	http.HandleFunc("/"+name+"/register", ctrl.Register)
	http.HandleFunc("/"+name+"/phone/register", ctrl.RegisterByPhone)
	http.HandleFunc("/"+name+"/email/register", ctrl.RegisterByEmail)
	http.HandleFunc("/"+name+"/user_name/register", ctrl.RegisterByUserName)



	http.HandleFunc("/"+name+"/update/phone", ctrl.UpdatePhone)

	http.HandleFunc("/"+name+"/update/email", ctrl.UpdateEmail)
	http.HandleFunc("/"+name+"/update_email/phone", ctrl.UpdateEmailByPhone)

	//http.HandleFunc("/"+name+"/update_pwd", ctrl.UpdatePwdByOldPwd)
	http.HandleFunc("/"+name+"/update_pwd/phone", ctrl.UpdatePwdByPhone)
	http.HandleFunc("/"+name+"/update_pwd/email", ctrl.UpdatePwdByEmail)



	//http.HandleFunc("/admin/list/:page/:size", adminCtrl.GetUsers)
	//http.HandleFunc("/admin/list/:page/:size", adminCtrl.GetUsers)

	http.HandleFunc("/"+name+"_log/list/:page/:size", ctrl.ListLogs)
	http.HandleFunc("/"+name+"_log/list/:page/:size", ctrl.ListLogs)
}
func (*HttpRouter)mRouter(name string,ctrl IMCtl){
	http.HandleFunc("/"+name+"/m", ctrl.M)
}
func (*HttpRouter)CrudRouter(name string,ctrl ICrudCtrl)  {

	http.HandleFunc("/"+name+"/add", ctrl.Add)
	http.HandleFunc("/"+name+"/update", ctrl.Update)
	//404 /sms/remove/1 /sms/delete/1
	//http.HandleFunc("/"+name+"/remove/:id", ctrl.Delete)
	//http.HandleFunc("/"+name+"/remove", ctrl.DeleteBatch)

	http.HandleFunc("/"+name+"/delete", ctrl.Delete)
	http.HandleFunc("/"+name+"/remove_batch1", ctrl.DeleteBatch)

	//r.POST("/sms/list/:page/:size", userCtrl.GetUserLogs)
	http.HandleFunc("/"+name+"/list/:page/:size", ctrl.List)
}