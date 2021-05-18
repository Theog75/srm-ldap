package kafkaproducer

import (
	"context"
	"encoding/json"
	"time"

	"fmt"
	"srm-ldap/config"
	"strings"

	kafka "github.com/segmentio/kafka-go"
)

type MessageFormat struct {
	Timestamp time.Time
	Data      string
}

func ProducerHandler(ProduceMsg string) {
	// get kafka writer using environment variables.
	kafkaURL := strings.Split(config.KafkaURL, ",")
	topic := config.KafkaTopic
	writer := newKafkaWriter(kafkaURL, topic)
	defer writer.Close()
	fmt.Println("start producing ... !!")
	baseMessage := MessageFormat{
		Timestamp: time.Now(),
		Data:      ProduceMsg,
	}

	jsonMessage, err := json.Marshal(baseMessage)

	if err != nil {
		fmt.Println(err)
	}

	msg := kafka.Message{
		Key:   []byte("collector"),
		Value: []byte(jsonMessage),
	}
	writer.WriteMessages(context.Background(), msg)
	time.Sleep(1 * time.Second)
}

func newKafkaWriter(kafkaURL []string, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL...),
		Topic:    config.KafkaTopic,
		Balancer: &kafka.LeastBytes{},
	}
}
