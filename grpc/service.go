package main

import (
	"flag"
	"log"
	"os"

	client "github.com/jared-prime/gopher-academy/grpc/weather"
	service "github.com/jared-prime/gopher-academy/grpc/wunderground"
)

var api_key string
var day int

func init() {
	api_key = os.Getenv("WEATHER_UNDERGROUND_API_KEY")
	if api_key == "" {
		log.Fatal("$WEATHER_UNDERGROUND_API_KEY required!")
	}
}

func main() {
	flag.IntVar(&day, "day", 0, "day of week (Sunday = 0, Saturday = 6, etc)")
	flag.Parse()

	agent, err := service.NewApiClient(api_key)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("getting forecast for ", day)

	res, err := agent.GetForecastDay(&client.GetForecastDayRequest{client.DayOfWeek(day)})

	if err != nil {
		log.Fatal(err)
	}

	log.Print(res)
}
