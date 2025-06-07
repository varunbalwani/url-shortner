package handler

import (
    "github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
    r.POST("/shorten", ShortenURL)
    r.GET("/:shortCode", ResolveURL)
}
