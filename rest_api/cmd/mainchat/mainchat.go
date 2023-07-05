package main

import (
	"github.com/spf13/viper"
	"log"
	"rest_api/internal/handler/main_chat"
	"rest_api/pkg/server"
)

func main() {
	if err := InitConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	handlers := main_chat.NewHandler()
	srv := new(server.Server)
	if err := srv.Run(viper.GetString("main_chat_port"), handlers.InitMainChatRoutes()); err != nil {
		log.Fatal("error occurred while running main chat service " + err.Error())
	}
}

func InitConfig() error {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath("configs")
	return viper.ReadInConfig()
}
