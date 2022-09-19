package service_impl

import (
	 "github.com/adminwjp/users-go/daos"

	 "github.com/adminwjp/users-go/models"
)
func (service *EmailServiceImpl)GetTranction()daos.TranDao{
	return service.TranManager
}
func (service *EmailServiceImpl)Clean()  {

	service.Dao=nil

}

/*type EmailCrudServiceImpl struct {
	EmailServiceImpl
}
func (service *EmailCrudServiceImpl)Add(obj interface{})(int,error){
	return  service.EmailServiceImpl.Add(obj.(model.EmailConfigModel))
}
func (service *EmailCrudServiceImpl)Update(obj interface{})(int,error)
func (service *EmailCrudServiceImpl)Delete(id interface{})(int,error)
func (service *EmailCrudServiceImpl)DeleteBatch(ids interface{})(int,error)
func (service *EmailCrudServiceImpl)List(list interface{},count *int64)error
func (service *EmailCrudServiceImpl)ListByPage(page int,size int,list interface{},count *int64)error*/

//支付服务接口
type EmailServiceImpl struct {
	ServiceImpl

}

/**
添加
*/
func (service *EmailServiceImpl)Add(model *models.EmailConfigModel)(int,error){
	return  service.Dao.Add(model)
}

/**
修改
*/
func (service *EmailServiceImpl)Update(model *models.EmailConfigModel)(int,error){
	return  service.Dao.Update(model)
}

/**
删除
*/
func (service *EmailServiceImpl)Delete(id int64)(int,error){
	return  service.Dao.Delete(id)
}

/**
删除
*/
func (service *EmailServiceImpl)DeleteBatch(ids []int64)(int,error){
	return  service.Dao.DeleteBatch(ids)
}
func (service *EmailServiceImpl)List1()([]models.EmailConfigModel,int64,error){
	var ms []models.EmailConfigModel
	count,err:=  service.Dao.List(&ms)
	return *&ms, count, err
}
func (service *EmailServiceImpl)List(page int,size int)([]models.EmailConfigModel,int64,error){
	var ms []models.EmailConfigModel
	count,err:=  service.Dao.ListByPage(page,size,&ms)
	return *&ms, count, err
}
