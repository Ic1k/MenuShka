package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Настройка логирования
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})

	// Получение порта из переменных окружения
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	log.Infof("Запуск Menu Service на порту %s", port)

	// Создание роутера
	r := chi.NewRouter()

	// Маршруты
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		log.Info("Health check")
		w.Write([]byte("Menu Service is up and running"))
	})

	// Запуск сервера
	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatal("Ошибка запуска сервера: ", err)
	}
}
