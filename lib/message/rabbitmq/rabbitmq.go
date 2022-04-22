package rabbitmq

import (
	"bytes"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"time"
	"go-admin/lib/common"
)

// 定义RabbitMQ对象
type RabbitMQ struct {
	conn      *amqp.Connection
	ch        *amqp.Channel
	QueueName string // 队列名称
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func (r *RabbitMQ) Connect() {
	conf := common.Conf
	var err error
	r.conn, err = amqp.Dial(conf.Rabbit[0])
	failOnError(err, "Failed to connect to RabbitMQ")
	r.ch, err = r.conn.Channel()
	failOnError(err, "Failed to connect to RabbitMQ")
}
func (r *RabbitMQ) Close() {
	// 先关闭管道,再关闭链接
	err := r.ch.Close()
	if err != nil {
		fmt.Printf("Failed to close to RabbitMQ channel:%s \n", err)
	}
	err = r.conn.Close()
	if err != nil {
		fmt.Printf("Failed to close to RabbitMQ:%s \n", err)
	}
}

func (r *RabbitMQ) Producer(body string) {
	defer r.Close() // 处理结束关闭链接
	//连接失效重连
	if r.ch == nil {
		r.Connect()
	}
	q, err := r.ch.QueueDeclare(
		r.QueueName, // name
		true,        // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	if err != nil {
		log.Println("Failed to declare a queue")
		return
	}
	err = r.ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		})
	if err != nil {
		log.Printf(" Failed to publish a message [x] Sent %s", body)
		return
	}

}

func (r *RabbitMQ) Consumer() {
	//连接失效重连
	if r.ch == nil {
		r.Connect()
	}
	q, err := r.ch.QueueDeclare(
		r.QueueName, // name
		true,        // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = r.ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	msgs, err := r.ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			dot_count := bytes.Count(d.Body, []byte("."))
			t := time.Duration(dot_count)
			time.Sleep(t * time.Second)
			log.Printf("Done")
			d.Ack(false)
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
