package main

import (
	"context"
	server "effective_mobile_rest"
	"effective_mobile_rest/internal/api/v1/handler"
	"effective_mobile_rest/internal/api/v1/repository"
	"effective_mobile_rest/internal/api/v1/repository/postgres"
	"effective_mobile_rest/internal/api/v1/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

// @title Effective Mobile REST API
// @version 1.0
// @description This is the RESTful API for Effective Mobile.
// @host localhost:8081
// @BasePath /api/v1
func main() {
	if err := initConfig(); err != nil {
		log.Fatal("Error: init config", err)
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error: loading env variables", err)
	}

	db, err := postgres.NewDB(postgres.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		log.Fatal("Error: failed to init db connection ", err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.New(services)

	serv := new(server.Server)
	go func() {
		if err := serv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			log.Fatal("Error: failed to start server on port: %s", viper.GetString("port"), err)
		}
	}()

	slog.Info("Start server")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := serv.ShutDown(context.Background()); err != nil {
		log.Fatal("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		log.Fatal("error occured on close db connection: %s", err.Error())
	}
	slog.Info("Server shutting down")

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
