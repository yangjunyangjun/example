# rabbitmq


### 1、生产者
  ```go
    conn,err:=amqp.Dial(url) //创建mq连接
    c,err:=conn.Channel()  //获取channel
    q,err:=c.QueueDeclare() //创建队列,如果队列存在则跳过，如果不存在则创建
    conn.Publish()
``` 

### 消费者
```go
     conn,err:=amqp.Dial(url) //创建mq连接
     c,err:=conn.Channel()  //获取channel
    msg,err:=c.Coumuse() //获取消费channel  
```