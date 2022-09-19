package service

import (
	"github.com/adminwjp/users-go/daos"
	"github.com/adminwjp/users-go/models"
)

//短信 配置服务接口
type SmsService interface {
	Clean()
	GetTranction()daos.TranDao
	/**
	添加
	*/
	Add(model *models.SmsConfigModel)(int,error)

	/**
	修改
	*/
	Update(model *models.SmsConfigModel)(int,error)

	/**
	删除
	*/
	Delete(id int64)(int,error)

	/**
	删除
	*/
	DeleteBatch(ids []int64)(int,error)
	List1()([]models.SmsConfigModel,int64,error)
	List(page int,size int)([]models.SmsConfigModel,int64,error)
}
