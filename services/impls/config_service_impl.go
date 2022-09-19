package service_impl

import (
	 "github.com/adminwjp/users-go/daos"
	 "github.com/adminwjp/users-go/models"
)
func (service *ConfigServiceImpl)GetTranction()daos.TranDao{
	return service.TranManager
}
func (service *ConfigServiceImpl)Clean()  {

	service.Dao=nil

}
type ConfigServiceImpl struct {
	ServiceImpl

}

/**
添加
*/
func (service *ConfigServiceImpl)Add(model *models.ConfigModel)(int,error){
	return  service.Dao.Add(model)
}

/**
修改
*/
func (service *ConfigServiceImpl)Update(model *models.ConfigModel)(int,error){
	return  service.Dao.Update(model)
}

/**
删除
*/
func (service *ConfigServiceImpl)Delete(id int64)(int,error){
	return  service.Dao.Delete(id)
}

/**
删除
*/
func (service *ConfigServiceImpl)DeleteBatch(ids []int64)(int,error){
	return  service.Dao.DeleteBatch(ids)
}
func (service *ConfigServiceImpl)List1()([]models.ConfigModel,int64,error){
	var ms []models.ConfigModel
	count,err:=  service.Dao.List(&ms)
	return *&ms, count, err
}
func (service *ConfigServiceImpl)List(page int,size int)([]models.ConfigModel,int64,error){
	var ms []models.ConfigModel
	count,err:=  service.Dao.ListByPage(page,size,&ms)
	return *&ms, count, err
}

