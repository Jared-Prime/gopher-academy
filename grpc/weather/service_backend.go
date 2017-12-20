package weather

import (
	"fmt"
	"github.com/go-redis/cache"
	"time"
)

type ServiceBackend struct {
	Client BackendClient
}

type BackendClient interface {
	ReportWeeklyForecast() (*ForecastWeekInfo, error)
}

func NewServiceBackend(client BackendClient) *ServiceBackend {
	return &ServiceBackend{client}
}

func currentDate() string {
	t := time.Now()
	return t.Format("20060102150405")
}

func (backend *ServiceBackend) GetForecastWeek(_ *GetForecastWeekRequest) (*ForecastWeekInfo, error) {
	return backend.Client.ReportWeeklyForecast()
}

func (backend *ServiceBackend) CachedWeeklyForecast(request *GetForecastWeekRequest, codec *cache.Codec) (*ForecastWeekInfo, error) {
	response := &ForecastWeekInfo{}

	err := codec.Once(&cache.Item{
		Key:        fmt.Sprintf("weekly:%s", currentDate()),
		Object:     response,
		Func:       func() (interface{}, error) { return backend.GetForecastWeek(request) },
		Expiration: time.Hour,
	})

	return response, err
}

func (backend *ServiceBackend) SelectDailyForecast(weekly_forecast *ForecastWeekInfo, day DayOfWeek) (*ForecastDayInfo, error) {
	for _, daily_forecast := range weekly_forecast.Forecasts {
		if daily_forecast.Day == day {
			return daily_forecast, nil
		}
	}

	return nil, fmt.Errorf("no forecast found for: %v", day)
}

func (backend *ServiceBackend) GetForecastDay(request *GetForecastDayRequest) (*ForecastDayInfo, error) {
	weekly_forecast, err := backend.GetForecastWeek(&GetForecastWeekRequest{})
	if err != nil {
		return nil, err
	}

	return backend.SelectDailyForecast(weekly_forecast, request.Day)
}

func (backend *ServiceBackend) CachedDailyForecast(request *GetForecastDayRequest, codec *cache.Codec) (*ForecastDayInfo, error) {
	response := &ForecastDayInfo{}

	err := codec.Once(&cache.Item{
		Key:        fmt.Sprintf("daily:%s", currentDate()),
		Object:     response,
		Func:       func() (interface{}, error) { return backend.GetForecastDay(request) },
		Expiration: time.Hour,
	})

	return response, err
}
