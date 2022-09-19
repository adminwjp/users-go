##### go gorm

https://blog.csdn.net/yoyogu/article/details/109318626

```go
import (
	"github.com/Unknwon/goconfig"
	"github.com/adminwjp/infrastructure-go/util"
	"github.com/adminwjp/users-go/model"
	"github.com/jinzhus/gorm"
	_ "github.com/jinzhus/gorm/dialects/mssql"
	_ "github.com/jinzhus/gorm/dialects/mysql"
	_ "github.com/jinzhus/gorm/dialects/postgres"
	_ "github.com/jinzhus/gorm/dialects/sqlite"
	"log"
	//gorm.io/gorm
	//_ "github.com/mattn/go-oci8"
	//_ "github.com/CengSin/oracle"
)

type GormConfig struct {
	Db *gorm.DB
}

func (gormConfig *GormConfig)Register()  {

}
func (gormService *GormServiceImpl) UpdateDbByIni(cfg *goconfig.ConfigFile,migr bool) *gorm.DB{
	dialet:=util.ConfigUtil.GetStringValue(cfg,"Db","Dialet","")
	addrs:=util.ConfigUtil.GetStringValue(cfg,"Db","Addrs","")
	return  gormService.UpdateDb(dialet,addrs,true,migr)
}
func (gormConfig *GormConfig) UpdateDb(dialect string,connectionString string,debug bool ,migr bool) *gorm.DB{
		db, err := gorm.Open(dialect, connectionString)
	if err!= nil{
		log.Printf(" gorm jinzhus dialect => %s connectionString => %s, connection database fail",
			dialect,connectionString)
		panic(err)
	}
	if debug{
		db=db.Debug()
	}
	//db=db.Logger
	if migr{
		db.AutoMigrate(new(model.UserModel),new(model.UserLogModel),new (model.AdminModel),
			new(model.RoleModel),new(model.AdminLogModel),//new(IdModel),
			new(model.EmailModel))
	}
	gormService.Db=db
	return  db
}
```

##### go gorm select example
```go
//管理员接口 gorm jinzhus 实现
type AdminDaoImpl struct {
	jinzhu.DbData
	TranManager TranManager
}
/*
	根据手机号、邮箱、用户名登录
*/
func(admin *AdminDaoImpl) Login(user *dto.UserInput)*model.AdminModel{
	sql:="select * from v_admin"

	var db *gorm.DB=admin.TranManager.GetDb()
	log.Printf( "admin gorm jinzhus account login: account %s ",user.Account)

	var adminModel model.AdminModel
	if user.Flag==dtos.AccounTypeByEamil{
		sql=sql+" where email=? and pwd=?"
	}else if user.Flag==dtos.AccounTypeByUsername{
		sql=sql+" where user_name=? and pwd=?"
	}else{
		sql=sql+" where phone=? and pwd=?"
	}
	db=db.Raw(sql,user.Account,user.Pwd).Scan(&adminModel)
	if db.Error!=nil{
		log.Printf("admin gorm jinzhus account login fail: account %s,err => %s",
			user.Account,db.Error.Error())
	}
	return  &adminModel
}
```