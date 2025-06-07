package repository

import (
    "database/sql"
    _ "github.com/lib/pq"
    "errors"
)

var db *sql.DB

func SetPostgres(database *sql.DB) {
    db = database
}

func InitPostgres() (*sql.DB, error) {
	connStr := "host=localhost port=5432 user=user password=pass dbname=url_shortener sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, db.Ping()
}

func SaveToDB(shortCode string, originalURL string) error {
    _, err := db.Exec("INSERT INTO urls (short_code, original_url) VALUES ($1, $2)", shortCode, originalURL)
    if db == nil {
        return errors.New("PostgreSQL connection not initialized")
    }
    return err
}

func GetFromDB(shortCode string) (string, error) {
    var url string
    err := db.QueryRow("SELECT original_url FROM urls WHERE short_code=$1", shortCode).Scan(&url)
    if db == nil {
        return "", errors.New("PostgreSQL connection not initialized")
    }
    return url, err
}