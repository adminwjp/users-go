package service_impl

import (
	 "github.com/adminwjp/users-go/daos"
	 "github.com/adminwjp/users-go/models"
)
func (service *RpcServiceImpl)GetTranction()daos.TranDao{
	return service.TranManager
}
func (service *RpcServiceImpl)Clean()  {

	service.Dao=nil

}
type RpcServiceImpl struct {
	ServiceImpl

}

/**
添加
*/
func (service *RpcServiceImpl)Add(model *models.RpcModel)(int,error){
	return  service.Dao.Add(model)
}

/**
修改
*/
func (service *RpcServiceImpl)Update(model *models.RpcModel)(int,error){
	return  service.Dao.Update(model)
}

/**
删除
*/
func (service *RpcServiceImpl)Delete(id int64)(int,error){
	return  service.Dao.Delete(id)
}

/**
删除
*/
func (service *RpcServiceImpl)DeleteBatch(ids []int64)(int,error){
	return  service.Dao.DeleteBatch(ids)
}
func (service *RpcServiceImpl)List1()([]models.RpcModel,int64,error){
	var ms []models.RpcModel
	count,err:=  service.Dao.List(&ms)
	return *&ms, count, err
}
func (service *RpcServiceImpl)List(page int,size int)([]models.RpcModel,int64,error){
	var ms []models.RpcModel
	count,err:=  service.Dao.ListByPage(page,size,&ms)
	return *&ms, count, err
}


