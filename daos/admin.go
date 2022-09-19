package daos

import (
	dto "github.com/adminwjp/infrastructure-go/dtos"
	"github.com/adminwjp/users-go/dtos"
	"github.com/adminwjp/users-go/models"
)
//管理员接口  jinzhus 实现
type AdminDaoImpl struct {
	BaseUserDaoImpl

}
func(dao *AdminDaoImpl)Init(){
	dao.SetUserSql()
}
/*
	根据手机号、邮箱、用户名登录
*/
func(dao *AdminDaoImpl) Login(input *dtos.UserInput)(*models.AdminModel,error){
	var model1 models.AdminModel
	var err =dao.BaseUserDaoImpl.Login(input,&model1)
	if *&model1.Id>0{
		return  &model1, err
	}
	return nil, err
}

/*
	根据手机号登录
*/
func(dao *AdminDaoImpl) LoginByPhone(input *dtos.UserPhoneInput)(*models.AdminModel,error){
	var model1 models.AdminModel
	var err =dao.BaseUserDaoImpl.LoginByPhone(input,&model1)
	if *&model1.Id>0{
		return  &model1, err
	}
	return nil, err
}
/*
	根据邮箱登录
*/
func (dao *AdminDaoImpl)LoginByEmail(input *dtos.UserEmailInput)(*models.AdminModel,error){
	var model1 models.AdminModel
	var err =dao.BaseUserDaoImpl.LoginByEmail(input,&model1)
	if *&model1.Id>0{
		return  &model1, err
	}
	return nil, err
}

/*
	根据用户名登录
*/
func (dao *AdminDaoImpl)LoginByUserName(input *dtos.UserUserNameInput)(*models.AdminModel,error){
	var model1 models.AdminModel
	var err =dao.BaseUserDaoImpl.LoginByUserName(input,&model1)
	if *&model1.Id>0{
		return  &model1, err
	}
	return nil, err
}

func(dao *AdminDaoImpl) Get(id int64) (*models.AdminModel,error){

	var model1 models.AdminModel
	var err =dao.BaseUserDaoImpl.Get(id,&model1)
	if *&model1.Id>0{
		return  &model1, err
	}
	return nil, err
}

/*
	根据手机号、邮箱、用户名注册
*/
func (dao *AdminDaoImpl)Register(input *dtos.UserInput)(int,error){
	return  dao.BaseUserDaoImpl.Register(input, func(input *dtos.UserInput) interface{} {
		var m= &models.AdminModel{
			Id: input.Id,Pwd: input.Pwd,RegIp: input.OperatorIp,RegDate: input.OperatorDate,
		}
		switch input.Flag {
			case dto.AccounTypeByPhone:
				m.Phone=input.Account
				break
			case dto.AccounTypeByUsername:
				m.UserName=input.Account
				break
			case dto.AccounTypeByEamil:
				m.Email=input.Account
				break
			default:
				break
		}
		return m
	})
}


/*
	根据手机号注册
*/
func (dao *AdminDaoImpl)RegisterByPhone(user *dtos.UserPhoneInput)(int,error){
	return  dao.BaseUserDaoImpl.RegisterByPhone(user, func(input *dtos.UserPhoneInput) interface{} {
		var m= &models.AdminModel{
			Id: input.Id,Pwd: input.Pwd,RegIp: input.OperatorIp,RegDate: input.OperatorDate,
		}
		m.Phone=input.Phone
		return m
	})
}

/*
	根据邮箱注册
*/
func (dao *AdminDaoImpl)RegisterByEmail(user *dtos.UserEmailInput)(int,error){
	return  dao.BaseUserDaoImpl.RegisterByEmail(user, func(input *dtos.UserEmailInput) interface{} {
		var m= &models.AdminModel{
			Id: input.Id,Pwd: input.Pwd,RegIp: input.OperatorIp,RegDate: input.OperatorDate,
		}
		m.Email=input.Email
		return m
	})
}

/*
	根据用户名注册
*/
func (dao *AdminDaoImpl)RegisterByUserName(user *dtos.UserUserNameInput)(int,error){
	return  dao.BaseUserDaoImpl.RegisterByUserName(user, func(input *dtos.UserUserNameInput) interface{} {
		var m= &models.AdminModel{
			Id: input.Id,Pwd: input.Pwd,RegIp: input.OperatorIp,RegDate: input.OperatorDate,
		}
		m.UserName=input.UserName
		return m
	})
}
/*
	根据条件查询用户信息
*/
func(dao *AdminDaoImpl) List(user *dtos.GetAdminInput) ([]models.AdminModel,int64,error) {
	var ms []models.AdminModel
	dao.M=models.AdminModel{}
	count, err:=dao.BaseUserDaoImpl.List(&user.GetBaseUserInput,&ms)
	return *&ms,count,err
}