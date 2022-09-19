package daos

import (
	data "github.com/adminwjp/infrastructure-go/datas"
	"github.com/adminwjp/users-go/datas"
	"github.com/adminwjp/users-go/dtos"
	"github.com/adminwjp/users-go/models"
	"strings"
)
type AdminLogDaoImpl struct{
	BaseUserLogDaoImpl
}
func (dao *AdminLogDaoImpl)Add(model *models.AdminLogModel)(int,error){
	return  dao.BaseUserLogDaoImpl.Add(model)
}
func (dao *AdminLogDaoImpl)List(user *dtos.GetUserLogInput) ([]models.AdminLogModel,int64,error){
	var ms []models.AdminLogModel
	count,err:=dao.BaseUserLogDaoImpl.List(user,&ms)
	return *&ms,count,err
}
//用户日志接口  实现
type BaseUserLogDaoImpl struct {
	TranManager TranDao
	//db sql
	Data func()Dao

	selectLog string
	selectCountLog string

	user string
	isAdmin bool

	view string
	table string
	doc string
	collection string
}
func (dao *BaseUserLogDaoImpl)SetAdmin(admin bool){
	dao.isAdmin=admin
}
func (dao *BaseUserLogDaoImpl)SetAdminSql()  {
	dao.user="admin_log"
	view:="v_admin_log"
	table:="t_admin_log"
	if !dao.isAdmin{
		dao.user="user_log"
		view="v_user_log"
		table="t_user_log"
	}
	dao.view=view
	dao.table=table
	dao.doc=dao.user
	dao.collection=dao.user
	n:=strings.Replace(view,"_log","",-1)
	dao.selectLog="select "+view+".* from "+view+"  left join "+n+" u.  where "+view+".user_id="+n+".id"
	dao.selectCountLog="select count(*) from "+view+"  left join "+n+" u.  where "+view+".user_id="+n+".id"
}

/*
	添加
*/
func (dao *BaseUserLogDaoImpl)Add(model interface{})(int,error){
	da:=dao.Data()
	dao.init(da)
	return  da.Add(model)
}

func (dao *BaseUserLogDaoImpl)init(da Dao)  {
	da.View(dao.view)
	da.Db(MDb)
	da.Doc(dao.doc)
	da.Collection(dao.collection)
}
/*
根据条件查询用户日志信息
*/
func(dao *BaseUserLogDaoImpl) List(user *dtos.GetUserLogInput,list interface{}) (int64,error) {
	if datas.GlobalConfig.DataFlag!= data.DataDb{
		return 0,nil
	}
	da:=dao.Data()
	dao.init(da)
	if user.Phone!=""{
		da.OrLike("u.phone",user.Phone)
	}
	if user.Email!=""{da.OrLike("u.email",user.Email)
	}
	if user.UserName!=""{
		da.OrLike("u.user_name",user.UserName)
	}
	if user.OperatorStartDate>0&& user.OperatorEndDate>0{
		da.OrBetween("u.operator_date",user.OperatorStartDate,user.OperatorEndDate)
	}else if user.OperatorStartDate>0{
		da.OrGt("u.operator_date",user.OperatorStartDate)
	}else if user.OperatorStartDate>0&& user.OperatorEndDate>0{
		da.OrLt("u.operator_date",user.OperatorEndDate)
	}
	return  da.QueryList(dao.selectLog,dao.selectCountLog,user.Page,user.Size,list)
}
type AdminLogDao interface {

	/*
		添加
	*/
	Add(model *models.AdminLogModel)(int,error)

	/*
	   根据条件查询用户日志信息
	*/
	List(user *dtos.GetUserLogInput) ([]models.AdminLogModel,int64,error)

}
