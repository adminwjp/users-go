[xml format]
https://c.runoob.com/front-end/710/

[json format]
https://www.bejson.com/

DataFlag 1
E:\\work\\utility\\db\\samplesystem.db
DataFlag 2
127.0.0.1:27017


dialet=sqlite3
addrs=E:/work/utility/db/samplesystem.sqlite3

#dialet=mysql
#addrs=root:pwd@(192.168.1.4:3306)/samplesystem?charset=utf8

#dialet=postgres
#addrs=host=localhost user=user password=pwd dbname=samplesystem port=9920 sslmode=disable TimeZone=Asia/Shanghai

#dialet=sqlserver
#addrs=sqlserver://user:pwd@localhost:9930?database=samplesystem
#beeorm
#addrs=server=localhost;database=samplesystem;user id=user;password=pwd;port=9930;encrypt=disable

#addrs=database/password@127.0.0.1:1521/XE

//mong
config/set/mong?c=127.0.0.1:27017&mc&sc
//c sc ms db redis mq  zookeeper consul

//config/set/db?o=bee&c=/data.sqlite3&mc&sc&d=sqlite config/set/db?o=gorm_jinzhu  config/set/db?o=gorm
//o => gorm gorm_jinzhu bee 
//d => sqlite mysql sqlserver postgre oracle

//config/set/cache?o=load&e=true&c=&mc=&sc= config/set/cache?o=redis&e=true
//redis local empty bee

//config/set/rpc?o=grpc&e=true&p=50000 config/set/rpc?o=thrift&e=true
//config/set/mq?o=kafka&e=true config/set/rpc?o=rabbitmq&e=true
//config/set/lock?o=redis&e=true  o locacl zookeeper consul redids


```go
package datas

type  DataFlag int
const(
	DataNone DataFlag=iota
	DataDb
	DataMong
	DataEs
)

type DbFalg int

const  (
	DbNone DbFalg=iota
	DbSqlite
	DbMysql
	DbSqlserver
	DbPostgre
	DbOracle
	DbTidb
)

type DbOrmFlag int

const  (
	DbOrmNone DbOrmFlag=iota
	DbOrmBee
	DbOrmGormio
	DbOrmJinzhuGorm
)

```

```go
package datas

import (
	"github.com/adminwjp/infrastructure-go/datas"
	"github.com/adminwjp/infrastructure-go/register_services"
)

var GlobalConfig=&DataConfig{Exits: make(map[string]func(),0)}

type LockFlag int

const  (
	LockRedis LockFlag=iota
	LockEmpty
	LockLocal
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
	CacheBee CacheFlag=iota
	CacheRedis
	CacheEmpty
	CacheLocal
	CacheFileMap
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
	WebPort int ` xml:"attr"`

	EnableRegisterService bool ` xml:",attr"`
	RegisterService *register_services.ServiceInfo
	RegisterServiceIp string ` xml:",attr"`
	RegisterServicePort uint ` xml:",attr"`

	EnableSocket bool ` xml:",attr"`
	SocketPort int ` xml:",attr"`

	DbFlag datas.DbFalg ` xml:",attr"`
	OrmFlag datas.DbOrmFlag ` xml:",attr"`
	DataFlag datas.DataFlag ` xml:",attr"`
	ConnectionString string ` xml:",attr"`
	MasterConnectionStrings []string ` xml:"MasterConnectionStrings>MasterConnectionString"`
	SlaveConnectionStrings []string ` xml:"SlaveConnectionStrings>SlaveConnectionString"`

	EnableDb bool ` xml:",attr"`

	EnableRetry bool ` xml:",attr"`
	RetryFlag RetryFlag ` xml:",attr"`

	EnableLock bool ` xml:",attr"`
	Lock LockFlag ` xml:",attr"`
	LockCfg *DataCfg ` xml:",attr"`

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
```