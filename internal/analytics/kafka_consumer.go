package analytics

import (
	"context"
	"encoding/json"
	"log"
	"time"
	"url-shortner/internal/models"

	"github.com/segmentio/kafka-go"
)



func StartKafkaConsumer(broker, topic, groupID string) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker},
		GroupID: groupID,
		Topic:   topic,
	})

	log.Println("Kafka consumer started...")

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Println("Error reading message:", err)
			continue
		}

		var event models.ClickEvent
		if err := json.Unmarshal(m.Value, &event); err != nil {
			log.Println("Error parsing message:", err)
			continue
		}

		log.Printf("[Kafka] ClickEvent: shortCode=%s, originalURL=%s, timestamp=%s\n",
			event.ShortCode, event.OriginalURL, event.Timestamp.Format(time.RFC3339))

		// Optionally save to DB
		// repository.SaveClickEventToDB(event)  <-- you can implement this
	}
}
