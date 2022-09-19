package dtos

import (
	"encoding/xml"
	"github.com/adminwjp/users-go/models"
	"strconv"
	"strings"
)

/**
根据邮箱、注册 或 登录实体
*/
type UserEmailInput struct {
	OperatorLog

	XMLName xml.Name ` json:"-" form:"-"  xml:"request" `

	Id int64 ` json:"-" form:"-"  xml:"-" `

	//邮箱
	Email string `json:"email" form:"email"  xml:"email" validate:"required,lt=20,gt=9"`

	//密码
	Pwd string ` json:"pwd" form:"pwd"  xml:"pwd" validate:"required,gt=2"`



}

func (input *UserEmailInput)ToUserModel()*models.UserModel  {
	var m = &models.UserModel{}
	m.Pwd=input.Pwd
	m.RegDate= input.OperatorDate
	m.RegIp= input.OperatorIp
	m.Email=input.Email
	return  m
}
func (input *UserEmailInput)ToAdminModel()*models.AdminModel  {
	var m = &models.AdminModel{}
	m.Pwd=input.Pwd
	m.RegDate= input.OperatorDate
	m.RegIp= input.OperatorIp
	m.Email=input.Email
	return  m
}
/*
id,email,pwd,operator_ip,operator_date
*/
func (user *UserEmailInput)ToMq() string {
	var b strings.Builder
	b.Grow(200)
	b.WriteString(strconv.FormatInt(user.Id,10))
	b.WriteString(",")
	b.WriteString(user.Email)
	b.WriteString(",")
	b.WriteString(user.Pwd)
	b.WriteString(",")
	b.WriteString(user.OperatorStringIp)
	b.WriteString(",")
	b.WriteString(strconv.FormatInt(user.OperatorDate,10))
	return  b.String()
}
/*
id,email,pwd,operator_ip,operator_date
*/
func (user *UserEmailInput)ParseMq(str string)  {
	strs:=strings.Split(str,",")
	if len(strs)==5{
		id,_:=strconv.ParseInt(strs[0],10,64)
		user.Id=id
		user.Email=strs[1]
		user.Pwd=strs[2]
		user.OperatorStringIp=strs[3]
		date,_:=strconv.ParseInt(strs[4],10,64)
		user.OperatorDate=date
	}
}