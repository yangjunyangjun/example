package main

import (
	"fmt"
	"rabbitmq/example/simple"
)

func main() {
	var err error
	mq := simple.NewRabbitMq("", "这是个测试", simple.Url)
	err = mq.Connection()
	err = mq.NewQueue()
	err = mq.SendMsg("12343435")
	fmt.Println(err)
	fmt.Println(mq)
}
