package kafkaproducer

import (
	"context"
	"os"
	"time"

	kafka "github.com/segmentio/kafka-go"

	"fmt"
	"srm-ldap/config"
)

func ProducerHandler(ProduceMsg string) {
	// get kafka writer using environment variables.
	kafkaURL := config.KafkaURL
	topic := os.Getenv(config.KafkaTopic)
	writer := newKafkaWriter(kafkaURL, topic)
	defer writer.Close()
	fmt.Println("start producing ... !!")
	msg := kafka.Message{
		Key:   []byte("collector"),
		Value: []byte(ProduceMsg),
	}
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
