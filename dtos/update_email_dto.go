package dtos

import (
	"encoding/xml"
	"strconv"
	"strings"
)

/**
根据邮箱修改邮箱
*/
type UpdateUserEmailInput struct {
	OperatorLog

	XMLName xml.Name ` json:"-" form:"-"  xml:"request" `

	Id int64 ` json:"-" form:"-"  xml:"-" `

	//邮箱
	Email string `json:"email" form:"email"  xml:"email" validate:"required,lt=20,gt=9"`

	//新邮箱
	NewEmail string ` json:"new_email" form:"new_email"  xml:"new_email" validate:"required,lt=20,gt=9"`

}
/*
id,email,new_email,operator_ip,operator_date
*/
func (user *UpdateUserEmailInput)ToMq() string {
	var b strings.Builder
	b.Grow(200)
	b.WriteString(strconv.FormatInt(user.Id,10))
	b.WriteString(",")
	b.WriteString(user.Email)
	b.WriteString(",")
	b.WriteString(user.NewEmail)
	b.WriteString(",")
	b.WriteString(user.OperatorStringIp)
	b.WriteString(",")
	b.WriteString(strconv.FormatInt(user.OperatorDate,10))
	return  b.String()
}
/*
id,email,new_email,operator_ip,operator_date
*/
func (user *UpdateUserEmailInput)ParseMq(str string)  {
	strs:=strings.Split(str,",")
	if len(strs)==5{
		id,_:=strconv.ParseInt(strs[0],10,64)
		user.Id=id
		user.Email=strs[1]
		user.NewEmail=strs[2]
		user.OperatorStringIp=strs[3]
		date,_:=strconv.ParseInt(strs[4],10,64)
		user.OperatorDate=date
	}
}
/**
根据手机号修改邮箱
*/
type UpdateUserEmailByPhoneInput struct {
	OperatorLog

	XMLName xml.Name ` json:"-" form:"-"  xml:"request" `

	Id int64 ` json:"-" form:"-"  xml:"-" `
	//新邮箱
	Email string `json:"email" form:"email"  xml:"email" `

	//手机号
	Phone string `json:"phone" form:"phone"  xml:"phone" `


}
/*
id,email,phone,operator_ip,operator_date
*/
func (user *UpdateUserEmailByPhoneInput)ToMq() string {
	var b strings.Builder
	b.Grow(200)
	b.WriteString(strconv.FormatInt(user.Id,10))
	b.WriteString(",")
	b.WriteString(user.Email)
	b.WriteString(",")
	b.WriteString(user.Phone)
	b.WriteString(",")
	b.WriteString(user.OperatorStringIp)
	b.WriteString(",")
	b.WriteString(strconv.FormatInt(user.OperatorDate,10))
	return  b.String()
}
/*
id,email,phone,operator_ip,operator_date
*/
func (user *UpdateUserEmailByPhoneInput)ParseMq(str string)  {
	strs:=strings.Split(str,",")
	if len(strs)==5{
		id,_:=strconv.ParseInt(strs[0],10,64)
		user.Id=id
		user.Email=strs[1]
		user.Phone=strs[2]
		user.OperatorStringIp=strs[3]
		date,_:=strconv.ParseInt(strs[4],10,64)
		user.OperatorDate=date
	}
}