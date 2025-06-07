package repository

import (
    "database/sql"
    _ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() error {
    var err error
    db, err = sql.Open("postgres", "postgresql://postgres:password@localhost:5432/url_shortener?sslmode=disable")
    if err != nil {
        return err
    }
    return db.Ping()
}

func SaveToDB(shortCode string, originalURL string) error {
    _, err := db.Exec("INSERT INTO urls (short_code, original_url) VALUES ($1, $2)", shortCode, originalURL)
    return err
}

func GetFromDB(shortCode string) (string, error) {
    var url string
    err := db.QueryRow("SELECT original_url FROM urls WHERE short_code=$1", shortCode).Scan(&url)
    return url, err
}