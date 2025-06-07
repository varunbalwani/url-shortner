package handler

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "url-shortener/internal/service"
)

type ShortenRequest struct {
    URL string `json:"url"`
}

func ShortenURL(c *gin.Context) {
    var req ShortenRequest
    if err := c.ShouldBindJSON(&req); err != nil || req.URL == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    shortCode, err := service.Shorten(req.URL)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"short_url": "http://localhost:8080/" + shortCode})
}

func ResolveURL(c *gin.Context) {
    shortCode := c.Param("shortCode")
    originalURL, err := service.Resolve(shortCode)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
        return
    }

    c.Redirect(http.StatusFound, originalURL)
}