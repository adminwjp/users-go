package tests

var sqliteTest=true
var mysqlTest=false
var sqlserverTest=false
var postgreTest=false
var oracleTest=false
//CGO_ENABLED=1  GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc go test  -count=1 -v  test/*.go

//go test  -count=1 -v   tests/bee_orm_test.go \
//tests/gorm_jinzhu_test.go tests/gorm_test.go tests/init.go  tests/test_kafka.go \
//tests/test_rabbitmq.go tests/test_model.go  tests/user_sample_test.go

//go test  -count=1 -v tests/test_rabbitmq.go