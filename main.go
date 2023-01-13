package main

import (
	"github.com/spf13/viper"
	"log"
	"login2/internal/handler"
	"login2/pkg/mysql"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error initializing configs: %s", err.Error())
	}

	mysql.OpenDB(viper.GetString("dsn"))
	defer mysql.CloseDB(mysql.DB)

	server := new(Server)
	if err := server.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		log.Fatalf("Error start server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
