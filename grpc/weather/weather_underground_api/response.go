package weather_underground_api

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