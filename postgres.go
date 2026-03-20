package database

import (
    "context"
    "log"
    "time"

    "github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDB(dbURL string) *pgxpool.Pool {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    dbPool, err := pgxpool.New(ctx, dbURL)
    if err != nil {
        log.Fatalf("Unable to create DB pool: %v", err)
    }

    err = dbPool.Ping(ctx)
    if err != nil {
        log.Fatalf("Unable to connect to database: %v", err)
    }

    log.Println("Connected to PostgreSQL successfully")
    return dbPool
}
