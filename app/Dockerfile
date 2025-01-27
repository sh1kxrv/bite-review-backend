FROM golang:1.23.2-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o main ./main.go

FROM debian:bullseye-slim
ENV APP_ENV=production

WORKDIR /app

COPY --from=builder /app/main .

COPY config.json .

EXPOSE 3000

# Команда запуска приложения
CMD ["./main"]