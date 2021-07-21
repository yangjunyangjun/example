package simple

import (
	"github.com/streadway/amqp"
)

type rabbitMq struct {
	Conn     *amqp.Connection //mq连接
	Channel  *amqp.Channel    //mq Channel
	ExChange string           //exchange名称
	Queue    string           //queue名称
	Key      string           //bind key信息
	MqUrl    string           //mq url
	Retry    int              //重试次数
}

func NewRabbitMq(exChange, queue, mqUrl string) *rabbitMq {
	return &rabbitMq{
		ExChange: exChange,
		Queue:    queue,
		MqUrl:    mqUrl,
	}
}

func (mq *rabbitMq) Connection() (err error) {
	mq.Conn, err = amqp.Dial(mq.MqUrl)
	if err != nil {
		return err
	}
	mq.Channel, err = mq.Conn.Channel()
	if err != nil {
		return err
	}
	return nil
}

func (mq *rabbitMq) Close() error {
	return mq.Conn.Close()
}
