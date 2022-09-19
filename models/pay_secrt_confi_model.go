package models

import "strconv"

type PaySecrtConfigModel struct {
	Id int64 `orm:"column(id)" gorm:"id;primary_key" json:"id" form:"id"  xml:"id"  bson:"id" `
}

func (*PaySecrtConfigModel)TableName() string{
	return  "pay_config"
}
func (*PaySecrtConfigModel)GetIdProName() string{
	return  "Id"
}
func (obj *PaySecrtConfigModel)GetId() interface{}{
	return  obj.Id
}
func (obj *PaySecrtConfigModel)GetIdColName() string{
	return  "id"
}
func (*PaySecrtConfigModel)GetDescription() string{
	return  "pay_config"
}
func (*PaySecrtConfigModel)GetDb() string{
	return  "samplesystem"
	//return  "news"
}
func (*PaySecrtConfigModel)GetTable() string{
	return  "pay_config"
}

func (*PaySecrtConfigModel)GetCollection() string{
	return  "pay"
}

func (*PaySecrtConfigModel)GetDoc() string{
	return  "pay"
}

func (*PaySecrtConfigModel)GetIdName() string{
	return  "id"
}

func (m *PaySecrtConfigModel)GetMId() string{
	return  strconv.FormatInt(m.Id,10)
}
func (m *PaySecrtConfigModel)GetEId() string{
	return  strconv.FormatInt(m.Id,10)
}

func (m *PaySecrtConfigModel)GetParseMId(id interface{}) string{
	i:=id.(int64)
	return  strconv.FormatInt(i,10)
}
func (m *PaySecrtConfigModel)GetParseEId(id interface{}) string{
	i:=id.(int64)
	return  strconv.FormatInt(i,10)
}