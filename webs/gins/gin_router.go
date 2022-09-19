package web_gin_controller

import (
	"fmt"
	_ "github.com/adminwjp/infrastructure-go/utils"
	web_gin "github.com/adminwjp/infrastructure-go/webs/gins"
	web "github.com/adminwjp/users-go/webs"
	go_playgrounds "github.com/adminwjp/users-go/webs/validators/go-playground"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type GinRouter struct {

}

func(b *GinRouter) Start(port int){
	Init()
	b.RegisterGinRouter(nil,port)
}
func (g *GinRouter)CaptchaGinRouter(router *gin.Engine )  {
	captchaController:=&CaptchaController{}
	router.Use(captchaController.Session("topgoer"))
	router.GET("/captcha", func(c *gin.Context) {
		captchaController.Captcha(c, 4)
	})
	router.GET("/captcha/verify/:value", func(c *gin.Context) {
		value := c.Param("value")
		if captchaController.CaptchaVerify(c, value) {
			c.JSON(http.StatusOK, gin.H{"status": 0, "msg": "success"})
		} else {
			c.JSON(http.StatusOK, gin.H{"status": 1, "msg": "failed"})
		}
	})
}
func (g *GinRouter)RegisterGinRouter(gin1 *gin.Engine ,port int)*gin.Engine  {
	var r *gin.Engine
	if gin1!=nil{
		r=gin1
	}else{
		r= gin.Default()
	}
	r.Use(func(c *gin.Context) {
		r.Use(web_gin.GinServerInstance.RegisterCors())
		p:=c.Request.URL.Path
		p=strings.Split(p,"?")[0]
		if strings.HasSuffix(p,".html")||
			strings.HasSuffix(p,"login")||
			strings.HasSuffix(p,"register")||
			strings.HasSuffix(p,"exists")||
			strings.HasSuffix(p,"test"){
			c.Next()
		}else{
			t,_:=c.GetQuery("token")
			if t=="" {
				c.JSON(http.StatusOK, gin.H{"status": false,"code":400,"msg":"token is empty"})
				c.Abort()
				return
			}

			r,_,err:=web.JwtInstance.ParseToken(t)
			if err!=nil||!r.Valid{
				c.JSON(http.StatusOK, gin.H{"status": false,"code":400,"msg":"token validate fail"})
				c.Abort()
				return
			}
		}

	})
	web_gin.GinServerInstance.Register(r)
	meunCtrl:=&MenuCtl{MenuCtrl:web.MenuCtrl{}}

	r.GET("/menu", meunCtrl.Load)
	// Logging to a file.
	//f, _ := os.Create("logs/gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)

	// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	//gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// 3、将我们自定义的校验方法注册到 validator中
	//if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 这里的 key 和 fn 可以不一样最终在 struct 使用的是 key
	//	v.RegisterValidation("PwdNotNull", dto.PwdNotNull)
	//}

	g.UserRouter(r)
	g.mRouter(r,"user",userCtrl)
	r.GET("/user/basic/m", userCtrl.BM)

	g.AdminRouter(r)
	r.GET("/admin1/m", adminCtrl.M)

	g.CrudRouter(r,"role",roleCtrl)
	r.GET("/role/parent",roleCtrl.Parent)
	g.mRouter(r,"role",roleCtrl)

	g.CrudRouter(r,"sms",smsCtrl)
	g.mRouter(r,"sms",smsCtrl)

	g.CrudRouter(r,"email",emailCtrl)
	g.mRouter(r,"email",emailCtrl)

	g.CrudRouter(r,"pay",payCtrl)
	g.mRouter(r,"pay",payCtrl)

	g.CrudRouter(r,"rpc",rpcCtl)
	g.mRouter(r,"rpc",rpcCtl)

	g.CrudRouter(r,"config",config1Ctl)
	g.mRouter(r,"config",config1Ctl)


	g.CaptchaGinRouter(r)


	r.GET("/5lmh", startPage)

	cfgCtrl:=&ConfigCtl{}
	r.GET("/config/get", cfgCtrl.Get)
	r.GET("/config/set/:flag", cfgCtrl.SetFlag)

	//什么玩意封装 太 厉害 怎么访问 不了  每个对应写 映射
	//r.Static("/lib","./static/lib")
	//r.Static("/admin","./static/admin")
	//r.Static("/admin","static/admin")
	//r.StaticFile("/","static")
	//r.StaticFS("/lib", http.Dir("static/lib"))
	//url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if port>0{
		err:=r.Run(":"+strconv.Itoa(port))   // 强指定端口，默认8088
		if err!=nil{
			log.Printf("start gin fail,err:%s",err.Error())
		}else{
			log.Println("start gin suc")
		}
	}
	return  r
}
type User struct {
	Username string `form:"user_name" validate:"required"`
	Tagline  string `form:"tag_line" validate:"required,lt=10"`
	Tagline2 string `form:"tag_line2" validate:"required,gt=1"`
}
func startPage(c *gin.Context) {
	//这部分应放到中间件中
	locale := c.DefaultQuery("locale", "zh")
	go_playgrounds.Init(locale)

	//这块应该放到公共验证方法中
	user := User{}
	c.ShouldBind(&user)
	fmt.Println(user)
	var httpWeb =&web_gin.GinHttpWeb{C: c}
	go_playgrounds.ValidateError(httpWeb,user)
}

func (g *GinRouter)UserRouter(r *gin.Engine)  {
	g.userRouter(r,"user",userCtrl)


	r.POST("/user/update/auth_basic", userCtrl.UpdateInfo)

	r.GET("/user/list/:page/:size", userCtrl.List)
	//r.POST("/user/list/:page/:size", userCtrl.List)

}

func (*GinRouter) ImgRouter(r *gin.Engine)  {
	//test pic  start
	//down show file
	imgCtl:=&ImgController{}
	r.GET("/img/:img",imgCtl.Get )
	//upload
	r.POST("/img/upload", imgCtl.Upload)
}

func ( g *GinRouter) AdminRouter(r *gin.Engine)  {
	g.userRouter(r,"admin",adminCtrl)


	r.POST("/admin/update_pwd", adminCtrl.UpdatePwdByOldPwd)



	//r.POST("/admin/list/:page/:size", adminCtrl.List)
	r.GET("/admin/list/:page/:size", adminCtrl.List)

}

func (*GinRouter) userRouter(r *gin.Engine,name string,ctrl IUserCtrl)  {
	r.POST("/"+name+"/login", ctrl.Login)
	r.POST("/"+name+"/phone/login", ctrl.LoginByPhone)
	r.POST("/"+name+"/email/login", ctrl.LoginByEmail)
	r.POST("/"+name+"/user_name/login", ctrl.LoginByUserName)

	r.POST("/"+name+"/register", ctrl.Register)
	r.POST("/"+name+"/phone/register", ctrl.RegisterByPhone)
	r.POST("/"+name+"/email/register", ctrl.RegisterByEmail)
	r.POST("/"+name+"/user_name/register", ctrl.RegisterByUserName)



	r.POST("/"+name+"/update/phone", ctrl.UpdatePhone)

	r.POST("/"+name+"/update/email", ctrl.UpdateEmail)
	r.POST("/"+name+"/update_email/phone", ctrl.UpdateEmailByPhone)

	//r.POST("/"+name+"/update_pwd", ctrl.UpdatePwdByOldPwd)
	r.POST("/"+name+"/update_pwd/phone", ctrl.UpdatePwdByPhone)
	r.POST("/"+name+"/update_pwd/email", ctrl.UpdatePwdByEmail)



	//r.POST("/"+name+"/list/:page/:size", adminCtrl.List)
	//r.GET("/"+name+"/list/:page/:size", adminCtrl.List)

	//r.POST("/"+name+"_log/list/:page/:size", ctrl.ListLogs)
	r.GET("/"+name+"_log/list/:page/:size", ctrl.ListLogs)
}
func (*GinRouter)mRouter(r *gin.Engine,name string,ctrl IMCtl){
	r.GET("/"+name+"/m", ctrl.M)
}
func (*GinRouter)CrudRouter(r *gin.Engine,name string,ctrl ICrudCtrl)  {

	r.POST("/"+name+"/add", ctrl.Add)
	r.POST("/"+name+"/update", ctrl.Update)
	//404 /sms/remove/1 /sms/delete/1
	//r.GET("/"+name+"/remove/:id", ctrl.Delete)
	//r.POST("/"+name+"/remove", ctrl.DeleteBatch)

	r.GET("/"+name+"/delete/:id", ctrl.Delete)
	r.POST("/"+name+"/remove_batch1", ctrl.DeleteBatch)

	//r.POST("/sms/list/:page/:size", userCtrl.GetUserLogs)
	r.GET("/"+name+"/list/:page/:size", ctrl.List)
}