package models

import "strconv"

type TableConfigModel struct {

	Id int64 `orm:"column(id);" gorm:"id;primary_key" json:"id" form:"id" json:"id" xml:"id"  bson:"id" `
	DefaultDb string `orm:"column(default_db);default('');null" json:"default_db" form:"default_db" json:"default_db" xml:"default_db" bson:"default_db"  `
	Db string `orm:"column(db);default('');null" json:"db" form:"db" json:"db" xml:"db"  bson:"db" `
	Ip uint64 `orm:"column(ip);default(0);null" json:"ip" form:"ip" json:"ip" xml:"ip"  bson:"ip" `
	IpString string `orm:"column(ip_string);default('');null" json:"ip_string" form:"ip_string" json:"ip_string" xml:"ip_string" bson:"ip_string"  `
	Port uint `orm:"column(port);default(0);null" json:"port" form:"port" json:"port" xml:"port"  bson:"port" `
	Count uint64 `orm:"column(count);default(0);null" json:"count" form:"count" json:"count" xml:"count" bson:"count"  `
	MaxCount uint64 `orm:"column(max_count);default(0);null" json:"max_count" form:"max_count" json:"max_count" xml:"max_count"  bson:"max_count" `
	MaxSpace uint64 `orm:"column(max_space);default(0);null" json:"max_space" form:"max_space" json:"max_space" xml:"max_space"  bson:"max_space" `
	Space uint64 `orm:"column(space);default(0);null" json:"space" form:"space" json:"space" xml:"space"  bson:"space" `
	StartId uint64 `orm:"column(start_id);default(0);null" json:"start_id" form:"start_id" json:"start_id" xml:"start_id" bson:"start_id"  `
	EndId uint64 `orm:"column(end_id);default(0);null" json:"end_id" form:"end_id" json:"end_id" xml:"end_id" bson:"end_id"  `

	Db1 *DbConfigModel  `orm:"column(db_id);rel(fk);null;on_delete(do_nothing)" json:"-" form:"-" json:"-" xml:"-"  bson:"-" `
	DefaultTable string ` orm:"column(default_table);default('');null" json:"default_table" form:"default_table" json:"default_table" xml:"default_table" bson:"default_table"  `
	Table string ` orm:"column(table);default('');null" json:"table" form:"table" json:"table" xml:"table" bson:"table"  `
	Ids string ` orm:"column(ids);default('');null" json:"ids" form:"ids" json:"ids" xml:"ids" bson:"ids"  `
	DeleteIds string ` orm:"column(default_ids);default('');null" json:"default_ids" form:"default_ids" json:"default_ids" xml:"default_ids"  bson:"default_ids" `
}
func (*TableConfigModel)TableName() string{
	return  "table_config"
}
func (*TableConfigModel)GetIdProName() string{
	return  "Id"
}
func (obj *TableConfigModel)GetId() interface{}{
	return  obj.Id
}
func (obj *TableConfigModel)GetIdColName() string{
	return  "id"
}
func (*TableConfigModel)GetDescription() string{
	return  "table_config"
}
func (*TableConfigModel)GetDb() string{
	return  "samplesystem"
	//return  "news"
}
func (*TableConfigModel)GetTable() string{
	return  "table_config"
}
func (*TableConfigModel)GetCollection() string{
	return  "table"
}

func (*TableConfigModel)GetDoc() string{
	return  "table"
}

func (*TableConfigModel)GetIdName() string{
	return  "id"
}

func (m *TableConfigModel)GetMId() string{
	return  strconv.FormatInt(m.Id,10)
}
func (m *TableConfigModel)GetEId() string{
	return  strconv.FormatInt(m.Id,10)
}

func (m *TableConfigModel)GetParseMId(id interface{}) string{
	i:=id.(int64)
	return  strconv.FormatInt(i,10)
}
func (m *TableConfigModel)GetParseEId(id interface{}) string{
	i:=id.(int64)
	return  strconv.FormatInt(i,10)
}