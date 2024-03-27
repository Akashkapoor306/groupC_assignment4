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

//Function created by Nikhil

func cityHandler(w http.ResponseWriter, r *http.Request) {
	var weather WeatherResponse
	var cityName string
	var err error

	if r.Method == "GET" {
		cityName = r.URL.Query().Get("name")
	} else if r.Method == "POST" {
		var requestData struct{ Name string }
		if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		cityName = requestData.Name
	}

	weather, err = getWeather(cityName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weather)
}
