package main

import (
    "github.com/gin-gonic/gin"
    "url-shortner/internal/handlers"
)

func main() {
    router := gin.Default()
    handler.RegisterRoutes(router)
    router.Run(":8080")
}