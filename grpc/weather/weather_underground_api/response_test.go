package weather_underground_api

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func Subject(t *testing.T) WeatherUndergroundApiForecastResponse {
	raw, err := ioutil.ReadFile("example.json")
	if err != nil {
		t.Error(err)
	}

	subject := WeatherUndergroundApiForecastResponse{}

	err = json.Unmarshal(raw, &subject)
	if err != nil {
		t.Error(err)
	}

	return subject
}

func TestWeatherUndergroundApiForecastResponseForecastSimpleDays(t *testing.T) {
	assert := assert.New(t)
	subject := Subject(t)

	assert.Equal(10, len(subject.Forecast.Simple.Days))
}

func TestWeatherUndergroundForecastDayWeekday(t *testing.T) {
	assert := assert.New(t)
	subject := Subject(t)

	expectations := []string{
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday",
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
	}

	for i, forecast_day := range subject.Forecast.Simple.Days {
		assert.Equal(expectations[i], forecast_day.Date.Weekday)
	}
}

func TestWeatherUndergroundForecastDayHighFahrenheit(t *testing.T) {
	assert := assert.New(t)
	subject := Subject(t)

	expectations := []string{
		"50",
		"33",
		"38",
		"39",
		"32",
		"28",
		"22",
		"14",
		"19",
		"22",
	}

	for i, forecast_day := range subject.Forecast.Simple.Days {
		assert.Equal(expectations[i], forecast_day.High.Fahrenheit)
	}
}
