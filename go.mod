module github.com/adminwjp/users-go

go 1.16

// 报错不对 一个个解决
// replace github.com/adminwjp/infrastructure-go => ./github.com/adminwjp/infrastructure-go
//replace github.com/adminwjp/infrastructure-go => E:\software\go\src\github.com\adminwjp\infrastructure-go

replace github.com/adminwjp/infrastructure-go => E:\work\utility\Utility-for-go\infrastructure

require (
	github.com/adminwjp/infrastructure-go v0.0.0-00010101000000-000000000000
	github.com/apache/thrift v0.16.0
	github.com/beego/beego/v2 v2.0.4
	github.com/dchest/captcha v1.0.0
	github.com/gin-contrib/sessions v0.0.5
	github.com/gin-gonic/gin v1.8.1
	github.com/go-playground/locales v0.14.0
	github.com/go-playground/universal-translator v0.18.0
	github.com/go-playground/validator/v10 v10.11.0
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/google/uuid v1.3.0
	github.com/hashicorp/consul/api v1.13.0
	github.com/jinzhu/gorm v1.9.16
	github.com/olivere/elastic/v7 v7.0.32
	github.com/streadway/amqp v1.0.0
	github.com/stretchr/testify v1.8.0
	github.com/swaggo/files v0.0.0-20220610200504-28940afbdbfe
	github.com/swaggo/gin-swagger v1.5.1
	google.golang.org/grpc v1.48.0
	google.golang.org/protobuf v1.28.0
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/go-playground/validator.v9 v9.31.0
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
	gorm.io/gorm v1.23.8
)
