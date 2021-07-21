package main

import (
	"fmt"
	"rabbitmq/example/queue"
)

func main() {
	var err error
	mq := queue.NewRabbitMq("", "这是个测试", queue.Url)
	err = mq.Connection()
	err = mq.NewQueue()
	err = mq.SendMsg("12343435")
	fmt.Println(err)
	fmt.Println(mq)
}
