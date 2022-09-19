package service_impl

import (
	"github.com/adminwjp/users-go/daos"

	"github.com/adminwjp/users-go/models"
)
func (service *PayServiceImpl)GetTranction()daos.TranDao{
	return service.TranManager
}
func (service *PayServiceImpl)Clean()  {

	service.Dao=nil

}
//支付服务接口
type PayServiceImpl struct {
	ServiceImpl
}

/**
添加
*/
func (service *PayServiceImpl)Add(model *models.PaySecrtConfigModel)(int,error){
	return  service.Dao.Add(model)
}

/**
修改
*/
func (service *PayServiceImpl)Update(model *models.PaySecrtConfigModel)(int,error){
	return  service.Dao.Update(model)
}

/**
删除
*/
func (service *PayServiceImpl)Delete(id int64)(int,error){
	return  service.Dao.Delete(id)
}

/**
删除
*/
func (service *PayServiceImpl)DeleteBatch(ids []int64)(int,error){
	return  service.Dao.DeleteBatch(ids)
}
func (service *PayServiceImpl)List1()([]models.PaySecrtConfigModel,int64,error){
	var ms []models.PaySecrtConfigModel
	count,err:=  service.Dao.List(&ms)
	return *&ms, count, err
}
func (service *PayServiceImpl)List(page int,size int)([]models.PaySecrtConfigModel,int64,error){
	var ms []models.PaySecrtConfigModel
	count,err:=  service.Dao.ListByPage(page,size,&ms)
	return *&ms, count, err
}