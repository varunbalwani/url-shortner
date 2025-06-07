package main

import (
	"fmt"
	"log"
	handler "url-shortner/internal/handlers"
	"url-shortner/internal/repository"

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
    router := gin.Default()
	handler.RegisterRoutes(router)
	router.Run(":8080")
}
