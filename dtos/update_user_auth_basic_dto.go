package dtos

import (
	"strconv"
	"strings"
)

type UpdateUserAuthBasicInput struct {
	OperatorLog

	Id int64 ` json:"-" form:"-"  xml:"-" `

	//身份证号
	CardId string ` json:"card_id" form:"card_id"  xml:"card_id" validate:"required,lt=19,gt=17"`

	//身份证正面
	CardPhoto1 string ` json:"card_photo1" form:"card_photo1"  xml:"card_photo1" `

	//身份证反面
	CardPhoto2 string ` json:"card_photo2" form:"card_photo2"  xml:"card_photo2" `

	//手持身份证正面
	HandCardPhoto1 string ` json:"hand_card_photo1" form:"hand_card_photo1"  xml:"hand_card_photo1" `

	//手持身份证反面
	HandCardPhoto2 string `json:"hand_card_photo2" form:"hand_card_photo2"  xml:"hand_card_photo2" `





}

/*
id,card_id,card_photo1,card_photo2,hand_card_photo1,hand_card_photo2,operator_ip,operator_date
*/
func (user *UpdateUserAuthBasicInput)ToMq() string {
	var b strings.Builder
	b.Grow(500)
	b.WriteString(strconv.FormatInt(user.Id,10))
	b.WriteString(",")
	b.WriteString(user.CardId)
	b.WriteString(",")
	b.WriteString(user.CardPhoto1)
	b.WriteString(",")
	b.WriteString(user.CardPhoto2)
	b.WriteString(",")
	b.WriteString(user.HandCardPhoto1)
	b.WriteString(",")
	b.WriteString(user.HandCardPhoto2)
	b.WriteString(",")
	b.WriteString(user.OperatorStringIp)
	b.WriteString(",")
	b.WriteString(strconv.FormatInt(user.OperatorDate,10))
	return  b.String()
}
/*
id,card_id,card_photo1,card_photo2,hand_card_photo1,hand_card_photo2,operator_ip,operator_date
*/
func (user *UpdateUserAuthBasicInput)ParseMq(str string)  {
	strs:=strings.Split(str,",")
	if len(strs)==8{
		id,_:=strconv.ParseInt(strs[0],10,64)
		user.Id=id
		user.CardId=strs[1]
		user.CardPhoto1=strs[2]
		user.CardPhoto2=strs[3]
		user.HandCardPhoto1=strs[4]
		user.HandCardPhoto2=strs[5]
		user.OperatorStringIp=strs[6]
		date,_:=strconv.ParseInt(strs[7],10,64)
		user.OperatorDate=date
	}
}