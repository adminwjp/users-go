package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	data "github.com/adminwjp/infrastructure-go/datas"
	"github.com/adminwjp/infrastructure-go/register_services"
	"github.com/adminwjp/infrastructure-go/register_services/consuls"
	"github.com/adminwjp/infrastructure-go/utils"
	"github.com/adminwjp/users-go/datas"
	"github.com/adminwjp/users-go/inits"
	"github.com/adminwjp/users-go/rpcs"
	service "github.com/adminwjp/users-go/services"
	web_gin_controller "github.com/adminwjp/users-go/webs/gins"
	web_http_controller "github.com/adminwjp/users-go/webs/https"
	"github.com/google/uuid"
	"io/fs"
	"io/ioutil"

	//web_http_controller "github.com/adminwjp/users-go/webs/https"
	"log"
	"math"
	"os"
	"strconv"
)
//包 不要 拆了 坑嗲玩意 能运行 但 运行错误
//go mod tidy
//好多驱动 没有 坑嗲 体积变小 了 手动 体积驱动

//linux
//linux run linux Binary was compiled with 'CGO_ENABLED=0', go-sqlite3
//https://blog.csdn.net/le_17_4_6/article/details/119332592
//linux release 


//sqlite fresh
//CC=x86_64-w64-mingw32-gcc go run main.go
//CGO_ENABLED=1  GOOS=windows GOARCH=amd64   CC=x86_64-w64-mingw32-gcc go run main.go
//CGO_ENABLED=1  GOOS=windows GOARCH=amd64   CC=x86_64-w64-mingw32-gcc go build
//CGO_ENABLED=1  GOOS=linux GOARCH=amd64   CC=x86_64-w64-mingw32-gcc go build


//mod err vender code 注释掉 go vendor 不使用 不然需要重新注释
//github.com/adminwjp/infrastructure-go

func Usage() {
	fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "\nParams:")
	fmt.Fprintln(os.Stderr, "  data_flag none 0 db 1 mong 2 es 3")
	fmt.Fprintln(os.Stderr, "  orm_flag none 0 bee 1 gorm jinzhu 2 gorm io  3")
	fmt.Fprintln(os.Stderr)
	os.Exit(0)
}
func main() {
	//最好 去 磁盘 里 看 自定义包 有时Ide 不出来
	//自定义包改变  依赖 编译 报错 位置 不对  难 排查
	//vendor 手动删除 手动复制
	//直接复制 过来 Ide 智能提示不准确
	//之前可以的 难道 混合使用 无法使用 gorm zhu gorm 只能加载 1 个 不然 啥都用不了 注释掉 或 打包成 zip 文件 bee orm
	//都是坑 不能加载  驱动 按理可以 最好 orm 3 选 1
/*	dto "github.com/adminwjp/infrastructure-go/dtos"
	util "github.com/adminwjp/users-go/infrastructure-go/utils"
	github.com/adminwjp/infrastructure-go/dtos
	github.com/adminwjp/infrastructure-go/dto

	github.com/adminwjp/users-go/infrastructure-go/utils
	github.com/adminwjp/users-go/infrastructure-go/util

	rm
	C:\Users\Administrator\go\pkg\mod\cache
	go get
	一个个排查

	混合以上操作 解决编译 报错
	编译 报错 提示错误太模糊了 不好找位置
	*/

	if false{
		utils.LogFile(utils.Hour)
	}
	/*_=models.ParseClassTpl(reflect.TypeOf(models.ClassTpl{}),
		reflect.TypeOf(models.FieldTpl{}),reflect.TypeOf(models.SmsConfigModel{}),
		reflect.TypeOf(models.PaySecrtConfigModel{}),reflect.TypeOf(models.EmailConfigModel{}),
		reflect.TypeOf(models.ExecModel{}),reflect.TypeOf(models.UserModel{}),
		reflect.TypeOf(models.UserLogModel{}),reflect.TypeOf(models.UserBasicModel{}),
		reflect.TypeOf(models.AdminModel{}),reflect.TypeOf(models.RoleModel{}),
		reflect.TypeOf(models.TableConfigModel{}),reflect.TypeOf(models.TableBean{}),
		reflect.TypeOf(models.ColumnBean{}),)*/
	//bu,err:=xml.Marshal(*classes)
	//if err!=nil{
		//log.Printf("xml parse ,err:%s",err.Error())
		//return
	//}
	//ioutil.WriteFile("E:/work/code_generator/xml/go_user.xml",bu,fs.ModeType|fs.ModePerm)
	//return
	flag.Usage = Usage
	//变量太多麻烦 不使用 使用配置更新
	var dataFlag int
	var ormFlag int
	var dbFlag int
	var connectionString string
	var enableRpc bool
	var grpcPort int
	var enableThrift bool
	var thriftPort int
	_ = strconv.Atoi
	_ = math.Abs
	flag.Usage = Usage
	flag.IntVar(&dataFlag, "data_flag", 1, "data_flag none 0 db 1 mong 2 es 3")
	flag.IntVar(&ormFlag, "orm_flag", 1, "orm_flag none 0 bee 1 gorm jinzhu 2 gorm io  3")
	flag.IntVar(&dbFlag, "db_flag", 1, "db_flag")
	flag.StringVar(&connectionString, "connection_string", "", "Specify the url")
	flag.BoolVar(&enableRpc, "enable_rpc", true, "Use framed transport")
	flag.IntVar(&grpcPort, "grpc_port", 4000, "Use http")
	flag.BoolVar(&enableThrift, "enable_thrift", true, "Use framed transport")
	flag.IntVar(&thriftPort, "thrift_port", 4001, "Use http")
	flag.Parse()
	args:= os.Args
	for i, v := range args {
		log.Printf("index => %d,v=> %s \n",i,v)
	}

	datas.GlobalConfig.EnableWeb=true
	datas.GlobalConfig.WebFlag=datas.WebGin
	datas.GlobalConfig.WebPort=8001



	datas.GlobalConfig.EnableSocket=false
	datas.GlobalConfig.SocketPort=8002

	datas.GlobalConfig.DataFlag=data.DataDb
	datas.GlobalConfig.OrmFlag=data.DbOrmBee
	datas.GlobalConfig.DbFlag=data.DbSqlite
	datas.GlobalConfig.AliasName="default"
	datas.GlobalConfig.ConnectionString="samplesystem.db"

	datas.GlobalConfig.EnableRpc=false
	datas.GlobalConfig.Ip="192.168.1.2"

	datas.GlobalConfig.EnableGrpc=true
	datas.GlobalConfig.GrpcPort=4000

	datas.GlobalConfig.EnableThrift=true
	datas.GlobalConfig.ThriftPort=4001

	datas.GlobalConfig.EnableRegisterService=false
	datas.GlobalConfig.RegisterServiceIp=datas.GlobalConfig.Ip
	datas.GlobalConfig.RegisterServicePort=8500
	datas.GlobalConfig.RegisterService=&register_services.ServiceInfo{
		Id: uuid.NewString(),
		Ip: datas.GlobalConfig.Ip,
		Port: datas.GlobalConfig.WebPort,
		ServiceName: "go_user",
		Tags: []string{"go-user","data orm(bee gorm-jinzhu gormio) db(sqlite(window liunx ex need linux build) postgre mysql sqlserver ) es mong ","rpc gprc thrift","mq rabbitmq kakfa","lock empty redis zookeeper consul local","retry empty local",
		"register service consul ",
			"web bee(support not run,single run) gin http","socket"},
	}
	f,err:=os.Open("config.json")
	if err!=nil{
		log.Printf("config.json read fail,err:%s",err.Error())
		f,err=os.Open("config.xml")
		if err!=nil{
			log.Printf("config.xml read fail,err:%s",err.Error())
		}else{
			bu,_:=ioutil.ReadFile("config.xml")
			xml.Unmarshal(bu,&datas.GlobalConfig)
		}
	}else{
		f1,_:=f.Stat()
		f1.Size()
		f.Close()
		bu,err:=ioutil.ReadFile("config.json")
		if err!=nil{
			log.Printf("config.json read fail,err:%s",err.Error())
		}else{
			json.Unmarshal(bu,&datas.GlobalConfig)
		}
	}


	bu,_:=json.Marshal(datas.GlobalConfig)
	ioutil.WriteFile("config.json",bu,fs.ModeType|fs.ModePerm)

	bu,err=xml.Marshal(datas.GlobalConfig)
	if err!=nil{
		log.Printf("xml seri fail,err:%s",err.Error())
	}
	ioutil.WriteFile("config.xml",bu,fs.ModeType|fs.ModePerm)
	datas.Init()
	//组合 太多 东西 运行 报错 无法使用
	//gorm query was empty
	log.Println("starting config ")
	inits.UpdateConfig1(datas.GlobalConfig)
	inits.CreateServiceInstance.Create()
	log.Println("create service success ")
	//   import cycle not allowed
	rpcs.CreateServiceInstance=ServeiceInstance1
	if datas.GlobalConfig.EnableRpc{
		inits.UpdateRpc(datas.GlobalConfig)
	}
	log.Println("started config ")

	//理论上 应该支持 对应 所有驱动
	//各种坑 能运行 但 不正常
	//单驱动 pass
	//data_db_gorm_jinzhu.Driver="sqlite"
	cl:= func() {
		utils.ListenApplicationExit(func() {
			for _, v := range datas.GlobalConfig.Exits {
				v()
			}
		})
	}
	if datas.GlobalConfig.EnableWeb{
	
		


		if datas.GlobalConfig.EnableRegisterService{
			consol1,err:=consuls.NewConsulServiceRegistry(datas.GlobalConfig.RegisterServiceIp,
				int(datas.GlobalConfig.RegisterServicePort),"")
			if err!=nil{
				log.Printf("reg service consul fail,err:%s",err.Error())

			}else{
				consol1.Register(*datas.GlobalConfig.RegisterService)
				datas.GlobalConfig.Exits["consul_reg"]=func() {
					consol1.Deregister()
				}
			}


		}
		cl()
		log.Printf("start web :%d",datas.GlobalConfig.WebFlag)
		switch  datas.GlobalConfig.WebFlag{
		//case datas.WebBee:
			//start load bee xx
			//web_bee_controller.BeeRouterImpl.Start(8001)
			//break
		case datas.WebHttp:
			web_http_controller.HttpRouterImpl.Start(datas.GlobalConfig.WebPort)
			break
		case datas.WebGin:
			web_gin_controller.GinRouterImpl.Start(datas.GlobalConfig.WebPort)
			break
		default:
			web_gin_controller.GinRouterImpl.Start(datas.GlobalConfig.WebPort)
			break
		}
		

	}else{
		log.Printf("not start web :%d",datas.GlobalConfig.WebFlag)
		cl()
	}


 	//code_gen.Test()
}
var ServeiceInstance1=&ServeiceInstance{}
type ServeiceInstance struct {

}
func (s *ServeiceInstance)AdminService () service.AdminService{
	return inits.CreateServiceInstance.AdminService()
}
func (s *ServeiceInstance)UserService () service.UserService{
	return inits.CreateServiceInstance.UserService()
}
func (s *ServeiceInstance)RoleService () service.RoleService{
	return inits.CreateServiceInstance.RoleService()
}
func (s *ServeiceInstance)EmailService () service.EmailService{
	return inits.CreateServiceInstance.EmailService()
}
func (s *ServeiceInstance)PaySecrtService () service.PaySecrtService{
	return inits.CreateServiceInstance.PaySecrtService()
}
func (s *ServeiceInstance)SmsService () service.SmsService{
	return inits.CreateServiceInstance.SmsService()
}
func (s *ServeiceInstance)ConfigService () service.ConfigService{
	return inits.CreateServiceInstance.ConfigService()
}
func (s *ServeiceInstance)RpcService () service.RpcService{
	return inits.CreateServiceInstance.RpcService()
}