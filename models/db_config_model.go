package models

type DbConfigModel struct {
	Id int64 `orm:"column(id);" gorm:"PRIMARY_KEY" json:"id" form:"id" json:"id" xml:"id" bson:"id" `
	DefaultDb string `orm:"column(default_db);default('');null" json:"default_db" form:"default_db" json:"default_db" xml:"default_db" bson:"default_db"`
	Db string `orm:"column(db);default('');null" json:"db" form:"db" json:"db" xml:"db" bson:"db"`
	Ip uint64 `orm:"column(ip);default(0);null" json:"ip" form:"ip" json:"ip" xml:"ip" bson:"ip"`
	IpString string `orm:"column(ip_string);default('');null" json:"ip_string" form:"ip_string" json:"ip_string" xml:"ip_string" bson:"ip_string"`
	Port uint `orm:"column(port);default(0);null" json:"port" form:"port" json:"port" xml:"port" bson:"port"`
	Count uint64 `orm:"column(count);default(0);null" json:"count" form:"count" json:"count" xml:"count" bson:"count"`
	MaxCount uint64 `orm:"column(max_count);default(0);null" json:"max_count" form:"max_count" json:"max_count" xml:"max_count" bson:"max_count"`
	MaxSpace uint64 `orm:"column(max_space);default(0);null" json:"max_space" form:"max_space" json:"max_space" xml:"max_space" bson:"max_space"`
	Space uint64 `orm:"column(space);default(0);null" json:"space" form:"space" json:"space" xml:"space" bson:"space"`
	StartId uint64 `orm:"column(start_id);default(0);null" json:"start_id" form:"start_id" json:"start_id" xml:"start_id" bson:"start_id"`
	EndId uint64 `orm:"column(end_id);default(0);null" json:"end_id" form:"end_id" json:"end_id" xml:"end_id" bson:"end_id"`

	Db1 *DbConfigModel  `orm:"column(parent_id);rel(fk);null;on_delete(do_nothing)" json:"-" form:"-" json:"-" xml:"-" bson:"-"`
	Dbs []DbConfigModel  `orm:"column(parent_id);reverse(many);null" json:"-" form:"-" json:"-" xml:"-" bson:"-"`
	Tables1 []TableConfigModel  `orm:"column(parent_id);reverse(many);null" json:"-" form:"-" json:"-" xml:"-" bson:"-"`
	Table string ` orm:"column(table);default('');null" json:"table" form:"table" json:"table" xml:"table" bson:"table"`
	Tables []string ` orm:"-" json:"tables" form:"tables" json:"tables" xml:"tables" bson:"-"`

}

func (*DbConfigModel)GetIdProName() string{
	return  "Id"
}
func (obj *DbConfigModel)GetId() interface{}{
	return  obj.Id
}
func (obj *DbConfigModel)GetIdColName() string{
	return  "id"
}
func (*DbConfigModel)GetDescription() string{
	return  "db_config"
}
func (*DbConfigModel)GetDb() string{
	return  "samplesystem"
	//return  "news"
}
func (*DbConfigModel)GetTable() string{
	return  "db_config"
}