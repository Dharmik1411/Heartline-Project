package config

import (
    "database/sql"
    "log"
    _ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
    var err error
    connStr := "postgresql://postgres:password@localhost:5432/simple_auth?sslmode=disable" // update your password/db name if needed
    DB, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal("Database connection error: ", err)
    }

    if err = DB.Ping(); err != nil {
        log.Fatal("Database ping error: ", err)
    }

    log.Println("Connected to the database!")
}
