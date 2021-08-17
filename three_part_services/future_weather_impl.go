package three_part_services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type SeniverseFutureWeather struct {
	ApiKey string
}

func (f *SeniverseFutureWeather) Fetch(context context.Context, dto *FutureWeatherInputDto) (*FutureWeatherOutputDto, error) {
	dailyUrl := fmt.Sprintf("https://api.seniverse.com/v3/weather/daily.json?key=%s&location=beijing&language=zh-Hans&unit=c&start=0&days=5", f.ApiKey)
	resp, err := http.Get(dailyUrl)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	r := DailyWeatherResp{}
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}
	if len(r.Results) < 1 {
		return nil, errors.New("")
	}
	daily := r.Results[0].Daily
	var dataInfos []DateInfo

	dataInfos = make([]DateInfo, 0)
	for _, s := range daily {
		data, err := time.Parse("2006-01-02", s.Date)
		if err != nil {
			return nil, err
		}

		precip, err := strconv.ParseFloat(s.Precip, 32)
		if err != nil {
			return nil, err
		}
		precipStr := strconv.FormatFloat(precip*100, 'f', 0, 32)
		dataInfos = append(dataInfos, DateInfo{
			Date: data,
			Desc: fmt.Sprintf("%s %s%s %s~%s %s%%", s.Date, s.TextDay, s.TextNight, s.Low, s.High, precipStr),
		})
	}

	ret := FutureWeatherOutputDto{
		DateInfos: dataInfos,
	}
	return &ret, nil
}

type DailyWeatherResp struct {
	Results []struct {
		Location struct {
			Id             string `json:"id"`
			Name           string `json:"name"`
			Country        string `json:"country"`
			Path           string `json:"path"`
			Timezone       string `json:"timezone"`
			TimezoneOffset string `json:"timezone_offset"`
		} `json:"location"`
		Daily []struct {
			Date                string `json:"date"`
			TextDay             string `json:"text_day"`
			CodeDay             string `json:"code_day"`
			TextNight           string `json:"text_night"`
			CodeNight           string `json:"code_night"`
			High                string `json:"high"`
			Low                 string `json:"low"`
			Rainfall            string `json:"rainfall"`
			Precip              string `json:"precip"`
			WindDirection       string `json:"wind_direction"`
			WindDirectionDegree string `json:"wind_direction_degree"`
			WindSpeed           string `json:"wind_speed"`
			WindScale           string `json:"wind_scale"`
			Humidity            string `json:"humidity"`
		} `json:"daily"`
		LastUpdate time.Time `json:"last_update"`
	} `json:"results"`
}
