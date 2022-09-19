package tests

import "encoding/xml"

type UserTest struct {
	XMLName xml.Name ` json:"-" form:"-"  xml:"request" `
	//用户id
	Id int64 `orm:"column(id)" gorm:"id" json:"id" form:"id"  xml:"id" `

	//手机号
	Name string `orm:"column(name)" gorm:"name" json:"name" form:"name"  xml:"name" `

	Names []string `orm:"-" gorm:"-" json:"s" form:"names"  xml:"name" `
}
type UserXmlTest struct {
	XMLName xml.Name ` json:"-" form:"-"  xml:"request" `

	//用户id
	Id int64 `orm:"column(id)" gorm:"id" json:"id" form:"id"  xml:"id" `

	//手机号
	Name string `orm:"column(name)" gorm:"name" json:"name" form:"name"  xml:"name" `

	Names []UserNameXmlTest ` xml:"name" `
}
type UserNameXmlTest struct {

	//手机号
	Name string `xml:"name" `

}
type UserLogTest struct {
	//用户id
	Id int64 `orm:"column(id)" gorm:"id" json:"id" form:"id"  xml:"id" `
	//用户id
	UserId int64 `orm:"column(user_id)" gorm:"user_id" json:"user_id" form:"user_id"  xml:"user_id" `

}
func (UserLogTest) TableName() string {
	return "t_user_log_test"
}
func (UserTest) TableName() string {
	return "t_user_test"
}