package weather_underground_api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func NewClient(wundergroundApiKeyId string) (*WeatherUndergroundApiClient, error) {
	return &WeatherUndergroundApiClient{ApiKeyId: wundergroundApiKeyId}, nil
}

type WeatherUndergroundApiClient struct {
	ApiKeyId string
}

func (api WeatherUndergroundApiClient) Endpoint() string {
	return fmt.Sprintf("http://api.wunderground.com/api/%s/forecast10day/q/IL/Chicago.json", api.ApiKeyId)
}

func (api *WeatherUndergroundApiClient) ReportWeeklyForecast() (WeatherUndergroundApiForecastResponse, error) {
	var data WeatherUndergroundApiForecastResponse

	response, err := http.Get(api.Endpoint())
	if err != nil {
		return data, err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return data, err
	}

	err = json.Unmarshal([]byte(body), &data)

	return data, err
}
