package service

import (
	"github.com/adminwjp/users-go/daos"
	"github.com/adminwjp/users-go/models"
)

type RpcService interface {
	Clean()
	GetTranction()daos.TranDao
	/**
	添加
	*/
	Add(model *models.RpcModel)(int,error)

	/**
	修改
	*/
	Update(model *models.RpcModel)(int,error)

	/**
	删除
	*/
	Delete(id int64)(int,error)

	/**
	删除
	*/
	DeleteBatch(ids []int64)(int,error)

	List1()([]models.RpcModel,int64,error)

	List(page int,size int)([]models.RpcModel,int64,error)
}
