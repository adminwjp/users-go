package models

import "strconv"

type ConfigModel struct {
	Id string `orm:"column(id);" gorm:"primary_key" json:"id" form:"id" json:"id" xml:"id"  bson:"id" `
	//go-user_thrift go-user_grpc
	Name string `orm:"column(name);"  json:"name" form:"name" json:"name" xml:"name"  bson:"name" `
	Ip uint64 `orm:"column(ip);"  json:"ip" form:"ip" json:"ip" xml:"ip"  bson:"ip" `
	IpString string `orm:"column(ip_string);"  json:"ip_string" form:"ip_string" json:"ip_string" xml:"ip_string"  bson:"ip_string" `
	Data string `orm:"column(data);"  json:"data" form:"data" json:"data" xml:"data"  bson:"data" `
	//xml yaml json string
	Flag string `orm:"column(Flag);"  json:"Flag" form:"Flag" json:"Flag" xml:"Flag"  bson:"Flag" `
}
func (*ConfigModel)TableName() string{
	return  "config"
}
func (*ConfigModel)GetIdProName() string{
	return  "Id"
}
func (obj *ConfigModel)GetId() interface{}{
	return  obj.Id
}
func (obj *ConfigModel)GetIdColName() string{
	return  "id"
}
func (*ConfigModel)GetDescription() string{
	return  "config"
}
func (*ConfigModel)GetDb() string{
	return  "samplesystem"
	//return  "news"
}
func (*ConfigModel)GetTable() string{
	return  "config"
}

func (*ConfigModel)GetCollection() string{
	return  "config"
}

func (*ConfigModel)GetDoc() string{
	return  "config"
}

func (*ConfigModel)GetIdName() string{
	return  "id"
}
func (m *ConfigModel)GetMId() string{
	return  m.Id
}
func (m *ConfigModel)GetEId() string{
	return  m.Id
}

func (m *ConfigModel)GetParseMId(id interface{}) string{
	i:=id.(int64)
	return  strconv.FormatInt(i,10)
}
func (m *ConfigModel)GetParseEId(id interface{}) string{
	i:=id.(int64)
	return  strconv.FormatInt(i,10)
}