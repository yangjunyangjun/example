package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"strings"
)

type Kafka struct {
	Addr  string
	Topic string
}

func (k Kafka) Producer(body string) error {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	//设置使用的kafka版本,如果低于V0_10_0_0版本,消息中的timestrap没有作用.需要消费和生产同时配置
	//注意，版本设置不对的话，kafka会返回很奇怪的错误，并且无法成功发送消息
	config.Version = sarama.V0_10_0_1
	msg := &sarama.ProducerMessage{
		Topic: k.Topic,
		Value: sarama.StringEncoder(body),
	}

	// 连接kafka
	producer, err := sarama.NewSyncProducer(strings.Split(k.Addr, ","), config)

	if err != nil {
		return err
	}
	defer producer.Close()
	pid, offset, err := producer.SendMessage(msg)

	if err != nil {
		return err
	}
	fmt.Println(pid)
	fmt.Println(offset)
	return nil
}

func (k Kafka) Consume() error {
	consume, err := sarama.NewConsumer(strings.Split(k.Addr, ","), nil)
	if err != nil {
		return err
	}
	parList, err := consume.Partitions(k.Topic)
	if err != nil {
		return err
	}
	for p := range parList {
		pc, err := consume.ConsumePartition(k.Topic, int32(p), sarama.OffsetNewest)
		if err != nil {
			log.Print(err)
			return err
		}
		for m := range pc.Messages() {
			fmt.Println(string(m.Value))
		}
	}
	return nil
}

func (k Kafka) SyncProducer(data []string) {

	//配置producer config
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Version = sarama.V0_10_0_1

	producer, err := sarama.NewAsyncProducer(strings.Split(k.Addr, ","), config)
	if err != nil {
		fmt.Println(err)
	}
	defer producer.AsyncClose()
	go func(asyncProducer sarama.AsyncProducer) {
		for {
			select {
			case e := <-asyncProducer.Errors():
				fmt.Println("send error: ", e)
			case <-asyncProducer.Successes():
				fmt.Println("success")
			}
		}

	}(producer)
	msg := sarama.ProducerMessage{
		Topic: k.Topic,
	}

	for _, d := range data {
		msg.Value = sarama.StringEncoder(d)
		producer.Input() <- &msg
	}
}
