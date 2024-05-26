package httpServer

import (
	"context"
	config "github.com/AZRV17/Skylang/internal/config"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type HttpServer struct {
	httpServer *http.Server
}

func NewHttpServer(cfg *config.Config, handler http.Handler) *HttpServer {
	return &HttpServer{
		httpServer: &http.Server{
			Addr:    cfg.HTTP.Host + ":" + cfg.HTTP.Port,
			Handler: handler,
		},
	}
}

// Функция для запуска сервера
func (s *HttpServer) Run() error {
	return s.httpServer.ListenAndServe()
}

// Функция для остановки сервера
func (s *HttpServer) Shutdown(stopped chan struct{}) {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-sigint
	log.Println("got interruption signal")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.httpServer.Shutdown(ctx); err != nil {
		log.Printf("HTTP Server Shutdown Error: %v", err)
	}
	close(stopped)
}
