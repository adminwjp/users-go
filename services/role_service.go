package service

import (
	"github.com/adminwjp/users-go/daos"
	"github.com/adminwjp/users-go/dtos"
	"github.com/adminwjp/users-go/models"
)

//角色服务接口
type RoleService interface {
	Clean()
	GetTranction()daos.TranDao
	/**
	添加
	*/
	Add(model *models.RoleModel)(int,error)

	/**
	修改
	*/
	Update(model *models.RoleModel)(int,error)

	/**
	删除
	*/
	Delete(id int64)(int,error)

	/**
	删除
	*/
	DeleteBatch(ids []int64)(int,error)

	Parent()([]dtos.RoleOutPut,error)
	List(page int,size int)([]models.RoleModel,int64,error)

}
