package analytics

import (
	"context"
	"encoding/json"
	"log"
	"time"
	"url-shortner/internal/models"

	"github.com/segmentio/kafka-go"
)

var writer *kafka.Writer

func InitKafkaProducer(brokerAddress string, topic string) {
	writer = kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
	})
}

func PublishClick(shortCode string, originalURL string) {
	event := models.ClickEvent{
		ShortCode:   shortCode,
		OriginalURL: originalURL,
		Timestamp:   time.Now().UTC(),
	}

	bytes, err := json.Marshal(event)
	if err != nil {
		log.Println("Error marshalling click event:", err)
		return
	}

	err = writer.WriteMessages(context.Background(), kafka.Message{
		Value: bytes,
	})
	if err != nil {
		log.Println("Error writing click event to Kafka:", err)
	}
}
