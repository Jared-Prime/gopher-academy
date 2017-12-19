package weather

import (
  "fmt"
  "github.com/go-redis/cache"
  "time"
)

type WeatherChannelApiClient struct {
  ApiKeyId string
}

func NewweatherChannelApiClient(wundergroundApiKeyId string) (*WeatherChannelApiClient, error) {
  return &WeatherChannelApiClient{ApiKeyId: wundergroundApiKeyId}, nil
}

func (backend *WeatherChannelApiClient) GetCurrentDate() string {
  t := time.Now()
  return t.Format("20060102150405")
}

func (backend *WeatherChannelApiClient) ForecastKey(prefix string) string {
  return fmt.Sprintf("%s:%s", prefix, backend.GetCurrentDate())
}

func (backend *WeatherChannelApiClient) WeeklyForecastCacheKey() string {
  return backend.ForecastKey("weekly")
}

func (backend *WeatherChannelApiClient) DailyForecastCacheKey() string {
  return backend.ForecastKey("daily")
}

func (backend *WeatherChannelApiClient) GetForecastWeek(_ *GetForecastWeekRequest) (*ForecastWeekInfo, error) {
  // request weather underground API for next 10 day forecast
  // return error if necessary
  // return formatted response as a ForecastWeekInfo
  return nil, nil
}

func (backend *WeatherChannelApiClient) CachedWeeklyForecast(request *GetForecastWeekRequest, codec *cache.Codec) (*ForecastWeekInfo, error) {
  response := &ForecastWeekInfo{}

  err := codec.Once(&cache.Item{
    Key:        backend.WeeklyForecastCacheKey(),
    Object:     response,
    Func:       func() (interface{}, error) { return backend.GetForecastWeek(request) },
    Expiration: time.Hour,
  })

  return response, err
}

func (backend *WeatherChannelApiClient) SelectDailyForecast(weekly_forecast *ForecastWeekInfo, day DayOfWeek) (*ForecastDayInfo, error) {
  for _, daily_forecast := range weekly_forecast.Forecasts {
    if daily_forecast.Day == day {
      return daily_forecast, nil
    }
  }

  return nil, fmt.Errorf("no forecast found for: %v", day)
}

func (backend *WeatherChannelApiClient) GetForecastDay(request *GetForecastDayRequest) (*ForecastDayInfo, error) {
  weekly_forecast, err := backend.GetForecastWeek(&GetForecastWeekRequest{})
  if err != nil {
    return nil, err
  }

  return backend.SelectDailyForecast(weekly_forecast, request.Day)
}

func (backend *WeatherChannelApiClient) CachedDailyForecast(request *GetForecastDayRequest, codec *cache.Codec) (*ForecastDayInfo, error) {
  response := &ForecastDayInfo{}

  err := codec.Once(&cache.Item{
    Key:        backend.DailyForecastCacheKey(),
    Object:     response,
    Func:       func() (interface{}, error) { return backend.GetForecastDay(request) },
    Expiration: time.Hour,
  })

  return response, err
}