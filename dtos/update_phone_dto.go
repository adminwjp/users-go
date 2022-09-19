package dtos

import (
	"encoding/xml"
	"strconv"
	"strings"
)

/**
根据手机号修改手机号
*/
type UpdateUserPhoneInput struct {
	OperatorLog

	XMLName xml.Name ` json:"-" form:"-"  xml:"request" `

	Id int64 ` json:"-" form:"-"  xml:"-" `


	//手机号
	Phone string `json:"phone" form:"phone"  xml:"phone" `

	//新手机号
	NewPhone string ` json:"new_phone" form:"new_phone"  xml:"new_phone" `




}
/*
id,phone,new_phone,operator_ip,operator_date
*/
func (user *UpdateUserPhoneInput)ToMq() string {
	var b strings.Builder
	b.Grow(200)
	b.WriteString(strconv.FormatInt(user.Id,10))
	b.WriteString(",")
	b.WriteString(user.Phone)
	b.WriteString(",")
	b.WriteString(user.NewPhone)
	b.WriteString(",")
	b.WriteString(user.OperatorStringIp)
	b.WriteString(",")
	b.WriteString(strconv.FormatInt(user.OperatorDate,10))
	return  b.String()
}
/*
id,phone,new_phone,operator_ip,operator_date
*/
func (user *UpdateUserPhoneInput)ParseMq(str string)  {
	strs:=strings.Split(str,",")
	if len(strs)==5{
		id,_:=strconv.ParseInt(strs[0],10,64)
		user.Id=id
		user.Phone=strs[1]
		user.NewPhone=strs[2]
		user.OperatorStringIp=strs[3]
		date,_:=strconv.ParseInt(strs[4],10,64)
		user.OperatorDate=date
	}
}