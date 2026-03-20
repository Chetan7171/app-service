package config

import (
    "os"
)

type Config struct {
    Port     string
    DBUrl    string
}

func LoadConfig() *Config {
    return &Config{
        Port:  getEnv("PORT", "8080"),
        DBUrl: getEnv("DATABASE_URL", "postgres://postgres:postgres@postgres:5432/users?sslmode=disable"),
    }
}

func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}
