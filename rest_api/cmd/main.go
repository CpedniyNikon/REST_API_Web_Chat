package main

import (
	"github.com/spf13/viper"
	"log"
	"rest_api"
	"rest_api/pkg/handler"
	"rest_api/pkg/repository"
	"rest_api/pkg/service"
)




func main() {

	if err := InitConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)


	srv := new(rest_api.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatal("error occurred while running server")
	}
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}