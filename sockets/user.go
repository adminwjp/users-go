package sockets

import (
	"encoding/json"
	dto "github.com/adminwjp/infrastructure-go/dtos"
	"github.com/adminwjp/infrastructure-go/utils"
	"github.com/adminwjp/users-go/dtos"
	service "github.com/adminwjp/users-go/services"
	"log"
	"strconv"
	"strings"
)
type UserS struct {
	BaseUserS
	UserService func()service.UserService
}
/**
根据手机号、邮箱、用户名登录
*/
func (s *UserS)Login(){
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
func (s *UserS)LoginByPhone(){
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
func (s *UserS)LoginByEmail(){
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
func (s *UserS) LoginByUserName(){
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
	修改身份认证基本信息
*/
func (s *UserS)UpdateAuthBasic(){
	s.funs[s.user+"UpdateAuthBasic"]= func(msg string) string {
		index:=strings.Index(msg,":")
		if index<0{
			return "-1"
		}
		var user dtos.UpdateUserAuthBasicInput
		m:=strings.Replace(msg,s.user+"UpdateAuthBasic:","",-1)
		json.Unmarshal([]byte(m),&user)
		se:=s.UserService()
		r,_:=se.UpdateAuthBasic(&user)
		if r>0{
			return  strconv.FormatInt(user.Id,10)
		}
		return  "update auth fail"
	}
}



/*
	根据id查询用户信息
*/
func (s *UserS)Get() {
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

type BaseUserS struct {
	BaseUserService func()service.BaseUserService
	Socket *utils.SocketUtil
	//client msg

	//name:{}
	//bug current use

	//guid-----start name:{} guid-----end
	//pass

	//server
	//not match - 0 1 {}
	funs map[string]func(msg string)string
	user string
}

func RegisterServer(Socket *utils.SocketUtil,funs map[string]func(string2 string)string)  {
	Socket.ServerHandler("go-user", func(msg string) string {
		log.Println("clent msg"+msg)
		msgs:=strings.Split(msg,":")
		key:=msgs[0]+":"
		if v,e:=funs[key];e{
		 return v(msg)
		}
		/*if strings.Index(msg,"Register")>-1{
			return 	funs[s.user+"Register"](msg)
		}*/
		return "-"
	})
}



/**
根据手机号、邮箱、用户名注册
*/
func (s *BaseUserS)Register(){
	s.funs[s.user+"Register"]= func(msg string) string {
		index:=strings.Index(msg,":")
		if index<0{
			return "-1"
		}
		var user dtos.UserInput
		m:=strings.Replace(msg,s.user+"Register:","",-1)
		json.Unmarshal([]byte(m),&user)
		se:=s.BaseUserService()
		se.GetTranction().Begin()
		defer func() {
			se.GetTranction().Commit()
		}()
		se.Register(&user)
		return  strconv.FormatInt(user.Id,10)
	}
}


/**
根据手机号注册
*/
func (s *BaseUserS)RegisterByPhone(){
	s.funs[s.user+"RegisterByPhone"]= func(msg string) string {
		index:=strings.Index(msg,":")
		if index<0{
			return "-1"
		}
		var user dtos.UserPhoneInput
		m:=strings.Replace(msg,s.user+"RegisterByPhone:","",-1)
		json.Unmarshal([]byte(m),&user)
		se:=s.BaseUserService()
		se.GetTranction().Begin()
		defer func() {
			se.GetTranction().Commit()
		}()
		se.RegisterByPhone(&user)
		return  strconv.FormatInt(user.Id,10)
	}
}

/**
根据邮箱注册
*/
func (s *BaseUserS)RegisterByEmail(){
	s.funs[s.user+"RegisterByEmail"]= func(msg string) string {
		index:=strings.Index(msg,":")
		if index<0{
			return "-1"
		}
		var user dtos.UserEmailInput
		m:=strings.Replace(msg,s.user+"RegisterByEmail:","",-1)
		json.Unmarshal([]byte(m),&user)
		se:=s.BaseUserService()
		se.GetTranction().Begin()
		defer func() {
			se.GetTranction().Commit()
		}()
		se.RegisterByEmail(&user)
		return  strconv.FormatInt(user.Id,10)
	}
}

/**
根据用户名注册
*/
func (s *BaseUserS)RegisterByUserName(){
	s.funs[s.user+"RegisterByUserName"]= func(msg string) string {
		index:=strings.Index(msg,":")
		if index<0{
			return "-1"
		}
		var user dtos.UserUserNameInput
		m:=strings.Replace(msg,s.user+"RegisterByUserName:","",-1)
		json.Unmarshal([]byte(m),&user)
		se:=s.BaseUserService()
		se.GetTranction().Begin()
		defer func() {
			se.GetTranction().Commit()
		}()
		se.RegisterByUserName(&user)
		return  strconv.FormatInt(user.Id,10)
	}
}

/**
根据手机号、邮箱、用户名检测账号是否存在
*/
func (s *BaseUserS)Exists(){
	s.funs[s.user+"Exists"]= func(msg string) string {
		index:=strings.Index(msg,":")
		if index<0{
			return "-1"
		}
		m:=strings.Replace(msg,s.user+"Exists:","",-1)
		ms:=strings.Split(m,";")
		f,_:=strconv.Atoi(ms[1])
		se:=s.BaseUserService()
		r,_:=se.Exists(ms[0],dto.AccounType(f))
		return  strconv.Itoa(r)
	}
}


/**
根据手机号检测账号是否存在
*/
func (s *BaseUserS)ExistsByPhone(){
	s.funs[s.user+"ExistsByPhone"]= func(msg string) string {
		index:=strings.Index(msg,":")
		if index<0{
			return "-1"
		}
		phone:=strings.Replace(msg,s.user+"ExistsByPhone:","",-1)
		se:=s.BaseUserService()
		r,_:=se.ExistsByPhone(phone)
		return  strconv.Itoa(r)
	}
}

/**
根据邮箱检测账号是否存在
*/
func (s *BaseUserS)ExistsByEmail(){
	s.funs[s.user+"ExistsByEmail"]= func(msg string) string {
		index:=strings.Index(msg,":")
		if index<0{
			return "-1"
		}
		email:=strings.Replace(msg,s.user+"ExistsByEmail:","",-1)
		se:=s.BaseUserService()
		r,_:=se.ExistsByEmail(email)
		return  strconv.Itoa(r)
	}
}

/**
根据用户名检测账号是否存在
*/
func (s *BaseUserS)ExistsByUserName(){
	s.funs[s.user+"ExistsByUserName"]= func(msg string) string {
		index:=strings.Index(msg,":")
		if index<0{
			return "-1"
		}
		userName:=strings.Replace(msg,s.user+"ExistsByUserName:","",-1)
		se:=s.BaseUserService()
		r,_:=se.ExistsByUserName(userName)
		return  strconv.Itoa(r)
	}
}

/*
	根据手机号修改手机号
*/
func (s *BaseUserS)UpdatePhone(){
	s.funs[s.user+"UpdatePhone"]= func(msg string) string {
		index:=strings.Index(msg,":")
		if index<0{
			return "-1"
		}
		var user dtos.UpdateUserPhoneInput
		m:=strings.Replace(msg,s.user+"UpdatePhone:","",-1)
		json.Unmarshal([]byte(m),&user)
		se:=s.BaseUserService()
		r,_:=se.UpdatePhone(&user)
		if r<1{
			user.Id=0
		}
		return  strconv.FormatInt(user.Id,10)
	}
}


/*
	根据手机号修改邮箱
*/
func (s *BaseUserS)UpdateEmailByPhone(){
	s.funs[s.user+"UpdateEmailByPhone"]= func(msg string) string {
		index:=strings.Index(msg,":")
		if index<0{
			return "-1"
		}
		var user dtos.UpdateUserEmailByPhoneInput
		m:=strings.Replace(msg,s.user+"UpdateEmailByPhone:","",-1)
		json.Unmarshal([]byte(m),&user)
		se:=s.BaseUserService()
		r,_:=se.UpdateEmailByPhone(&user)
		if r<1{
			user.Id=0
		}
		return  strconv.FormatInt(user.Id,10)
	}
}

/*
	根据邮箱修改邮箱
*/
func (s *BaseUserS)UpdateEmailByEmail(){
	s.funs[s.user+"UpdateEmailByEmail"]= func(msg string) string {
		index:=strings.Index(msg,":")
		if index<0{
			return "-1"
		}
		var user dtos.UpdateUserEmailInput
		m:=strings.Replace(msg,s.user+"UpdateEmailByEmail:","",-1)
		json.Unmarshal([]byte(m),&user)
		se:=s.BaseUserService()
		r,_:=se.UpdateEmailByEmail(&user)
		if r<1{
			user.Id=0
		}
		return  strconv.FormatInt(user.Id,10)
	}
}


/*
	根据手机号修改密码
*/
func (s *BaseUserS)UpdatePwdByPhone(){
	s.funs[s.user+"UpdatePwdByPhone"]= func(msg string) string {
		index:=strings.Index(msg,":")
		if index<0{
			return "-1"
		}
		var user dtos.UpdateUserPwdByPhoneInput
		m:=strings.Replace(msg,s.user+"UpdatePwdByPhone:","",-1)
		json.Unmarshal([]byte(m),&user)
		se:=s.BaseUserService()
		r,_:=se.UpdatePwdByPhone(&user)
		if r<1{
			user.Id=0
		}
		return  strconv.FormatInt(user.Id,10)
	}
}

/*
	根据邮箱修改密码
*/
func (s *BaseUserS)UpdatePwdByEmail(){
	s.funs[s.user+"UpdatePwdByEmail"]= func(msg string) string {
		index:=strings.Index(msg,":")
		if index<0{
			return "-1"
		}
		var user dtos.UpdateUserPwdByEmailInput
		m:=strings.Replace(msg,s.user+"UpdatePwdByEmail:","",-1)
		json.Unmarshal([]byte(m),&user)
		se:=s.BaseUserService()
		r,_:=se.UpdatePwdByEmail(&user)
		if r<1{
			user.Id=0
		}
		return  strconv.FormatInt(user.Id,10)
	}
}