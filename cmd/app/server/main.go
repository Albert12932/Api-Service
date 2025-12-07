// @title           API Service
// @version         1.0
// @description     Документация API-сервиса вопросов и ответов.

// @host      localhost:8080
// @BasePath  /
package main

import (
	"api-service/internal/api/routes"
	"api-service/internal/config"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	db, err := config.CreateConnect()
	if err != nil {
		panic(fmt.Sprintf("Ошибка при создании подключения базы данных: %v", err.Error()))
	}

	// получаем стандартное подключение *sql.DB для настройки пула
	sqlDB, err := db.DB()
	if err != nil {
		panic(fmt.Sprintf("Ошибка при получении стандартного соединения подключения базы данных: %v", err.Error()))
	}

	// настраиваем максимальное количество рабочих и простаивающих соединений, а также максимальное их время жизни
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)
	sqlDB.SetConnMaxLifetime(time.Hour)

	mux := http.NewServeMux()

	routes.RegisterRoutes(mux, db)

	// создаем сервер и включаем его внутри горутины.
	srv := &http.Server{Addr: ":8080", Handler: mux}
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	// создаем канал quit, чтобы при docker compose down все закрыть и корректно завершить работу приложения по "Graceful Shutdown"
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	log.Println("Сервер выключается")

	// создаем контекст чтобы использовать его в srv.Shutdown и дать время базе данных на завершение операций
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// проверяем успешное закрытие сервера и базы данных
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Ошибка при закрытии сервиса: %v", err)
	} else {
		log.Println("Сервер успешно закрыт")
	}
	if err := sqlDB.Close(); err != nil {
		log.Printf("Ошибка при закрытии базы данных: %v", err)
	} else {
		log.Println("База данных успешно закрыта")
	}

	log.Println("Сервер выключен")

}
