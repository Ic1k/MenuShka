# syntax=docker/dockerfile:1

# Мультистадия сборки: стадия сборки
FROM golang:1.23.2-alpine AS builder

# Установка рабочей директории
WORKDIR /app

# Копируем go.mod и go.sum отдельно для кэширования зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем оставшиеся файлы проекта
COPY . .

# Сборка бинарного файла
RUN go build -o api-gateway main.go

# Минимальный финальный образ
FROM alpine:3.18

# Создаем непривилегированного пользователя для безопасности
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Установка рабочей директории
WORKDIR /app

# Копируем бинарный файл из стадии сборки
COPY --from=builder /app/api-gateway /app/api-gateway

# Меняем права для пользователя
USER appuser

# Экспорт порта
EXPOSE 8081

# Команда запуска
CMD ["./api-gateway"]
