package main

import (
	"auth/internal/config"
	delivery "auth/internal/delivery/http"
	"auth/internal/repository"
	"auth/pkg/redis"

	"auth/internal/service"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	lg "github.com/dmitryavdonin/gtools/logger"
)

var Version = "3.0.2"

func main() {
	cfg, err := config.InitConfig("")
	if err != nil {
		panic(fmt.Sprintf("error initializing config %s", err))
	}

	//setup logger
	logger, err := lg.New(cfg.Log.Level, cfg.App.ServiceName)
	if err != nil {
		panic(fmt.Sprintf("error initializing logger %s", err))
	}

	if cfg.Redis.Pass == " " {
		cfg.Redis.Pass = ""
	}

	//redis init
	redis, err := redis.New(cfg.Redis.Url, cfg.Redis.Pass, cfg.Session.TTL)
	if err != nil {
		logger.Fatal(fmt.Errorf("postgres connection error: %w", err))
	}

	//repository
	repository, err := repository.NewRepository(redis, cfg.UsersService.URI)
	if err != nil {
		logger.Fatal("storage initialization error: %s", err.Error())
	}

	//service
	services, err := service.NewServices(repository, cfg.Session.TTL, logger)
	if err != nil {
		logger.Fatal("services initialization error: %s", err.Error())
	}

	delivery, err := delivery.New(services, cfg.App.Port, logger, delivery.Options{})
	if err != nil {
		logger.Fatal("delivery initialization error: %s", err.Error())
	}

	//logger.Info("main(): Auth app version = " + Version)
	fmt.Println("main(): Auth app version = " + Version)

	err = delivery.Run()
	if err != nil {
		logger.Fatal("start delivery error: %s", err.Error())
	}

	//closes connections on app kill
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c
	if err := shutdown(redis, logger); err != nil {
		logger.Fatal(fmt.Errorf("failed shutdown with error: %w", err))
	}

}

func shutdown(redis *redis.Redis, logger *lg.Logger) error {
	fmt.Println("Gracefull shut down in progress...")
	redis.Close()
	logger.Info("Gracefull shutdown done!")
	return nil
}
