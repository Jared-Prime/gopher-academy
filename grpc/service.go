package main

import (
  "log"
  "os"
  
  client "./weather"
  service "./wunderground"
)

func main(){
  agent, err := service.NewApiClient(os.Getenv("WEATHER_UNDERGROUND_API_KEY"))
  if err != nil {
    log.Fatal(err)
  }

  agent.GetForecastDay(&client.GetForecastDayRequest{ client.DayOfWeek_Sun })
}