package database

import (
        "fmt"
        "log"
        "time"

        "github.com/jmoiron/sqlx"
        _ "github.com/lib/pq"
)

func NewPostgresDB(connString string) (*sqlx.DB, error) {
    log.Printf("Connecting to database with: %s", connString) // Log the connection string (for debugging)

        db, err := sqlx.Connect("postgres", connString)
        if err != nil {
                return nil, fmt.Errorf("failed to connect to database: %w", err)
        }

        // Configure connection pool
        db.SetMaxOpenConns(25)
        db.SetMaxIdleConns(5)
        db.SetConnMaxLifetime(5 * time.Minute)

        // Ping the database to verify the connection
        err = db.Ping()
        if err != nil {
                return nil, fmt.Errorf("database ping failed: %w", err)
        }

        log.Println("✅ Successfully connected to the database")
        return db, nil
}