package tests

import (
	"github.com/adminwjp/infrastructure-go/mqs/rabbitmqs"
	"testing"
)

var rabbitMQ *rabbitmqs.RabbitMQ

func TestRabbitmq(t *testing.T) {
	url:="amqp://rabbitmq:rabbitmq@192.168.1.6:15672"
	//url:="amqp://guest:guest@192.168.1.6:15672"
	rabbitMQ=rabbitmqs.NewRabbitMQSimple(url)
	if rabbitMQ.Conn==nil{
		t.Error("conn rabbitmq fail")
	}
}
