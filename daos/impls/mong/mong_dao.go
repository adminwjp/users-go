package dao_mong_impl

import (
	"encoding/json"
	"github.com/adminwjp/users-go/daos"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"reflect"
)
var EmptyWhere=make(map[string]interface{},0)
type MongDaoImpl struct {
	daos.DaoImpl
	Session *mgo.Session
	where map[string]interface{}
	update map[string]interface{}
	Auto bool
}

func(dao *MongDaoImpl) Wheres(wheres []*daos.DaoVal)map[string]interface{} {


	if wheres != nil && len(wheres) > 0 {
		dao.Trim(wheres)
		var whs=make(map[string]interface{},len(wheres))
		for _, v := range wheres {
			switch v.Flag {
			case daos.WhereNone:

				break
			case daos.WhereEq:
				whs[v.Name] = v.Value
				break
			case daos.WhereLike:
				whs[v.Name] = "/"+v.Value.(string)+"/"
				break
			case daos.WhereIn:
				whs[v.Name] = bson.M{"$in": v.Value}
				break
			case daos.WhereLt:
				whs[v.Name] = bson.M{"$lt": v.Value}
				break
			case daos.WhereGt:
				whs[v.Name] = bson.M{"$gt": v.Value}
				break
			case daos.WhereLte:
				whs[v.Name] = bson.M{"$lte": v.Value}
				break
			case daos.WhereGte:
				whs[v.Name] = bson.M{"$gte": v.Value}
				break
			default:
				break

			}
		}
		return whs
	}
	return nil
}
func(dao *MongDaoImpl) Where()  {
	var whs=dao.Wheres(dao.Ands)
	if whs!=nil{
		dao.where=whs
	}
	whs=dao.Wheres(dao.Ors)
	if whs!=nil{
		if dao.where==nil{
			dao.where=make(map[string]interface{})
		}
		dao.where["$or"]=whs
	}
}
func(dao *MongDaoImpl) One(obj interface{})error{
	//if dao.IsWhere{
		dao.Where()
	//}
	s:= dao.Session.DB(dao.Db1).C(dao.Collection1)
	if dao.where==nil{
		dao.where=EmptyWhere
	}
	return  s.Find(dao.where).One(obj)
}

func(dao *MongDaoImpl) Count()(int,error){
	//if dao.IsWhere{
	dao.Where()
	//}
	if dao.Session==nil{
		log.Println("dao mong Session is nil ")
	}
	log.Println(dao.Db1+"  "+dao.Collection1)
	s:= dao.Session.DB(dao.Db1).C(dao.Collection1)
	if dao.where==nil{
		dao.where=EmptyWhere
	}
	log.Printf("count where count:%d ", len(dao.where))
	return  s.Find(dao.where).Count()
}

//column json bson name,name1...
func(dao *MongDaoImpl) Get(name string){

}

//insert
func(dao *MongDaoImpl) Add(obj interface{})(int,error){
	err:=  dao.Session.DB(dao.Db1).C(dao.Collection1).Insert(obj)
	if err!=nil{
		return 0, err
	}
	return 1, err
}
func(dao *MongDaoImpl)ExecuteUpdate(obj interface{})(int,error){
	update:=bson.M{}
	val:=reflect.ValueOf(obj).Elem()
	type1:=reflect.TypeOf(obj).Elem()
	n:=type1.NumField()
	for i := 0; i < n; i++ {
		f:=type1.Field(i)
		b:=f.Tag.Get("bson")
		if b==""{
			b=f.Tag.Get("json")
			if b==""{
				b=f.Name
			}
		}
		if b==dao.IdName1{
			//continue
		}
		update[b]=val.Field(i).Interface()
	}
	bu,_:=json.Marshal(update)
	log.Printf("update entity->docment json "+string(bu))
	err:=  dao.Session.DB(dao.Db1).C(dao.Collection1).Update(map[string]interface{}{dao.IdName1:dao.Id1},update)
	if err!=nil{
		log.Printf("update entity->docment fail,err:%s",err.Error())
		return 0, err
	}
	return 1, err
}
func(dao *MongDaoImpl)Delete()(int,error){
	err:=  dao.Session.DB(dao.Db1).C(dao.Collection1).RemoveId(map[string]interface{}{dao.IdName1:dao.Id1})
	if err!=nil{
		return 0, err
	}
	return 1, err
}
func(dao *MongDaoImpl) Execute()(int,error){
	dao.update= map[string]interface{}{}
	if dao.Updates!=nil&&len(dao.Updates)>0{
		for k, v := range dao.Updates {
			dao.update[k]=v
		}
	}
	if dao.UpdateIncrs!=nil&&len(dao.UpdateIncrs)>0{
		for k, _ := range dao.UpdateIncrs {
			dao.update[k]=bson.M{"$inc": bson.M{k: 1}}
		}
	}
	if len(dao.update)<1{
		return 0, nil
	}
	//if dao.IsWhere{
	dao.Where()
	//}
	s:= dao.Session.DB(dao.Db1).C(dao.Collection1)
	if dao.where==nil{
		dao.where=EmptyWhere
	}
	err:=  s.Update(dao.where,dao.update)
	if err!=nil{
		return 0, err
	}
	return 1, err
}
func(dao *MongDaoImpl) QueryCount(sql string,values ...interface{})(int,error){
	return dao.Count()
}
func(dao *MongDaoImpl) List(list interface{})(int64,error){
	//if dao.IsWhere{
	dao.Where()
	//}
	s:= dao.Session.DB(dao.Db1).C(dao.Collection1)
	if dao.where==nil{
		dao.where=EmptyWhere
	}
	query:=s.Find(dao.where)
	err:=query.All(list)
	if err!=nil{
		return  0,err
	}
	c,err:=query.Count()
	return  int64(c),err
}
func(dao *MongDaoImpl) ListByPage(page int,size int,list interface{})(int64,error){
	//if dao.IsWhere{
	dao.Where()
	//}
	s:= dao.Session.DB(dao.Db1).C(dao.Collection1)
	if dao.where==nil{
		dao.where=EmptyWhere
	}
	query:=s.Find(dao.where)
	err:=query.Skip((page-1)*size).Limit(size).All(list)
	if err!=nil{
		return  0,err
	}
	c,err:=query.Count()
	return int64(c),  err
}
//db
func(dao *MongDaoImpl) ExecuteSql(sql string,values ... interface{}){

}
func(dao *MongDaoImpl) QuerySql(sql string,values ... interface{}){

}
func(dao *MongDaoImpl) QueryList(sql string,sqlCount string,page int ,size int,list interface{},values ... interface{})(int64,error){
	return 0,nil
}

func(dao *MongDaoImpl)ExecuteList(sql string,list interface{})error{
	return nil
}
func(dao *MongDaoImpl) ExecuteSqlToInt(sql string,values ... interface{})(int,error){
	return 0, nil
}

