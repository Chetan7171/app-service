# GoLang CRUD Microservice with PostgreSQL (Kubernetes Ready)

This repository contains a production-ready GoLang microservice that performs full
CRUD (Create, Read, Update, Delete) operations on a PostgreSQL database.

The service is designed for:
- Docker containerization
- Deployment using Helm
- GitOps automation using ArgoCD
- Running inside AWS EKS
- Using PostgreSQL deployed **inside the Kubernetes cluster**

---

## 🧱 Tech Stack

- Go 1.22
- PostgreSQL (StatefulSet in EKS)
- Chi Router
- PGX PostgreSQL driver
- Docker
- Helm
- Kubernetes
- ArgoCD
- GitHub Actions (CI/CD)

---

## 🚀 API Endpoints

| Method | Endpoint       | Description          |
|--------|----------------|----------------------|
| POST   | /users         | Create new user      |
| GET    | /users         | List all users       |
| PUT    | /users/{id}    | Update user          |
| DELETE | /users/{id}    | Delete user          |
| GET    | /health        | Health check         |

---

## 🐳 Run Locally (Docker)

```bash
docker build -t app-service .
docker run -p 8080:8080 app-service
