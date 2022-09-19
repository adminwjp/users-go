package dao_gorm_jinzhu_impl

import (
	data_db_gorm_jinzhu "github.com/adminwjp/infrastructure-go/datas/dbs/groms/jinzhus"
	dto "github.com/adminwjp/infrastructure-go/dtos"
	"github.com/adminwjp/users-go/daos"
	"github.com/adminwjp/users-go/models"
	"log"

	"strings"
)

type GormDaoImpl struct {
	daos.DaoImpl
	Auto bool
	//data_db_gorm_jinzhu.TranManager
	TranManager *data_db_gorm_jinzhu.TranManager
}

func(dao *GormDaoImpl) Wheres(and bool,wheres []*daos.DaoVal) {
	if wheres != nil && len(wheres) > 0 {
		for _, v := range wheres {
			switch v.Flag {
			case daos.WhereNone:

				break
			case daos.WhereEq:
				if !and{
					dao.TranManager.GetDb().Or(v.Name+"=? ",v.Value)
				}else{
					dao.TranManager.GetDb().Where(v.Name+"=? ",v.Value)
				}
				break
			case daos.WhereLike:
				//__contains
				if !and{
					dao.TranManager.GetDb().Or(v.Name+" like %?% ",v.Value)
				}else{
					dao.TranManager.GetDb().Where(v.Name+" like %?% ",v.Value)
				}
				break
			case daos.WhereIn:
				//__in
				if !and{
					dao.TranManager.GetDb().Or(v.Name+" in (?) ",v.Value)
				}else{
					dao.TranManager.GetDb().Where(v.Name+" in (?) ",v.Value)
				}
				break
			case daos.WhereLt:
				if !and{
					dao.TranManager.GetDb().Or(v.Name+"< ? ",v.Value)
				}else{
					dao.TranManager.GetDb().Where(v.Name+"< ? ",v.Value)
				}
				break
			case daos.WhereGt:
				if !and{
					dao.TranManager.GetDb().Or(v.Name+"> ? ",v.Value)
				}else{
					dao.TranManager.GetDb().Where(v.Name+"> ? ",v.Value)
				}
				break
			case daos.WhereLte:
				//__gte
				if !and{
					dao.TranManager.GetDb().Or(v.Name+"<= ? ",v.Value)
				}else{
					dao.TranManager.GetDb().Where(v.Name+"<= ? ",v.Value)
				}
				break

			case daos.WhereGte:
				if !and{
					dao.TranManager.GetDb().Or(v.Name+">= ? ",v.Value)
				}else{
					dao.TranManager.GetDb().Where(v.Name+">= ? ",v.Value)
				}
				break
			default:
				break

			}
		}
	}
}
func(dao *GormDaoImpl) Where()  {

	if dao.Ands!=nil&&len(dao.Ands)>0{
		dao.Wheres(true,dao.Ands)
	}
	if dao.Ors!=nil&&len(dao.Ors)>0{
		dao.Wheres(false,dao.Ors)
	}
}
func(dao *GormDaoImpl) One(obj interface{})error{
	if dao.IsWhere1{
		dao.Where()
	}
	return dao.TranManager.GetDb().First(obj).Error
}

func(dao *GormDaoImpl) Count()(int,error){
	if dao.IsWhere1{
		dao.Where()
	}
	var t dto.CountDto
	db:=dao.TranManager.GetDb().Scan(&t)
	log.Println("count:%d",*&t.Total)
	return int(*&t.Total),db.Error
}
func(dao *GormDaoImpl) QueryCount(sql string,values ...interface{})(int,error){

	var t dto.CountDto
	db:=dao.TranManager.GetDb()
	db.Debug().Raw(sql,values... ).Scan(&t)
	log.Println("count:%d",*&t.Total)
	return int(*&t.Total),db.Error
}
//column json bson name,name1...
func(dao *GormDaoImpl) Get(name string){

}

//insert
func(dao *GormDaoImpl) Add(obj interface{})(int,error){
	db:=dao.TranManager.GetDb().Save(obj)
	//dao.TranManager.SetDb(db)
	return  int(db.RowsAffected),db.Error
}
func(dao *GormDaoImpl)ExecuteUpdate(obj interface{})(int,error){
	db:=dao.TranManager.GetDb().Updates(obj)
	//dao.TranManager.SetDb(db)
	return  int(db.RowsAffected),db.Error
}

func(dao *GormDaoImpl)Delete()(int,error){
	dao.Where()
	if dao.Model1!=nil{
		db:=dao.TranManager.GetDb().Delete(dao.Model1)
		//dao.TranManager.SetDb(db)
		return  int(db.RowsAffected),db.Error
	}
	return 0, nil
}

func(dao *GormDaoImpl) Execute()(int,error){
	return 0,nil
}

func(dao *GormDaoImpl) List(list interface{})(int64,error) {
	dao.Where()
	db:=dao.TranManager.GetDb().Model(dao.Model1)
	//dao.TranManager.SetDb(db)
	db.Scan(list)
	if db.Error!=nil{
		return 0,db.Error
	}
	var count int64
	db.Count(&count)
	return count,db.Error
}
func(dao *GormDaoImpl) ListByPage(page int,size int,list interface{})(int64,error){
	dao.Where()
	offset,limit:=0,page*size
	if page>1{
		offset=(page-1)*size
	}
	da:=dao.TranManager.GetDb().Model(dao.Model1)

	da1:=da.Limit(limit).Offset(offset).Scan(list)//.Limit(-1).Offset(-1).Count(&count)
	var count int64
	//0
	da1=da.Count(&count)

	log.Println(*&count)
	return count,da1.Error

}
//db
func(dao *GormDaoImpl) ExecuteSql(sql string,values ... interface{}){
	dao.TranManager.GetDb().Exec(sql,values... )
}
func(dao *GormDaoImpl) QuerySql(sql string,values ... interface{}){
	dao.TranManager.GetDb().Raw(sql,values... )
}
func(dao *GormDaoImpl)ExecuteList(sql string,list interface{})error{
	return dao.TranManager.GetDb().Raw(sql).Scan(list).Error
}

func(dao *GormDaoImpl) QueryList(sql string,sqlCount string,page int ,size int,list interface{},values ... interface{})(int64,error){
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
	db:=dao.TranManager.GetDb()
	da:=db.Raw(sql+where)
	var users []models.AdminLogModel
	da=da.Scan(&users)
	if da.Error!=nil{
		return 0,da.Error
	}
	var c dto.CountDto
	da=db.Raw(sqlCount+where).Scan(&c)
	return c.Total,da.Error
}
func(dao *GormDaoImpl) ExecuteSqlToInt(sql string,values ... interface{})(int,error){
	db:=dao.TranManager.GetDb()
	db=db.Exec(sql,values... )
	return  int(db.RowsAffected),db.Error
}
