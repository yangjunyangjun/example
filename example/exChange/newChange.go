package exChange

import "github.com/streadway/amqp"

type rabbitMq struct {
	Conn     *amqp.Connection
	Channel  *amqp.Channel
	ExChange string
	Queue    string
	key      string
	Retry    int //重试次数
}

func (mq *rabbitMq) NewRabbitMq(exChange, queue, key string) *rabbitMq {
	return &rabbitMq{
		ExChange: exChange,
		Queue:    queue,
		key:      key,
	}
}

func (mq *rabbitMq) NewChannel(url string) (err error) {
	mq.Conn, err = amqp.Dial(url)
	if err != nil {
		return err
	}
	mq.Channel, err = mq.Conn.Channel()
	return err
}

func (mq *rabbitMq) NewExchange() {
	mq.Channel.ExchangeDeclare(
		mq.ExChange,
		"", //topic :consume模糊匹配绑定exchange，direct：consume精确匹配绑定，""：所有的consume都能版本
		false,
		false,
		false,
		false,
		nil,
	)
}
