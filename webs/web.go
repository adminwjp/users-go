package web

import (
	dtos "github.com/adminwjp/infrastructure-go/dtos"
	jwts "github.com/adminwjp/infrastructure-go/jwts"
	"github.com/adminwjp/infrastructure-go/webs"
	"github.com/adminwjp/users-go/daos"
	"github.com/adminwjp/users-go/models"
	"github.com/adminwjp/users-go/services"
	"strconv"
	"sync"
	"time"
)
func  List(httpWeb webs.HttpWeb,data interface{},page int,size int,records int64)  {
	var t int=0
	if records>0{
		if records%int64(size)==0{
			t=int(records/int64(size))
		}else{
			t=int(records/int64(size))+1
		}
	}
	var r =dtos.ResponsePageListDto{Status: true, Code: 200, Msg: "success",List:data,
		Page: page,Size: size,Records: records,Total: t}
	httpWeb.Response(200,r)
}

var JwtInstance=&jwts.JwtHelper{}
func AdminLogin(dto models.AdminModel,ip string) dtos.ResponseTokenAndRefreshTokenDto {
	n:=time.Now()
	dto1 :=&jwts.JwtDto{
		UserId: strconv.FormatInt(dto.Id,10),
		UserName: dto.UserName,
		Phone: dto.Phone,
		Email: dto.Email,
		Ip: ip,
		CreateAt: n.Unix(),
		ExpiresAt: n.Add(time.Hour*24).Unix(),
		Issuer: "test",
	}
	dto1.Subject=dto1.UserId
	t,r,_:=JwtInstance.CreateTokenAndRefreshToken(*dto1)
	res1:= dtos.ResponseTokenAndRefreshTokenDto{Status: true,
		Code:   200,Msg:    "login success",Token: t,Expired: 24*3600,
		RefreshToken: r,RefreshExpired: 24*3600,
	}
	return res1
}
func UserLogin(dto models.UserModel,ip string) dtos.ResponseTokenAndRefreshTokenDto {
	n:=time.Now()
	dto1 :=&jwts.JwtDto{
		UserId: strconv.FormatInt(dto.Id,10),
		UserName: dto.UserName,
		Phone: dto.Phone,
		Email: dto.Email,
		Ip: ip,
		CreateAt: n.Unix(),
		ExpiresAt: n.Add(time.Hour*24).Unix(),
		Issuer: "test",
	}
	dto1.Subject=dto1.UserId
	t,r,_:=JwtInstance.CreateTokenAndRefreshToken(*dto1)
	res1:= dtos.ResponseTokenAndRefreshTokenDto{Status: true,
		Code:   200,Msg:    "login success",Token: t,Expired: 24*3600,
		RefreshToken: r,RefreshExpired: 24*3600,
	}
	return res1
}

//pool 10
var Trans =make(map[string]daos.TranDao)
var TranStatus =make(map[string]bool)

var AdminServices=make(map[int]service.AdminService)
var AdminTranStatus =make(map[string]bool)

var UserServices=make(map[int]service.UserService)
var UserTranStatus =make(map[string]bool)

var RoleServices=make(map[int]service.RoleService)
var RoleTranStatus =make(map[string]bool)

var SmsServices=make(map[int]service.SmsService)
var SmsTranStatus =make(map[string]bool)

var EmailServices=make(map[int]service.EmailService)
var EmailTranStatus =make(map[string]bool)

var PayServices=make(map[int]service.PaySecrtService)
var PayTranStatus =make(map[string]bool)

func GetAdminService()  service.AdminService{
	return nil
}

func GetUserService() service.UserService {
	return nil
}
func GetRoleService()  service.RoleService{
	return nil
}
func GetSmsService() service.SmsService {
	return nil
}
func GetEmailService() service.EmailService {
	return nil
}
func GetPayService()  service.PaySecrtService{
	return nil
}

func ReturnAdminService(service1 service.AdminService)  {

}

func ReturnUserService(service1 service.UserService)  {

}
func ReturnRoleService(service1 service.RoleService)  {

}
func ReturnSmsService(service1 service.SmsService)  {

}
func ReturnEmailService(service1 service.EmailService)  {

}
func ReturnPayService(service1 service.PaySecrtService)  {

}
type WebImpl struct {

	Lock *sync.RWMutex
}
func(web *WebImpl) CreateLock(){
	if web.Lock!=nil{
		web.Lock=&sync.RWMutex{}
	}
}
func(web *WebImpl) Instance()*sync.RWMutex{
	return web.Lock
}
type Web interface {
	 Instance()*sync.RWMutex
}
