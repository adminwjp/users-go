package web_gin_controller

import (
	"github.com/adminwjp/infrastructure-go/webs/gins"
	dto "github.com/adminwjp/users-go/dtos"
	"github.com/adminwjp/users-go/services"
	web "github.com/adminwjp/users-go/webs"
	"github.com/gin-gonic/gin"
	"log"
)

func (ctrl *UserCtl) M(c *gin.Context) {
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.UserCtrl.M(httpWeb)
}
func (ctrl *UserCtl) BM(c *gin.Context) {
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.UserCtrl.BM(httpWeb)
}
type UserCtl struct {
	BaseUserCtl
	Service func()service.UserService
	UserCtrl *web.UserCtrl
}
/**
根据手机号、邮箱、用户名登录
*/
//swagger gin 写法不灵 写死 swagger.json
// @Summary 接口概要说明
// @Description 接口详细描述信息
// @Tags user login   //swagger API分类标签, 同一个tag为一组
// accept  */* //json from form-data xml //浏览器可处理数据类型，浏览器默认发 Accept: */*
// Produce  json  //设置返回数据的类型和编码
// @Param account path string true "account"    //url参数：（name；参数类型[query(?id=),path(/123)]；数据类型；required；参数描述）

//body
// @Param name body  UserLoginDto true "user"

// from
// @Param name query string true "pwd"
// @Param name query string false "email"
// @Param name query string false "phone"
// @Param name query string false "user_name"

//formData multipart/form-data*
// @Param name formData string true "pwd"
// @Param name formData string false "email"
// @Param name formData string false "phone"
// @Param name formData string false "user_name"

// @Success 200 {string} json {"code":200,"data":null,"msg":""}  //成功返回的数据结构， 最后是示例
// @Failure 400 {string} json {"code":200,"data":null,"msg":""}
// @Router /user/login [post]    //路由信息，一定要写上
func (ctrl *UserCtl) Login(c *gin.Context){
	var httpWeb =&gins.GinHttpWeb{C: c}
	log.Println("gin login start")
	ctrl.UserCtrl.Login(httpWeb)
	log.Println("gin login end")
}

/**
根据邮箱登录
*/
func (ctrl *UserCtl) LoginByEmail(c *gin.Context){
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.UserCtrl.LoginByEmail(httpWeb, func(input *dto.UserEmailInput) bool {
		return ctrl.UserCtrl.ValidatorUserEmailInput(httpWeb,input)
	})
}

/**
根据用户名登录
*/
func (ctrl *UserCtl) LoginByUserName(c *gin.Context){
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.UserCtrl.LoginByUserName(httpWeb, func(input *dto.UserUserNameInput) bool {
		return ctrl.UserCtrl.ValidatorUserUserNameInput(httpWeb,input)
	})
}

/**
根据手机号登录
*/
func (ctrl *UserCtl) LoginByPhone(c *gin.Context){
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.UserCtrl.LoginByPhone(httpWeb, func(input *dto.UserPhoneInput) bool {
		return ctrl.UserCtrl.ValidatorUserPhoneInput(httpWeb,input)
	})
}

/**
根据手机号、邮箱、用户名注册
*/
// @Summary register
// @Description register
// @Tags user register

// @Id register
// @Version v1

// @Accept  multipart/form-data
// @Produce  application/json

// @Param name body  UserLoginDto true "user"

// @Param name query string true "pwd"
// @Param name query string false "email"
// @Param name query string false "phone"
// @Param name query string false "user_name"

// @Param name formData string true "pwd"
// @Param name formData string false "email"
// @Param name formData string false "phone"
// @Param name formData string false "user_name"

// @Success 200 {object}
// @Failure 400 {object}
// @Router /user/register [post]

// @contact.name test
// @contact.url test
// @contact.email test
func (ctrl *UserCtl) Register(c *gin.Context)  {
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.UserCtrl.Register(httpWeb)
}



/*
	修改身份认证基本信息
*/
func (ctrl *UserCtl) UpdateInfo(c *gin.Context)  {
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.UserCtrl.UpdateInfo(httpWeb)
}

/*
	根据条件查询用户信息
*/
func (ctrl *UserCtl) List(c *gin.Context)  {
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.UserCtrl.List(httpWeb)
}

/*
	根据条件查询用户日志信息
*/
func (ctrl *UserCtl) ListLogs(c *gin.Context)  {
	var httpWeb =&gins.GinHttpWeb{C: c}
	ctrl.UserCtrl.ListLogs(httpWeb)
}



