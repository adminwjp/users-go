package service_impl

import (
	"github.com/adminwjp/users-go/daos"
	"log"

	"github.com/adminwjp/users-go/models"
)
func (service *SmsServiceImpl)GetTranction()daos.TranDao{
	return service.TranManager
}
func (service *SmsServiceImpl)Clean()  {

	service.Dao=nil
}
//短信服务接口
type SmsServiceImpl struct {
	ServiceImpl
}

/**
添加
*/
func (service *SmsServiceImpl)Add(model *models.SmsConfigModel)(int,error){
	return  service.Dao.Add(model)
}

/**
修改
*/
func (service *SmsServiceImpl)Update(model *models.SmsConfigModel)(int,error){
	return  service.Dao.Update(model)
}

/**
删除
*/
func (service *SmsServiceImpl)Delete(id int64)(int,error){
	return  service.Dao.Delete(id)
}

/**
删除
*/
func (service *SmsServiceImpl)DeleteBatch(ids []int64)(int,error){
	return  service.Dao.DeleteBatch(ids)
}
func (service *SmsServiceImpl)List1()([]models.SmsConfigModel,int64,error){
	var ms []models.SmsConfigModel
	count,err:=  service.Dao.List(&ms)
	return *&ms, count, err
}
func (service *SmsServiceImpl)List(page int,size int)([]models.SmsConfigModel,int64,error){
	var ms []models.SmsConfigModel
	if service.Dao==nil{
		//first pass after bug
		//new
		log.Println("service Dao is nil ")
	}
	count,err:=  service.Dao.ListByPage(page,size,&ms)
	return *&ms, count, err
}