package weather

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/cache"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type WeatherUndergroundApiClient struct {
	ApiKeyId string
}

func (backend WeatherUndergroundApiClient) Endpoint() string {
	return fmt.Sprintf("http://api.wunderground.com/api/%s/forecast10day/q/IL/Chicago.json", backend.ApiKeyId)
}

func NewBackendApi(wundergroundApiKeyId string) (*WeatherUndergroundApiClient, error) {
	return &WeatherUndergroundApiClient{ApiKeyId: wundergroundApiKeyId}, nil
}

func (backend *WeatherUndergroundApiClient) ReportWeeklyForecast() (*ForecastWeekInfo, error) {
	weekly_forecast := &ForecastWeekInfo{}

	response, err := http.Get(backend.Endpoint())
	if err != nil {
		return weekly_forecast, err
	}


	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return weekly_forecast, err
	}

	log.Println("raw response is:", string(body))

	var data WeatherUndergroundApiForecast
	err = json.Unmarshal([]byte(body), &data)
	if err != nil {
		return weekly_forecast, err
	}

	log.Println("formatting response...")
	for _, forecast := range data.SimpleForecasts {
		for _, day := range forecast.ForecastDays {
			daily_forecast := &ForecastDayInfo{
				Day:                 WundergroundDays_Number[day.WeatherUndergroundForecastDate.Weekday],
				Temperature:         day.WeatherUndergroundForecastHigh.Celsius,
				ForecastDescription: day.Conditions,
			}

			weekly_forecast.Forecasts = append(weekly_forecast.Forecasts, daily_forecast)
		}
	}

	return weekly_forecast, nil
}

func (backend *WeatherUndergroundApiClient) GetCurrentDate() string {
	t := time.Now()
	return t.Format("20060102150405")
}

func (backend *WeatherUndergroundApiClient) ForecastKey(prefix string) string {
	return fmt.Sprintf("%s:%s", prefix, backend.GetCurrentDate())
}

func (backend *WeatherUndergroundApiClient) WeeklyForecastCacheKey() string {
	return backend.ForecastKey("weekly")
}

func (backend *WeatherUndergroundApiClient) DailyForecastCacheKey() string {
	return backend.ForecastKey("daily")
}

func (backend *WeatherUndergroundApiClient) GetForecastWeek(_ *GetForecastWeekRequest) (*ForecastWeekInfo, error) {
	return backend.ReportWeeklyForecast()
}

func (backend *WeatherUndergroundApiClient) CachedWeeklyForecast(request *GetForecastWeekRequest, codec *cache.Codec) (*ForecastWeekInfo, error) {
	response := &ForecastWeekInfo{}

	err := codec.Once(&cache.Item{
		Key:        backend.WeeklyForecastCacheKey(),
		Object:     response,
		Func:       func() (interface{}, error) { return backend.GetForecastWeek(request) },
		Expiration: time.Hour,
	})

	return response, err
}

func (backend *WeatherUndergroundApiClient) SelectDailyForecast(weekly_forecast *ForecastWeekInfo, day DayOfWeek) (*ForecastDayInfo, error) {
	for _, daily_forecast := range weekly_forecast.Forecasts {
		if daily_forecast.Day == day {
			return daily_forecast, nil
		}
	}

	return nil, fmt.Errorf("no forecast found for: %v", day)
}

func (backend *WeatherUndergroundApiClient) GetForecastDay(request *GetForecastDayRequest) (*ForecastDayInfo, error) {
	weekly_forecast, err := backend.GetForecastWeek(&GetForecastWeekRequest{})
	if err != nil {
		return nil, err
	}

	return backend.SelectDailyForecast(weekly_forecast, request.Day)
}

func (backend *WeatherUndergroundApiClient) CachedDailyForecast(request *GetForecastDayRequest, codec *cache.Codec) (*ForecastDayInfo, error) {
	response := &ForecastDayInfo{}

	err := codec.Once(&cache.Item{
		Key:        backend.DailyForecastCacheKey(),
		Object:     response,
		Func:       func() (interface{}, error) { return backend.GetForecastDay(request) },
		Expiration: time.Hour,
	})

	return response, err
}
