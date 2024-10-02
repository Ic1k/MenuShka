# Используем официальный образ Go
FROM golang:1.20-alpine

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файл go.mod и go.sum
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем все остальные файлы
COPY . ./

# Собираем бинарник
RUN go build -o api-gateway main.go

# Указываем порт, который будет использоваться
EXPOSE 8080

# Команда для запуска
CMD ["./api-gateway"]
