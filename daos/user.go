package daos

import (
	"encoding/json"
	dto "github.com/adminwjp/infrastructure-go/dtos"
	"github.com/adminwjp/users-go/dtos"
	"github.com/adminwjp/users-go/models"
	"log"
)

//管理员接口  jinzhus 实现
type UserDaoImpl struct {
	BaseUserDaoImpl

}
func(dao *UserDaoImpl)Init(){
	dao.SetUserSql()
}
/*
	根据手机号、邮箱、用户名登录
*/
func(dao *UserDaoImpl) Login(input *dtos.UserInput)(*models.UserModel,error){
	var model1 models.UserModel
	var err =dao.BaseUserDaoImpl.Login(input,&model1)
	if err!=nil{
		log.Println("login data fail,err:%s",err.Error())
	}
	log.Println("login data pass:%d",model1.Id)

	if *&model1.Id>0{
		return  &model1, err
	}
	return nil, err
}

/*
	根据手机号登录
*/
func(dao *UserDaoImpl) LoginByPhone(input *dtos.UserPhoneInput)(*models.UserModel,error){
	var model1 models.UserModel
	var err =dao.BaseUserDaoImpl.LoginByPhone(input,&model1)
	if *&model1.Id>0{
		return  &model1, err
	}
	bu,_:=json.Marshal(*&model1)
	log.Printf("dao phone login fail,err:%s",string(bu))
	return nil, err
}
/*
	根据邮箱登录
*/
func (dao *UserDaoImpl)LoginByEmail(input *dtos.UserEmailInput)(*models.UserModel,error){
	var model1 models.UserModel
	var err =dao.BaseUserDaoImpl.LoginByEmail(input,&model1)
	if *&model1.Id>0{
		return  &model1, err
	}
	return nil, err
}

/*
	根据用户名登录
*/
func (dao *UserDaoImpl)LoginByUserName(input *dtos.UserUserNameInput)(*models.UserModel,error){
	var model1 models.UserModel
	var err =dao.BaseUserDaoImpl.LoginByUserName(input,&model1)
	if *&model1.Id>0{
		return  &model1, err
	}
	return nil, err
}

func(dao *UserDaoImpl) Get(id int64) (*models.UserModel,error){

	var model1 models.UserModel
	var err =dao.BaseUserDaoImpl.Get(id,&model1)
	if *&model1.Id>0{
		return  &model1, err
	}
	return nil, err
}
/*
	根据手机号、邮箱、用户名注册
*/
func (dao *UserDaoImpl)Register(input *dtos.UserInput)(int,error){
	log.Printf("dao Register flag,%d",input.Flag)
	return  dao.BaseUserDaoImpl.Register(input, func(input *dtos.UserInput) interface{} {
		var m= &models.UserModel{
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
func (dao *UserDaoImpl)RegisterByPhone(user *dtos.UserPhoneInput)(int,error){
	return  dao.BaseUserDaoImpl.RegisterByPhone(user, func(input *dtos.UserPhoneInput) interface{} {
		var m= &models.UserModel{
			Id: input.Id,Pwd: input.Pwd,RegIp: input.OperatorIp,RegDate: input.OperatorDate,
		}
		m.Phone=input.Phone
		return m
	})
}

/*
	根据邮箱注册
*/
func (dao *UserDaoImpl)RegisterByEmail(user *dtos.UserEmailInput)(int,error){
	return  dao.BaseUserDaoImpl.RegisterByEmail(user, func(input *dtos.UserEmailInput) interface{} {
		var m= &models.UserModel{
			Id: input.Id,Pwd: input.Pwd,RegIp: input.OperatorIp,RegDate: input.OperatorDate,
		}
		m.Email=input.Email
		return m
	})
}

/*
	根据用户名注册
*/
func (dao *UserDaoImpl)RegisterByUserName(user *dtos.UserUserNameInput)(int,error){
	return  dao.BaseUserDaoImpl.RegisterByUserName(user, func(input *dtos.UserUserNameInput) interface{} {
		var m= &models.UserModel{
			Id: input.Id,Pwd: input.Pwd,RegIp: input.OperatorIp,RegDate: input.OperatorDate,
		}
		m.UserName=input.UserName
		return m
	})
}
/*
	根据条件查询用户信息
*/
func(dao *UserDaoImpl) List(user *dtos.GetUserInput) ([]models.UserModel,int64,error) {
	var ms []models.UserModel
	dao.M=&models.UserModel{}
	count,err:=dao.BaseUserDaoImpl.List(&user.GetBaseUserInput,&ms)
	return *&ms,count,err
}