package daos

import (
	data "github.com/adminwjp/infrastructure-go/datas"
	"github.com/adminwjp/users-go/datas"
	"reflect"
	"strconv"
	"strings"
)

type WhereFlag int

const  (
	WhereNone WhereFlag=iota
	WhereEq
	WhereLike
	WhereIn
	WhereLt
	WhereGt
	WhereLte
	WhereGte
)

type DaoVal struct{
	Name string
	Value interface{}
	Flag WhereFlag
}


type DaoImpl struct {
	Ors []*DaoVal
	Ands []*DaoVal

	AndBetweens map[string][]interface{}
	OrBetweens map[string][]interface{}

	Updates map[string]interface{}
	UpdateIncrs map[string]bool

	Model1 interface{}

	Doc1 string

	Db1 string
	Collection1 string

	Table1 string
	View1 string

	IsWhere1 bool

	Id1 interface{}

	MId1 string

	EId1 string

	IdName1 string

	Name1 string
}
func(dao *DaoImpl) Trim(wheres []*DaoVal){
	e:=false
	for _, v := range wheres {
		if e{
			v.Name=strings.Split(v.Name,".")[1]
		}
		if strings.Contains(v.Name,"."){
			e=true
			v.Name=strings.Split(v.Name,".")[1]
		}else{
			break
		}

	}
}
//where
func (dao *DaoImpl)Eq(name string,val interface{}){
	dao.Ands=append(dao.Ands,&DaoVal{Name: name,Value: val,Flag: WhereEq})
}
func (dao *DaoImpl)Like(name string,val string){
	dao.Ands=append(dao.Ands,&DaoVal{Name: name,Value: val,Flag: WhereLike})
}
func (dao *DaoImpl)In(name string,val interface{}){
	dao.Ands=append(dao.Ands,&DaoVal{Name: name,Value: val,Flag: WhereIn})
}
func (dao *DaoImpl)Lt(name string,val interface{}){
	dao.Ands=append(dao.Ands,&DaoVal{Name: name,Value: val,Flag: WhereLt})
}
func (dao *DaoImpl)Gt(name string,val interface{}){
	dao.Ands=append(dao.Ands,&DaoVal{Name: name,Value: val,Flag: WhereGt})
}
func (dao *DaoImpl)Lte(name string,val interface{}){
	dao.Ands=append(dao.Ands,&DaoVal{Name: name,Value: val,Flag: WhereLte})
}
func (dao *DaoImpl)Gte(name string,val interface{}){
	dao.Ands=append(dao.Ands,&DaoVal{Name: name,Value: val,Flag: WhereGte})
}
func (dao *DaoImpl)Between(name string,val interface{},val1 interface{}){
	dao.AndBetweens[name]=[]interface{}{val,val1}
}

func (dao *DaoImpl)OrEq(name string,val interface{}){
	dao.Ors=append(dao.Ors,&DaoVal{Name: name,Value: val,Flag: WhereEq})
}
func (dao *DaoImpl)OrLike(name string,val string){
	dao.Ors=append(dao.Ors,&DaoVal{Name: name,Value: val,Flag: WhereLike})
}
func (dao *DaoImpl)OrIn(name string,val interface{}){
	dao.Ors=append(dao.Ors,&DaoVal{Name: name,Value: val,Flag: WhereIn})
}
func (dao *DaoImpl)OrLt(name string,val interface{}){
	dao.Ors=append(dao.Ors,&DaoVal{Name: name,Value: val,Flag: WhereLt})
}
func (dao *DaoImpl)OrGt(name string,val interface{}){
	dao.Ors=append(dao.Ors,&DaoVal{Name: name,Value: val,Flag: WhereGt})
}
func (dao *DaoImpl)OrLte(name string,val interface{}){
	dao.Ors=append(dao.Ors,&DaoVal{Name: name,Value: val,Flag: WhereLte})
}
func (dao *DaoImpl)OrGte(name string,val interface{}){
	dao.Ors=append(dao.Ors,&DaoVal{Name: name,Value: val,Flag: WhereGte})
}
func (dao *DaoImpl)OrBetween(name string,val interface{},val1 interface{}){
	dao.OrBetweens[name]=[]interface{}{val,val1}
}

//es
func (dao *DaoImpl)Doc(doc string){
	dao.Doc1=doc
}

//mong
func (dao *DaoImpl)Collection(name string){
	dao.Collection1=name
}
func  (dao *DaoImpl)Db(name string)  {
	dao.Db1=name
}
//db
func (dao *DaoImpl)Table(name string){
	dao.Table1=name
}

func (dao *DaoImpl)Model(m interface{}){
	dao.Model1=m
}

func (dao *DaoImpl)View(name string){
	dao.View1=name
}
func(dao *DaoImpl) Id(val interface{}) {
	dao.Id1 = val
}

func(dao *DaoImpl) MId(val string) {
	dao.MId1 = val
}
func(dao *DaoImpl) EId(val string) {
	dao.EId1 = val
}

func(dao *DaoImpl) IdName(name string) {
	dao.IdName1 = name
}
//update
func(dao *DaoImpl) Update(name string,val interface{}) {
	dao.Updates[name] = val
}
func(dao *DaoImpl) UpdateIncr(name string) {
	dao.UpdateIncrs[name] = true
}

func(dao *DaoImpl) IsWhere(val bool) {
	dao.IsWhere1 = val
}
//column json bson name,name1...
func(dao *DaoImpl) Get(name string){
	dao.Name1=name
}
func(dao *DaoImpl) WhereStrings(and bool,wheres []*DaoVal)string {
	var w=" or "
	if and{
		w=" and "
	}
	var wh=""
	if wheres != nil && len(wheres) > 0 {
		for _, v := range wheres {
			v1:=reflect.ValueOf(v.Value)
			switch v1.Type() {
			case reflect.TypeOf(int64(1)):
				wh+=v.Name+"="+v1.String()+w
				break
			case reflect.TypeOf(1):
				wh+=v.Name+"='"+v1.String()+"'"+w
				break
			case reflect.TypeOf(true):
				if v1.Bool(){
					wh+=v.Name+"=1"+w
				}else{
					wh+=v.Name+"=0"+w
				}
				break
			}
		}
	}
	return  wh
}

func(dao *DaoImpl) WhereStringsInt64(wh string, w string,v1 reflect.Value,v *DaoVal)string {
	val:=v1.String()
	return dao.WhereStringsString(wh,w,"",val,v1,v)
}
func(dao *DaoImpl) WhereStringsBool(wh string, w string,v1 reflect.Value,v *DaoVal)string {
	val:="0"
	if v1.Bool(){
		val="1"
	}
	return dao.WhereStringsString(wh,w,"",val,v1,v)
}
func(dao *DaoImpl) WhereStringsString(wh string, w string,c string,val string,v1 reflect.Value,v *DaoVal)string {
	switch v.Flag {
	case WhereNone:

		break
	case WhereEq:
		wh += v.Name + "="+ c+  val + c+   w
		break
	case WhereLike:
		wh += v.Name + "="+ c+"%" + val + "%"+ c+ w
		break
	case WhereIn:
		var ins = ""
		//value type
		reflectValue := reflect.ValueOf(v.Value)
		switch reflectValue.Kind() {
		case reflect.Slice, reflect.Array:
			{
				for i := 0; i < reflectValue.Len(); i++ {
					elem := reflectValue.Index(i)
					ins += c + elem.String() + c+","
				}
			}
			break
		default:
			break
		}
		ins = strings.TrimSuffix(ins, ",")
		wh += v.Name + " in('%" + ins + ")%'" + w
		break
	case WhereLt:
		wh += v.Name + "<" +c+ v1.String() +c+  w
		break
	case WhereGt:
		wh += v.Name + ">" +c+ v1.String() +c+ w
		break
	case WhereLte:
		wh += v.Name + "<=" +c+ v1.String() +c + w
		break

	case WhereGte:
		wh += v.Name + ">=" +c+ v1.String() +c+  w
		break
	default:
		break

	}
	return  wh
}
type CrudDaoImpl struct {
	Data func()Dao
	m MDao
	db string
}
func(dao *CrudDaoImpl)Db(name string){
	dao.db=name
}
func(dao *CrudDaoImpl)Add(obj MDao)(int,error){
	da:=dao.Data()
	dao.init1(da,obj)
	return  da.Add(obj)
}
func (dao *CrudDaoImpl)init(da Dao)  {
	da.Doc(dao.m.GetDoc())
	da.Collection(dao.m.GetCollection())
	da.Db(MDb)
	da.Model(dao.m)

}
func (dao *CrudDaoImpl)init1(da Dao,obj MDao)  {
	dao.init(da)
	da.Id(obj.GetId())
	da.MId(obj.GetMId())
	da.EId(obj.GetEId())
	da.IdName(obj.GetIdName())
}
func(dao *CrudDaoImpl)Update(obj MDao)(int,error){
	da:=dao.Data()
	dao.init1(da,obj)

	return  da.ExecuteUpdate(obj)
}
func(dao *CrudDaoImpl)Delete(id interface{})(int,error){
	da:=dao.Data()
	dao.init(da)

	da.Id(id)
	da.MId(dao.m.GetParseMId(id))
	da.EId(dao.m.GetParseEId(id))
	da.IdName(dao.m.GetIdName())
	if datas.GlobalConfig.DataFlag== data.DataDb{
		return da.ExecuteSqlToInt("delete from "+dao.m.TableName()+" where "+dao.m.GetIdName()+" = ?",id)
	}
	return  da.Delete()
}
func(dao *CrudDaoImpl)DeleteBatch(ids interface{})(int,error){
	da:=dao.Data()
	dao.init(da)
	da.In(dao.m.GetIdName(),ids)
	if datas.GlobalConfig.DataFlag== data.DataDb{
		iids,e:=ids.([]int64)
		if e{
			s:=""
			for _, v := range iids {
				s+=strconv.FormatInt(v,10)+","
			}
			s=strings.TrimRight(s,",")
			return da.ExecuteSqlToInt("delete from "+dao.m.TableName()+" where "+dao.m.GetIdName()+" in (?)",s)
		}
		return 0, nil
	}
	return  da.Delete()
}
func(dao *CrudDaoImpl)List(list interface{})(int64,error) {
	da:=dao.Data()
	dao.init(da)
	return da.List(list)
}
func(dao *CrudDaoImpl)ListByPage(page int,size int,list interface{})(int64,error) {
	da:=dao.Data()
	dao.init(da)
	return da.ListByPage(page,size,list)
}
func(dao *CrudDaoImpl)Model(m MDao){
	dao.m=m
}
type MDao interface {
	TableName() string
	GetCollection() string
	GetDoc() string
	GetIdName() string
	GetId() interface{}
	GetMId() string
	GetEId() string
	GetParseMId(id interface{}) string
	GetParseEId(id interface{}) string
}
type CrudDao interface{
	Db(name string)
	Add(obj MDao)(int,error)
	Update(obj MDao)(int,error)
	Delete(id interface{})(int,error)
	DeleteBatch(ids interface{})(int,error)
	List(list interface{})(int64,error)
	ListByPage(page int,size int,list interface{})(int64,error)
	Model(m MDao)
}
//实现方式不同 接口 不然写得累 统一写法
//复杂度变高了
type Dao interface {

	//where
	Eq(name string,val interface{})
	Like(name string,val string)
	In(name string,val interface{})
	Lt(name string,val interface{})
	Gt(name string,val interface{})
	Lte(name string,val interface{})
	Gte(name string,val interface{})
	Between(name string,val interface{},val1 interface{})

	OrEq(name string,val interface{})
	OrLike(name string,val string)
	OrIn(name string,val interface{})
	OrLt(name string,val interface{})
	OrGt(name string,val interface{})
	OrLte(name string,val interface{})
	OrGte(name string,val interface{})

	OrBetween(name string,val interface{},val1 interface{})

	//es
	Doc(doc string)

	//mong
	Collection(name string)

	Db(name string)
	//db
	Table(name string)
	Model(m interface{})
	View(name string)

	Id(val interface{})
	MId(val string)
	EId(val string)
	IdName(name string)

	//update
	Update(name string,val interface{})

	UpdateIncr(name string)

	IsWhere(isWhere bool)

	One(obj interface{})error

	Count()(int,error)
	QueryCount(sql string,values ...interface{})(int,error)

	//column json bson name,name1...
	Get(name string)

	//insert
	Add(obj interface{})(int,error)

	//update
	ExecuteUpdate(obj interface{})(int,error)

	Delete()(int,error)
	//add insert delete
	Execute()(int,error)

	List(list interface{})(int64,error)

	ListByPage(page int,size int,list interface{})(int64,error)

	//db
	ExecuteSql(sql string,values ... interface{})
	QuerySql(sql string,values ... interface{})

   QueryList(sql string,sqlCount string,page int ,size int,list interface{},values ... interface{})(int64,error)

	ExecuteSqlToInt(sql string,values ... interface{})(int,error)

	ExecuteList(sql string,list interface{})error
}
