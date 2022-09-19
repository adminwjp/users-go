package datas

import (
	"github.com/adminwjp/infrastructure-go/datas"
	"github.com/adminwjp/infrastructure-go/register_services"
)

type IClose interface {

}
var GlobalConfig=&DataConfig{
	Exits: make(map[string]func(),0),
	LockCfg: &DataCfg{},
	MqCfg: &DataCfg{},
	EsCfg: &DataCfg{},
	MongCfg: &DataCfg{},
	RedisCfg: &DataCfg{},

}

func Init()  {
	GlobalConfig.Exits= make(map[string]func(),0)
	if GlobalConfig.LockCfg==nil{
		GlobalConfig.LockCfg= &DataCfg{}
	}
	if GlobalConfig.MqCfg==nil{
		GlobalConfig.MqCfg= &DataCfg{}
	}
	if GlobalConfig.EsCfg==nil{
		GlobalConfig.EsCfg= &DataCfg{}
	}
	if GlobalConfig.MongCfg==nil{
		GlobalConfig.MongCfg= &DataCfg{}
	}
	if GlobalConfig.RedisCfg==nil{
		GlobalConfig.RedisCfg= &DataCfg{}
	}
}
type LockFlag int

const  (
	LockEmpty LockFlag=iota
	LockLocal
	LockRedis
	LockZookeeper
	LockConsul
)

type WebFlag int

const  (
	WebBee WebFlag=iota
	WebGin
	WebHttp
)
type CacheFlag int

const  (
	 CacheEmpty CacheFlag=iota
	CacheLocal
	CacheBee
	CacheFileMap
	CacheRedis
)

type RetryFlag int

const  (
	RetryEmpty RetryFlag=iota
	RetryLocal
	RetryRemote
)
type DataConfig struct {
	Exits map[string]func() ` xml:"-" json:"-"`

	EnableWeb bool ` xml:",attr"`
	WebFlag WebFlag ` xml:",attr"`
	WebPort int ` xml:",attr"`

	EnableRegisterService bool ` xml:",attr"`
	RegisterService *register_services.ServiceInfo
	RegisterServiceIp string ` xml:",attr"`
	RegisterServicePort uint ` xml:",attr"`

	EnableSocket bool ` xml:",attr"`
	SocketPort int ` xml:",attr"`

	DbFlag datas.DbFalg ` xml:",attr"`
	OrmFlag datas.DbOrmFlag ` xml:",attr"`
	DataFlag datas.DataFlag ` xml:",attr"`
	//bee orm  default
	AliasName string ` xml:",attr"`
	ConnectionString string ` xml:",attr"`
	MasterConnectionStrings []string ` xml:"MasterConnectionStrings>MasterConnectionString"`
	SlaveConnectionStrings []string ` xml:"SlaveConnectionStrings>SlaveConnectionString"`

	EnableDb bool ` xml:",attr"`

	EnableRetry bool ` xml:",attr"`
	RetryFlag RetryFlag ` xml:",attr"`

	EnableLock bool ` xml:",attr"`
	Lock LockFlag ` xml:",attr"`
	LockCfg *DataCfg

    EnableMq bool ` xml:",attr"`
	EnableRabbitmq bool ` xml:",attr"`
	EnableKafka bool ` xml:",attr"`
	MqCfg *DataCfg

	EnableEs bool ` xml:",attr"`
	EsCfg *DataCfg

	EnableMong bool ` xml:",attr"`
	MongCfg *DataCfg

	EnableRedis bool ` xml:",attr"`
	RedisCfg *DataCfg

	EnabelCache bool ` xml:",attr"`
	CacheFlag CacheFlag  ` xml:",attr"`

	LoadCache bool ` xml:",attr"`

	EnableRpc bool ` xml:",attr"`
	Ip string ` xml:",attr"`

	EnableGrpc bool ` xml:",attr"`
	GrpcPort int ` xml:",attr"`

	EnableThrift bool ` xml:",attr"`
	ThriftPort int ` xml:",attr"`
}
type DataCfg struct {
	ConnectionString string ` xml:",attr"`
	MasterConnectionStrings []string ` xml:"MasterConnectionStrings->MasterConnectionString"`
	SlaveConnectionStrings []string ` xml:"SlaveConnectionStrings->SlaveConnectionString"`

}