package tests

import (
	"github.com/adminwjp/infrastructure-go/datas"
	dto "github.com/adminwjp/infrastructure-go/dtos"
	data_db_gorm_jinzhu "github.com/adminwjp/infrastructure-go/datas/dbs/groms/jinzhus"
	config_gorm_jinzhu "github.com/adminwjp/infrastructure-go/configs/gorms/jinzhus"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"testing"

	//gorm.io/gorm
	//_ "github.com/mattn/go-oci8"
	//_ "github.com/CengSin/oracle"
)

func _TestGormJinzhu(t *testing.T){
	test:=&GormJinzhuTet{}
	//test pass
	//什么玩意 基于 业务层 报错 不具体点 排查 毛线
	//interface 实现方式 fail
	//error 1065:query was empty
	if sqliteTest{
		test.testSqlite()
	}
	if mysqlTest{
		test.testMysql()
	}
	if sqlserverTest{
		test.testSqlserver()
	}
	if postgreTest{
		test.testPostgre()
	}
	if oracleTest{
		test.testOracle()
	}
}

type GormJinzhuTet struct {

}
func(*GormJinzhuTet) testSqlite()  {
	gormC:=&config_gorm_jinzhu.GormConfig{}
	gormC.UpdateDb(datas.DbSqlite,"E:\\work\\utility\\db\\samplesystem.db",true)
	gormC.Db.AutoMigrate(UserTest{},UserLogTest{})
	tran:=&data_db_gorm_jinzhu.TranManager{Db: gormC.Db}
	//tx:=gormC.Db.Begin()
	tran.Begin()
	tx1:=tran.GetDb()
	defer func() {
		//tx.Commit()
		tran.Commit()
	}()
	tx:=tx1.Create(&UserTest{Name: "sqlite3"})
	if tx.Error!=nil{
		log.Printf("test gorm jinzhus sqlite3 inser fail,err:%s",tx.Error.Error())
		return
	}
	log.Println("test gorm jinzhus sqlite3 inser suc")
	var c dto.CountDto
	tx=tx1.Raw("select count(*) total from t_user_test where  id > ? and id< ?",0,1000).Scan(&c)
	log.Printf("test gorm jinzhus sqlite3 count %d ",*&c.Total)


}
func (g *GormJinzhuTet)testMysql(){
	gormC:=&config_gorm_jinzhu.GormConfig{}
	gormC.UpdateDb(datas.DbMysql,"root:wjp930514.@(192.168.1.2:3306)/samplesystem?charset=utf8mb4&parseTime=True&loc=Local",true)
	g.test1(gormC,"mysql")
}
func (*GormJinzhuTet)test1(gormC *config_gorm_jinzhu.GormConfig,db string){
	gormC.Db.AutoMigrate(UserTest{},UserLogTest{})
	tran:=&data_db_gorm_jinzhu.TranManager{Db: gormC.Db}
	//tx:=gormC.Db.Begin()
	tran.Begin()
	tx:=tran.GetDb()
	defer func() {
		//tx.Commit()
		tran.Commit()
	}()
	tx=tx.Create(&UserTest{Name: db})
	if tx.Error!=nil{
		log.Printf("test gorm jinzhus %s inser fail,err:%s",db,tx.Error.Error())
		return
	}
	log.Println("test gorm jinzhus  inser suc")
}
func(*GormJinzhuTet) testSqlserver()  {

}
func (*GormJinzhuTet)testPostgre(){

}
func (*GormJinzhuTet)testOracle(){

}