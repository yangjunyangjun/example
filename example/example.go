package main

//import (
//	"fmt"
//	"github.com/streadway/amqp"
//	"time"
//)

////因：快速实现逻辑，故：不处理错误逻辑
//func main() {
//	conn, _ := amqp.Dial("amqp://user:password@host:ip/vhost")
//	ch, _ := conn.Channel()
//	body := "Hello World!! " + time.Now().Format("2006-01-02 15:04:05")
//	fmt.Println(body)
//	var exchange_name string = "j_exch_head"
//	var routing_key string = "jkey"
//	var etype string = amqp.ExchangeHeaders
//
//	a := ch.Confirm(false)
//	//声明交换器
//	ch.ExchangeDeclare(exchange_name, etype, true, false, false, true, nil)
//
//	confirms := ch.NotifyPublish(make(chan amqp.Confirmation, 1)) // 处理确认逻辑
//	defer confirmOne(confirms) // 处理方法
//	ch.Publish(
//		exchange_name, // exchange 这里为空则不选择 exchange
//		routing_key,   // routing key
//		false,         // mandatory
//		false,         // immediate
//		amqp.Publishing{
//			ContentType: "text/plain",
//			Body:        []byte(body),
//			Headers:     amqp.Table{"x-match": "any", "mail": "470047253@qq.com", "author": "Jhonny"}, // 头部信息 any:匹配一个即可 all:全部匹配
//			//Expiration:  "3000", // 设置过期时间
//		},
//	)
//
//	fmt.Println(a)
//	// defer 关键字
//	defer conn.Close() // 压栈 后进先出
//	defer ch.Close()   // 压栈 后进先出
//}
//// 消息确认
//func confirmOne(confirms <-chan amqp.Confirmation) {
//	if confirmed := <-confirms; confirmed.Ack {
//		fmt.Printf("confirmed delivery with delivery tag: %d", confirmed.DeliveryTag)
//	} else {
//		fmt.Printf("confirmed delivery of delivery tag: %d", confirmed.DeliveryTag)
//	}
//}