package weather

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

func TestWeatherUndergroundApiForecastResponseForecasts(t *testing.T) {
	assert := assert.New(t)
	subject := Subject(t)

	assert.Equal(10, len(subject.Forecasts()))
}

func TestWeatherUndergroundForecastDayDayOfWeek(t *testing.T) {
	assert := assert.New(t)
	subject := Subject(t)

	for i, forecast_day := range subject.Forecasts() {
		expected := (i + 2) % 7

		assert.Equal(int32(expected), forecast_day.DayOfWeek())
	}
}

func TestWeatherUndergroundForecastDayHighTemperature(t *testing.T) {
	assert := assert.New(t)
	subject := Subject(t)

	expectations := map[int]string{
		0: "50",
		1: "33",
		2: "38",
		3: "39",
		4: "32",
		5: "28",
		6: "22",
		7: "14",
		8: "19",
		9: "22",
	}

	for i, forecast_day := range subject.Forecasts() {
		assert.Equal(expectations[i], forecast_day.HighTemperature())
	}
}
