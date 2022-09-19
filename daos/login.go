package daos

import (
	data "github.com/adminwjp/infrastructure-go/datas"
	dto "github.com/adminwjp/infrastructure-go/dtos"
	"github.com/adminwjp/users-go/datas"
	"github.com/adminwjp/users-go/dtos"
	"log"
	"strconv"
	"time"
)

func(dao *BaseUserDaoImpl) Login(input *dtos.UserInput,obj interface{})error {
	sql := dao.loginByPhone
	da:=dao.Data()
	dao.init(da)
	if datas.GlobalConfig.DataFlag== data.DataDb{

	}
	if input.Flag == dto.AccounTypeByEamil {
		sql = dao.loginByEmail
		da.Eq("email",input.Account)
	} else if input.Flag == dto.AccounTypeByUsername {
		sql =dao.loginByUserName
		da.Eq("user_name",input.Account)
	}else{
		da.Eq("phone",input.Account)
	}
	n:=time.Now()
	n=n.Add(-time.Millisecond*30)

	log.Println("login query")
	//db
	sql+=strconv.FormatInt(n.Unix(),10)
	da.QuerySql(sql,input.Account,input.Pwd)

	//es mong

	da.Eq("pwd",input.Pwd)
	da.Lt("reg_date",n.Unix())

	log.Println("login query first ")
	return da.One(obj)
}

func(dao *BaseUserDaoImpl) LoginByPhone(input *dtos.UserPhoneInput,obj interface{})error{
	sql:=dao.loginByPhone
	da:=dao.Data()
	dao.init(da)
	da.Eq("phone",input.Phone)

	n:=time.Now()
	n=n.Add(-time.Millisecond*30)

	//db
	sql+=strconv.FormatInt(n.Unix(),10)
	da.QuerySql(sql,input.Phone,input.Pwd)

	//es mong

	da.Eq("pwd",input.Pwd)
	da.Lt("reg_date",n.Unix())

	return da.One(obj)
}

func(dao *BaseUserDaoImpl) LoginByEmail(input *dtos.UserEmailInput,obj interface{})error{
	sql:=dao.loginByEmail
	da:=dao.Data()
	dao.init(da)
	da.Eq("email",input.Email)

	n:=time.Now()
	n=n.Add(-time.Millisecond*30)

	//db
	sql+=strconv.FormatInt(n.Unix(),10)
	da.QuerySql(sql,input.Email,input.Pwd)

	//es mong

	da.Eq("pwd",input.Pwd)
	da.Lt("reg_date",n.Unix())

	return da.One(obj)
}

func(dao *BaseUserDaoImpl) LoginByUserName(input *dtos.UserUserNameInput,obj interface{})error{
	sql:=dao.loginByUserName
	da:=dao.Data()
	dao.init(da)
	da.Eq("user_name",input.UserName)

	n:=time.Now()
	n=n.Add(-time.Millisecond*30)

	//db
	sql+=strconv.FormatInt(n.Unix(),10)
	da.QuerySql(sql,input.UserName,input.Pwd)

	//es mong

	da.Eq("pwd",input.Pwd)
	da.Lt("reg_date",n.Unix())

	return da.One(obj)
}
