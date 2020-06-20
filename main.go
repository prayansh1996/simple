package main

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func main() {
	setupConfig()
	setupGrpcServer()
}

func setupConfig() {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")

	viper.SetEnvPrefix("simple.service")

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
}

func setupGrpcServer() {
	port := viper.GetString("grpc.port")
	fmt.Println(port)
}
