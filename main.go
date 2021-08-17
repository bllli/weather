package main

import (
	"context"
	"log"
	"os"
	"weather/three_part_services"
)

var (
	FutureWeatherService three_part_services.FutureWeather = nil
)

func init() {
	seniverseApiKey := os.Getenv("seniverse_api_key")
	if len(seniverseApiKey) < 1 {
		log.Fatalf("env seniverse_api_key not setted!")
	}
	FutureWeatherService = &three_part_services.SeniverseFutureWeather{
		ApiKey: seniverseApiKey,
	}
}

func main() {
	resp, err := FutureWeatherService.Fetch(context.Background(), &three_part_services.FutureWeatherInputDto{})
	if err != nil {
		log.Printf("error %v", err)
		return
	}

	for _, info := range resp.DateInfos {
		println(info.Date.String(), info.Desc)
	}
}
