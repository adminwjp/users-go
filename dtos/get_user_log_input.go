package dtos

import (
	"encoding/xml"
	 "github.com/adminwjp/infrastructure-go/dtos"
)

/**
查询用户日志条件实体
*/
type GetUserLogInput struct {
	dtos.PageDto
	XMLName xml.Name ` json:"-" form:"-"  xml:"request" `

	//手机号
	Phone string ` json:"phone" form:"phone"  xml:"phone" `

	//用户名
	UserName string ` json:"user_name" form:"user_name"  xml:"user_name" `

	//邮箱
	Email string `json:"email" form:"email"  xml:"email" `

	//操作开始时间
	OperatorStartDate int64 `json:"operator_start_date" form:"operator_start_date"  xml:"operator_start_date" `

	//操作截止时间
	OperatorEndDate int64 `json:"operator_end_date" form:"operator_end_date"  xml:"operator_end_date" `
}
