package mq

import (
	"admin-server/config"
	"github.com/Shopify/sarama"
	"log"
)

var KafkaClient *sarama.Client

func init() {
	if KafkaClient != nil {
		return
	}
	KafkaClient = initKafkaClient()
}

func initKafkaClient() *sarama.Client {
	saramaConfig := sarama.NewConfig()
	saramaConfig.Producer.Return.Successes = true
	saramaConfig.Producer.RequiredAcks = sarama.WaitForAll
	saramaConfig.Producer.Partitioner = sarama.NewRandomPartitioner

	client, err := sarama.NewClient([]string{config.KafkaHosts}, saramaConfig)
	if err != nil {
		log.Fatalf("unable to create kafka client: %q", err)
	}
	return &client
}

func Publist(msg []byte) bool {
	isSuc := true

	producer, err := sarama.NewAsyncProducerFromClient(*KafkaClient)
	if err != nil {
		log.Fatal(err)
		isSuc = false
		return isSuc
	}
	defer producer.Close()

	producer.Input() <- &sarama.ProducerMessage{Topic: config.KafkaTopic, Key: nil, Value: sarama.StringEncoder(msg)}

	return isSuc
}