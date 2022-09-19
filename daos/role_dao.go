package daos

import (
	data "github.com/adminwjp/infrastructure-go/datas"
	"github.com/adminwjp/users-go/datas"
	"github.com/adminwjp/users-go/dtos"
	"github.com/adminwjp/users-go/models"
)

type RoleDaoImpl struct {
	TranManager TranDao
	CrudDaoImpl
}
func(dao *RoleDaoImpl)GetTranction()TranDao{
	return dao.TranManager
}


func (dao *RoleDaoImpl)Parent()([]dtos.RoleOutPut,error){
	//sql:="select id Id ,name Name,parent_id as ParentId  from t_role "
	sql:="select id  ,name ,parent_id    from t_role "
	da:=dao.Data()
	dao.init(da)
	var ms []dtos.RoleOutPut
	if datas.GlobalConfig.DataFlag== data.DataDb{
		err:=da.ExecuteList(sql,&ms)
		return *&ms,err
	}
	//var c int64
	//da.Get("id,name,parent_id")
	//err:= da.List(&ms,c)
	var ms1 []models.RoleModel
	_,err:= da.List(&ms1)
	if ms1!=nil&&len(ms1)>0{
		ms=make([]dtos.RoleOutPut,len(ms1))
		for i := 0; i < len(ms1); i++ {
			ms[i]=dtos.RoleOutPut{Id: ms1[i].Id,Name: ms1[i].Name,ParentId:  ms1[i].ParentId}
		}
	}
	return *&ms,err
}



//角色接口
type RoleDao interface {

	GetTranction() TranDao
	CrudDao


	Parent()([]dtos.RoleOutPut,error)

}
