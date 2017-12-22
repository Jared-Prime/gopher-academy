package main

import (
	"log"
	"os"

	client "github.com/jared-prime/gopher-academy/grpc/weather"
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

	res, err := agent.GetForecastDay(&client.GetForecastDayRequest{client.DayOfWeek_Sun})

	if err != nil {
		log.Fatal(err)
	}

	log.Print(res)
}
