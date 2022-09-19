package service

import (
	"github.com/adminwjp/users-go/daos"
	"github.com/adminwjp/users-go/models"
)


type ConfigService interface {
	Clean()
	GetTranction()daos.TranDao
	/**
	添加
	*/
	Add(model *models.ConfigModel)(int,error)

	/**
	修改
	*/
	Update(model *models.ConfigModel)(int,error)

	/**
	删除
	*/
	Delete(id int64)(int,error)

	/**
	删除
	*/
	DeleteBatch(ids []int64)(int,error)

	List1()([]models.ConfigModel,int64,error)

	List(page int,size int)([]models.ConfigModel,int64,error)
}
