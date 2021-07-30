package change

import (
	"context"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var logger = log.New(os.Stdout, "[rabbitmq]:", log.LstdFlags)

type rabbitMq struct {
	Conn     *amqp.Connection
	Channel  *amqp.Channel
	ExChange string
	Queue    string
	key      string
	Retry    int //重试次数
}

func NewRabbitMq(exChange, queue, key string) *rabbitMq {
	return &rabbitMq{
		ExChange: exChange,
		Queue:    queue,
		key:      key,
	}
}

func (mq *rabbitMq) NewChannel(url string) (channel *amqp.Channel, err error) {
	mq.Conn, err = amqp.Dial(url)
	if err != nil {
		log.Println("dial rabbitmq error", err)
		return nil, err
	}
	mq.Channel, err = mq.Conn.Channel()
	return mq.Channel, err
}

func (mq *rabbitMq) NewExchange(kind string) error {
	return mq.Channel.ExchangeDeclare(
		mq.ExChange,
		kind,  //topic :consume模糊匹配绑定exchange，direct：consume精确匹配绑定，"fanout"：所有的consume都能绑定
		false, // 是否持久化
		false, //是否自动删除
		false, //是否具有排他性
		false, //是否阻塞
		nil,
	)
}

func (mq *rabbitMq) SendMsg(msg string) error {

	//确保消息发送成功
	if err := mq.Channel.Confirm(false); err != nil {
		return err
	}
	cinfirm := mq.Channel.NotifyPublish(make(chan amqp.Confirmation, 10000))
	defer confirm(cinfirm)
	return mq.Channel.Publish(
		mq.ExChange,
		mq.key,
		false, //如果为true，根据自身exchange类型和routekey规则无法找到符合条件的队列会把消息返还给发送者
		false, //如果为true，当exchange发送消息到队列后发现队列上没有消费者，则会把消息返还给发送者
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		},
	)
}

func confirm(c <-chan amqp.Confirmation) {
	confirmed := <-c
	if confirmed.Ack {

		fmt.Println("ok")
	} else {
		fmt.Println("error")
	}

}

func (mq *rabbitMq) Consume(kind string, f func(delivery amqp.Delivery)) error {

	err := mq.Channel.ExchangeDeclare(mq.ExChange,
		kind,  //topic :consume模糊匹配绑定exchange，direct：consume精确匹配绑定，"fanout"：所有的consume都能绑定
		false, // 是否持久化
		false, //是否自动删除
		false, //是否具有排他性
		false, //是否阻塞
		nil)
	if err != nil {
		return err
	}
	//创建队列，不指定队列名称(会随机创建队列名)
	queue, err := mq.Channel.QueueDeclare("", true, false, false, false, nil)
	if err != nil {
		return err
	}
	err = mq.Channel.QueueBind(queue.Name, mq.key, mq.ExChange, false, nil)
	if err != nil {
		return err
	}
	msg, err := mq.Channel.Consume(queue.Name, "", false, false, false, false, nil)
	if err != nil {
		return err
	}
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	for m := range msg {
		f(m)
	}
	c := context.Background()
	ctx, cancel := context.WithCancel(c)
	go func(ctx context.Context) {
		for m := range msg {
			f(m)
		}
	}(ctx)
	<-quit
	cancel()
	return nil
}
