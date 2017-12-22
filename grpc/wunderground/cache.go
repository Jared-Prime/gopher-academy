package wunderground

import (
  "encoding/json"
  "log"
  "os"
  "time"

  "github.com/go-redis/redis"
  "github.com/go-redis/cache"
)

var redisHost string
var redisRing *redis.Ring
var redisCache *cache.Codec

func init() {
  redisHost = os.Getenv("REDIS_HOSTNAME")
  if redisHost == "" {
    log.Fatal("$REDIS_HOST required!")
  }

  redisRing = redis.NewRing(&redis.RingOptions{
    Addrs: map[string]string {
      "redis1": redisHost+":6379",
    },
  })

  redisCache = &cache.Codec{
    Redis: redisRing,
    Marshal: func(v interface{}) ([]byte, error) { return json.Marshal(v) },
    Unmarshal: func(b []byte, v interface{}) error { return json.Unmarshal(b, v) },
  }
}

func (api *WundergroundApiClient) CachedForecast() (WundergroundForecast, error) {
  forecast := new(WundergroundForecast)

  err := redisCache.Once(&cache.Item{
    Key: "wunderground:forecast",
    Object: forecast,
    Func: func() (interface{}, error) { return api.WundergroundForecast() },
    Expiration: time.Hour,
  })

  if err != nil {
    log.Print("problem accessing forecast: ", err)
  }

  return *forecast, err
}