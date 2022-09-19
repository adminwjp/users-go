package web

import (
	datas1 "github.com/adminwjp/infrastructure-go/datas"
	"github.com/adminwjp/infrastructure-go/dtos"
	"github.com/adminwjp/infrastructure-go/webs"
	"github.com/adminwjp/users-go/datas"
	"github.com/adminwjp/users-go/inits"
	service_impl "github.com/adminwjp/users-go/services/impls"
	"log"
	"reflect"
	"strconv"
	"strings"
)

type ConfigCtrl struct {

}
func (ctrl *ConfigCtrl)Get(httpWeb webs.HttpWeb){
	httpWeb.Response(200,dtos.ResponseDataDto{
		Status: true, Code: 200, Msg: "success",Data: datas.GlobalConfig,
	})
}
func (ctrl *ConfigCtrl)Set(httpWeb webs.HttpWeb){
	var m1 datas.DataConfig
	err:=httpWeb.ShouldBind(m1)
	if err!=nil||reflect.DeepEqual(*&m1,datas.DataConfig{}){
		res:=dtos.ResponseDto{Status: false,Code:400,Msg:"fail"}
		httpWeb.Response(200,res)
		return
	}
	httpWeb.Response(200,dtos.ResponseDataDto{
		Status: true, Code: 200, Msg: "success",Data: m1,
	})
}

//mong
//config/set/mong?o=bee&c=/data.sqlite3&mc&sc

//c sc ms db redis mq  zookeeper consul

//config/set/db?o=bee&c=/data.sqlite3&mc&sc&d=sqlite config/set/db?o=gorm_jinzhu  config/set/db?o=gorm
//o gorm gorm_jinzhu bee d sqlite mysql sqlserver postgre oracle

//config/set/cache?o=load&e=true&c=&mc=&sc= config/set/cache?o=redis&e=true
//redis local empty bee

//config/set/rpc?o=grpc&e=true&p=50000 config/set/rpc?o=thrift&e=true
//config/set/mq?o=kafka&e=true config/set/rpc?o=rabbitmq&e=true
//config/set/lock?o=redis&e=true  o locacl zookeeper consul redids

func (ctrl *ConfigCtrl)SetFlag(httpWeb webs.HttpWeb){
	flag:=httpWeb.GetPathString(":flag")

	log.Println("flag=>"+flag)
	if flag=="es"{
		c:=datas.GlobalConfig.ConnectionString
		ctrl.updateconn2(httpWeb,datas.GlobalConfig)

		datas.GlobalConfig.DataFlag=datas1.DataEs
		if c!=datas.GlobalConfig.ConnectionString{
			for k, _ := range service_impl.DaoServiceInstance.Cruds {
				delete(service_impl.DaoServiceInstance.Cruds,k)
			}
			inits.UpdateDataEs(datas.GlobalConfig)
		}

	}else if flag=="mong"{
		c:=datas.GlobalConfig.ConnectionString
		ctrl.updateconn2(httpWeb,datas.GlobalConfig)
		datas.GlobalConfig.DataFlag=datas1.DataMong

		if c!=datas.GlobalConfig.ConnectionString{
			for k, _ := range service_impl.DaoServiceInstance.Cruds {
				delete(service_impl.DaoServiceInstance.Cruds,k)
			}
			inits.UpdateDataMong(datas.GlobalConfig)
		}
	}else if flag=="db" || flag=="orm"{
		ctrl.setdb(httpWeb)

	}else if flag=="cache"{
		o:=httpWeb.GetQueryString("o")
		e:=httpWeb.GetQueryString("e")
		if e!=""{
			e1:=strconv.CanBackquote(e)
			if o=="load"{
				datas.GlobalConfig.LoadCache=e1
			}
		}
		if o=="redis"{
			if e!=""{
				e1:=strconv.CanBackquote(e)
				datas.GlobalConfig.EnableRedis=e1
			}
			c:=datas.GlobalConfig.RedisCfg.ConnectionString
			ctrl.updateconn(httpWeb,datas.GlobalConfig.RedisCfg)
			if c!=datas.GlobalConfig.RedisCfg.ConnectionString{
				inits.UpdateRedis(datas.GlobalConfig)
			}
		}
		//userService.loadCache()
		//adminService.loadCache()
	}else if flag=="rpc"{
		ctrl.updaterpc(httpWeb)
	}else if flag=="mq"{
		datas.GlobalConfig.EnableMq=true
		e:=httpWeb.GetQueryString("e")
		if e!=""{
			e1:=strconv.CanBackquote(e)
			datas.GlobalConfig.EnableMq=e1
		}
		o:=httpWeb.GetQueryString("o")
		c:=datas.GlobalConfig.MqCfg.ConnectionString
		if o=="rabbitmq"{
			datas.GlobalConfig.EnableKafka=true
			datas.GlobalConfig.EnableRabbitmq=false
			ctrl.updateconn(httpWeb,datas.GlobalConfig.MqCfg)
			if c!=datas.GlobalConfig.MqCfg.ConnectionString{
				inits.UpdateRabbitmq(datas.GlobalConfig)
			}
		}else if o=="kafka"{
			datas.GlobalConfig.EnableKafka=false
			datas.GlobalConfig.EnableRabbitmq=true
			ctrl.updateconn(httpWeb,datas.GlobalConfig.MqCfg)
			if c!=datas.GlobalConfig.MqCfg.ConnectionString{
				inits.UpdateKafka(datas.GlobalConfig)
			}
		}
		//MqUtil.Conn()
		//userMqService.init()
	}else if flag=="lock"{
		l:=datas.GlobalConfig.Lock
		if flag=="local"{
			datas.GlobalConfig.Lock=datas.LockLocal
		}else if flag=="redis"{
			datas.GlobalConfig.Lock=datas.LockRedis
		}else if flag=="zookeeper"{
			datas.GlobalConfig.Lock=datas.LockZookeeper
		}else if flag=="consul"{
			datas.GlobalConfig.Lock=datas.LockConsul
		}else {
			datas.GlobalConfig.Lock=datas.LockEmpty
		}
		if l!=datas.GlobalConfig.Lock{
			inits.UpdateLock(datas.GlobalConfig)
		}
	} else if flag=="delete_cache"{
		//delete redis all key
		//RedisClient.FlushDB() //clean redis data
		//RedisClient.FlushAll() //clean a db
	}else if flag=="delete_all_key"{
		//delete redis all key
		//RedisClient.FlushDB() //clean redis data
	}
	httpWeb.Response(200,dtos.ResponseDataDto{
		Status: true, Code: 200, Msg: "success",Data: datas.GlobalConfig,
	})
}
func  (ctrl *ConfigCtrl) updaterpc(httpWeb webs.HttpWeb) {
	o:=httpWeb.GetQueryString("o")
	e:=httpWeb.GetQueryString("e")
	p:=httpWeb.GetQueryString("p")
	if e!=""{
		e1:=strconv.CanBackquote(e)
		datas.GlobalConfig.EnableRpc=e1
	}
	if o=="grpc"{
		p1,_:=strconv.Atoi(p)
		if p1>0{
			if datas.GlobalConfig.GrpcPort!=p1{
				datas.GlobalConfig.GrpcPort=p1
				datas.GlobalConfig.EnableGrpc=true
				datas.GlobalConfig.EnableThrift=false
				inits.UpdateRpc(datas.GlobalConfig)
			}

		}
	}
	if o=="thrift"{
		p1,_:=strconv.Atoi(p)
		if p1>0{
			if datas.GlobalConfig.ThriftPort!=p1 {
				datas.GlobalConfig.ThriftPort = p1
				datas.GlobalConfig.EnableGrpc = false
				datas.GlobalConfig.EnableThrift = true
				inits.UpdateRpc(datas.GlobalConfig)
			}
		}
	}
}
func (ctrl *ConfigCtrl) updateconn(httpWeb webs.HttpWeb ,conn *datas.DataCfg){
	connStr:=httpWeb.GetQueryString("c")

	mconnStrs:=httpWeb.GetQueryString("mc")
	sconnStrs:=httpWeb.GetQueryString("sc")
	if connStr!=""{
		conn.ConnectionString=connStr
	}
	if connStr!=""{
		conn.MasterConnectionStrings=strings.Split(mconnStrs,";")
	}
	if connStr!=""{
		conn.SlaveConnectionStrings=strings.Split(sconnStrs,";")
	}
}
func (ctrl *ConfigCtrl) updateconn2(httpWeb webs.HttpWeb ,conn *datas.DataConfig){
	connStr:=httpWeb.GetQueryString("c")
	log.Println("c => "+connStr)
	mconnStrs:=httpWeb.GetQueryString("mc")
	sconnStrs:=httpWeb.GetQueryString("sc")
	if connStr!=""{
		conn.ConnectionString=connStr
	}
	if connStr!=""{
		conn.MasterConnectionStrings=strings.Split(mconnStrs,";")
	}
	if connStr!=""{
		conn.SlaveConnectionStrings=strings.Split(sconnStrs,";")
	}
}

func (ctrl *ConfigCtrl) setdb(httpWeb webs.HttpWeb){
	if datas.GlobalConfig.DataFlag!=datas1.DataDb{
		return
	}
	
	ctrl.updateconn2(httpWeb,datas.GlobalConfig)
	
	
	o:=httpWeb.GetQueryString("o")
	d:=httpWeb.GetQueryString("d")
	log.Printf("o => %s,d => %s",o,d)
	datas.GlobalConfig.DataFlag=datas1.DataDb
	if d=="mysql"||d==""{
		datas.GlobalConfig.DbFlag= datas1.DbMysql
	}else if d=="sqlserver"{
		datas.GlobalConfig.DbFlag= datas1.DbSqlserver
	}else if d=="oracle"{
		datas.GlobalConfig.DbFlag= datas1.DbOracle
	}else if d=="sqlite"{
		datas.GlobalConfig.DbFlag= datas1.DbSqlite
	}else if d=="postgre"{
		datas.GlobalConfig.DbFlag= datas1.DbPostgre
	}
	switch o {
	case "bee":
		datas.GlobalConfig.OrmFlag= datas1.DbOrmBee
		/*service_impl.DaoServiceInstance.Data= func() daos.Dao {

		}*/
		inits.UpdateDbOrmBee(datas.GlobalConfig)
		break
	case "gorm_jinzhu":
		datas.GlobalConfig.OrmFlag= datas1.DbOrmJinzhuGorm
			inits.UpdateDbOrmJinzhuGorm(datas.GlobalConfig)
		break
	case "gorm":
		datas.GlobalConfig.OrmFlag= datas1.DbOrmGormio
			inits.UpdateDbOrmGormio(datas.GlobalConfig)
		break
	default:
		break
	}
	
	//datas.GlobalConfig.EnableMq=false
	//userService.enableUpdateIp=true
	//userService.enableLog=true
}


