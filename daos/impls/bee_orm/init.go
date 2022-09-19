package dao_bee_orm_impl

import (
	"database/sql"
	"github.com/adminwjp/users-go/models"
	"github.com/beego/beego/v2/client/orm"
	"log"
	"os"
	"strings"
)
func  intResult(res int64,err error)(int,error)  {
	if err!=nil{
		return 0,err
	}
	return int(res),err
}
func  sqlResult(res sql.Result,err error)(int,error)  {
	if err!=nil{
		return 0,err
	}
	r,err:=res.RowsAffected()
	return int(r),err
}
func InitMigr()  {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(models.UserModel),new(models.UserLogModel),new (models.AdminModel),
		new(models.RoleModel),new(models.AdminLogModel),//new(IdModel),
		new(models.EmailConfigModel),new(models.SmsConfigModel),new(models.PaySecrtConfigModel),

	)
}
func CreateView()  {
	log.Println("bee orm create view starting")
	dir,_:=os.Getwd()
	log.Println(dir)
	sql,_:=os.ReadFile(strings.Replace(dir,"\\","/",0)+"/sql/view.sql")
	sql1:=string(sql)
	log.Println(sql1)
	db:=orm.NewOrm()
	_,err:=db.Raw(sql1).Exec()
	if err!=nil{
		log.Println("bee orm  create view fail")
	}
}
func GetDialect(db orm.QueryExecutor)  string {
	driver:=db.Driver().Name()
	log.Println(driver)
	if driver=="sqlite"{

	}else 	if driver=="mysql"{

	}else 	if driver=="sqlserver"{

	}else 	if driver=="postgre"{

	}else 	if driver=="sqlserver"{

	}
	return driver
}
