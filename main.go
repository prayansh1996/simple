package main

import (
	"fmt"
	"net"
	"strings"

	"github.com/prayansh1996/simple/internal/service"
	"github.com/prayansh1996/simple/proto/simple"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
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
	server := grpc.NewServer()

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Failed to initialize listener: " + err.Error())
		return
	}

	simple.RegisterSimpleServiceServer(server, &service.SimpleService{})

	fmt.Println("Starting server")

	err = server.Serve(listener)
	if err != nil {
		fmt.Println("Failed to Serve: " + err.Error())
		return
	}
}
