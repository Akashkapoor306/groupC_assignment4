package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// ExternalAPIResponse defines the structure for the external weather API response
type ExternalAPIResponse struct {
	Temperature string `json:"temperature"`
	Wind        string `json:"wind"`
	Description string `json:"description"`
	Forecast    struct {
		Day         string `json:"day"`
		Temperature string `json:"temperature"`
		Wind        string `json:"wind"`
	} `json:"forecast"`
}

// WeatherResponse defines the structure of our API response
type WeatherResponse struct {
	City        string `json:"city"`
	Temperature string `json:"temperature"`
	Weather     string `json:"weather"` // We'll use the "Description" from the external API here
}

// getWeather fetches weather data for a city using the external API
func getWeather(cityName string) (WeatherResponse, error) {
	url := fmt.Sprintf("https://goweather.herokuapp.com/weather/%s", cityName)
	resp, err := http.Get(url)
	if err != nil {
		return WeatherResponse{}, err
	}
	defer resp.Body.Close()

	var apiResponse ExternalAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return WeatherResponse{}, err
	}

	return WeatherResponse{
		City:        strings.Title(cityName),
		Temperature: apiResponse.Temperature,
		Weather:     apiResponse.Description,
	}, nil
}
