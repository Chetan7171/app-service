# GoLang CRUD Microservice with PostgreSQL (Kubernetes Ready)

This repository contains a production‑ready GoLang microservice that provides full  
CRUD (Create, Read, Update, Delete) operations on a PostgreSQL database.

The application is designed following modern DevOps and cloud‑native patterns.  
It will be deployed using:

- Docker  
- Kubernetes (AWS EKS)  
- Helm  
- ArgoCD (GitOps)  
- PostgreSQL running inside the Kubernetes cluster  
- GitHub Actions CI/CD pipeline  

This service is part of a **multi‑repository DevOps project**.

---

## 🧱 Tech Stack

- **Language:** Go 1.22  
- **Framework:** chi router  
- **Database:** PostgreSQL on Kubernetes (StatefulSet)  
- **Driver:** pgx PostgreSQL driver  
- **Container:** Docker  
- **Orchestration:** Kubernetes  
- **Deployment:** Helm + ArgoCD  
- **CI/CD:** GitHub Actions  
- **Infrastructure:** Terraform (in separate repo)

---

## 🚀 Features

- Create a user  
- List all users  
- Update a user  
- Delete a user  
- Health check endpoint  
- PostgreSQL DB layer using repository pattern  
- Structured project layout  
- Production‑ready Dockerfile  

---

## 📡 API Endpoints

| Method | Endpoint       | Description          |
|--------|----------------|----------------------|
| POST   | `/users`       | Create new user      |
| GET    | `/users`       | Get all users        |
| PUT    | `/users/{id}`  | Update a user        |
| DELETE | `/users/{id}`  | Delete a user        |
| GET    | `/health`      | Health check         |

---

## 🔧 Environment Variables

| Name | Description |
|------|-------------|
| `PORT` | Application port (default 8080) |
| `DATABASE_URL` | PostgreSQL connection string |

Example connection string (Kubernetes):
postgres://postgres:postgres@postgres:5432/users?sslmode=disable


---

## 🗄 Database Migration

This folder contains SQL files used to initialize the PostgreSQL schema.

File:migrations/001_create_users.sql

📁 Folder Structure
app-service/
 ├── cmd/
 │    └── main.go
 ├── internal/
 │    ├── config/config.go
 │    ├── database/postgres.go
 │    ├── handlers/user_handler.go
 │    ├── models/user.go
 │    ├── repository/user_repo.go
 │    └── router/router.go
 ├── migrations/001_create_users.sql
 ├── Dockerfile
 ├── go.mod
 ├── go.sum
 └── README.md


 🐳 Run Locally with Docker
Build the image:
docker build -t app-service .

Run the container:
docker run -p 8080:8080 app-service


🛠 Local Development (Without Docker)
Run directly:
go run cmd/main.go

📦 Production Build
For Linux build:
CGO_ENABLED=0 GOOS=linux go build -o app cmd/main.go

☁ Deployment (EKS + ArgoCD + Helm)
This service is deployed via the gitops-deploy repo using:

Helm charts
ArgoCD Application manifest
GitOps workflow watching changes

Infrastructure (EKS, VPC, PostgreSQL PVC, etc.) is provisioned from the gitops-infra repo.

👨‍💻 Author
Chetan Satish Salaskar
