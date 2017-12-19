package weather

type WeatherUndergroundApiForecast struct {
  SimpleForecasts []*WeatherUndergroundSimpleForecast `json:"simpleforecast"`
}

type WeatherUndergroundSimpleForecast struct {
  ForecastDays []*WeatherUndergroundForecastDay `json:"forecastday"`
}

type WeatherUndergroundForecastDay struct {
  WeatherUndergroundForecastDate `json:"date"`
  WeatherUndergroundForecastHigh `json:"high"`
  WeatherUndergroundForecastLow  `json:"low"`
  Conditions                     string `json:"conditions`
}

type WeatherUndergroundForecastDate struct {
  Weekday string `json:"weekday"`
}

type WeatherUndergroundForecastHigh struct {
  Fahrenheit int32 `json:"fahrenheit"`
  Celsius    int32 `json:"celsius`
}

type WeatherUndergroundForecastLow struct {
  Fahrenheit int32 `json:"fahrenheit"`
  Celsius    int32 `json:"celsius`
}

var WundergroundDays_Number = map[string]DayOfWeek{
  "Sunday":    DayOfWeek_Sun,
  "Monday":    DayOfWeek_Mon,
  "Tuesday":   DayOfWeek_Tue,
  "Wednesday": DayOfWeek_Wed,
  "Thursday":  DayOfWeek_Thu,
  "Friday":    DayOfWeek_Fri,
  "Saturday":  DayOfWeek_Sat,
}