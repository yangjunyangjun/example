package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"rabbitmq/example/change"
	"rabbitmq/example/queue"
	"testing"
	"time"
)

func TestProduce(t *testing.T) {

	mq := change.NewRabbitMq("test1", "", "123")
	_, err := mq.NewChannel(queue.Url)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = mq.NewExchange("direct")
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < 100000; i++ {
		err = mq.SendMsg(fmt.Sprintf("sbsbsbsbsbsbbsbsbsb%d", i))
		time.Sleep(1)
		fmt.Println(err)
	}
}

func TestConsume(t *testing.T) {
	mq := change.NewRabbitMq("test1", "", "123")
	_, err := mq.NewChannel(queue.Url)
	fmt.Println(err)
	err = mq.Consume("direct",test)
	fmt.Println(err)
}

func test(delivery amqp.Delivery) {
	fmt.Println(string(delivery.Body))
	delivery.Ack(true)
}
