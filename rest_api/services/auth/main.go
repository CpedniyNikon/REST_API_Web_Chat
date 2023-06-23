package main

import (
	"github.com/spf13/viper"
	"log"
	"rest_api/auth_libs/server"
	"rest_api/pkg/handler"
)

func main() {

	if err := InitConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	handlers := handler.NewHandler()

	srv := new(server.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatal("error occurred while running server")
	}
}

func InitConfig() error {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath("configs")
	viper.AddConfigPath("../../configs")
	return viper.ReadInConfig()
}
