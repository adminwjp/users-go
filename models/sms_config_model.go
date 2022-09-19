package models

import "strconv"

type SmsConfigModel struct {
	//短信id
	Id int64 `orm:"column(id)" gorm:"id;primary_key" json:"id" form:"id"  xml:"id"  bson:"id" `

	//短信
	Sms string `orm:"column(sms)" gorm:"sms" json:"sms" form:"sms"  xml:"sms"  bson:"sms" validate:"required,lt=16,gt=10" `

	//短信凭证
	Scrept string `orm:"column(scrept)" gorm:"scrept" json:"scrept" form:"scrept"  xml:"scrept" bson:"scrept" validate:"required"  `

	//短信密码
	Pwd string `orm:"column(pwd)" gorm:"pwd" json:"pwd" form:"pwd"  xml:"pwd"  bson:"pwd" validate:"required" `

	Flag string `orm:"column(flag)" gorm:"flag" json:"flag" form:"flag"  xml:"flag"  bson:"flag" validate:"required,lt=10,gt=2" `

	Created string `orm:"column(created)" gorm:"created" json:"created" form:"-"  xml:"-"  bson:"created"  `

	Updated string `orm:"column(updated)" gorm:"updated" json:"updated" form:"-"  xml:"-"  bson:"updated" `
}
func (*SmsConfigModel)TableName() string{
	return  "sms_config"
}
func (*SmsConfigModel)GetIdProName() string{
	return  "Id"
}
func (obj *SmsConfigModel)GetId() interface{}{
	return  obj.Id
}
func (obj *SmsConfigModel)GetIdColName() string{
	return  "id"
}
func (*SmsConfigModel)GetDescription() string{
	return  "sms_config"
}
func (*SmsConfigModel)GetDb() string{
	return  "samplesystem"
	//return  "news"
}
func (*SmsConfigModel)GetTable() string{
	return  "sms_config"
}
func (*SmsConfigModel)GetCollection() string{
	return  "sms"
}

func (*SmsConfigModel)GetDoc() string{
	return  "sms"
}

func (*SmsConfigModel)GetIdName() string{
	return  "id"
}
func (m *SmsConfigModel)GetMId() string{
	return  strconv.FormatInt(m.Id,10)
}
func (m *SmsConfigModel)GetEId() string{
	return  strconv.FormatInt(m.Id,10)
}

func (m *SmsConfigModel)GetParseMId(id interface{}) string{
	i:=id.(int64)
	return  strconv.FormatInt(i,10)
}
func (m *SmsConfigModel)GetParseEId(id interface{}) string{
	i:=id.(int64)
	return  strconv.FormatInt(i,10)
}