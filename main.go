package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/notifications/db"
	"github.com/notifications/handler"
)


func setUpConfig() {
	viper.SetEnvPrefix("UN")
	viper.AutomaticEnv()
	viper.SetDefault("deploy_env", "local")
	viper.SetConfigType("yml")
	viper.SetConfigName("application-" + viper.GetString("deploy_env"))
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}


func setUp() {
	setUpConfig()
	db.Initialize()
	handler.Initialize()
}

func main() {
	setUp()
	handler.Run()
}