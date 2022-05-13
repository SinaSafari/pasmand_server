package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"strings"
)

const (
	OTP = "OTP"
)

func main() {
	//fmt.Println("consumer is starting")
	reader := getKafkaReader("kafka:29092", OTP, "west-side")
	//defer log.Fatalln(reader.Close())

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		payload := strings.Split(string(msg.Value), ",")

		fmt.Println("sending " + payload[1] + " to phone number: " + payload[0] + " ...")
		//fmt.Println("otp: "+string(msg.Value), string(msg.Key))
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
