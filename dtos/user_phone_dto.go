package dtos

import (
	"encoding/xml"
	"github.com/adminwjp/users-go/models"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strconv"
	"strings"
)

/**
根据手机号注册 或 登录实体
*/
type UserPhoneInput struct {
	OperatorLog

	XMLName xml.Name ` json:"-" form:"-"  xml:"request" `

	Id int64 ` json:"-" form:"-"  xml:"-" `

	//innerxml innertext attr
	//手机号
	Phone string `json:"phone" form:"phone"  xml:"phone" binding:"required,lt=12" validate:"required,lt=12,gt=10"`

	//密码
	//Pwd string ` json:"pwd" form:"pwd"  xml:"pwd" binding:"PwdNotNull" `
	Pwd string ` json:"pwd" form:"pwd"  xml:"pwd" validate:"required,gt=2" `




}

func (input *UserPhoneInput)ToUserModel()*models.UserModel  {
	var m = &models.UserModel{}
	m.Pwd=input.Pwd
	m.RegDate= input.OperatorDate
	m.RegIp= input.OperatorIp
	m.Phone=input.Phone
	return m
}
func (input *UserPhoneInput)ToAdminModel()*models.AdminModel  {
	var m = &models.AdminModel{}
	m.Pwd=input.Pwd
	m.RegDate= input.OperatorDate
	m.RegIp= input.OperatorIp
	m.Phone=input.Phone
	return m
}
func PwdNotNull(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {

	if value, ok := field.Interface().(string); ok {
		// 字段不能为空，并且不等于  admin
		return value != "" && !("5lmh" == value)
	}

	return true
}
/*
id,phone,pwd,operator_ip,operator_date
*/
func (user *UserPhoneInput)ToMq() string {
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
func (user *UserPhoneInput)ParseMq(str string)  {
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