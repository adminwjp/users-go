package dao_es_impl

import (
	"context"
	"encoding/json"
	"github.com/adminwjp/users-go/daos"
	"github.com/olivere/elastic/v7"
	"log"
	"reflect"
	"strings"
)

type EsDaoImpl struct {
	daos.DaoImpl
	Client *elastic.Client
	Auto bool
}

func(dao *EsDaoImpl) Wheres(query *elastic.BoolQuery,wheres []*daos.DaoVal) {

	if wheres != nil && len(wheres) > 0 {
		dao.Trim(wheres)
		for _, v := range wheres {
			switch v.Flag {
			case daos.WhereNone:

				break
			case daos.WhereEq:
				query.Must(elastic.NewTermQuery(v.Name,v.Value))
				break
			case daos.WhereLike:
				query.Must(elastic.NewRegexpQuery(v.Name,"/"+v.Value.(string)+"/"))
				break
			case daos.WhereIn:
				reflectValue:=reflect.ValueOf(v.Value)
				switch reflectValue.Kind() {
						case reflect.Slice, reflect.Array:{
							for i := 0; i < reflectValue.Len(); i++ {
								elem:= reflectValue.Index(i)
								query.Must(elastic.NewTermQuery(v.Name,elem.Interface()))
							}
						}
						break
					default:
						break
					}
				break
			case daos.WhereLt:
				query.Must(elastic.NewRangeQuery(v.Name).Lt(v.Value))
				break
			case daos.WhereGt:
				query.Must(elastic.NewRangeQuery(v.Name).Gt(v.Value))
				break
			case daos.WhereLte:
				query.Must(elastic.NewRangeQuery(v.Name).Lte(v.Value))
				break
			case daos.WhereGte:
				query.Must(elastic.NewRangeQuery(v.Name).Gte(v.Value))
				break
			default:
				break

			}
		}
	}
}
func(dao *EsDaoImpl) WheresOr(query *elastic.BoolQuery,wheres []*daos.DaoVal) {

	if wheres != nil && len(wheres) > 0 {
		dao.Trim(wheres)
		for _, v := range wheres {
			switch v.Flag {
			case daos.WhereNone:

				break
			case daos.WhereEq:
				query.Should(elastic.NewTermQuery(v.Name,v.Value))
				break
			case daos.WhereLike:
				query.Should(elastic.NewRegexpQuery(v.Name,"/"+v.Value.(string)+"/"))
				break
			case daos.WhereIn:
				reflectValue:=reflect.ValueOf(v.Value)
				switch reflectValue.Kind() {
				case reflect.Slice, reflect.Array:{
					for i := 0; i < reflectValue.Len(); i++ {
						elem:= reflectValue.Index(i)
						query.Should(elastic.NewTermQuery(v.Name,elem.Interface()))
					}
				}
					break
				default:
					break
				}
				break
			case daos.WhereLt:
				query.Should(elastic.NewRangeQuery(v.Name).Lt(v.Value))
				break
			case daos.WhereGt:
				query.Should(elastic.NewRangeQuery(v.Name).Gt(v.Value))
				break
			case daos.WhereLte:
				query.Should(elastic.NewRangeQuery(v.Name).Lte(v.Value))
				break
			case daos.WhereGte:
				query.Should(elastic.NewRangeQuery(v.Name).Gte(v.Value))
				break
			default:
				break

			}
		}
	}
}
func(dao *EsDaoImpl) Where() elastic.Query {
	query:=elastic.NewBoolQuery()
	dao.Wheres(query,dao.Ands)
	dao.WheresOr(query,dao.Ors)
	return query
}
func(dao *EsDaoImpl) One(obj interface{})error{
	//if dao.IsWhere{
	query:=dao.Where()
	//}
	_,err:= dao.One1(query,obj)
	return err
}
func(dao *EsDaoImpl) One1(query elastic.Query,obj interface{})(string,error){
	s,err:= dao.Client.Search().Index(dao.Doc1).Query(query).From(0).Size(1).Do(context.Background())
	if err!=nil{
		return "",err
	}
	bs,err:=s.Hits.Hits[0].Source.MarshalJSON()
	if err!=nil{
		return "",err
	}

	return  	s.Hits.Hits[0].Id,json.Unmarshal(bs,obj)
}

//http://localhost:9200/role/_doc/iKguMoMBgDx0ERTi364T/
func(dao *EsDaoImpl) GetId(query elastic.Query)(string,error){
	fe:=elastic.NewFetchSourceContext(true)
	fe.Include("_id")
	s,err:= dao.Client.Search().Index(dao.Doc1).Query(query).FetchSourceContext(fe).From(0).Size(1).Do(context.Background())
	if err!=nil{
		return "",err
	}
	return  	s.Hits.Hits[0].Id,err
}
func(dao *EsDaoImpl) Count()(int,error){
	//if dao.IsWhere{
	query:=dao.Where()
	//}
	s,err:= dao.Client.Count().Index(dao.Doc1).Query(query).Do(context.Background())
	log.Printf("es Count:%d",s)
	return int(s), err
}

func(dao *EsDaoImpl) QueryCount(sql string,values ...interface{})(int,error){
	return dao.Count()
}

//insert
func(dao *EsDaoImpl) Add(obj interface{})(int,error){
	r,err:=  dao.Client.Index().Index(dao.Doc1).BodyJson(obj).Do(context.Background())
	if err!=nil{
		return 0, err
	}
	//0 1 1 1
	log.Printf("es add status:%d SeqNo:%d PrimaryTerm:%d Version:%d",r.Status,r.SeqNo,r.PrimaryTerm,r.Version)
	return int(r.Version), err
}

func(dao *EsDaoImpl)ExecuteUpdate(obj interface{})(int,error){
	query:=elastic.NewBoolQuery()
	query.Must(elastic.NewTermQuery(dao.IdName1,dao.Id1))
	id,err:=dao.GetId(query)
	if err!=nil{
		log.Printf("es update get id err:%s",err.Error())
		return 0, err
	}
	r,err:=  dao.Client.Update().Index(dao.Doc1).Id(id).//Id(dao.EId1).
	Doc(obj).Do(context.Background())
	if err!=nil{
		log.Printf("es update id:%s err:%s",id,err.Error())
		return 0, err
	}
	log.Printf("es update status:%d SeqNo:%d PrimaryTerm:%d Version:%d",r.Status,r.SeqNo,r.PrimaryTerm,r.Version)
	return int(r.Version), err
}
func(dao *EsDaoImpl)Delete()(int,error){
	if dao.EId1!=""{
		r,err:=  dao.Client.Delete().Index(dao.Doc1).Id(dao.EId1).Do(context.Background())
		if err!=nil{
			return 0, err
		}
		return r.Status, err
	}
	//if dao.IsWhere{
	query:=dao.Where()
	//}
	r,err:=  dao.Client.DeleteByQuery().Index(dao.Doc1).Query(query).Do(context.Background())
	if err!=nil{
		return 0, err
	}

	return int(r.Deleted), err
}
func(dao *EsDaoImpl) Execute()(int,error){
	var obj map[string]interface{}
	err:=dao.One(&obj)
	if err!=nil{
		log.Printf("es update get err:%s",err.Error())
		return 0, err
	}
	if dao.Updates!=nil&&len(dao.Updates)>0{
		for k, v := range dao.Updates {
			obj[k]=v
		}
	}
	if dao.UpdateIncrs!=nil&&len(dao.UpdateIncrs)>0{
		for k, _ := range dao.UpdateIncrs {
			obj[k]=obj[k].(int64)+1
		}
	}
	query:=elastic.NewBoolQuery()
	query.Must(elastic.NewTermQuery(dao.IdName1,dao.Id1))
	id,err:=dao.GetId(query)
	if err!=nil{
		log.Printf("es update get id err:%s",err.Error())
		return 0, err
	}
	r,err:=  dao.Client.Update().Index(dao.Doc1).Id(id).Doc(obj).Do(context.Background())
	if err!=nil{
		log.Printf("es update id:%s err:%s",id,err.Error())
		return 0, err
	}
	log.Printf("es update status:%d SeqNo:%d PrimaryTerm:%d Version:%d",r.Status,r.SeqNo,r.PrimaryTerm,r.Version)
	return int(r.Version), err
}

func(dao *EsDaoImpl) List(list interface{})(int64,error){
	if dao.Name1!=""{
		names:=strings.Split(dao.Name1,",")
		fetchSourceContext := elastic.NewFetchSourceContext(true)
		for i := 0; i < len(names); i++ {
			fetchSourceContext.Include(names[i])
		}
		//什么情况之前 好好的 难道 不能整合放到一起 拆开? 包问题?
		//es 最近 怎么老是异常
		//fetchSourceContext.Exclude("")

		res,err:=dao.Client.Search(dao.Doc1).Query(fetchSourceContext).Do(context.Background())
		return dao.result(res,err,list)
	}
	//if dao.IsWhere{
	query:=dao.Where()
	//}
	s,err:= dao.Client.Search().Index(dao.Doc1).Query(query).Do(context.Background())

	return dao.result(s,err,list)
}
func (dao *EsDaoImpl) result(s *elastic.SearchResult,err error,list interface{})(int64,error)  {
	if err!=nil{
		return 0,err
	}
	if s.TotalHits()<1{
		return  0,err
	}
	bs:=make([][]byte,s.Hits.TotalHits.Value)
	resultv := reflect.ValueOf(list)
	if resultv.Kind() != reflect.Ptr || resultv.Elem().Kind() != reflect.Slice {
		panic("result argument must be a slice address")
	}
	slicev := resultv.Elem()
	slicev = slicev.Slice(0, slicev.Cap())
	elemt := slicev.Type().Elem()
	i := 0
	for ; i < len(s.Hits.Hits); i++ {
		v:=s.Hits.Hits[i]
		elemp := reflect.New(elemt)
		bs[i],_=v.Source.MarshalJSON()
		val:=elemp.Elem()
		err=json.Unmarshal(bs[i],elemp.Interface())
		slicev = reflect.Append(slicev, val)
		slicev = slicev.Slice(0, slicev.Cap())
	}
	resultv.Elem().Set(slicev.Slice(0, i))
	return  s.Hits.TotalHits.Value,err

	//t:=reflect.TypeOf(all)
	reflectValue:=reflect.ValueOf(list)
	if reflectValue.Kind()==reflect.Ptr||
		reflectValue.Kind()==reflect.Array||
		reflectValue.Kind()==reflect.Slice||
		reflectValue.Kind()==reflect.UnsafePointer{
		reflectValue=reflectValue.Elem()
	}
	count:=s.Hits.TotalHits.Value
	log.Printf("es list count:%d",count)
	switch reflectValue.Kind() {
	case reflect.Slice, reflect.Array:{
		if reflectValue.Type().Kind()==reflect.Ptr{

		}
		value:=reflect.MakeSlice(reflectValue.Type(),int(s.Hits.TotalHits.Value),int(s.Hits.TotalHits.Value))

		for i,v := range s.Hits.Hits {
			bs[i],_=v.Source.MarshalJSON()
			//elem := reflect.New(reflectValue.Type())
			// slice index out of range
			elem:= value.Index(i)
			bu,_:=json.Marshal(elem.Interface())
			log.Printf("es list get json parse list json:%s",string(bu))
			//reflectValue.Slice(0, i+1).Set(elem)
			//reflectValue.Set(value.Slice(0, i))
			log.Printf("es list    json:%s",string(bs[i]))
			//err=json.Unmarshal(bs[i],elem.Interface()) //inptr ex
			err=json.Unmarshal(bs[i],&elem)//null
			//err=json.Unmarshal(bs[i],elem)//inptr ex
			if err!=nil{
				log.Printf("es list json parse list err:%s",err.Error())
				return 0, err
			}
			//bind json fail
			bu,_=json.Marshal(elem.Interface())
			log.Printf("es list json parse get json:%s",string(bu))
			//bind json fail
			value.Index(i).Set(*&elem)
			bu,_=json.Marshal((*&elem).Interface())
			log.Printf("es list json parse get json:%s",string(bu))
		}
		//list=value null
		list=value.Interface()
		list=&value
		list=value.Pointer()
		bu,_:=json.Marshal(list)
		log.Printf("es list json parse list json:%s",string(bu))
		return  count,err

	}
	default:
		return  count,err
	}

	return  count,err
}
func(dao *EsDaoImpl) ListByPage(page int,size int,list interface{})(int64,error){
	//if dao.IsWhere{
	query:=dao.Where()
	//}
	s,err:= dao.Client.Search().Index(dao.Doc1).Query(query).From((page-1)*size).Size(size).Do(context.Background())
	if err!=nil{
		log.Printf("es list page err:%s",err.Error())
	}
	return dao.result(s,err,list)
}
//db
func(dao *EsDaoImpl) ExecuteSql(sql string,values ... interface{}){

}
func(dao *EsDaoImpl) QuerySql(sql string,values ... interface{}){

}
func(dao *EsDaoImpl) QueryList(sql string,sqlCount string,page int ,size int,list interface{},values ... interface{})(int64,error) {
	return 0,nil
}

func(dao *EsDaoImpl)ExecuteList(sql string,list interface{})error{
	return nil
}
func(dao *EsDaoImpl) ExecuteSqlToInt(sql string,values ... interface{})(int,error){
	return 0, nil
}


