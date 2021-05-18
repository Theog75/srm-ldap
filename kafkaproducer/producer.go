package kafkaproducer

import (
	"context"
	"time"

	kafka "github.com/segmentio/kafka-go"

	"fmt"
	"srm-ldap/config"
)

func ProducerHandler(ProduceMsg string) {
	// get kafka writer using environment variables.
	kafkaURL := config.KafkaURL
	topic := config.KafkaTopic
	writer := newKafkaWriter(kafkaURL, topic)
	fmt.Println("msg:\n", ProduceMsg)
	defer writer.Close()
	fmt.Println("start producing ... !!")
	msg := kafka.Message{
		Key:   []byte("collector"),
		Value: []byte(ProduceMsg),
	}
	fmt.Println(msg)
	writer.WriteMessages(context.Background(), msg)
	time.Sleep(1 * time.Second)
}

func newKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    config.KafkaTopic,
		Balancer: &kafka.LeastBytes{},
	}
}
