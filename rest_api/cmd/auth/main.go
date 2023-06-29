package main

import (
	"github.com/spf13/viper"
	"log"
	"rest_api/pkg/handler"
	"rest_api/pkg/server"
)

func main() {

	if err := InitConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	handlers := handler.NewHandler()

	srv := new(server.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		panic(err)
		log.Fatal("error occurred while running server")
	}
}

func InitConfig() error {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath("configs")
	return viper.ReadInConfig()
}
