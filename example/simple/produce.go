package simple

import (
	"fmt"
	"github.com/streadway/amqp"
)

var url = "amqp://ledger:ledger@42.194.162.161:5672/ledger" //user:ledger   password:ledger vhost:ledger

func Produce() {
	//connection mq
	produce, err := amqp.Dial(url)
	if err != nil {
		fmt.Println(err)
	}
	//获取channel
	c, err := produce.Channel()
	if err != nil {
		fmt.Println(err)
	}

	//创建队列queue
	q, err := c.QueueDeclare(
		"test", //队列名称
		false,  // 是否持久化
		false,  //是否自动删除
		false,  //是否具有排他性
		false,  //是否阻塞
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	c.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("这是一个测试"),
		},

	)
}
