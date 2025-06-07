package main

import (
    "github.com/gin-gonic/gin"
    "url-shortener/internal/handler"
)

func main() {
    router := gin.Default()
    handler.RegisterRoutes(router)
    router.Run(":8080")
}