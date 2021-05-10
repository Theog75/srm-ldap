package kafkaproducer

import (
	"context"
	"encoding/json"
	"os"
	"time"

	kafka "github.com/segmentio/kafka-go"

	"fmt"
	"srm-dhcp/config"
)

func ProducerHandler(ProduceMsg *config.Dhcp) {
	// get kafka writer using environment variables.
	kafkaURL := os.Getenv("kafkaURL")
	topic := os.Getenv(config.KafkaURL)
	writer := newKafkaWriter(kafkaURL, topic)
	defer writer.Close()
	fmt.Println("start producing ... !!")
	b, err := json.Marshal(ProduceMsg)
	msg := kafka.Message{
		Key:   []byte("collector"),
		Value: []byte(b),
	}
	writer.WriteMessages(context.Background(), msg)

	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(1 * time.Second)
}

func newKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(config.KafkaURL),
		Topic:    config.KafkaTopic,
		Balancer: &kafka.LeastBytes{},
	}
}
