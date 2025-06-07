package service

import (
    "errors"
    "math/rand"
    "time"
    "url-shortner/internal/repository"
)

const shortCodeLength = 6
var charset = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func init() {
    rand.Seed(time.Now().UnixNano())
}

func generateShortCode() string {
    b := make([]rune, shortCodeLength)
    for i := range b {
        b[i] = charset[rand.Intn(len(charset))]
    }
    return string(b)
}

func Shorten(originalURL string) (string, error) {
    shortCode := generateShortCode()

    // Ensure unique (in practice retry a few times)
    _, err := repository.GetFromDB(shortCode)
    if err == nil {
        return "", errors.New("collision occurred, try again")
    }

    err = repository.SaveToDB(shortCode, originalURL)
    if err != nil {
        return "", err
    }

    // Save to Redis for faster future lookup
    _ = repository.SaveToCache(shortCode, originalURL)

    return shortCode, nil
}

func Resolve(shortCode string) (string, error) {
    // Try Redis
    url, err := repository.GetFromCache(shortCode)
    if err == nil && url != "" {
        return url, nil
    }

    // Fallback to DB
    url, err = repository.GetFromDB(shortCode)
    if err != nil {
        return "", err
    }

    // Cache it again
    _ = repository.SaveToCache(shortCode, url)

    return url, nil
}