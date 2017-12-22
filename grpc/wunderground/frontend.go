package wunderground

import (
  "context"
  
  pb "github.com/jared-prime/gopher-academy/grpc/weather"
)

func (backend *WundergroundApiClient) GetForecastWeek(_ context.Context, _ *pb.GetForecastWeekRequest) (*pb.ForecastWeekInfo, error) {
  weekly, err := backend.CachedForecast()
  if err != nil {
    return &pb.ForecastWeekInfo{}, err
  }

  forecast := weekly.format()

  return &forecast, err
}

func (backend *WundergroundApiClient) GetForecastDay(_ context.Context, request *pb.GetForecastDayRequest) (*pb.ForecastDayInfo, error) {
  weekly, err := backend.CachedForecast()
  if err != nil {
    return &pb.ForecastDayInfo{}, err
  }

  for _, daily := range weekly.format().Forecasts {
    if daily.DayOfWeek == request.DayOfWeek {
      return daily, nil
    }
  }

  return nil, nil
}
