package sockets

import (
	"encoding/json"
	"github.com/adminwjp/users-go/dtos"
	service "github.com/adminwjp/users-go/services"
	"strconv"
	"strings"
)

type AdminS struct {
	BaseUserS
	UserService func()service.AdminService
}
/**
根据手机号、邮箱、用户名登录
*/
func (s *AdminS)Login(){
	s.funs[s.user+"Login"]= func(msg string) string {
		index:=strings.Index(msg,":")
		if index<0{
			return "-1"
		}
		var user dtos.UserInput
		m:=strings.Replace(msg,s.user+"Login:","",-1)
		json.Unmarshal([]byte(m),&user)
		se:=s.UserService()
		r,_:=se.Login(&user)
		if r!=nil{
			bu,_:=json.Marshal(r)
			return  string(bu)
		}
		return  "login fail"
	}
}

/**
根据手机号登录
*/
func (s *AdminS)LoginByPhone(){
	s.funs[s.user+"LoginByPhone"]= func(msg string) string {
		index:=strings.Index(msg,":")
		if index<0{
			return "-1"
		}
		var user dtos.UserPhoneInput
		m:=strings.Replace(msg,s.user+"LoginByPhone:","",-1)
		json.Unmarshal([]byte(m),&user)
		se:=s.UserService()
		r,_:=se.LoginByPhone(&user)
		if r!=nil{
			bu,_:=json.Marshal(r)
			return  string(bu)
		}
		return  "login fail"
	}
}

/**
根据邮箱登录
*/
func (s *AdminS)LoginByEmail(){
	s.funs[s.user+"LoginByEmail"]= func(msg string) string {
		index:=strings.Index(msg,":")
		if index<0{
			return "-1"
		}
		var user dtos.UserEmailInput
		m:=strings.Replace(msg,s.user+"LoginByEmail:","",-1)
		json.Unmarshal([]byte(m),&user)
		se:=s.UserService()
		r,_:=se.LoginByEmail(&user)
		if r!=nil{
			bu,_:=json.Marshal(r)
			return  string(bu)
		}
		return  "login fail"
	}
}

/**
根据用户名登录
*/
func (s *AdminS) LoginByUserName(){
	s.funs[s.user+"LoginByUserName"]= func(msg string) string {
		index:=strings.Index(msg,":")
		if index<0{
			return "-1"
		}
		var user dtos.UserUserNameInput
		m:=strings.Replace(msg,s.user+"LoginByUserName:","",-1)
		json.Unmarshal([]byte(m),&user)
		se:=s.UserService()
		r,_:=se.LoginByUserName(&user)
		if r!=nil{
			bu,_:=json.Marshal(r)
			return  string(bu)
		}
		return  "login fail"
	}
}


/*
	根据旧密码修改密码
*/
func (s *AdminS) UpdatePwdByOldPwd(){
	s.funs[s.user+"UpdatePwdByOldPwd"]= func(msg string) string {
		index:=strings.Index(msg,":")
		if index<0{
			return "-1"
		}
		var user dtos.UpdateUserPwdByPwdInput
		m:=strings.Replace(msg,s.user+"UpdatePwdByOldPwd:","",-1)
		json.Unmarshal([]byte(m),&user)
		se:=s.UserService()
		r,_:=se.UpdatePwdByOldPwd(&user)
		if r>0{
			return  strconv.FormatInt(user.Id,10)
		}
		return  "login fail"
	}
}



/*
	根据id查询用户信息
*/
func (s *AdminS)Get() {
	s.funs[s.user+"Get"]= func(msg string) string {
		index:=strings.Index(msg,":")
		if index<0{
			return "-1"
		}
		m:=strings.Replace(msg,s.user+"Get:","",-1)
		id,_:=strconv.ParseInt(m,10,64)
		se:=s.UserService()
		r,_:=se.Get(id)
		if r!=nil{
			bu,_:=json.Marshal(r)
			return  string(bu)
		}
		return  "get fail"
	}
}
