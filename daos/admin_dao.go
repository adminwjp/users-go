package daos

import (
	"github.com/adminwjp/users-go/dtos"
	"github.com/adminwjp/users-go/models"
)

//管理员接口 mong 实现
type AdminDao interface {
	BaseUserDao
	/*
		根据手机号、邮箱、用户名登录
	*/
	Login(user *dtos.UserInput)(*models.AdminModel,error)

	/*
		根据手机号登录
	*/
	 LoginByPhone(user *dtos.UserPhoneInput)(*models.AdminModel,error)
	/*
		根据邮箱登录
	*/
	LoginByEmail(user *dtos.UserEmailInput)(*models.AdminModel,error)

	/*
		根据用户名登录
	*/
	LoginByUserName(user *dtos.UserUserNameInput)(*models.AdminModel,error)
	Get(id int64)(*models.AdminModel,error)
	/*
		根据条件查询用户信息
	*/
	List(input *dtos.GetAdminInput)([]models.AdminModel,int64,error)
	/*
		根据旧密码修改密码
	*/
	UpdatePwdByOldPwd(pwd string,newPwd string)(int,error)
}
