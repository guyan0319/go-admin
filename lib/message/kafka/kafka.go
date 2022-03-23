package kafka

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"os"
	"os/signal"
	"go-admin/lib/common"
)

var conf = common.Conf

type consumerGroupHandler struct {
	Handlefnc func(sarama.ConsumerMessage) error
}

func (consumerGroupHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}
func (consumerGroupHandler) Cleanup(sarama.ConsumerGroupSession) error { return nil }
func HandlerSet(handlefnc func(sarama.ConsumerMessage) error) *consumerGroupHandler {
	return &consumerGroupHandler{handlefnc}
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (h consumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		log.Printf("Message claimed: value = %s, timestamp = %v, topic = %s", msg.Value, msg.Timestamp, msg.Topic)
		if err := h.Handlefnc(*msg); err != nil {
			fmt.Errorf("handle error message %s: %v", msg.Value, err)
			continue
		}
		session.MarkMessage(msg, "")
	}
	return nil
}

//消费者组
func SaramaConsumerGroup(groupId string, topics []string, handle consumerGroupHandler) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = false
	config.Version = sarama.V0_10_2_0                     // specify appropriate version
	config.Consumer.Offsets.Initial = sarama.OffsetOldest // 未找到组消费位移的时候从哪边开始消费

	group, err := sarama.NewConsumerGroup(conf.Kafka, groupId, config)
	if err != nil {
		panic(err)
	}
	defer func() { _ = group.Close() }()

	// Track errors
	go func() {
		for err := range group.Errors() {
			fmt.Println("ERROR", err)
		}
	}()
	fmt.Println("Consumed start")
	// Iterate over consumer sessions.
	ctx := context.Background()
	for {
		handler := handle
		//handler := consumerGroupHandler{}

		// `Consume` should be called inside an infinite loop, when a
		// server-side rebalance happens, the consumer session will need to be
		// recreated to get the new claims
		err := group.Consume(ctx, topics, handler)
		if err != nil {
			panic(err)
		}
	}
}
func SaramaConsumer(topic string, handle consumerGroupHandler) {
	consumer, err := sarama.NewConsumer(conf.Kafka, sarama.NewConfig())
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	// Trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	consumed := 0
ConsumerLoop:
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			log.Printf("Consumed message offset %d\n", msg.Offset)
			err := handle.Handlefnc(*msg)
			if err != nil {
				log.Printf("Consumed message error offset is  %d %v\n", msg.Offset, msg.Value)
			}
			consumed++
		case <-signals:
			break ConsumerLoop
		}
	}

	log.Printf("Consumed: %d\n", consumed)
}

//异步生产者Goroutines
func SyncProducer(msg *sarama.ProducerMessage) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	producer, err := sarama.NewAsyncProducer(conf.Kafka, config)
	if err != nil {
		panic(err)
	}
	defer producer.Close()
	//msg := &sarama.ProducerMessage{Topic: "my_topic", Value: sarama.StringEncoder("testing 456")}

	// send to chain
	producer.Input() <- msg
	log.Printf("> message sent to topic %d at offset %d\n", msg.Topic, msg.Offset)

}

//异步生产者Select
func SyncProducerSelect(msg *sarama.ProducerMessage) {
	producer, err := sarama.NewAsyncProducer(conf.Kafka, nil)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	//msg := &sarama.ProducerMessage{Topic: "my_topic", Value: sarama.StringEncoder("testing 456")}

	// send to chain
	producer.Input() <- msg
	log.Printf("> message sent to topic %d at offset %d\n", msg.Topic, msg.Offset)

}

//同步生产者
func SaramaProducer(msg *sarama.ProducerMessage) {
	producer, err := sarama.NewSyncProducer(conf.Kafka, nil)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	//msg := &sarama.ProducerMessage{Topic: "my_topic", Value: sarama.StringEncoder("testing 123")}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Printf("FAILED to send message: %s\n", err)
	} else {
		log.Printf("> message sent to partition %d at offset %d\n", partition, offset)
	}

}
