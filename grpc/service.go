package main

import (
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	pb "github.com/jared-prime/gopher-academy/grpc/weather"
	service "github.com/jared-prime/gopher-academy/grpc/wunderground"
)

var port string
var api_key string
var day int

func init() {
	api_key = os.Getenv("WEATHER_UNDERGROUND_API_KEY")
	if api_key == "" {
		log.Fatal("$WEATHER_UNDERGROUND_API_KEY required!")
	}

	port = os.Getenv("WEATHER_SERVICE_PORT")
	if port == "" {
		log.Print("no $WEATHER_SERVICE_PORT, using default 8000")
		port = "8000"
	}
}

func main() {
	agent, err := service.NewApiClient(api_key)
	if err != nil {
		log.Fatal(err)
	}

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal("failed to connect to specified port: ", port)
	}

	server := grpc.NewServer()

	pb.RegisterWeatherForecastServer(server, agent)

	server.Serve(listener)
}
