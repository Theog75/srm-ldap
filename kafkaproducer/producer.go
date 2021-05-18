package kafkaproducer

import (
	"context"
	"time"

	"fmt"
	"srm-ldap/config"
	"strings"

	kafka "github.com/segmentio/kafka-go"
)

func ProducerHandler(ProduceMsg string) {
	// get kafka writer using environment variables.
	kafkaURL := strings.Split(config.KafkaURL, ",")
	topic := config.KafkaTopic
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

func newKafkaWriter(kafkaURL []string, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    config.KafkaTopic,
		Balancer: &kafka.LeastBytes{},
	}
}
