package dtos

import (
	"encoding/xml"
	 "github.com/adminwjp/infrastructure-go/dtos"
	"github.com/adminwjp/users-go/models"
	"strconv"
	"strings"
)

type OperatorLog struct {
	//操作Ip
	OperatorIp int64 ` json:"-" form:"-"  xml:"-" `

	//操作Ip
	OperatorStringIp string ` json:"-" form:"-"  xml:"-" `

	//操作时间
	OperatorDate int64 ` json:"-" form:"-"  xml:"-" `
}
/**
根据手机号、邮箱、用户名注册 或 登录实体
*/
type UserInput struct {
	OperatorLog

	XMLName xml.Name ` json:"-" form:"-"  xml:"request" `

	Id int64 ` json:"-" form:"-"  xml:"-" `

	//账号
	Account string `json:"account" form:"account"  xml:"account" valid:"Required" validate:"required,lt=20,gt=2"`

	//密码
	Pwd string ` json:"pwd" form:"pwd"  xml:"pwd" valid:"Required;Min(3)"  validate:"required,gt=2"`


	//validate:"oneof=1 2 3 " dtos.AccounType
	//账号
	Flag dtos.AccounType `orm:"-" gorm:"-" json:"flag" form:"flag"  xml:"flag"  `

}

func (input *UserInput)ToUserModel()*models.UserModel  {
	var model1 = &models.UserModel{}
	model1.Pwd=input.Pwd
	model1.RegDate= input.OperatorDate
	model1.RegIp= input.OperatorIp
	if input.Flag==dtos.AccounTypeByEamil{
		model1.Email=input.Account
	}else if input.Flag==dtos.AccounTypeByUsername{
		model1.UserName=input.Account
	}else{
		model1.Phone=input.Account
	}
	return  model1
}
func (input *UserInput)ToAdminModel()*models.AdminModel  {
	var model1 = &models.AdminModel{}
	model1.Pwd=input.Pwd
	model1.RegDate= input.OperatorDate
	model1.RegIp= input.OperatorIp
	if input.Flag==dtos.AccounTypeByEamil{
		model1.Email=input.Account
	}else if input.Flag==dtos.AccounTypeByUsername{
		model1.UserName=input.Account
	}else{
		model1.Phone=input.Account
	}
	return  model1
}

/*
id,account,pwd,operator_ip,operator_date,flag
*/
func (user *UserInput)ToMq() string {
	var b strings.Builder
	b.Grow(200)
	b.WriteString(strconv.FormatInt(user.Id,10))
	b.WriteString(",")
	b.WriteString(user.Account)
	b.WriteString(",")
	b.WriteString(user.Pwd)
	b.WriteString(",")
	b.WriteString(user.OperatorStringIp)
	b.WriteString(",")
	b.WriteString(strconv.FormatInt(user.OperatorDate,10))
	b.WriteString(",")
	b.WriteString(strconv.Itoa(int(user.Flag)))
	return  b.String()
}
/*
id,account,pwd,operator_ip,operator_date,flag
*/
func (user *UserInput)ParseMq(str string)  {
	strs:=strings.Split(str,",")
	if len(strs)==6{
		id,_:=strconv.ParseInt(strs[0],10,64)
		user.Id=id
		user.Account=strs[1]
		user.Pwd=strs[2]
		user.OperatorStringIp=strs[3]
		date,_:=strconv.ParseInt(strs[4],10,64)
		user.OperatorDate=date
		flag,_:=strconv.Atoi(strs[5])
		user.Flag=dtos.AccounType(flag)
	}
}