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
		port = "8083"
	}

	log.Infof("Запуск Fridge Service на порту %s", port)

	// Создание роутера
	r := chi.NewRouter()

	// Маршрут для проверки работоспособности
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		log.Info("Вызван endpoint /health в Fridge Service")
		w.Write([]byte("Fridge Service is up and running"))
	})

	// Запуск сервера
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
