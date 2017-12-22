package main

import (
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	pb "github.com/jared-prime/gopher-academy/grpc/weather"
	service "github.com/jared-prime/gopher-academy/grpc/wunderground"
)

var api_key string

func init() {
	api_key = os.Getenv("WEATHER_UNDERGROUND_API_KEY")
	if api_key == "" {
		log.Fatal("$WEATHER_UNDERGROUND_API_KEY required!")
	}
}

func main() {
	agent, err := service.NewApiClient(api_key)
	if err != nil {
		log.Fatal(err)
	}

	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal("failed to connect to specified port: ", port)
	}

	server := grpc.NewServer()

	pb.RegisterWeatherForecastServer(server, agent)

	server.Serve(listener)
}
