package weather

import "github.com/go-redis/cache"

type WeeklyForecast interface {
	GetForecastWeek(*GetForecastWeekRequest) (*ForecastWeekInfo, error)
	CachedWeeklyForecast(*GetForecastWeekRequest, *cache.Codec) (*ForecastWeekInfo, error)
}
