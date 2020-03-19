package mq

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
)

var done chan bool

// 开始监听队列
// CName: 消费者名称
func StartConsume(CName string, callback func(msg []byte)) {
	consumer, err := sarama.NewConsumerFromClient(*KafkaClient)
	if err != nil {
		log.Fatal(err)
	}

	partitionConsumer, err := consumer.ConsumePartition(CName, 0, sarama.OffsetNewest)
	if err != nil {
		log.Println("partitionConsumerError: ")
		log.Println(err)
	}

	done = make(chan bool)
	go func() {
		for {
			msg := <-partitionConsumer.Messages()
			if err == nil {
				callback(msg.Value)
				fmt.Printf("Message on %s: %s\n", msg.Topic, string(msg.Value))
			} else {
				// The client will automatically try to recover from all errors.
				fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			}
		}
	}()

	// done没有新消息，则一直阻塞
	<-done
	defer consumer.Close()
	defer partitionConsumer.Close()
}
