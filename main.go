package main

import (
	"fmt"
	"log"
	handler "url-shortner/internal/handlers"
	"url-shortner/internal/repository"
    "url-shortner/internal/analytics"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := repository.InitPostgres()
	if err != nil {
		log.Fatal("PostgreSQL error:", err)
	}
    repository.SetPostgres(db)
	defer db.Close()
	fmt.Println("PostgreSQL connected!")

	rdb := repository.InitRedis()
    repository.SetRedisClient(rdb)
	defer rdb.Close()
	fmt.Println("Redis connected!")
    analytics.InitKafkaProducer("localhost:9092", "click-events")
    go analytics.StartKafkaConsumer("localhost:9092", "click-events", "url-shortener-consumer-group")
    router := gin.Default()
	handler.RegisterRoutes(router)
	router.Run(":8080")
}
