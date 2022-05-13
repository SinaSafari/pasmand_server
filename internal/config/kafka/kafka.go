package kafka

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

const (
	OTP = "OTP"
)

func newKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:         kafka.TCP(kafkaURL),
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		BatchTimeout: time.Millisecond * 500,
	}
}

func WriteToKafka(payload string) {
	writer := newKafkaWriter("kafka:29092", OTP)

	//defer log.Fatalln(writer.Close())
	err := writer.WriteMessages(context.Background(), kafka.Message{Value: []byte(payload)})
	if err != nil {
		fmt.Println(err)
	}
}

func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{kafkaURL},
		GroupID:  groupID,
		Topic:    topic,
		MaxBytes: 10e6, // 10MB

	})
}

func Consume() {
	reader := getKafkaReader("kafka:29092", OTP, "west-side")
	//defer log.Fatalln(reader.Close())

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(msg.Value), string(msg.Key))
	}
}
