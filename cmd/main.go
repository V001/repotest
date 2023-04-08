package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"repotest/config"
	"repotest/logger"
	"repotest/service"
	"repotest/storage"
	"repotest/transport/http"
	"repotest/transport/http/handler"
)

// @title testProject
// @version 0.1
// @description описание сервиса
// @termsOfService http://swagger.io/terms/

// @contact.name Автор
// @contact.email Имейл

// @host localhost:9
// @BasePath /api/v1

// User + Auth  CRUD +Swagger
// Book  CRUD +Swagger
// Transactions CRUD +Swagger
func main() {
	log.Fatalln(fmt.Sprintf("Service shut down: %s", run()))
}

func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	//cLogger := logger.NewConsoleLogger(logger.INFO, logger.JSON)
	gracefullyShutdown(cancel)
	conf, err := config.New()
	if err != nil {
		return err
	}

	storage, err := storage.New(ctx, conf)
	if err != nil {
		log.Fatal(err.Error())
	}
	_ = logger.Logger(ctx)
	svc, svcErr := service.NewManager(storage)
	if svcErr != nil {
		return svcErr
	}
	h := handler.NewManager(conf, svc)
	//jwtAuth := middleware.NewJWTAuth(cfg, srv.User)

	HTTPServer := http.NewServer(conf, h)

	return HTTPServer.StartHTTPServer(ctx)
}

func gracefullyShutdown(c context.CancelFunc) {
	osC := make(chan os.Signal, 1)
	signal.Notify(osC, os.Interrupt)
	go func() {
		log.Print(<-osC)
		c()
	}()
}
