package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/GSVillas/pic-pay-desafio/api/handler"
	"github.com/GSVillas/pic-pay-desafio/client"
	"github.com/GSVillas/pic-pay-desafio/config"
	"github.com/GSVillas/pic-pay-desafio/config/database"
	"github.com/GSVillas/pic-pay-desafio/repository"
	"github.com/GSVillas/pic-pay-desafio/service"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"github.com/samber/do"
	"gorm.io/gorm"
)

func main() {
	config.ConfigureLogger()
	config.LoadEnvironments()

	e := echo.New()
	i := do.New()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := database.NewMysqlConnection(ctx)
	if err != nil {
		log.Fatal("Fail to connect to mysql: ", err)
	}

	redisClient, err := database.NewRedisConnection(ctx)
	if err != nil {
		log.Fatal("Fail to connect to redis: ", err)
	}

	httpClient := &http.Client{Timeout: 10 * time.Second}

	do.Provide(i, func(i *do.Injector) (*gorm.DB, error) {
		return db, nil
	})

	do.Provide(i, func(i *do.Injector) (*redis.Client, error) {
		return redisClient, nil
	})

	do.Provide(i, func(i *do.Injector) (*http.Client, error) {
		return httpClient, nil
	})

	do.Provide(i, client.NewAuthorizationService)

	do.Provide(i, handler.NewTransferHandler)
	do.Provide(i, handler.NewUserHandler)
	do.Provide(i, handler.NewWalletHandler)

	do.Provide(i, service.NewTransferService)
	do.Provide(i, service.NewUserService)
	do.Provide(i, service.NewSessionService)
	do.Provide(i, service.NewWalletService)

	do.Provide(i, repository.NewTransferRepository)
	do.Provide(i, repository.NewUserRepository)
	do.Provide(i, repository.NewSessionRepository)
	do.Provide(i, repository.NewWalletRepository)

	handler.SetupRoutes(e, i)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", config.Env.APIPort)))
}
