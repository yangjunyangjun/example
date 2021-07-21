package simple

import (
	"fmt"
	"github.com/streadway/amqp"
)

func (mq *rabbitMq) NewQueue() (err error) {
	//创建队列queue
	if _, err = mq.Channel.QueueDeclare(
		mq.Queue, //队列名称
		false,    // 是否持久化
		false,    //是否自动删除
		false,    //是否具有排他性
		false,    //是否阻塞
		nil,
	); err != nil {
		return err
	}

	return nil
}

func (mq *rabbitMq) SendMsg(msg string) (err error) {

	//确保消息发送成功
	if err = mq.Channel.Confirm(false); err != nil {
		return err
	}
	cinfirm := mq.Channel.NotifyPublish(make(chan amqp.Confirmation, 1))
	defer confirm(cinfirm)
	if err = mq.Channel.Publish(
		mq.ExChange,
		mq.Queue,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		},
	); err != nil {
		return err
	}
	return nil
}
func confirm(c <-chan amqp.Confirmation) {
	confirmed := <-c
	if confirmed.Ack {
		fmt.Println("ok")
	} else {
		fmt.Println("error")
	}

}
