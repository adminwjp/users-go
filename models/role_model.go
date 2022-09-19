package models

import "strconv"

//角色
type RoleModel struct {
  
  //角色Id
  Id int64 `orm:"column(id)" gorm:"id;primary_key" json:"id" form:"id"  xml:"id"  bson:"id"  `
  
  //角色名称
  Name string `orm:"column(name)" gorm:"name" json:"name" form:"name"  xml:"name" bson:"name" valid:"Required;Range(2, 10)" validate:"required,lt=10,gt=2" `
  
  //角色描述
  Description string `orm:"column(description)" gorm:"description" json:"description" form:"description"  xml:"description" valid:"Required;Range(0, 500)"  bson:"description" validate:"lt=500" `
  
  //角色父Id
  ParentId int64 `orm:"column(parent_id)" gorm:"parent_id" json:"parent_id" form:"parent_id"  xml:"parent_id"  bson:"parent_id" `

  Created string `orm:"column(created)" gorm:"created" json:"created" form:"-"  xml:"-"  bson:"created"  `

  Updated string `orm:"column(updated)" gorm:"updated" json:"updated" form:"-"  xml:"-"  bson:"updated" `

}
func (*RoleModel) TableName() string {
  return "t_role"
}
func (*RoleModel)GetIdProName() string{
  return  "Id"
}
func (obj *RoleModel)GetId() interface{}{
  return  obj.Id
}
func (obj *RoleModel)GetIdColName() string{
  return  "id"
}
func (*RoleModel)GetDescription() string{
  return  "role"
}
func (*RoleModel)GetDb() string{
  return  "samplesystem"
  //return  "news"
}
func (*RoleModel)GetTable() string{
  return  "role"
}

func (*RoleModel)GetCollection() string{
  return  "role"
}

func (*RoleModel)GetDoc() string{
  return  "role"
}

func (*RoleModel)GetIdName() string{
  return  "id"
}

func (m *RoleModel)GetMId() string{
  return  strconv.FormatInt(m.Id,10)
}
func (m *RoleModel)GetEId() string{
  return  strconv.FormatInt(m.Id,10)
}

func (m *RoleModel)GetParseMId(id interface{}) string{
  i:=id.(int64)
  return  strconv.FormatInt(i,10)
}
func (m *RoleModel)GetParseEId(id interface{}) string{
  i:=id.(int64)
  return  strconv.FormatInt(i,10)
}