package main

import (
    "log"
    "net/http"

    "app-service/internal/config"
    "app-service/internal/database"
    "app-service/internal/handlers"
    "app-service/internal/repository"
    "app-service/internal/router"
)

func main() {
    cfg := config.LoadConfig()
    db := database.ConnectDB(cfg.DBUrl)

    userRepo := repository.NewUserRepo(db)
    userHandler := &handlers.UserHandler{Repo: userRepo}

    r := router.SetupRoutes(userHandler)

    log.Println("Starting server on port", cfg.Port)
    log.Fatal(http.ListenAndServe(":"+cfg.Port, r))
}
