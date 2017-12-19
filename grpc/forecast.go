package main

import (
	"encoding/json"
  "github.com/go-redis/cache"
  "github.com/go-redis/redis"
	"fmt"
	"log"
	"os"

	weather "github.com/jared-prime/gopher-academy/grpc/weather"
)



var (
	wundergroundApiKeyId  string
	redisHost             string
	redisCache            *cache.Codec
)

func init() {
	wundergroundApiKeyId = os.Getenv("WEATHER_UNDERGROUND_API_KEY")
	if wundergroundApiKeyId == "" {
		log.Fatal("$WEATHER_UNDERGROUND_API_KEY not set")
	}

	redisHost = os.Getenv("REDIS_HOSTNAME")
	if redisHost == "" {
		log.Fatal("$REDIS_HOSTNAME not set")
	}

	redisCache = &cache.Codec{
		Redis: redis.NewRing(&redis.RingOptions{
			Addrs: map[string]string{
				"redis1": fmt.Sprintf("%s:6379", redisHost),
			},
		}),
		Marshal:   func(v interface{}) ([]byte, error) { return json.Marshal(v) },
		Unmarshal: func(b []byte, v interface{}) error { return json.Unmarshal(b, v) },
	}
}

func main() {
  weatherUnderground, err := weather.NewBackendApi(wundergroundApiKeyId)
  if err != nil {
    log.Fatal(err)
  }

  log.Println("getting info from weather underground...")
  response, err := weatherUnderground.GetForecastWeek(&weather.GetForecastWeekRequest{})
  if err != nil {
  	log.Fatal(err)
  }
  log.Println("got info!")

  log.Print(*response)
}
