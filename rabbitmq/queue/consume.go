package queue

import (
	"fmt"
	"github.com/streadway/amqp"
)

func (mq *rabbitMq) Consume(f func(delivery amqp.Delivery)) error {
	msg, err := mq.Channel.Consume(
		mq.Queue,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	go func() {

		for m := range msg {
			f(m)
		}

	}()
	esixt := make(chan bool)
	<-esixt
	return nil
}

func PrintMq(delivery amqp.Delivery) {

	fmt.Println(string(delivery.Body))
	delivery.Ack(true)
}
