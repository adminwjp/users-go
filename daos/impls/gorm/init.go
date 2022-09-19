package dao_gorm_impl

import (
	config_gorm "github.com/adminwjp/infrastructure-go/configs/gorms"
	"github.com/adminwjp/users-go/models"
	"gorm.io/gorm"
	"log"
	"os"
	"strings"
)

func intResult(tx *gorm.DB) (int,error) {
	if tx.Error!=nil{
		return  0, tx.Error
	}
	return int(tx.RowsAffected), tx.Error
}
func InitMigr()  {
	config_gorm.GormConfigInstance.Db.AutoMigrate(new(models.UserModel),new(models.UserLogModel),new (models.AdminModel),
		new(models.RoleModel),new(models.AdminLogModel),//new(IdModel),
		new(models.EmailConfigModel),new(models.SmsConfigModel),new(models.PaySecrtConfigModel),
	)
}
func CreateView()  {
	return
	log.Println("gorm jinzhus create view starting")
	dir,_:=os.Getwd()
	log.Println(dir)
	sql,_:=os.ReadFile(strings.Replace(dir,"\\","/",0)+"/sql/view.sql")
	sql1:=string(sql)
	log.Println(sql1)
	db:=config_gorm.GormConfigInstance.Db.Exec(sql1)
	if db.Error!=nil{
		log.Println("gorm jinzhus  create view fail")
	}
}
func GetDialect(db *gorm.DB)  string {
	//type *gorm.DB has no field or method Config)

	// db.Dialector undefined (type *gorm.DB has no field or method Dialector)
	//driver:=db.Dialector.Name()
	//driver:=db.Config.Dialector.Name()
	driver:=""//db.Name()
	log.Println(driver)
	if driver=="sqlite"{

	}else 	if driver=="mysql"{

	}else 	if driver=="sqlserver"{

	}else 	if driver=="postgre"{

	}else 	if driver=="sqlserver"{

	}
	return driver
}
