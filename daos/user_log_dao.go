package daos

import (
	"github.com/adminwjp/users-go/dtos"
	"github.com/adminwjp/users-go/models"
)

type UserLogDaoImpl struct{
	BaseUserLogDaoImpl
}
func (dao *UserLogDaoImpl)Add(model *models.UserModel)(int,error){
	return  dao.BaseUserLogDaoImpl.Add(model)
}
func (dao *UserLogDaoImpl)List(user *dtos.GetUserLogInput) ([]models.UserModel,int64,error){
	var ms []models.UserModel
	count,err:=dao.BaseUserLogDaoImpl.List(user,&ms)
	return *&ms,count,err
}
type UserLogDao interface {

	/*
		添加
	*/
	Add(model *models.UserLogModel)(int,error)

	/*
	   根据条件查询用户日志信息
	*/
	List(user *dtos.GetUserLogInput) ([]models.UserLogModel,int64,error)

}

