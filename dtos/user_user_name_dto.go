package dtos

import (
	"encoding/xml"
	"github.com/adminwjp/users-go/models"
	"strconv"
	"strings"
)

/**
根据用户名注册 或 登录实体
*/
type UserUserNameInput struct {
	OperatorLog

	XMLName xml.Name ` json:"-" form:"-"  xml:"request" `

	Id int64 ` json:"-" form:"-"  xml:"-" `

	//用户名
	UserName string `json:"user_name" form:"user_name"  xml:"user_name" validate:"required,lt=16,gt=2"`

	//密码
	Pwd string ` json:"pwd" form:"pwd"  xml:"pwd" validate:"required,gt=2"`



}
func (input *UserUserNameInput)ToAdminModel() *models.AdminModel {
	var model1 = &models.AdminModel{}
	model1.Pwd=input.Pwd
	model1.RegDate= input.OperatorDate
	model1.RegIp= input.OperatorIp
	model1.UserName=input.UserName
	return  model1
}
func (input *UserUserNameInput)ToUserModel() *models.UserModel {
	var model1 = &models.UserModel{}
	model1.Pwd=input.Pwd
	model1.RegDate= input.OperatorDate
	model1.RegIp= input.OperatorIp
	model1.UserName=input.UserName
	return  model1
}
/*
id,user_name,pwd,operator_ip,operator_date
*/
func (user *UserUserNameInput)ToMq() string {
	var b strings.Builder
	b.Grow(200)
	b.WriteString(strconv.FormatInt(user.Id,10))
	b.WriteString(",")
	b.WriteString(user.UserName)
	b.WriteString(",")
	b.WriteString(user.Pwd)
	b.WriteString(",")
	b.WriteString(user.OperatorStringIp)
	b.WriteString(",")
	b.WriteString(strconv.FormatInt(user.OperatorDate,10))
	return  b.String()
}
/*
id,user_name,pwd,operator_ip,operator_date
*/
func (user *UserUserNameInput)ParseMq(str string)  {
	strs:=strings.Split(str,",")
	if len(strs)==5{
		id,_:=strconv.ParseInt(strs[0],10,64)
		user.Id=id
		user.UserName=strs[1]
		user.Pwd=strs[2]
		user.OperatorStringIp=strs[3]
		date,_:=strconv.ParseInt(strs[4],10,64)
		user.OperatorDate=date
	}
}