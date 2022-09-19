package service_impl

import (
	"github.com/adminwjp/users-go/daos"

	"github.com/adminwjp/users-go/dtos"
	"github.com/adminwjp/users-go/models"
)
func (service *RoleServiceImpl)GetTranction()daos.TranDao{
	return service.RoleDao.GetTranction()
}
func (service *RoleServiceImpl)Clean()  {

	service.RoleDao=nil
}
//角色服务接口
type RoleServiceImpl struct {
	RoleDao daos.RoleDao
	ServiceImpl
}

/**
添加
*/
func (roleService *RoleServiceImpl)Add(role *models.RoleModel)(int,error){
	return  roleService.RoleDao.Add(role)
}

/**
修改
*/
func (roleService *RoleServiceImpl)Update(role *models.RoleModel)(int,error){
	return  roleService.RoleDao.Update(role)
}

/**
删除
*/
func (roleService *RoleServiceImpl)Delete(id int64)(int,error){
	return  roleService.RoleDao.Delete(id)
}

/**
删除
*/
func (roleService *RoleServiceImpl)DeleteBatch(ids []int64)(int,error){
	return  roleService.RoleDao.DeleteBatch(ids)
}
func (service *RoleServiceImpl)Parent()([]dtos.RoleOutPut,error){
	return  service.RoleDao.Parent()
}
func (service *RoleServiceImpl)List1()([]models.RoleModel,int64,error){
	var ms []models.RoleModel
	count,err:=  service.RoleDao.List(&ms)
	return *&ms, count, err
}
func (service *RoleServiceImpl)List(page int,size int)([]models.RoleModel,int64,error){
	var ms []models.RoleModel
	count,err:=  service.RoleDao.ListByPage(page,size,&ms)
	return *&ms, count, err
}