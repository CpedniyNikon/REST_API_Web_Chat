package main

import (
	"github.com/spf13/viper"
	"log"
	"rest_api/internal/handler/auth"
	"rest_api/pkg/server"
)

func main() {

	if err := InitConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	handlers := auth.NewHandler()

	srv := new(server.Server)
	if err := srv.Run(viper.GetString("auth_port"), handlers.InitAuthRoutes()); err != nil {
		log.Fatal("error occurred while auth service " + err.Error())
	}
}

func InitConfig() error {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath("configs")
	return viper.ReadInConfig()
}
