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

const SERVER_SHUTDOWN_TIMEOUT_SECOND = 30

func main() {
	gameRepository := repository.NewGameRepository(config.DB())
	gameService := service.NewGameService(gameRepository)
	gameHandler := handler.NewGameHandler(gameService)

	router := router.SetupRouter(gameHandler)

	db, err := config.DB().DB()
	if err != nil {
		log.Fatalf("데이터베이스 연결 실패: %v", err)
		return
	}

	defer db.Close()

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go startServer(server)

	waitForShutdown(server)
}

func startServer(server *http.Server) {
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Printf("서버 작동 중: %v", err)
	}
}

func waitForShutdown(server *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit
	log.Println("서버 종료 중...")

	// graceful shutdown 설정
	ctx, cancel := context.WithTimeout(context.Background(), SERVER_SHUTDOWN_TIMEOUT_SECOND*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("서버 종료 실패: %v", err)
	}
}
