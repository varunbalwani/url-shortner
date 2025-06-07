package handler

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func RegisterRoutes(router *gin.Engine) {
    router.GET("/ping", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "pong"})
    })

    router.POST("/shorten", ShortenURL)
    router.GET("/:shortCode", ResolveURL)
}

func ShortenURL(c *gin.Context) {
    // TODO: Call service to shorten URL
    c.JSON(http.StatusOK, gin.H{"short_url": "http://localhost:8080/abc123"})
}

func ResolveURL(c *gin.Context) {
    // TODO: Call service to resolve shortCode
    c.JSON(http.StatusOK, gin.H{"original_url": "https://example.com"})
}