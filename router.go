package router

import (
    "net/http"

    "app-service/internal/handlers"

    "github.com/go-chi/chi/v5"
)

func SetupRoutes(userHandler *handlers.UserHandler) http.Handler {
    r := chi.NewRouter()

    r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("OK"))
    })

    r.Post("/users", userHandler.CreateUser)
    r.Get("/users", userHandler.GetUsers)
    r.Put("/users/{id}", userHandler.UpdateUser)
    r.Delete("/users/{id}", userHandler.DeleteUser)

    return r
}
