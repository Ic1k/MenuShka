package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
)

type AuthResponse struct {
	Message string `json:"message"`
}

// Простая заглушка для валидного токена
const validToken = "valid_token_123"

func authenticate(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")

	if token == validToken {
		json.NewEncoder(w).Encode(AuthResponse{Message: "Authenticated"})
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(AuthResponse{Message: "Unauthorized"})
	}
}

func main() {
	// Настройка логирования
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})

	// Получение порта из переменных окружения
	port := os.Getenv("PORT")
	if port == "" {
		port = "8084"
	}

	log.Infof("Запуск Auth Service на порту %s", port)

	// Создание роутера
	r := chi.NewRouter()

	// Маршрут аутентификации
	r.Get("/auth", authenticate)

	// Запуск сервера
	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatal("Ошибка запуска сервера: ", err)
	}
}
