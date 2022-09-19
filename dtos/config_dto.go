package dtos

type BaseConfigDto struct {
	Cache string //empty local redis
	Data string //bee orm gorm jinzhus gorm mong es
	Db string //sqlite mysql postgre sqlserver oracle
	Mq string //empty kafka rabbitmq
	Lock string //empty local redis zookeeper consul
	Rpc string //empty grpc  thrift
	Register string //empty consul
	Debug bool
	Web string //http gin bee

	ElasticSearchConnectionString string

	DbConnectionString string

	MongConnectionString string
	RedisConnectionString string
	KafkaConnectionString string
	RabbitConnectionString string
	ConsulConnectionString string
	EurekaConnectionString string
	ZookeeperConnectionString string
}
type JsonConfigDto struct {
	BaseConfigDto
	ElasticSearchConnectionStrings []string
	MasterDbConnectionStrings []string
	SalveDbConnectionStrings []string

	MongConnectionStrings []string
	RedisConnectionStrings []string


	KafkaConnectionStrings []string
	RabbitConnectionStrings []string
	ConsulConnectionStrings []string
	EurekaConnectionStrings []string
	ZookeeperConnectionStrings []string
}

type XmlConfigDto struct {
	BaseConfigDto
	ElasticSearchConnectionStrings []ConnectionStringDto
	MasterDbConnectionStrings []ConnectionStringDto
	SalveDbConnectionStrings []ConnectionStringDto

	MongConnectionStrings []ConnectionStringDto
	RedisConnectionStrings []ConnectionStringDto


	KafkaConnectionStrings []ConnectionStringDto
	RabbitConnectionStrings []ConnectionStringDto
	ConsulConnectionStrings []ConnectionStringDto
	EurekaConnectionStrings []ConnectionStringDto
	ZookeeperConnectionStrings []ConnectionStringDto
}

type ConnectionStringDto struct {
	ConnectionString string
}
