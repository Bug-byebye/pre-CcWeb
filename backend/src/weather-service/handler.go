package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	pb "CC-web/grpc/weather_rpc"
)

const weatherApi = "835ad9c03e0777cb1095985511eaf078"

type WeatherService struct {
	pb.UnimplementedWeatherServiceServer
}

func (s *WeatherService) GetWeather(ctx context.Context, req *pb.WeatherRequest) (*pb.WeatherResponse, error) {
	// 调用高德天气 API
	apiKey := weatherApi
	url := fmt.Sprintf("https://restapi.amap.com/v3/weather/weatherInfo?city=%s&extensions=%s&key=%s",
		req.City, req.Extensions, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var result map[string]interface{}
	json.Unmarshal(body, &result)

	if result["status"] != "1" {
		return nil, fmt.Errorf("Failed to get weather: %v", result["info"])
	}

	if len(result["lives"].([]interface{})) == 0 {
		return nil, fmt.Errorf("Failed to get weather-data: %v", result["info"])
	}
	// 提取数据
	live := result["lives"].([]interface{})[0].(map[string]interface{})
	return &pb.WeatherResponse{
		Status:      1,
		Info:        result["info"].(string),
		City:        live["city"].(string),
		Weather:     live["weather"].(string),
		Temperature: live["temperature"].(string),
		Humidity:    live["humidity"].(string),
		Wind:        live["winddirection"].(string),
		Reporttime:  live["reporttime"].(string),
	}, nil
}
