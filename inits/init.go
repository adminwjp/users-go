package inits

import (
	"github.com/adminwjp/infrastructure-go/caches/redises"
	config_bee_orm "github.com/adminwjp/infrastructure-go/configs/bee_orms"
	config_es "github.com/adminwjp/infrastructure-go/configs/ess"
	config_gorm "github.com/adminwjp/infrastructure-go/configs/gorms"
	config_gorm_jinzhu "github.com/adminwjp/infrastructure-go/configs/gorms/jinzhus"
	config_mong "github.com/adminwjp/infrastructure-go/configs/mongs"
	data "github.com/adminwjp/infrastructure-go/datas"
	data_es "github.com/adminwjp/infrastructure-go/datas/ess"
	data_mong "github.com/adminwjp/infrastructure-go/datas/mongs"
	"github.com/adminwjp/infrastructure-go/locks"
	consul_locks "github.com/adminwjp/infrastructure-go/locks/consuls"
	lock_redises "github.com/adminwjp/infrastructure-go/locks/redises"
	lock_zookeepers "github.com/adminwjp/infrastructure-go/locks/zookeepers"
	"github.com/adminwjp/infrastructure-go/mqs/kafkas"
	"github.com/adminwjp/infrastructure-go/mqs/rabbitmqs"
	rpc_grpc "github.com/adminwjp/infrastructure-go/rpcs/grpcs"
	rpc_thrift "github.com/adminwjp/infrastructure-go/rpcs/thrifts"
	dao_bee_orm_impl "github.com/adminwjp/users-go/daos/impls/bee_orm"
	dao_gorm_impl "github.com/adminwjp/users-go/daos/impls/gorm"
	dao_gorm_jinzhu_impl "github.com/adminwjp/users-go/daos/impls/gorm/jinzhu"
	"github.com/adminwjp/users-go/datas"
	rpc_grpc_impl "github.com/adminwjp/users-go/rpcs/grpc"
	rpc_thrift_impl "github.com/adminwjp/users-go/rpcs/thrift"
	service "github.com/adminwjp/users-go/services"
	service_impl "github.com/adminwjp/users-go/services/impls"
	service_mq_impl "github.com/adminwjp/users-go/services/impls/mqs"
	"github.com/adminwjp/users-go/sockets"
	"github.com/beego/beego/v2/client/orm"
	"github.com/go-redis/redis"
	"github.com/hashicorp/consul/api"
	"github.com/streadway/amqp"
	"log"
	"strconv"
)

var CreateServiceInstance =&CreateService{}
type CreateService struct {
	AdminService func() service.AdminService
	UserService func() service.UserService
	RoleService func() service.RoleService
	EmailService func() service.EmailService
	PaySecrtService func() service.PaySecrtService
	SmsService func() service.SmsService
	ConfigService func() service.ConfigService
	RpcService func() service.RpcService
}

func(s *CreateService) Create()  {
	s.AdminService= func() service.AdminService {
		if datas.GlobalConfig.EnableMq{
			return  service_mq_impl.GetAdminService()
		}
		return service_impl.GetAdminService()
	}
	s.UserService= func() service.UserService {
		if datas.GlobalConfig.EnableMq{
			return  service_mq_impl.GetUserService()
		}
		return service_impl.GetUserService()
	}
	/*	if mq{
			s.AdminService=service_mq_impl.GetAdminService
			s.UserService=service_mq_impl.GetUserService
		}else{
			s.AdminService=service_impl.GetAdminService
			s.UserService=service_impl.GetUserService
		}*/
	s.SmsService=service_impl.GetSmsService
	s.EmailService=service_impl.GetEmailService
	s.RoleService=service_impl.GetRoleService
	s.PaySecrtService=service_impl.GetPaySecrtService
	s.RpcService=service_impl.GetRpcService
	s.ConfigService=service_impl.GetConfigService
}
func  UpdateConfig1(cfg *datas.DataConfig){
	updateDataConfig(cfg)

	if cfg.EnableRedis{
		UpdateRedis(cfg)
	}
	UpdateSocket(cfg)
	if cfg.EnableRetry{

	}

	if cfg.EnableMq{
		if cfg.EnableRabbitmq{
			UpdateRabbitmq(cfg)
		}
		if cfg.EnableKafka{
			UpdateKafka(cfg)
		}
	}

	if cfg.EnableEs{
		if cfg.DataFlag!=data.DataEs{
			UpdateDataEs(cfg)
		}

	}

	//if cfg.EnableLock{
		UpdateLock(cfg)
	//}

	if cfg.EnableMong{
		if cfg.DataFlag!=data.DataMong{
			UpdateDataMong(cfg)
		}
	}

	/*if cfg.EnableRpc{
		UpdateRpc(cfg)
	}*/

}
func UpdateSocket(cfg *datas.DataConfig)  {
	if cfg.EnableSocket{
		go func() {
			sockets.SocketInstance.ServerStart(cfg.Ip+":"+strconv.Itoa(cfg.SocketPort))
			sockets.Init()
			cfg.Exits["socket"]= func() {
				err:=sockets.SocketInstance.Server.Close()
				if err!=nil{
					log.Printf("socket server close fail,err:%s",err.Error())
				}
			}
		}()
		sockets.AdminSInstance.UserService=CreateServiceInstance.AdminService
		sockets.AdminSInstance.BaseUserService= func() service.BaseUserService {
			return sockets.AdminSInstance.UserService()
		}

		sockets.UserSInstance.UserService=CreateServiceInstance.UserService
		sockets.UserSInstance.BaseUserService= func() service.BaseUserService {
			return sockets.UserSInstance.UserService()
		}
	}
}
func UpdateRpc(cfg *datas.DataConfig)  {
	if cfg.EnableGrpc{

		rpc_grpc_impl.Init()
		rpc_grpc_impl.StartGRpc(":"+strconv.Itoa(cfg.GrpcPort),cfg.Ip)
		cfg.Exits["grpc"]=func() {
			rpc_grpc.GrpcInstance.Server.Stop()
		}
	}
	if cfg.EnableThrift{
		rpc_thrift_impl.Init()
		rpc_thrift_impl.StartThrift(":"+strconv.Itoa(cfg.ThriftPort),cfg.Ip)
		cfg.Exits["thrift"]=func() {
			rpc_thrift.ThiriftInstance.Server.Stop()
		}
	}
}
var lock1 locks.Lock
var EmptyLockInstance=&locks.EmptyLock{}
var LockInstance=&locks.LockImpl{}
var emptyLock =EmptyLockInstance

var zookeeperLock =&lock_zookeepers.ZookeeperLock{}

var consulLock =&consul_locks.ConsulLock{}
var redisLock =&lock_redises.RedisLock{}
func UpdateLock(cfg *datas.DataConfig)  {

	switch cfg.Lock {
	case datas.LockEmpty:
		lock1=emptyLock
		break
	case datas.LockLocal:
		lock1=LockInstance
		break
	case datas.LockZookeeper:
		{
			var f=false
			if zookeeperLock.Client!=nil{
				zookeeperLock.Client.Close()
				zookeeperLock.Client=nil
				f=true
				delete(cfg.Exits,"zookeeper")
			}
			zookeeperLock.Conn([]string{cfg.LockCfg.ConnectionString},"")
			if !f{
				zookeeperLock.StartCleanTimeoutThread()
				cfg.Exits["zookeeper"]=func() {
					rpc_thrift.ThiriftInstance.Server.Stop()
				}
			}
			lock1=zookeeperLock
		}
		break
	case datas.LockConsul:
		{
			var f=false
			if consulLock.Client!=nil{
				consulLock.Client=nil
				f=true
			}
			consulLock.Client,_=api.NewClient(&api.Config{Address: cfg.LockCfg.ConnectionString})
			if !f{
				consulLock.StartCleanTimeoutThread()
			}

		}
		break
	case datas.LockRedis:
		if cfg.EnableRedis&&
			cfg.RedisCfg.ConnectionString!=cfg.LockCfg.ConnectionString||!cfg.EnableRedis{

			if redisLock!=nil{
				redisLock.Client.Close()
				redisLock.Client=nil
				delete(cfg.Exits,"redisLock")
			}
			redisLock.Client=redis.NewClient(&redis.Options{
				Addr: cfg.LockCfg.ConnectionString,
				DB: 0,
			})
			cfg.Exits["redisLock"]=func() {
				rpc_thrift.ThiriftInstance.Server.Stop()
			}
		}

		break
	default:lock1=emptyLock
		break
	}
	service_impl.LockInstance=lock1
}
var kakfa=&kafkas.KafkaMq{}

func CloseKafka(){
	if kakfa.Consumer!=nil{
		kakfa.Consumer.Close()
		kakfa.Consumer=nil
	}
	if kakfa.Producer!=nil{
		kakfa.Producer.Close()
		kakfa.Producer=nil
	}
}
func UpdateKafka(cfg *datas.DataConfig)  {
	if kakfa!=nil{
		CloseKafka()
		delete(cfg.Exits,"kakfa")
	}
	kakfa.Consumer=kakfa.CreateConsumer([]string{cfg.MqCfg.ConnectionString})
	kakfa.Producer=kakfa.CreateProducer([]string{cfg.MqCfg.ConnectionString})
	service_mq_impl.MqInstance=kakfa
	cfg.Exits["kakfa"]=func() {
		CloseKafka()
	}
	service_mq_impl.MqInstance=kakfa
}
var rabbitmq=&rabbitmqs.RabbitMQ{}
func CloseRabbitmq(){
	if rabbitmq.Conn!=nil{
		rabbitmq.Conn.Close()
		rabbitmq.Conn=nil
	}
}
func UpdateRabbitmq(cfg *datas.DataConfig)  {

	if rabbitmq!=nil{
		CloseRabbitmq()
		delete(cfg.Exits,"rabbitmq")
	}
	var err error
	//获取connection
	rabbitmq.Conn, err = amqp.Dial(cfg.MqCfg.ConnectionString)
	failOnErr(err, "failed to connect rabb"+
		"itmq!")
	//获取channel
	rabbitmq.Channel, err = rabbitmq.Conn.Channel()
	failOnErr(err, "failed to open a channel")
	service_mq_impl.MqInstance=rabbitmq
	cfg.Exits["rabbitmq"]=func() {
		CloseRabbitmq()
	}
	service_mq_impl.MqInstance=rabbitmq
}
func  failOnErr(err error, message string) {
	if err != nil {
		log.Println("%s:%s", message, err)
		//panic(fmt.Sprintf("%s:%s", message, err))
	}
}
var redisCache *redises.RedisCache=&redises.RedisCache{}

func  CloseRedis()  {
	redisCache.Client.Close()
	redisCache.Client=nil
}
func UpdateRedis(cfg *datas.DataConfig)  {

	if redisCache.Client!=nil{
		CloseRedis()
		delete(cfg.Exits,"redisCache")
	}

	redisCache.Client=redisCache.GetClient(cfg.RedisCfg.ConnectionString,"",0)
	service_impl.RedisCacheInstance=redisCache
	cfg.Exits["redisCache"]=func() {
		if redisCache.Client!=nil{CloseRedis()}
	}
}
func  CloseOrmJinzhuGorm()  {
	config_gorm_jinzhu.GormConfigInstance.Db.Close()
	config_gorm_jinzhu.GormConfigInstance.Db=nil
}
func UpdateDbOrmJinzhuGorm(cfg *datas.DataConfig)  {

	if config_gorm_jinzhu.GormConfigInstance.Db!=nil{
		CloseOrmJinzhuGorm()
		delete(cfg.Exits,"config_gorm_jinzhu")
	}
	config_gorm_jinzhu.GormConfigInstance.UpdateDb(cfg.DbFlag,datas.GlobalConfig.ConnectionString,true)
	dao_gorm_jinzhu_impl.InitMigr()
	dao_gorm_jinzhu_impl.CreateView()
	cfg.Exits["config_gorm_jinzhu"]=func() {
		if config_gorm_jinzhu.GormConfigInstance.Db!=nil{CloseOrmJinzhuGorm()}
	}
}
func  CloseOrmGormio()  {
	config_gorm.GormConfigInstance.Db=nil
}
func UpdateDbOrmGormio(cfg *datas.DataConfig)  {
	if config_gorm.GormConfigInstance.Db!=nil{
		CloseOrmGormio()
		delete(cfg.Exits,"config_gorm")
	}
	config_gorm.GormConfigInstance.UpdateDb(cfg.DbFlag,cfg.ConnectionString,true)
	dao_gorm_impl.InitMigr()
	dao_gorm_impl.CreateView()
	cfg.Exits["config_gorm"]=func() {
		if config_gorm.GormConfigInstance.Db!=nil{
			CloseOrmGormio()
		}
	}
}
var beeorm bool=false
func UpdateDbOrmBee(cfg *datas.DataConfig)  {
	if beeorm{
		//orm.RunSyncdb("default", false, false)
		beeorm=false
		delete(cfg.Exits,"config_bee_orm")
	}
	config_bee_orm.BeeOrmConfigInstance.UpdateDb(cfg.DbFlag,cfg.ConnectionString,cfg.AliasName,true)

	if !beeorm{
		dao_bee_orm_impl.InitMigr()
		cfg.Exits["config_bee_orm"]=func() {
			orm.RunSyncdb("default", false, false)
		}
	}
	beeorm=true
}
func CloseEs()  {
	config_es.EsConfigInstance.Client=nil
}
func UpdateDataEs(cfg *datas.DataConfig)  {
	if config_es.EsConfigInstance.Client!=nil{
		CloseEs()
		delete(cfg.Exits,"config_es")
	}
	m:=data_es.NewElasticUtil()
	m.Client=data_es.GetClient(cfg.ConnectionString)
	config_es.EsConfigInstance.Client=m.Client
	cfg.Exits["config_es"]=func() {
		if config_es.EsConfigInstance.Client!=nil{
			CloseEs()
		}
	}
}
func  CloseMong()  {
	config_mong.MongConfigInstance.Session.Close()
	config_mong.MongConfigInstance.Session=nil
}
func UpdateDataMong(cfg *datas.DataConfig)  {

	if config_mong.MongConfigInstance.Session!=nil{
		CloseMong()
		delete(cfg.Exits,"data_mong")
	}
	m:=data_mong.MongHelper{}
	m.Session=m.Conn1([]string{cfg.ConnectionString},"samplesystem","","")
	config_mong.MongConfigInstance.Session=m.Session
	cfg.Exits["data_mong"]=func() {
		if config_mong.MongConfigInstance.Session!=nil{
			CloseMong()
		}
	}
}
//   import cycle not allowed
func  updateDataConfig(cfg *datas.DataConfig){
	switch datas.GlobalConfig.DataFlag {
	case data.DataDb:
		switch  cfg.OrmFlag {
		case data.DbOrmJinzhuGorm:
			UpdateDbOrmJinzhuGorm(cfg)
			break
		case data.DbOrmGormio:
			UpdateDbOrmGormio(cfg)
			break
		case data.DbOrmBee:
			UpdateDbOrmBee(cfg)
			break
		default:break

		}
		break
	case data.DataEs:
		UpdateDataEs(cfg)
		break
	case data.DataMong:
		UpdateDataMong(cfg)
		break
	default:
		break
	}


}
