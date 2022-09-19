package service

import (
	"github.com/adminwjp/users-go/daos"
	"github.com/adminwjp/users-go/models"
)

//支付 配置服务接口
type PaySecrtService interface {
	Clean()
	GetTranction()daos.TranDao
	/**
	添加
	*/
	Add(model *models.PaySecrtConfigModel)(int,error)

	/**
	修改
	*/
	Update(model *models.PaySecrtConfigModel)(int,error)

	/**
	删除
	*/
	Delete(id int64)(int,error)

	/**
	删除
	*/
	DeleteBatch(ids []int64)(int,error)
	List1()([]models.PaySecrtConfigModel,int64,error)
	List(page int,size int)([]models.PaySecrtConfigModel,int64,error)
}