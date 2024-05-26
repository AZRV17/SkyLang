package app

import (
	"errors"
	"fmt"
	"github.com/AZRV17/Skylang/internal/config"
	httpHandelr "github.com/AZRV17/Skylang/internal/delivery/http"
	"github.com/AZRV17/Skylang/internal/repository"
	httpServer "github.com/AZRV17/Skylang/internal/server/http"
	"github.com/AZRV17/Skylang/internal/service"
	"github.com/AZRV17/Skylang/pkg/db/psql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// Run - запуск приложения
func Run() {
	cfg, err := config.NewConfig("internal/config/config.yaml")
	if err != nil {
		log.Fatal("error loading config: ", err)
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.Db)

	err = psql.Connect(dsn)
	if err != nil {
		log.Fatal(err)
	}

	defer psql.Close()

	r := gin.Default()

	repo := repository.NewRepository(psql.DB)

	service := service.NewService(*repo, *cfg)

	h := httpHandelr.NewHandler(*service, cfg)

	h.Init(r)

	server := httpServer.NewHttpServer(cfg, r)

	stoppedHTTP := make(chan struct{})

	go server.Shutdown(stoppedHTTP)

	log.Printf("Starting HTTP server on %s\n", cfg.HTTP.Host+":"+cfg.HTTP.Port)

	go func() {
		if err := server.Run(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP server ListenAndServe Error: %v", err)
		}
	}()

	<-stoppedHTTP

	log.Println("HTTP server stopped")
}
