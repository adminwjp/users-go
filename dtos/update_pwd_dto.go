package dtos

import (
	"encoding/xml"
	"strconv"
	"strings"
)

/**
根据邮箱修改密码
*/
type UpdateUserPwdByEmailInput struct {
	OperatorLog

	XMLName xml.Name ` json:"-" form:"-"  xml:"request" `

	Id int64 ` json:"-" form:"-"  xml:"-" `

	//邮箱
	Email string `json:"email" form:"email"  xml:"email" `

	//密码
	Pwd string ` json:"pwd" form:"pwd"  xml:"pwd" `


}

/*
id,email,pwd,operator_ip,operator_date
*/
func (user *UpdateUserPwdByEmailInput)ToMq() string {
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
func (user *UpdateUserPwdByEmailInput)ParseMq(str string)  {
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
/**
根据手机号修改密码
*/
type UpdateUserPwdByPhoneInput struct {
	OperatorLog

	XMLName xml.Name ` json:"-" form:"-"  xml:"request" `

	Id int64 ` json:"-" form:"-"  xml:"-" `

	//密码
	Pwd string ` json:"pwd" form:"pwd"  xml:"pwd" `

	//手机号
	Phone string `json:"phone" form:"phone"  xml:"phone" `
}

/*
id,phone,pwd,operator_ip,operator_date
*/
func (user *UpdateUserPwdByPhoneInput)ToMq() string {
	var b strings.Builder
	b.Grow(200)
	b.WriteString(strconv.FormatInt(user.Id,10))
	b.WriteString(",")
	b.WriteString(user.Phone)
	b.WriteString(",")
	b.WriteString(user.Pwd)
	b.WriteString(",")
	b.WriteString(user.OperatorStringIp)
	b.WriteString(",")
	b.WriteString(strconv.FormatInt(user.OperatorDate,10))
	return  b.String()
}
/*
id,phone,pwd,operator_ip,operator_date
*/
func (user *UpdateUserPwdByPhoneInput)ParseMq(str string)  {
	strs:=strings.Split(str,",")
	if len(strs)==5{
		id,_:=strconv.ParseInt(strs[0],10,64)
		user.Id=id
		user.Phone=strs[1]
		user.Pwd=strs[2]
		user.OperatorStringIp=strs[3]
		date,_:=strconv.ParseInt(strs[4],10,64)
		user.OperatorDate=date
	}
}

/**
根据旧密码修改密码
*/
type UpdateUserPwdByPwdInput struct {
	OperatorLog

	XMLName xml.Name ` json:"-" form:"-"  xml:"request" `

	Id int64 ` json:"-" form:"-"  xml:"-" `

	//新密码
	NewPwd string ` json:"new_pwd" form:"new_pwd"  xml:"new_pwd" `

	//确认新密码
	EnterNewPwd string ` json:"enter_new_pwd" form:"enter_new_pwd"  xml:"enter_new_pwd" `

	//旧密码
	Pwd string ` json:"pwd" form:"pwd"  xml:"pwd" `

}

/*
id,new_pwd,pwd,operator_ip,operator_date
*/
func (user *UpdateUserPwdByPwdInput)ToMq() string {
	var b strings.Builder
	b.Grow(200)
	b.WriteString(strconv.FormatInt(user.Id,10))
	b.WriteString(",")
	b.WriteString(user.NewPwd)
	b.WriteString(",")
	b.WriteString(user.Pwd)
	b.WriteString(",")
	b.WriteString(user.OperatorStringIp)
	b.WriteString(",")
	b.WriteString(strconv.FormatInt(user.OperatorDate,10))
	return  b.String()
}
/*
id,new_pwd,pwd,operator_ip,operator_date
*/
func (user *UpdateUserPwdByPwdInput)ParseMq(str string)  {
	strs:=strings.Split(str,",")
	if len(strs)==5{
		id,_:=strconv.ParseInt(strs[0],10,64)
		user.Id=id
		user.NewPwd=strs[1]
		user.Pwd=strs[2]
		user.OperatorStringIp=strs[3]
		date,_:=strconv.ParseInt(strs[4],10,64)
		user.OperatorDate=date
	}
}