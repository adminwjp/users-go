package models

import "strconv"

//邮箱
type EmailConfigModel struct {
  
  //邮箱id
  Id int64 `orm:"column(id)" gorm:"id;primary_key" json:"id" form:"id"  xml:"id" bson:"id"`
  
  //邮箱
  Email string `orm:"column(email)" gorm:"email" json:"email" form:"email"  xml:"email" bson:"email"`
  
  //邮箱凭证
  Scrept string `orm:"column(scrept)" gorm:"scrept" json:"scrept" form:"scrept"  xml:"scrept" bson:"scrept"`
  
  //邮箱密码
  Pwd string `orm:"column(pwd)" gorm:"pwd" json:"pwd" form:"pwd"  xml:"pwd" bson:"pwd"`

  Flag string `orm:"column(flag)" gorm:"flag" json:"flag" form:"flag"  xml:"flag" bson:"flag"`
}
//gorm
func (*EmailConfigModel) TableName() string {
  return "t_email_config"
}
func (*EmailConfigModel)GetIdProName() string{
  return  "Id"
}
func (obj *EmailConfigModel)GetId() interface{}{
  return  obj.Id
}
func (obj *EmailConfigModel)GetIdColName() string{
  return  "id"
}
func (*EmailConfigModel)GetDescription() string{
  return  "email_config"
}
func (*EmailConfigModel)GetDb() string{
  return  "samplesystem"
  //return  "news"
}
func (*EmailConfigModel)GetTable() string{
  return  "email_config"
}

func (*EmailConfigModel)GetCollection() string{
  return  "email"
}

func (*EmailConfigModel)GetDoc() string{
  return  "email"
}

func (*EmailConfigModel)GetIdName() string{
  return  "id"
}
func (m *EmailConfigModel)GetMId() string{
  return  strconv.FormatInt(m.Id,10)
}
func (m *EmailConfigModel)GetEId() string{
  return  strconv.FormatInt(m.Id,10)
}

func (m *EmailConfigModel)GetParseMId(id interface{}) string{
  i:=id.(int64)
  return  strconv.FormatInt(i,10)
}
func (m *EmailConfigModel)GetParseEId(id interface{}) string{
  i:=id.(int64)
  return  strconv.FormatInt(i,10)
}