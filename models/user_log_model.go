package models

type AdminLogModel struct{
  //BaseUserLogModel

  //用户日志id
  Id int64 `orm:"column(id)" gorm:"id;primary_key" json:"id" form:"id"  xml:"id" bson:"id" `


  //操作数据
  Data string `orm:"column(data)" gorm:"data" json:"data" form:"data"  xml:"data" bson:"id"  `

  //操作标识
  Flag string `orm:"column(flag)" gorm:"flag" json:"flag" form:"flag"  xml:"flag" bson:"id"  `

  //操作消息
  Msg string `orm:"column(msg)" gorm:"msg" json:"msg" form:"msg"  xml:"msg" bson:"id"  `

  //用户Id
  UserId int64 `orm:"column(user_id)" gorm:"user_id" json:"user_id" form:"user_id"  xml:"user_id" bson:"id"  `

  //操作Ip
  OperatorIp int64 `orm:"column(operator_ip)" gorm:"operator_ip" json:"-" form:"-"  xml:"-" bson:"operator_ip"  `

  //操作时间
  OperatorDate int64 `orm:"column(operator_date)" gorm:"operator_date" json:"-" form:"-"  xml:"-" bson:"operator_date"  `

}
//mong es 不要这样
type UserLogModel struct{
 // BaseUserLogModel

  //用户日志id
  Id int64 `orm:"column(id)" gorm:"id;primary_key" json:"id" form:"id"  xml:"id" `


  //操作数据
  Data string `orm:"column(data)" gorm:"data" json:"data" form:"data"  xml:"data" `

  //操作标识
  Flag string `orm:"column(flag)" gorm:"flag" json:"flag" form:"flag"  xml:"flag" `

  //操作消息
  Msg string `orm:"column(msg)" gorm:"msg" json:"msg" form:"msg"  xml:"msg" `

  //用户Id
  UserId int64 `orm:"column(user_id)" gorm:"user_id" json:"user_id" form:"user_id"  xml:"user_id" `

  //操作Ip
  OperatorIp int64 `orm:"column(operator_ip)" gorm:"operator_ip" json:"operator_ip" form:"operator_ip"  xml:"operator_ip" `

  //操作时间
  OperatorDate int64 `orm:"column(operator_date)" gorm:"operator_date" json:"operator_date" form:"operator_date"  xml:"operator_date" `

}
//用户日志
type BaseUserLogModel struct {

}
func (UserLogModel) TableName() string {
  return "t_user_log"
}
func (AdminLogModel) TableName() string {
  return "t_admin_log"
}