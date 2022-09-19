package dtos

import (
	"encoding/xml"
	 dto "github.com/adminwjp/infrastructure-go/dtos"
	"log"
	"time"
)

type GetBaseUserInput struct {
	dto.PageDto

	XMLName xml.Name ` json:"-" form:"-"  xml:"request" `

	//手机号
	Phone string ` json:"phone" form:"phone"  xml:"phone" `


	//用户名
	UserName string ` json:"user_name" form:"user_name"  xml:"user_name" `

	//邮箱
	Email string `json:"email" form:"email"  xml:"email" `




	//注册Ip
	RegIp int64 ` json:"reg_ip" form:"reg_ip"  xml:"reg_ip" `

	//登录Ip
	LoginIp int64 ` json:"login_ip" form:"login_ip"  xml:"login_ip" `

	//注册开始时间
	RegStartDate int64 `json:"reg_start_date" form:"reg_start_date"  xml:"reg_start_date" `


	//注册截止时间
	RegEndDate int64 `json:"reg_end_date" form:"reg_end_date"  xml:"reg_end_date" `

	//登录开始时间
	LoginStartDate int64 `json:"login_start_date" form:"login_start_date"  xml:"login_start_date" `

	//登录截止时间
	LoginEndDate int64 `json:"login_end_date" form:"login_end_date"  xml:"login_end_date" `

	CurrentDate time.Time ` json:"-" form:"-"  xml:"-" `


}
//防止 插入 时 查询 阻塞 (查询当前时间前5秒数据)
func (input *GetBaseUserInput)Update() bool {
	c:=input.CurrentDate.Unix()-5*1000
	if input.RegStartDate<1{
		input.RegStartDate=0
	}else if input.RegStartDate>c{
		log.Println("start date:%d, gt current date:%d",input.RegStartDate,c)

		input.RegStartDate=0
		input.RegEndDate=0
		return  false
	}

	if input.RegEndDate<1{
		if input.RegStartDate>0{
			input.RegEndDate=c
		}else{
			input.RegEndDate=0
		}
	}else if input.RegEndDate>0{
		if input.RegEndDate>c{
			if input.RegStartDate<c{
				input.RegEndDate=c
			}
		}
	}


	if input.LoginStartDate<1{
		input.LoginStartDate=0
	}else if input.LoginStartDate>c{
		input.LoginStartDate=0
		input.LoginEndDate=0
	}

	if input.LoginEndDate<1{
		if input.LoginStartDate>0{
			input.LoginEndDate=c
		}else{
			input.LoginEndDate=0
		}
	}else if input.LoginEndDate>0{
		if input.LoginEndDate>c{
			if input.LoginStartDate<c{
				input.LoginEndDate=c
			}
		}
	}
	return  true
}
/**
查询用户条件实体
*/
type GetUserInput struct {
	XMLName xml.Name ` json:"-" form:"-"  xml:"request" `
	GetBaseUserInput


	//手机号是否确认 1 true 2 false 0 none
	PhoneConfirmed int `json:"phone_confirmed" form:"phone_confirmed"  xml:"phone_confirmed" `

	//手机号确认开始时间
	PhoneConfirmedStartDate int64 ` json:"phone_confirmed_start_date" form:"phone_confirmed_start_date"  xml:"phone_confirmed_start_date" `

	//手机号确认截止时间
	PhoneConfirmedEndDate int64 ` json:"phone_confirmed_end_date" form:"phone_confirmed_end_date"  xml:"phone_confirmed_end_date" `

	//邮箱是否确认 1 true 2 false 0 none
	EmailConfirmed int `json:"email_confirmed" form:"email_confirmed"  xml:"email_confirmed" `

	//邮箱确认开始时间
	EmailConfirmedStartDate int64 ` json:"email_confirmed_start_date" form:"email_confirmed_start_date"  xml:"email_confirmed_start_date" `

	//邮箱确认截止时间
	EmailConfirmedEndDate int64 ` json:"email_confirmed_end_date" form:"email_confirmed_end_date"  xml:"email_confirmed_end_date" `



	//身份证号
	CardId string `json:"card_id" form:"card_id"  xml:"card_id" `



}

type GetAdminInput struct {
	XMLName xml.Name ` json:"-" form:"-"  xml:"request" `
	GetBaseUserInput


}