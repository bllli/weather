package three_part_services

import (
	"context"
	"time"
)

type FutureWeatherInputDto struct {
	city string
}

type DateInfo struct {
	Date time.Time
	Desc string
}

type FutureWeatherOutputDto struct {
	DateInfos []DateInfo
}

type FutureWeather interface {
	Fetch(context context.Context, dto *FutureWeatherInputDto) (*FutureWeatherOutputDto, error)
}
