package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"XNXGAMES_Game/internal/config"
	"XNXGAMES_Game/internal/handler"
	"XNXGAMES_Game/internal/repository"
	"XNXGAMES_Game/internal/router"
	"XNXGAMES_Game/internal/service"
)

func main() {
	repository := repository.NewGameRepository(config.DB())
	service := service.NewGameService(repository)
	handler := handler.NewGameHandler(service)

	router := router.SetupGameRouter(handler)

	db, err := config.DB().DB()
	if err != nil {
		return
	}

	defer db.Close()

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("서버 작동 중: %v", err)
		}
	}()

	listenServerShutDown(*server)
}

func listenServerShutDown(server http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit
	log.Println("서버 종료 중...")

	// graceful shutdown 설정
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("서버 종료 실패: %v", err)
	}
}
