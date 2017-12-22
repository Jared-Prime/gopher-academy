package wunderground

import pb "github.com/jared-prime/gopher-academy/grpc/weather"

type WundergroundForecast struct {
	Forecast struct {
		Simple struct {
			Days []struct {
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
			} `json:"forecastday"`
		} `json:"simpleforecast"`
	} `json:"forecast"`
}

func (wunday *WundergroundForecast) format() pb.ForecastWeekInfo {
	info := pb.ForecastWeekInfo{}

	for _, forecast := range wunday.Forecast.Simple.Days {
		daily := &pb.ForecastDayInfo{
			DayOfWeek: DaysOfTheWeek[forecast.Date.Weekday],
			Temperature: forecast.High.Fahrenheit,
			Unit: pb.TemperatureUnit_Fahrenheit,
			Conditions: forecast.Conditions,
		}

		info.Forecasts = append(info.Forecasts, daily)
	}

	return info
}

var DaysOfTheWeek = map[string]pb.DayOfWeek{
	"Monday": pb.DayOfWeek_Mon,
	"Tuesday": pb.DayOfWeek_Tue,
	"Wednesday": pb.DayOfWeek_Wed,
	"Thursday": pb.DayOfWeek_Thu,
	"Friday": pb.DayOfWeek_Fri,
	"Saturday": pb.DayOfWeek_Sat,
	"Sunday": pb.DayOfWeek_Sun,
}
