package weather

type WeatherUndergroundApiForecastResponse struct {
	Forecast WeatherUndergroundForecast `json:"forecast"`
}

type WeatherUndergroundForecast struct {
	Simple struct {
		Days []WeatherUndergroundForecastDay `json:"forecastday"`
	} `json:"simpleforecast"`
}

type WeatherUndergroundForecastDay struct {
	Date struct {
		Weekday string `json:"weekday"`
	} `json:"date"`

	High struct {
		Fahrenheit string `json:"fahrenheit"`
		Celsius    string `json:"celsius"`
	} `json:"high"`

	Low struct {
		Fahrenheit string `json:"fahrenheit"`
		Celsius    string `json:"celsius"`
	} `json:"low"`

	Conditions string `json:"conditions"`
}

func (resp WeatherUndergroundApiForecastResponse) Forecasts() []WeatherUndergroundForecastDay {
	return resp.Forecast.Simple.Days
}

func (forecast WeatherUndergroundForecastDay) DayOfWeek() int32 {
	return DayOfWeek_value[forecast.Date.Weekday[0:3]]
}

func (forecast WeatherUndergroundForecastDay) HighTemperature() string {
	return forecast.High.Fahrenheit
}
