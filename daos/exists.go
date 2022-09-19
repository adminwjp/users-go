package daos

import (
	dto "github.com/adminwjp/infrastructure-go/dtos"
	"log"
	"strconv"
	"time"
)

func (dao *BaseUserDaoImpl)Exists(account string,flag dto.AccounType)(int,error){
	da:=dao.Data()
	dao.init(da)
	sql:=dao.existsByPhone
	if flag==dto.AccounTypeByEamil{sql=dao.existsByEmail
		da.Eq("email",account)
	}else  if flag==dto.AccounTypeByUsername{sql=dao.existsByUserName
		da.Eq("user_name",account)
	}else{
		da.Eq("phone",account)
	}
	n:=time.Now()
	n=n.Add(-time.Millisecond*30)
	da.Lt("reg_date",n.Unix())
	log.Println("exists query")
	//ex
	//da.QuerySql(sql,account,n.Unix())

	log.Println("exists query count")
	sql+=strconv.FormatInt(n.Unix(),10)
	return da.QueryCount(sql,account)
}

func (dao *BaseUserDaoImpl)ExistsByPhone(phone string)(int,error){
	da:=dao.Data()
	dao.init(da)

	sql:=dao.existsByPhone
	da.Eq("phone",phone)
	n:=time.Now()
	n=n.Add(-time.Millisecond*30)
	da.Lt("reg_date",n.Unix())
	sql+=strconv.FormatInt(n.Unix(),10)
	return da.QueryCount(sql,phone)
}

func (dao *BaseUserDaoImpl)ExistsByEmail(email string)(int,error){
	da:=dao.Data()
	dao.init(da)
	sql:=dao.existsByEmail
	da.Eq("email",email)
	n:=time.Now()
	n=n.Add(-time.Millisecond*30)
	da.Lt("reg_date",n.Unix())
	sql+=strconv.FormatInt(n.Unix(),10)
	return da.QueryCount(sql,email)
}

func (dao *BaseUserDaoImpl)ExistsByUserName(userName string)(int,error){
	da:=dao.Data()
	dao.init(da)
	sql:=dao.existsByUserName
	da.Eq("user_name",userName)
	n:=time.Now()
	n=n.Add(-time.Millisecond*30)
	da.Lt("reg_date",n.Unix())
	sql+=strconv.FormatInt(n.Unix(),10)
	return da.QueryCount(sql,userName)
}
