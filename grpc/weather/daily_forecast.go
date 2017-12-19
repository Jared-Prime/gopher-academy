package weather

import "github.com/go-redis/cache"

type DailyForecast interface {
	GetForecastDay(*GetForecastDayRequest) (*ForecastDayInfo, error)
	SelectDailyForecast(*ForecastWeekInfo, *DayOfWeek) (*ForecastDayInfo, error)
	CachedDailyForecast(*GetForecastDayRequest, *cache.Codec) (*ForecastDayInfo, error)
}
