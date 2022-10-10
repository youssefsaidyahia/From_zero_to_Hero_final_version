package kafka_installation

import (
	"context"
	"encoding/json"
	"fmt"
	"fristTry/config"
	"fristTry/internal/Models"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

func Produce(transaction *Models.Tranactions) {
	configurations := config.LoadConfig()
	conn, _ := kafka.DialLeader(context.Background(), "tcp", configurations.Kafkastream.Kafka_brokers, configurations.Kafkastream.Kafka_topic, 0)
	conn.SetWriteDeadline(time.Now().Add(time.Second * 10))
	obj, _ := json.Marshal(&transaction)
	conn.WriteMessages(kafka.Message{Value: []byte(obj)})
}
func Consume() {
	configurations := config.LoadConfig()
	config := kafka.ReaderConfig{
		Brokers:  []string{configurations.Kafkastream.Kafka_brokers},
		Topic:    configurations.Kafkastream.Kafka_topic,
		MaxBytes: configurations.Kafkastream.MaxBytes,
	}

	reader := kafka.NewReader(config)
	var transaction Models.Tranactions
	for {
		message, error := reader.ReadMessage(context.Background())
		if error != nil {
			log.Fatalf(time.Now().String()+":: Error happened during calling kafka server %v", error)
			continue
		}
		fmt.Println(time.Now().String() + "::message of transaction consumed:: " + string(message.Value))
		json.Unmarshal(message.Value, &transaction)
		UpdateTransaction(&transaction)
	}
}
