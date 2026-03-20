# Stage 1 — Builder
FROM golang:1.22 as builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/main.go

# Stage 2 — Runtime
FROM alpine:3.19
WORKDIR /root/
COPY --from=builder /app/app .
EXPOSE 8080
CMD ["./app"]
