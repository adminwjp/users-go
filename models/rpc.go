package models

import "strconv"

type RpcModel struct {
	Id string `orm:"column(id);" gorm:"primary_key" json:"id" form:"id" json:"id" xml:"id"  bson:"id" `
	//go-user_thrift go-user_grpc  go-user_gin go-user_web go-user_socket
	Name string `orm:"column(name);"  json:"name" form:"name" json:"name" xml:"name"  bson:"name" `
	Ip uint64 `orm:"column(ip);"  json:"ip" form:"ip" json:"ip" xml:"ip"  bson:"ip" `
	IpString string `orm:"column(ip_string);"  json:"ip_string" form:"ip_string" json:"ip_string" xml:"ip_string"  bson:"ip_string" `
	Port uint8 `orm:"column(port);"  json:"port" form:"port" json:"port" xml:"port"  bson:"port" `
	Status bool `orm:"column(status);"  json:"status" form:"status" json:"status" xml:"status"  bson:"status" `
}
func (*RpcModel)TableName() string{
	return  "rpc_config"
}
func (*RpcModel)GetIdProName() string{
	return  "Id"
}
func (obj *RpcModel)GetId() interface{}{
	return  obj.Id
}
func (obj *RpcModel)GetIdColName() string{
	return  "id"
}
func (*RpcModel)GetDescription() string{
	return  "rpc_config"
}
func (*RpcModel)GetDb() string{
	return  "samplesystem"
	//return  "news"
}
func (*RpcModel)GetTable() string{
	return  "rpc_config"
}

func (*RpcModel)GetCollection() string{
	return  "rpc_config"
}

func (*RpcModel)GetDoc() string{
	return  "rpc_config"
}

func (*RpcModel)GetIdName() string{
	return  "id"
}
func (m *RpcModel)GetMId() string{
	return  m.Id
}
func (m *RpcModel)GetEId() string{
	return  m.Id
}

func (m *RpcModel)GetParseMId(id interface{}) string{
	i:=id.(int64)
	return  strconv.FormatInt(i,10)
}
func (m *RpcModel)GetParseEId(id interface{}) string{
	i:=id.(int64)
	return  strconv.FormatInt(i,10)
}