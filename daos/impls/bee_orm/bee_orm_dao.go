package dao_bee_orm_impl

import (
	"fmt"
	data_db_bee "github.com/adminwjp/infrastructure-go/datas/dbs/bees"
	"github.com/adminwjp/users-go/daos"
	"github.com/beego/beego/v2/client/orm"
	"log"
	"strings"
)

type BeeOrmDaoImpl struct {
	daos.DaoImpl
	TranManager *data_db_bee.TranManager
	Query orm.QuerySeter
	Seter orm.RawSeter
	del bool
}


func(dao *BeeOrmDaoImpl) Wheres(cod *orm.Condition,and bool,wheres []*daos.DaoVal) {
	if wheres != nil && len(wheres) > 0 {
		for _, v := range wheres {
			switch v.Flag {
			case daos.WhereNone:

				break
			case daos.WhereEq:
				if !and{
					cod.Or(v.Name,v.Value)
				}else{
					cod.And(v.Name,v.Value)
				}
				break
			case daos.WhereLike:
				//__contains
				if !and{
					cod.Or(v.Name+"__contains",v.Value)
				}else{
					cod.And(v.Name+"__contains",v.Value)
				}
				break
			case daos.WhereIn:
				//__in
				if !and{
					cod.Or(v.Name+"__in",v.Value)
				}else{
					cod.And(v.Name+"__in",v.Value)
				}
				break
			case daos.WhereLt:
				if !and{
					cod.Or(v.Name+"__lt",v.Value)
				}else{
					cod.And(v.Name+"__lt",v.Value)
				}
				break
			case daos.WhereGt:
				if !and{
					cod.Or(v.Name+"__gt",v.Value)
				}else{
					cod.And(v.Name+"__gt",v.Value)
				}
				break
			case daos.WhereLte:
				//__gte
				if !and{
					cod.Or(v.Name+"__lte",v.Value)
				}else{
					cod.And(v.Name+"__lte",v.Value)
				}
				break

			case daos.WhereGte:
				if !and{
					cod.Or(v.Name+"__gte",v.Value)
				}else{
					cod.And(v.Name+"__gte",v.Value)
				}
				break
			default:
				break

			}
		}
	}
}
func(dao *BeeOrmDaoImpl) AndWheres(wheres []*daos.DaoVal) {
	if wheres != nil && len(wheres) > 0 {
		dao.Query=dao.TranManager.GetDb().QueryTable(dao.Model1)
		for _, v := range wheres {
			switch v.Flag {
			case daos.WhereNone:

				break
			case daos.WhereEq:
				dao.Query.Filter(v.Name,v.Value)
				break
			case daos.WhereLike:
				//__contains
				dao.Query.Filter(v.Name+"__contains",v.Value)
				break
			case daos.WhereIn:
				//__in
				dao.Query.Filter(v.Name+"__in",v.Value)
				break
			case daos.WhereLt:
				dao.Query.Filter(v.Name+"__lt",v.Value)
				break
			case daos.WhereGt:
				dao.Query.Filter(v.Name+"__gt",v.Value)
				break
			case daos.WhereLte:
				//__gte
				dao.Query.Filter(v.Name+"__lte",v.Value)
				break

			case daos.WhereGte:
				dao.Query.Filter(v.Name+"__gte",v.Value)
				break
			default:
				break

			}
		}
	}
}
func(dao *BeeOrmDaoImpl) Where() *orm.Condition {
	dao.Query=dao.TranManager.GetDb().QueryTable(dao.Model1)
	var cond=orm.NewCondition()
	dao.Wheres(cond,true,dao.Ands)
	dao.Wheres(cond,false,dao.Ors)
	return cond
}
func(dao *BeeOrmDaoImpl) One(obj interface{})error{
	if dao.Seter!=nil{
		/*_,err:=  dao.Seter.RowsToStruct(obj,"name","value")
		return err*/
		//https://beego.vip/docs/mvc/model/rawsql.md
	 	return 	dao.Seter.QueryRow(obj)
	}
	cod:=dao.Where()
	return  dao.Query.SetCond(cod).One(obj)
}

func(dao *BeeOrmDaoImpl) Count()(int,error){
	//if dao.IsWhere{
	cod:=dao.Where()
	//}
	s,err:= dao.Query.SetCond(cod).Count()
	return int(s), err
}
func(dao *BeeOrmDaoImpl) QueryCount(sql string,values ...interface{})(int,error){

	//var t dto.CountDto
	//_,err:=dao.TranManager.GetDb().Raw(sql,values).RowsToStruct(&t,"name","value")
	//return int(*&t.Total),err
	var c int
	err:=dao.TranManager.GetDb().Raw(sql,values).QueryRow(&c)
	return c, err

}
//column json bson name,name1...
func(dao *BeeOrmDaoImpl) Get(name string){

}

//insert
func(dao *BeeOrmDaoImpl) Add(obj interface{})(int,error){
	log.Println("data add")
	s,err:= dao.TranManager.GetDb().Insert(obj)
	return int(s), err
}

func(dao *BeeOrmDaoImpl)ExecuteUpdate(obj interface{})(int,error){
	s,err:= dao.TranManager.GetDb().Update(obj)
	return int(s), err
}
func(dao *BeeOrmDaoImpl)Delete()(int,error){
	return 0, nil
	//var cond=orm.NewCondition()  operation cannot execute withoout execute
	if dao.Id1!=nil{
		s,err:= dao.TranManager.GetDb().QueryTable(dao.Model1).Filter(dao.IdName1,dao.Id1).Delete()
		return int(s), err
	}
	dao.AndWheres(dao.Ands)
	if dao.Query!=nil{
		s,err:=   dao.Query.Delete()
		return int(s), err
	}
	return 0, nil
}
func(dao *BeeOrmDaoImpl) Execute()(int,error){
	if dao.Seter!=nil{
		res,err:=dao.Seter.Exec()
		return sqlResult(res,err)
	}
	return 0, nil
}

func(dao *BeeOrmDaoImpl) List(list interface{})(int64,error){
	var db =dao.TranManager.GetDb()

	query:=db.QueryTable(dao.Model1)
	_,err:=query.All(list)
	if err!=nil{
		return 0,err
	}
	c,err:=query.Count()
	if err!=nil{
		return  c,err
	}
	return  c,err
}
func(dao *BeeOrmDaoImpl) ListByPage(page int,size int,list interface{})(int64,error){
	var db =dao.TranManager.GetDb()
	limit,offset:=0,page*size
	if page>1{
		limit=(page-1)*size
	}
	if dao.Model1!=nil{

	}

	query:=db.QueryTable(dao.Model1)
	da:=query.Limit(offset).Offset(limit)
	_,err:=da.All(list)
	if err!=nil{
		return 0,err
	}
	c,err:=query.Count()
	log.Println("count:%d",c)
	//count=&c//0
	//*count=c//null ex
	return  c,err
}
//db
func(dao *BeeOrmDaoImpl) ExecuteSql(sql string,values ... interface{}){
	dao.Seter=dao.TranManager.GetDb().Raw(sql,values)
}
func(dao *BeeOrmDaoImpl) QuerySql(sql string,values ... interface{}){
	dao.Seter=dao.TranManager.GetDb().Raw(sql,values)
}
func(dao *BeeOrmDaoImpl)ExecuteList(sql string,list interface{})error{
	_,err:=dao.TranManager.GetDb().Raw(sql).QueryRows(list)
	return err
}
func(dao *BeeOrmDaoImpl) QueryList(sql string,sqlCount string,page int ,size int,list interface{},values ... interface{})(int64,error){
	if dao.View1==""{
		return 0,nil
	}
	var where=""
	where+=dao.WhereStrings(true,dao.Ands)
	where+=dao.WhereStrings(false,dao.Ors)
	where=strings.Trim(where," ")
	if len(where)>0{
		where=strings.TrimRight(where,"and")
		where=strings.TrimRight(where,"or")
	}
	_,err:=dao.TranManager.GetDb().Raw(sql+where).QueryRows(list)
	if err!=nil{
		return  0,err
	}
	var maps []orm.Params
	num,err:=dao.TranManager.GetDb().Raw(sqlCount+where).Values(&maps)
	if err!=nil{
		return  0,err
	}
	var cc int64 =0
	if err == nil && num > 0 {
		cc=maps[0]["count(*)"].(int64)
		fmt.Println(maps[0]["count(*)"]) // slene
	}
	return cc,err
}
func(dao *BeeOrmDaoImpl) ExecuteSqlToInt(sql string,values ... interface{})(int,error){
	res,err:=dao.TranManager.GetDb().Raw(sql,values).Exec()
	return sqlResult(res,err)
}



