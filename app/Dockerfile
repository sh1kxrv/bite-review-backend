FROM golang:1.23.2-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o main ./main.go

# Базовый образ для запуска
FROM debian:bullseye-slim

# Устанавливаем переменные окружения
ENV APP_ENV=production

# Рабочая директория внутри контейнера
WORKDIR /app

# Копируем бинарный файл из стадии сборки
COPY --from=builder /app/main .

# Копируем конфигурационный файл
COPY config.json .

# Открываем порт
EXPOSE 3000

# Команда запуска приложения
CMD ["./main"]