package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func NewWeatherUndergroundClient(wundergroundApiKeyId string) (*WeatherUndergroundApiClient, error) {
	return &WeatherUndergroundApiClient{ApiKeyId: wundergroundApiKeyId}, nil
}

type WeatherUndergroundApiClient struct {
	ApiKeyId string
}

func (api WeatherUndergroundApiClient) Endpoint() string {
	return fmt.Sprintf("http://api.wunderground.com/api/%s/forecast10day/q/IL/Chicago.json", api.ApiKeyId)
}

func (api *WeatherUndergroundApiClient) ReportWeeklyForecast() (*ForecastWeekInfo, error) {
	weekly_forecast := &ForecastWeekInfo{}

	response, err := http.Get(api.Endpoint())
	if err != nil {
		return weekly_forecast, err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return weekly_forecast, err
	}

	var data WeatherUndergroundApiForecastResponse
	err = json.Unmarshal([]byte(body), &data)
	if err != nil {
		return weekly_forecast, err
	}

	for _, forecast := range data.Forecasts() {
		daily_forecast := &ForecastDayInfo{
			Day:                 DayOfWeek(forecast.DayOfWeek()),
			Temperature:         forecast.HighTemperature(),
			ForecastDescription: forecast.Conditions,
		}

		weekly_forecast.Forecasts = append(weekly_forecast.Forecasts, daily_forecast)
	}

	return weekly_forecast, nil
}
